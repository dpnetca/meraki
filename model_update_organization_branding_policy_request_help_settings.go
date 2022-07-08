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

// UpdateOrganizationBrandingPolicyRequestHelpSettings     Settings for describing the modifications to various Help page features. Each property in this object accepts one of     'default or inherit' (do not modify functionality), 'hide' (remove the section from Dashboard), or 'show' (always show     the section on Dashboard). Some properties in this object also accept custom HTML used to replace the section on     Dashboard; see the documentation for each property to see the allowed values. 
type UpdateOrganizationBrandingPolicyRequestHelpSettings struct {
	//     The Help tab, under which all support information resides. If this tab is hidden, no other 'Help' branding     customizations will be visible. Can be one of 'default or inherit', 'hide' or 'show'. 
	HelpTab *string `json:"helpTab,omitempty"`
	//     The 'Help -> Get Help' subtab on which Cisco Meraki KB, Product Manuals, and Support/Case Information are displayed. Note     that if this subtab is hidden, branding customizations for the KB on 'Get help', Cisco Meraki product documentation,     and support contact info will not be visible. Can be one of 'default or inherit', 'hide' or 'show'. 
	GetHelpSubtab *string `json:"getHelpSubtab,omitempty"`
	//     The 'Help -> Community' subtab which provides a link to Meraki Community. Can be one of 'default or inherit', 'hide' or 'show'. 
	CommunitySubtab *string `json:"communitySubtab,omitempty"`
	//     The 'Help -> Cases' Dashboard subtab on which Cisco Meraki support cases for this organization can be managed. Can be one     of 'default or inherit', 'hide' or 'show'. 
	CasesSubtab *string `json:"casesSubtab,omitempty"`
	//     The 'Help -> Data protection requests' Dashboard subtab on which requests to delete, restrict, or export end-user data can     be audited. Can be one of 'default or inherit', 'hide' or 'show'. 
	DataProtectionRequestsSubtab *string `json:"dataProtectionRequestsSubtab,omitempty"`
	//     The KB search box which appears on the Help page. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string. 
	GetHelpSubtabKnowledgeBaseSearch *string `json:"getHelpSubtabKnowledgeBaseSearch,omitempty"`
	//     The universal search box always visible on Dashboard will, by default, present results from the Meraki KB. This configures     whether these Meraki KB results should be returned. Can be one of 'default or inherit', 'hide' or 'show'. 
	UniversalSearchKnowledgeBaseSearch *string `json:"universalSearchKnowledgeBaseSearch,omitempty"`
	//     The 'Product Manuals' section of the 'Help -> Get Help' subtab. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string. 
	CiscoMerakiProductDocumentation *string `json:"ciscoMerakiProductDocumentation,omitempty"`
	//     The 'Contact Meraki Support' section of the 'Help -> Get Help' subtab. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string. 
	SupportContactInfo *string `json:"supportContactInfo,omitempty"`
	//     The 'Help -> New features' subtab where new Dashboard features are detailed. Can be one of 'default or inherit', 'hide' or 'show'. 
	NewFeaturesSubtab *string `json:"newFeaturesSubtab,omitempty"`
	//     The 'Help -> Firewall info' subtab where necessary upstream firewall rules for communication to the Cisco Meraki cloud are     listed. Can be one of 'default or inherit', 'hide' or 'show'. 
	FirewallInfoSubtab *string `json:"firewallInfoSubtab,omitempty"`
	//     The 'Help -> API docs' subtab where a detailed description of the Dashboard API is listed. Can be one of     'default or inherit', 'hide' or 'show'. 
	ApiDocsSubtab *string `json:"apiDocsSubtab,omitempty"`
	//     The 'Help -> Replacement info' subtab where important information regarding device replacements is detailed. Can be one of     'default or inherit', 'hide' or 'show'. 
	HardwareReplacementsSubtab *string `json:"hardwareReplacementsSubtab,omitempty"`
	//     The 'SM Forums' subtab which links to community-based support for Cisco Meraki Systems Manager. Only configurable for     organizations that contain Systems Manager networks. Can be one of 'default or inherit', 'hide' or 'show'. 
	SmForums *string `json:"smForums,omitempty"`
}

// NewUpdateOrganizationBrandingPolicyRequestHelpSettings instantiates a new UpdateOrganizationBrandingPolicyRequestHelpSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationBrandingPolicyRequestHelpSettings() *UpdateOrganizationBrandingPolicyRequestHelpSettings {
	this := UpdateOrganizationBrandingPolicyRequestHelpSettings{}
	return &this
}

