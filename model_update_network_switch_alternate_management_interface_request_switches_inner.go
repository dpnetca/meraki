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

// checks if the UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner{}

// UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner struct for UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner
type UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner struct {
	// Switch serial number
	Serial string `json:"serial"`
	// Switch alternative management IP. To remove a prior IP setting, provide an empty string
	AlternateManagementIp string `json:"alternateManagementIp"`
	// Switch subnet mask must be in IP format. Only and must be specified for Polaris switches
	SubnetMask *string `json:"subnetMask,omitempty"`
	// Switch gateway must be in IP format. Only and must be specified for Polaris switches
	Gateway *string `json:"gateway,omitempty"`
}

// NewUpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner instantiates a new UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner(serial string, alternateManagementIp string) *UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner {
	this := UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner{}
	this.Serial = serial
	this.AlternateManagementIp = alternateManagementIp
	return &this
}

// NewUpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInnerWithDefaults instantiates a new UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInnerWithDefaults() *UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner {
	this := UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner{}
	return &this
}

// GetSerial returns the Serial field value
func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) GetSerial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Serial
}

// GetSerialOk returns a tuple with the Serial field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) GetSerialOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Serial, true
}

// SetSerial sets field value
func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) SetSerial(v string) {
	o.Serial = v
}

// GetAlternateManagementIp returns the AlternateManagementIp field value
func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) GetAlternateManagementIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlternateManagementIp
}

// GetAlternateManagementIpOk returns a tuple with the AlternateManagementIp field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) GetAlternateManagementIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlternateManagementIp, true
}

// SetAlternateManagementIp sets field value
func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) SetAlternateManagementIp(v string) {
	o.AlternateManagementIp = v
}

// GetSubnetMask returns the SubnetMask field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) GetSubnetMask() string {
	if o == nil || IsNil(o.SubnetMask) {
		var ret string
		return ret
	}
	return *o.SubnetMask
}

// GetSubnetMaskOk returns a tuple with the SubnetMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) GetSubnetMaskOk() (*string, bool) {
	if o == nil || IsNil(o.SubnetMask) {
		return nil, false
	}
	return o.SubnetMask, true
}

// HasSubnetMask returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) HasSubnetMask() bool {
	if o != nil && !IsNil(o.SubnetMask) {
		return true
	}

	return false
}

// SetSubnetMask gets a reference to the given string and assigns it to the SubnetMask field.
func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) SetSubnetMask(v string) {
	o.SubnetMask = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) GetGateway() string {
	if o == nil || IsNil(o.Gateway) {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) GetGatewayOk() (*string, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) SetGateway(v string) {
	o.Gateway = &v
}

func (o UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["serial"] = o.Serial
	toSerialize["alternateManagementIp"] = o.AlternateManagementIp
	if !IsNil(o.SubnetMask) {
		toSerialize["subnetMask"] = o.SubnetMask
	}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	return toSerialize, nil
}

type NullableUpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner struct {
	value *UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner
	isSet bool
}

func (v NullableUpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) Get() *UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner {
	return v.value
}

func (v *NullableUpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) Set(val *UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner(val *UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) *NullableUpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner {
	return &NullableUpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


