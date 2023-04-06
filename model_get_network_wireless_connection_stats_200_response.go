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

// checks if the GetNetworkWirelessConnectionStats200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkWirelessConnectionStats200Response{}

// GetNetworkWirelessConnectionStats200Response struct for GetNetworkWirelessConnectionStats200Response
type GetNetworkWirelessConnectionStats200Response struct {
	// The number of failed association attempts
	Assoc *int32 `json:"assoc,omitempty"`
	// The number of failed authentication attempts
	Auth *int32 `json:"auth,omitempty"`
	// The number of failed DHCP attempts
	Dhcp *int32 `json:"dhcp,omitempty"`
	// The number of failed DNS attempts
	Dns *int32 `json:"dns,omitempty"`
	// The number of successful connection attempts
	Success *int32 `json:"success,omitempty"`
}

// NewGetNetworkWirelessConnectionStats200Response instantiates a new GetNetworkWirelessConnectionStats200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWirelessConnectionStats200Response() *GetNetworkWirelessConnectionStats200Response {
	this := GetNetworkWirelessConnectionStats200Response{}
	return &this
}

// NewGetNetworkWirelessConnectionStats200ResponseWithDefaults instantiates a new GetNetworkWirelessConnectionStats200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWirelessConnectionStats200ResponseWithDefaults() *GetNetworkWirelessConnectionStats200Response {
	this := GetNetworkWirelessConnectionStats200Response{}
	return &this
}

// GetAssoc returns the Assoc field value if set, zero value otherwise.
func (o *GetNetworkWirelessConnectionStats200Response) GetAssoc() int32 {
	if o == nil || IsNil(o.Assoc) {
		var ret int32
		return ret
	}
	return *o.Assoc
}

// GetAssocOk returns a tuple with the Assoc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessConnectionStats200Response) GetAssocOk() (*int32, bool) {
	if o == nil || IsNil(o.Assoc) {
		return nil, false
	}
	return o.Assoc, true
}

// HasAssoc returns a boolean if a field has been set.
func (o *GetNetworkWirelessConnectionStats200Response) HasAssoc() bool {
	if o != nil && !IsNil(o.Assoc) {
		return true
	}

	return false
}

// SetAssoc gets a reference to the given int32 and assigns it to the Assoc field.
func (o *GetNetworkWirelessConnectionStats200Response) SetAssoc(v int32) {
	o.Assoc = &v
}

// GetAuth returns the Auth field value if set, zero value otherwise.
func (o *GetNetworkWirelessConnectionStats200Response) GetAuth() int32 {
	if o == nil || IsNil(o.Auth) {
		var ret int32
		return ret
	}
	return *o.Auth
}

// GetAuthOk returns a tuple with the Auth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessConnectionStats200Response) GetAuthOk() (*int32, bool) {
	if o == nil || IsNil(o.Auth) {
		return nil, false
	}
	return o.Auth, true
}

// HasAuth returns a boolean if a field has been set.
func (o *GetNetworkWirelessConnectionStats200Response) HasAuth() bool {
	if o != nil && !IsNil(o.Auth) {
		return true
	}

	return false
}

// SetAuth gets a reference to the given int32 and assigns it to the Auth field.
func (o *GetNetworkWirelessConnectionStats200Response) SetAuth(v int32) {
	o.Auth = &v
}

// GetDhcp returns the Dhcp field value if set, zero value otherwise.
func (o *GetNetworkWirelessConnectionStats200Response) GetDhcp() int32 {
	if o == nil || IsNil(o.Dhcp) {
		var ret int32
		return ret
	}
	return *o.Dhcp
}

// GetDhcpOk returns a tuple with the Dhcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessConnectionStats200Response) GetDhcpOk() (*int32, bool) {
	if o == nil || IsNil(o.Dhcp) {
		return nil, false
	}
	return o.Dhcp, true
}

// HasDhcp returns a boolean if a field has been set.
func (o *GetNetworkWirelessConnectionStats200Response) HasDhcp() bool {
	if o != nil && !IsNil(o.Dhcp) {
		return true
	}

	return false
}

// SetDhcp gets a reference to the given int32 and assigns it to the Dhcp field.
func (o *GetNetworkWirelessConnectionStats200Response) SetDhcp(v int32) {
	o.Dhcp = &v
}

// GetDns returns the Dns field value if set, zero value otherwise.
func (o *GetNetworkWirelessConnectionStats200Response) GetDns() int32 {
	if o == nil || IsNil(o.Dns) {
		var ret int32
		return ret
	}
	return *o.Dns
}

// GetDnsOk returns a tuple with the Dns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessConnectionStats200Response) GetDnsOk() (*int32, bool) {
	if o == nil || IsNil(o.Dns) {
		return nil, false
	}
	return o.Dns, true
}

// HasDns returns a boolean if a field has been set.
func (o *GetNetworkWirelessConnectionStats200Response) HasDns() bool {
	if o != nil && !IsNil(o.Dns) {
		return true
	}

	return false
}

// SetDns gets a reference to the given int32 and assigns it to the Dns field.
func (o *GetNetworkWirelessConnectionStats200Response) SetDns(v int32) {
	o.Dns = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *GetNetworkWirelessConnectionStats200Response) GetSuccess() int32 {
	if o == nil || IsNil(o.Success) {
		var ret int32
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessConnectionStats200Response) GetSuccessOk() (*int32, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *GetNetworkWirelessConnectionStats200Response) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given int32 and assigns it to the Success field.
func (o *GetNetworkWirelessConnectionStats200Response) SetSuccess(v int32) {
	o.Success = &v
}

func (o GetNetworkWirelessConnectionStats200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkWirelessConnectionStats200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Assoc) {
		toSerialize["assoc"] = o.Assoc
	}
	if !IsNil(o.Auth) {
		toSerialize["auth"] = o.Auth
	}
	if !IsNil(o.Dhcp) {
		toSerialize["dhcp"] = o.Dhcp
	}
	if !IsNil(o.Dns) {
		toSerialize["dns"] = o.Dns
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return toSerialize, nil
}

type NullableGetNetworkWirelessConnectionStats200Response struct {
	value *GetNetworkWirelessConnectionStats200Response
	isSet bool
}

func (v NullableGetNetworkWirelessConnectionStats200Response) Get() *GetNetworkWirelessConnectionStats200Response {
	return v.value
}

func (v *NullableGetNetworkWirelessConnectionStats200Response) Set(val *GetNetworkWirelessConnectionStats200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWirelessConnectionStats200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWirelessConnectionStats200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWirelessConnectionStats200Response(val *GetNetworkWirelessConnectionStats200Response) *NullableGetNetworkWirelessConnectionStats200Response {
	return &NullableGetNetworkWirelessConnectionStats200Response{value: val, isSet: true}
}

func (v NullableGetNetworkWirelessConnectionStats200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWirelessConnectionStats200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