// NewUpdateOrganizationBrandingPolicyRequestHelpSettingsWithDefaults instantiates a new UpdateOrganizationBrandingPolicyRequestHelpSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationBrandingPolicyRequestHelpSettingsWithDefaults() *UpdateOrganizationBrandingPolicyRequestHelpSettings {
	this := UpdateOrganizationBrandingPolicyRequestHelpSettings{}
	return &this
}

// GetHelpTab returns the HelpTab field value if set, zero value otherwise.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) GetHelpTab() string {
	if o == nil || o.HelpTab == nil {
		var ret string
		return ret
	}
	return *o.HelpTab
}

// GetHelpTabOk returns a tuple with the HelpTab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) GetHelpTabOk() (*string, bool) {
	if o == nil || o.HelpTab == nil {
		return nil, false
	}
	return o.HelpTab, true
}

// HasHelpTab returns a boolean if a field has been set.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) HasHelpTab() bool {
	if o != nil && o.HelpTab != nil {
		return true
	}

	return false
}

// SetHelpTab gets a reference to the given string and assigns it to the HelpTab field.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) SetHelpTab(v string) {
	o.HelpTab = &v
}

// GetGetHelpSubtab returns the GetHelpSubtab field value if set, zero value otherwise.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) GetGetHelpSubtab() string {
	if o == nil || o.GetHelpSubtab == nil {
		var ret string
		return ret
	}
	return *o.GetHelpSubtab
}

// GetGetHelpSubtabOk returns a tuple with the GetHelpSubtab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) GetGetHelpSubtabOk() (*string, bool) {
	if o == nil || o.GetHelpSubtab == nil {
		return nil, false
	}
	return o.GetHelpSubtab, true
}

// HasGetHelpSubtab returns a boolean if a field has been set.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) HasGetHelpSubtab() bool {
	if o != nil && o.GetHelpSubtab != nil {
		return true
	}

	return false
}

// SetGetHelpSubtab gets a reference to the given string and assigns it to the GetHelpSubtab field.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) SetGetHelpSubtab(v string) {
	o.GetHelpSubtab = &v
}

// GetCommunitySubtab returns the CommunitySubtab field value if set, zero value otherwise.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) GetCommunitySubtab() string {
	if o == nil || o.CommunitySubtab == nil {
		var ret string
		return ret
	}
	return *o.CommunitySubtab
}

// GetCommunitySubtabOk returns a tuple with the CommunitySubtab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) GetCommunitySubtabOk() (*string, bool) {
	if o == nil || o.CommunitySubtab == nil {
		return nil, false
	}
	return o.CommunitySubtab, true
}

// HasCommunitySubtab returns a boolean if a field has been set.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) HasCommunitySubtab() bool {
	if o != nil && o.CommunitySubtab != nil {
		return true
	}

	return false
}

// SetCommunitySubtab gets a reference to the given string and assigns it to the CommunitySubtab field.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) SetCommunitySubtab(v string) {
	o.CommunitySubtab = &v
}

// GetCasesSubtab returns the CasesSubtab field value if set, zero value otherwise.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) GetCasesSubtab() string {
	if o == nil || o.CasesSubtab == nil {
		var ret string
		return ret
	}
	return *o.CasesSubtab
}

// GetCasesSubtabOk returns a tuple with the CasesSubtab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) GetCasesSubtabOk() (*string, bool) {
	if o == nil || o.CasesSubtab == nil {
		return nil, false
	}
	return o.CasesSubtab, true
}

// HasCasesSubtab returns a boolean if a field has been set.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) HasCasesSubtab() bool {
	if o != nil && o.CasesSubtab != nil {
		return true
	}

	return false
}

// SetCasesSubtab gets a reference to the given string and assigns it to the CasesSubtab field.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) SetCasesSubtab(v string) {
	o.CasesSubtab = &v
}

// GetDataProtectionRequestsSubtab returns the DataProtectionRequestsSubtab field value if set, zero value otherwise.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) GetDataProtectionRequestsSubtab() string {
	if o == nil || o.DataProtectionRequestsSubtab == nil {
		var ret string
		return ret
	}
	return *o.DataProtectionRequestsSubtab
}

// GetDataProtectionRequestsSubtabOk returns a tuple with the DataProtectionRequestsSubtab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) GetDataProtectionRequestsSubtabOk() (*string, bool) {
	if o == nil || o.DataProtectionRequestsSubtab == nil {
		return nil, false
	}
	return o.DataProtectionRequestsSubtab, true
}

