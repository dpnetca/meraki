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

// checks if the GetNetworkWirelessSsidSplashSettings200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkWirelessSsidSplashSettings200Response{}

// GetNetworkWirelessSsidSplashSettings200Response struct for GetNetworkWirelessSsidSplashSettings200Response
type GetNetworkWirelessSsidSplashSettings200Response struct {
	// SSID number
	SsidNumber *int32 `json:"ssidNumber,omitempty"`
	// The type of splash page for this SSID
	SplashPage *string `json:"splashPage,omitempty"`
	// Boolean indicating whether the users will be redirected to the custom splash url
	UseSplashUrl *bool `json:"useSplashUrl,omitempty"`
	// The custom splash URL of the click-through splash page.
	SplashUrl *string `json:"splashUrl,omitempty"`
	// Splash timeout in minutes.
	SplashTimeout *int32 `json:"splashTimeout,omitempty"`
	// The custom redirect URL where the users will go after the splash page.
	RedirectUrl *string `json:"redirectUrl,omitempty"`
	// The Boolean indicating whether the the user will be redirected to the custom redirect URL after the splash page.
	UseRedirectUrl *bool `json:"useRedirectUrl,omitempty"`
	// The welcome message for the users on the splash page.
	WelcomeMessage *string `json:"welcomeMessage,omitempty"`
	SplashLogo *GetNetworkWirelessSsidSplashSettings200ResponseSplashLogo `json:"splashLogo,omitempty"`
	SplashImage *GetNetworkWirelessSsidSplashSettings200ResponseSplashImage `json:"splashImage,omitempty"`
	SplashPrepaidFront *GetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront `json:"splashPrepaidFront,omitempty"`
	GuestSponsorship *GetNetworkWirelessSsidSplashSettings200ResponseGuestSponsorship `json:"guestSponsorship,omitempty"`
	// How restricted allowing traffic should be. If true, all traffic types are blocked until the splash page is acknowledged. If false, all non-HTTP traffic is allowed before the splash page is acknowledged.
	BlockAllTrafficBeforeSignOn *bool `json:"blockAllTrafficBeforeSignOn,omitempty"`
	// How login attempts should be handled when the controller is unreachable.
	ControllerDisconnectionBehavior *string `json:"controllerDisconnectionBehavior,omitempty"`
	// Whether or not to allow simultaneous logins from different devices.
	AllowSimultaneousLogins *bool `json:"allowSimultaneousLogins,omitempty"`
	Billing *GetNetworkWirelessSsidSplashSettings200ResponseBilling `json:"billing,omitempty"`
	SentryEnrollment *GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment `json:"sentryEnrollment,omitempty"`
	SelfRegistration *GetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration `json:"selfRegistration,omitempty"`
}

// NewGetNetworkWirelessSsidSplashSettings200Response instantiates a new GetNetworkWirelessSsidSplashSettings200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWirelessSsidSplashSettings200Response() *GetNetworkWirelessSsidSplashSettings200Response {
	this := GetNetworkWirelessSsidSplashSettings200Response{}
	return &this
}

// NewGetNetworkWirelessSsidSplashSettings200ResponseWithDefaults instantiates a new GetNetworkWirelessSsidSplashSettings200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWirelessSsidSplashSettings200ResponseWithDefaults() *GetNetworkWirelessSsidSplashSettings200Response {
	this := GetNetworkWirelessSsidSplashSettings200Response{}
	return &this
}

// GetSsidNumber returns the SsidNumber field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidSplashSettings200Response) GetSsidNumber() int32 {
	if o == nil || IsNil(o.SsidNumber) {
		var ret int32
		return ret
	}
	return *o.SsidNumber
}

// GetSsidNumberOk returns a tuple with the SsidNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidSplashSettings200Response) GetSsidNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.SsidNumber) {
		return nil, false
	}
	return o.SsidNumber, true
}

// HasSsidNumber returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidSplashSettings200Response) HasSsidNumber() bool {
	if o != nil && !IsNil(o.SsidNumber) {
		return true
	}

	return false
}

// SetSsidNumber gets a reference to the given int32 and assigns it to the SsidNumber field.
func (o *GetNetworkWirelessSsidSplashSettings200Response) SetSsidNumber(v int32) {
	o.SsidNumber = &v
}

