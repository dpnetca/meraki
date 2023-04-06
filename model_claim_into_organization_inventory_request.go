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

// checks if the ClaimIntoOrganizationInventoryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClaimIntoOrganizationInventoryRequest{}

// ClaimIntoOrganizationInventoryRequest struct for ClaimIntoOrganizationInventoryRequest
type ClaimIntoOrganizationInventoryRequest struct {
	// The numbers of the orders that should be claimed
	Orders []string `json:"orders,omitempty"`
	// The serials of the devices that should be claimed
	Serials []string `json:"serials,omitempty"`
	// The licenses that should be claimed
	Licenses []ClaimIntoOrganizationInventoryRequestLicensesInner `json:"licenses,omitempty"`
}

// NewClaimIntoOrganizationInventoryRequest instantiates a new ClaimIntoOrganizationInventoryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClaimIntoOrganizationInventoryRequest() *ClaimIntoOrganizationInventoryRequest {
	this := ClaimIntoOrganizationInventoryRequest{}
	return &this
}

// NewClaimIntoOrganizationInventoryRequestWithDefaults instantiates a new ClaimIntoOrganizationInventoryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClaimIntoOrganizationInventoryRequestWithDefaults() *ClaimIntoOrganizationInventoryRequest {
	this := ClaimIntoOrganizationInventoryRequest{}
	return &this
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *ClaimIntoOrganizationInventoryRequest) GetOrders() []string {
	if o == nil || IsNil(o.Orders) {
		var ret []string
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimIntoOrganizationInventoryRequest) GetOrdersOk() ([]string, bool) {
	if o == nil || IsNil(o.Orders) {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *ClaimIntoOrganizationInventoryRequest) HasOrders() bool {
	if o != nil && !IsNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []string and assigns it to the Orders field.
func (o *ClaimIntoOrganizationInventoryRequest) SetOrders(v []string) {
	o.Orders = v
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *ClaimIntoOrganizationInventoryRequest) GetSerials() []string {
	if o == nil || IsNil(o.Serials) {
		var ret []string
		return ret
	}
	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimIntoOrganizationInventoryRequest) GetSerialsOk() ([]string, bool) {
	if o == nil || IsNil(o.Serials) {
		return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *ClaimIntoOrganizationInventoryRequest) HasSerials() bool {
	if o != nil && !IsNil(o.Serials) {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *ClaimIntoOrganizationInventoryRequest) SetSerials(v []string) {
	o.Serials = v
}

// GetLicenses returns the Licenses field value if set, zero value otherwise.
func (o *ClaimIntoOrganizationInventoryRequest) GetLicenses() []ClaimIntoOrganizationInventoryRequestLicensesInner {
	if o == nil || IsNil(o.Licenses) {
		var ret []ClaimIntoOrganizationInventoryRequestLicensesInner
		return ret
	}
	return o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimIntoOrganizationInventoryRequest) GetLicensesOk() ([]ClaimIntoOrganizationInventoryRequestLicensesInner, bool) {
	if o == nil || IsNil(o.Licenses) {
		return nil, false
	}
	return o.Licenses, true
}

// HasLicenses returns a boolean if a field has been set.
func (o *ClaimIntoOrganizationInventoryRequest) HasLicenses() bool {
	if o != nil && !IsNil(o.Licenses) {
		return true
	}

	return false
}

// SetLicenses gets a reference to the given []ClaimIntoOrganizationInventoryRequestLicensesInner and assigns it to the Licenses field.
func (o *ClaimIntoOrganizationInventoryRequest) SetLicenses(v []ClaimIntoOrganizationInventoryRequestLicensesInner) {
	o.Licenses = v
}

func (o ClaimIntoOrganizationInventoryRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClaimIntoOrganizationInventoryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Orders) {
		toSerialize["orders"] = o.Orders
	}
	if !IsNil(o.Serials) {
		toSerialize["serials"] = o.Serials
	}
	if !IsNil(o.Licenses) {
		toSerialize["licenses"] = o.Licenses
	}
	return toSerialize, nil
}

type NullableClaimIntoOrganizationInventoryRequest struct {
	value *ClaimIntoOrganizationInventoryRequest
	isSet bool
}

func (v NullableClaimIntoOrganizationInventoryRequest) Get() *ClaimIntoOrganizationInventoryRequest {
	return v.value
}

func (v *NullableClaimIntoOrganizationInventoryRequest) Set(val *ClaimIntoOrganizationInventoryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableClaimIntoOrganizationInventoryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableClaimIntoOrganizationInventoryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClaimIntoOrganizationInventoryRequest(val *ClaimIntoOrganizationInventoryRequest) *NullableClaimIntoOrganizationInventoryRequest {
	return &NullableClaimIntoOrganizationInventoryRequest{value: val, isSet: true}
}

func (v NullableClaimIntoOrganizationInventoryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClaimIntoOrganizationInventoryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


