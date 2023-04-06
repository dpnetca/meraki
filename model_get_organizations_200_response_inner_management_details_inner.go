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

// checks if the GetOrganizations200ResponseInnerManagementDetailsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizations200ResponseInnerManagementDetailsInner{}

// GetOrganizations200ResponseInnerManagementDetailsInner struct for GetOrganizations200ResponseInnerManagementDetailsInner
type GetOrganizations200ResponseInnerManagementDetailsInner struct {
	// Name of management data
	Name *string `json:"name,omitempty"`
	// Value of management data
	Value *string `json:"value,omitempty"`
}

// NewGetOrganizations200ResponseInnerManagementDetailsInner instantiates a new GetOrganizations200ResponseInnerManagementDetailsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizations200ResponseInnerManagementDetailsInner() *GetOrganizations200ResponseInnerManagementDetailsInner {
	this := GetOrganizations200ResponseInnerManagementDetailsInner{}
	return &this
}

// NewGetOrganizations200ResponseInnerManagementDetailsInnerWithDefaults instantiates a new GetOrganizations200ResponseInnerManagementDetailsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizations200ResponseInnerManagementDetailsInnerWithDefaults() *GetOrganizations200ResponseInnerManagementDetailsInner {
	this := GetOrganizations200ResponseInnerManagementDetailsInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizations200ResponseInnerManagementDetailsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizations200ResponseInnerManagementDetailsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizations200ResponseInnerManagementDetailsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizations200ResponseInnerManagementDetailsInner) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *GetOrganizations200ResponseInnerManagementDetailsInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizations200ResponseInnerManagementDetailsInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *GetOrganizations200ResponseInnerManagementDetailsInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *GetOrganizations200ResponseInnerManagementDetailsInner) SetValue(v string) {
	o.Value = &v
}

func (o GetOrganizations200ResponseInnerManagementDetailsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizations200ResponseInnerManagementDetailsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableGetOrganizations200ResponseInnerManagementDetailsInner struct {
	value *GetOrganizations200ResponseInnerManagementDetailsInner
	isSet bool
}

func (v NullableGetOrganizations200ResponseInnerManagementDetailsInner) Get() *GetOrganizations200ResponseInnerManagementDetailsInner {
	return v.value
}

func (v *NullableGetOrganizations200ResponseInnerManagementDetailsInner) Set(val *GetOrganizations200ResponseInnerManagementDetailsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizations200ResponseInnerManagementDetailsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizations200ResponseInnerManagementDetailsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizations200ResponseInnerManagementDetailsInner(val *GetOrganizations200ResponseInnerManagementDetailsInner) *NullableGetOrganizations200ResponseInnerManagementDetailsInner {
	return &NullableGetOrganizations200ResponseInnerManagementDetailsInner{value: val, isSet: true}
}

func (v NullableGetOrganizations200ResponseInnerManagementDetailsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizations200ResponseInnerManagementDetailsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

