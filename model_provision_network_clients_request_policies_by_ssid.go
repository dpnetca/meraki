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

// checks if the ProvisionNetworkClientsRequestPoliciesBySsid type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisionNetworkClientsRequestPoliciesBySsid{}

// ProvisionNetworkClientsRequestPoliciesBySsid An object, describing the policy-connection associations for each active SSID within the network. Keys should be the number of enabled SSIDs, mapping to an object describing the client's policy
type ProvisionNetworkClientsRequestPoliciesBySsid struct {
	Var0 *ProvisionNetworkClientsRequestPoliciesBySsid0 `json:"0,omitempty"`
	Var1 *ProvisionNetworkClientsRequestPoliciesBySsid0 `json:"1,omitempty"`
	Var2 *ProvisionNetworkClientsRequestPoliciesBySsid0 `json:"2,omitempty"`
	Var3 *ProvisionNetworkClientsRequestPoliciesBySsid0 `json:"3,omitempty"`
	Var4 *ProvisionNetworkClientsRequestPoliciesBySsid0 `json:"4,omitempty"`
	Var5 *ProvisionNetworkClientsRequestPoliciesBySsid0 `json:"5,omitempty"`
	Var6 *ProvisionNetworkClientsRequestPoliciesBySsid0 `json:"6,omitempty"`
	Var7 *ProvisionNetworkClientsRequestPoliciesBySsid0 `json:"7,omitempty"`
	Var8 *ProvisionNetworkClientsRequestPoliciesBySsid0 `json:"8,omitempty"`
	Var9 *ProvisionNetworkClientsRequestPoliciesBySsid0 `json:"9,omitempty"`
	Var10 *ProvisionNetworkClientsRequestPoliciesBySsid0 `json:"10,omitempty"`
	Var11 *ProvisionNetworkClientsRequestPoliciesBySsid0 `json:"11,omitempty"`
	Var12 *ProvisionNetworkClientsRequestPoliciesBySsid0 `json:"12,omitempty"`
	Var13 *ProvisionNetworkClientsRequestPoliciesBySsid0 `json:"13,omitempty"`
	Var14 *ProvisionNetworkClientsRequestPoliciesBySsid0 `json:"14,omitempty"`
}

// NewProvisionNetworkClientsRequestPoliciesBySsid instantiates a new ProvisionNetworkClientsRequestPoliciesBySsid object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisionNetworkClientsRequestPoliciesBySsid() *ProvisionNetworkClientsRequestPoliciesBySsid {
	this := ProvisionNetworkClientsRequestPoliciesBySsid{}
	return &this
}

// NewProvisionNetworkClientsRequestPoliciesBySsidWithDefaults instantiates a new ProvisionNetworkClientsRequestPoliciesBySsid object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisionNetworkClientsRequestPoliciesBySsidWithDefaults() *ProvisionNetworkClientsRequestPoliciesBySsid {
	this := ProvisionNetworkClientsRequestPoliciesBySsid{}
	return &this
}

// GetVar0 returns the Var0 field value if set, zero value otherwise.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) GetVar0() ProvisionNetworkClientsRequestPoliciesBySsid0 {
	if o == nil || IsNil(o.Var0) {
		var ret ProvisionNetworkClientsRequestPoliciesBySsid0
		return ret
	}
	return *o.Var0
}

// GetVar0Ok returns a tuple with the Var0 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) GetVar0Ok() (*ProvisionNetworkClientsRequestPoliciesBySsid0, bool) {
	if o == nil || IsNil(o.Var0) {
		return nil, false
	}
	return o.Var0, true
}

// HasVar0 returns a boolean if a field has been set.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) HasVar0() bool {
	if o != nil && !IsNil(o.Var0) {
		return true
	}

	return false
}

// SetVar0 gets a reference to the given ProvisionNetworkClientsRequestPoliciesBySsid0 and assigns it to the Var0 field.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) SetVar0(v ProvisionNetworkClientsRequestPoliciesBySsid0) {
	o.Var0 = &v
}

// GetVar1 returns the Var1 field value if set, zero value otherwise.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) GetVar1() ProvisionNetworkClientsRequestPoliciesBySsid0 {
	if o == nil || IsNil(o.Var1) {
		var ret ProvisionNetworkClientsRequestPoliciesBySsid0
		return ret
	}
	return *o.Var1
}

