package collection

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	cb "github.com/clearblade/Go-SDK"
	"github.com/clearblade/cbtest/modules/system"
)

// IDByName gets a collection ID from a given name.
// Panics on failure.
func IDByName(t *testing.T, system *system.EphemeralSystem, devClient *cb.DevClient, collectionName string) string {
	id, err := IDByNameE(t, system, devClient, collectionName)
	require.NoError(t, err)
	return id
}

// IDByNameE gets a collection ID from a given name.
// Returns error on failure.
func IDByNameE(t *testing.T, system *system.EphemeralSystem, devClient *cb.DevClient, collectionName string) (string, error) {

	allCollections, err := devClient.GetAllCollections(system.SystemKey())
	if err != nil {
		return "", err
	}

	for _, rawColl := range allCollections {

		coll, ok := rawColl.(map[string]interface{})
		if !ok {
			return "", fmt.Errorf("could not parse collections")
		}

		collID, ok := coll["collectionID"].(string)
		if !ok {
			return "", fmt.Errorf("could not parse collection ID")
		}

		collName, ok := coll["name"].(string)
		if !ok {
			return "", fmt.Errorf("could not parse collection name")
		}

		if collName == collectionName {
			return collID, nil
		}
	}

	return "", fmt.Errorf("collection ID for name \"%s\" not found", collectionName)
}
