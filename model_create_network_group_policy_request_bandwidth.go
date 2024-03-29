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

// checks if the CreateNetworkGroupPolicyRequestBandwidth type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkGroupPolicyRequestBandwidth{}

// CreateNetworkGroupPolicyRequestBandwidth     The bandwidth settings for clients bound to your group policy. 
type CreateNetworkGroupPolicyRequestBandwidth struct {
	// How bandwidth limits are enforced. Can be 'network default', 'ignore' or 'custom'.
	Settings *string `json:"settings,omitempty"`
	BandwidthLimits *CreateNetworkGroupPolicyRequestBandwidthBandwidthLimits `json:"bandwidthLimits,omitempty"`
}

// NewCreateNetworkGroupPolicyRequestBandwidth instantiates a new CreateNetworkGroupPolicyRequestBandwidth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkGroupPolicyRequestBandwidth() *CreateNetworkGroupPolicyRequestBandwidth {
	this := CreateNetworkGroupPolicyRequestBandwidth{}
	return &this
}

// NewCreateNetworkGroupPolicyRequestBandwidthWithDefaults instantiates a new CreateNetworkGroupPolicyRequestBandwidth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkGroupPolicyRequestBandwidthWithDefaults() *CreateNetworkGroupPolicyRequestBandwidth {
	this := CreateNetworkGroupPolicyRequestBandwidth{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *CreateNetworkGroupPolicyRequestBandwidth) GetSettings() string {
	if o == nil || IsNil(o.Settings) {
		var ret string
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequestBandwidth) GetSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *CreateNetworkGroupPolicyRequestBandwidth) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given string and assigns it to the Settings field.
func (o *CreateNetworkGroupPolicyRequestBandwidth) SetSettings(v string) {
	o.Settings = &v
}

// GetBandwidthLimits returns the BandwidthLimits field value if set, zero value otherwise.
func (o *CreateNetworkGroupPolicyRequestBandwidth) GetBandwidthLimits() CreateNetworkGroupPolicyRequestBandwidthBandwidthLimits {
	if o == nil || IsNil(o.BandwidthLimits) {
		var ret CreateNetworkGroupPolicyRequestBandwidthBandwidthLimits
		return ret
	}
	return *o.BandwidthLimits
}

// GetBandwidthLimitsOk returns a tuple with the BandwidthLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequestBandwidth) GetBandwidthLimitsOk() (*CreateNetworkGroupPolicyRequestBandwidthBandwidthLimits, bool) {
	if o == nil || IsNil(o.BandwidthLimits) {
		return nil, false
	}
	return o.BandwidthLimits, true
}

// HasBandwidthLimits returns a boolean if a field has been set.
func (o *CreateNetworkGroupPolicyRequestBandwidth) HasBandwidthLimits() bool {
	if o != nil && !IsNil(o.BandwidthLimits) {
		return true
	}

	return false
}

// SetBandwidthLimits gets a reference to the given CreateNetworkGroupPolicyRequestBandwidthBandwidthLimits and assigns it to the BandwidthLimits field.
func (o *CreateNetworkGroupPolicyRequestBandwidth) SetBandwidthLimits(v CreateNetworkGroupPolicyRequestBandwidthBandwidthLimits) {
	o.BandwidthLimits = &v
}

func (o CreateNetworkGroupPolicyRequestBandwidth) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkGroupPolicyRequestBandwidth) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !IsNil(o.BandwidthLimits) {
		toSerialize["bandwidthLimits"] = o.BandwidthLimits
	}
	return toSerialize, nil
}

type NullableCreateNetworkGroupPolicyRequestBandwidth struct {
	value *CreateNetworkGroupPolicyRequestBandwidth
	isSet bool
}

func (v NullableCreateNetworkGroupPolicyRequestBandwidth) Get() *CreateNetworkGroupPolicyRequestBandwidth {
	return v.value
}

func (v *NullableCreateNetworkGroupPolicyRequestBandwidth) Set(val *CreateNetworkGroupPolicyRequestBandwidth) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkGroupPolicyRequestBandwidth) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkGroupPolicyRequestBandwidth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkGroupPolicyRequestBandwidth(val *CreateNetworkGroupPolicyRequestBandwidth) *NullableCreateNetworkGroupPolicyRequestBandwidth {
	return &NullableCreateNetworkGroupPolicyRequestBandwidth{value: val, isSet: true}
}

func (v NullableCreateNetworkGroupPolicyRequestBandwidth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkGroupPolicyRequestBandwidth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


