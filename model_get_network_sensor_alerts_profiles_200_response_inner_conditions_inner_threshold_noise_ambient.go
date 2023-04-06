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

// checks if the GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient{}

// GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient Ambient noise threshold. One of 'level' or 'quality' must be provided.
type GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient struct {
	// Alerting threshold as adjusted decibels.
	Level *int32 `json:"level,omitempty"`
	// Alerting threshold as a qualitative ambient noise level.
	Quality *string `json:"quality,omitempty"`
}

// NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient instantiates a new GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient() *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient {
	this := GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient{}
	return &this
}

// NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbientWithDefaults instantiates a new GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbientWithDefaults() *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient {
	this := GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient{}
	return &this
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient) GetLevel() int32 {
	if o == nil || IsNil(o.Level) {
		var ret int32
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient) GetLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given int32 and assigns it to the Level field.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient) SetLevel(v int32) {
	o.Level = &v
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient) GetQuality() string {
	if o == nil || IsNil(o.Quality) {
		var ret string
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient) GetQualityOk() (*string, bool) {
	if o == nil || IsNil(o.Quality) {
		return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient) HasQuality() bool {
	if o != nil && !IsNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given string and assigns it to the Quality field.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient) SetQuality(v string) {
	o.Quality = &v
}

func (o GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	if !IsNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	return toSerialize, nil
}

type NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient struct {
	value *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient
	isSet bool
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient) Get() *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient {
	return v.value
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient) Set(val *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient(val *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient) *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient {
	return &NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient{value: val, isSet: true}
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