// HasDataProtectionRequestsSubtab returns a boolean if a field has been set.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) HasDataProtectionRequestsSubtab() bool {
	if o != nil && o.DataProtectionRequestsSubtab != nil {
		return true
	}

	return false
}

// SetDataProtectionRequestsSubtab gets a reference to the given string and assigns it to the DataProtectionRequestsSubtab field.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) SetDataProtectionRequestsSubtab(v string) {
	o.DataProtectionRequestsSubtab = &v
}

// GetGetHelpSubtabKnowledgeBaseSearch returns the GetHelpSubtabKnowledgeBaseSearch field value if set, zero value otherwise.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) GetGetHelpSubtabKnowledgeBaseSearch() string {
	if o == nil || o.GetHelpSubtabKnowledgeBaseSearch == nil {
		var ret string
		return ret
	}
	return *o.GetHelpSubtabKnowledgeBaseSearch
}

// GetGetHelpSubtabKnowledgeBaseSearchOk returns a tuple with the GetHelpSubtabKnowledgeBaseSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) GetGetHelpSubtabKnowledgeBaseSearchOk() (*string, bool) {
	if o == nil || o.GetHelpSubtabKnowledgeBaseSearch == nil {
		return nil, false
	}
	return o.GetHelpSubtabKnowledgeBaseSearch, true
}

// HasGetHelpSubtabKnowledgeBaseSearch returns a boolean if a field has been set.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) HasGetHelpSubtabKnowledgeBaseSearch() bool {
	if o != nil && o.GetHelpSubtabKnowledgeBaseSearch != nil {
		return true
	}

	return false
}

// SetGetHelpSubtabKnowledgeBaseSearch gets a reference to the given string and assigns it to the GetHelpSubtabKnowledgeBaseSearch field.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) SetGetHelpSubtabKnowledgeBaseSearch(v string) {
	o.GetHelpSubtabKnowledgeBaseSearch = &v
}

// GetUniversalSearchKnowledgeBaseSearch returns the UniversalSearchKnowledgeBaseSearch field value if set, zero value otherwise.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) GetUniversalSearchKnowledgeBaseSearch() string {
	if o == nil || o.UniversalSearchKnowledgeBaseSearch == nil {
		var ret string
		return ret
	}
	return *o.UniversalSearchKnowledgeBaseSearch
}

// GetUniversalSearchKnowledgeBaseSearchOk returns a tuple with the UniversalSearchKnowledgeBaseSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) GetUniversalSearchKnowledgeBaseSearchOk() (*string, bool) {
	if o == nil || o.UniversalSearchKnowledgeBaseSearch == nil {
		return nil, false
	}
	return o.UniversalSearchKnowledgeBaseSearch, true
}

// HasUniversalSearchKnowledgeBaseSearch returns a boolean if a field has been set.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) HasUniversalSearchKnowledgeBaseSearch() bool {
	if o != nil && o.UniversalSearchKnowledgeBaseSearch != nil {
		return true
	}

	return false
}

// SetUniversalSearchKnowledgeBaseSearch gets a reference to the given string and assigns it to the UniversalSearchKnowledgeBaseSearch field.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) SetUniversalSearchKnowledgeBaseSearch(v string) {
	o.UniversalSearchKnowledgeBaseSearch = &v
}

// GetCiscoMerakiProductDocumentation returns the CiscoMerakiProductDocumentation field value if set, zero value otherwise.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) GetCiscoMerakiProductDocumentation() string {
	if o == nil || o.CiscoMerakiProductDocumentation == nil {
		var ret string
		return ret
	}
	return *o.CiscoMerakiProductDocumentation
}

// GetCiscoMerakiProductDocumentationOk returns a tuple with the CiscoMerakiProductDocumentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) GetCiscoMerakiProductDocumentationOk() (*string, bool) {
	if o == nil || o.CiscoMerakiProductDocumentation == nil {
		return nil, false
	}
	return o.CiscoMerakiProductDocumentation, true
}

// HasCiscoMerakiProductDocumentation returns a boolean if a field has been set.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) HasCiscoMerakiProductDocumentation() bool {
	if o != nil && o.CiscoMerakiProductDocumentation != nil {
		return true
	}

	return false
}

// SetCiscoMerakiProductDocumentation gets a reference to the given string and assigns it to the CiscoMerakiProductDocumentation field.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) SetCiscoMerakiProductDocumentation(v string) {
	o.CiscoMerakiProductDocumentation = &v
}

