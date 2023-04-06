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

// checks if the CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersion{}

// CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersion The version to be updated to for switch devices
type CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersion struct {
	// The version ID
	Id string `json:"id"`
}

// NewCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersion instantiates a new CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersion(id string) *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersion {
	this := CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersion{}
	this.Id = id
	return &this
}

// NewCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersionWithDefaults instantiates a new CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersionWithDefaults() *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersion {
	this := CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersion{}
	return &this
}

// GetId returns the Id field value
func (o *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersion) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersion) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersion) SetId(v string) {
	o.Id = v
}

func (o CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersion struct {
	value *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersion
	isSet bool
}

func (v NullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersion) Get() *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersion {
	return v.value
}

func (v *NullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersion) Set(val *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersion(val *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersion) *NullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersion {
	return &NullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersion{value: val, isSet: true}
}

func (v NullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


