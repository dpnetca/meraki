/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 July, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.23.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package meraki

import (
	"encoding/json"
)

// CreateNetworkWirelessRfProfile201ResponsePerSsidSettings14 Settings for SSID 14
type CreateNetworkWirelessRfProfile201ResponsePerSsidSettings14 struct {
	// Name of SSID
	Name *string `json:"name,omitempty"`
	// Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	MinBitrate *int32 `json:"minBitrate,omitempty"`
	// Choice between 'dual', '2.4ghz' or '5ghz'.
	BandOperationMode *string `json:"bandOperationMode,omitempty"`
	// Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	BandSteeringEnabled *bool `json:"bandSteeringEnabled,omitempty"`
}

// NewCreateNetworkWirelessRfProfile201ResponsePerSsidSettings14 instantiates a new CreateNetworkWirelessRfProfile201ResponsePerSsidSettings14 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkWirelessRfProfile201ResponsePerSsidSettings14() *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings14 {
	this := CreateNetworkWirelessRfProfile201ResponsePerSsidSettings14{}
	return &this
}

// NewCreateNetworkWirelessRfProfile201ResponsePerSsidSettings14WithDefaults instantiates a new CreateNetworkWirelessRfProfile201ResponsePerSsidSettings14 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkWirelessRfProfile201ResponsePerSsidSettings14WithDefaults() *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings14 {
	this := CreateNetworkWirelessRfProfile201ResponsePerSsidSettings14{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings14) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings14) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings14) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings14) SetName(v string) {
	o.Name = &v
}

// GetMinBitrate returns the MinBitrate field value if set, zero value otherwise.
func (o *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings14) GetMinBitrate() int32 {
	if o == nil || o.MinBitrate == nil {
		var ret int32
		return ret
	}
	return *o.MinBitrate
}

// GetMinBitrateOk returns a tuple with the MinBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings14) GetMinBitrateOk() (*int32, bool) {
	if o == nil || o.MinBitrate == nil {
		return nil, false
	}
	return o.MinBitrate, true
}

// HasMinBitrate returns a boolean if a field has been set.
func (o *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings14) HasMinBitrate() bool {
	if o != nil && o.MinBitrate != nil {
		return true
	}

	return false
}

// SetMinBitrate gets a reference to the given int32 and assigns it to the MinBitrate field.
func (o *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings14) SetMinBitrate(v int32) {
	o.MinBitrate = &v
}

// GetBandOperationMode returns the BandOperationMode field value if set, zero value otherwise.
func (o *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings14) GetBandOperationMode() string {
	if o == nil || o.BandOperationMode == nil {
		var ret string
		return ret
	}
	return *o.BandOperationMode
}

// GetBandOperationModeOk returns a tuple with the BandOperationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings14) GetBandOperationModeOk() (*string, bool) {
	if o == nil || o.BandOperationMode == nil {
		return nil, false
	}
	return o.BandOperationMode, true
}

// HasBandOperationMode returns a boolean if a field has been set.
func (o *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings14) HasBandOperationMode() bool {
	if o != nil && o.BandOperationMode != nil {
		return true
	}

	return false
}

// SetBandOperationMode gets a reference to the given string and assigns it to the BandOperationMode field.
func (o *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings14) SetBandOperationMode(v string) {
	o.BandOperationMode = &v
}

// GetBandSteeringEnabled returns the BandSteeringEnabled field value if set, zero value otherwise.
func (o *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings14) GetBandSteeringEnabled() bool {
	if o == nil || o.BandSteeringEnabled == nil {
		var ret bool
		return ret
	}
	return *o.BandSteeringEnabled
}

// GetBandSteeringEnabledOk returns a tuple with the BandSteeringEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings14) GetBandSteeringEnabledOk() (*bool, bool) {
	if o == nil || o.BandSteeringEnabled == nil {
		return nil, false
	}
	return o.BandSteeringEnabled, true
}

// HasBandSteeringEnabled returns a boolean if a field has been set.
func (o *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings14) HasBandSteeringEnabled() bool {
	if o != nil && o.BandSteeringEnabled != nil {
		return true
	}

	return false
}

// SetBandSteeringEnabled gets a reference to the given bool and assigns it to the BandSteeringEnabled field.
func (o *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings14) SetBandSteeringEnabled(v bool) {
	o.BandSteeringEnabled = &v
}

func (o CreateNetworkWirelessRfProfile201ResponsePerSsidSettings14) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.MinBitrate != nil {
		toSerialize["minBitrate"] = o.MinBitrate
	}
	if o.BandOperationMode != nil {
		toSerialize["bandOperationMode"] = o.BandOperationMode
	}
	if o.BandSteeringEnabled != nil {
		toSerialize["bandSteeringEnabled"] = o.BandSteeringEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableCreateNetworkWirelessRfProfile201ResponsePerSsidSettings14 struct {
	value *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings14
	isSet bool
}

func (v NullableCreateNetworkWirelessRfProfile201ResponsePerSsidSettings14) Get() *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings14 {
	return v.value
}

func (v *NullableCreateNetworkWirelessRfProfile201ResponsePerSsidSettings14) Set(val *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings14) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkWirelessRfProfile201ResponsePerSsidSettings14) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkWirelessRfProfile201ResponsePerSsidSettings14) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkWirelessRfProfile201ResponsePerSsidSettings14(val *CreateNetworkWirelessRfProfile201ResponsePerSsidSettings14) *NullableCreateNetworkWirelessRfProfile201ResponsePerSsidSettings14 {
	return &NullableCreateNetworkWirelessRfProfile201ResponsePerSsidSettings14{value: val, isSet: true}
}

func (v NullableCreateNetworkWirelessRfProfile201ResponsePerSsidSettings14) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkWirelessRfProfile201ResponsePerSsidSettings14) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


