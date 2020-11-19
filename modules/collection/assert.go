package collection

import (
	"github.com/clearblade/cbtest"
	"github.com/clearblade/cbtest/provider"
	"github.com/stretchr/testify/assert"
)

// AssertHasTotal returns true if the given collection has the desired
// number of rows.
func AssertHasTotal(t cbtest.T, provider provider.ConfigAndClient, collectionID string, total int) bool {
	t.Helper()
	actual, err := TotalE(t, provider, collectionID)
	assert.NoError(t, err)
	return assert.Equal(t, total, actual)
}
