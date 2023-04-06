/*
Meraki Dashboard API

Testing PayloadTemplatesApiService

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

func Test_meraki_PayloadTemplatesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PayloadTemplatesApiService CreateNetworkWebhooksPayloadTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.PayloadTemplatesApi.CreateNetworkWebhooksPayloadTemplate(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PayloadTemplatesApiService DeleteNetworkWebhooksPayloadTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var payloadTemplateId string

		httpRes, err := apiClient.PayloadTemplatesApi.DeleteNetworkWebhooksPayloadTemplate(context.Background(), networkId, payloadTemplateId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PayloadTemplatesApiService GetNetworkWebhooksPayloadTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var payloadTemplateId string

		resp, httpRes, err := apiClient.PayloadTemplatesApi.GetNetworkWebhooksPayloadTemplate(context.Background(), networkId, payloadTemplateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PayloadTemplatesApiService GetNetworkWebhooksPayloadTemplates", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.PayloadTemplatesApi.GetNetworkWebhooksPayloadTemplates(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PayloadTemplatesApiService UpdateNetworkWebhooksPayloadTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var payloadTemplateId string

		resp, httpRes, err := apiClient.PayloadTemplatesApi.UpdateNetworkWebhooksPayloadTemplate(context.Background(), networkId, payloadTemplateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
