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

// checks if the UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediate{}

// UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediate Immediate WAN transition terminates all flows (new and existing) on current WAN when it is deemed unreliable.
type UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediate struct {
	// Toggle for enabling or disabling immediate WAN failover and failback
	Enabled bool `json:"enabled"`
}

// NewUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediate instantiates a new UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediate(enabled bool) *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediate {
	this := UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediate{}
	this.Enabled = enabled
	return &this
}

// NewUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediateWithDefaults instantiates a new UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediateWithDefaults() *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediate {
	this := UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediate{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediate) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediate) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediate) SetEnabled(v bool) {
	o.Enabled = v
}

func (o UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediate struct {
	value *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediate
	isSet bool
}

func (v NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediate) Get() *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediate {
	return v.value
}

func (v *NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediate) Set(val *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediate) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediate) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediate(val *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediate) *NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediate {
	return &NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediate{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

