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

// CreateNetworkFloorPlanRequestBottomLeftCorner The longitude and latitude of the bottom left corner of your floor plan.
type CreateNetworkFloorPlanRequestBottomLeftCorner struct {
	// Latitude
	Lat *float32 `json:"lat,omitempty"`
	// Longitude
	Lng *float32 `json:"lng,omitempty"`
}

// NewCreateNetworkFloorPlanRequestBottomLeftCorner instantiates a new CreateNetworkFloorPlanRequestBottomLeftCorner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkFloorPlanRequestBottomLeftCorner() *CreateNetworkFloorPlanRequestBottomLeftCorner {
	this := CreateNetworkFloorPlanRequestBottomLeftCorner{}
	return &this
}

// NewCreateNetworkFloorPlanRequestBottomLeftCornerWithDefaults instantiates a new CreateNetworkFloorPlanRequestBottomLeftCorner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkFloorPlanRequestBottomLeftCornerWithDefaults() *CreateNetworkFloorPlanRequestBottomLeftCorner {
	this := CreateNetworkFloorPlanRequestBottomLeftCorner{}
	return &this
}

// GetLat returns the Lat field value if set, zero value otherwise.
func (o *CreateNetworkFloorPlanRequestBottomLeftCorner) GetLat() float32 {
	if o == nil || o.Lat == nil {
		var ret float32
		return ret
	}
	return *o.Lat
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkFloorPlanRequestBottomLeftCorner) GetLatOk() (*float32, bool) {
	if o == nil || o.Lat == nil {
		return nil, false
	}
	return o.Lat, true
}

// HasLat returns a boolean if a field has been set.
func (o *CreateNetworkFloorPlanRequestBottomLeftCorner) HasLat() bool {
	if o != nil && o.Lat != nil {
		return true
	}

	return false
}

// SetLat gets a reference to the given float32 and assigns it to the Lat field.
func (o *CreateNetworkFloorPlanRequestBottomLeftCorner) SetLat(v float32) {
	o.Lat = &v
}

// GetLng returns the Lng field value if set, zero value otherwise.
func (o *CreateNetworkFloorPlanRequestBottomLeftCorner) GetLng() float32 {
	if o == nil || o.Lng == nil {
		var ret float32
		return ret
	}
	return *o.Lng
}

// GetLngOk returns a tuple with the Lng field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkFloorPlanRequestBottomLeftCorner) GetLngOk() (*float32, bool) {
	if o == nil || o.Lng == nil {
		return nil, false
	}
	return o.Lng, true
}

// HasLng returns a boolean if a field has been set.
func (o *CreateNetworkFloorPlanRequestBottomLeftCorner) HasLng() bool {
	if o != nil && o.Lng != nil {
		return true
	}

	return false
}

// SetLng gets a reference to the given float32 and assigns it to the Lng field.
func (o *CreateNetworkFloorPlanRequestBottomLeftCorner) SetLng(v float32) {
	o.Lng = &v
}

func (o CreateNetworkFloorPlanRequestBottomLeftCorner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Lat != nil {
		toSerialize["lat"] = o.Lat
	}
	if o.Lng != nil {
		toSerialize["lng"] = o.Lng
	}
	return json.Marshal(toSerialize)
}

type NullableCreateNetworkFloorPlanRequestBottomLeftCorner struct {
	value *CreateNetworkFloorPlanRequestBottomLeftCorner
	isSet bool
}

func (v NullableCreateNetworkFloorPlanRequestBottomLeftCorner) Get() *CreateNetworkFloorPlanRequestBottomLeftCorner {
	return v.value
}

func (v *NullableCreateNetworkFloorPlanRequestBottomLeftCorner) Set(val *CreateNetworkFloorPlanRequestBottomLeftCorner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkFloorPlanRequestBottomLeftCorner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkFloorPlanRequestBottomLeftCorner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkFloorPlanRequestBottomLeftCorner(val *CreateNetworkFloorPlanRequestBottomLeftCorner) *NullableCreateNetworkFloorPlanRequestBottomLeftCorner {
	return &NullableCreateNetworkFloorPlanRequestBottomLeftCorner{value: val, isSet: true}
}

func (v NullableCreateNetworkFloorPlanRequestBottomLeftCorner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkFloorPlanRequestBottomLeftCorner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


