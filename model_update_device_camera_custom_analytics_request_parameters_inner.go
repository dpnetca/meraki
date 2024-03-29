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

// checks if the UpdateDeviceCameraCustomAnalyticsRequestParametersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeviceCameraCustomAnalyticsRequestParametersInner{}

// UpdateDeviceCameraCustomAnalyticsRequestParametersInner struct for UpdateDeviceCameraCustomAnalyticsRequestParametersInner
type UpdateDeviceCameraCustomAnalyticsRequestParametersInner struct {
	// Name of the parameter
	Name string `json:"name"`
	// Value of the parameter
	Value string `json:"value"`
}

// NewUpdateDeviceCameraCustomAnalyticsRequestParametersInner instantiates a new UpdateDeviceCameraCustomAnalyticsRequestParametersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceCameraCustomAnalyticsRequestParametersInner(name string, value string) *UpdateDeviceCameraCustomAnalyticsRequestParametersInner {
	this := UpdateDeviceCameraCustomAnalyticsRequestParametersInner{}
	this.Name = name
	this.Value = value
	return &this
}

// NewUpdateDeviceCameraCustomAnalyticsRequestParametersInnerWithDefaults instantiates a new UpdateDeviceCameraCustomAnalyticsRequestParametersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceCameraCustomAnalyticsRequestParametersInnerWithDefaults() *UpdateDeviceCameraCustomAnalyticsRequestParametersInner {
	this := UpdateDeviceCameraCustomAnalyticsRequestParametersInner{}
	return &this
}

// GetName returns the Name field value
func (o *UpdateDeviceCameraCustomAnalyticsRequestParametersInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateDeviceCameraCustomAnalyticsRequestParametersInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateDeviceCameraCustomAnalyticsRequestParametersInner) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *UpdateDeviceCameraCustomAnalyticsRequestParametersInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *UpdateDeviceCameraCustomAnalyticsRequestParametersInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *UpdateDeviceCameraCustomAnalyticsRequestParametersInner) SetValue(v string) {
	o.Value = v
}

func (o UpdateDeviceCameraCustomAnalyticsRequestParametersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeviceCameraCustomAnalyticsRequestParametersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableUpdateDeviceCameraCustomAnalyticsRequestParametersInner struct {
	value *UpdateDeviceCameraCustomAnalyticsRequestParametersInner
	isSet bool
}

func (v NullableUpdateDeviceCameraCustomAnalyticsRequestParametersInner) Get() *UpdateDeviceCameraCustomAnalyticsRequestParametersInner {
	return v.value
}

func (v *NullableUpdateDeviceCameraCustomAnalyticsRequestParametersInner) Set(val *UpdateDeviceCameraCustomAnalyticsRequestParametersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceCameraCustomAnalyticsRequestParametersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceCameraCustomAnalyticsRequestParametersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceCameraCustomAnalyticsRequestParametersInner(val *UpdateDeviceCameraCustomAnalyticsRequestParametersInner) *NullableUpdateDeviceCameraCustomAnalyticsRequestParametersInner {
	return &NullableUpdateDeviceCameraCustomAnalyticsRequestParametersInner{value: val, isSet: true}
}

func (v NullableUpdateDeviceCameraCustomAnalyticsRequestParametersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceCameraCustomAnalyticsRequestParametersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


