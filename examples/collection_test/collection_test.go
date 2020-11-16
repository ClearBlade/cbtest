// Package adder_test showcases a test that passes parameters to a code service.
package collection_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/clearblade/cbtest"
	"github.com/clearblade/cbtest/modules/auth"
	"github.com/clearblade/cbtest/modules/collection"
	"github.com/clearblade/cbtest/modules/service"
	"github.com/clearblade/cbtest/modules/system"
)

const (
	AdderService      = "adder"
	ResultsCollection = "results"
)

func TestCollection(t *testing.T) {

	table := []struct {
		lhs  float64
		rhs  float64
		want float64
	}{
		{0, 0, 0},
		{1, 1, 2},
		{3, 4, 7},
	}

	// import into new system
	s := system.ImportSystem(t, "./extra")

	// destroy the system after the test
	defer cbtest.Destroy(t, s)

	// obtain developer client from the ephemeral system
	devClient := auth.LoginAsDev(t, s)

	// calls all the required operations from table
	for _, tt := range table {
		payload := map[string]interface{}{"lhs": tt.lhs, "rhs": tt.rhs}
		resp, err := devClient.CallService(s.SystemKey(), AdderService, payload, false)
		require.NoError(t, err)
		service.AssertResponseEqual(t, tt.want, resp)
	}

	// fetch all collection data
	collID := collection.IDByName(t, s, devClient, ResultsCollection)
	data, err := devClient.GetData(collID, nil)
	require.NoError(t, err)

	// assert on the collection data
	collection.AssertHasLength(t, 3, data)
}
