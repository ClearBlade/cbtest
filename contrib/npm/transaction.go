package npm

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/clearblade/cbtest"
	"github.com/clearblade/cbtest/internal/fsutil"
)

// checkNPM returns error if the given path does not contain an NPM project.
func checkNPM(path string) error {

	packageJSONPath := filepath.Join(path, "package.json")

	_, err := os.Stat(packageJSONPath)
	if err != nil {
		return fmt.Errorf("not a npm project: %s", err)
	}

	return nil
}

// Transaction represents a transaction on a NPM project.
type Transaction struct {
	t        cbtest.T
	dir      string
	failfast bool
	output   string
	err      error
}

// NewTransaction returns a new *Transaction instance.
func NewTransaction(t cbtest.T, dir string, failfast bool) (*Transaction, error) {

	absdir, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}

	if !fsutil.IsDir(absdir) {
		return nil, fmt.Errorf("not a directory: %s", absdir)
	}

	err = checkNPM(absdir)
	if err != nil {
		return nil, err
	}

	return &Transaction{t: t, dir: absdir, failfast: failfast}, nil
}

// process processes the output of a command an adjusts the internal transaction
// state accordingly.
func (tx *Transaction) process(output string, err error) {
	tx.t.Helper()

	tx.output = output

	if err == nil {
		return
	}

	tx.err = err

	if tx.failfast {
		tx.t.Error(err)
		tx.t.Error(output)
		tx.t.FailNow()
	}
}

// Dir returns the directory for this transaction.
func (tx *Transaction) Dir() string {
	return tx.dir
}

// Failfast returns true if the transaction is set to fail fast.
func (tx *Transaction) Failfast() bool {
	return tx.failfast
}

// Output returns the output of the last command in the transaction.
func (tx *Transaction) Output() string {
	return tx.output
}

// Error returns the error of the last command in the transaction.
func (tx *Transaction) Error() error {
	return tx.err
}

// Install runs npm install.
func (tx *Transaction) Install() *Transaction {
	tx.t.Helper()

	if tx.err != nil {
		return tx
	}

	cmd := exec.Command("npm", "install")
	cmd.Dir = tx.dir

	tx.t.Log("Running npm install...")
	output, err := cmd.CombinedOutput()
	tx.process(string(output), err)
	return tx
}

// Run runs the given npm script.
func (tx *Transaction) Run(name string, arg ...string) *Transaction {
	tx.t.Helper()

	if tx.err != nil {
		return tx
	}

	cmdArg := make([]string, 0, len(arg)+2)
	cmdArg = append(cmdArg, "run")
	cmdArg = append(cmdArg, name)
	cmdArg = append(cmdArg, arg...)

	cmd := exec.Command("npm", cmdArg...)
	cmd.Dir = tx.dir

	tx.t.Logf("Running npm command `%s`...", cmd.String())
	output, err := cmd.CombinedOutput()
	tx.process(string(output), err)
	return tx
}
