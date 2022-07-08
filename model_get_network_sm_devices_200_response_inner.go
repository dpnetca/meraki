/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 July, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.23.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package meraki

import (
	"encoding/json"
)

// GetNetworkSmDevices200ResponseInner struct for GetNetworkSmDevices200ResponseInner
type GetNetworkSmDevices200ResponseInner struct {
	// The Meraki Id of the device record.
	Id *string `json:"id,omitempty"`
	// The name of the device.
	Name *string `json:"name,omitempty"`
	// An array of tags associated with the device.
	Tags []string `json:"tags,omitempty"`
	// The name of the SSID the device was last connected to.
	Ssid *string `json:"ssid,omitempty"`
	// The MAC of the device.
	WifiMac *string `json:"wifiMac,omitempty"`
	// The name of the device OS.
	OsName *string `json:"osName,omitempty"`
	// The device model.
	SystemModel *string `json:"systemModel,omitempty"`
	// The UUID of the device.
	Uuid *string `json:"uuid,omitempty"`
	// The device serial number.
	SerialNumber *string `json:"serialNumber,omitempty"`
	// The device serial.
	Serial *string `json:"serial,omitempty"`
	// The IP address of the device.
	Ip *string `json:"ip,omitempty"`
	// Notes associated with the device.
	Notes *string `json:"notes,omitempty"`
}

// NewGetNetworkSmDevices200ResponseInner instantiates a new GetNetworkSmDevices200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSmDevices200ResponseInner() *GetNetworkSmDevices200ResponseInner {
	this := GetNetworkSmDevices200ResponseInner{}
	return &this
}

// NewGetNetworkSmDevices200ResponseInnerWithDefaults instantiates a new GetNetworkSmDevices200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSmDevices200ResponseInnerWithDefaults() *GetNetworkSmDevices200ResponseInner {
	this := GetNetworkSmDevices200ResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNetworkSmDevices200ResponseInner) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDevices200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNetworkSmDevices200ResponseInner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetNetworkSmDevices200ResponseInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkSmDevices200ResponseInner) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDevices200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkSmDevices200ResponseInner) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkSmDevices200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GetNetworkSmDevices200ResponseInner) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDevices200ResponseInner) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GetNetworkSmDevices200ResponseInner) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GetNetworkSmDevices200ResponseInner) SetTags(v []string) {
	o.Tags = v
}

// GetSsid returns the Ssid field value if set, zero value otherwise.
func (o *GetNetworkSmDevices200ResponseInner) GetSsid() string {
	if o == nil || o.Ssid == nil {
		var ret string
		return ret
	}
	return *o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDevices200ResponseInner) GetSsidOk() (*string, bool) {
	if o == nil || o.Ssid == nil {
		return nil, false
	}
	return o.Ssid, true
}

// HasSsid returns a boolean if a field has been set.
func (o *GetNetworkSmDevices200ResponseInner) HasSsid() bool {
	if o != nil && o.Ssid != nil {
		return true
	}

	return false
}

// SetSsid gets a reference to the given string and assigns it to the Ssid field.
func (o *GetNetworkSmDevices200ResponseInner) SetSsid(v string) {
	o.Ssid = &v
}

// GetWifiMac returns the WifiMac field value if set, zero value otherwise.
func (o *GetNetworkSmDevices200ResponseInner) GetWifiMac() string {
	if o == nil || o.WifiMac == nil {
		var ret string
		return ret
	}
	return *o.WifiMac
}

// GetWifiMacOk returns a tuple with the WifiMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDevices200ResponseInner) GetWifiMacOk() (*string, bool) {
	if o == nil || o.WifiMac == nil {
		return nil, false
	}
	return o.WifiMac, true
}

// HasWifiMac returns a boolean if a field has been set.
func (o *GetNetworkSmDevices200ResponseInner) HasWifiMac() bool {
	if o != nil && o.WifiMac != nil {
		return true
	}

	return false
}

// SetWifiMac gets a reference to the given string and assigns it to the WifiMac field.
func (o *GetNetworkSmDevices200ResponseInner) SetWifiMac(v string) {
	o.WifiMac = &v
}

