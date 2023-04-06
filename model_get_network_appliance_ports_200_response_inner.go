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

// checks if the GetNetworkAppliancePorts200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkAppliancePorts200ResponseInner{}

// GetNetworkAppliancePorts200ResponseInner struct for GetNetworkAppliancePorts200ResponseInner
type GetNetworkAppliancePorts200ResponseInner struct {
	// Number of the port
	Number *int32 `json:"number,omitempty"`
	// The status of the port
	Enabled *bool `json:"enabled,omitempty"`
	// The type of the port: 'access' or 'trunk'.
	Type *string `json:"type,omitempty"`
	// Whether the trunk port can drop all untagged traffic.
	DropUntaggedTraffic *bool `json:"dropUntaggedTraffic,omitempty"`
	// Native VLAN when the port is in Trunk mode. Access VLAN when the port is in Access mode.
	Vlan *int32 `json:"vlan,omitempty"`
	// Comma-delimited list of the VLAN ID's allowed on the port, or 'all' to permit all VLAN's on the port.
	AllowedVlans *string `json:"allowedVlans,omitempty"`
	// The name of the policy. Only applicable to Access ports.
	AccessPolicy *string `json:"accessPolicy,omitempty"`
}

// NewGetNetworkAppliancePorts200ResponseInner instantiates a new GetNetworkAppliancePorts200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkAppliancePorts200ResponseInner() *GetNetworkAppliancePorts200ResponseInner {
	this := GetNetworkAppliancePorts200ResponseInner{}
	return &this
}

// NewGetNetworkAppliancePorts200ResponseInnerWithDefaults instantiates a new GetNetworkAppliancePorts200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkAppliancePorts200ResponseInnerWithDefaults() *GetNetworkAppliancePorts200ResponseInner {
	this := GetNetworkAppliancePorts200ResponseInner{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *GetNetworkAppliancePorts200ResponseInner) GetNumber() int32 {
	if o == nil || IsNil(o.Number) {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkAppliancePorts200ResponseInner) GetNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *GetNetworkAppliancePorts200ResponseInner) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *GetNetworkAppliancePorts200ResponseInner) SetNumber(v int32) {
	o.Number = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetNetworkAppliancePorts200ResponseInner) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkAppliancePorts200ResponseInner) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetNetworkAppliancePorts200ResponseInner) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetNetworkAppliancePorts200ResponseInner) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetNetworkAppliancePorts200ResponseInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkAppliancePorts200ResponseInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetNetworkAppliancePorts200ResponseInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetNetworkAppliancePorts200ResponseInner) SetType(v string) {
	o.Type = &v
}

// GetDropUntaggedTraffic returns the DropUntaggedTraffic field value if set, zero value otherwise.
func (o *GetNetworkAppliancePorts200ResponseInner) GetDropUntaggedTraffic() bool {
	if o == nil || IsNil(o.DropUntaggedTraffic) {
		var ret bool
		return ret
	}
	return *o.DropUntaggedTraffic
}

// GetDropUntaggedTrafficOk returns a tuple with the DropUntaggedTraffic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkAppliancePorts200ResponseInner) GetDropUntaggedTrafficOk() (*bool, bool) {
	if o == nil || IsNil(o.DropUntaggedTraffic) {
		return nil, false
	}
	return o.DropUntaggedTraffic, true
}

// HasDropUntaggedTraffic returns a boolean if a field has been set.
func (o *GetNetworkAppliancePorts200ResponseInner) HasDropUntaggedTraffic() bool {
	if o != nil && !IsNil(o.DropUntaggedTraffic) {
		return true
	}

	return false
}

