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

// checks if the UpdateNetworkCameraWirelessProfileRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkCameraWirelessProfileRequest{}

// UpdateNetworkCameraWirelessProfileRequest struct for UpdateNetworkCameraWirelessProfileRequest
type UpdateNetworkCameraWirelessProfileRequest struct {
	// The name of the camera wireless profile.
	Name *string `json:"name,omitempty"`
	Ssid *CreateNetworkCameraWirelessProfileRequestSsid `json:"ssid,omitempty"`
	Identity *CreateNetworkCameraWirelessProfileRequestIdentity `json:"identity,omitempty"`
}

// NewUpdateNetworkCameraWirelessProfileRequest instantiates a new UpdateNetworkCameraWirelessProfileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkCameraWirelessProfileRequest() *UpdateNetworkCameraWirelessProfileRequest {
	this := UpdateNetworkCameraWirelessProfileRequest{}
	return &this
}

// NewUpdateNetworkCameraWirelessProfileRequestWithDefaults instantiates a new UpdateNetworkCameraWirelessProfileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkCameraWirelessProfileRequestWithDefaults() *UpdateNetworkCameraWirelessProfileRequest {
	this := UpdateNetworkCameraWirelessProfileRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateNetworkCameraWirelessProfileRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkCameraWirelessProfileRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateNetworkCameraWirelessProfileRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateNetworkCameraWirelessProfileRequest) SetName(v string) {
	o.Name = &v
}

// GetSsid returns the Ssid field value if set, zero value otherwise.
func (o *UpdateNetworkCameraWirelessProfileRequest) GetSsid() CreateNetworkCameraWirelessProfileRequestSsid {
	if o == nil || IsNil(o.Ssid) {
		var ret CreateNetworkCameraWirelessProfileRequestSsid
		return ret
	}
	return *o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkCameraWirelessProfileRequest) GetSsidOk() (*CreateNetworkCameraWirelessProfileRequestSsid, bool) {
	if o == nil || IsNil(o.Ssid) {
		return nil, false
	}
	return o.Ssid, true
}

// HasSsid returns a boolean if a field has been set.
func (o *UpdateNetworkCameraWirelessProfileRequest) HasSsid() bool {
	if o != nil && !IsNil(o.Ssid) {
		return true
	}

	return false
}

// SetSsid gets a reference to the given CreateNetworkCameraWirelessProfileRequestSsid and assigns it to the Ssid field.
func (o *UpdateNetworkCameraWirelessProfileRequest) SetSsid(v CreateNetworkCameraWirelessProfileRequestSsid) {
	o.Ssid = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *UpdateNetworkCameraWirelessProfileRequest) GetIdentity() CreateNetworkCameraWirelessProfileRequestIdentity {
	if o == nil || IsNil(o.Identity) {
		var ret CreateNetworkCameraWirelessProfileRequestIdentity
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkCameraWirelessProfileRequest) GetIdentityOk() (*CreateNetworkCameraWirelessProfileRequestIdentity, bool) {
	if o == nil || IsNil(o.Identity) {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *UpdateNetworkCameraWirelessProfileRequest) HasIdentity() bool {
	if o != nil && !IsNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given CreateNetworkCameraWirelessProfileRequestIdentity and assigns it to the Identity field.
func (o *UpdateNetworkCameraWirelessProfileRequest) SetIdentity(v CreateNetworkCameraWirelessProfileRequestIdentity) {
	o.Identity = &v
}

func (o UpdateNetworkCameraWirelessProfileRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkCameraWirelessProfileRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Ssid) {
		toSerialize["ssid"] = o.Ssid
	}
	if !IsNil(o.Identity) {
		toSerialize["identity"] = o.Identity
	}
	return toSerialize, nil
}

type NullableUpdateNetworkCameraWirelessProfileRequest struct {
	value *UpdateNetworkCameraWirelessProfileRequest
	isSet bool
}

func (v NullableUpdateNetworkCameraWirelessProfileRequest) Get() *UpdateNetworkCameraWirelessProfileRequest {
	return v.value
}

func (v *NullableUpdateNetworkCameraWirelessProfileRequest) Set(val *UpdateNetworkCameraWirelessProfileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkCameraWirelessProfileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkCameraWirelessProfileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkCameraWirelessProfileRequest(val *UpdateNetworkCameraWirelessProfileRequest) *NullableUpdateNetworkCameraWirelessProfileRequest {
	return &NullableUpdateNetworkCameraWirelessProfileRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkCameraWirelessProfileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkCameraWirelessProfileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


