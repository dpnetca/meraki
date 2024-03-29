/*
Meraki Dashboard API

Testing OptInsApiService

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

func Test_meraki_OptInsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OptInsApiService CreateOrganizationEarlyAccessFeaturesOptIn", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.OptInsApi.CreateOrganizationEarlyAccessFeaturesOptIn(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OptInsApiService DeleteOrganizationEarlyAccessFeaturesOptIn", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var optInId string

		httpRes, err := apiClient.OptInsApi.DeleteOrganizationEarlyAccessFeaturesOptIn(context.Background(), organizationId, optInId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OptInsApiService GetOrganizationEarlyAccessFeaturesOptIn", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var optInId string

		resp, httpRes, err := apiClient.OptInsApi.GetOrganizationEarlyAccessFeaturesOptIn(context.Background(), organizationId, optInId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OptInsApiService GetOrganizationEarlyAccessFeaturesOptIns", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.OptInsApi.GetOrganizationEarlyAccessFeaturesOptIns(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OptInsApiService UpdateOrganizationEarlyAccessFeaturesOptIn", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var optInId string

		resp, httpRes, err := apiClient.OptInsApi.UpdateOrganizationEarlyAccessFeaturesOptIn(context.Background(), organizationId, optInId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
