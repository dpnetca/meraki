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

// checks if the CreateNetworkWirelessRfProfile201ResponsePerSsidSettings0 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkWirelessRfProfile201ResponsePerSsidSettings0{}

// CreateNetworkWirelessRfProfile201ResponsePerSsidSettings0 Settings for SSID 0
type CreateNetworkWirelessRfProfile201ResponsePerSsidSettings0 struct {
	// Name of SSID
	Name *string `json:"name,omitempty"`
	// Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	MinBitrate *int32 `json:"minBitrate,omitempty"`
	// Choice between 'dual', '2.4ghz' or '5ghz'.
	BandOperationMode *string `json:"bandOperationMode,omitempty"`
	// Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	BandSteeringEnabled *bool `json:"bandSteeringEnabled,omitempty"`
}

// NewCreateNetworkWirelessRfProfile201ResponsePerSsidSettings0 instantiates a new CreateNetworkWirelessRfProfile201ResponsePerSsidSettings0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkWirelessRfProfile201ResponsePerSsidSettings0() *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings0 {
	this := CreateNetworkWirelessRfProfile201ResponsePerSsidSettings0{}
	return &this
}

// NewCreateNetworkWirelessRfProfile201ResponsePerSsidSettings0WithDefaults instantiates a new CreateNetworkWirelessRfProfile201ResponsePerSsidSettings0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkWirelessRfProfile201ResponsePerSsidSettings0WithDefaults() *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings0 {
	this := CreateNetworkWirelessRfProfile201ResponsePerSsidSettings0{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings0) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings0) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings0) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings0) SetName(v string) {
	o.Name = &v
}

// GetMinBitrate returns the MinBitrate field value if set, zero value otherwise.
func (o *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings0) GetMinBitrate() int32 {
	if o == nil || IsNil(o.MinBitrate) {
		var ret int32
		return ret
	}
	return *o.MinBitrate
}

// GetMinBitrateOk returns a tuple with the MinBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings0) GetMinBitrateOk() (*int32, bool) {
	if o == nil || IsNil(o.MinBitrate) {
		return nil, false
	}
	return o.MinBitrate, true
}

// HasMinBitrate returns a boolean if a field has been set.
func (o *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings0) HasMinBitrate() bool {
	if o != nil && !IsNil(o.MinBitrate) {
		return true
	}

	return false
}

// SetMinBitrate gets a reference to the given int32 and assigns it to the MinBitrate field.
func (o *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings0) SetMinBitrate(v int32) {
	o.MinBitrate = &v
}

// GetBandOperationMode returns the BandOperationMode field value if set, zero value otherwise.
func (o *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings0) GetBandOperationMode() string {
	if o == nil || IsNil(o.BandOperationMode) {
		var ret string
		return ret
	}
	return *o.BandOperationMode
}

// GetBandOperationModeOk returns a tuple with the BandOperationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings0) GetBandOperationModeOk() (*string, bool) {
	if o == nil || IsNil(o.BandOperationMode) {
		return nil, false
	}
	return o.BandOperationMode, true
}

// HasBandOperationMode returns a boolean if a field has been set.
func (o *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings0) HasBandOperationMode() bool {
	if o != nil && !IsNil(o.BandOperationMode) {
		return true
	}

	return false
}

// SetBandOperationMode gets a reference to the given string and assigns it to the BandOperationMode field.
func (o *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings0) SetBandOperationMode(v string) {
	o.BandOperationMode = &v
}

// GetBandSteeringEnabled returns the BandSteeringEnabled field value if set, zero value otherwise.
func (o *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings0) GetBandSteeringEnabled() bool {
	if o == nil || IsNil(o.BandSteeringEnabled) {
		var ret bool
		return ret
	}
	return *o.BandSteeringEnabled
}

// GetBandSteeringEnabledOk returns a tuple with the BandSteeringEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings0) GetBandSteeringEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BandSteeringEnabled) {
		return nil, false
	}
	return o.BandSteeringEnabled, true
}

// HasBandSteeringEnabled returns a boolean if a field has been set.
func (o *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings0) HasBandSteeringEnabled() bool {
	if o != nil && !IsNil(o.BandSteeringEnabled) {
		return true
	}

	return false
}

// SetBandSteeringEnabled gets a reference to the given bool and assigns it to the BandSteeringEnabled field.
func (o *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings0) SetBandSteeringEnabled(v bool) {
	o.BandSteeringEnabled = &v
}

func (o CreateNetworkWirelessRfProfile201ResponsePerSsidSettings0) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkWirelessRfProfile201ResponsePerSsidSettings0) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.MinBitrate) {
		toSerialize["minBitrate"] = o.MinBitrate
	}
	if !IsNil(o.BandOperationMode) {
		toSerialize["bandOperationMode"] = o.BandOperationMode
	}
	if !IsNil(o.BandSteeringEnabled) {
		toSerialize["bandSteeringEnabled"] = o.BandSteeringEnabled
	}
	return toSerialize, nil
}

type NullableCreateNetworkWirelessRfProfile201ResponsePerSsidSettings0 struct {
	value *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings0
	isSet bool
}

func (v NullableCreateNetworkWirelessRfProfile201ResponsePerSsidSettings0) Get() *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings0 {
	return v.value
}

func (v *NullableCreateNetworkWirelessRfProfile201ResponsePerSsidSettings0) Set(val *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings0) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkWirelessRfProfile201ResponsePerSsidSettings0) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkWirelessRfProfile201ResponsePerSsidSettings0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkWirelessRfProfile201ResponsePerSsidSettings0(val *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings0) *NullableCreateNetworkWirelessRfProfile201ResponsePerSsidSettings0 {
	return &NullableCreateNetworkWirelessRfProfile201ResponsePerSsidSettings0{value: val, isSet: true}
}

func (v NullableCreateNetworkWirelessRfProfile201ResponsePerSsidSettings0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkWirelessRfProfile201ResponsePerSsidSettings0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


