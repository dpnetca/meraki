/*
Meraki Dashboard API

Testing WirelessProfilesApiService

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

func Test_meraki_WirelessProfilesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test WirelessProfilesApiService CreateNetworkCameraWirelessProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.WirelessProfilesApi.CreateNetworkCameraWirelessProfile(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WirelessProfilesApiService DeleteNetworkCameraWirelessProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var wirelessProfileId string

		httpRes, err := apiClient.WirelessProfilesApi.DeleteNetworkCameraWirelessProfile(context.Background(), networkId, wirelessProfileId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WirelessProfilesApiService GetDeviceCameraWirelessProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.WirelessProfilesApi.GetDeviceCameraWirelessProfiles(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WirelessProfilesApiService GetNetworkCameraWirelessProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var wirelessProfileId string

		resp, httpRes, err := apiClient.WirelessProfilesApi.GetNetworkCameraWirelessProfile(context.Background(), networkId, wirelessProfileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WirelessProfilesApiService GetNetworkCameraWirelessProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.WirelessProfilesApi.GetNetworkCameraWirelessProfiles(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WirelessProfilesApiService UpdateDeviceCameraWirelessProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.WirelessProfilesApi.UpdateDeviceCameraWirelessProfiles(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WirelessProfilesApiService UpdateNetworkCameraWirelessProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var wirelessProfileId string

		resp, httpRes, err := apiClient.WirelessProfilesApi.UpdateNetworkCameraWirelessProfile(context.Background(), networkId, wirelessProfileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
