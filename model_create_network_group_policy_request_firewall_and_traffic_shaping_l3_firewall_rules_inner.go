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

// checks if the CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner{}

// CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner struct for CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner
type CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner struct {
	// Description of the rule (optional)
	Comment *string `json:"comment,omitempty"`
	// 'allow' or 'deny' traffic specified by this rule
	Policy string `json:"policy"`
	// The type of protocol (must be 'tcp', 'udp', 'icmp', 'icmp6' or 'any')
	Protocol string `json:"protocol"`
	// Destination port (integer in the range 1-65535), a port range (e.g. 8080-9090), or 'any'
	DestPort *string `json:"destPort,omitempty"`
	// Destination IP address (in IP or CIDR notation), a fully-qualified domain name (FQDN, if your network supports it) or 'any'.
	DestCidr string `json:"destCidr"`
}

// NewCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner instantiates a new CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner(policy string, protocol string, destCidr string) *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner {
	this := CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner{}
	this.Policy = policy
	this.Protocol = protocol
	this.DestCidr = destCidr
	return &this
}

// NewCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInnerWithDefaults instantiates a new CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInnerWithDefaults() *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner {
	this := CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) SetComment(v string) {
	o.Comment = &v
}

// GetPolicy returns the Policy field value
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) GetPolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) GetPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Policy, true
}

// SetPolicy sets field value
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) SetPolicy(v string) {
	o.Policy = v
}

// GetProtocol returns the Protocol field value
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) SetProtocol(v string) {
	o.Protocol = v
}

// GetDestPort returns the DestPort field value if set, zero value otherwise.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) GetDestPort() string {
	if o == nil || IsNil(o.DestPort) {
		var ret string
		return ret
	}
	return *o.DestPort
}

// GetDestPortOk returns a tuple with the DestPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) GetDestPortOk() (*string, bool) {
	if o == nil || IsNil(o.DestPort) {
		return nil, false
	}
	return o.DestPort, true
}

// HasDestPort returns a boolean if a field has been set.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) HasDestPort() bool {
	if o != nil && !IsNil(o.DestPort) {
		return true
	}

	return false
}

// SetDestPort gets a reference to the given string and assigns it to the DestPort field.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) SetDestPort(v string) {
	o.DestPort = &v
}

// GetDestCidr returns the DestCidr field value
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) GetDestCidr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestCidr
}

// GetDestCidrOk returns a tuple with the DestCidr field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) GetDestCidrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestCidr, true
}

// SetDestCidr sets field value
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) SetDestCidr(v string) {
	o.DestCidr = v
}

func (o CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	toSerialize["policy"] = o.Policy
	toSerialize["protocol"] = o.Protocol
	if !IsNil(o.DestPort) {
		toSerialize["destPort"] = o.DestPort
	}
	toSerialize["destCidr"] = o.DestCidr
	return toSerialize, nil
}

type NullableCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner struct {
	value *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner
	isSet bool
}

func (v NullableCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) Get() *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner {
	return v.value
}

func (v *NullableCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) Set(val *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner(val *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) *NullableCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner {
	return &NullableCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner{value: val, isSet: true}
}

func (v NullableCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


