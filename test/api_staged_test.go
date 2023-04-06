/*
Meraki Dashboard API

Testing StagedApiService

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

func Test_meraki_StagedApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test StagedApiService CreateNetworkFirmwareUpgradesStagedEvent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.StagedApi.CreateNetworkFirmwareUpgradesStagedEvent(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StagedApiService CreateNetworkFirmwareUpgradesStagedGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.StagedApi.CreateNetworkFirmwareUpgradesStagedGroup(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StagedApiService DeferNetworkFirmwareUpgradesStagedEvents", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.StagedApi.DeferNetworkFirmwareUpgradesStagedEvents(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StagedApiService DeleteNetworkFirmwareUpgradesStagedGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var groupId string

		httpRes, err := apiClient.StagedApi.DeleteNetworkFirmwareUpgradesStagedGroup(context.Background(), networkId, groupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StagedApiService GetNetworkFirmwareUpgradesStagedEvents", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.StagedApi.GetNetworkFirmwareUpgradesStagedEvents(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StagedApiService GetNetworkFirmwareUpgradesStagedGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var groupId string

		resp, httpRes, err := apiClient.StagedApi.GetNetworkFirmwareUpgradesStagedGroup(context.Background(), networkId, groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StagedApiService GetNetworkFirmwareUpgradesStagedGroups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.StagedApi.GetNetworkFirmwareUpgradesStagedGroups(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StagedApiService GetNetworkFirmwareUpgradesStagedStages", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.StagedApi.GetNetworkFirmwareUpgradesStagedStages(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StagedApiService RollbacksNetworkFirmwareUpgradesStagedEvents", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.StagedApi.RollbacksNetworkFirmwareUpgradesStagedEvents(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StagedApiService UpdateNetworkFirmwareUpgradesStagedEvents", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.StagedApi.UpdateNetworkFirmwareUpgradesStagedEvents(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StagedApiService UpdateNetworkFirmwareUpgradesStagedGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var groupId string

		resp, httpRes, err := apiClient.StagedApi.UpdateNetworkFirmwareUpgradesStagedGroup(context.Background(), networkId, groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StagedApiService UpdateNetworkFirmwareUpgradesStagedStages", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.StagedApi.UpdateNetworkFirmwareUpgradesStagedStages(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
