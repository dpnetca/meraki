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

// GetAdministeredIdentitiesMe200ResponseAuthentication Authentication info
type GetAdministeredIdentitiesMe200ResponseAuthentication struct {
	// Authentication mode
	Mode *string `json:"mode,omitempty"`
	Api *GetAdministeredIdentitiesMe200ResponseAuthenticationApi `json:"api,omitempty"`
	TwoFactor *GetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactor `json:"twoFactor,omitempty"`
	Saml *GetAdministeredIdentitiesMe200ResponseAuthenticationSaml `json:"saml,omitempty"`
}

// NewGetAdministeredIdentitiesMe200ResponseAuthentication instantiates a new GetAdministeredIdentitiesMe200ResponseAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAdministeredIdentitiesMe200ResponseAuthentication() *GetAdministeredIdentitiesMe200ResponseAuthentication {
	this := GetAdministeredIdentitiesMe200ResponseAuthentication{}
	return &this
}

// NewGetAdministeredIdentitiesMe200ResponseAuthenticationWithDefaults instantiates a new GetAdministeredIdentitiesMe200ResponseAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAdministeredIdentitiesMe200ResponseAuthenticationWithDefaults() *GetAdministeredIdentitiesMe200ResponseAuthentication {
	this := GetAdministeredIdentitiesMe200ResponseAuthentication{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *GetAdministeredIdentitiesMe200ResponseAuthentication) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredIdentitiesMe200ResponseAuthentication) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *GetAdministeredIdentitiesMe200ResponseAuthentication) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *GetAdministeredIdentitiesMe200ResponseAuthentication) SetMode(v string) {
	o.Mode = &v
}

// GetApi returns the Api field value if set, zero value otherwise.
func (o *GetAdministeredIdentitiesMe200ResponseAuthentication) GetApi() GetAdministeredIdentitiesMe200ResponseAuthenticationApi {
	if o == nil || o.Api == nil {
		var ret GetAdministeredIdentitiesMe200ResponseAuthenticationApi
		return ret
	}
	return *o.Api
}

// GetApiOk returns a tuple with the Api field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredIdentitiesMe200ResponseAuthentication) GetApiOk() (*GetAdministeredIdentitiesMe200ResponseAuthenticationApi, bool) {
	if o == nil || o.Api == nil {
		return nil, false
	}
	return o.Api, true
}

// HasApi returns a boolean if a field has been set.
func (o *GetAdministeredIdentitiesMe200ResponseAuthentication) HasApi() bool {
	if o != nil && o.Api != nil {
		return true
	}

	return false
}

// SetApi gets a reference to the given GetAdministeredIdentitiesMe200ResponseAuthenticationApi and assigns it to the Api field.
func (o *GetAdministeredIdentitiesMe200ResponseAuthentication) SetApi(v GetAdministeredIdentitiesMe200ResponseAuthenticationApi) {
	o.Api = &v
}

// GetTwoFactor returns the TwoFactor field value if set, zero value otherwise.
func (o *GetAdministeredIdentitiesMe200ResponseAuthentication) GetTwoFactor() GetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactor {
	if o == nil || o.TwoFactor == nil {
		var ret GetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactor
		return ret
	}
	return *o.TwoFactor
}

// GetTwoFactorOk returns a tuple with the TwoFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredIdentitiesMe200ResponseAuthentication) GetTwoFactorOk() (*GetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactor, bool) {
	if o == nil || o.TwoFactor == nil {
		return nil, false
	}
	return o.TwoFactor, true
}

// HasTwoFactor returns a boolean if a field has been set.
func (o *GetAdministeredIdentitiesMe200ResponseAuthentication) HasTwoFactor() bool {
	if o != nil && o.TwoFactor != nil {
		return true
	}

	return false
}

// SetTwoFactor gets a reference to the given GetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactor and assigns it to the TwoFactor field.
func (o *GetAdministeredIdentitiesMe200ResponseAuthentication) SetTwoFactor(v GetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactor) {
	o.TwoFactor = &v
}

// GetSaml returns the Saml field value if set, zero value otherwise.
func (o *GetAdministeredIdentitiesMe200ResponseAuthentication) GetSaml() GetAdministeredIdentitiesMe200ResponseAuthenticationSaml {
	if o == nil || o.Saml == nil {
		var ret GetAdministeredIdentitiesMe200ResponseAuthenticationSaml
		return ret
	}
	return *o.Saml
}

// GetSamlOk returns a tuple with the Saml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredIdentitiesMe200ResponseAuthentication) GetSamlOk() (*GetAdministeredIdentitiesMe200ResponseAuthenticationSaml, bool) {
	if o == nil || o.Saml == nil {
		return nil, false
	}
	return o.Saml, true
}

// HasSaml returns a boolean if a field has been set.
func (o *GetAdministeredIdentitiesMe200ResponseAuthentication) HasSaml() bool {
	if o != nil && o.Saml != nil {
		return true
	}

	return false
}

// SetSaml gets a reference to the given GetAdministeredIdentitiesMe200ResponseAuthenticationSaml and assigns it to the Saml field.
func (o *GetAdministeredIdentitiesMe200ResponseAuthentication) SetSaml(v GetAdministeredIdentitiesMe200ResponseAuthenticationSaml) {
	o.Saml = &v
}

func (o GetAdministeredIdentitiesMe200ResponseAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.Api != nil {
		toSerialize["api"] = o.Api
	}
	if o.TwoFactor != nil {
		toSerialize["twoFactor"] = o.TwoFactor
	}
	if o.Saml != nil {
		toSerialize["saml"] = o.Saml
	}
	return json.Marshal(toSerialize)
}

type NullableGetAdministeredIdentitiesMe200ResponseAuthentication struct {
	value *GetAdministeredIdentitiesMe200ResponseAuthentication
	isSet bool
}

func (v NullableGetAdministeredIdentitiesMe200ResponseAuthentication) Get() *GetAdministeredIdentitiesMe200ResponseAuthentication {
	return v.value
}

func (v *NullableGetAdministeredIdentitiesMe200ResponseAuthentication) Set(val *GetAdministeredIdentitiesMe200ResponseAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAdministeredIdentitiesMe200ResponseAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAdministeredIdentitiesMe200ResponseAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAdministeredIdentitiesMe200ResponseAuthentication(val *GetAdministeredIdentitiesMe200ResponseAuthentication) *NullableGetAdministeredIdentitiesMe200ResponseAuthentication {
	return &NullableGetAdministeredIdentitiesMe200ResponseAuthentication{value: val, isSet: true}
}

func (v NullableGetAdministeredIdentitiesMe200ResponseAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAdministeredIdentitiesMe200ResponseAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

