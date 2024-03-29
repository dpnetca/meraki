/*
Meraki Dashboard API

Testing PoliciesApiService

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

func Test_meraki_PoliciesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PoliciesApiService CreateOrganizationAdaptivePolicyPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.PoliciesApi.CreateOrganizationAdaptivePolicyPolicy(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoliciesApiService DeleteOrganizationAdaptivePolicyPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var id string

		httpRes, err := apiClient.PoliciesApi.DeleteOrganizationAdaptivePolicyPolicy(context.Background(), organizationId, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoliciesApiService GetNetworkPoliciesByClient", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.PoliciesApi.GetNetworkPoliciesByClient(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoliciesApiService GetOrganizationAdaptivePolicyPolicies", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.PoliciesApi.GetOrganizationAdaptivePolicyPolicies(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoliciesApiService GetOrganizationAdaptivePolicyPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var id string

		resp, httpRes, err := apiClient.PoliciesApi.GetOrganizationAdaptivePolicyPolicy(context.Background(), organizationId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoliciesApiService UpdateOrganizationAdaptivePolicyPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var id string

		resp, httpRes, err := apiClient.PoliciesApi.UpdateOrganizationAdaptivePolicyPolicy(context.Background(), organizationId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
