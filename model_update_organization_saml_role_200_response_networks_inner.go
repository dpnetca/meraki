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

// checks if the UpdateOrganizationSamlRole200ResponseNetworksInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrganizationSamlRole200ResponseNetworksInner{}

// UpdateOrganizationSamlRole200ResponseNetworksInner struct for UpdateOrganizationSamlRole200ResponseNetworksInner
type UpdateOrganizationSamlRole200ResponseNetworksInner struct {
	// The network ID
	Id *string `json:"id,omitempty"`
	// The privilege of the SAML administrator on the network
	Access *string `json:"access,omitempty"`
}

// NewUpdateOrganizationSamlRole200ResponseNetworksInner instantiates a new UpdateOrganizationSamlRole200ResponseNetworksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationSamlRole200ResponseNetworksInner() *UpdateOrganizationSamlRole200ResponseNetworksInner {
	this := UpdateOrganizationSamlRole200ResponseNetworksInner{}
	return &this
}

// NewUpdateOrganizationSamlRole200ResponseNetworksInnerWithDefaults instantiates a new UpdateOrganizationSamlRole200ResponseNetworksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationSamlRole200ResponseNetworksInnerWithDefaults() *UpdateOrganizationSamlRole200ResponseNetworksInner {
	this := UpdateOrganizationSamlRole200ResponseNetworksInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateOrganizationSamlRole200ResponseNetworksInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationSamlRole200ResponseNetworksInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateOrganizationSamlRole200ResponseNetworksInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UpdateOrganizationSamlRole200ResponseNetworksInner) SetId(v string) {
	o.Id = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *UpdateOrganizationSamlRole200ResponseNetworksInner) GetAccess() string {
	if o == nil || IsNil(o.Access) {
		var ret string
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationSamlRole200ResponseNetworksInner) GetAccessOk() (*string, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *UpdateOrganizationSamlRole200ResponseNetworksInner) HasAccess() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given string and assigns it to the Access field.
func (o *UpdateOrganizationSamlRole200ResponseNetworksInner) SetAccess(v string) {
	o.Access = &v
}

func (o UpdateOrganizationSamlRole200ResponseNetworksInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrganizationSamlRole200ResponseNetworksInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	return toSerialize, nil
}

type NullableUpdateOrganizationSamlRole200ResponseNetworksInner struct {
	value *UpdateOrganizationSamlRole200ResponseNetworksInner
	isSet bool
}

func (v NullableUpdateOrganizationSamlRole200ResponseNetworksInner) Get() *UpdateOrganizationSamlRole200ResponseNetworksInner {
	return v.value
}

func (v *NullableUpdateOrganizationSamlRole200ResponseNetworksInner) Set(val *UpdateOrganizationSamlRole200ResponseNetworksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationSamlRole200ResponseNetworksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationSamlRole200ResponseNetworksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationSamlRole200ResponseNetworksInner(val *UpdateOrganizationSamlRole200ResponseNetworksInner) *NullableUpdateOrganizationSamlRole200ResponseNetworksInner {
	return &NullableUpdateOrganizationSamlRole200ResponseNetworksInner{value: val, isSet: true}
}

func (v NullableUpdateOrganizationSamlRole200ResponseNetworksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationSamlRole200ResponseNetworksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


