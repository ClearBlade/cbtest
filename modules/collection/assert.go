package collection

import (
	"github.com/clearblade/cbtest"
	"github.com/clearblade/cbtest/provider"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// AssertHasLength returns true if the given collection has the desired
// number of rows.
func AssertHasLength(t cbtest.T, provider provider.ConfigAndClient, collectionID string, length int) bool {
	t.Helper()

	devClient := provider.Client(t)

	data, err := devClient.GetDataTotal(collectionID, nil)
	require.NoError(t, err)
	count, ok := data["count"]
	require.True(t, ok, "could not get collection count")

	// NOTE: cast to float64 because response `count` is a float64
	return assert.Equal(t, float64(length), count)
}
