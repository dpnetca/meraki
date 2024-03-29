/*
Meraki Dashboard API

Testing CameraApiService

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

func Test_meraki_CameraApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CameraApiService CreateNetworkCameraQualityRetentionProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.CameraApi.CreateNetworkCameraQualityRetentionProfile(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CameraApiService CreateNetworkCameraWirelessProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.CameraApi.CreateNetworkCameraWirelessProfile(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CameraApiService CreateOrganizationCameraCustomAnalyticsArtifact", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.CameraApi.CreateOrganizationCameraCustomAnalyticsArtifact(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CameraApiService DeleteNetworkCameraQualityRetentionProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var qualityRetentionProfileId string

		httpRes, err := apiClient.CameraApi.DeleteNetworkCameraQualityRetentionProfile(context.Background(), networkId, qualityRetentionProfileId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CameraApiService DeleteNetworkCameraWirelessProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var wirelessProfileId string

		httpRes, err := apiClient.CameraApi.DeleteNetworkCameraWirelessProfile(context.Background(), networkId, wirelessProfileId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CameraApiService DeleteOrganizationCameraCustomAnalyticsArtifact", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var artifactId string

		httpRes, err := apiClient.CameraApi.DeleteOrganizationCameraCustomAnalyticsArtifact(context.Background(), organizationId, artifactId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CameraApiService GenerateDeviceCameraSnapshot", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.CameraApi.GenerateDeviceCameraSnapshot(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CameraApiService GetDeviceCameraAnalyticsLive", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.CameraApi.GetDeviceCameraAnalyticsLive(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CameraApiService GetDeviceCameraAnalyticsOverview", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.CameraApi.GetDeviceCameraAnalyticsOverview(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CameraApiService GetDeviceCameraAnalyticsRecent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.CameraApi.GetDeviceCameraAnalyticsRecent(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CameraApiService GetDeviceCameraAnalyticsZoneHistory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string
		var zoneId string

		resp, httpRes, err := apiClient.CameraApi.GetDeviceCameraAnalyticsZoneHistory(context.Background(), serial, zoneId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CameraApiService GetDeviceCameraAnalyticsZones", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.CameraApi.GetDeviceCameraAnalyticsZones(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CameraApiService GetDeviceCameraCustomAnalytics", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.CameraApi.GetDeviceCameraCustomAnalytics(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CameraApiService GetDeviceCameraQualityAndRetention", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.CameraApi.GetDeviceCameraQualityAndRetention(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CameraApiService GetDeviceCameraSense", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.CameraApi.GetDeviceCameraSense(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CameraApiService GetDeviceCameraSenseObjectDetectionModels", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.CameraApi.GetDeviceCameraSenseObjectDetectionModels(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CameraApiService GetDeviceCameraVideoLink", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.CameraApi.GetDeviceCameraVideoLink(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CameraApiService GetDeviceCameraVideoSettings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.CameraApi.GetDeviceCameraVideoSettings(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CameraApiService GetDeviceCameraWirelessProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.CameraApi.GetDeviceCameraWirelessProfiles(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CameraApiService GetNetworkCameraQualityRetentionProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var qualityRetentionProfileId string

		resp, httpRes, err := apiClient.CameraApi.GetNetworkCameraQualityRetentionProfile(context.Background(), networkId, qualityRetentionProfileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CameraApiService GetNetworkCameraQualityRetentionProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.CameraApi.GetNetworkCameraQualityRetentionProfiles(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CameraApiService GetNetworkCameraSchedules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.CameraApi.GetNetworkCameraSchedules(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CameraApiService GetNetworkCameraWirelessProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var wirelessProfileId string

		resp, httpRes, err := apiClient.CameraApi.GetNetworkCameraWirelessProfile(context.Background(), networkId, wirelessProfileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CameraApiService GetNetworkCameraWirelessProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.CameraApi.GetNetworkCameraWirelessProfiles(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CameraApiService GetOrganizationCameraCustomAnalyticsArtifact", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var artifactId string

		resp, httpRes, err := apiClient.CameraApi.GetOrganizationCameraCustomAnalyticsArtifact(context.Background(), organizationId, artifactId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CameraApiService GetOrganizationCameraCustomAnalyticsArtifacts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.CameraApi.GetOrganizationCameraCustomAnalyticsArtifacts(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CameraApiService GetOrganizationCameraOnboardingStatuses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.CameraApi.GetOrganizationCameraOnboardingStatuses(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CameraApiService UpdateDeviceCameraCustomAnalytics", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.CameraApi.UpdateDeviceCameraCustomAnalytics(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CameraApiService UpdateDeviceCameraQualityAndRetention", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.CameraApi.UpdateDeviceCameraQualityAndRetention(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CameraApiService UpdateDeviceCameraSense", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.CameraApi.UpdateDeviceCameraSense(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CameraApiService UpdateDeviceCameraVideoSettings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.CameraApi.UpdateDeviceCameraVideoSettings(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CameraApiService UpdateDeviceCameraWirelessProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.CameraApi.UpdateDeviceCameraWirelessProfiles(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CameraApiService UpdateNetworkCameraQualityRetentionProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var qualityRetentionProfileId string

		resp, httpRes, err := apiClient.CameraApi.UpdateNetworkCameraQualityRetentionProfile(context.Background(), networkId, qualityRetentionProfileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CameraApiService UpdateNetworkCameraWirelessProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var wirelessProfileId string

		resp, httpRes, err := apiClient.CameraApi.UpdateNetworkCameraWirelessProfile(context.Background(), networkId, wirelessProfileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CameraApiService UpdateOrganizationCameraOnboardingStatuses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.CameraApi.UpdateOrganizationCameraOnboardingStatuses(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
