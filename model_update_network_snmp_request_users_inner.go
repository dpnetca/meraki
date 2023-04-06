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

// checks if the UpdateNetworkSnmpRequestUsersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkSnmpRequestUsersInner{}

// UpdateNetworkSnmpRequestUsersInner struct for UpdateNetworkSnmpRequestUsersInner
type UpdateNetworkSnmpRequestUsersInner struct {
	// The username for the SNMP user. Required.
	Username string `json:"username"`
	// The passphrase for the SNMP user. Required.
	Passphrase string `json:"passphrase"`
}

// NewUpdateNetworkSnmpRequestUsersInner instantiates a new UpdateNetworkSnmpRequestUsersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSnmpRequestUsersInner(username string, passphrase string) *UpdateNetworkSnmpRequestUsersInner {
	this := UpdateNetworkSnmpRequestUsersInner{}
	this.Username = username
	this.Passphrase = passphrase
	return &this
}

// NewUpdateNetworkSnmpRequestUsersInnerWithDefaults instantiates a new UpdateNetworkSnmpRequestUsersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSnmpRequestUsersInnerWithDefaults() *UpdateNetworkSnmpRequestUsersInner {
	this := UpdateNetworkSnmpRequestUsersInner{}
	return &this
}

// GetUsername returns the Username field value
func (o *UpdateNetworkSnmpRequestUsersInner) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSnmpRequestUsersInner) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *UpdateNetworkSnmpRequestUsersInner) SetUsername(v string) {
	o.Username = v
}

// GetPassphrase returns the Passphrase field value
func (o *UpdateNetworkSnmpRequestUsersInner) GetPassphrase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSnmpRequestUsersInner) GetPassphraseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Passphrase, true
}

// SetPassphrase sets field value
func (o *UpdateNetworkSnmpRequestUsersInner) SetPassphrase(v string) {
	o.Passphrase = v
}

func (o UpdateNetworkSnmpRequestUsersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkSnmpRequestUsersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["username"] = o.Username
	toSerialize["passphrase"] = o.Passphrase
	return toSerialize, nil
}

type NullableUpdateNetworkSnmpRequestUsersInner struct {
	value *UpdateNetworkSnmpRequestUsersInner
	isSet bool
}

func (v NullableUpdateNetworkSnmpRequestUsersInner) Get() *UpdateNetworkSnmpRequestUsersInner {
	return v.value
}

func (v *NullableUpdateNetworkSnmpRequestUsersInner) Set(val *UpdateNetworkSnmpRequestUsersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSnmpRequestUsersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSnmpRequestUsersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSnmpRequestUsersInner(val *UpdateNetworkSnmpRequestUsersInner) *NullableUpdateNetworkSnmpRequestUsersInner {
	return &NullableUpdateNetworkSnmpRequestUsersInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkSnmpRequestUsersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSnmpRequestUsersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


