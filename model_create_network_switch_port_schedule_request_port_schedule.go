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

// checks if the CreateNetworkSwitchPortScheduleRequestPortSchedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkSwitchPortScheduleRequestPortSchedule{}

// CreateNetworkSwitchPortScheduleRequestPortSchedule     The schedule for switch port scheduling. Schedules are applied to days of the week.     When it's empty, default schedule with all days of a week are configured.     Any unspecified day in the schedule is added as a default schedule configuration of the day. 
type CreateNetworkSwitchPortScheduleRequestPortSchedule struct {
	Monday *CreateNetworkGroupPolicyRequestSchedulingMonday `json:"monday,omitempty"`
	Tuesday *CreateNetworkGroupPolicyRequestSchedulingTuesday `json:"tuesday,omitempty"`
	Wednesday *CreateNetworkGroupPolicyRequestSchedulingWednesday `json:"wednesday,omitempty"`
	Thursday *CreateNetworkGroupPolicyRequestSchedulingThursday `json:"thursday,omitempty"`
	Friday *CreateNetworkGroupPolicyRequestSchedulingFriday `json:"friday,omitempty"`
	Saturday *CreateNetworkGroupPolicyRequestSchedulingSaturday `json:"saturday,omitempty"`
	Sunday *CreateNetworkGroupPolicyRequestSchedulingSunday `json:"sunday,omitempty"`
}

// NewCreateNetworkSwitchPortScheduleRequestPortSchedule instantiates a new CreateNetworkSwitchPortScheduleRequestPortSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkSwitchPortScheduleRequestPortSchedule() *CreateNetworkSwitchPortScheduleRequestPortSchedule {
	this := CreateNetworkSwitchPortScheduleRequestPortSchedule{}
	return &this
}

// NewCreateNetworkSwitchPortScheduleRequestPortScheduleWithDefaults instantiates a new CreateNetworkSwitchPortScheduleRequestPortSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkSwitchPortScheduleRequestPortScheduleWithDefaults() *CreateNetworkSwitchPortScheduleRequestPortSchedule {
	this := CreateNetworkSwitchPortScheduleRequestPortSchedule{}
	return &this
}

// GetMonday returns the Monday field value if set, zero value otherwise.
func (o *CreateNetworkSwitchPortScheduleRequestPortSchedule) GetMonday() CreateNetworkGroupPolicyRequestSchedulingMonday {
	if o == nil || IsNil(o.Monday) {
		var ret CreateNetworkGroupPolicyRequestSchedulingMonday
		return ret
	}
	return *o.Monday
}

// GetMondayOk returns a tuple with the Monday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchPortScheduleRequestPortSchedule) GetMondayOk() (*CreateNetworkGroupPolicyRequestSchedulingMonday, bool) {
	if o == nil || IsNil(o.Monday) {
		return nil, false
	}
	return o.Monday, true
}

// HasMonday returns a boolean if a field has been set.
func (o *CreateNetworkSwitchPortScheduleRequestPortSchedule) HasMonday() bool {
	if o != nil && !IsNil(o.Monday) {
		return true
	}

	return false
}

// SetMonday gets a reference to the given CreateNetworkGroupPolicyRequestSchedulingMonday and assigns it to the Monday field.
func (o *CreateNetworkSwitchPortScheduleRequestPortSchedule) SetMonday(v CreateNetworkGroupPolicyRequestSchedulingMonday) {
	o.Monday = &v
}

// GetTuesday returns the Tuesday field value if set, zero value otherwise.
func (o *CreateNetworkSwitchPortScheduleRequestPortSchedule) GetTuesday() CreateNetworkGroupPolicyRequestSchedulingTuesday {
	if o == nil || IsNil(o.Tuesday) {
		var ret CreateNetworkGroupPolicyRequestSchedulingTuesday
		return ret
	}
	return *o.Tuesday
}

// GetTuesdayOk returns a tuple with the Tuesday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchPortScheduleRequestPortSchedule) GetTuesdayOk() (*CreateNetworkGroupPolicyRequestSchedulingTuesday, bool) {
	if o == nil || IsNil(o.Tuesday) {
		return nil, false
	}
	return o.Tuesday, true
}

// HasTuesday returns a boolean if a field has been set.
func (o *CreateNetworkSwitchPortScheduleRequestPortSchedule) HasTuesday() bool {
	if o != nil && !IsNil(o.Tuesday) {
		return true
	}

	return false
}

