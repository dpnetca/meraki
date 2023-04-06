/*
Meraki Dashboard API

Testing SplashApiService

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

func Test_meraki_SplashApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SplashApiService GetNetworkWirelessSsidSplashSettings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string

		resp, httpRes, err := apiClient.SplashApi.GetNetworkWirelessSsidSplashSettings(context.Background(), networkId, number).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SplashApiService UpdateNetworkWirelessSsidSplashSettings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string

		resp, httpRes, err := apiClient.SplashApi.UpdateNetworkWirelessSsidSplashSettings(context.Background(), networkId, number).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
