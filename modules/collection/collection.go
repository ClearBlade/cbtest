package collection

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	cb "github.com/clearblade/Go-SDK"
)

// IDByName gets a collection ID from a given name.
// Panics on failure.
func IDByName(t *testing.T, devClient *cb.DevClient, collectionName string) string {
	id, err := IDByNameE(t, devClient, collectionName)
	require.NoError(t, err)
	return id
}

// IDByNameE gets a collection ID from a given name.
// Returns error on failure.
func IDByNameE(t *testing.T, devClient *cb.DevClient, collectionID string) (string, error) {
	return "", fmt.Errorf("not implemented yet")
}
