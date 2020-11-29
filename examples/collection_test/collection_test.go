// Package collection_test showcases a test that checks a collection.
package collection_test

import (
	"testing"

	"github.com/clearblade/cbtest/modules/auth"
	"github.com/clearblade/cbtest/modules/collection"
	"github.com/clearblade/cbtest/modules/service"
	"github.com/clearblade/cbtest/modules/should"
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
	s := system.UseOrImport(t, "./extra")

	// close the system after the test
	defer system.Finish(t, s)

	// obtain developer client from the ephemeral system
	devClient := auth.LoginAsDev(t, s)

	// calls all the required operations from table
	for _, tt := range table {
		payload := map[string]interface{}{"lhs": tt.lhs, "rhs": tt.rhs}
		resp, err := devClient.CallService(s.SystemKey(), AdderService, payload, false)
		should.NoError(t, err)
		should.Expect(t, resp, service.ResponseSuccess(tt.want))
	}

	// fetch all collection data
	collID := collection.IDByName(t, s, ResultsCollection)

	// fetch data total for the collection
	data, err := devClient.GetDataTotal(collID, nil)
	should.NoError(t, err)

	// assert that the total respose contains expected value
	should.Expect(t, data, collection.HaveTotal(len(table)))
}
