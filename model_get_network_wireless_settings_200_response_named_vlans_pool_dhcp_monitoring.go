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

// checks if the GetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring{}

// GetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring Named VLAN Pool DHCP Monitoring settings.
type GetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring struct {
	// Whether or not devices using named VLAN pools should remove dirty VLANs from the pool, thereby preventing clients from being assigned to VLANs where they would be unable to obtain an IP address via DHCP
	Enabled *bool `json:"enabled,omitempty"`
	// The duration in minutes that devices will refrain from using dirty VLANs before adding them back to the pool.
	Duration *int32 `json:"duration,omitempty"`
}

// NewGetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring instantiates a new GetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring() *GetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring {
	this := GetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring{}
	return &this
}

// NewGetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoringWithDefaults instantiates a new GetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoringWithDefaults() *GetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring {
	this := GetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *GetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring) GetDuration() int32 {
	if o == nil || IsNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring) GetDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *GetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *GetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring) SetDuration(v int32) {
	o.Duration = &v
}

func (o GetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	return toSerialize, nil
}

type NullableGetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring struct {
	value *GetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring
	isSet bool
}

func (v NullableGetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring) Get() *GetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring {
	return v.value
}

func (v *NullableGetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring) Set(val *GetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring(val *GetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring) *NullableGetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring {
	return &NullableGetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring{value: val, isSet: true}
}

func (v NullableGetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


