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

// checks if the UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1{}

// UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1 The bandwidth settings for the 'wan1' uplink
type UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1 struct {
	// The maximum upload limit (integer, in Kbps). null indicates no limit
	LimitUp *int32 `json:"limitUp,omitempty"`
	// The maximum download limit (integer, in Kbps). null indicates no limit
	LimitDown *int32 `json:"limitDown,omitempty"`
}

// NewUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1 instantiates a new UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1() *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1 {
	this := UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1{}
	return &this
}

// NewUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1WithDefaults instantiates a new UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1WithDefaults() *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1 {
	this := UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1{}
	return &this
}

// GetLimitUp returns the LimitUp field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1) GetLimitUp() int32 {
	if o == nil || IsNil(o.LimitUp) {
		var ret int32
		return ret
	}
	return *o.LimitUp
}

// GetLimitUpOk returns a tuple with the LimitUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1) GetLimitUpOk() (*int32, bool) {
	if o == nil || IsNil(o.LimitUp) {
		return nil, false
	}
	return o.LimitUp, true
}

// HasLimitUp returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1) HasLimitUp() bool {
	if o != nil && !IsNil(o.LimitUp) {
		return true
	}

	return false
}

// SetLimitUp gets a reference to the given int32 and assigns it to the LimitUp field.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1) SetLimitUp(v int32) {
	o.LimitUp = &v
}

// GetLimitDown returns the LimitDown field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1) GetLimitDown() int32 {
	if o == nil || IsNil(o.LimitDown) {
		var ret int32
		return ret
	}
	return *o.LimitDown
}

// GetLimitDownOk returns a tuple with the LimitDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1) GetLimitDownOk() (*int32, bool) {
	if o == nil || IsNil(o.LimitDown) {
		return nil, false
	}
	return o.LimitDown, true
}

// HasLimitDown returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1) HasLimitDown() bool {
	if o != nil && !IsNil(o.LimitDown) {
		return true
	}

	return false
}

// SetLimitDown gets a reference to the given int32 and assigns it to the LimitDown field.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1) SetLimitDown(v int32) {
	o.LimitDown = &v
}

func (o UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LimitUp) {
		toSerialize["limitUp"] = o.LimitUp
	}
	if !IsNil(o.LimitDown) {
		toSerialize["limitDown"] = o.LimitDown
	}
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1 struct {
	value *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1
	isSet bool
}

func (v NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1) Get() *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1 {
	return v.value
}

func (v *NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1) Set(val *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1(val *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1) *NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1 {
	return &NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