// GetSupportContactInfo returns the SupportContactInfo field value if set, zero value otherwise.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) GetSupportContactInfo() string {
	if o == nil || o.SupportContactInfo == nil {
		var ret string
		return ret
	}
	return *o.SupportContactInfo
}

// GetSupportContactInfoOk returns a tuple with the SupportContactInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) GetSupportContactInfoOk() (*string, bool) {
	if o == nil || o.SupportContactInfo == nil {
		return nil, false
	}
	return o.SupportContactInfo, true
}

// HasSupportContactInfo returns a boolean if a field has been set.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) HasSupportContactInfo() bool {
	if o != nil && o.SupportContactInfo != nil {
		return true
	}

	return false
}

// SetSupportContactInfo gets a reference to the given string and assigns it to the SupportContactInfo field.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) SetSupportContactInfo(v string) {
	o.SupportContactInfo = &v
}

// GetNewFeaturesSubtab returns the NewFeaturesSubtab field value if set, zero value otherwise.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) GetNewFeaturesSubtab() string {
	if o == nil || o.NewFeaturesSubtab == nil {
		var ret string
		return ret
	}
	return *o.NewFeaturesSubtab
}

// GetNewFeaturesSubtabOk returns a tuple with the NewFeaturesSubtab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) GetNewFeaturesSubtabOk() (*string, bool) {
	if o == nil || o.NewFeaturesSubtab == nil {
		return nil, false
	}
	return o.NewFeaturesSubtab, true
}

// HasNewFeaturesSubtab returns a boolean if a field has been set.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) HasNewFeaturesSubtab() bool {
	if o != nil && o.NewFeaturesSubtab != nil {
		return true
	}

	return false
}

// SetNewFeaturesSubtab gets a reference to the given string and assigns it to the NewFeaturesSubtab field.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) SetNewFeaturesSubtab(v string) {
	o.NewFeaturesSubtab = &v
}

// GetFirewallInfoSubtab returns the FirewallInfoSubtab field value if set, zero value otherwise.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) GetFirewallInfoSubtab() string {
	if o == nil || o.FirewallInfoSubtab == nil {
		var ret string
		return ret
	}
	return *o.FirewallInfoSubtab
}

// GetFirewallInfoSubtabOk returns a tuple with the FirewallInfoSubtab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) GetFirewallInfoSubtabOk() (*string, bool) {
	if o == nil || o.FirewallInfoSubtab == nil {
		return nil, false
	}
	return o.FirewallInfoSubtab, true
}

// HasFirewallInfoSubtab returns a boolean if a field has been set.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) HasFirewallInfoSubtab() bool {
	if o != nil && o.FirewallInfoSubtab != nil {
		return true
	}

	return false
}

// SetFirewallInfoSubtab gets a reference to the given string and assigns it to the FirewallInfoSubtab field.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) SetFirewallInfoSubtab(v string) {
	o.FirewallInfoSubtab = &v
}

// GetApiDocsSubtab returns the ApiDocsSubtab field value if set, zero value otherwise.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) GetApiDocsSubtab() string {
	if o == nil || o.ApiDocsSubtab == nil {
		var ret string
		return ret
	}
	return *o.ApiDocsSubtab
}

// GetApiDocsSubtabOk returns a tuple with the ApiDocsSubtab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) GetApiDocsSubtabOk() (*string, bool) {
	if o == nil || o.ApiDocsSubtab == nil {
		return nil, false
	}
	return o.ApiDocsSubtab, true
}

// HasApiDocsSubtab returns a boolean if a field has been set.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) HasApiDocsSubtab() bool {
	if o != nil && o.ApiDocsSubtab != nil {
		return true
	}

	return false
}

// SetApiDocsSubtab gets a reference to the given string and assigns it to the ApiDocsSubtab field.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) SetApiDocsSubtab(v string) {
	o.ApiDocsSubtab = &v
}

// GetHardwareReplacementsSubtab returns the HardwareReplacementsSubtab field value if set, zero value otherwise.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) GetHardwareReplacementsSubtab() string {
	if o == nil || o.HardwareReplacementsSubtab == nil {
		var ret string
		return ret
	}
	return *o.HardwareReplacementsSubtab
}

// GetHardwareReplacementsSubtabOk returns a tuple with the HardwareReplacementsSubtab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) GetHardwareReplacementsSubtabOk() (*string, bool) {
	if o == nil || o.HardwareReplacementsSubtab == nil {
		return nil, false
	}
	return o.HardwareReplacementsSubtab, true
}

