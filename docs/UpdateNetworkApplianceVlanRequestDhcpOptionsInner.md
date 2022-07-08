# UpdateNetworkApplianceVlanRequestDhcpOptionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | The code for the DHCP option. This should be an integer between 2 and 254. | 
**Type** | **string** | The type for the DHCP option. One of: &#39;text&#39;, &#39;ip&#39;, &#39;hex&#39; or &#39;integer&#39; | 
**Value** | **string** | The value for the DHCP option | 

## Methods

### NewUpdateNetworkApplianceVlanRequestDhcpOptionsInner

`func NewUpdateNetworkApplianceVlanRequestDhcpOptionsInner(code string, type_ string, value string, ) *UpdateNetworkApplianceVlanRequestDhcpOptionsInner`

NewUpdateNetworkApplianceVlanRequestDhcpOptionsInner instantiates a new UpdateNetworkApplianceVlanRequestDhcpOptionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkApplianceVlanRequestDhcpOptionsInnerWithDefaults

`func NewUpdateNetworkApplianceVlanRequestDhcpOptionsInnerWithDefaults() *UpdateNetworkApplianceVlanRequestDhcpOptionsInner`

NewUpdateNetworkApplianceVlanRequestDhcpOptionsInnerWithDefaults instantiates a new UpdateNetworkApplianceVlanRequestDhcpOptionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *UpdateNetworkApplianceVlanRequestDhcpOptionsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateNetworkApplianceVlanRequestDhcpOptionsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateNetworkApplianceVlanRequestDhcpOptionsInner) SetCode(v string)`

SetCode sets Code field to given value.


### GetType

`func (o *UpdateNetworkApplianceVlanRequestDhcpOptionsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateNetworkApplianceVlanRequestDhcpOptionsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateNetworkApplianceVlanRequestDhcpOptionsInner) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *UpdateNetworkApplianceVlanRequestDhcpOptionsInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateNetworkApplianceVlanRequestDhcpOptionsInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateNetworkApplianceVlanRequestDhcpOptionsInner) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


