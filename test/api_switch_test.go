/*
Meraki Dashboard API

Testing SwitchApiService

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

func Test_meraki_SwitchApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SwitchApiService AddNetworkSwitchStack", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var switchStackId string

		resp, httpRes, err := apiClient.SwitchApi.AddNetworkSwitchStack(context.Background(), networkId, switchStackId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService CloneOrganizationSwitchDevices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.SwitchApi.CloneOrganizationSwitchDevices(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService CreateDeviceSwitchRoutingInterface", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.SwitchApi.CreateDeviceSwitchRoutingInterface(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService CreateDeviceSwitchRoutingStaticRoute", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.SwitchApi.CreateDeviceSwitchRoutingStaticRoute(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService CreateNetworkSwitchAccessPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SwitchApi.CreateNetworkSwitchAccessPolicy(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SwitchApi.CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService CreateNetworkSwitchLinkAggregation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SwitchApi.CreateNetworkSwitchLinkAggregation(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService CreateNetworkSwitchPortSchedule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SwitchApi.CreateNetworkSwitchPortSchedule(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService CreateNetworkSwitchQosRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SwitchApi.CreateNetworkSwitchQosRule(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService CreateNetworkSwitchRoutingMulticastRendezvousPoint", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SwitchApi.CreateNetworkSwitchRoutingMulticastRendezvousPoint(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService CreateNetworkSwitchStack", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SwitchApi.CreateNetworkSwitchStack(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService CreateNetworkSwitchStackRoutingInterface", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var switchStackId string

		resp, httpRes, err := apiClient.SwitchApi.CreateNetworkSwitchStackRoutingInterface(context.Background(), networkId, switchStackId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService CreateNetworkSwitchStackRoutingStaticRoute", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var switchStackId string

		resp, httpRes, err := apiClient.SwitchApi.CreateNetworkSwitchStackRoutingStaticRoute(context.Background(), networkId, switchStackId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService CycleDeviceSwitchPorts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.SwitchApi.CycleDeviceSwitchPorts(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService DeleteDeviceSwitchRoutingInterface", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string
		var interfaceId string

		httpRes, err := apiClient.SwitchApi.DeleteDeviceSwitchRoutingInterface(context.Background(), serial, interfaceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService DeleteDeviceSwitchRoutingStaticRoute", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string
		var staticRouteId string

		httpRes, err := apiClient.SwitchApi.DeleteDeviceSwitchRoutingStaticRoute(context.Background(), serial, staticRouteId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService DeleteNetworkSwitchAccessPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var accessPolicyNumber string

		httpRes, err := apiClient.SwitchApi.DeleteNetworkSwitchAccessPolicy(context.Background(), networkId, accessPolicyNumber).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var trustedServerId string

		httpRes, err := apiClient.SwitchApi.DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(context.Background(), networkId, trustedServerId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService DeleteNetworkSwitchLinkAggregation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var linkAggregationId string

		httpRes, err := apiClient.SwitchApi.DeleteNetworkSwitchLinkAggregation(context.Background(), networkId, linkAggregationId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService DeleteNetworkSwitchPortSchedule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var portScheduleId string

		httpRes, err := apiClient.SwitchApi.DeleteNetworkSwitchPortSchedule(context.Background(), networkId, portScheduleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService DeleteNetworkSwitchQosRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var qosRuleId string

		httpRes, err := apiClient.SwitchApi.DeleteNetworkSwitchQosRule(context.Background(), networkId, qosRuleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService DeleteNetworkSwitchRoutingMulticastRendezvousPoint", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var rendezvousPointId string

		httpRes, err := apiClient.SwitchApi.DeleteNetworkSwitchRoutingMulticastRendezvousPoint(context.Background(), networkId, rendezvousPointId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService DeleteNetworkSwitchStack", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var switchStackId string

		httpRes, err := apiClient.SwitchApi.DeleteNetworkSwitchStack(context.Background(), networkId, switchStackId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService DeleteNetworkSwitchStackRoutingInterface", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var switchStackId string
		var interfaceId string

		httpRes, err := apiClient.SwitchApi.DeleteNetworkSwitchStackRoutingInterface(context.Background(), networkId, switchStackId, interfaceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService DeleteNetworkSwitchStackRoutingStaticRoute", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var switchStackId string
		var staticRouteId string

		httpRes, err := apiClient.SwitchApi.DeleteNetworkSwitchStackRoutingStaticRoute(context.Background(), networkId, switchStackId, staticRouteId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetDeviceSwitchPort", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string
		var portId string

		resp, httpRes, err := apiClient.SwitchApi.GetDeviceSwitchPort(context.Background(), serial, portId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetDeviceSwitchPorts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.SwitchApi.GetDeviceSwitchPorts(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetDeviceSwitchPortsStatuses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.SwitchApi.GetDeviceSwitchPortsStatuses(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetDeviceSwitchPortsStatusesPackets", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.SwitchApi.GetDeviceSwitchPortsStatusesPackets(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetDeviceSwitchRoutingInterface", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string
		var interfaceId string

		resp, httpRes, err := apiClient.SwitchApi.GetDeviceSwitchRoutingInterface(context.Background(), serial, interfaceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetDeviceSwitchRoutingInterfaceDhcp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string
		var interfaceId string

		resp, httpRes, err := apiClient.SwitchApi.GetDeviceSwitchRoutingInterfaceDhcp(context.Background(), serial, interfaceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetDeviceSwitchRoutingInterfaces", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.SwitchApi.GetDeviceSwitchRoutingInterfaces(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetDeviceSwitchRoutingStaticRoute", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string
		var staticRouteId string

		resp, httpRes, err := apiClient.SwitchApi.GetDeviceSwitchRoutingStaticRoute(context.Background(), serial, staticRouteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetDeviceSwitchRoutingStaticRoutes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.SwitchApi.GetDeviceSwitchRoutingStaticRoutes(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetDeviceSwitchWarmSpare", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.SwitchApi.GetDeviceSwitchWarmSpare(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetNetworkSwitchAccessControlLists", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SwitchApi.GetNetworkSwitchAccessControlLists(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetNetworkSwitchAccessPolicies", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SwitchApi.GetNetworkSwitchAccessPolicies(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetNetworkSwitchAccessPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var accessPolicyNumber string

		resp, httpRes, err := apiClient.SwitchApi.GetNetworkSwitchAccessPolicy(context.Background(), networkId, accessPolicyNumber).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetNetworkSwitchAlternateManagementInterface", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SwitchApi.GetNetworkSwitchAlternateManagementInterface(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetNetworkSwitchDhcpServerPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SwitchApi.GetNetworkSwitchDhcpServerPolicy(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SwitchApi.GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SwitchApi.GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetNetworkSwitchDhcpV4ServersSeen", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SwitchApi.GetNetworkSwitchDhcpV4ServersSeen(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetNetworkSwitchDscpToCosMappings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SwitchApi.GetNetworkSwitchDscpToCosMappings(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetNetworkSwitchLinkAggregations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SwitchApi.GetNetworkSwitchLinkAggregations(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetNetworkSwitchMtu", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SwitchApi.GetNetworkSwitchMtu(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetNetworkSwitchPortSchedules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SwitchApi.GetNetworkSwitchPortSchedules(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetNetworkSwitchQosRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var qosRuleId string

		resp, httpRes, err := apiClient.SwitchApi.GetNetworkSwitchQosRule(context.Background(), networkId, qosRuleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetNetworkSwitchQosRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SwitchApi.GetNetworkSwitchQosRules(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetNetworkSwitchQosRulesOrder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SwitchApi.GetNetworkSwitchQosRulesOrder(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetNetworkSwitchRoutingMulticast", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SwitchApi.GetNetworkSwitchRoutingMulticast(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetNetworkSwitchRoutingMulticastRendezvousPoint", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var rendezvousPointId string

		resp, httpRes, err := apiClient.SwitchApi.GetNetworkSwitchRoutingMulticastRendezvousPoint(context.Background(), networkId, rendezvousPointId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetNetworkSwitchRoutingMulticastRendezvousPoints", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SwitchApi.GetNetworkSwitchRoutingMulticastRendezvousPoints(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetNetworkSwitchRoutingOspf", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SwitchApi.GetNetworkSwitchRoutingOspf(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetNetworkSwitchSettings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SwitchApi.GetNetworkSwitchSettings(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetNetworkSwitchStack", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var switchStackId string

		resp, httpRes, err := apiClient.SwitchApi.GetNetworkSwitchStack(context.Background(), networkId, switchStackId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetNetworkSwitchStackRoutingInterface", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var switchStackId string
		var interfaceId string

		resp, httpRes, err := apiClient.SwitchApi.GetNetworkSwitchStackRoutingInterface(context.Background(), networkId, switchStackId, interfaceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetNetworkSwitchStackRoutingInterfaceDhcp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var switchStackId string
		var interfaceId string

		resp, httpRes, err := apiClient.SwitchApi.GetNetworkSwitchStackRoutingInterfaceDhcp(context.Background(), networkId, switchStackId, interfaceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetNetworkSwitchStackRoutingInterfaces", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var switchStackId string

		resp, httpRes, err := apiClient.SwitchApi.GetNetworkSwitchStackRoutingInterfaces(context.Background(), networkId, switchStackId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetNetworkSwitchStackRoutingStaticRoute", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var switchStackId string
		var staticRouteId string

		resp, httpRes, err := apiClient.SwitchApi.GetNetworkSwitchStackRoutingStaticRoute(context.Background(), networkId, switchStackId, staticRouteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetNetworkSwitchStackRoutingStaticRoutes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var switchStackId string

		resp, httpRes, err := apiClient.SwitchApi.GetNetworkSwitchStackRoutingStaticRoutes(context.Background(), networkId, switchStackId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetNetworkSwitchStacks", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SwitchApi.GetNetworkSwitchStacks(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetNetworkSwitchStormControl", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SwitchApi.GetNetworkSwitchStormControl(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetNetworkSwitchStp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SwitchApi.GetNetworkSwitchStp(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetOrganizationConfigTemplateSwitchProfilePort", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var configTemplateId string
		var profileId string
		var portId string

		resp, httpRes, err := apiClient.SwitchApi.GetOrganizationConfigTemplateSwitchProfilePort(context.Background(), organizationId, configTemplateId, profileId, portId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetOrganizationConfigTemplateSwitchProfilePorts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var configTemplateId string
		var profileId string

		resp, httpRes, err := apiClient.SwitchApi.GetOrganizationConfigTemplateSwitchProfilePorts(context.Background(), organizationId, configTemplateId, profileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetOrganizationConfigTemplateSwitchProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var configTemplateId string

		resp, httpRes, err := apiClient.SwitchApi.GetOrganizationConfigTemplateSwitchProfiles(context.Background(), organizationId, configTemplateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService GetOrganizationSwitchPortsBySwitch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.SwitchApi.GetOrganizationSwitchPortsBySwitch(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService RemoveNetworkSwitchStack", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var switchStackId string

		resp, httpRes, err := apiClient.SwitchApi.RemoveNetworkSwitchStack(context.Background(), networkId, switchStackId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService UpdateDeviceSwitchPort", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string
		var portId string

		resp, httpRes, err := apiClient.SwitchApi.UpdateDeviceSwitchPort(context.Background(), serial, portId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService UpdateDeviceSwitchRoutingInterface", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string
		var interfaceId string

		resp, httpRes, err := apiClient.SwitchApi.UpdateDeviceSwitchRoutingInterface(context.Background(), serial, interfaceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService UpdateDeviceSwitchRoutingInterfaceDhcp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string
		var interfaceId string

		resp, httpRes, err := apiClient.SwitchApi.UpdateDeviceSwitchRoutingInterfaceDhcp(context.Background(), serial, interfaceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService UpdateDeviceSwitchRoutingStaticRoute", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string
		var staticRouteId string

		resp, httpRes, err := apiClient.SwitchApi.UpdateDeviceSwitchRoutingStaticRoute(context.Background(), serial, staticRouteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService UpdateDeviceSwitchWarmSpare", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.SwitchApi.UpdateDeviceSwitchWarmSpare(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService UpdateNetworkSwitchAccessControlLists", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SwitchApi.UpdateNetworkSwitchAccessControlLists(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService UpdateNetworkSwitchAccessPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var accessPolicyNumber string

		resp, httpRes, err := apiClient.SwitchApi.UpdateNetworkSwitchAccessPolicy(context.Background(), networkId, accessPolicyNumber).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService UpdateNetworkSwitchAlternateManagementInterface", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SwitchApi.UpdateNetworkSwitchAlternateManagementInterface(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService UpdateNetworkSwitchDhcpServerPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SwitchApi.UpdateNetworkSwitchDhcpServerPolicy(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var trustedServerId string

		resp, httpRes, err := apiClient.SwitchApi.UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(context.Background(), networkId, trustedServerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService UpdateNetworkSwitchDscpToCosMappings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SwitchApi.UpdateNetworkSwitchDscpToCosMappings(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService UpdateNetworkSwitchLinkAggregation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var linkAggregationId string

		resp, httpRes, err := apiClient.SwitchApi.UpdateNetworkSwitchLinkAggregation(context.Background(), networkId, linkAggregationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService UpdateNetworkSwitchMtu", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SwitchApi.UpdateNetworkSwitchMtu(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService UpdateNetworkSwitchPortSchedule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var portScheduleId string

		resp, httpRes, err := apiClient.SwitchApi.UpdateNetworkSwitchPortSchedule(context.Background(), networkId, portScheduleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService UpdateNetworkSwitchQosRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var qosRuleId string

		resp, httpRes, err := apiClient.SwitchApi.UpdateNetworkSwitchQosRule(context.Background(), networkId, qosRuleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService UpdateNetworkSwitchQosRulesOrder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SwitchApi.UpdateNetworkSwitchQosRulesOrder(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService UpdateNetworkSwitchRoutingMulticast", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SwitchApi.UpdateNetworkSwitchRoutingMulticast(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService UpdateNetworkSwitchRoutingMulticastRendezvousPoint", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var rendezvousPointId string

		resp, httpRes, err := apiClient.SwitchApi.UpdateNetworkSwitchRoutingMulticastRendezvousPoint(context.Background(), networkId, rendezvousPointId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService UpdateNetworkSwitchRoutingOspf", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SwitchApi.UpdateNetworkSwitchRoutingOspf(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService UpdateNetworkSwitchSettings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SwitchApi.UpdateNetworkSwitchSettings(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService UpdateNetworkSwitchStackRoutingInterface", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var switchStackId string
		var interfaceId string

		resp, httpRes, err := apiClient.SwitchApi.UpdateNetworkSwitchStackRoutingInterface(context.Background(), networkId, switchStackId, interfaceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService UpdateNetworkSwitchStackRoutingInterfaceDhcp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var switchStackId string
		var interfaceId string

		resp, httpRes, err := apiClient.SwitchApi.UpdateNetworkSwitchStackRoutingInterfaceDhcp(context.Background(), networkId, switchStackId, interfaceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService UpdateNetworkSwitchStackRoutingStaticRoute", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var switchStackId string
		var staticRouteId string

		resp, httpRes, err := apiClient.SwitchApi.UpdateNetworkSwitchStackRoutingStaticRoute(context.Background(), networkId, switchStackId, staticRouteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService UpdateNetworkSwitchStormControl", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SwitchApi.UpdateNetworkSwitchStormControl(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService UpdateNetworkSwitchStp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SwitchApi.UpdateNetworkSwitchStp(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchApiService UpdateOrganizationConfigTemplateSwitchProfilePort", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var configTemplateId string
		var profileId string
		var portId string

		resp, httpRes, err := apiClient.SwitchApi.UpdateOrganizationConfigTemplateSwitchProfilePort(context.Background(), organizationId, configTemplateId, profileId, portId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
