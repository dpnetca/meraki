/*
Meraki Dashboard API

Testing EarlyAccessApiService

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

func Test_meraki_EarlyAccessApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EarlyAccessApiService CreateOrganizationEarlyAccessFeaturesOptIn", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.EarlyAccessApi.CreateOrganizationEarlyAccessFeaturesOptIn(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EarlyAccessApiService DeleteOrganizationEarlyAccessFeaturesOptIn", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var optInId string

		httpRes, err := apiClient.EarlyAccessApi.DeleteOrganizationEarlyAccessFeaturesOptIn(context.Background(), organizationId, optInId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EarlyAccessApiService GetOrganizationEarlyAccessFeatures", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.EarlyAccessApi.GetOrganizationEarlyAccessFeatures(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EarlyAccessApiService GetOrganizationEarlyAccessFeaturesOptIn", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var optInId string

		resp, httpRes, err := apiClient.EarlyAccessApi.GetOrganizationEarlyAccessFeaturesOptIn(context.Background(), organizationId, optInId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EarlyAccessApiService GetOrganizationEarlyAccessFeaturesOptIns", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.EarlyAccessApi.GetOrganizationEarlyAccessFeaturesOptIns(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EarlyAccessApiService UpdateOrganizationEarlyAccessFeaturesOptIn", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var optInId string

		resp, httpRes, err := apiClient.EarlyAccessApi.UpdateOrganizationEarlyAccessFeaturesOptIn(context.Background(), organizationId, optInId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