// SetTuesday gets a reference to the given CreateNetworkGroupPolicyRequestSchedulingTuesday and assigns it to the Tuesday field.
func (o *CreateNetworkSwitchPortScheduleRequestPortSchedule) SetTuesday(v CreateNetworkGroupPolicyRequestSchedulingTuesday) {
	o.Tuesday = &v
}

// GetWednesday returns the Wednesday field value if set, zero value otherwise.
func (o *CreateNetworkSwitchPortScheduleRequestPortSchedule) GetWednesday() CreateNetworkGroupPolicyRequestSchedulingWednesday {
	if o == nil || IsNil(o.Wednesday) {
		var ret CreateNetworkGroupPolicyRequestSchedulingWednesday
		return ret
	}
	return *o.Wednesday
}

// GetWednesdayOk returns a tuple with the Wednesday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchPortScheduleRequestPortSchedule) GetWednesdayOk() (*CreateNetworkGroupPolicyRequestSchedulingWednesday, bool) {
	if o == nil || IsNil(o.Wednesday) {
		return nil, false
	}
	return o.Wednesday, true
}

// HasWednesday returns a boolean if a field has been set.
func (o *CreateNetworkSwitchPortScheduleRequestPortSchedule) HasWednesday() bool {
	if o != nil && !IsNil(o.Wednesday) {
		return true
	}

	return false
}

// SetWednesday gets a reference to the given CreateNetworkGroupPolicyRequestSchedulingWednesday and assigns it to the Wednesday field.
func (o *CreateNetworkSwitchPortScheduleRequestPortSchedule) SetWednesday(v CreateNetworkGroupPolicyRequestSchedulingWednesday) {
	o.Wednesday = &v
}

// GetThursday returns the Thursday field value if set, zero value otherwise.
func (o *CreateNetworkSwitchPortScheduleRequestPortSchedule) GetThursday() CreateNetworkGroupPolicyRequestSchedulingThursday {
	if o == nil || IsNil(o.Thursday) {
		var ret CreateNetworkGroupPolicyRequestSchedulingThursday
		return ret
	}
	return *o.Thursday
}

// GetThursdayOk returns a tuple with the Thursday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchPortScheduleRequestPortSchedule) GetThursdayOk() (*CreateNetworkGroupPolicyRequestSchedulingThursday, bool) {
	if o == nil || IsNil(o.Thursday) {
		return nil, false
	}
	return o.Thursday, true
}

// HasThursday returns a boolean if a field has been set.
func (o *CreateNetworkSwitchPortScheduleRequestPortSchedule) HasThursday() bool {
	if o != nil && !IsNil(o.Thursday) {
		return true
	}

	return false
}

// SetThursday gets a reference to the given CreateNetworkGroupPolicyRequestSchedulingThursday and assigns it to the Thursday field.
func (o *CreateNetworkSwitchPortScheduleRequestPortSchedule) SetThursday(v CreateNetworkGroupPolicyRequestSchedulingThursday) {
	o.Thursday = &v
}

// GetFriday returns the Friday field value if set, zero value otherwise.
func (o *CreateNetworkSwitchPortScheduleRequestPortSchedule) GetFriday() CreateNetworkGroupPolicyRequestSchedulingFriday {
	if o == nil || IsNil(o.Friday) {
		var ret CreateNetworkGroupPolicyRequestSchedulingFriday
		return ret
	}
	return *o.Friday
}

// GetFridayOk returns a tuple with the Friday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchPortScheduleRequestPortSchedule) GetFridayOk() (*CreateNetworkGroupPolicyRequestSchedulingFriday, bool) {
	if o == nil || IsNil(o.Friday) {
		return nil, false
	}
	return o.Friday, true
}

// HasFriday returns a boolean if a field has been set.
func (o *CreateNetworkSwitchPortScheduleRequestPortSchedule) HasFriday() bool {
	if o != nil && !IsNil(o.Friday) {
		return true
	}

	return false
}

// SetFriday gets a reference to the given CreateNetworkGroupPolicyRequestSchedulingFriday and assigns it to the Friday field.
func (o *CreateNetworkSwitchPortScheduleRequestPortSchedule) SetFriday(v CreateNetworkGroupPolicyRequestSchedulingFriday) {
	o.Friday = &v
}

