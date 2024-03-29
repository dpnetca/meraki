/*
Meraki Dashboard API

Testing ActionBatchesApiService

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

func Test_meraki_ActionBatchesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ActionBatchesApiService CreateOrganizationActionBatch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.ActionBatchesApi.CreateOrganizationActionBatch(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ActionBatchesApiService DeleteOrganizationActionBatch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var actionBatchId string

		httpRes, err := apiClient.ActionBatchesApi.DeleteOrganizationActionBatch(context.Background(), organizationId, actionBatchId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ActionBatchesApiService GetOrganizationActionBatch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var actionBatchId string

		resp, httpRes, err := apiClient.ActionBatchesApi.GetOrganizationActionBatch(context.Background(), organizationId, actionBatchId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ActionBatchesApiService GetOrganizationActionBatches", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.ActionBatchesApi.GetOrganizationActionBatches(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ActionBatchesApiService UpdateOrganizationActionBatch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var actionBatchId string

		resp, httpRes, err := apiClient.ActionBatchesApi.UpdateOrganizationActionBatch(context.Background(), organizationId, actionBatchId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
