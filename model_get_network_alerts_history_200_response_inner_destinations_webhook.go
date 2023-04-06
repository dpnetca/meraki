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

// checks if the GetNetworkAlertsHistory200ResponseInnerDestinationsWebhook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkAlertsHistory200ResponseInnerDestinationsWebhook{}

// GetNetworkAlertsHistory200ResponseInnerDestinationsWebhook webhook destinations for this alert
type GetNetworkAlertsHistory200ResponseInnerDestinationsWebhook struct {
	// time when the alert was sent to the user(s) for this channel
	SentAt *string `json:"sentAt,omitempty"`
}

// NewGetNetworkAlertsHistory200ResponseInnerDestinationsWebhook instantiates a new GetNetworkAlertsHistory200ResponseInnerDestinationsWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkAlertsHistory200ResponseInnerDestinationsWebhook() *GetNetworkAlertsHistory200ResponseInnerDestinationsWebhook {
	this := GetNetworkAlertsHistory200ResponseInnerDestinationsWebhook{}
	return &this
}

// NewGetNetworkAlertsHistory200ResponseInnerDestinationsWebhookWithDefaults instantiates a new GetNetworkAlertsHistory200ResponseInnerDestinationsWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkAlertsHistory200ResponseInnerDestinationsWebhookWithDefaults() *GetNetworkAlertsHistory200ResponseInnerDestinationsWebhook {
	this := GetNetworkAlertsHistory200ResponseInnerDestinationsWebhook{}
	return &this
}

// GetSentAt returns the SentAt field value if set, zero value otherwise.
func (o *GetNetworkAlertsHistory200ResponseInnerDestinationsWebhook) GetSentAt() string {
	if o == nil || IsNil(o.SentAt) {
		var ret string
		return ret
	}
	return *o.SentAt
}

// GetSentAtOk returns a tuple with the SentAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkAlertsHistory200ResponseInnerDestinationsWebhook) GetSentAtOk() (*string, bool) {
	if o == nil || IsNil(o.SentAt) {
		return nil, false
	}
	return o.SentAt, true
}

// HasSentAt returns a boolean if a field has been set.
func (o *GetNetworkAlertsHistory200ResponseInnerDestinationsWebhook) HasSentAt() bool {
	if o != nil && !IsNil(o.SentAt) {
		return true
	}

	return false
}

// SetSentAt gets a reference to the given string and assigns it to the SentAt field.
func (o *GetNetworkAlertsHistory200ResponseInnerDestinationsWebhook) SetSentAt(v string) {
	o.SentAt = &v
}

func (o GetNetworkAlertsHistory200ResponseInnerDestinationsWebhook) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkAlertsHistory200ResponseInnerDestinationsWebhook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SentAt) {
		toSerialize["sentAt"] = o.SentAt
	}
	return toSerialize, nil
}

type NullableGetNetworkAlertsHistory200ResponseInnerDestinationsWebhook struct {
	value *GetNetworkAlertsHistory200ResponseInnerDestinationsWebhook
	isSet bool
}

func (v NullableGetNetworkAlertsHistory200ResponseInnerDestinationsWebhook) Get() *GetNetworkAlertsHistory200ResponseInnerDestinationsWebhook {
	return v.value
}

func (v *NullableGetNetworkAlertsHistory200ResponseInnerDestinationsWebhook) Set(val *GetNetworkAlertsHistory200ResponseInnerDestinationsWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkAlertsHistory200ResponseInnerDestinationsWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkAlertsHistory200ResponseInnerDestinationsWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkAlertsHistory200ResponseInnerDestinationsWebhook(val *GetNetworkAlertsHistory200ResponseInnerDestinationsWebhook) *NullableGetNetworkAlertsHistory200ResponseInnerDestinationsWebhook {
	return &NullableGetNetworkAlertsHistory200ResponseInnerDestinationsWebhook{value: val, isSet: true}
}

func (v NullableGetNetworkAlertsHistory200ResponseInnerDestinationsWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkAlertsHistory200ResponseInnerDestinationsWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