// HasHardwareReplacementsSubtab returns a boolean if a field has been set.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) HasHardwareReplacementsSubtab() bool {
	if o != nil && o.HardwareReplacementsSubtab != nil {
		return true
	}

	return false
}

// SetHardwareReplacementsSubtab gets a reference to the given string and assigns it to the HardwareReplacementsSubtab field.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) SetHardwareReplacementsSubtab(v string) {
	o.HardwareReplacementsSubtab = &v
}

// GetSmForums returns the SmForums field value if set, zero value otherwise.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) GetSmForums() string {
	if o == nil || o.SmForums == nil {
		var ret string
		return ret
	}
	return *o.SmForums
}

// GetSmForumsOk returns a tuple with the SmForums field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) GetSmForumsOk() (*string, bool) {
	if o == nil || o.SmForums == nil {
		return nil, false
	}
	return o.SmForums, true
}

// HasSmForums returns a boolean if a field has been set.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) HasSmForums() bool {
	if o != nil && o.SmForums != nil {
		return true
	}

	return false
}

// SetSmForums gets a reference to the given string and assigns it to the SmForums field.
func (o *UpdateOrganizationBrandingPolicyRequestHelpSettings) SetSmForums(v string) {
	o.SmForums = &v
}

func (o UpdateOrganizationBrandingPolicyRequestHelpSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HelpTab != nil {
		toSerialize["helpTab"] = o.HelpTab
	}
	if o.GetHelpSubtab != nil {
		toSerialize["getHelpSubtab"] = o.GetHelpSubtab
	}
	if o.CommunitySubtab != nil {
		toSerialize["communitySubtab"] = o.CommunitySubtab
	}
	if o.CasesSubtab != nil {
		toSerialize["casesSubtab"] = o.CasesSubtab
	}
	if o.DataProtectionRequestsSubtab != nil {
		toSerialize["dataProtectionRequestsSubtab"] = o.DataProtectionRequestsSubtab
	}
	if o.GetHelpSubtabKnowledgeBaseSearch != nil {
		toSerialize["getHelpSubtabKnowledgeBaseSearch"] = o.GetHelpSubtabKnowledgeBaseSearch
	}
	if o.UniversalSearchKnowledgeBaseSearch != nil {
		toSerialize["universalSearchKnowledgeBaseSearch"] = o.UniversalSearchKnowledgeBaseSearch
	}
	if o.CiscoMerakiProductDocumentation != nil {
		toSerialize["ciscoMerakiProductDocumentation"] = o.CiscoMerakiProductDocumentation
	}
	if o.SupportContactInfo != nil {
		toSerialize["supportContactInfo"] = o.SupportContactInfo
	}
	if o.NewFeaturesSubtab != nil {
		toSerialize["newFeaturesSubtab"] = o.NewFeaturesSubtab
	}
	if o.FirewallInfoSubtab != nil {
		toSerialize["firewallInfoSubtab"] = o.FirewallInfoSubtab
	}
	if o.ApiDocsSubtab != nil {
		toSerialize["apiDocsSubtab"] = o.ApiDocsSubtab
	}
	if o.HardwareReplacementsSubtab != nil {
		toSerialize["hardwareReplacementsSubtab"] = o.HardwareReplacementsSubtab
	}
	if o.SmForums != nil {
		toSerialize["smForums"] = o.SmForums
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateOrganizationBrandingPolicyRequestHelpSettings struct {
	value *UpdateOrganizationBrandingPolicyRequestHelpSettings
	isSet bool
}

func (v NullableUpdateOrganizationBrandingPolicyRequestHelpSettings) Get() *UpdateOrganizationBrandingPolicyRequestHelpSettings {
	return v.value
}

func (v *NullableUpdateOrganizationBrandingPolicyRequestHelpSettings) Set(val *UpdateOrganizationBrandingPolicyRequestHelpSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationBrandingPolicyRequestHelpSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationBrandingPolicyRequestHelpSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationBrandingPolicyRequestHelpSettings(val *UpdateOrganizationBrandingPolicyRequestHelpSettings) *NullableUpdateOrganizationBrandingPolicyRequestHelpSettings {
	return &NullableUpdateOrganizationBrandingPolicyRequestHelpSettings{value: val, isSet: true}
}

func (v NullableUpdateOrganizationBrandingPolicyRequestHelpSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationBrandingPolicyRequestHelpSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


