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

// checks if the CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorization{}

// CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorization VTY AAA authorization
type CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorization struct {
	Group *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthenticationGroup `json:"group,omitempty"`
}

// NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorization instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorization() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorization {
	this := CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorization{}
	return &this
}

// NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorizationWithDefaults instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorizationWithDefaults() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorization {
	this := CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorization{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorization) GetGroup() CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthenticationGroup {
	if o == nil || IsNil(o.Group) {
		var ret CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthenticationGroup
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorization) GetGroupOk() (*CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthenticationGroup, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorization) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthenticationGroup and assigns it to the Group field.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorization) SetGroup(v CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthenticationGroup) {
	o.Group = &v
}

func (o CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	return toSerialize, nil
}

type NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorization struct {
	value *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorization
	isSet bool
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorization) Get() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorization {
	return v.value
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorization) Set(val *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorization) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorization) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorization(val *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorization) *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorization {
	return &NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorization{value: val, isSet: true}
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

