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

// checks if the GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices{}

// GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices The devices and Switch Stacks assigned to the Group
type GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices struct {
	// Data Array of Devices containing the name and serial
	Devices []GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner `json:"devices,omitempty"`
	// Data Array of Switch Stacks containing the name and id
	SwitchStacks []GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner `json:"switchStacks,omitempty"`
}

// NewGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices instantiates a new GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices() *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices {
	this := GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices{}
	return &this
}

// NewGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesWithDefaults instantiates a new GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesWithDefaults() *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices {
	this := GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices{}
	return &this
}

// GetDevices returns the Devices field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices) GetDevices() []GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner {
	if o == nil || IsNil(o.Devices) {
		var ret []GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner
		return ret
	}
	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices) GetDevicesOk() ([]GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner, bool) {
	if o == nil || IsNil(o.Devices) {
		return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices) HasDevices() bool {
	if o != nil && !IsNil(o.Devices) {
		return true
	}

	return false
}

// SetDevices gets a reference to the given []GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner and assigns it to the Devices field.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices) SetDevices(v []GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner) {
	o.Devices = v
}

// GetSwitchStacks returns the SwitchStacks field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices) GetSwitchStacks() []GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner {
	if o == nil || IsNil(o.SwitchStacks) {
		var ret []GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner
		return ret
	}
	return o.SwitchStacks
}

// GetSwitchStacksOk returns a tuple with the SwitchStacks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices) GetSwitchStacksOk() ([]GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner, bool) {
	if o == nil || IsNil(o.SwitchStacks) {
		return nil, false
	}
	return o.SwitchStacks, true
}

// HasSwitchStacks returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices) HasSwitchStacks() bool {
	if o != nil && !IsNil(o.SwitchStacks) {
		return true
	}

	return false
}

// SetSwitchStacks gets a reference to the given []GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner and assigns it to the SwitchStacks field.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices) SetSwitchStacks(v []GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner) {
	o.SwitchStacks = v
}

func (o GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Devices) {
		toSerialize["devices"] = o.Devices
	}
	if !IsNil(o.SwitchStacks) {
		toSerialize["switchStacks"] = o.SwitchStacks
	}
	return toSerialize, nil
}

type NullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices struct {
	value *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices
	isSet bool
}

func (v NullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices) Get() *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices {
	return v.value
}

func (v *NullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices) Set(val *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices(val *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices) *NullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices {
	return &NullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices{value: val, isSet: true}
}

func (v NullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


