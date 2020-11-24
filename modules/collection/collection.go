// Package collection contains functions and assertions for interacting with
// collections.
package collection

import (
	"fmt"

	"github.com/stretchr/testify/require"

	"github.com/clearblade/cbtest"
	"github.com/clearblade/cbtest/provider"
)

// IDByName gets a collection ID from a given name.
// Panics on failure.
func IDByName(t cbtest.T, provider provider.ConfigAndClient, collectionName string) string {
	id, err := IDByNameE(t, provider, collectionName)
	require.NoError(t, err)
	return id
}

// IDByNameE gets a collection ID from a given name.
// Returns error on failure.
func IDByNameE(t cbtest.T, provider provider.ConfigAndClient, collectionName string) (string, error) {

	config := provider.Config(t)
	devClient, err := provider.ClientE(t)
	if err != nil {
		return "", err
	}

	allCollections, err := devClient.GetAllCollections(config.SystemKey)
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

// Total returns the total number of rows in the collection.
// Panics on failure.
func Total(t cbtest.T, provider provider.ConfigAndClient, collectionID string) int {
	count, err := TotalE(t, provider, collectionID)
	require.NoError(t, err)
	return count
}

// TotalE returns the total number of rows in the collection.
// Returns error on failure.
func TotalE(t cbtest.T, provider provider.ConfigAndClient, collectionID string) (int, error) {

	devClient, err := provider.ClientE(t)
	if err != nil {
		return 0, err
	}

	data, err := devClient.GetDataTotal(collectionID, nil)
	if err != nil {
		return 0, nil
	}

	count, ok := data["count"].(float64) // NOTE: response comes as float64
	if !ok {
		return 0, fmt.Errorf("could not get collection row count")
	}

	return int(count), nil
}