// GetVar1Ok returns a tuple with the Var1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) GetVar1Ok() (*ProvisionNetworkClientsRequestPoliciesBySsid0, bool) {
	if o == nil || IsNil(o.Var1) {
		return nil, false
	}
	return o.Var1, true
}

// HasVar1 returns a boolean if a field has been set.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) HasVar1() bool {
	if o != nil && !IsNil(o.Var1) {
		return true
	}

	return false
}

// SetVar1 gets a reference to the given ProvisionNetworkClientsRequestPoliciesBySsid0 and assigns it to the Var1 field.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) SetVar1(v ProvisionNetworkClientsRequestPoliciesBySsid0) {
	o.Var1 = &v
}

// GetVar2 returns the Var2 field value if set, zero value otherwise.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) GetVar2() ProvisionNetworkClientsRequestPoliciesBySsid0 {
	if o == nil || IsNil(o.Var2) {
		var ret ProvisionNetworkClientsRequestPoliciesBySsid0
		return ret
	}
	return *o.Var2
}

// GetVar2Ok returns a tuple with the Var2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) GetVar2Ok() (*ProvisionNetworkClientsRequestPoliciesBySsid0, bool) {
	if o == nil || IsNil(o.Var2) {
		return nil, false
	}
	return o.Var2, true
}

// HasVar2 returns a boolean if a field has been set.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) HasVar2() bool {
	if o != nil && !IsNil(o.Var2) {
		return true
	}

	return false
}

// SetVar2 gets a reference to the given ProvisionNetworkClientsRequestPoliciesBySsid0 and assigns it to the Var2 field.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) SetVar2(v ProvisionNetworkClientsRequestPoliciesBySsid0) {
	o.Var2 = &v
}

// GetVar3 returns the Var3 field value if set, zero value otherwise.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) GetVar3() ProvisionNetworkClientsRequestPoliciesBySsid0 {
	if o == nil || IsNil(o.Var3) {
		var ret ProvisionNetworkClientsRequestPoliciesBySsid0
		return ret
	}
	return *o.Var3
}

// GetVar3Ok returns a tuple with the Var3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) GetVar3Ok() (*ProvisionNetworkClientsRequestPoliciesBySsid0, bool) {
	if o == nil || IsNil(o.Var3) {
		return nil, false
	}
	return o.Var3, true
}

// HasVar3 returns a boolean if a field has been set.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) HasVar3() bool {
	if o != nil && !IsNil(o.Var3) {
		return true
	}

	return false
}

// SetVar3 gets a reference to the given ProvisionNetworkClientsRequestPoliciesBySsid0 and assigns it to the Var3 field.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) SetVar3(v ProvisionNetworkClientsRequestPoliciesBySsid0) {
	o.Var3 = &v
}

// GetVar4 returns the Var4 field value if set, zero value otherwise.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) GetVar4() ProvisionNetworkClientsRequestPoliciesBySsid0 {
	if o == nil || IsNil(o.Var4) {
		var ret ProvisionNetworkClientsRequestPoliciesBySsid0
		return ret
	}
	return *o.Var4
}

// GetVar4Ok returns a tuple with the Var4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) GetVar4Ok() (*ProvisionNetworkClientsRequestPoliciesBySsid0, bool) {
	if o == nil || IsNil(o.Var4) {
		return nil, false
	}
	return o.Var4, true
}

// HasVar4 returns a boolean if a field has been set.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) HasVar4() bool {
	if o != nil && !IsNil(o.Var4) {
		return true
	}

	return false
}

// SetVar4 gets a reference to the given ProvisionNetworkClientsRequestPoliciesBySsid0 and assigns it to the Var4 field.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) SetVar4(v ProvisionNetworkClientsRequestPoliciesBySsid0) {
	o.Var4 = &v
}

// GetVar5 returns the Var5 field value if set, zero value otherwise.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) GetVar5() ProvisionNetworkClientsRequestPoliciesBySsid0 {
	if o == nil || IsNil(o.Var5) {
		var ret ProvisionNetworkClientsRequestPoliciesBySsid0
		return ret
	}
	return *o.Var5
}

