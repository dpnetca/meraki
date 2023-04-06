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

// checks if the GetOrganizations200ResponseInnerCloud type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizations200ResponseInnerCloud{}

// GetOrganizations200ResponseInnerCloud Data for this organization
type GetOrganizations200ResponseInnerCloud struct {
	Region *GetOrganizations200ResponseInnerCloudRegion `json:"region,omitempty"`
}

// NewGetOrganizations200ResponseInnerCloud instantiates a new GetOrganizations200ResponseInnerCloud object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizations200ResponseInnerCloud() *GetOrganizations200ResponseInnerCloud {
	this := GetOrganizations200ResponseInnerCloud{}
	return &this
}

// NewGetOrganizations200ResponseInnerCloudWithDefaults instantiates a new GetOrganizations200ResponseInnerCloud object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizations200ResponseInnerCloudWithDefaults() *GetOrganizations200ResponseInnerCloud {
	this := GetOrganizations200ResponseInnerCloud{}
	return &this
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *GetOrganizations200ResponseInnerCloud) GetRegion() GetOrganizations200ResponseInnerCloudRegion {
	if o == nil || IsNil(o.Region) {
		var ret GetOrganizations200ResponseInnerCloudRegion
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizations200ResponseInnerCloud) GetRegionOk() (*GetOrganizations200ResponseInnerCloudRegion, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *GetOrganizations200ResponseInnerCloud) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given GetOrganizations200ResponseInnerCloudRegion and assigns it to the Region field.
func (o *GetOrganizations200ResponseInnerCloud) SetRegion(v GetOrganizations200ResponseInnerCloudRegion) {
	o.Region = &v
}

func (o GetOrganizations200ResponseInnerCloud) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizations200ResponseInnerCloud) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	return toSerialize, nil
}

type NullableGetOrganizations200ResponseInnerCloud struct {
	value *GetOrganizations200ResponseInnerCloud
	isSet bool
}

func (v NullableGetOrganizations200ResponseInnerCloud) Get() *GetOrganizations200ResponseInnerCloud {
	return v.value
}

func (v *NullableGetOrganizations200ResponseInnerCloud) Set(val *GetOrganizations200ResponseInnerCloud) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizations200ResponseInnerCloud) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizations200ResponseInnerCloud) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizations200ResponseInnerCloud(val *GetOrganizations200ResponseInnerCloud) *NullableGetOrganizations200ResponseInnerCloud {
	return &NullableGetOrganizations200ResponseInnerCloud{value: val, isSet: true}
}

func (v NullableGetOrganizations200ResponseInnerCloud) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizations200ResponseInnerCloud) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


