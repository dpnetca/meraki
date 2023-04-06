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

// checks if the GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue{}

// GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue Value of traffic filter
type GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue struct {
	// ID of 'applicationCategory' or 'application' type traffic filter
	Id *string `json:"id,omitempty"`
	// Protocol of 'custom' type traffic filter. Must be one of: 'tcp', 'udp', 'icmp', 'icmp6' or 'any'
	Protocol *string `json:"protocol,omitempty"`
	Source *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource `json:"source,omitempty"`
	Destination *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestination `json:"destination,omitempty"`
}

// NewGetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue instantiates a new GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue() *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue {
	this := GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue{}
	return &this
}

// NewGetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValueWithDefaults instantiates a new GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValueWithDefaults() *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue {
	this := GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) SetId(v string) {
	o.Id = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) SetProtocol(v string) {
	o.Protocol = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) GetSource() GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource {
	if o == nil || IsNil(o.Source) {
		var ret GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) GetSourceOk() (*GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource and assigns it to the Source field.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) SetSource(v GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource) {
	o.Source = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) GetDestination() GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestination {
	if o == nil || IsNil(o.Destination) {
		var ret GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestination
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) GetDestinationOk() (*GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestination, bool) {
	if o == nil || IsNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) HasDestination() bool {
	if o != nil && !IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestination and assigns it to the Destination field.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) SetDestination(v GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestination) {
	o.Destination = &v
}

func (o GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	return toSerialize, nil
}

type NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue struct {
	value *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue
	isSet bool
}

func (v NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) Get() *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue {
	return v.value
}

func (v *NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) Set(val *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue(val *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) *NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue {
	return &NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue{value: val, isSet: true}
}

func (v NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


