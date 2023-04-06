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

// checks if the GetNetworkPoliciesByClient200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkPoliciesByClient200ResponseInner{}

// GetNetworkPoliciesByClient200ResponseInner struct for GetNetworkPoliciesByClient200ResponseInner
type GetNetworkPoliciesByClient200ResponseInner struct {
	// Name of client
	Name *string `json:"name,omitempty"`
	// ID of client
	ClientId *string `json:"clientId,omitempty"`
	// Assigned policies
	Assigned []GetNetworkPoliciesByClient200ResponseInnerAssignedInner `json:"assigned,omitempty"`
}

// NewGetNetworkPoliciesByClient200ResponseInner instantiates a new GetNetworkPoliciesByClient200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkPoliciesByClient200ResponseInner() *GetNetworkPoliciesByClient200ResponseInner {
	this := GetNetworkPoliciesByClient200ResponseInner{}
	return &this
}

// NewGetNetworkPoliciesByClient200ResponseInnerWithDefaults instantiates a new GetNetworkPoliciesByClient200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkPoliciesByClient200ResponseInnerWithDefaults() *GetNetworkPoliciesByClient200ResponseInner {
	this := GetNetworkPoliciesByClient200ResponseInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkPoliciesByClient200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkPoliciesByClient200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkPoliciesByClient200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkPoliciesByClient200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *GetNetworkPoliciesByClient200ResponseInner) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkPoliciesByClient200ResponseInner) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *GetNetworkPoliciesByClient200ResponseInner) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *GetNetworkPoliciesByClient200ResponseInner) SetClientId(v string) {
	o.ClientId = &v
}

// GetAssigned returns the Assigned field value if set, zero value otherwise.
func (o *GetNetworkPoliciesByClient200ResponseInner) GetAssigned() []GetNetworkPoliciesByClient200ResponseInnerAssignedInner {
	if o == nil || IsNil(o.Assigned) {
		var ret []GetNetworkPoliciesByClient200ResponseInnerAssignedInner
		return ret
	}
	return o.Assigned
}

// GetAssignedOk returns a tuple with the Assigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkPoliciesByClient200ResponseInner) GetAssignedOk() ([]GetNetworkPoliciesByClient200ResponseInnerAssignedInner, bool) {
	if o == nil || IsNil(o.Assigned) {
		return nil, false
	}
	return o.Assigned, true
}

// HasAssigned returns a boolean if a field has been set.
func (o *GetNetworkPoliciesByClient200ResponseInner) HasAssigned() bool {
	if o != nil && !IsNil(o.Assigned) {
		return true
	}

	return false
}

// SetAssigned gets a reference to the given []GetNetworkPoliciesByClient200ResponseInnerAssignedInner and assigns it to the Assigned field.
func (o *GetNetworkPoliciesByClient200ResponseInner) SetAssigned(v []GetNetworkPoliciesByClient200ResponseInnerAssignedInner) {
	o.Assigned = v
}

func (o GetNetworkPoliciesByClient200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkPoliciesByClient200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !IsNil(o.Assigned) {
		toSerialize["assigned"] = o.Assigned
	}
	return toSerialize, nil
}

type NullableGetNetworkPoliciesByClient200ResponseInner struct {
	value *GetNetworkPoliciesByClient200ResponseInner
	isSet bool
}

func (v NullableGetNetworkPoliciesByClient200ResponseInner) Get() *GetNetworkPoliciesByClient200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkPoliciesByClient200ResponseInner) Set(val *GetNetworkPoliciesByClient200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkPoliciesByClient200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkPoliciesByClient200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkPoliciesByClient200ResponseInner(val *GetNetworkPoliciesByClient200ResponseInner) *NullableGetNetworkPoliciesByClient200ResponseInner {
	return &NullableGetNetworkPoliciesByClient200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkPoliciesByClient200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkPoliciesByClient200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


