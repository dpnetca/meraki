# GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IkeCipherAlgo** | Pointer to **[]string** | This is the cipher algorithm to be used in Phase 1. The value should be an array with one of the following algorithms: &#39;aes256&#39;, &#39;aes192&#39;, &#39;aes128&#39;, &#39;tripledes&#39;, &#39;des&#39; | [optional] 
**IkeAuthAlgo** | Pointer to **[]string** | This is the authentication algorithm to be used in Phase 1. The value should be an array with one of the following algorithms: &#39;sha256&#39;, &#39;sha1&#39;, &#39;md5&#39; | [optional] 
**IkePrfAlgo** | Pointer to **[]string** | [optional] This is the pseudo-random function to be used in IKE_SA. The value should be an array with one of the following algorithms: &#39;prfsha256&#39;, &#39;prfsha1&#39;, &#39;prfmd5&#39;, &#39;default&#39;. The &#39;default&#39; option can be used to default to the Authentication algorithm. | [optional] 
**IkeDiffieHellmanGroup** | Pointer to **[]string** | This is the Diffie-Hellman group to be used in Phase 1. The value should be an array with one of the following algorithms: &#39;group14&#39;, &#39;group5&#39;, &#39;group2&#39;, &#39;group1&#39; | [optional] 
**IkeLifetime** | Pointer to **int32** | The lifetime of the Phase 1 SA in seconds. | [optional] 
**ChildCipherAlgo** | Pointer to **[]string** | This is the cipher algorithms to be used in Phase 2. The value should be an array with one or more of the following algorithms: &#39;aes256&#39;, &#39;aes192&#39;, &#39;aes128&#39;, &#39;tripledes&#39;, &#39;des&#39;, &#39;null&#39; | [optional] 
**ChildAuthAlgo** | Pointer to **[]string** | This is the authentication algorithms to be used in Phase 2. The value should be an array with one of the following algorithms: &#39;sha256&#39;, &#39;sha1&#39;, &#39;md5&#39; | [optional] 
**ChildPfsGroup** | Pointer to **[]string** | This is the Diffie-Hellman group to be used for Perfect Forward Secrecy in Phase 2. The value should be an array with one of the following values: &#39;disabled&#39;,&#39;group14&#39;, &#39;group5&#39;, &#39;group2&#39;, &#39;group1&#39; | [optional] 
**ChildLifetime** | Pointer to **int32** | The lifetime of the Phase 2 SA in seconds. | [optional] 

## Methods

### NewGetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies

`func NewGetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies() *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies`

NewGetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies instantiates a new GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPoliciesWithDefaults

`func NewGetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPoliciesWithDefaults() *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies`

NewGetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPoliciesWithDefaults instantiates a new GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIkeCipherAlgo

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies) GetIkeCipherAlgo() []string`

GetIkeCipherAlgo returns the IkeCipherAlgo field if non-nil, zero value otherwise.

### GetIkeCipherAlgoOk

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies) GetIkeCipherAlgoOk() (*[]string, bool)`

GetIkeCipherAlgoOk returns a tuple with the IkeCipherAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeCipherAlgo

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies) SetIkeCipherAlgo(v []string)`

SetIkeCipherAlgo sets IkeCipherAlgo field to given value.

### HasIkeCipherAlgo

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies) HasIkeCipherAlgo() bool`

HasIkeCipherAlgo returns a boolean if a field has been set.

### GetIkeAuthAlgo

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies) GetIkeAuthAlgo() []string`

GetIkeAuthAlgo returns the IkeAuthAlgo field if non-nil, zero value otherwise.

### GetIkeAuthAlgoOk

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies) GetIkeAuthAlgoOk() (*[]string, bool)`

GetIkeAuthAlgoOk returns a tuple with the IkeAuthAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeAuthAlgo

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies) SetIkeAuthAlgo(v []string)`

SetIkeAuthAlgo sets IkeAuthAlgo field to given value.

### HasIkeAuthAlgo

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies) HasIkeAuthAlgo() bool`

HasIkeAuthAlgo returns a boolean if a field has been set.

### GetIkePrfAlgo

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies) GetIkePrfAlgo() []string`

GetIkePrfAlgo returns the IkePrfAlgo field if non-nil, zero value otherwise.

### GetIkePrfAlgoOk

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies) GetIkePrfAlgoOk() (*[]string, bool)`