// GetSplashPage returns the SplashPage field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidSplashSettings200Response) GetSplashPage() string {
	if o == nil || IsNil(o.SplashPage) {
		var ret string
		return ret
	}
	return *o.SplashPage
}

// GetSplashPageOk returns a tuple with the SplashPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidSplashSettings200Response) GetSplashPageOk() (*string, bool) {
	if o == nil || IsNil(o.SplashPage) {
		return nil, false
	}
	return o.SplashPage, true
}

// HasSplashPage returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidSplashSettings200Response) HasSplashPage() bool {
	if o != nil && !IsNil(o.SplashPage) {
		return true
	}

	return false
}

// SetSplashPage gets a reference to the given string and assigns it to the SplashPage field.
func (o *GetNetworkWirelessSsidSplashSettings200Response) SetSplashPage(v string) {
	o.SplashPage = &v
}

// GetUseSplashUrl returns the UseSplashUrl field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidSplashSettings200Response) GetUseSplashUrl() bool {
	if o == nil || IsNil(o.UseSplashUrl) {
		var ret bool
		return ret
	}
	return *o.UseSplashUrl
}

// GetUseSplashUrlOk returns a tuple with the UseSplashUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidSplashSettings200Response) GetUseSplashUrlOk() (*bool, bool) {
	if o == nil || IsNil(o.UseSplashUrl) {
		return nil, false
	}
	return o.UseSplashUrl, true
}

// HasUseSplashUrl returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidSplashSettings200Response) HasUseSplashUrl() bool {
	if o != nil && !IsNil(o.UseSplashUrl) {
		return true
	}

	return false
}

// SetUseSplashUrl gets a reference to the given bool and assigns it to the UseSplashUrl field.
func (o *GetNetworkWirelessSsidSplashSettings200Response) SetUseSplashUrl(v bool) {
	o.UseSplashUrl = &v
}

// GetSplashUrl returns the SplashUrl field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidSplashSettings200Response) GetSplashUrl() string {
	if o == nil || IsNil(o.SplashUrl) {
		var ret string
		return ret
	}
	return *o.SplashUrl
}

// GetSplashUrlOk returns a tuple with the SplashUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidSplashSettings200Response) GetSplashUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SplashUrl) {
		return nil, false
	}
	return o.SplashUrl, true
}

// HasSplashUrl returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidSplashSettings200Response) HasSplashUrl() bool {
	if o != nil && !IsNil(o.SplashUrl) {
		return true
	}

	return false
}

// SetSplashUrl gets a reference to the given string and assigns it to the SplashUrl field.
func (o *GetNetworkWirelessSsidSplashSettings200Response) SetSplashUrl(v string) {
	o.SplashUrl = &v
}

// GetSplashTimeout returns the SplashTimeout field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidSplashSettings200Response) GetSplashTimeout() int32 {
	if o == nil || IsNil(o.SplashTimeout) {
		var ret int32
		return ret
	}
	return *o.SplashTimeout
}

// GetSplashTimeoutOk returns a tuple with the SplashTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidSplashSettings200Response) GetSplashTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.SplashTimeout) {
		return nil, false
	}
	return o.SplashTimeout, true
}

// HasSplashTimeout returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidSplashSettings200Response) HasSplashTimeout() bool {
	if o != nil && !IsNil(o.SplashTimeout) {
		return true
	}

	return false
}

// SetSplashTimeout gets a reference to the given int32 and assigns it to the SplashTimeout field.
func (o *GetNetworkWirelessSsidSplashSettings200Response) SetSplashTimeout(v int32) {
	o.SplashTimeout = &v
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidSplashSettings200Response) GetRedirectUrl() string {
	if o == nil || IsNil(o.RedirectUrl) {
		var ret string
		return ret
	}
	return *o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidSplashSettings200Response) GetRedirectUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RedirectUrl) {
		return nil, false
	}
	return o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidSplashSettings200Response) HasRedirectUrl() bool {
	if o != nil && !IsNil(o.RedirectUrl) {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given string and assigns it to the RedirectUrl field.
func (o *GetNetworkWirelessSsidSplashSettings200Response) SetRedirectUrl(v string) {
	o.RedirectUrl = &v
}

// GetUseRedirectUrl returns the UseRedirectUrl field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidSplashSettings200Response) GetUseRedirectUrl() bool {
	if o == nil || IsNil(o.UseRedirectUrl) {
		var ret bool
		return ret
	}
	return *o.UseRedirectUrl
}

// GetUseRedirectUrlOk returns a tuple with the UseRedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidSplashSettings200Response) GetUseRedirectUrlOk() (*bool, bool) {
	if o == nil || IsNil(o.UseRedirectUrl) {
		return nil, false
	}
	return o.UseRedirectUrl, true
}