// GetOsName returns the OsName field value if set, zero value otherwise.
func (o *GetNetworkSmDevices200ResponseInner) GetOsName() string {
	if o == nil || o.OsName == nil {
		var ret string
		return ret
	}
	return *o.OsName
}

// GetOsNameOk returns a tuple with the OsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDevices200ResponseInner) GetOsNameOk() (*string, bool) {
	if o == nil || o.OsName == nil {
		return nil, false
	}
	return o.OsName, true
}

// HasOsName returns a boolean if a field has been set.
func (o *GetNetworkSmDevices200ResponseInner) HasOsName() bool {
	if o != nil && o.OsName != nil {
		return true
	}

	return false
}

// SetOsName gets a reference to the given string and assigns it to the OsName field.
func (o *GetNetworkSmDevices200ResponseInner) SetOsName(v string) {
	o.OsName = &v
}

// GetSystemModel returns the SystemModel field value if set, zero value otherwise.
func (o *GetNetworkSmDevices200ResponseInner) GetSystemModel() string {
	if o == nil || o.SystemModel == nil {
		var ret string
		return ret
	}
	return *o.SystemModel
}

// GetSystemModelOk returns a tuple with the SystemModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDevices200ResponseInner) GetSystemModelOk() (*string, bool) {
	if o == nil || o.SystemModel == nil {
		return nil, false
	}
	return o.SystemModel, true
}

// HasSystemModel returns a boolean if a field has been set.
func (o *GetNetworkSmDevices200ResponseInner) HasSystemModel() bool {
	if o != nil && o.SystemModel != nil {
		return true
	}

	return false
}

// SetSystemModel gets a reference to the given string and assigns it to the SystemModel field.
func (o *GetNetworkSmDevices200ResponseInner) SetSystemModel(v string) {
	o.SystemModel = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *GetNetworkSmDevices200ResponseInner) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDevices200ResponseInner) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *GetNetworkSmDevices200ResponseInner) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *GetNetworkSmDevices200ResponseInner) SetUuid(v string) {
	o.Uuid = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *GetNetworkSmDevices200ResponseInner) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDevices200ResponseInner) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *GetNetworkSmDevices200ResponseInner) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *GetNetworkSmDevices200ResponseInner) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *GetNetworkSmDevices200ResponseInner) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDevices200ResponseInner) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *GetNetworkSmDevices200ResponseInner) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *GetNetworkSmDevices200ResponseInner) SetSerial(v string) {
	o.Serial = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *GetNetworkSmDevices200ResponseInner) GetIp() string {
	if o == nil || o.Ip == nil {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDevices200ResponseInner) GetIpOk() (*string, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *GetNetworkSmDevices200ResponseInner) HasIp() bool {
	if o != nil && o.Ip != nil {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *GetNetworkSmDevices200ResponseInner) SetIp(v string) {
	o.Ip = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *GetNetworkSmDevices200ResponseInner) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDevices200ResponseInner) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *GetNetworkSmDevices200ResponseInner) HasNotes() bool {
	if o != nil && o.Notes != nil {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *GetNetworkSmDevices200ResponseInner) SetNotes(v string) {
	o.Notes = &v
}

func (o GetNetworkSmDevices200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Ssid != nil {
		toSerialize["ssid"] = o.Ssid
	}
	if o.WifiMac != nil {
		toSerialize["wifiMac"] = o.WifiMac
	}
	if o.OsName != nil {
		toSerialize["osName"] = o.OsName
	}
	if o.SystemModel != nil {
		toSerialize["systemModel"] = o.SystemModel
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.SerialNumber != nil {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	if o.Serial != nil {
		toSerialize["serial"] = o.Serial
	}
	if o.Ip != nil {
		toSerialize["ip"] = o.Ip
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkSmDevices200ResponseInner struct {
	value *GetNetworkSmDevices200ResponseInner
	isSet bool
}

func (v NullableGetNetworkSmDevices200ResponseInner) Get() *GetNetworkSmDevices200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkSmDevices200ResponseInner) Set(val *GetNetworkSmDevices200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSmDevices200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSmDevices200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSmDevices200ResponseInner(val *GetNetworkSmDevices200ResponseInner) *NullableGetNetworkSmDevices200ResponseInner {
	return &NullableGetNetworkSmDevices200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkSmDevices200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSmDevices200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


