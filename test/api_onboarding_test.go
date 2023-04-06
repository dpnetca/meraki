/*
Meraki Dashboard API

Testing OnboardingApiService

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

func Test_meraki_OnboardingApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OnboardingApiService CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.OnboardingApi.CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OnboardingApiService CreateOrganizationInventoryOnboardingCloudMonitoringImport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.OnboardingApi.CreateOrganizationInventoryOnboardingCloudMonitoringImport(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OnboardingApiService CreateOrganizationInventoryOnboardingCloudMonitoringPrepare", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.OnboardingApi.CreateOrganizationInventoryOnboardingCloudMonitoringPrepare(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OnboardingApiService GetOrganizationCameraOnboardingStatuses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.OnboardingApi.GetOrganizationCameraOnboardingStatuses(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OnboardingApiService GetOrganizationInventoryOnboardingCloudMonitoringImports", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.OnboardingApi.GetOrganizationInventoryOnboardingCloudMonitoringImports(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OnboardingApiService GetOrganizationInventoryOnboardingCloudMonitoringNetworks", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.OnboardingApi.GetOrganizationInventoryOnboardingCloudMonitoringNetworks(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OnboardingApiService UpdateOrganizationCameraOnboardingStatuses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.OnboardingApi.UpdateOrganizationCameraOnboardingStatuses(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}