package flow

import (
	"fmt"
	"io"
	"path"
	"regexp"
	"runtime"
	"strings"
	"sync"
)

// testingIndent for compatibility with the testing package identation.
var testingIndent = "    "

// innerIndent for indenting any subsequent lines.
var innerIndent = "    "

// spaceRegex is a regular expression for whitespace.
var spaceRegex = regexp.MustCompile(`\s+`)

// T serves as flow controller and test handler (akin to testing.T). A new instance
// is passed to each of the workers in a given flow.
type T struct {
	parent      *T
	name        string
	depth       int
	helperNames map[string]struct{}
	mu          *sync.RWMutex
	output      io.Writer
	outputmu    *sync.RWMutex
	failed      bool
}

// newFlowT returns a new *flow.T without a parent.
func newFlowT(name string, output io.Writer) *T {
	return &T{
		name:        name,
		helperNames: make(map[string]struct{}),
		mu:          &sync.RWMutex{},
		output:      output,
		outputmu:    &sync.RWMutex{},
	}
}

// newChildFlowT returns a new *flow.T with the given parent.
func newChildFlowT(parent *T, name string) *T {
	return &T{
		parent:      parent,
		name:        name,
		depth:       parent.depth + 1,
		helperNames: make(map[string]struct{}),
		mu:          &sync.RWMutex{},
		output:      parent.output,
		outputmu:    parent.outputmu,
	}
}

// Error outputs the given args and marks the flow as failed.
func (t *T) Error(args ...interface{}) {
	t.Fail()
	t.log(fmt.Sprint(args...))
}

// Errorf outputs the given format, args, and marks the flow as failed.
func (t *T) Errorf(format string, args ...interface{}) {
	t.Fail()
	t.log(fmt.Sprintf(format, args...))
}

// Fail fails the current flow but continues execution.
func (t *T) Fail() {

	if t.parent != nil {
		t.parent.Fail()
	}

	t.mu.Lock()
	t.failed = true
	t.mu.Unlock()
}

// FailNow fails the current flow, and halts execution.
func (t *T) FailNow() {
	t.Fail()
	runtime.Goexit()
}

// Failed returns whenever the flow has failed.
func (t *T) Failed() bool {
	return t.failed
}

// Helper marks the calling function as a helper function, meaning its name
// will be skipped when printing messages.
func (t *T) Helper() {
	t.mu.Lock()
	defer t.mu.Unlock()

	caller, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("testflow: no caller found")
	}

	t.helperNames[pcToName(caller)] = struct{}{}
}

// Name returns the name of the current flow.
func (t *T) Name() string {
	parentPrefix := ""
	if t.parent != nil {
		parentPrefix = t.parent.Name() + "/"
	}

	formatted := fmt.Sprintf("%s%s", parentPrefix, t.name)
	return spaceRegex.ReplaceAllString(formatted, "_")
}

// Log outputs the given args.
func (t *T) Log(args ...interface{}) {
	t.log(fmt.Sprint(args...))
}

// Logf outputs the given format and args.
func (t *T) Logf(format string, args ...interface{}) {
	t.log(fmt.Sprintf(format, args...))
}

// log outputs the given string.
func (t *T) log(s string) {
	t.logDepth(s, 2) // skip log + public function
}

// logDepth outputs the given string skipping the given amount of callers.
func (t *T) logDepth(s string, depth int) {

	frame := t.frameSkip(depth + 1) // skip logDepth
	file := path.Base(frame.File)
	line := frame.Line

	var indent strings.Builder
	indent.WriteString("\n")
	indent.WriteString(testingIndent)
	indent.WriteString(innerIndent)

	formatted := strings.ReplaceAll(s, "\n", indent.String())

	t.outputmu.Lock()
	fmt.Fprintf(t.output, testingIndent)
	fmt.Fprintf(t.output, "%s:%d [%s]: %s\n", file, line, t.Name(), formatted)
	t.outputmu.Unlock()
}

// frameSkip searches, starting from skip frames, for the first caller not marked
// as a helper function and returns its frame.
func (t *T) frameSkip(skip int) runtime.Frame {

	t.mu.Lock()
	defer t.mu.Unlock()

	pc := make([]uintptr, 50)
	n := runtime.Callers(skip+2, pc) // skip Callers + frameSkip
	if n == 0 {
		panic("testflow: no callers found")
	}

	frames := runtime.CallersFrames(pc)

	var firstFrame, frame runtime.Frame
	var more bool

	frame, more = frames.Next()
	for more {

		if frame.PC == 0 {
			firstFrame = frame
		}

		if _, ok := t.helperNames[frame.Function]; !ok {
			return frame
		}

		frame, more = frames.Next()
	}

	return firstFrame
}

// pcToName returns the function name that contains the given program counter.
func pcToName(pc uintptr) string {
	pcs := []uintptr{pc}
	frames := runtime.CallersFrames(pcs)
	frame, _ := frames.Next()
	return frame.Function
}
