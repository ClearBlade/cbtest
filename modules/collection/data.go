package collection

import (
	cb "github.com/clearblade/Go-SDK"
	"github.com/clearblade/cbtest"
	"github.com/clearblade/cbtest/provider"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/require"
)

// DataResponse represents a response from a collection get data call.
type DataResponse struct {
	Page        int                      `json:"CURRENTPAGE" mapstructure:"CURRENTPAGE"`
	PrevPageURL string                   `json:"PREVPAGEURL" mapstructure:"PREVPAGEURL"`
	NextPageURL string                   `json:"NEXTPAGEURL" mapstructure:"NEXTPAGEURL"`
	Items       []map[string]interface{} `json:"DATA" mapstructure:"DATA"`
	TotalItems  int                      `json:"TOTAL" mapstructure:"TOTAL"`
}

// GetData wraps around ClearBlade Go SDK GetData and provides the response as a structure.
// Panics on failure.
func GetData(t cbtest.T, provider provider.ConfigAndClient, collectionID string, query *cb.Query) *DataResponse {
	resp, err := GetDataE(t, provider, collectionID, query)
	require.NoError(t, err)
	return resp
}

// GetDataE wraps around ClearBlade Go SDK GetData and provides the response as a structure.
// Returns error on failure.
func GetDataE(t cbtest.T, provider provider.ConfigAndClient, collectionID string, query *cb.Query) (*DataResponse, error) {

	devClient, err := provider.ClientE(t)
	if err != nil {
		return nil, err
	}

	rawdata, err := devClient.GetData(collectionID, query)
	if err != nil {
		return nil, err
	}

	resp := DataResponse{}
	err = mapstructure.Decode(rawdata, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
