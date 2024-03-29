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

// checks if the CreateNetworkMqttBrokerRequestSecurity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkMqttBrokerRequestSecurity{}

// CreateNetworkMqttBrokerRequestSecurity Security settings of the MQTT broker.
type CreateNetworkMqttBrokerRequestSecurity struct {
	// Security protocol of the MQTT broker.
	Mode *string `json:"mode,omitempty"`
	Security *CreateNetworkMqttBrokerRequestSecuritySecurity `json:"security,omitempty"`
}

// NewCreateNetworkMqttBrokerRequestSecurity instantiates a new CreateNetworkMqttBrokerRequestSecurity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkMqttBrokerRequestSecurity() *CreateNetworkMqttBrokerRequestSecurity {
	this := CreateNetworkMqttBrokerRequestSecurity{}
	return &this
}

// NewCreateNetworkMqttBrokerRequestSecurityWithDefaults instantiates a new CreateNetworkMqttBrokerRequestSecurity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkMqttBrokerRequestSecurityWithDefaults() *CreateNetworkMqttBrokerRequestSecurity {
	this := CreateNetworkMqttBrokerRequestSecurity{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *CreateNetworkMqttBrokerRequestSecurity) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkMqttBrokerRequestSecurity) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *CreateNetworkMqttBrokerRequestSecurity) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *CreateNetworkMqttBrokerRequestSecurity) SetMode(v string) {
	o.Mode = &v
}

// GetSecurity returns the Security field value if set, zero value otherwise.
func (o *CreateNetworkMqttBrokerRequestSecurity) GetSecurity() CreateNetworkMqttBrokerRequestSecuritySecurity {
	if o == nil || IsNil(o.Security) {
		var ret CreateNetworkMqttBrokerRequestSecuritySecurity
		return ret
	}
	return *o.Security
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkMqttBrokerRequestSecurity) GetSecurityOk() (*CreateNetworkMqttBrokerRequestSecuritySecurity, bool) {
	if o == nil || IsNil(o.Security) {
		return nil, false
	}
	return o.Security, true
}

// HasSecurity returns a boolean if a field has been set.
func (o *CreateNetworkMqttBrokerRequestSecurity) HasSecurity() bool {
	if o != nil && !IsNil(o.Security) {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given CreateNetworkMqttBrokerRequestSecuritySecurity and assigns it to the Security field.
func (o *CreateNetworkMqttBrokerRequestSecurity) SetSecurity(v CreateNetworkMqttBrokerRequestSecuritySecurity) {
	o.Security = &v
}

func (o CreateNetworkMqttBrokerRequestSecurity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkMqttBrokerRequestSecurity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.Security) {
		toSerialize["security"] = o.Security
	}
	return toSerialize, nil
}

type NullableCreateNetworkMqttBrokerRequestSecurity struct {
	value *CreateNetworkMqttBrokerRequestSecurity
	isSet bool
}

func (v NullableCreateNetworkMqttBrokerRequestSecurity) Get() *CreateNetworkMqttBrokerRequestSecurity {
	return v.value
}

func (v *NullableCreateNetworkMqttBrokerRequestSecurity) Set(val *CreateNetworkMqttBrokerRequestSecurity) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkMqttBrokerRequestSecurity) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkMqttBrokerRequestSecurity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkMqttBrokerRequestSecurity(val *CreateNetworkMqttBrokerRequestSecurity) *NullableCreateNetworkMqttBrokerRequestSecurity {
	return &NullableCreateNetworkMqttBrokerRequestSecurity{value: val, isSet: true}
}

func (v NullableCreateNetworkMqttBrokerRequestSecurity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkMqttBrokerRequestSecurity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


