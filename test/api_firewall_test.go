/*
Meraki Dashboard API

Testing FirewallApiService

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

func Test_meraki_FirewallApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FirewallApiService GetNetworkApplianceFirewallCellularFirewallRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.FirewallApi.GetNetworkApplianceFirewallCellularFirewallRules(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallApiService GetNetworkApplianceFirewallFirewalledService", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var service string

		resp, httpRes, err := apiClient.FirewallApi.GetNetworkApplianceFirewallFirewalledService(context.Background(), networkId, service).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallApiService GetNetworkApplianceFirewallFirewalledServices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.FirewallApi.GetNetworkApplianceFirewallFirewalledServices(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallApiService GetNetworkApplianceFirewallInboundCellularFirewallRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.FirewallApi.GetNetworkApplianceFirewallInboundCellularFirewallRules(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallApiService GetNetworkApplianceFirewallInboundFirewallRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.FirewallApi.GetNetworkApplianceFirewallInboundFirewallRules(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallApiService GetNetworkApplianceFirewallL3FirewallRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.FirewallApi.GetNetworkApplianceFirewallL3FirewallRules(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallApiService GetNetworkApplianceFirewallL7FirewallRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.FirewallApi.GetNetworkApplianceFirewallL7FirewallRules(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallApiService GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.FirewallApi.GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallApiService GetNetworkApplianceFirewallOneToManyNatRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.FirewallApi.GetNetworkApplianceFirewallOneToManyNatRules(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallApiService GetNetworkApplianceFirewallOneToOneNatRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.FirewallApi.GetNetworkApplianceFirewallOneToOneNatRules(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallApiService GetNetworkApplianceFirewallPortForwardingRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.FirewallApi.GetNetworkApplianceFirewallPortForwardingRules(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallApiService GetNetworkApplianceFirewallSettings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.FirewallApi.GetNetworkApplianceFirewallSettings(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallApiService GetNetworkWirelessSsidFirewallL3FirewallRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string

		resp, httpRes, err := apiClient.FirewallApi.GetNetworkWirelessSsidFirewallL3FirewallRules(context.Background(), networkId, number).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallApiService GetNetworkWirelessSsidFirewallL7FirewallRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string

		resp, httpRes, err := apiClient.FirewallApi.GetNetworkWirelessSsidFirewallL7FirewallRules(context.Background(), networkId, number).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallApiService UpdateNetworkApplianceFirewallCellularFirewallRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.FirewallApi.UpdateNetworkApplianceFirewallCellularFirewallRules(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallApiService UpdateNetworkApplianceFirewallFirewalledService", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var service string

		resp, httpRes, err := apiClient.FirewallApi.UpdateNetworkApplianceFirewallFirewalledService(context.Background(), networkId, service).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallApiService UpdateNetworkApplianceFirewallInboundCellularFirewallRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.FirewallApi.UpdateNetworkApplianceFirewallInboundCellularFirewallRules(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallApiService UpdateNetworkApplianceFirewallInboundFirewallRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.FirewallApi.UpdateNetworkApplianceFirewallInboundFirewallRules(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallApiService UpdateNetworkApplianceFirewallL3FirewallRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.FirewallApi.UpdateNetworkApplianceFirewallL3FirewallRules(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallApiService UpdateNetworkApplianceFirewallL7FirewallRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.FirewallApi.UpdateNetworkApplianceFirewallL7FirewallRules(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallApiService UpdateNetworkApplianceFirewallOneToManyNatRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.FirewallApi.UpdateNetworkApplianceFirewallOneToManyNatRules(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallApiService UpdateNetworkApplianceFirewallOneToOneNatRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.FirewallApi.UpdateNetworkApplianceFirewallOneToOneNatRules(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallApiService UpdateNetworkApplianceFirewallPortForwardingRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.FirewallApi.UpdateNetworkApplianceFirewallPortForwardingRules(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallApiService UpdateNetworkApplianceFirewallSettings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.FirewallApi.UpdateNetworkApplianceFirewallSettings(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallApiService UpdateNetworkWirelessSsidFirewallL3FirewallRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string

		resp, httpRes, err := apiClient.FirewallApi.UpdateNetworkWirelessSsidFirewallL3FirewallRules(context.Background(), networkId, number).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallApiService UpdateNetworkWirelessSsidFirewallL7FirewallRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string

		resp, httpRes, err := apiClient.FirewallApi.UpdateNetworkWirelessSsidFirewallL7FirewallRules(context.Background(), networkId, number).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