GetIkePrfAlgoOk returns a tuple with the IkePrfAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkePrfAlgo

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies) SetIkePrfAlgo(v []string)`

SetIkePrfAlgo sets IkePrfAlgo field to given value.

### HasIkePrfAlgo

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies) HasIkePrfAlgo() bool`

HasIkePrfAlgo returns a boolean if a field has been set.

### GetIkeDiffieHellmanGroup

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies) GetIkeDiffieHellmanGroup() []string`

GetIkeDiffieHellmanGroup returns the IkeDiffieHellmanGroup field if non-nil, zero value otherwise.

### GetIkeDiffieHellmanGroupOk

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies) GetIkeDiffieHellmanGroupOk() (*[]string, bool)`

GetIkeDiffieHellmanGroupOk returns a tuple with the IkeDiffieHellmanGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeDiffieHellmanGroup

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies) SetIkeDiffieHellmanGroup(v []string)`

SetIkeDiffieHellmanGroup sets IkeDiffieHellmanGroup field to given value.

### HasIkeDiffieHellmanGroup

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies) HasIkeDiffieHellmanGroup() bool`

HasIkeDiffieHellmanGroup returns a boolean if a field has been set.

### GetIkeLifetime

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies) GetIkeLifetime() int32`

GetIkeLifetime returns the IkeLifetime field if non-nil, zero value otherwise.

### GetIkeLifetimeOk

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies) GetIkeLifetimeOk() (*int32, bool)`

GetIkeLifetimeOk returns a tuple with the IkeLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeLifetime

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies) SetIkeLifetime(v int32)`

SetIkeLifetime sets IkeLifetime field to given value.

### HasIkeLifetime

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies) HasIkeLifetime() bool`

HasIkeLifetime returns a boolean if a field has been set.

### GetChildCipherAlgo

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies) GetChildCipherAlgo() []string`

GetChildCipherAlgo returns the ChildCipherAlgo field if non-nil, zero value otherwise.

### GetChildCipherAlgoOk

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies) GetChildCipherAlgoOk() (*[]string, bool)`

GetChildCipherAlgoOk returns a tuple with the ChildCipherAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildCipherAlgo

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies) SetChildCipherAlgo(v []string)`

SetChildCipherAlgo sets ChildCipherAlgo field to given value.

### HasChildCipherAlgo

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies) HasChildCipherAlgo() bool`

HasChildCipherAlgo returns a boolean if a field has been set.

### GetChildAuthAlgo

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies) GetChildAuthAlgo() []string`

GetChildAuthAlgo returns the ChildAuthAlgo field if non-nil, zero value otherwise.

### GetChildAuthAlgoOk

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies) GetChildAuthAlgoOk() (*[]string, bool)`

GetChildAuthAlgoOk returns a tuple with the ChildAuthAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildAuthAlgo

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies) SetChildAuthAlgo(v []string)`

SetChildAuthAlgo sets ChildAuthAlgo field to given value.

### HasChildAuthAlgo

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies) HasChildAuthAlgo() bool`

HasChildAuthAlgo returns a boolean if a field has been set.

### GetChildPfsGroup

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies) GetChildPfsGroup() []string`

GetChildPfsGroup returns the ChildPfsGroup field if non-nil, zero value otherwise.

### GetChildPfsGroupOk

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies) GetChildPfsGroupOk() (*[]string, bool)`

GetChildPfsGroupOk returns a tuple with the ChildPfsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildPfsGroup

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies) SetChildPfsGroup(v []string)`

SetChildPfsGroup sets ChildPfsGroup field to given value.

### HasChildPfsGroup

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies) HasChildPfsGroup() bool`

HasChildPfsGroup returns a boolean if a field has been set.

### GetChildLifetime

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies) GetChildLifetime() int32`

GetChildLifetime returns the ChildLifetime field if non-nil, zero value otherwise.

### GetChildLifetimeOk

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies) GetChildLifetimeOk() (*int32, bool)`

GetChildLifetimeOk returns a tuple with the ChildLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildLifetime

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies) SetChildLifetime(v int32)`

SetChildLifetime sets ChildLifetime field to given value.

### HasChildLifetime

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies) HasChildLifetime() bool`

HasChildLifetime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


