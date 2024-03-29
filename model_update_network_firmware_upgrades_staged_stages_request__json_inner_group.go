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

// checks if the UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroup{}

// UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroup The Staged Upgrade Group
type UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroup struct {
	// ID of the Staged Upgrade Group
	Id string `json:"id"`
}

// NewUpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroup instantiates a new UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroup(id string) *UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroup {
	this := UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroup{}
	this.Id = id
	return &this
}

// NewUpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroupWithDefaults instantiates a new UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroupWithDefaults() *UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroup {
	this := UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroup{}
	return &this
}

// GetId returns the Id field value
func (o *UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroup) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroup) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroup) SetId(v string) {
	o.Id = v
}

func (o UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableUpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroup struct {
	value *UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroup
	isSet bool
}

func (v NullableUpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroup) Get() *UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroup {
	return v.value
}

func (v *NullableUpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroup) Set(val *UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroup(val *UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroup) *NullableUpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroup {
	return &NullableUpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroup{value: val, isSet: true}
}

func (v NullableUpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


