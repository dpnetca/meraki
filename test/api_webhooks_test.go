/*
Meraki Dashboard API

Testing WebhooksApiService

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

func Test_meraki_WebhooksApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test WebhooksApiService CreateNetworkWebhooksHttpServer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.WebhooksApi.CreateNetworkWebhooksHttpServer(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhooksApiService CreateNetworkWebhooksPayloadTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.WebhooksApi.CreateNetworkWebhooksPayloadTemplate(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhooksApiService CreateNetworkWebhooksWebhookTest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.WebhooksApi.CreateNetworkWebhooksWebhookTest(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhooksApiService DeleteNetworkWebhooksHttpServer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var httpServerId string

		httpRes, err := apiClient.WebhooksApi.DeleteNetworkWebhooksHttpServer(context.Background(), networkId, httpServerId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhooksApiService DeleteNetworkWebhooksPayloadTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var payloadTemplateId string

		httpRes, err := apiClient.WebhooksApi.DeleteNetworkWebhooksPayloadTemplate(context.Background(), networkId, payloadTemplateId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhooksApiService GetNetworkWebhooksHttpServer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var httpServerId string

		resp, httpRes, err := apiClient.WebhooksApi.GetNetworkWebhooksHttpServer(context.Background(), networkId, httpServerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhooksApiService GetNetworkWebhooksHttpServers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.WebhooksApi.GetNetworkWebhooksHttpServers(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhooksApiService GetNetworkWebhooksPayloadTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var payloadTemplateId string

		resp, httpRes, err := apiClient.WebhooksApi.GetNetworkWebhooksPayloadTemplate(context.Background(), networkId, payloadTemplateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhooksApiService GetNetworkWebhooksPayloadTemplates", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.WebhooksApi.GetNetworkWebhooksPayloadTemplates(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhooksApiService GetNetworkWebhooksWebhookTest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var webhookTestId string

		resp, httpRes, err := apiClient.WebhooksApi.GetNetworkWebhooksWebhookTest(context.Background(), networkId, webhookTestId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhooksApiService GetOrganizationWebhooksAlertTypes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.WebhooksApi.GetOrganizationWebhooksAlertTypes(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhooksApiService GetOrganizationWebhooksLogs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.WebhooksApi.GetOrganizationWebhooksLogs(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhooksApiService UpdateNetworkWebhooksHttpServer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var httpServerId string

		resp, httpRes, err := apiClient.WebhooksApi.UpdateNetworkWebhooksHttpServer(context.Background(), networkId, httpServerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhooksApiService UpdateNetworkWebhooksPayloadTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var payloadTemplateId string

		resp, httpRes, err := apiClient.WebhooksApi.UpdateNetworkWebhooksPayloadTemplate(context.Background(), networkId, payloadTemplateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
