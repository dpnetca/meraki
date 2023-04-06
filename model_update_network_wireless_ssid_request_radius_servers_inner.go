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

// checks if the UpdateNetworkWirelessSsidRequestRadiusServersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidRequestRadiusServersInner{}

// UpdateNetworkWirelessSsidRequestRadiusServersInner struct for UpdateNetworkWirelessSsidRequestRadiusServersInner
type UpdateNetworkWirelessSsidRequestRadiusServersInner struct {
	// IP address of your RADIUS server
	Host string `json:"host"`
	// UDP port the RADIUS server listens on for Access-requests
	Port *int32 `json:"port,omitempty"`
	// RADIUS client shared secret
	Secret *string `json:"secret,omitempty"`
	// Use RADSEC (TLS over TCP) to connect to this RADIUS server. Requires radiusProxyEnabled.
	RadsecEnabled *bool `json:"radsecEnabled,omitempty"`
	// The ID of the Openroaming Certificate attached to radius server.
	OpenRoamingCertificateId *int32 `json:"openRoamingCertificateId,omitempty"`
	// Certificate used for authorization for the RADSEC Server
	CaCertificate *string `json:"caCertificate,omitempty"`
}

// NewUpdateNetworkWirelessSsidRequestRadiusServersInner instantiates a new UpdateNetworkWirelessSsidRequestRadiusServersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidRequestRadiusServersInner(host string) *UpdateNetworkWirelessSsidRequestRadiusServersInner {
	this := UpdateNetworkWirelessSsidRequestRadiusServersInner{}
	this.Host = host
	return &this
}

// NewUpdateNetworkWirelessSsidRequestRadiusServersInnerWithDefaults instantiates a new UpdateNetworkWirelessSsidRequestRadiusServersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidRequestRadiusServersInnerWithDefaults() *UpdateNetworkWirelessSsidRequestRadiusServersInner {
	this := UpdateNetworkWirelessSsidRequestRadiusServersInner{}
	return &this
}

// GetHost returns the Host field value
func (o *UpdateNetworkWirelessSsidRequestRadiusServersInner) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestRadiusServersInner) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *UpdateNetworkWirelessSsidRequestRadiusServersInner) SetHost(v string) {
	o.Host = v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestRadiusServersInner) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestRadiusServersInner) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestRadiusServersInner) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *UpdateNetworkWirelessSsidRequestRadiusServersInner) SetPort(v int32) {
	o.Port = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestRadiusServersInner) GetSecret() string {
	if o == nil || IsNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestRadiusServersInner) GetSecretOk() (*string, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestRadiusServersInner) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *UpdateNetworkWirelessSsidRequestRadiusServersInner) SetSecret(v string) {
	o.Secret = &v
}

// GetRadsecEnabled returns the RadsecEnabled field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestRadiusServersInner) GetRadsecEnabled() bool {
	if o == nil || IsNil(o.RadsecEnabled) {
		var ret bool
		return ret
	}
	return *o.RadsecEnabled
}

// GetRadsecEnabledOk returns a tuple with the RadsecEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestRadiusServersInner) GetRadsecEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.RadsecEnabled) {
		return nil, false
	}
	return o.RadsecEnabled, true
}

// HasRadsecEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestRadiusServersInner) HasRadsecEnabled() bool {
	if o != nil && !IsNil(o.RadsecEnabled) {
		return true
	}

	return false
}

// SetRadsecEnabled gets a reference to the given bool and assigns it to the RadsecEnabled field.
func (o *UpdateNetworkWirelessSsidRequestRadiusServersInner) SetRadsecEnabled(v bool) {
	o.RadsecEnabled = &v
}

// GetOpenRoamingCertificateId returns the OpenRoamingCertificateId field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestRadiusServersInner) GetOpenRoamingCertificateId() int32 {
	if o == nil || IsNil(o.OpenRoamingCertificateId) {
		var ret int32
		return ret
	}
	return *o.OpenRoamingCertificateId
}

// GetOpenRoamingCertificateIdOk returns a tuple with the OpenRoamingCertificateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestRadiusServersInner) GetOpenRoamingCertificateIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OpenRoamingCertificateId) {
		return nil, false
	}
	return o.OpenRoamingCertificateId, true
}

// HasOpenRoamingCertificateId returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestRadiusServersInner) HasOpenRoamingCertificateId() bool {
	if o != nil && !IsNil(o.OpenRoamingCertificateId) {
		return true
	}

	return false
}

// SetOpenRoamingCertificateId gets a reference to the given int32 and assigns it to the OpenRoamingCertificateId field.
func (o *UpdateNetworkWirelessSsidRequestRadiusServersInner) SetOpenRoamingCertificateId(v int32) {
	o.OpenRoamingCertificateId = &v
}

// GetCaCertificate returns the CaCertificate field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestRadiusServersInner) GetCaCertificate() string {
	if o == nil || IsNil(o.CaCertificate) {
		var ret string
		return ret
	}
	return *o.CaCertificate
}

// GetCaCertificateOk returns a tuple with the CaCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestRadiusServersInner) GetCaCertificateOk() (*string, bool) {
	if o == nil || IsNil(o.CaCertificate) {
		return nil, false
	}
	return o.CaCertificate, true
}

// HasCaCertificate returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestRadiusServersInner) HasCaCertificate() bool {
	if o != nil && !IsNil(o.CaCertificate) {
		return true
	}

	return false
}

// SetCaCertificate gets a reference to the given string and assigns it to the CaCertificate field.
func (o *UpdateNetworkWirelessSsidRequestRadiusServersInner) SetCaCertificate(v string) {
	o.CaCertificate = &v
}

func (o UpdateNetworkWirelessSsidRequestRadiusServersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidRequestRadiusServersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["host"] = o.Host
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	if !IsNil(o.RadsecEnabled) {
		toSerialize["radsecEnabled"] = o.RadsecEnabled
	}
	if !IsNil(o.OpenRoamingCertificateId) {
		toSerialize["openRoamingCertificateId"] = o.OpenRoamingCertificateId
	}
	if !IsNil(o.CaCertificate) {
		toSerialize["caCertificate"] = o.CaCertificate
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidRequestRadiusServersInner struct {
	value *UpdateNetworkWirelessSsidRequestRadiusServersInner
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidRequestRadiusServersInner) Get() *UpdateNetworkWirelessSsidRequestRadiusServersInner {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidRequestRadiusServersInner) Set(val *UpdateNetworkWirelessSsidRequestRadiusServersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidRequestRadiusServersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidRequestRadiusServersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidRequestRadiusServersInner(val *UpdateNetworkWirelessSsidRequestRadiusServersInner) *NullableUpdateNetworkWirelessSsidRequestRadiusServersInner {
	return &NullableUpdateNetworkWirelessSsidRequestRadiusServersInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidRequestRadiusServersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidRequestRadiusServersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


