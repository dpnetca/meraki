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

// GetOrganizationLoginSecurity200ResponseApiAuthentication Details for indicating whether organization will restrict access to API (but not Dashboard) to certain IP addresses.
type GetOrganizationLoginSecurity200ResponseApiAuthentication struct {
	IpRestrictionsForKeys *GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys `json:"ipRestrictionsForKeys,omitempty"`
}

// NewGetOrganizationLoginSecurity200ResponseApiAuthentication instantiates a new GetOrganizationLoginSecurity200ResponseApiAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationLoginSecurity200ResponseApiAuthentication() *GetOrganizationLoginSecurity200ResponseApiAuthentication {
	this := GetOrganizationLoginSecurity200ResponseApiAuthentication{}
	return &this
}

// NewGetOrganizationLoginSecurity200ResponseApiAuthenticationWithDefaults instantiates a new GetOrganizationLoginSecurity200ResponseApiAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationLoginSecurity200ResponseApiAuthenticationWithDefaults() *GetOrganizationLoginSecurity200ResponseApiAuthentication {
	this := GetOrganizationLoginSecurity200ResponseApiAuthentication{}
	return &this
}

// GetIpRestrictionsForKeys returns the IpRestrictionsForKeys field value if set, zero value otherwise.
func (o *GetOrganizationLoginSecurity200ResponseApiAuthentication) GetIpRestrictionsForKeys() GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys {
	if o == nil || o.IpRestrictionsForKeys == nil {
		var ret GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys
		return ret
	}
	return *o.IpRestrictionsForKeys
}

// GetIpRestrictionsForKeysOk returns a tuple with the IpRestrictionsForKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLoginSecurity200ResponseApiAuthentication) GetIpRestrictionsForKeysOk() (*GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys, bool) {
	if o == nil || o.IpRestrictionsForKeys == nil {
		return nil, false
	}
	return o.IpRestrictionsForKeys, true
}

// HasIpRestrictionsForKeys returns a boolean if a field has been set.
func (o *GetOrganizationLoginSecurity200ResponseApiAuthentication) HasIpRestrictionsForKeys() bool {
	if o != nil && o.IpRestrictionsForKeys != nil {
		return true
	}

	return false
}

// SetIpRestrictionsForKeys gets a reference to the given GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys and assigns it to the IpRestrictionsForKeys field.
func (o *GetOrganizationLoginSecurity200ResponseApiAuthentication) SetIpRestrictionsForKeys(v GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys) {
	o.IpRestrictionsForKeys = &v
}

func (o GetOrganizationLoginSecurity200ResponseApiAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IpRestrictionsForKeys != nil {
		toSerialize["ipRestrictionsForKeys"] = o.IpRestrictionsForKeys
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrganizationLoginSecurity200ResponseApiAuthentication struct {
	value *GetOrganizationLoginSecurity200ResponseApiAuthentication
	isSet bool
}

func (v NullableGetOrganizationLoginSecurity200ResponseApiAuthentication) Get() *GetOrganizationLoginSecurity200ResponseApiAuthentication {
	return v.value
}

func (v *NullableGetOrganizationLoginSecurity200ResponseApiAuthentication) Set(val *GetOrganizationLoginSecurity200ResponseApiAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationLoginSecurity200ResponseApiAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationLoginSecurity200ResponseApiAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationLoginSecurity200ResponseApiAuthentication(val *GetOrganizationLoginSecurity200ResponseApiAuthentication) *NullableGetOrganizationLoginSecurity200ResponseApiAuthentication {
	return &NullableGetOrganizationLoginSecurity200ResponseApiAuthentication{value: val, isSet: true}
}

func (v NullableGetOrganizationLoginSecurity200ResponseApiAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationLoginSecurity200ResponseApiAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


