/*
Meraki Dashboard API

Testing RendezvousPointsApiService

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

func Test_meraki_RendezvousPointsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RendezvousPointsApiService CreateNetworkSwitchRoutingMulticastRendezvousPoint", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.RendezvousPointsApi.CreateNetworkSwitchRoutingMulticastRendezvousPoint(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RendezvousPointsApiService DeleteNetworkSwitchRoutingMulticastRendezvousPoint", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var rendezvousPointId string

		httpRes, err := apiClient.RendezvousPointsApi.DeleteNetworkSwitchRoutingMulticastRendezvousPoint(context.Background(), networkId, rendezvousPointId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RendezvousPointsApiService GetNetworkSwitchRoutingMulticastRendezvousPoint", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var rendezvousPointId string

		resp, httpRes, err := apiClient.RendezvousPointsApi.GetNetworkSwitchRoutingMulticastRendezvousPoint(context.Background(), networkId, rendezvousPointId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RendezvousPointsApiService GetNetworkSwitchRoutingMulticastRendezvousPoints", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.RendezvousPointsApi.GetNetworkSwitchRoutingMulticastRendezvousPoints(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RendezvousPointsApiService UpdateNetworkSwitchRoutingMulticastRendezvousPoint", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var rendezvousPointId string

		resp, httpRes, err := apiClient.RendezvousPointsApi.UpdateNetworkSwitchRoutingMulticastRendezvousPoint(context.Background(), networkId, rendezvousPointId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}