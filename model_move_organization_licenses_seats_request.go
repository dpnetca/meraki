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

// MoveOrganizationLicensesSeatsRequest struct for MoveOrganizationLicensesSeatsRequest
type MoveOrganizationLicensesSeatsRequest struct {
	// The ID of the organization to move the SM seats to
	DestOrganizationId string `json:"destOrganizationId"`
	// The ID of the SM license to move the seats from
	LicenseId string `json:"licenseId"`
	// The number of seats to move to the new organization. Must be less than or equal to the total number of seats of the license
	SeatCount int32 `json:"seatCount"`
}

// NewMoveOrganizationLicensesSeatsRequest instantiates a new MoveOrganizationLicensesSeatsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoveOrganizationLicensesSeatsRequest(destOrganizationId string, licenseId string, seatCount int32) *MoveOrganizationLicensesSeatsRequest {
	this := MoveOrganizationLicensesSeatsRequest{}
	this.DestOrganizationId = destOrganizationId
	this.LicenseId = licenseId
	this.SeatCount = seatCount
	return &this
}

// NewMoveOrganizationLicensesSeatsRequestWithDefaults instantiates a new MoveOrganizationLicensesSeatsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoveOrganizationLicensesSeatsRequestWithDefaults() *MoveOrganizationLicensesSeatsRequest {
	this := MoveOrganizationLicensesSeatsRequest{}
	return &this
}

// GetDestOrganizationId returns the DestOrganizationId field value
func (o *MoveOrganizationLicensesSeatsRequest) GetDestOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestOrganizationId
}

// GetDestOrganizationIdOk returns a tuple with the DestOrganizationId field value
// and a boolean to check if the value has been set.
func (o *MoveOrganizationLicensesSeatsRequest) GetDestOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestOrganizationId, true
}

// SetDestOrganizationId sets field value
func (o *MoveOrganizationLicensesSeatsRequest) SetDestOrganizationId(v string) {
	o.DestOrganizationId = v
}

// GetLicenseId returns the LicenseId field value
func (o *MoveOrganizationLicensesSeatsRequest) GetLicenseId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LicenseId
}

// GetLicenseIdOk returns a tuple with the LicenseId field value
// and a boolean to check if the value has been set.
func (o *MoveOrganizationLicensesSeatsRequest) GetLicenseIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LicenseId, true
}

// SetLicenseId sets field value
func (o *MoveOrganizationLicensesSeatsRequest) SetLicenseId(v string) {
	o.LicenseId = v
}

// GetSeatCount returns the SeatCount field value
func (o *MoveOrganizationLicensesSeatsRequest) GetSeatCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SeatCount
}

// GetSeatCountOk returns a tuple with the SeatCount field value
// and a boolean to check if the value has been set.
func (o *MoveOrganizationLicensesSeatsRequest) GetSeatCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeatCount, true
}

// SetSeatCount sets field value
func (o *MoveOrganizationLicensesSeatsRequest) SetSeatCount(v int32) {
	o.SeatCount = v
}

func (o MoveOrganizationLicensesSeatsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["destOrganizationId"] = o.DestOrganizationId
	}
	if true {
		toSerialize["licenseId"] = o.LicenseId
	}
	if true {
		toSerialize["seatCount"] = o.SeatCount
	}
	return json.Marshal(toSerialize)
}

type NullableMoveOrganizationLicensesSeatsRequest struct {
	value *MoveOrganizationLicensesSeatsRequest
	isSet bool
}

func (v NullableMoveOrganizationLicensesSeatsRequest) Get() *MoveOrganizationLicensesSeatsRequest {
	return v.value
}

func (v *NullableMoveOrganizationLicensesSeatsRequest) Set(val *MoveOrganizationLicensesSeatsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMoveOrganizationLicensesSeatsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMoveOrganizationLicensesSeatsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoveOrganizationLicensesSeatsRequest(val *MoveOrganizationLicensesSeatsRequest) *NullableMoveOrganizationLicensesSeatsRequest {
	return &NullableMoveOrganizationLicensesSeatsRequest{value: val, isSet: true}
}

func (v NullableMoveOrganizationLicensesSeatsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoveOrganizationLicensesSeatsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


