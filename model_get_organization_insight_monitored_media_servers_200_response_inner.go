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

// checks if the GetOrganizationInsightMonitoredMediaServers200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationInsightMonitoredMediaServers200ResponseInner{}

// GetOrganizationInsightMonitoredMediaServers200ResponseInner struct for GetOrganizationInsightMonitoredMediaServers200ResponseInner
type GetOrganizationInsightMonitoredMediaServers200ResponseInner struct {
	// Monitored media server id
	Id *string `json:"id,omitempty"`
	// The name of the VoIP provider
	Name *string `json:"name,omitempty"`
	// The IP address (IPv4 only) or hostname of the media server to monitor
	Address *string `json:"address,omitempty"`
	// Indicates that if the media server doesn't respond to ICMP pings, the nearest hop will be used in its stead
	BestEffortMonitoringEnabled *bool `json:"bestEffortMonitoringEnabled,omitempty"`
}

// NewGetOrganizationInsightMonitoredMediaServers200ResponseInner instantiates a new GetOrganizationInsightMonitoredMediaServers200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationInsightMonitoredMediaServers200ResponseInner() *GetOrganizationInsightMonitoredMediaServers200ResponseInner {
	this := GetOrganizationInsightMonitoredMediaServers200ResponseInner{}
	return &this
}

// NewGetOrganizationInsightMonitoredMediaServers200ResponseInnerWithDefaults instantiates a new GetOrganizationInsightMonitoredMediaServers200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationInsightMonitoredMediaServers200ResponseInnerWithDefaults() *GetOrganizationInsightMonitoredMediaServers200ResponseInner {
	this := GetOrganizationInsightMonitoredMediaServers200ResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetOrganizationInsightMonitoredMediaServers200ResponseInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationInsightMonitoredMediaServers200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetOrganizationInsightMonitoredMediaServers200ResponseInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetOrganizationInsightMonitoredMediaServers200ResponseInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizationInsightMonitoredMediaServers200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationInsightMonitoredMediaServers200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizationInsightMonitoredMediaServers200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizationInsightMonitoredMediaServers200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *GetOrganizationInsightMonitoredMediaServers200ResponseInner) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationInsightMonitoredMediaServers200ResponseInner) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *GetOrganizationInsightMonitoredMediaServers200ResponseInner) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *GetOrganizationInsightMonitoredMediaServers200ResponseInner) SetAddress(v string) {
	o.Address = &v
}

// GetBestEffortMonitoringEnabled returns the BestEffortMonitoringEnabled field value if set, zero value otherwise.
func (o *GetOrganizationInsightMonitoredMediaServers200ResponseInner) GetBestEffortMonitoringEnabled() bool {
	if o == nil || IsNil(o.BestEffortMonitoringEnabled) {
		var ret bool
		return ret
	}
	return *o.BestEffortMonitoringEnabled
}

// GetBestEffortMonitoringEnabledOk returns a tuple with the BestEffortMonitoringEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationInsightMonitoredMediaServers200ResponseInner) GetBestEffortMonitoringEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BestEffortMonitoringEnabled) {
		return nil, false
	}
	return o.BestEffortMonitoringEnabled, true
}

// HasBestEffortMonitoringEnabled returns a boolean if a field has been set.
func (o *GetOrganizationInsightMonitoredMediaServers200ResponseInner) HasBestEffortMonitoringEnabled() bool {
	if o != nil && !IsNil(o.BestEffortMonitoringEnabled) {
		return true
	}

	return false
}

// SetBestEffortMonitoringEnabled gets a reference to the given bool and assigns it to the BestEffortMonitoringEnabled field.
func (o *GetOrganizationInsightMonitoredMediaServers200ResponseInner) SetBestEffortMonitoringEnabled(v bool) {
	o.BestEffortMonitoringEnabled = &v
}

func (o GetOrganizationInsightMonitoredMediaServers200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationInsightMonitoredMediaServers200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.BestEffortMonitoringEnabled) {
		toSerialize["bestEffortMonitoringEnabled"] = o.BestEffortMonitoringEnabled
	}
	return toSerialize, nil
}

type NullableGetOrganizationInsightMonitoredMediaServers200ResponseInner struct {
	value *GetOrganizationInsightMonitoredMediaServers200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationInsightMonitoredMediaServers200ResponseInner) Get() *GetOrganizationInsightMonitoredMediaServers200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationInsightMonitoredMediaServers200ResponseInner) Set(val *GetOrganizationInsightMonitoredMediaServers200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationInsightMonitoredMediaServers200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationInsightMonitoredMediaServers200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationInsightMonitoredMediaServers200ResponseInner(val *GetOrganizationInsightMonitoredMediaServers200ResponseInner) *NullableGetOrganizationInsightMonitoredMediaServers200ResponseInner {
	return &NullableGetOrganizationInsightMonitoredMediaServers200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationInsightMonitoredMediaServers200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationInsightMonitoredMediaServers200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

