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

// checks if the GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner{}

// GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner struct for GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner
type GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner struct {
	// Option name.
	Name *string `json:"name,omitempty"`
	// Option value.
	Value *string `json:"value,omitempty"`
}

// NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner instantiates a new GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner {
	this := GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner{}
	return &this
}

// NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInnerWithDefaults instantiates a new GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInnerWithDefaults() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner {
	this := GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner) SetValue(v string) {
	o.Value = &v
}

func (o GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner struct {
	value *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner
	isSet bool
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner) Get() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner {
	return v.value
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner) Set(val *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner(val *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner) *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner {
	return &NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


