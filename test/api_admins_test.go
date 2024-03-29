/*
Meraki Dashboard API

Testing AdminsApiService

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

func Test_meraki_AdminsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AdminsApiService CreateOrganizationAdmin", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.AdminsApi.CreateOrganizationAdmin(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminsApiService DeleteOrganizationAdmin", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var adminId string

		httpRes, err := apiClient.AdminsApi.DeleteOrganizationAdmin(context.Background(), organizationId, adminId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminsApiService GetOrganizationAdmins", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.AdminsApi.GetOrganizationAdmins(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminsApiService UpdateOrganizationAdmin", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var adminId string

		resp, httpRes, err := apiClient.AdminsApi.UpdateOrganizationAdmin(context.Background(), organizationId, adminId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
