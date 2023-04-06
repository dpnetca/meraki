/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 April, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package meraki

import (
	"encoding/json"
	"time"
)

// checks if the GetNetworkWirelessUsageHistory200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkWirelessUsageHistory200ResponseInner{}

// GetNetworkWirelessUsageHistory200ResponseInner struct for GetNetworkWirelessUsageHistory200ResponseInner
type GetNetworkWirelessUsageHistory200ResponseInner struct {
	// The start time of the query range
	StartTs *time.Time `json:"startTs,omitempty"`
	// The end time of the query range
	EndTs *time.Time `json:"endTs,omitempty"`
	// Total usage in kilobytes-per-second
	TotalKbps *int32 `json:"totalKbps,omitempty"`
	// Sent kilobytes-per-second
	SentKbps *int32 `json:"sentKbps,omitempty"`
	// Received kilobytes-per-second
	ReceivedKbps *int32 `json:"receivedKbps,omitempty"`
}

// NewGetNetworkWirelessUsageHistory200ResponseInner instantiates a new GetNetworkWirelessUsageHistory200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWirelessUsageHistory200ResponseInner() *GetNetworkWirelessUsageHistory200ResponseInner {
	this := GetNetworkWirelessUsageHistory200ResponseInner{}
	return &this
}

// NewGetNetworkWirelessUsageHistory200ResponseInnerWithDefaults instantiates a new GetNetworkWirelessUsageHistory200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWirelessUsageHistory200ResponseInnerWithDefaults() *GetNetworkWirelessUsageHistory200ResponseInner {
	this := GetNetworkWirelessUsageHistory200ResponseInner{}
	return &this
}

// GetStartTs returns the StartTs field value if set, zero value otherwise.
func (o *GetNetworkWirelessUsageHistory200ResponseInner) GetStartTs() time.Time {
	if o == nil || IsNil(o.StartTs) {
		var ret time.Time
		return ret
	}
	return *o.StartTs
}

// GetStartTsOk returns a tuple with the StartTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessUsageHistory200ResponseInner) GetStartTsOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTs) {
		return nil, false
	}
	return o.StartTs, true
}

// HasStartTs returns a boolean if a field has been set.
func (o *GetNetworkWirelessUsageHistory200ResponseInner) HasStartTs() bool {
	if o != nil && !IsNil(o.StartTs) {
		return true
	}

	return false
}

// SetStartTs gets a reference to the given time.Time and assigns it to the StartTs field.
func (o *GetNetworkWirelessUsageHistory200ResponseInner) SetStartTs(v time.Time) {
	o.StartTs = &v
}

// GetEndTs returns the EndTs field value if set, zero value otherwise.
func (o *GetNetworkWirelessUsageHistory200ResponseInner) GetEndTs() time.Time {
	if o == nil || IsNil(o.EndTs) {
		var ret time.Time
		return ret
	}
	return *o.EndTs
}

// GetEndTsOk returns a tuple with the EndTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessUsageHistory200ResponseInner) GetEndTsOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTs) {
		return nil, false
	}
	return o.EndTs, true
}

// HasEndTs returns a boolean if a field has been set.
func (o *GetNetworkWirelessUsageHistory200ResponseInner) HasEndTs() bool {
	if o != nil && !IsNil(o.EndTs) {
		return true
	}

	return false
}

// SetEndTs gets a reference to the given time.Time and assigns it to the EndTs field.
func (o *GetNetworkWirelessUsageHistory200ResponseInner) SetEndTs(v time.Time) {
	o.EndTs = &v
}

// GetTotalKbps returns the TotalKbps field value if set, zero value otherwise.
func (o *GetNetworkWirelessUsageHistory200ResponseInner) GetTotalKbps() int32 {
	if o == nil || IsNil(o.TotalKbps) {
		var ret int32
		return ret
	}
	return *o.TotalKbps
}

// GetTotalKbpsOk returns a tuple with the TotalKbps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessUsageHistory200ResponseInner) GetTotalKbpsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalKbps) {
		return nil, false
	}
	return o.TotalKbps, true
}

