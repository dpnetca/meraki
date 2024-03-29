/*
Meraki Dashboard API

Testing DeviceProfilesApiService

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

func Test_meraki_DeviceProfilesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DeviceProfilesApiService GetNetworkSmDeviceDeviceProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var deviceId string

		resp, httpRes, err := apiClient.DeviceProfilesApi.GetNetworkSmDeviceDeviceProfiles(context.Background(), networkId, deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeviceProfilesApiService GetNetworkSmUserDeviceProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var userId string

		resp, httpRes, err := apiClient.DeviceProfilesApi.GetNetworkSmUserDeviceProfiles(context.Background(), networkId, userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
