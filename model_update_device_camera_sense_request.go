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

// checks if the UpdateDeviceCameraSenseRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeviceCameraSenseRequest{}

// UpdateDeviceCameraSenseRequest struct for UpdateDeviceCameraSenseRequest
type UpdateDeviceCameraSenseRequest struct {
	// Boolean indicating if sense(license) is enabled(true) or disabled(false) on the camera
	SenseEnabled *bool `json:"senseEnabled,omitempty"`
	// The ID of the MQTT broker to be enabled on the camera. A value of null will disable MQTT on the camera
	MqttBrokerId *string `json:"mqttBrokerId,omitempty"`
	AudioDetection *UpdateDeviceCameraSenseRequestAudioDetection `json:"audioDetection,omitempty"`
	// The ID of the object detection model
	DetectionModelId *string `json:"detectionModelId,omitempty"`
}

// NewUpdateDeviceCameraSenseRequest instantiates a new UpdateDeviceCameraSenseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceCameraSenseRequest() *UpdateDeviceCameraSenseRequest {
	this := UpdateDeviceCameraSenseRequest{}
	return &this
}

// NewUpdateDeviceCameraSenseRequestWithDefaults instantiates a new UpdateDeviceCameraSenseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceCameraSenseRequestWithDefaults() *UpdateDeviceCameraSenseRequest {
	this := UpdateDeviceCameraSenseRequest{}
	return &this
}

// GetSenseEnabled returns the SenseEnabled field value if set, zero value otherwise.
func (o *UpdateDeviceCameraSenseRequest) GetSenseEnabled() bool {
	if o == nil || IsNil(o.SenseEnabled) {
		var ret bool
		return ret
	}
	return *o.SenseEnabled
}

// GetSenseEnabledOk returns a tuple with the SenseEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceCameraSenseRequest) GetSenseEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SenseEnabled) {
		return nil, false
	}
	return o.SenseEnabled, true
}

// HasSenseEnabled returns a boolean if a field has been set.
func (o *UpdateDeviceCameraSenseRequest) HasSenseEnabled() bool {
	if o != nil && !IsNil(o.SenseEnabled) {
		return true
	}

	return false
}

// SetSenseEnabled gets a reference to the given bool and assigns it to the SenseEnabled field.
func (o *UpdateDeviceCameraSenseRequest) SetSenseEnabled(v bool) {
	o.SenseEnabled = &v
}

// GetMqttBrokerId returns the MqttBrokerId field value if set, zero value otherwise.
func (o *UpdateDeviceCameraSenseRequest) GetMqttBrokerId() string {
	if o == nil || IsNil(o.MqttBrokerId) {
		var ret string
		return ret
	}
	return *o.MqttBrokerId
}

// GetMqttBrokerIdOk returns a tuple with the MqttBrokerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceCameraSenseRequest) GetMqttBrokerIdOk() (*string, bool) {
	if o == nil || IsNil(o.MqttBrokerId) {
		return nil, false
	}
	return o.MqttBrokerId, true
}

// HasMqttBrokerId returns a boolean if a field has been set.
func (o *UpdateDeviceCameraSenseRequest) HasMqttBrokerId() bool {
	if o != nil && !IsNil(o.MqttBrokerId) {
		return true
	}

	return false
}

// SetMqttBrokerId gets a reference to the given string and assigns it to the MqttBrokerId field.
func (o *UpdateDeviceCameraSenseRequest) SetMqttBrokerId(v string) {
	o.MqttBrokerId = &v
}

// GetAudioDetection returns the AudioDetection field value if set, zero value otherwise.
func (o *UpdateDeviceCameraSenseRequest) GetAudioDetection() UpdateDeviceCameraSenseRequestAudioDetection {
	if o == nil || IsNil(o.AudioDetection) {
		var ret UpdateDeviceCameraSenseRequestAudioDetection
		return ret
	}
	return *o.AudioDetection
}

// GetAudioDetectionOk returns a tuple with the AudioDetection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceCameraSenseRequest) GetAudioDetectionOk() (*UpdateDeviceCameraSenseRequestAudioDetection, bool) {
	if o == nil || IsNil(o.AudioDetection) {
		return nil, false
	}
	return o.AudioDetection, true
}

// HasAudioDetection returns a boolean if a field has been set.
func (o *UpdateDeviceCameraSenseRequest) HasAudioDetection() bool {
	if o != nil && !IsNil(o.AudioDetection) {
		return true
	}

	return false
}

// SetAudioDetection gets a reference to the given UpdateDeviceCameraSenseRequestAudioDetection and assigns it to the AudioDetection field.
func (o *UpdateDeviceCameraSenseRequest) SetAudioDetection(v UpdateDeviceCameraSenseRequestAudioDetection) {
	o.AudioDetection = &v
}

// GetDetectionModelId returns the DetectionModelId field value if set, zero value otherwise.
func (o *UpdateDeviceCameraSenseRequest) GetDetectionModelId() string {
	if o == nil || IsNil(o.DetectionModelId) {
		var ret string
		return ret
	}
	return *o.DetectionModelId
}

// GetDetectionModelIdOk returns a tuple with the DetectionModelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceCameraSenseRequest) GetDetectionModelIdOk() (*string, bool) {
	if o == nil || IsNil(o.DetectionModelId) {
		return nil, false
	}
	return o.DetectionModelId, true
}

// HasDetectionModelId returns a boolean if a field has been set.
func (o *UpdateDeviceCameraSenseRequest) HasDetectionModelId() bool {
	if o != nil && !IsNil(o.DetectionModelId) {
		return true
	}

	return false
}

// SetDetectionModelId gets a reference to the given string and assigns it to the DetectionModelId field.
func (o *UpdateDeviceCameraSenseRequest) SetDetectionModelId(v string) {
	o.DetectionModelId = &v
}

func (o UpdateDeviceCameraSenseRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeviceCameraSenseRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SenseEnabled) {
		toSerialize["senseEnabled"] = o.SenseEnabled
	}
	if !IsNil(o.MqttBrokerId) {
		toSerialize["mqttBrokerId"] = o.MqttBrokerId
	}
	if !IsNil(o.AudioDetection) {
		toSerialize["audioDetection"] = o.AudioDetection
	}
	if !IsNil(o.DetectionModelId) {
		toSerialize["detectionModelId"] = o.DetectionModelId
	}
	return toSerialize, nil
}

type NullableUpdateDeviceCameraSenseRequest struct {
	value *UpdateDeviceCameraSenseRequest
	isSet bool
}

func (v NullableUpdateDeviceCameraSenseRequest) Get() *UpdateDeviceCameraSenseRequest {
	return v.value
}

func (v *NullableUpdateDeviceCameraSenseRequest) Set(val *UpdateDeviceCameraSenseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceCameraSenseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceCameraSenseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceCameraSenseRequest(val *UpdateDeviceCameraSenseRequest) *NullableUpdateDeviceCameraSenseRequest {
	return &NullableUpdateDeviceCameraSenseRequest{value: val, isSet: true}
}

func (v NullableUpdateDeviceCameraSenseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceCameraSenseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


