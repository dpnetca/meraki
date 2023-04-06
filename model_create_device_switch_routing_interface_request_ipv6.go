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

// checks if the CreateDeviceSwitchRoutingInterfaceRequestIpv6 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDeviceSwitchRoutingInterfaceRequestIpv6{}

// CreateDeviceSwitchRoutingInterfaceRequestIpv6 The IPv6 settings of the interface.
type CreateDeviceSwitchRoutingInterfaceRequestIpv6 struct {
	// The IPv6 assignment mode for the interface. Can be either 'eui-64' or 'static'.
	AssignmentMode *string `json:"assignmentMode,omitempty"`
	// The IPv6 prefix of the interface. Required if IPv6 object is included.
	Prefix *string `json:"prefix,omitempty"`
	// The IPv6 address of the interface. Required if assignmentMode is 'static'. Must not be included if           assignmentMode is 'eui-64'.
	Address *string `json:"address,omitempty"`
	// The IPv6 default gateway of the interface. Required if prefix is defined and this is the first           interface with IPv6 configured for the switch.
	Gateway *string `json:"gateway,omitempty"`
}

// NewCreateDeviceSwitchRoutingInterfaceRequestIpv6 instantiates a new CreateDeviceSwitchRoutingInterfaceRequestIpv6 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDeviceSwitchRoutingInterfaceRequestIpv6() *CreateDeviceSwitchRoutingInterfaceRequestIpv6 {
	this := CreateDeviceSwitchRoutingInterfaceRequestIpv6{}
	return &this
}

// NewCreateDeviceSwitchRoutingInterfaceRequestIpv6WithDefaults instantiates a new CreateDeviceSwitchRoutingInterfaceRequestIpv6 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDeviceSwitchRoutingInterfaceRequestIpv6WithDefaults() *CreateDeviceSwitchRoutingInterfaceRequestIpv6 {
	this := CreateDeviceSwitchRoutingInterfaceRequestIpv6{}
	return &this
}

// GetAssignmentMode returns the AssignmentMode field value if set, zero value otherwise.
func (o *CreateDeviceSwitchRoutingInterfaceRequestIpv6) GetAssignmentMode() string {
	if o == nil || IsNil(o.AssignmentMode) {
		var ret string
		return ret
	}
	return *o.AssignmentMode
}

// GetAssignmentModeOk returns a tuple with the AssignmentMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceSwitchRoutingInterfaceRequestIpv6) GetAssignmentModeOk() (*string, bool) {
	if o == nil || IsNil(o.AssignmentMode) {
		return nil, false
	}
	return o.AssignmentMode, true
}

// HasAssignmentMode returns a boolean if a field has been set.
func (o *CreateDeviceSwitchRoutingInterfaceRequestIpv6) HasAssignmentMode() bool {
	if o != nil && !IsNil(o.AssignmentMode) {
		return true
	}

	return false
}

// SetAssignmentMode gets a reference to the given string and assigns it to the AssignmentMode field.
func (o *CreateDeviceSwitchRoutingInterfaceRequestIpv6) SetAssignmentMode(v string) {
	o.AssignmentMode = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *CreateDeviceSwitchRoutingInterfaceRequestIpv6) GetPrefix() string {
	if o == nil || IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceSwitchRoutingInterfaceRequestIpv6) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *CreateDeviceSwitchRoutingInterfaceRequestIpv6) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *CreateDeviceSwitchRoutingInterfaceRequestIpv6) SetPrefix(v string) {
	o.Prefix = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *CreateDeviceSwitchRoutingInterfaceRequestIpv6) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceSwitchRoutingInterfaceRequestIpv6) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *CreateDeviceSwitchRoutingInterfaceRequestIpv6) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *CreateDeviceSwitchRoutingInterfaceRequestIpv6) SetAddress(v string) {
	o.Address = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *CreateDeviceSwitchRoutingInterfaceRequestIpv6) GetGateway() string {
	if o == nil || IsNil(o.Gateway) {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceSwitchRoutingInterfaceRequestIpv6) GetGatewayOk() (*string, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *CreateDeviceSwitchRoutingInterfaceRequestIpv6) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *CreateDeviceSwitchRoutingInterfaceRequestIpv6) SetGateway(v string) {
	o.Gateway = &v
}

func (o CreateDeviceSwitchRoutingInterfaceRequestIpv6) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDeviceSwitchRoutingInterfaceRequestIpv6) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssignmentMode) {
		toSerialize["assignmentMode"] = o.AssignmentMode
	}
	if !IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	return toSerialize, nil
}

type NullableCreateDeviceSwitchRoutingInterfaceRequestIpv6 struct {
	value *CreateDeviceSwitchRoutingInterfaceRequestIpv6
	isSet bool
}

func (v NullableCreateDeviceSwitchRoutingInterfaceRequestIpv6) Get() *CreateDeviceSwitchRoutingInterfaceRequestIpv6 {
	return v.value
}

func (v *NullableCreateDeviceSwitchRoutingInterfaceRequestIpv6) Set(val *CreateDeviceSwitchRoutingInterfaceRequestIpv6) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDeviceSwitchRoutingInterfaceRequestIpv6) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDeviceSwitchRoutingInterfaceRequestIpv6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDeviceSwitchRoutingInterfaceRequestIpv6(val *CreateDeviceSwitchRoutingInterfaceRequestIpv6) *NullableCreateDeviceSwitchRoutingInterfaceRequestIpv6 {
	return &NullableCreateDeviceSwitchRoutingInterfaceRequestIpv6{value: val, isSet: true}
}

func (v NullableCreateDeviceSwitchRoutingInterfaceRequestIpv6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDeviceSwitchRoutingInterfaceRequestIpv6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


