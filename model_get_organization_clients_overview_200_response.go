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

// GetOrganizationClientsOverview200Response struct for GetOrganizationClientsOverview200Response
type GetOrganizationClientsOverview200Response struct {
	Usage *GetOrganizationClientsOverview200ResponseUsage `json:"usage,omitempty"`
	Counts *GetOrganizationClientsOverview200ResponseCounts `json:"counts,omitempty"`
}

// NewGetOrganizationClientsOverview200Response instantiates a new GetOrganizationClientsOverview200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationClientsOverview200Response() *GetOrganizationClientsOverview200Response {
	this := GetOrganizationClientsOverview200Response{}
	return &this
}

// NewGetOrganizationClientsOverview200ResponseWithDefaults instantiates a new GetOrganizationClientsOverview200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationClientsOverview200ResponseWithDefaults() *GetOrganizationClientsOverview200Response {
	this := GetOrganizationClientsOverview200Response{}
	return &this
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *GetOrganizationClientsOverview200Response) GetUsage() GetOrganizationClientsOverview200ResponseUsage {
	if o == nil || o.Usage == nil {
		var ret GetOrganizationClientsOverview200ResponseUsage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationClientsOverview200Response) GetUsageOk() (*GetOrganizationClientsOverview200ResponseUsage, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *GetOrganizationClientsOverview200Response) HasUsage() bool {
	if o != nil && o.Usage != nil {
		return true
	}

	return false
}

// SetUsage gets a reference to the given GetOrganizationClientsOverview200ResponseUsage and assigns it to the Usage field.
func (o *GetOrganizationClientsOverview200Response) SetUsage(v GetOrganizationClientsOverview200ResponseUsage) {
	o.Usage = &v
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *GetOrganizationClientsOverview200Response) GetCounts() GetOrganizationClientsOverview200ResponseCounts {
	if o == nil || o.Counts == nil {
		var ret GetOrganizationClientsOverview200ResponseCounts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationClientsOverview200Response) GetCountsOk() (*GetOrganizationClientsOverview200ResponseCounts, bool) {
	if o == nil || o.Counts == nil {
		return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *GetOrganizationClientsOverview200Response) HasCounts() bool {
	if o != nil && o.Counts != nil {
		return true
	}

	return false
}

// SetCounts gets a reference to the given GetOrganizationClientsOverview200ResponseCounts and assigns it to the Counts field.
func (o *GetOrganizationClientsOverview200Response) SetCounts(v GetOrganizationClientsOverview200ResponseCounts) {
	o.Counts = &v
}

func (o GetOrganizationClientsOverview200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Usage != nil {
		toSerialize["usage"] = o.Usage
	}
	if o.Counts != nil {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrganizationClientsOverview200Response struct {
	value *GetOrganizationClientsOverview200Response
	isSet bool
}

func (v NullableGetOrganizationClientsOverview200Response) Get() *GetOrganizationClientsOverview200Response {
	return v.value
}

func (v *NullableGetOrganizationClientsOverview200Response) Set(val *GetOrganizationClientsOverview200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationClientsOverview200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationClientsOverview200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationClientsOverview200Response(val *GetOrganizationClientsOverview200Response) *NullableGetOrganizationClientsOverview200Response {
	return &NullableGetOrganizationClientsOverview200Response{value: val, isSet: true}
}

func (v NullableGetOrganizationClientsOverview200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationClientsOverview200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

