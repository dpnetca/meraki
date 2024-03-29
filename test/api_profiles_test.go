/*
Meraki Dashboard API

Testing ProfilesApiService

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

func Test_meraki_ProfilesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProfilesApiService CreateNetworkSensorAlertsProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.ProfilesApi.CreateNetworkSensorAlertsProfile(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProfilesApiService CreateOrganizationAlertsProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.ProfilesApi.CreateOrganizationAlertsProfile(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProfilesApiService DeleteNetworkSensorAlertsProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var id string

		httpRes, err := apiClient.ProfilesApi.DeleteNetworkSensorAlertsProfile(context.Background(), networkId, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProfilesApiService DeleteOrganizationAlertsProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var alertConfigId string

		httpRes, err := apiClient.ProfilesApi.DeleteOrganizationAlertsProfile(context.Background(), organizationId, alertConfigId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProfilesApiService GetNetworkSensorAlertsProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var id string

		resp, httpRes, err := apiClient.ProfilesApi.GetNetworkSensorAlertsProfile(context.Background(), networkId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProfilesApiService GetNetworkSensorAlertsProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.ProfilesApi.GetNetworkSensorAlertsProfiles(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProfilesApiService GetNetworkSmProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.ProfilesApi.GetNetworkSmProfiles(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProfilesApiService GetOrganizationAlertsProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.ProfilesApi.GetOrganizationAlertsProfiles(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProfilesApiService GetOrganizationConfigTemplateSwitchProfilePort", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var configTemplateId string
		var profileId string
		var portId string

		resp, httpRes, err := apiClient.ProfilesApi.GetOrganizationConfigTemplateSwitchProfilePort(context.Background(), organizationId, configTemplateId, profileId, portId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProfilesApiService GetOrganizationConfigTemplateSwitchProfilePorts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var configTemplateId string
		var profileId string

		resp, httpRes, err := apiClient.ProfilesApi.GetOrganizationConfigTemplateSwitchProfilePorts(context.Background(), organizationId, configTemplateId, profileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProfilesApiService GetOrganizationConfigTemplateSwitchProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var configTemplateId string

		resp, httpRes, err := apiClient.ProfilesApi.GetOrganizationConfigTemplateSwitchProfiles(context.Background(), organizationId, configTemplateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProfilesApiService UpdateNetworkSensorAlertsProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var id string

		resp, httpRes, err := apiClient.ProfilesApi.UpdateNetworkSensorAlertsProfile(context.Background(), networkId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProfilesApiService UpdateOrganizationAlertsProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var alertConfigId string

		resp, httpRes, err := apiClient.ProfilesApi.UpdateOrganizationAlertsProfile(context.Background(), organizationId, alertConfigId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProfilesApiService UpdateOrganizationConfigTemplateSwitchProfilePort", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var configTemplateId string
		var profileId string
		var portId string

		resp, httpRes, err := apiClient.ProfilesApi.UpdateOrganizationConfigTemplateSwitchProfilePort(context.Background(), organizationId, configTemplateId, profileId, portId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
