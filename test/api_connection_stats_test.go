/*
Meraki Dashboard API

Testing ConnectionStatsApiService

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

func Test_meraki_ConnectionStatsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ConnectionStatsApiService GetDeviceWirelessConnectionStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.ConnectionStatsApi.GetDeviceWirelessConnectionStats(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectionStatsApiService GetNetworkWirelessClientConnectionStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var clientId string

		resp, httpRes, err := apiClient.ConnectionStatsApi.GetNetworkWirelessClientConnectionStats(context.Background(), networkId, clientId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectionStatsApiService GetNetworkWirelessClientsConnectionStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.ConnectionStatsApi.GetNetworkWirelessClientsConnectionStats(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectionStatsApiService GetNetworkWirelessConnectionStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.ConnectionStatsApi.GetNetworkWirelessConnectionStats(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectionStatsApiService GetNetworkWirelessDevicesConnectionStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.ConnectionStatsApi.GetNetworkWirelessDevicesConnectionStats(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