// HasUseRedirectUrl returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidSplashSettings200Response) HasUseRedirectUrl() bool {
	if o != nil && !IsNil(o.UseRedirectUrl) {
		return true
	}

	return false
}

// SetUseRedirectUrl gets a reference to the given bool and assigns it to the UseRedirectUrl field.
func (o *GetNetworkWirelessSsidSplashSettings200Response) SetUseRedirectUrl(v bool) {
	o.UseRedirectUrl = &v
}

// GetWelcomeMessage returns the WelcomeMessage field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidSplashSettings200Response) GetWelcomeMessage() string {
	if o == nil || IsNil(o.WelcomeMessage) {
		var ret string
		return ret
	}
	return *o.WelcomeMessage
}

// GetWelcomeMessageOk returns a tuple with the WelcomeMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidSplashSettings200Response) GetWelcomeMessageOk() (*string, bool) {
	if o == nil || IsNil(o.WelcomeMessage) {
		return nil, false
	}
	return o.WelcomeMessage, true
}

// HasWelcomeMessage returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidSplashSettings200Response) HasWelcomeMessage() bool {
	if o != nil && !IsNil(o.WelcomeMessage) {
		return true
	}

	return false
}

// SetWelcomeMessage gets a reference to the given string and assigns it to the WelcomeMessage field.
func (o *GetNetworkWirelessSsidSplashSettings200Response) SetWelcomeMessage(v string) {
	o.WelcomeMessage = &v
}

// GetSplashLogo returns the SplashLogo field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidSplashSettings200Response) GetSplashLogo() GetNetworkWirelessSsidSplashSettings200ResponseSplashLogo {
	if o == nil || IsNil(o.SplashLogo) {
		var ret GetNetworkWirelessSsidSplashSettings200ResponseSplashLogo
		return ret
	}
	return *o.SplashLogo
}

// GetSplashLogoOk returns a tuple with the SplashLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidSplashSettings200Response) GetSplashLogoOk() (*GetNetworkWirelessSsidSplashSettings200ResponseSplashLogo, bool) {
	if o == nil || IsNil(o.SplashLogo) {
		return nil, false
	}
	return o.SplashLogo, true
}

// HasSplashLogo returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidSplashSettings200Response) HasSplashLogo() bool {
	if o != nil && !IsNil(o.SplashLogo) {
		return true
	}

	return false
}

// SetSplashLogo gets a reference to the given GetNetworkWirelessSsidSplashSettings200ResponseSplashLogo and assigns it to the SplashLogo field.
func (o *GetNetworkWirelessSsidSplashSettings200Response) SetSplashLogo(v GetNetworkWirelessSsidSplashSettings200ResponseSplashLogo) {
	o.SplashLogo = &v
}

// GetSplashImage returns the SplashImage field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidSplashSettings200Response) GetSplashImage() GetNetworkWirelessSsidSplashSettings200ResponseSplashImage {
	if o == nil || IsNil(o.SplashImage) {
		var ret GetNetworkWirelessSsidSplashSettings200ResponseSplashImage
		return ret
	}
	return *o.SplashImage
}

// GetSplashImageOk returns a tuple with the SplashImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidSplashSettings200Response) GetSplashImageOk() (*GetNetworkWirelessSsidSplashSettings200ResponseSplashImage, bool) {
	if o == nil || IsNil(o.SplashImage) {
		return nil, false
	}
	return o.SplashImage, true
}

// HasSplashImage returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidSplashSettings200Response) HasSplashImage() bool {
	if o != nil && !IsNil(o.SplashImage) {
		return true
	}

	return false
}

// SetSplashImage gets a reference to the given GetNetworkWirelessSsidSplashSettings200ResponseSplashImage and assigns it to the SplashImage field.
func (o *GetNetworkWirelessSsidSplashSettings200Response) SetSplashImage(v GetNetworkWirelessSsidSplashSettings200ResponseSplashImage) {
	o.SplashImage = &v
}

