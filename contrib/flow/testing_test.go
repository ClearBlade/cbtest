package flow

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var fourSpaces = "    "
var newline = "\n"

func TestT_FailingChildFailsParent(t *testing.T) {

	output := strings.Builder{}
	parent := newTWithOutput("parent", &output)
	child := newChildT(parent, "child")

	child.Fail()
	assert.True(t, child.Failed())
	assert.True(t, parent.Failed())
}

func TestT_FailingSiblingFailsParent(t *testing.T) {

	output := strings.Builder{}
	parent := newTWithOutput("parent", &output)
	child := newChildT(parent, "child")
	sibling := newSiblingT(child, "sibling")

	sibling.Fail()
	assert.True(t, sibling.Failed())
	assert.False(t, child.Failed())
	assert.True(t, parent.Failed())
}

func TestT_Helper(t *testing.T) {

	output := strings.Builder{}
	flowT := newTWithOutput("root", &output)

	logOne := func(t *T) {
		t.Helper()
		t.Log("one")
	}

	logTwo := func(t *T) {
		t.Helper()
		logOne(t)
		t.Log("two")
	}

	logThree := func(t *T) {
		t.Helper()
		logTwo(t)
		t.Log("three")
	}

	logThree(flowT)

	expected := strings.Join([]string{
		fourSpaces, "testing_test.go:59 [root]: one", newline,
		fourSpaces, "testing_test.go:59 [root]: two", newline,
		fourSpaces, "testing_test.go:59 [root]: three", newline,
	}, "")

	assert.Equal(t, expected, output.String())
}

func TestT_LogNested(t *testing.T) {

	output := strings.Builder{}
	parent := newTWithOutput("parent", &output)
	child := newChildT(parent, "child")
	sibling := newSiblingT(child, "sibling")

	parent.Log("parent")
	child.Log("child")
	sibling.Log("sibling")

	expected := strings.Join([]string{
		fourSpaces, "testing_test.go:77 [parent]: parent", newline,
		fourSpaces, "testing_test.go:78 [parent/child]: child", newline,
		fourSpaces, "testing_test.go:79 [parent/sibling]: sibling", newline,
	}, "")

	assert.Equal(t, expected, output.String())
}
