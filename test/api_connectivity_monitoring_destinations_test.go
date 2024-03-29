/*
Meraki Dashboard API

Testing ConnectivityMonitoringDestinationsApiService

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

func Test_meraki_ConnectivityMonitoringDestinationsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ConnectivityMonitoringDestinationsApiService GetNetworkApplianceConnectivityMonitoringDestinations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.ConnectivityMonitoringDestinationsApi.GetNetworkApplianceConnectivityMonitoringDestinations(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectivityMonitoringDestinationsApiService GetNetworkCellularGatewayConnectivityMonitoringDestinations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.ConnectivityMonitoringDestinationsApi.GetNetworkCellularGatewayConnectivityMonitoringDestinations(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectivityMonitoringDestinationsApiService UpdateNetworkApplianceConnectivityMonitoringDestinations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.ConnectivityMonitoringDestinationsApi.UpdateNetworkApplianceConnectivityMonitoringDestinations(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectivityMonitoringDestinationsApiService UpdateNetworkCellularGatewayConnectivityMonitoringDestinations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.ConnectivityMonitoringDestinationsApi.UpdateNetworkCellularGatewayConnectivityMonitoringDestinations(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
