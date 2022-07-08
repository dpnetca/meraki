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

// CreateDeviceSwitchRoutingInterfaceRequestOspfSettings The OSPF routing settings of the interface.
type CreateDeviceSwitchRoutingInterfaceRequestOspfSettings struct {
	// The OSPF area to which this interface should belong. Can be either 'disabled' or the identifier of an           existing OSPF area. Defaults to 'disabled'.
	Area *string `json:"area,omitempty"`
	// The path cost for this interface. Defaults to 1, but can be increased up to 65535           to give lower priority.
	Cost *int32 `json:"cost,omitempty"`
	// When enabled, OSPF will not run on the interface, but the subnet will still be advertised.
	IsPassiveEnabled *bool `json:"isPassiveEnabled,omitempty"`
}

// NewCreateDeviceSwitchRoutingInterfaceRequestOspfSettings instantiates a new CreateDeviceSwitchRoutingInterfaceRequestOspfSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDeviceSwitchRoutingInterfaceRequestOspfSettings() *CreateDeviceSwitchRoutingInterfaceRequestOspfSettings {
	this := CreateDeviceSwitchRoutingInterfaceRequestOspfSettings{}
	return &this
}

// NewCreateDeviceSwitchRoutingInterfaceRequestOspfSettingsWithDefaults instantiates a new CreateDeviceSwitchRoutingInterfaceRequestOspfSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDeviceSwitchRoutingInterfaceRequestOspfSettingsWithDefaults() *CreateDeviceSwitchRoutingInterfaceRequestOspfSettings {
	this := CreateDeviceSwitchRoutingInterfaceRequestOspfSettings{}
	return &this
}

// GetArea returns the Area field value if set, zero value otherwise.
func (o *CreateDeviceSwitchRoutingInterfaceRequestOspfSettings) GetArea() string {
	if o == nil || o.Area == nil {
		var ret string
		return ret
	}
	return *o.Area
}

// GetAreaOk returns a tuple with the Area field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceSwitchRoutingInterfaceRequestOspfSettings) GetAreaOk() (*string, bool) {
	if o == nil || o.Area == nil {
		return nil, false
	}
	return o.Area, true
}

// HasArea returns a boolean if a field has been set.
func (o *CreateDeviceSwitchRoutingInterfaceRequestOspfSettings) HasArea() bool {
	if o != nil && o.Area != nil {
		return true
	}

	return false
}

// SetArea gets a reference to the given string and assigns it to the Area field.
func (o *CreateDeviceSwitchRoutingInterfaceRequestOspfSettings) SetArea(v string) {
	o.Area = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *CreateDeviceSwitchRoutingInterfaceRequestOspfSettings) GetCost() int32 {
	if o == nil || o.Cost == nil {
		var ret int32
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceSwitchRoutingInterfaceRequestOspfSettings) GetCostOk() (*int32, bool) {
	if o == nil || o.Cost == nil {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *CreateDeviceSwitchRoutingInterfaceRequestOspfSettings) HasCost() bool {
	if o != nil && o.Cost != nil {
		return true
	}

	return false
}

// SetCost gets a reference to the given int32 and assigns it to the Cost field.
func (o *CreateDeviceSwitchRoutingInterfaceRequestOspfSettings) SetCost(v int32) {
	o.Cost = &v
}

// GetIsPassiveEnabled returns the IsPassiveEnabled field value if set, zero value otherwise.
func (o *CreateDeviceSwitchRoutingInterfaceRequestOspfSettings) GetIsPassiveEnabled() bool {
	if o == nil || o.IsPassiveEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsPassiveEnabled
}

// GetIsPassiveEnabledOk returns a tuple with the IsPassiveEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceSwitchRoutingInterfaceRequestOspfSettings) GetIsPassiveEnabledOk() (*bool, bool) {
	if o == nil || o.IsPassiveEnabled == nil {
		return nil, false
	}
	return o.IsPassiveEnabled, true
}

// HasIsPassiveEnabled returns a boolean if a field has been set.
func (o *CreateDeviceSwitchRoutingInterfaceRequestOspfSettings) HasIsPassiveEnabled() bool {
	if o != nil && o.IsPassiveEnabled != nil {
		return true
	}

	return false
}

// SetIsPassiveEnabled gets a reference to the given bool and assigns it to the IsPassiveEnabled field.
func (o *CreateDeviceSwitchRoutingInterfaceRequestOspfSettings) SetIsPassiveEnabled(v bool) {
	o.IsPassiveEnabled = &v
}

func (o CreateDeviceSwitchRoutingInterfaceRequestOspfSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Area != nil {
		toSerialize["area"] = o.Area
	}
	if o.Cost != nil {
		toSerialize["cost"] = o.Cost
	}
	if o.IsPassiveEnabled != nil {
		toSerialize["isPassiveEnabled"] = o.IsPassiveEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableCreateDeviceSwitchRoutingInterfaceRequestOspfSettings struct {
	value *CreateDeviceSwitchRoutingInterfaceRequestOspfSettings
	isSet bool
}

func (v NullableCreateDeviceSwitchRoutingInterfaceRequestOspfSettings) Get() *CreateDeviceSwitchRoutingInterfaceRequestOspfSettings {
	return v.value
}

func (v *NullableCreateDeviceSwitchRoutingInterfaceRequestOspfSettings) Set(val *CreateDeviceSwitchRoutingInterfaceRequestOspfSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDeviceSwitchRoutingInterfaceRequestOspfSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDeviceSwitchRoutingInterfaceRequestOspfSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDeviceSwitchRoutingInterfaceRequestOspfSettings(val *CreateDeviceSwitchRoutingInterfaceRequestOspfSettings) *NullableCreateDeviceSwitchRoutingInterfaceRequestOspfSettings {
	return &NullableCreateDeviceSwitchRoutingInterfaceRequestOspfSettings{value: val, isSet: true}
}

func (v NullableCreateDeviceSwitchRoutingInterfaceRequestOspfSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDeviceSwitchRoutingInterfaceRequestOspfSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

