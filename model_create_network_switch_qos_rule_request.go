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

// checks if the CreateNetworkSwitchQosRuleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkSwitchQosRuleRequest{}

// CreateNetworkSwitchQosRuleRequest struct for CreateNetworkSwitchQosRuleRequest
type CreateNetworkSwitchQosRuleRequest struct {
	// The VLAN of the incoming packet. A null value will match any VLAN.
	Vlan int32 `json:"vlan"`
	// The protocol of the incoming packet. Can be one of \"ANY\", \"TCP\" or \"UDP\". Default value is \"ANY\"
	Protocol *string `json:"protocol,omitempty"`
	// The source port of the incoming packet. Applicable only if protocol is TCP or UDP.
	SrcPort *int32 `json:"srcPort,omitempty"`
	// The source port range of the incoming packet. Applicable only if protocol is set to TCP or UDP. Example: 70-80
	SrcPortRange *string `json:"srcPortRange,omitempty"`
	// The destination port of the incoming packet. Applicable only if protocol is TCP or UDP.
	DstPort *int32 `json:"dstPort,omitempty"`
	// The destination port range of the incoming packet. Applicable only if protocol is set to TCP or UDP. Example: 70-80
	DstPortRange *string `json:"dstPortRange,omitempty"`
	// DSCP tag. Set this to -1 to trust incoming DSCP. Default value is 0
	Dscp *int32 `json:"dscp,omitempty"`
}

// NewCreateNetworkSwitchQosRuleRequest instantiates a new CreateNetworkSwitchQosRuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkSwitchQosRuleRequest(vlan int32) *CreateNetworkSwitchQosRuleRequest {
	this := CreateNetworkSwitchQosRuleRequest{}
	this.Vlan = vlan
	return &this
}

// NewCreateNetworkSwitchQosRuleRequestWithDefaults instantiates a new CreateNetworkSwitchQosRuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkSwitchQosRuleRequestWithDefaults() *CreateNetworkSwitchQosRuleRequest {
	this := CreateNetworkSwitchQosRuleRequest{}
	return &this
}

// GetVlan returns the Vlan field value
func (o *CreateNetworkSwitchQosRuleRequest) GetVlan() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchQosRuleRequest) GetVlanOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vlan, true
}

// SetVlan sets field value
func (o *CreateNetworkSwitchQosRuleRequest) SetVlan(v int32) {
	o.Vlan = v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *CreateNetworkSwitchQosRuleRequest) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchQosRuleRequest) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *CreateNetworkSwitchQosRuleRequest) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *CreateNetworkSwitchQosRuleRequest) SetProtocol(v string) {
	o.Protocol = &v
}

// GetSrcPort returns the SrcPort field value if set, zero value otherwise.
func (o *CreateNetworkSwitchQosRuleRequest) GetSrcPort() int32 {
	if o == nil || IsNil(o.SrcPort) {
		var ret int32
		return ret
	}
	return *o.SrcPort
}

// GetSrcPortOk returns a tuple with the SrcPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchQosRuleRequest) GetSrcPortOk() (*int32, bool) {
	if o == nil || IsNil(o.SrcPort) {
		return nil, false
	}
	return o.SrcPort, true
}

// HasSrcPort returns a boolean if a field has been set.
func (o *CreateNetworkSwitchQosRuleRequest) HasSrcPort() bool {
	if o != nil && !IsNil(o.SrcPort) {
		return true
	}

	return false
}

// SetSrcPort gets a reference to the given int32 and assigns it to the SrcPort field.
func (o *CreateNetworkSwitchQosRuleRequest) SetSrcPort(v int32) {
	o.SrcPort = &v
}

// GetSrcPortRange returns the SrcPortRange field value if set, zero value otherwise.
func (o *CreateNetworkSwitchQosRuleRequest) GetSrcPortRange() string {
	if o == nil || IsNil(o.SrcPortRange) {
		var ret string
		return ret
	}
	return *o.SrcPortRange
}

// GetSrcPortRangeOk returns a tuple with the SrcPortRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchQosRuleRequest) GetSrcPortRangeOk() (*string, bool) {
	if o == nil || IsNil(o.SrcPortRange) {
		return nil, false
	}
	return o.SrcPortRange, true
}

// HasSrcPortRange returns a boolean if a field has been set.
func (o *CreateNetworkSwitchQosRuleRequest) HasSrcPortRange() bool {
	if o != nil && !IsNil(o.SrcPortRange) {
		return true
	}

	return false
}

