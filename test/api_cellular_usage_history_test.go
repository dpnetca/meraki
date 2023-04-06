/*
Meraki Dashboard API

Testing CellularUsageHistoryApiService

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

func Test_meraki_CellularUsageHistoryApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CellularUsageHistoryApiService GetNetworkSmDeviceCellularUsageHistory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var deviceId string

		resp, httpRes, err := apiClient.CellularUsageHistoryApi.GetNetworkSmDeviceCellularUsageHistory(context.Background(), networkId, deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
