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

// checks if the GetDeviceLiveToolsPing200ResponseResultsLoss type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeviceLiveToolsPing200ResponseResultsLoss{}

// GetDeviceLiveToolsPing200ResponseResultsLoss Lost packets
type GetDeviceLiveToolsPing200ResponseResultsLoss struct {
	// Percentage of packets lost
	Percentage *float32 `json:"percentage,omitempty"`
}

// NewGetDeviceLiveToolsPing200ResponseResultsLoss instantiates a new GetDeviceLiveToolsPing200ResponseResultsLoss object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeviceLiveToolsPing200ResponseResultsLoss() *GetDeviceLiveToolsPing200ResponseResultsLoss {
	this := GetDeviceLiveToolsPing200ResponseResultsLoss{}
	return &this
}

// NewGetDeviceLiveToolsPing200ResponseResultsLossWithDefaults instantiates a new GetDeviceLiveToolsPing200ResponseResultsLoss object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeviceLiveToolsPing200ResponseResultsLossWithDefaults() *GetDeviceLiveToolsPing200ResponseResultsLoss {
	this := GetDeviceLiveToolsPing200ResponseResultsLoss{}
	return &this
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *GetDeviceLiveToolsPing200ResponseResultsLoss) GetPercentage() float32 {
	if o == nil || IsNil(o.Percentage) {
		var ret float32
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceLiveToolsPing200ResponseResultsLoss) GetPercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.Percentage) {
		return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *GetDeviceLiveToolsPing200ResponseResultsLoss) HasPercentage() bool {
	if o != nil && !IsNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given float32 and assigns it to the Percentage field.
func (o *GetDeviceLiveToolsPing200ResponseResultsLoss) SetPercentage(v float32) {
	o.Percentage = &v
}

func (o GetDeviceLiveToolsPing200ResponseResultsLoss) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDeviceLiveToolsPing200ResponseResultsLoss) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	return toSerialize, nil
}

type NullableGetDeviceLiveToolsPing200ResponseResultsLoss struct {
	value *GetDeviceLiveToolsPing200ResponseResultsLoss
	isSet bool
}

func (v NullableGetDeviceLiveToolsPing200ResponseResultsLoss) Get() *GetDeviceLiveToolsPing200ResponseResultsLoss {
	return v.value
}

func (v *NullableGetDeviceLiveToolsPing200ResponseResultsLoss) Set(val *GetDeviceLiveToolsPing200ResponseResultsLoss) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeviceLiveToolsPing200ResponseResultsLoss) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeviceLiveToolsPing200ResponseResultsLoss) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeviceLiveToolsPing200ResponseResultsLoss(val *GetDeviceLiveToolsPing200ResponseResultsLoss) *NullableGetDeviceLiveToolsPing200ResponseResultsLoss {
	return &NullableGetDeviceLiveToolsPing200ResponseResultsLoss{value: val, isSet: true}
}

func (v NullableGetDeviceLiveToolsPing200ResponseResultsLoss) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeviceLiveToolsPing200ResponseResultsLoss) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


