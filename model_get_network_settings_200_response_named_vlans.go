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

// checks if the GetNetworkSettings200ResponseNamedVlans type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSettings200ResponseNamedVlans{}

// GetNetworkSettings200ResponseNamedVlans A hash of Named VLANs options applied to the Network.
type GetNetworkSettings200ResponseNamedVlans struct {
	// Enables / disables Named VLANs on the Network.
	Enabled bool `json:"enabled"`
}

// NewGetNetworkSettings200ResponseNamedVlans instantiates a new GetNetworkSettings200ResponseNamedVlans object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSettings200ResponseNamedVlans(enabled bool) *GetNetworkSettings200ResponseNamedVlans {
	this := GetNetworkSettings200ResponseNamedVlans{}
	this.Enabled = enabled
	return &this
}

// NewGetNetworkSettings200ResponseNamedVlansWithDefaults instantiates a new GetNetworkSettings200ResponseNamedVlans object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSettings200ResponseNamedVlansWithDefaults() *GetNetworkSettings200ResponseNamedVlans {
	this := GetNetworkSettings200ResponseNamedVlans{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *GetNetworkSettings200ResponseNamedVlans) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *GetNetworkSettings200ResponseNamedVlans) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *GetNetworkSettings200ResponseNamedVlans) SetEnabled(v bool) {
	o.Enabled = v
}

func (o GetNetworkSettings200ResponseNamedVlans) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSettings200ResponseNamedVlans) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableGetNetworkSettings200ResponseNamedVlans struct {
	value *GetNetworkSettings200ResponseNamedVlans
	isSet bool
}

func (v NullableGetNetworkSettings200ResponseNamedVlans) Get() *GetNetworkSettings200ResponseNamedVlans {
	return v.value
}

func (v *NullableGetNetworkSettings200ResponseNamedVlans) Set(val *GetNetworkSettings200ResponseNamedVlans) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSettings200ResponseNamedVlans) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSettings200ResponseNamedVlans) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSettings200ResponseNamedVlans(val *GetNetworkSettings200ResponseNamedVlans) *NullableGetNetworkSettings200ResponseNamedVlans {
	return &NullableGetNetworkSettings200ResponseNamedVlans{value: val, isSet: true}
}

func (v NullableGetNetworkSettings200ResponseNamedVlans) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSettings200ResponseNamedVlans) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


