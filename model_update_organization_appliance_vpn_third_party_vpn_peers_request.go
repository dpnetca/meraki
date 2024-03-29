/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 April, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package meraki

import (
	"encoding/json"
)

// checks if the UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest{}

// UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest struct for UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest
type UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest struct {
	// The list of VPN peers
	Peers []UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner `json:"peers"`
}

// NewUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest instantiates a new UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest(peers []UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest {
	this := UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest{}
	this.Peers = peers
	return &this
}

// NewUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestWithDefaults instantiates a new UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestWithDefaults() *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest {
	this := UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest{}
	return &this
}

// GetPeers returns the Peers field value
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest) GetPeers() []UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner {
	if o == nil {
		var ret []UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner
		return ret
	}

	return o.Peers
}

// GetPeersOk returns a tuple with the Peers field value
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest) GetPeersOk() ([]UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Peers, true
}

// SetPeers sets field value
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest) SetPeers(v []UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) {
	o.Peers = v
}

func (o UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["peers"] = o.Peers
	return toSerialize, nil
}

type NullableUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest struct {
	value *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest
	isSet bool
}

func (v NullableUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest) Get() *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest {
	return v.value
}

func (v *NullableUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest) Set(val *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest(val *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest) *NullableUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest {
	return &NullableUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest{value: val, isSet: true}
}

func (v NullableUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


