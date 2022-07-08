/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 July, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.23.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package meraki

import (
	"encoding/json"
)

// UpdateNetworkWirelessSsidRequestDot11w The current setting for Protected Management Frames (802.11w).
type UpdateNetworkWirelessSsidRequestDot11w struct {
	// Whether 802.11w is enabled or not.
	Enabled *bool `json:"enabled,omitempty"`
	// (Optional) Whether 802.11w is required or not.
	Required *bool `json:"required,omitempty"`
}

// NewUpdateNetworkWirelessSsidRequestDot11w instantiates a new UpdateNetworkWirelessSsidRequestDot11w object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidRequestDot11w() *UpdateNetworkWirelessSsidRequestDot11w {
	this := UpdateNetworkWirelessSsidRequestDot11w{}
	return &this
}

// NewUpdateNetworkWirelessSsidRequestDot11wWithDefaults instantiates a new UpdateNetworkWirelessSsidRequestDot11w object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidRequestDot11wWithDefaults() *UpdateNetworkWirelessSsidRequestDot11w {
	this := UpdateNetworkWirelessSsidRequestDot11w{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestDot11w) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestDot11w) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestDot11w) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateNetworkWirelessSsidRequestDot11w) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestDot11w) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestDot11w) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestDot11w) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *UpdateNetworkWirelessSsidRequestDot11w) SetRequired(v bool) {
	o.Required = &v
}

func (o UpdateNetworkWirelessSsidRequestDot11w) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkWirelessSsidRequestDot11w struct {
	value *UpdateNetworkWirelessSsidRequestDot11w
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidRequestDot11w) Get() *UpdateNetworkWirelessSsidRequestDot11w {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidRequestDot11w) Set(val *UpdateNetworkWirelessSsidRequestDot11w) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidRequestDot11w) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidRequestDot11w) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidRequestDot11w(val *UpdateNetworkWirelessSsidRequestDot11w) *NullableUpdateNetworkWirelessSsidRequestDot11w {
	return &NullableUpdateNetworkWirelessSsidRequestDot11w{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidRequestDot11w) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidRequestDot11w) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


