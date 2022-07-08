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

// UpdateOrganizationEarlyAccessFeaturesOptInRequest struct for UpdateOrganizationEarlyAccessFeaturesOptInRequest
type UpdateOrganizationEarlyAccessFeaturesOptInRequest struct {
	// A list of network IDs to apply the opt-in to
	LimitScopeToNetworks []string `json:"limitScopeToNetworks,omitempty"`
}

// NewUpdateOrganizationEarlyAccessFeaturesOptInRequest instantiates a new UpdateOrganizationEarlyAccessFeaturesOptInRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationEarlyAccessFeaturesOptInRequest() *UpdateOrganizationEarlyAccessFeaturesOptInRequest {
	this := UpdateOrganizationEarlyAccessFeaturesOptInRequest{}
	return &this
}

// NewUpdateOrganizationEarlyAccessFeaturesOptInRequestWithDefaults instantiates a new UpdateOrganizationEarlyAccessFeaturesOptInRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationEarlyAccessFeaturesOptInRequestWithDefaults() *UpdateOrganizationEarlyAccessFeaturesOptInRequest {
	this := UpdateOrganizationEarlyAccessFeaturesOptInRequest{}
	return &this
}

// GetLimitScopeToNetworks returns the LimitScopeToNetworks field value if set, zero value otherwise.
func (o *UpdateOrganizationEarlyAccessFeaturesOptInRequest) GetLimitScopeToNetworks() []string {
	if o == nil || o.LimitScopeToNetworks == nil {
		var ret []string
		return ret
	}
	return o.LimitScopeToNetworks
}

// GetLimitScopeToNetworksOk returns a tuple with the LimitScopeToNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationEarlyAccessFeaturesOptInRequest) GetLimitScopeToNetworksOk() ([]string, bool) {
	if o == nil || o.LimitScopeToNetworks == nil {
		return nil, false
	}
	return o.LimitScopeToNetworks, true
}

// HasLimitScopeToNetworks returns a boolean if a field has been set.
func (o *UpdateOrganizationEarlyAccessFeaturesOptInRequest) HasLimitScopeToNetworks() bool {
	if o != nil && o.LimitScopeToNetworks != nil {
		return true
	}

	return false
}

// SetLimitScopeToNetworks gets a reference to the given []string and assigns it to the LimitScopeToNetworks field.
func (o *UpdateOrganizationEarlyAccessFeaturesOptInRequest) SetLimitScopeToNetworks(v []string) {
	o.LimitScopeToNetworks = v
}

func (o UpdateOrganizationEarlyAccessFeaturesOptInRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LimitScopeToNetworks != nil {
		toSerialize["limitScopeToNetworks"] = o.LimitScopeToNetworks
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateOrganizationEarlyAccessFeaturesOptInRequest struct {
	value *UpdateOrganizationEarlyAccessFeaturesOptInRequest
	isSet bool
}

func (v NullableUpdateOrganizationEarlyAccessFeaturesOptInRequest) Get() *UpdateOrganizationEarlyAccessFeaturesOptInRequest {
	return v.value
}

func (v *NullableUpdateOrganizationEarlyAccessFeaturesOptInRequest) Set(val *UpdateOrganizationEarlyAccessFeaturesOptInRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationEarlyAccessFeaturesOptInRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationEarlyAccessFeaturesOptInRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationEarlyAccessFeaturesOptInRequest(val *UpdateOrganizationEarlyAccessFeaturesOptInRequest) *NullableUpdateOrganizationEarlyAccessFeaturesOptInRequest {
	return &NullableUpdateOrganizationEarlyAccessFeaturesOptInRequest{value: val, isSet: true}
}

func (v NullableUpdateOrganizationEarlyAccessFeaturesOptInRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationEarlyAccessFeaturesOptInRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


