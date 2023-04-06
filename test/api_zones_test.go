/*
Meraki Dashboard API

Testing ZonesApiService

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

func Test_meraki_ZonesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ZonesApiService GetDeviceCameraAnalyticsZoneHistory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string
		var zoneId string

		resp, httpRes, err := apiClient.ZonesApi.GetDeviceCameraAnalyticsZoneHistory(context.Background(), serial, zoneId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ZonesApiService GetDeviceCameraAnalyticsZones", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.ZonesApi.GetDeviceCameraAnalyticsZones(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}