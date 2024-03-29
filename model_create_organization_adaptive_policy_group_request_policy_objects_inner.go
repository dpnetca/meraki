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

// checks if the CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner{}

// CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner struct for CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner
type CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner struct {
	// The ID of the policy object
	Id *string `json:"id,omitempty"`
	// The name of the policy object
	Name *string `json:"name,omitempty"`
}

// NewCreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner instantiates a new CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner() *CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner {
	this := CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner{}
	return &this
}

// NewCreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInnerWithDefaults instantiates a new CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInnerWithDefaults() *CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner {
	this := CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner) SetName(v string) {
	o.Name = &v
}

func (o CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableCreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner struct {
	value *CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner
	isSet bool
}

func (v NullableCreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner) Get() *CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner {
	return v.value
}

func (v *NullableCreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner) Set(val *CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner(val *CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner) *NullableCreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner {
	return &NullableCreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner{value: val, isSet: true}
}

func (v NullableCreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