// GetVar5Ok returns a tuple with the Var5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) GetVar5Ok() (*ProvisionNetworkClientsRequestPoliciesBySsid0, bool) {
	if o == nil || IsNil(o.Var5) {
		return nil, false
	}
	return o.Var5, true
}

// HasVar5 returns a boolean if a field has been set.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) HasVar5() bool {
	if o != nil && !IsNil(o.Var5) {
		return true
	}

	return false
}

// SetVar5 gets a reference to the given ProvisionNetworkClientsRequestPoliciesBySsid0 and assigns it to the Var5 field.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) SetVar5(v ProvisionNetworkClientsRequestPoliciesBySsid0) {
	o.Var5 = &v
}

// GetVar6 returns the Var6 field value if set, zero value otherwise.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) GetVar6() ProvisionNetworkClientsRequestPoliciesBySsid0 {
	if o == nil || IsNil(o.Var6) {
		var ret ProvisionNetworkClientsRequestPoliciesBySsid0
		return ret
	}
	return *o.Var6
}

// GetVar6Ok returns a tuple with the Var6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) GetVar6Ok() (*ProvisionNetworkClientsRequestPoliciesBySsid0, bool) {
	if o == nil || IsNil(o.Var6) {
		return nil, false
	}
	return o.Var6, true
}

// HasVar6 returns a boolean if a field has been set.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) HasVar6() bool {
	if o != nil && !IsNil(o.Var6) {
		return true
	}

	return false
}

// SetVar6 gets a reference to the given ProvisionNetworkClientsRequestPoliciesBySsid0 and assigns it to the Var6 field.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) SetVar6(v ProvisionNetworkClientsRequestPoliciesBySsid0) {
	o.Var6 = &v
}

// GetVar7 returns the Var7 field value if set, zero value otherwise.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) GetVar7() ProvisionNetworkClientsRequestPoliciesBySsid0 {
	if o == nil || IsNil(o.Var7) {
		var ret ProvisionNetworkClientsRequestPoliciesBySsid0
		return ret
	}
	return *o.Var7
}

// GetVar7Ok returns a tuple with the Var7 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) GetVar7Ok() (*ProvisionNetworkClientsRequestPoliciesBySsid0, bool) {
	if o == nil || IsNil(o.Var7) {
		return nil, false
	}
	return o.Var7, true
}

// HasVar7 returns a boolean if a field has been set.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) HasVar7() bool {
	if o != nil && !IsNil(o.Var7) {
		return true
	}

	return false
}

// SetVar7 gets a reference to the given ProvisionNetworkClientsRequestPoliciesBySsid0 and assigns it to the Var7 field.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) SetVar7(v ProvisionNetworkClientsRequestPoliciesBySsid0) {
	o.Var7 = &v
}

// GetVar8 returns the Var8 field value if set, zero value otherwise.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) GetVar8() ProvisionNetworkClientsRequestPoliciesBySsid0 {
	if o == nil || IsNil(o.Var8) {
		var ret ProvisionNetworkClientsRequestPoliciesBySsid0
		return ret
	}
	return *o.Var8
}

// GetVar8Ok returns a tuple with the Var8 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) GetVar8Ok() (*ProvisionNetworkClientsRequestPoliciesBySsid0, bool) {
	if o == nil || IsNil(o.Var8) {
		return nil, false
	}
	return o.Var8, true
}

// HasVar8 returns a boolean if a field has been set.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) HasVar8() bool {
	if o != nil && !IsNil(o.Var8) {
		return true
	}

	return false
}

// SetVar8 gets a reference to the given ProvisionNetworkClientsRequestPoliciesBySsid0 and assigns it to the Var8 field.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) SetVar8(v ProvisionNetworkClientsRequestPoliciesBySsid0) {
	o.Var8 = &v
}

// GetVar9 returns the Var9 field value if set, zero value otherwise.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) GetVar9() ProvisionNetworkClientsRequestPoliciesBySsid0 {
	if o == nil || IsNil(o.Var9) {
		var ret ProvisionNetworkClientsRequestPoliciesBySsid0
		return ret
	}
	return *o.Var9
}

