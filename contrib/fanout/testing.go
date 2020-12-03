package fanout

import (
	"testing"

	"github.com/clearblade/cbtest"
)

//go:generate mockery --name tWithRunParallel --inpackage --testonly

// tWithRunParallel specifies the usual cbtest.T interface with extra methods for running
// the subtest in parallel.
type tWithRunParallel interface {
	cbtest.T
	Run(name string, f func(*testing.T)) bool
	Parallel()
}