// SetDropUntaggedTraffic gets a reference to the given bool and assigns it to the DropUntaggedTraffic field.
func (o *GetNetworkAppliancePorts200ResponseInner) SetDropUntaggedTraffic(v bool) {
	o.DropUntaggedTraffic = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *GetNetworkAppliancePorts200ResponseInner) GetVlan() int32 {
	if o == nil || IsNil(o.Vlan) {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkAppliancePorts200ResponseInner) GetVlanOk() (*int32, bool) {
	if o == nil || IsNil(o.Vlan) {
		return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *GetNetworkAppliancePorts200ResponseInner) HasVlan() bool {
	if o != nil && !IsNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *GetNetworkAppliancePorts200ResponseInner) SetVlan(v int32) {
	o.Vlan = &v
}

// GetAllowedVlans returns the AllowedVlans field value if set, zero value otherwise.
func (o *GetNetworkAppliancePorts200ResponseInner) GetAllowedVlans() string {
	if o == nil || IsNil(o.AllowedVlans) {
		var ret string
		return ret
	}
	return *o.AllowedVlans
}

// GetAllowedVlansOk returns a tuple with the AllowedVlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkAppliancePorts200ResponseInner) GetAllowedVlansOk() (*string, bool) {
	if o == nil || IsNil(o.AllowedVlans) {
		return nil, false
	}
	return o.AllowedVlans, true
}

// HasAllowedVlans returns a boolean if a field has been set.
func (o *GetNetworkAppliancePorts200ResponseInner) HasAllowedVlans() bool {
	if o != nil && !IsNil(o.AllowedVlans) {
		return true
	}

	return false
}

// SetAllowedVlans gets a reference to the given string and assigns it to the AllowedVlans field.
func (o *GetNetworkAppliancePorts200ResponseInner) SetAllowedVlans(v string) {
	o.AllowedVlans = &v
}

// GetAccessPolicy returns the AccessPolicy field value if set, zero value otherwise.
func (o *GetNetworkAppliancePorts200ResponseInner) GetAccessPolicy() string {
	if o == nil || IsNil(o.AccessPolicy) {
		var ret string
		return ret
	}
	return *o.AccessPolicy
}

// GetAccessPolicyOk returns a tuple with the AccessPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkAppliancePorts200ResponseInner) GetAccessPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.AccessPolicy) {
		return nil, false
	}
	return o.AccessPolicy, true
}

// HasAccessPolicy returns a boolean if a field has been set.
func (o *GetNetworkAppliancePorts200ResponseInner) HasAccessPolicy() bool {
	if o != nil && !IsNil(o.AccessPolicy) {
		return true
	}

	return false
}

// SetAccessPolicy gets a reference to the given string and assigns it to the AccessPolicy field.
func (o *GetNetworkAppliancePorts200ResponseInner) SetAccessPolicy(v string) {
	o.AccessPolicy = &v
}

func (o GetNetworkAppliancePorts200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkAppliancePorts200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.DropUntaggedTraffic) {
		toSerialize["dropUntaggedTraffic"] = o.DropUntaggedTraffic
	}
	if !IsNil(o.Vlan) {
		toSerialize["vlan"] = o.Vlan
	}
	if !IsNil(o.AllowedVlans) {
		toSerialize["allowedVlans"] = o.AllowedVlans
	}
	if !IsNil(o.AccessPolicy) {
		toSerialize["accessPolicy"] = o.AccessPolicy
	}
	return toSerialize, nil
}

type NullableGetNetworkAppliancePorts200ResponseInner struct {
	value *GetNetworkAppliancePorts200ResponseInner
	isSet bool
}

func (v NullableGetNetworkAppliancePorts200ResponseInner) Get() *GetNetworkAppliancePorts200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkAppliancePorts200ResponseInner) Set(val *GetNetworkAppliancePorts200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkAppliancePorts200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkAppliancePorts200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkAppliancePorts200ResponseInner(val *GetNetworkAppliancePorts200ResponseInner) *NullableGetNetworkAppliancePorts200ResponseInner {
	return &NullableGetNetworkAppliancePorts200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkAppliancePorts200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkAppliancePorts200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


