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

// checks if the GetNetworkSensorAlertsProfiles200ResponseInnerRecipients type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSensorAlertsProfiles200ResponseInnerRecipients{}

// GetNetworkSensorAlertsProfiles200ResponseInnerRecipients List of recipients that will recieve the alert.
type GetNetworkSensorAlertsProfiles200ResponseInnerRecipients struct {
	// A list of emails that will receive information about the alert.
	Emails []string `json:"emails,omitempty"`
	// A list of SMS numbers that will receive information about the alert.
	SmsNumbers []string `json:"smsNumbers,omitempty"`
	// A list of webhook endpoint IDs that will receive information about the alert.
	HttpServerIds []string `json:"httpServerIds,omitempty"`
}

// NewGetNetworkSensorAlertsProfiles200ResponseInnerRecipients instantiates a new GetNetworkSensorAlertsProfiles200ResponseInnerRecipients object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSensorAlertsProfiles200ResponseInnerRecipients() *GetNetworkSensorAlertsProfiles200ResponseInnerRecipients {
	this := GetNetworkSensorAlertsProfiles200ResponseInnerRecipients{}
	return &this
}

// NewGetNetworkSensorAlertsProfiles200ResponseInnerRecipientsWithDefaults instantiates a new GetNetworkSensorAlertsProfiles200ResponseInnerRecipients object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSensorAlertsProfiles200ResponseInnerRecipientsWithDefaults() *GetNetworkSensorAlertsProfiles200ResponseInnerRecipients {
	this := GetNetworkSensorAlertsProfiles200ResponseInnerRecipients{}
	return &this
}

// GetEmails returns the Emails field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerRecipients) GetEmails() []string {
	if o == nil || IsNil(o.Emails) {
		var ret []string
		return ret
	}
	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerRecipients) GetEmailsOk() ([]string, bool) {
	if o == nil || IsNil(o.Emails) {
		return nil, false
	}
	return o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerRecipients) HasEmails() bool {
	if o != nil && !IsNil(o.Emails) {
		return true
	}

	return false
}

// SetEmails gets a reference to the given []string and assigns it to the Emails field.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerRecipients) SetEmails(v []string) {
	o.Emails = v
}

// GetSmsNumbers returns the SmsNumbers field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerRecipients) GetSmsNumbers() []string {
	if o == nil || IsNil(o.SmsNumbers) {
		var ret []string
		return ret
	}
	return o.SmsNumbers
}

// GetSmsNumbersOk returns a tuple with the SmsNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerRecipients) GetSmsNumbersOk() ([]string, bool) {
	if o == nil || IsNil(o.SmsNumbers) {
		return nil, false
	}
	return o.SmsNumbers, true
}

// HasSmsNumbers returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerRecipients) HasSmsNumbers() bool {
	if o != nil && !IsNil(o.SmsNumbers) {
		return true
	}

	return false
}

// SetSmsNumbers gets a reference to the given []string and assigns it to the SmsNumbers field.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerRecipients) SetSmsNumbers(v []string) {
	o.SmsNumbers = v
}

// GetHttpServerIds returns the HttpServerIds field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerRecipients) GetHttpServerIds() []string {
	if o == nil || IsNil(o.HttpServerIds) {
		var ret []string
		return ret
	}
	return o.HttpServerIds
}

// GetHttpServerIdsOk returns a tuple with the HttpServerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerRecipients) GetHttpServerIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.HttpServerIds) {
		return nil, false
	}
	return o.HttpServerIds, true
}

// HasHttpServerIds returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerRecipients) HasHttpServerIds() bool {
	if o != nil && !IsNil(o.HttpServerIds) {
		return true
	}

	return false
}

// SetHttpServerIds gets a reference to the given []string and assigns it to the HttpServerIds field.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerRecipients) SetHttpServerIds(v []string) {
	o.HttpServerIds = v
}

func (o GetNetworkSensorAlertsProfiles200ResponseInnerRecipients) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSensorAlertsProfiles200ResponseInnerRecipients) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Emails) {
		toSerialize["emails"] = o.Emails
	}
	if !IsNil(o.SmsNumbers) {
		toSerialize["smsNumbers"] = o.SmsNumbers
	}
	if !IsNil(o.HttpServerIds) {
		toSerialize["httpServerIds"] = o.HttpServerIds
	}
	return toSerialize, nil
}

type NullableGetNetworkSensorAlertsProfiles200ResponseInnerRecipients struct {
	value *GetNetworkSensorAlertsProfiles200ResponseInnerRecipients
	isSet bool
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerRecipients) Get() *GetNetworkSensorAlertsProfiles200ResponseInnerRecipients {
	return v.value
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerRecipients) Set(val *GetNetworkSensorAlertsProfiles200ResponseInnerRecipients) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerRecipients) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerRecipients) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSensorAlertsProfiles200ResponseInnerRecipients(val *GetNetworkSensorAlertsProfiles200ResponseInnerRecipients) *NullableGetNetworkSensorAlertsProfiles200ResponseInnerRecipients {
	return &NullableGetNetworkSensorAlertsProfiles200ResponseInnerRecipients{value: val, isSet: true}
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerRecipients) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerRecipients) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


