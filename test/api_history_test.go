/*
Meraki Dashboard API

Testing HistoryApiService

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

func Test_meraki_HistoryApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test HistoryApiService GetDeviceCameraAnalyticsZoneHistory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string
		var zoneId string

		resp, httpRes, err := apiClient.HistoryApi.GetDeviceCameraAnalyticsZoneHistory(context.Background(), serial, zoneId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HistoryApiService GetNetworkAlertsHistory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.HistoryApi.GetNetworkAlertsHistory(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HistoryApiService GetOrganizationSensorReadingsHistory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.HistoryApi.GetOrganizationSensorReadingsHistory(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
