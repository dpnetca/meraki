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

// checks if the UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgrade type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgrade{}

// UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgrade The pending firmware upgrade if it exists
type UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgrade struct {
	// The time of the last successful upgrade
	Time *string `json:"time,omitempty"`
	ToVersion *UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion `json:"toVersion,omitempty"`
}

// NewUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgrade instantiates a new UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgrade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgrade() *UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgrade {
	this := UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgrade{}
	return &this
}

// NewUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeWithDefaults instantiates a new UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgrade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeWithDefaults() *UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgrade {
	this := UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgrade{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgrade) GetTime() string {
	if o == nil || IsNil(o.Time) {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgrade) GetTimeOk() (*string, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgrade) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgrade) SetTime(v string) {
	o.Time = &v
}

// GetToVersion returns the ToVersion field value if set, zero value otherwise.
func (o *UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgrade) GetToVersion() UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion {
	if o == nil || IsNil(o.ToVersion) {
		var ret UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion
		return ret
	}
	return *o.ToVersion
}

// GetToVersionOk returns a tuple with the ToVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgrade) GetToVersionOk() (*UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion, bool) {
	if o == nil || IsNil(o.ToVersion) {
		return nil, false
	}
	return o.ToVersion, true
}

// HasToVersion returns a boolean if a field has been set.
func (o *UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgrade) HasToVersion() bool {
	if o != nil && !IsNil(o.ToVersion) {
		return true
	}

	return false
}

// SetToVersion gets a reference to the given UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion and assigns it to the ToVersion field.
func (o *UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgrade) SetToVersion(v UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion) {
	o.ToVersion = &v
}

func (o UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgrade) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgrade) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.ToVersion) {
		toSerialize["toVersion"] = o.ToVersion
	}
	return toSerialize, nil
}

type NullableUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgrade struct {
	value *UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgrade
	isSet bool
}

func (v NullableUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgrade) Get() *UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgrade {
	return v.value
}

func (v *NullableUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgrade) Set(val *UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgrade) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgrade) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgrade(val *UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgrade) *NullableUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgrade {
	return &NullableUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgrade{value: val, isSet: true}
}

func (v NullableUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


