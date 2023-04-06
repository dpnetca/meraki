/*
Meraki Dashboard API

Testing DscpTaggingOptionsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package meraki

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/dpnetca/meraki"
)

func Test_meraki_DscpTaggingOptionsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DscpTaggingOptionsApiService GetNetworkTrafficShapingDscpTaggingOptions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.DscpTaggingOptionsApi.GetNetworkTrafficShapingDscpTaggingOptions(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}