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

// checks if the UpdateNetworkApplianceFirewallOneToManyNatRulesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceFirewallOneToManyNatRulesRequest{}

// UpdateNetworkApplianceFirewallOneToManyNatRulesRequest struct for UpdateNetworkApplianceFirewallOneToManyNatRulesRequest
type UpdateNetworkApplianceFirewallOneToManyNatRulesRequest struct {
	// An array of 1:Many nat rules
	Rules []UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner `json:"rules"`
}

// NewUpdateNetworkApplianceFirewallOneToManyNatRulesRequest instantiates a new UpdateNetworkApplianceFirewallOneToManyNatRulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceFirewallOneToManyNatRulesRequest(rules []UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner) *UpdateNetworkApplianceFirewallOneToManyNatRulesRequest {
	this := UpdateNetworkApplianceFirewallOneToManyNatRulesRequest{}
	this.Rules = rules
	return &this
}

// NewUpdateNetworkApplianceFirewallOneToManyNatRulesRequestWithDefaults instantiates a new UpdateNetworkApplianceFirewallOneToManyNatRulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceFirewallOneToManyNatRulesRequestWithDefaults() *UpdateNetworkApplianceFirewallOneToManyNatRulesRequest {
	this := UpdateNetworkApplianceFirewallOneToManyNatRulesRequest{}
	return &this
}

// GetRules returns the Rules field value
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesRequest) GetRules() []UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner {
	if o == nil {
		var ret []UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesRequest) GetRulesOk() ([]UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesRequest) SetRules(v []UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner) {
	o.Rules = v
}

func (o UpdateNetworkApplianceFirewallOneToManyNatRulesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceFirewallOneToManyNatRulesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rules"] = o.Rules
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceFirewallOneToManyNatRulesRequest struct {
	value *UpdateNetworkApplianceFirewallOneToManyNatRulesRequest
	isSet bool
}

func (v NullableUpdateNetworkApplianceFirewallOneToManyNatRulesRequest) Get() *UpdateNetworkApplianceFirewallOneToManyNatRulesRequest {
	return v.value
}

func (v *NullableUpdateNetworkApplianceFirewallOneToManyNatRulesRequest) Set(val *UpdateNetworkApplianceFirewallOneToManyNatRulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceFirewallOneToManyNatRulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceFirewallOneToManyNatRulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceFirewallOneToManyNatRulesRequest(val *UpdateNetworkApplianceFirewallOneToManyNatRulesRequest) *NullableUpdateNetworkApplianceFirewallOneToManyNatRulesRequest {
	return &NullableUpdateNetworkApplianceFirewallOneToManyNatRulesRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceFirewallOneToManyNatRulesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceFirewallOneToManyNatRulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


