/*
Meraki Dashboard API

Testing CustomPerformanceClassesApiService

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

func Test_meraki_CustomPerformanceClassesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CustomPerformanceClassesApiService CreateNetworkApplianceTrafficShapingCustomPerformanceClass", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.CustomPerformanceClassesApi.CreateNetworkApplianceTrafficShapingCustomPerformanceClass(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomPerformanceClassesApiService DeleteNetworkApplianceTrafficShapingCustomPerformanceClass", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var customPerformanceClassId string

		httpRes, err := apiClient.CustomPerformanceClassesApi.DeleteNetworkApplianceTrafficShapingCustomPerformanceClass(context.Background(), networkId, customPerformanceClassId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomPerformanceClassesApiService GetNetworkApplianceTrafficShapingCustomPerformanceClass", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var customPerformanceClassId string

		resp, httpRes, err := apiClient.CustomPerformanceClassesApi.GetNetworkApplianceTrafficShapingCustomPerformanceClass(context.Background(), networkId, customPerformanceClassId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomPerformanceClassesApiService GetNetworkApplianceTrafficShapingCustomPerformanceClasses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.CustomPerformanceClassesApi.GetNetworkApplianceTrafficShapingCustomPerformanceClasses(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomPerformanceClassesApiService UpdateNetworkApplianceTrafficShapingCustomPerformanceClass", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var customPerformanceClassId string

		resp, httpRes, err := apiClient.CustomPerformanceClassesApi.UpdateNetworkApplianceTrafficShapingCustomPerformanceClass(context.Background(), networkId, customPerformanceClassId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