// HasTotalKbps returns a boolean if a field has been set.
func (o *GetNetworkWirelessUsageHistory200ResponseInner) HasTotalKbps() bool {
	if o != nil && !IsNil(o.TotalKbps) {
		return true
	}

	return false
}

// SetTotalKbps gets a reference to the given int32 and assigns it to the TotalKbps field.
func (o *GetNetworkWirelessUsageHistory200ResponseInner) SetTotalKbps(v int32) {
	o.TotalKbps = &v
}

// GetSentKbps returns the SentKbps field value if set, zero value otherwise.
func (o *GetNetworkWirelessUsageHistory200ResponseInner) GetSentKbps() int32 {
	if o == nil || IsNil(o.SentKbps) {
		var ret int32
		return ret
	}
	return *o.SentKbps
}

// GetSentKbpsOk returns a tuple with the SentKbps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessUsageHistory200ResponseInner) GetSentKbpsOk() (*int32, bool) {
	if o == nil || IsNil(o.SentKbps) {
		return nil, false
	}
	return o.SentKbps, true
}

// HasSentKbps returns a boolean if a field has been set.
func (o *GetNetworkWirelessUsageHistory200ResponseInner) HasSentKbps() bool {
	if o != nil && !IsNil(o.SentKbps) {
		return true
	}

	return false
}

// SetSentKbps gets a reference to the given int32 and assigns it to the SentKbps field.
func (o *GetNetworkWirelessUsageHistory200ResponseInner) SetSentKbps(v int32) {
	o.SentKbps = &v
}

// GetReceivedKbps returns the ReceivedKbps field value if set, zero value otherwise.
func (o *GetNetworkWirelessUsageHistory200ResponseInner) GetReceivedKbps() int32 {
	if o == nil || IsNil(o.ReceivedKbps) {
		var ret int32
		return ret
	}
	return *o.ReceivedKbps
}

// GetReceivedKbpsOk returns a tuple with the ReceivedKbps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessUsageHistory200ResponseInner) GetReceivedKbpsOk() (*int32, bool) {
	if o == nil || IsNil(o.ReceivedKbps) {
		return nil, false
	}
	return o.ReceivedKbps, true
}

// HasReceivedKbps returns a boolean if a field has been set.
func (o *GetNetworkWirelessUsageHistory200ResponseInner) HasReceivedKbps() bool {
	if o != nil && !IsNil(o.ReceivedKbps) {
		return true
	}

	return false
}

// SetReceivedKbps gets a reference to the given int32 and assigns it to the ReceivedKbps field.
func (o *GetNetworkWirelessUsageHistory200ResponseInner) SetReceivedKbps(v int32) {
	o.ReceivedKbps = &v
}

func (o GetNetworkWirelessUsageHistory200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkWirelessUsageHistory200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartTs) {
		toSerialize["startTs"] = o.StartTs
	}
	if !IsNil(o.EndTs) {
		toSerialize["endTs"] = o.EndTs
	}
	if !IsNil(o.TotalKbps) {
		toSerialize["totalKbps"] = o.TotalKbps
	}
	if !IsNil(o.SentKbps) {
		toSerialize["sentKbps"] = o.SentKbps
	}
	if !IsNil(o.ReceivedKbps) {
		toSerialize["receivedKbps"] = o.ReceivedKbps
	}
	return toSerialize, nil
}

type NullableGetNetworkWirelessUsageHistory200ResponseInner struct {
	value *GetNetworkWirelessUsageHistory200ResponseInner
	isSet bool
}

func (v NullableGetNetworkWirelessUsageHistory200ResponseInner) Get() *GetNetworkWirelessUsageHistory200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkWirelessUsageHistory200ResponseInner) Set(val *GetNetworkWirelessUsageHistory200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWirelessUsageHistory200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWirelessUsageHistory200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWirelessUsageHistory200ResponseInner(val *GetNetworkWirelessUsageHistory200ResponseInner) *NullableGetNetworkWirelessUsageHistory200ResponseInner {
	return &NullableGetNetworkWirelessUsageHistory200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkWirelessUsageHistory200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWirelessUsageHistory200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


