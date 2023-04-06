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

// checks if the GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner{}

// GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner struct for GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner
type GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner struct {
	Network *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerNetwork `json:"network,omitempty"`
	// Name of the switch
	Name *string `json:"name,omitempty"`
	// Mac address of the switch
	Mac *string `json:"mac,omitempty"`
	// Model of the switch
	Model *string `json:"model,omitempty"`
	Usage *GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerUsage `json:"usage,omitempty"`
}

// NewGetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner instantiates a new GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner() *GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner {
	this := GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner{}
	return &this
}

// NewGetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerWithDefaults instantiates a new GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerWithDefaults() *GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner {
	this := GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner{}
	return &this
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner) GetNetwork() GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerNetwork {
	if o == nil || IsNil(o.Network) {
		var ret GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner) GetNetworkOk() (*GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerNetwork, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerNetwork and assigns it to the Network field.
func (o *GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner) SetNetwork(v GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerNetwork) {
	o.Network = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner) GetMac() string {
	if o == nil || IsNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner) HasMac() bool {
	if o != nil && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner) SetMac(v string) {
	o.Mac = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner) SetModel(v string) {
	o.Model = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner) GetUsage() GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerUsage {
	if o == nil || IsNil(o.Usage) {
		var ret GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerUsage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner) GetUsageOk() (*GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerUsage, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerUsage and assigns it to the Usage field.
func (o *GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner) SetUsage(v GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerUsage) {
	o.Usage = &v
}

func (o GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	return toSerialize, nil
}

type NullableGetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner struct {
	value *GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner) Get() *GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner) Set(val *GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner(val *GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner) *NullableGetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner {
	return &NullableGetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


