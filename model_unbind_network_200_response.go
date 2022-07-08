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

// UnbindNetwork200Response struct for UnbindNetwork200Response
type UnbindNetwork200Response struct {
	// Network ID
	Id *string `json:"id,omitempty"`
	// Organization ID
	OrganizationId *string `json:"organizationId,omitempty"`
	// Network name
	Name *string `json:"name,omitempty"`
	// List of the product types that the network supports
	ProductTypes []string `json:"productTypes,omitempty"`
	// Timezone of the network
	TimeZone *string `json:"timeZone,omitempty"`
	// Network tags
	Tags []string `json:"tags,omitempty"`
	// Enrollment string for the network
	EnrollmentString *string `json:"enrollmentString,omitempty"`
	// URL to the network Dashboard UI
	Url *string `json:"url,omitempty"`
	// Notes for the network
	Notes *string `json:"notes,omitempty"`
	// If the network is bound to a config template
	IsBoundToConfigTemplate *bool `json:"isBoundToConfigTemplate,omitempty"`
}

// NewUnbindNetwork200Response instantiates a new UnbindNetwork200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnbindNetwork200Response() *UnbindNetwork200Response {
	this := UnbindNetwork200Response{}
	return &this
}

// NewUnbindNetwork200ResponseWithDefaults instantiates a new UnbindNetwork200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnbindNetwork200ResponseWithDefaults() *UnbindNetwork200Response {
	this := UnbindNetwork200Response{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnbindNetwork200Response) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnbindNetwork200Response) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnbindNetwork200Response) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UnbindNetwork200Response) SetId(v string) {
	o.Id = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *UnbindNetwork200Response) GetOrganizationId() string {
	if o == nil || o.OrganizationId == nil {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnbindNetwork200Response) GetOrganizationIdOk() (*string, bool) {
	if o == nil || o.OrganizationId == nil {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *UnbindNetwork200Response) HasOrganizationId() bool {
	if o != nil && o.OrganizationId != nil {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *UnbindNetwork200Response) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UnbindNetwork200Response) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnbindNetwork200Response) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UnbindNetwork200Response) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UnbindNetwork200Response) SetName(v string) {
	o.Name = &v
}

// GetProductTypes returns the ProductTypes field value if set, zero value otherwise.
func (o *UnbindNetwork200Response) GetProductTypes() []string {
	if o == nil || o.ProductTypes == nil {
		var ret []string
		return ret
	}
	return o.ProductTypes
}

// GetProductTypesOk returns a tuple with the ProductTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnbindNetwork200Response) GetProductTypesOk() ([]string, bool) {
	if o == nil || o.ProductTypes == nil {
		return nil, false
	}
	return o.ProductTypes, true
}

// HasProductTypes returns a boolean if a field has been set.
func (o *UnbindNetwork200Response) HasProductTypes() bool {
	if o != nil && o.ProductTypes != nil {
		return true
	}

	return false
}

// SetProductTypes gets a reference to the given []string and assigns it to the ProductTypes field.
func (o *UnbindNetwork200Response) SetProductTypes(v []string) {
	o.ProductTypes = v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *UnbindNetwork200Response) GetTimeZone() string {
	if o == nil || o.TimeZone == nil {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnbindNetwork200Response) GetTimeZoneOk() (*string, bool) {
	if o == nil || o.TimeZone == nil {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *UnbindNetwork200Response) HasTimeZone() bool {
	if o != nil && o.TimeZone != nil {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *UnbindNetwork200Response) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UnbindNetwork200Response) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnbindNetwork200Response) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UnbindNetwork200Response) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *UnbindNetwork200Response) SetTags(v []string) {
	o.Tags = v
}

// GetEnrollmentString returns the EnrollmentString field value if set, zero value otherwise.
func (o *UnbindNetwork200Response) GetEnrollmentString() string {
	if o == nil || o.EnrollmentString == nil {
		var ret string
		return ret
	}
	return *o.EnrollmentString
}

// GetEnrollmentStringOk returns a tuple with the EnrollmentString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnbindNetwork200Response) GetEnrollmentStringOk() (*string, bool) {
	if o == nil || o.EnrollmentString == nil {
		return nil, false
	}
	return o.EnrollmentString, true
}

// HasEnrollmentString returns a boolean if a field has been set.
func (o *UnbindNetwork200Response) HasEnrollmentString() bool {
	if o != nil && o.EnrollmentString != nil {
		return true
	}

	return false
}

// SetEnrollmentString gets a reference to the given string and assigns it to the EnrollmentString field.
func (o *UnbindNetwork200Response) SetEnrollmentString(v string) {
	o.EnrollmentString = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *UnbindNetwork200Response) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnbindNetwork200Response) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *UnbindNetwork200Response) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *UnbindNetwork200Response) SetUrl(v string) {
	o.Url = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *UnbindNetwork200Response) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnbindNetwork200Response) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *UnbindNetwork200Response) HasNotes() bool {
	if o != nil && o.Notes != nil {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *UnbindNetwork200Response) SetNotes(v string) {
	o.Notes = &v
}

// GetIsBoundToConfigTemplate returns the IsBoundToConfigTemplate field value if set, zero value otherwise.
func (o *UnbindNetwork200Response) GetIsBoundToConfigTemplate() bool {
	if o == nil || o.IsBoundToConfigTemplate == nil {
		var ret bool
		return ret
	}
	return *o.IsBoundToConfigTemplate
}

// GetIsBoundToConfigTemplateOk returns a tuple with the IsBoundToConfigTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnbindNetwork200Response) GetIsBoundToConfigTemplateOk() (*bool, bool) {
	if o == nil || o.IsBoundToConfigTemplate == nil {
		return nil, false
	}
	return o.IsBoundToConfigTemplate, true
}

// HasIsBoundToConfigTemplate returns a boolean if a field has been set.
func (o *UnbindNetwork200Response) HasIsBoundToConfigTemplate() bool {
	if o != nil && o.IsBoundToConfigTemplate != nil {
		return true
	}

	return false
}

// SetIsBoundToConfigTemplate gets a reference to the given bool and assigns it to the IsBoundToConfigTemplate field.
func (o *UnbindNetwork200Response) SetIsBoundToConfigTemplate(v bool) {
	o.IsBoundToConfigTemplate = &v
}

func (o UnbindNetwork200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.OrganizationId != nil {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ProductTypes != nil {
		toSerialize["productTypes"] = o.ProductTypes
	}
	if o.TimeZone != nil {
		toSerialize["timeZone"] = o.TimeZone
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.EnrollmentString != nil {
		toSerialize["enrollmentString"] = o.EnrollmentString
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	if o.IsBoundToConfigTemplate != nil {
		toSerialize["isBoundToConfigTemplate"] = o.IsBoundToConfigTemplate
	}
	return json.Marshal(toSerialize)
}

type NullableUnbindNetwork200Response struct {
	value *UnbindNetwork200Response
	isSet bool
}

func (v NullableUnbindNetwork200Response) Get() *UnbindNetwork200Response {
	return v.value
}

func (v *NullableUnbindNetwork200Response) Set(val *UnbindNetwork200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUnbindNetwork200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUnbindNetwork200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnbindNetwork200Response(val *UnbindNetwork200Response) *NullableUnbindNetwork200Response {
	return &NullableUnbindNetwork200Response{value: val, isSet: true}
}

func (v NullableUnbindNetwork200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnbindNetwork200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

