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

// checks if the UpdateNetworkSwitchStpRequestStpBridgePriorityInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkSwitchStpRequestStpBridgePriorityInner{}

// UpdateNetworkSwitchStpRequestStpBridgePriorityInner struct for UpdateNetworkSwitchStpRequestStpBridgePriorityInner
type UpdateNetworkSwitchStpRequestStpBridgePriorityInner struct {
	// List of switch profile IDs
	SwitchProfiles []string `json:"switchProfiles,omitempty"`
	// List of switch serial numbers
	Switches []string `json:"switches,omitempty"`
	// List of stack IDs
	Stacks []string `json:"stacks,omitempty"`
	// STP priority for switch, stacks, or switch profiles
	StpPriority int32 `json:"stpPriority"`
}

// NewUpdateNetworkSwitchStpRequestStpBridgePriorityInner instantiates a new UpdateNetworkSwitchStpRequestStpBridgePriorityInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSwitchStpRequestStpBridgePriorityInner(stpPriority int32) *UpdateNetworkSwitchStpRequestStpBridgePriorityInner {
	this := UpdateNetworkSwitchStpRequestStpBridgePriorityInner{}
	this.StpPriority = stpPriority
	return &this
}

// NewUpdateNetworkSwitchStpRequestStpBridgePriorityInnerWithDefaults instantiates a new UpdateNetworkSwitchStpRequestStpBridgePriorityInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSwitchStpRequestStpBridgePriorityInnerWithDefaults() *UpdateNetworkSwitchStpRequestStpBridgePriorityInner {
	this := UpdateNetworkSwitchStpRequestStpBridgePriorityInner{}
	return &this
}

// GetSwitchProfiles returns the SwitchProfiles field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchStpRequestStpBridgePriorityInner) GetSwitchProfiles() []string {
	if o == nil || IsNil(o.SwitchProfiles) {
		var ret []string
		return ret
	}
	return o.SwitchProfiles
}

// GetSwitchProfilesOk returns a tuple with the SwitchProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchStpRequestStpBridgePriorityInner) GetSwitchProfilesOk() ([]string, bool) {
	if o == nil || IsNil(o.SwitchProfiles) {
		return nil, false
	}
	return o.SwitchProfiles, true
}

// HasSwitchProfiles returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchStpRequestStpBridgePriorityInner) HasSwitchProfiles() bool {
	if o != nil && !IsNil(o.SwitchProfiles) {
		return true
	}

	return false
}

// SetSwitchProfiles gets a reference to the given []string and assigns it to the SwitchProfiles field.
func (o *UpdateNetworkSwitchStpRequestStpBridgePriorityInner) SetSwitchProfiles(v []string) {
	o.SwitchProfiles = v
}

// GetSwitches returns the Switches field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchStpRequestStpBridgePriorityInner) GetSwitches() []string {
	if o == nil || IsNil(o.Switches) {
		var ret []string
		return ret
	}
	return o.Switches
}

// GetSwitchesOk returns a tuple with the Switches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchStpRequestStpBridgePriorityInner) GetSwitchesOk() ([]string, bool) {
	if o == nil || IsNil(o.Switches) {
		return nil, false
	}
	return o.Switches, true
}

// HasSwitches returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchStpRequestStpBridgePriorityInner) HasSwitches() bool {
	if o != nil && !IsNil(o.Switches) {
		return true
	}

	return false
}

// SetSwitches gets a reference to the given []string and assigns it to the Switches field.
func (o *UpdateNetworkSwitchStpRequestStpBridgePriorityInner) SetSwitches(v []string) {
	o.Switches = v
}

// GetStacks returns the Stacks field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchStpRequestStpBridgePriorityInner) GetStacks() []string {
	if o == nil || IsNil(o.Stacks) {
		var ret []string
		return ret
	}
	return o.Stacks
}

// GetStacksOk returns a tuple with the Stacks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchStpRequestStpBridgePriorityInner) GetStacksOk() ([]string, bool) {
	if o == nil || IsNil(o.Stacks) {
		return nil, false
	}
	return o.Stacks, true
}

// HasStacks returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchStpRequestStpBridgePriorityInner) HasStacks() bool {
	if o != nil && !IsNil(o.Stacks) {
		return true
	}

	return false
}

// SetStacks gets a reference to the given []string and assigns it to the Stacks field.
func (o *UpdateNetworkSwitchStpRequestStpBridgePriorityInner) SetStacks(v []string) {
	o.Stacks = v
}

// GetStpPriority returns the StpPriority field value
func (o *UpdateNetworkSwitchStpRequestStpBridgePriorityInner) GetStpPriority() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StpPriority
}

// GetStpPriorityOk returns a tuple with the StpPriority field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchStpRequestStpBridgePriorityInner) GetStpPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StpPriority, true
}

// SetStpPriority sets field value
func (o *UpdateNetworkSwitchStpRequestStpBridgePriorityInner) SetStpPriority(v int32) {
	o.StpPriority = v
}

func (o UpdateNetworkSwitchStpRequestStpBridgePriorityInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkSwitchStpRequestStpBridgePriorityInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SwitchProfiles) {
		toSerialize["switchProfiles"] = o.SwitchProfiles
	}
	if !IsNil(o.Switches) {
		toSerialize["switches"] = o.Switches
	}
	if !IsNil(o.Stacks) {
		toSerialize["stacks"] = o.Stacks
	}
	toSerialize["stpPriority"] = o.StpPriority
	return toSerialize, nil
}

type NullableUpdateNetworkSwitchStpRequestStpBridgePriorityInner struct {
	value *UpdateNetworkSwitchStpRequestStpBridgePriorityInner
	isSet bool
}

func (v NullableUpdateNetworkSwitchStpRequestStpBridgePriorityInner) Get() *UpdateNetworkSwitchStpRequestStpBridgePriorityInner {
	return v.value
}

func (v *NullableUpdateNetworkSwitchStpRequestStpBridgePriorityInner) Set(val *UpdateNetworkSwitchStpRequestStpBridgePriorityInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSwitchStpRequestStpBridgePriorityInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSwitchStpRequestStpBridgePriorityInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSwitchStpRequestStpBridgePriorityInner(val *UpdateNetworkSwitchStpRequestStpBridgePriorityInner) *NullableUpdateNetworkSwitchStpRequestStpBridgePriorityInner {
	return &NullableUpdateNetworkSwitchStpRequestStpBridgePriorityInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkSwitchStpRequestStpBridgePriorityInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSwitchStpRequestStpBridgePriorityInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


