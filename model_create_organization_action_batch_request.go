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

// checks if the CreateOrganizationActionBatchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationActionBatchRequest{}

// CreateOrganizationActionBatchRequest struct for CreateOrganizationActionBatchRequest
type CreateOrganizationActionBatchRequest struct {
	// Set to true for immediate execution. Set to false if the action should be previewed before executing. This property cannot be unset once it is true. Defaults to false.
	Confirmed *bool `json:"confirmed,omitempty"`
	// Set to true to force the batch to run synchronous. There can be at most 20 actions in synchronous batch. Defaults to false.
	Synchronous *bool `json:"synchronous,omitempty"`
	// A set of changes to make as part of this action (<a href='https://developer.cisco.com/meraki/api/#/rest/guides/action-batches/'>more details</a>)
	Actions []CreateOrganizationActionBatchRequestActionsInner `json:"actions"`
}

// NewCreateOrganizationActionBatchRequest instantiates a new CreateOrganizationActionBatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationActionBatchRequest(actions []CreateOrganizationActionBatchRequestActionsInner) *CreateOrganizationActionBatchRequest {
	this := CreateOrganizationActionBatchRequest{}
	this.Actions = actions
	return &this
}

// NewCreateOrganizationActionBatchRequestWithDefaults instantiates a new CreateOrganizationActionBatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationActionBatchRequestWithDefaults() *CreateOrganizationActionBatchRequest {
	this := CreateOrganizationActionBatchRequest{}
	return &this
}

// GetConfirmed returns the Confirmed field value if set, zero value otherwise.
func (o *CreateOrganizationActionBatchRequest) GetConfirmed() bool {
	if o == nil || IsNil(o.Confirmed) {
		var ret bool
		return ret
	}
	return *o.Confirmed
}

// GetConfirmedOk returns a tuple with the Confirmed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationActionBatchRequest) GetConfirmedOk() (*bool, bool) {
	if o == nil || IsNil(o.Confirmed) {
		return nil, false
	}
	return o.Confirmed, true
}

// HasConfirmed returns a boolean if a field has been set.
func (o *CreateOrganizationActionBatchRequest) HasConfirmed() bool {
	if o != nil && !IsNil(o.Confirmed) {
		return true
	}

	return false
}

// SetConfirmed gets a reference to the given bool and assigns it to the Confirmed field.
func (o *CreateOrganizationActionBatchRequest) SetConfirmed(v bool) {
	o.Confirmed = &v
}

// GetSynchronous returns the Synchronous field value if set, zero value otherwise.
func (o *CreateOrganizationActionBatchRequest) GetSynchronous() bool {
	if o == nil || IsNil(o.Synchronous) {
		var ret bool
		return ret
	}
	return *o.Synchronous
}

// GetSynchronousOk returns a tuple with the Synchronous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationActionBatchRequest) GetSynchronousOk() (*bool, bool) {
	if o == nil || IsNil(o.Synchronous) {
		return nil, false
	}
	return o.Synchronous, true
}

// HasSynchronous returns a boolean if a field has been set.
func (o *CreateOrganizationActionBatchRequest) HasSynchronous() bool {
	if o != nil && !IsNil(o.Synchronous) {
		return true
	}

	return false
}

// SetSynchronous gets a reference to the given bool and assigns it to the Synchronous field.
func (o *CreateOrganizationActionBatchRequest) SetSynchronous(v bool) {
	o.Synchronous = &v
}

// GetActions returns the Actions field value
func (o *CreateOrganizationActionBatchRequest) GetActions() []CreateOrganizationActionBatchRequestActionsInner {
	if o == nil {
		var ret []CreateOrganizationActionBatchRequestActionsInner
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationActionBatchRequest) GetActionsOk() ([]CreateOrganizationActionBatchRequestActionsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Actions, true
}

// SetActions sets field value
func (o *CreateOrganizationActionBatchRequest) SetActions(v []CreateOrganizationActionBatchRequestActionsInner) {
	o.Actions = v
}

func (o CreateOrganizationActionBatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationActionBatchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Confirmed) {
		toSerialize["confirmed"] = o.Confirmed
	}
	if !IsNil(o.Synchronous) {
		toSerialize["synchronous"] = o.Synchronous
	}
	toSerialize["actions"] = o.Actions
	return toSerialize, nil
}

type NullableCreateOrganizationActionBatchRequest struct {
	value *CreateOrganizationActionBatchRequest
	isSet bool
}

func (v NullableCreateOrganizationActionBatchRequest) Get() *CreateOrganizationActionBatchRequest {
	return v.value
}

func (v *NullableCreateOrganizationActionBatchRequest) Set(val *CreateOrganizationActionBatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationActionBatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationActionBatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationActionBatchRequest(val *CreateOrganizationActionBatchRequest) *NullableCreateOrganizationActionBatchRequest {
	return &NullableCreateOrganizationActionBatchRequest{value: val, isSet: true}
}

func (v NullableCreateOrganizationActionBatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationActionBatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