// GetSplashPrepaidFront returns the SplashPrepaidFront field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidSplashSettings200Response) GetSplashPrepaidFront() GetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront {
	if o == nil || IsNil(o.SplashPrepaidFront) {
		var ret GetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront
		return ret
	}
	return *o.SplashPrepaidFront
}

// GetSplashPrepaidFrontOk returns a tuple with the SplashPrepaidFront field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidSplashSettings200Response) GetSplashPrepaidFrontOk() (*GetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront, bool) {
	if o == nil || IsNil(o.SplashPrepaidFront) {
		return nil, false
	}
	return o.SplashPrepaidFront, true
}

// HasSplashPrepaidFront returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidSplashSettings200Response) HasSplashPrepaidFront() bool {
	if o != nil && !IsNil(o.SplashPrepaidFront) {
		return true
	}

	return false
}

// SetSplashPrepaidFront gets a reference to the given GetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront and assigns it to the SplashPrepaidFront field.
func (o *GetNetworkWirelessSsidSplashSettings200Response) SetSplashPrepaidFront(v GetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront) {
	o.SplashPrepaidFront = &v
}

// GetGuestSponsorship returns the GuestSponsorship field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidSplashSettings200Response) GetGuestSponsorship() GetNetworkWirelessSsidSplashSettings200ResponseGuestSponsorship {
	if o == nil || IsNil(o.GuestSponsorship) {
		var ret GetNetworkWirelessSsidSplashSettings200ResponseGuestSponsorship
		return ret
	}
	return *o.GuestSponsorship
}

// GetGuestSponsorshipOk returns a tuple with the GuestSponsorship field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidSplashSettings200Response) GetGuestSponsorshipOk() (*GetNetworkWirelessSsidSplashSettings200ResponseGuestSponsorship, bool) {
	if o == nil || IsNil(o.GuestSponsorship) {
		return nil, false
	}
	return o.GuestSponsorship, true
}

// HasGuestSponsorship returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidSplashSettings200Response) HasGuestSponsorship() bool {
	if o != nil && !IsNil(o.GuestSponsorship) {
		return true
	}

	return false
}

// SetGuestSponsorship gets a reference to the given GetNetworkWirelessSsidSplashSettings200ResponseGuestSponsorship and assigns it to the GuestSponsorship field.
func (o *GetNetworkWirelessSsidSplashSettings200Response) SetGuestSponsorship(v GetNetworkWirelessSsidSplashSettings200ResponseGuestSponsorship) {
	o.GuestSponsorship = &v
}

// GetBlockAllTrafficBeforeSignOn returns the BlockAllTrafficBeforeSignOn field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidSplashSettings200Response) GetBlockAllTrafficBeforeSignOn() bool {
	if o == nil || IsNil(o.BlockAllTrafficBeforeSignOn) {
		var ret bool
		return ret
	}
	return *o.BlockAllTrafficBeforeSignOn
}

// GetBlockAllTrafficBeforeSignOnOk returns a tuple with the BlockAllTrafficBeforeSignOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidSplashSettings200Response) GetBlockAllTrafficBeforeSignOnOk() (*bool, bool) {
	if o == nil || IsNil(o.BlockAllTrafficBeforeSignOn) {
		return nil, false
	}
	return o.BlockAllTrafficBeforeSignOn, true
}

// HasBlockAllTrafficBeforeSignOn returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidSplashSettings200Response) HasBlockAllTrafficBeforeSignOn() bool {
	if o != nil && !IsNil(o.BlockAllTrafficBeforeSignOn) {
		return true
	}

	return false
}

// SetBlockAllTrafficBeforeSignOn gets a reference to the given bool and assigns it to the BlockAllTrafficBeforeSignOn field.
func (o *GetNetworkWirelessSsidSplashSettings200Response) SetBlockAllTrafficBeforeSignOn(v bool) {
	o.BlockAllTrafficBeforeSignOn = &v
}

// GetControllerDisconnectionBehavior returns the ControllerDisconnectionBehavior field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidSplashSettings200Response) GetControllerDisconnectionBehavior() string {
	if o == nil || IsNil(o.ControllerDisconnectionBehavior) {
		var ret string
		return ret
	}
	return *o.ControllerDisconnectionBehavior
}