// GetVar9Ok returns a tuple with the Var9 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) GetVar9Ok() (*ProvisionNetworkClientsRequestPoliciesBySsid0, bool) {
	if o == nil || IsNil(o.Var9) {
		return nil, false
	}
	return o.Var9, true
}

// HasVar9 returns a boolean if a field has been set.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) HasVar9() bool {
	if o != nil && !IsNil(o.Var9) {
		return true
	}

	return false
}

// SetVar9 gets a reference to the given ProvisionNetworkClientsRequestPoliciesBySsid0 and assigns it to the Var9 field.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) SetVar9(v ProvisionNetworkClientsRequestPoliciesBySsid0) {
	o.Var9 = &v
}

// GetVar10 returns the Var10 field value if set, zero value otherwise.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) GetVar10() ProvisionNetworkClientsRequestPoliciesBySsid0 {
	if o == nil || IsNil(o.Var10) {
		var ret ProvisionNetworkClientsRequestPoliciesBySsid0
		return ret
	}
	return *o.Var10
}

// GetVar10Ok returns a tuple with the Var10 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) GetVar10Ok() (*ProvisionNetworkClientsRequestPoliciesBySsid0, bool) {
	if o == nil || IsNil(o.Var10) {
		return nil, false
	}
	return o.Var10, true
}

// HasVar10 returns a boolean if a field has been set.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) HasVar10() bool {
	if o != nil && !IsNil(o.Var10) {
		return true
	}

	return false
}

// SetVar10 gets a reference to the given ProvisionNetworkClientsRequestPoliciesBySsid0 and assigns it to the Var10 field.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) SetVar10(v ProvisionNetworkClientsRequestPoliciesBySsid0) {
	o.Var10 = &v
}

// GetVar11 returns the Var11 field value if set, zero value otherwise.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) GetVar11() ProvisionNetworkClientsRequestPoliciesBySsid0 {
	if o == nil || IsNil(o.Var11) {
		var ret ProvisionNetworkClientsRequestPoliciesBySsid0
		return ret
	}
	return *o.Var11
}

// GetVar11Ok returns a tuple with the Var11 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) GetVar11Ok() (*ProvisionNetworkClientsRequestPoliciesBySsid0, bool) {
	if o == nil || IsNil(o.Var11) {
		return nil, false
	}
	return o.Var11, true
}

// HasVar11 returns a boolean if a field has been set.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) HasVar11() bool {
	if o != nil && !IsNil(o.Var11) {
		return true
	}

	return false
}

// SetVar11 gets a reference to the given ProvisionNetworkClientsRequestPoliciesBySsid0 and assigns it to the Var11 field.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) SetVar11(v ProvisionNetworkClientsRequestPoliciesBySsid0) {
	o.Var11 = &v
}

// GetVar12 returns the Var12 field value if set, zero value otherwise.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) GetVar12() ProvisionNetworkClientsRequestPoliciesBySsid0 {
	if o == nil || IsNil(o.Var12) {
		var ret ProvisionNetworkClientsRequestPoliciesBySsid0
		return ret
	}
	return *o.Var12
}

// GetVar12Ok returns a tuple with the Var12 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) GetVar12Ok() (*ProvisionNetworkClientsRequestPoliciesBySsid0, bool) {
	if o == nil || IsNil(o.Var12) {
		return nil, false
	}
	return o.Var12, true
}

// HasVar12 returns a boolean if a field has been set.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) HasVar12() bool {
	if o != nil && !IsNil(o.Var12) {
		return true
	}

	return false
}

// SetVar12 gets a reference to the given ProvisionNetworkClientsRequestPoliciesBySsid0 and assigns it to the Var12 field.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) SetVar12(v ProvisionNetworkClientsRequestPoliciesBySsid0) {
	o.Var12 = &v
}

// GetVar13 returns the Var13 field value if set, zero value otherwise.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) GetVar13() ProvisionNetworkClientsRequestPoliciesBySsid0 {
	if o == nil || IsNil(o.Var13) {
		var ret ProvisionNetworkClientsRequestPoliciesBySsid0
		return ret
	}
	return *o.Var13
}

// GetVar13Ok returns a tuple with the Var13 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) GetVar13Ok() (*ProvisionNetworkClientsRequestPoliciesBySsid0, bool) {
	if o == nil || IsNil(o.Var13) {
		return nil, false
	}
	return o.Var13, true
}