// GetSaturday returns the Saturday field value if set, zero value otherwise.
func (o *CreateNetworkSwitchPortScheduleRequestPortSchedule) GetSaturday() CreateNetworkGroupPolicyRequestSchedulingSaturday {
	if o == nil || IsNil(o.Saturday) {
		var ret CreateNetworkGroupPolicyRequestSchedulingSaturday
		return ret
	}
	return *o.Saturday
}

// GetSaturdayOk returns a tuple with the Saturday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchPortScheduleRequestPortSchedule) GetSaturdayOk() (*CreateNetworkGroupPolicyRequestSchedulingSaturday, bool) {
	if o == nil || IsNil(o.Saturday) {
		return nil, false
	}
	return o.Saturday, true
}

// HasSaturday returns a boolean if a field has been set.
func (o *CreateNetworkSwitchPortScheduleRequestPortSchedule) HasSaturday() bool {
	if o != nil && !IsNil(o.Saturday) {
		return true
	}

	return false
}

// SetSaturday gets a reference to the given CreateNetworkGroupPolicyRequestSchedulingSaturday and assigns it to the Saturday field.
func (o *CreateNetworkSwitchPortScheduleRequestPortSchedule) SetSaturday(v CreateNetworkGroupPolicyRequestSchedulingSaturday) {
	o.Saturday = &v
}

// GetSunday returns the Sunday field value if set, zero value otherwise.
func (o *CreateNetworkSwitchPortScheduleRequestPortSchedule) GetSunday() CreateNetworkGroupPolicyRequestSchedulingSunday {
	if o == nil || IsNil(o.Sunday) {
		var ret CreateNetworkGroupPolicyRequestSchedulingSunday
		return ret
	}
	return *o.Sunday
}

// GetSundayOk returns a tuple with the Sunday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchPortScheduleRequestPortSchedule) GetSundayOk() (*CreateNetworkGroupPolicyRequestSchedulingSunday, bool) {
	if o == nil || IsNil(o.Sunday) {
		return nil, false
	}
	return o.Sunday, true
}

// HasSunday returns a boolean if a field has been set.
func (o *CreateNetworkSwitchPortScheduleRequestPortSchedule) HasSunday() bool {
	if o != nil && !IsNil(o.Sunday) {
		return true
	}

	return false
}

// SetSunday gets a reference to the given CreateNetworkGroupPolicyRequestSchedulingSunday and assigns it to the Sunday field.
func (o *CreateNetworkSwitchPortScheduleRequestPortSchedule) SetSunday(v CreateNetworkGroupPolicyRequestSchedulingSunday) {
	o.Sunday = &v
}

func (o CreateNetworkSwitchPortScheduleRequestPortSchedule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkSwitchPortScheduleRequestPortSchedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Monday) {
		toSerialize["monday"] = o.Monday
	}
	if !IsNil(o.Tuesday) {
		toSerialize["tuesday"] = o.Tuesday
	}
	if !IsNil(o.Wednesday) {
		toSerialize["wednesday"] = o.Wednesday
	}
	if !IsNil(o.Thursday) {
		toSerialize["thursday"] = o.Thursday
	}
	if !IsNil(o.Friday) {
		toSerialize["friday"] = o.Friday
	}
	if !IsNil(o.Saturday) {
		toSerialize["saturday"] = o.Saturday
	}
	if !IsNil(o.Sunday) {
		toSerialize["sunday"] = o.Sunday
	}
	return toSerialize, nil
}

type NullableCreateNetworkSwitchPortScheduleRequestPortSchedule struct {
	value *CreateNetworkSwitchPortScheduleRequestPortSchedule
	isSet bool
}

func (v NullableCreateNetworkSwitchPortScheduleRequestPortSchedule) Get() *CreateNetworkSwitchPortScheduleRequestPortSchedule {
	return v.value
}

func (v *NullableCreateNetworkSwitchPortScheduleRequestPortSchedule) Set(val *CreateNetworkSwitchPortScheduleRequestPortSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkSwitchPortScheduleRequestPortSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkSwitchPortScheduleRequestPortSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkSwitchPortScheduleRequestPortSchedule(val *CreateNetworkSwitchPortScheduleRequestPortSchedule) *NullableCreateNetworkSwitchPortScheduleRequestPortSchedule {
	return &NullableCreateNetworkSwitchPortScheduleRequestPortSchedule{value: val, isSet: true}
}

func (v NullableCreateNetworkSwitchPortScheduleRequestPortSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkSwitchPortScheduleRequestPortSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