// GetControllerDisconnectionBehaviorOk returns a tuple with the ControllerDisconnectionBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidSplashSettings200Response) GetControllerDisconnectionBehaviorOk() (*string, bool) {
	if o == nil || IsNil(o.ControllerDisconnectionBehavior) {
		return nil, false
	}
	return o.ControllerDisconnectionBehavior, true
}

// HasControllerDisconnectionBehavior returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidSplashSettings200Response) HasControllerDisconnectionBehavior() bool {
	if o != nil && !IsNil(o.ControllerDisconnectionBehavior) {
		return true
	}

	return false
}

// SetControllerDisconnectionBehavior gets a reference to the given string and assigns it to the ControllerDisconnectionBehavior field.
func (o *GetNetworkWirelessSsidSplashSettings200Response) SetControllerDisconnectionBehavior(v string) {
	o.ControllerDisconnectionBehavior = &v
}

// GetAllowSimultaneousLogins returns the AllowSimultaneousLogins field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidSplashSettings200Response) GetAllowSimultaneousLogins() bool {
	if o == nil || IsNil(o.AllowSimultaneousLogins) {
		var ret bool
		return ret
	}
	return *o.AllowSimultaneousLogins
}

// GetAllowSimultaneousLoginsOk returns a tuple with the AllowSimultaneousLogins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidSplashSettings200Response) GetAllowSimultaneousLoginsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowSimultaneousLogins) {
		return nil, false
	}
	return o.AllowSimultaneousLogins, true
}

// HasAllowSimultaneousLogins returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidSplashSettings200Response) HasAllowSimultaneousLogins() bool {
	if o != nil && !IsNil(o.AllowSimultaneousLogins) {
		return true
	}

	return false
}

// SetAllowSimultaneousLogins gets a reference to the given bool and assigns it to the AllowSimultaneousLogins field.
func (o *GetNetworkWirelessSsidSplashSettings200Response) SetAllowSimultaneousLogins(v bool) {
	o.AllowSimultaneousLogins = &v
}

// GetBilling returns the Billing field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidSplashSettings200Response) GetBilling() GetNetworkWirelessSsidSplashSettings200ResponseBilling {
	if o == nil || IsNil(o.Billing) {
		var ret GetNetworkWirelessSsidSplashSettings200ResponseBilling
		return ret
	}
	return *o.Billing
}

// GetBillingOk returns a tuple with the Billing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidSplashSettings200Response) GetBillingOk() (*GetNetworkWirelessSsidSplashSettings200ResponseBilling, bool) {
	if o == nil || IsNil(o.Billing) {
		return nil, false
	}
	return o.Billing, true
}

// HasBilling returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidSplashSettings200Response) HasBilling() bool {
	if o != nil && !IsNil(o.Billing) {
		return true
	}

	return false
}

// SetBilling gets a reference to the given GetNetworkWirelessSsidSplashSettings200ResponseBilling and assigns it to the Billing field.
func (o *GetNetworkWirelessSsidSplashSettings200Response) SetBilling(v GetNetworkWirelessSsidSplashSettings200ResponseBilling) {
	o.Billing = &v
}

// GetSentryEnrollment returns the SentryEnrollment field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidSplashSettings200Response) GetSentryEnrollment() GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment {
	if o == nil || IsNil(o.SentryEnrollment) {
		var ret GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment
		return ret
	}
	return *o.SentryEnrollment
}

// GetSentryEnrollmentOk returns a tuple with the SentryEnrollment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidSplashSettings200Response) GetSentryEnrollmentOk() (*GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment, bool) {
	if o == nil || IsNil(o.SentryEnrollment) {
		return nil, false
	}
	return o.SentryEnrollment, true
}

// HasSentryEnrollment returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidSplashSettings200Response) HasSentryEnrollment() bool {
	if o != nil && !IsNil(o.SentryEnrollment) {
		return true
	}

	return false
}

// SetSentryEnrollment gets a reference to the given GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment and assigns it to the SentryEnrollment field.
func (o *GetNetworkWirelessSsidSplashSettings200Response) SetSentryEnrollment(v GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment) {
	o.SentryEnrollment = &v
}