// SetSrcPortRange gets a reference to the given string and assigns it to the SrcPortRange field.
func (o *CreateNetworkSwitchQosRuleRequest) SetSrcPortRange(v string) {
	o.SrcPortRange = &v
}

// GetDstPort returns the DstPort field value if set, zero value otherwise.
func (o *CreateNetworkSwitchQosRuleRequest) GetDstPort() int32 {
	if o == nil || IsNil(o.DstPort) {
		var ret int32
		return ret
	}
	return *o.DstPort
}

// GetDstPortOk returns a tuple with the DstPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchQosRuleRequest) GetDstPortOk() (*int32, bool) {
	if o == nil || IsNil(o.DstPort) {
		return nil, false
	}
	return o.DstPort, true
}

// HasDstPort returns a boolean if a field has been set.
func (o *CreateNetworkSwitchQosRuleRequest) HasDstPort() bool {
	if o != nil && !IsNil(o.DstPort) {
		return true
	}

	return false
}

// SetDstPort gets a reference to the given int32 and assigns it to the DstPort field.
func (o *CreateNetworkSwitchQosRuleRequest) SetDstPort(v int32) {
	o.DstPort = &v
}

// GetDstPortRange returns the DstPortRange field value if set, zero value otherwise.
func (o *CreateNetworkSwitchQosRuleRequest) GetDstPortRange() string {
	if o == nil || IsNil(o.DstPortRange) {
		var ret string
		return ret
	}
	return *o.DstPortRange
}

// GetDstPortRangeOk returns a tuple with the DstPortRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchQosRuleRequest) GetDstPortRangeOk() (*string, bool) {
	if o == nil || IsNil(o.DstPortRange) {
		return nil, false
	}
	return o.DstPortRange, true
}

// HasDstPortRange returns a boolean if a field has been set.
func (o *CreateNetworkSwitchQosRuleRequest) HasDstPortRange() bool {
	if o != nil && !IsNil(o.DstPortRange) {
		return true
	}

	return false
}

// SetDstPortRange gets a reference to the given string and assigns it to the DstPortRange field.
func (o *CreateNetworkSwitchQosRuleRequest) SetDstPortRange(v string) {
	o.DstPortRange = &v
}

// GetDscp returns the Dscp field value if set, zero value otherwise.
func (o *CreateNetworkSwitchQosRuleRequest) GetDscp() int32 {
	if o == nil || IsNil(o.Dscp) {
		var ret int32
		return ret
	}
	return *o.Dscp
}

// GetDscpOk returns a tuple with the Dscp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchQosRuleRequest) GetDscpOk() (*int32, bool) {
	if o == nil || IsNil(o.Dscp) {
		return nil, false
	}
	return o.Dscp, true
}

// HasDscp returns a boolean if a field has been set.
func (o *CreateNetworkSwitchQosRuleRequest) HasDscp() bool {
	if o != nil && !IsNil(o.Dscp) {
		return true
	}

	return false
}

// SetDscp gets a reference to the given int32 and assigns it to the Dscp field.
func (o *CreateNetworkSwitchQosRuleRequest) SetDscp(v int32) {
	o.Dscp = &v
}

func (o CreateNetworkSwitchQosRuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkSwitchQosRuleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["vlan"] = o.Vlan
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.SrcPort) {
		toSerialize["srcPort"] = o.SrcPort
	}
	if !IsNil(o.SrcPortRange) {
		toSerialize["srcPortRange"] = o.SrcPortRange
	}
	if !IsNil(o.DstPort) {
		toSerialize["dstPort"] = o.DstPort
	}
	if !IsNil(o.DstPortRange) {
		toSerialize["dstPortRange"] = o.DstPortRange
	}
	if !IsNil(o.Dscp) {
		toSerialize["dscp"] = o.Dscp
	}
	return toSerialize, nil
}

type NullableCreateNetworkSwitchQosRuleRequest struct {
	value *CreateNetworkSwitchQosRuleRequest
	isSet bool
}

func (v NullableCreateNetworkSwitchQosRuleRequest) Get() *CreateNetworkSwitchQosRuleRequest {
	return v.value
}

func (v *NullableCreateNetworkSwitchQosRuleRequest) Set(val *CreateNetworkSwitchQosRuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkSwitchQosRuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkSwitchQosRuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkSwitchQosRuleRequest(val *CreateNetworkSwitchQosRuleRequest) *NullableCreateNetworkSwitchQosRuleRequest {
	return &NullableCreateNetworkSwitchQosRuleRequest{value: val, isSet: true}
}

func (v NullableCreateNetworkSwitchQosRuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkSwitchQosRuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


