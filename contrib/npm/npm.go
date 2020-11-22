package npm

import (
	"github.com/clearblade/cbtest"
	"github.com/stretchr/testify/require"
)

// Use creates a new *NPMTransaction.
// Panics on failure.
func Use(t cbtest.T, dir string) *Transaction {
	t.Helper()
	tran, err := NewTransaction(t, dir, true)
	require.NoError(t, err)
	return tran
}

// UseE creates a new *NPMTransaction.
// Returns error on failure.
func UseE(t cbtest.T, dir string) (*Transaction, error) {
	t.Helper()
	return NewTransaction(t, dir, false)
}