// HasVar13 returns a boolean if a field has been set.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) HasVar13() bool {
	if o != nil && !IsNil(o.Var13) {
		return true
	}

	return false
}

// SetVar13 gets a reference to the given ProvisionNetworkClientsRequestPoliciesBySsid0 and assigns it to the Var13 field.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) SetVar13(v ProvisionNetworkClientsRequestPoliciesBySsid0) {
	o.Var13 = &v
}

// GetVar14 returns the Var14 field value if set, zero value otherwise.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) GetVar14() ProvisionNetworkClientsRequestPoliciesBySsid0 {
	if o == nil || IsNil(o.Var14) {
		var ret ProvisionNetworkClientsRequestPoliciesBySsid0
		return ret
	}
	return *o.Var14
}

// GetVar14Ok returns a tuple with the Var14 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) GetVar14Ok() (*ProvisionNetworkClientsRequestPoliciesBySsid0, bool) {
	if o == nil || IsNil(o.Var14) {
		return nil, false
	}
	return o.Var14, true
}

// HasVar14 returns a boolean if a field has been set.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) HasVar14() bool {
	if o != nil && !IsNil(o.Var14) {
		return true
	}

	return false
}

// SetVar14 gets a reference to the given ProvisionNetworkClientsRequestPoliciesBySsid0 and assigns it to the Var14 field.
func (o *ProvisionNetworkClientsRequestPoliciesBySsid) SetVar14(v ProvisionNetworkClientsRequestPoliciesBySsid0) {
	o.Var14 = &v
}

func (o ProvisionNetworkClientsRequestPoliciesBySsid) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisionNetworkClientsRequestPoliciesBySsid) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Var0) {
		toSerialize["0"] = o.Var0
	}
	if !IsNil(o.Var1) {
		toSerialize["1"] = o.Var1
	}
	if !IsNil(o.Var2) {
		toSerialize["2"] = o.Var2
	}
	if !IsNil(o.Var3) {
		toSerialize["3"] = o.Var3
	}
	if !IsNil(o.Var4) {
		toSerialize["4"] = o.Var4
	}
	if !IsNil(o.Var5) {
		toSerialize["5"] = o.Var5
	}
	if !IsNil(o.Var6) {
		toSerialize["6"] = o.Var6
	}
	if !IsNil(o.Var7) {
		toSerialize["7"] = o.Var7
	}
	if !IsNil(o.Var8) {
		toSerialize["8"] = o.Var8
	}
	if !IsNil(o.Var9) {
		toSerialize["9"] = o.Var9
	}
	if !IsNil(o.Var10) {
		toSerialize["10"] = o.Var10
	}
	if !IsNil(o.Var11) {
		toSerialize["11"] = o.Var11
	}
	if !IsNil(o.Var12) {
		toSerialize["12"] = o.Var12
	}
	if !IsNil(o.Var13) {
		toSerialize["13"] = o.Var13
	}
	if !IsNil(o.Var14) {
		toSerialize["14"] = o.Var14
	}
	return toSerialize, nil
}

type NullableProvisionNetworkClientsRequestPoliciesBySsid struct {
	value *ProvisionNetworkClientsRequestPoliciesBySsid
	isSet bool
}

func (v NullableProvisionNetworkClientsRequestPoliciesBySsid) Get() *ProvisionNetworkClientsRequestPoliciesBySsid {
	return v.value
}

func (v *NullableProvisionNetworkClientsRequestPoliciesBySsid) Set(val *ProvisionNetworkClientsRequestPoliciesBySsid) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisionNetworkClientsRequestPoliciesBySsid) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisionNetworkClientsRequestPoliciesBySsid) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisionNetworkClientsRequestPoliciesBySsid(val *ProvisionNetworkClientsRequestPoliciesBySsid) *NullableProvisionNetworkClientsRequestPoliciesBySsid {
	return &NullableProvisionNetworkClientsRequestPoliciesBySsid{value: val, isSet: true}
}

func (v NullableProvisionNetworkClientsRequestPoliciesBySsid) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisionNetworkClientsRequestPoliciesBySsid) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