// GetSelfRegistration returns the SelfRegistration field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidSplashSettings200Response) GetSelfRegistration() GetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration {
	if o == nil || IsNil(o.SelfRegistration) {
		var ret GetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration
		return ret
	}
	return *o.SelfRegistration
}

// GetSelfRegistrationOk returns a tuple with the SelfRegistration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidSplashSettings200Response) GetSelfRegistrationOk() (*GetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration, bool) {
	if o == nil || IsNil(o.SelfRegistration) {
		return nil, false
	}
	return o.SelfRegistration, true
}

// HasSelfRegistration returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidSplashSettings200Response) HasSelfRegistration() bool {
	if o != nil && !IsNil(o.SelfRegistration) {
		return true
	}

	return false
}

// SetSelfRegistration gets a reference to the given GetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration and assigns it to the SelfRegistration field.
func (o *GetNetworkWirelessSsidSplashSettings200Response) SetSelfRegistration(v GetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration) {
	o.SelfRegistration = &v
}

func (o GetNetworkWirelessSsidSplashSettings200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkWirelessSsidSplashSettings200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SsidNumber) {
		toSerialize["ssidNumber"] = o.SsidNumber
	}
	if !IsNil(o.SplashPage) {
		toSerialize["splashPage"] = o.SplashPage
	}
	if !IsNil(o.UseSplashUrl) {
		toSerialize["useSplashUrl"] = o.UseSplashUrl
	}
	if !IsNil(o.SplashUrl) {
		toSerialize["splashUrl"] = o.SplashUrl
	}
	if !IsNil(o.SplashTimeout) {
		toSerialize["splashTimeout"] = o.SplashTimeout
	}
	if !IsNil(o.RedirectUrl) {
		toSerialize["redirectUrl"] = o.RedirectUrl
	}
	if !IsNil(o.UseRedirectUrl) {
		toSerialize["useRedirectUrl"] = o.UseRedirectUrl
	}
	if !IsNil(o.WelcomeMessage) {
		toSerialize["welcomeMessage"] = o.WelcomeMessage
	}
	if !IsNil(o.SplashLogo) {
		toSerialize["splashLogo"] = o.SplashLogo
	}
	if !IsNil(o.SplashImage) {
		toSerialize["splashImage"] = o.SplashImage
	}
	if !IsNil(o.SplashPrepaidFront) {
		toSerialize["splashPrepaidFront"] = o.SplashPrepaidFront
	}
	if !IsNil(o.GuestSponsorship) {
		toSerialize["guestSponsorship"] = o.GuestSponsorship
	}
	if !IsNil(o.BlockAllTrafficBeforeSignOn) {
		toSerialize["blockAllTrafficBeforeSignOn"] = o.BlockAllTrafficBeforeSignOn
	}
	if !IsNil(o.ControllerDisconnectionBehavior) {
		toSerialize["controllerDisconnectionBehavior"] = o.ControllerDisconnectionBehavior
	}
	if !IsNil(o.AllowSimultaneousLogins) {
		toSerialize["allowSimultaneousLogins"] = o.AllowSimultaneousLogins
	}
	if !IsNil(o.Billing) {
		toSerialize["billing"] = o.Billing
	}
	if !IsNil(o.SentryEnrollment) {
		toSerialize["sentryEnrollment"] = o.SentryEnrollment
	}
	if !IsNil(o.SelfRegistration) {
		toSerialize["selfRegistration"] = o.SelfRegistration
	}
	return toSerialize, nil
}

type NullableGetNetworkWirelessSsidSplashSettings200Response struct {
	value *GetNetworkWirelessSsidSplashSettings200Response
	isSet bool
}

func (v NullableGetNetworkWirelessSsidSplashSettings200Response) Get() *GetNetworkWirelessSsidSplashSettings200Response {
	return v.value
}

func (v *NullableGetNetworkWirelessSsidSplashSettings200Response) Set(val *GetNetworkWirelessSsidSplashSettings200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWirelessSsidSplashSettings200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWirelessSsidSplashSettings200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWirelessSsidSplashSettings200Response(val *GetNetworkWirelessSsidSplashSettings200Response) *NullableGetNetworkWirelessSsidSplashSettings200Response {
	return &NullableGetNetworkWirelessSsidSplashSettings200Response{value: val, isSet: true}
}

func (v NullableGetNetworkWirelessSsidSplashSettings200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWirelessSsidSplashSettings200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


