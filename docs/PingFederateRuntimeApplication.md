# PingFederateRuntimeApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextRoot** | Pointer to **string** | The base path of the PingFederate runtime application. Default value is &#39;/&#39;. | [optional] 
**PrimaryVirtualHostId** | **int64** | The ID of the primary virtual host to use for front channel requests to the PA proxied PingFederate runtime application. This virtual host will be used for the default OpenID Connect Issuer when an application specific issuer is not configured. | 
**AdditionalVirtualHostIds** | Pointer to **[]int64** | Additional virtual host IDs that can be used to proxy the PingFederate runtime application. | [optional] 
**Policy** | Pointer to [**[]PolicyItem**](PolicyItem.md) | A List of PolicyItems associated with the PingFederate runtime application. | [optional] 
**ClientCertHeaderNames** | Pointer to **[]string** | The header names to contain PEM-encoded client certificates. The list position correlates to the index in the client certificate chain. For example, the first element always maps to the leaf certificate. | [optional] 
**CaseSensitive** | Pointer to **bool** | True if the context root is case sensitive. | [optional] 

## Methods

### NewPingFederateRuntimeApplication

`func NewPingFederateRuntimeApplication(primaryVirtualHostId int64, ) *PingFederateRuntimeApplication`

NewPingFederateRuntimeApplication instantiates a new PingFederateRuntimeApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPingFederateRuntimeApplicationWithDefaults

`func NewPingFederateRuntimeApplicationWithDefaults() *PingFederateRuntimeApplication`

NewPingFederateRuntimeApplicationWithDefaults instantiates a new PingFederateRuntimeApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextRoot

`func (o *PingFederateRuntimeApplication) GetContextRoot() string`

GetContextRoot returns the ContextRoot field if non-nil, zero value otherwise.

### GetContextRootOk

`func (o *PingFederateRuntimeApplication) GetContextRootOk() (*string, bool)`

GetContextRootOk returns a tuple with the ContextRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextRoot

`func (o *PingFederateRuntimeApplication) SetContextRoot(v string)`

SetContextRoot sets ContextRoot field to given value.

### HasContextRoot

`func (o *PingFederateRuntimeApplication) HasContextRoot() bool`

HasContextRoot returns a boolean if a field has been set.

### GetPrimaryVirtualHostId

`func (o *PingFederateRuntimeApplication) GetPrimaryVirtualHostId() int64`

GetPrimaryVirtualHostId returns the PrimaryVirtualHostId field if non-nil, zero value otherwise.

### GetPrimaryVirtualHostIdOk

`func (o *PingFederateRuntimeApplication) GetPrimaryVirtualHostIdOk() (*int64, bool)`

GetPrimaryVirtualHostIdOk returns a tuple with the PrimaryVirtualHostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryVirtualHostId

`func (o *PingFederateRuntimeApplication) SetPrimaryVirtualHostId(v int64)`

SetPrimaryVirtualHostId sets PrimaryVirtualHostId field to given value.


### GetAdditionalVirtualHostIds

`func (o *PingFederateRuntimeApplication) GetAdditionalVirtualHostIds() []int64`

GetAdditionalVirtualHostIds returns the AdditionalVirtualHostIds field if non-nil, zero value otherwise.

### GetAdditionalVirtualHostIdsOk

`func (o *PingFederateRuntimeApplication) GetAdditionalVirtualHostIdsOk() (*[]int64, bool)`

GetAdditionalVirtualHostIdsOk returns a tuple with the AdditionalVirtualHostIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalVirtualHostIds

`func (o *PingFederateRuntimeApplication) SetAdditionalVirtualHostIds(v []int64)`

SetAdditionalVirtualHostIds sets AdditionalVirtualHostIds field to given value.

### HasAdditionalVirtualHostIds

`func (o *PingFederateRuntimeApplication) HasAdditionalVirtualHostIds() bool`

HasAdditionalVirtualHostIds returns a boolean if a field has been set.

### GetPolicy

`func (o *PingFederateRuntimeApplication) GetPolicy() []PolicyItem`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *PingFederateRuntimeApplication) GetPolicyOk() (*[]PolicyItem, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *PingFederateRuntimeApplication) SetPolicy(v []PolicyItem)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *PingFederateRuntimeApplication) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetClientCertHeaderNames

`func (o *PingFederateRuntimeApplication) GetClientCertHeaderNames() []string`

GetClientCertHeaderNames returns the ClientCertHeaderNames field if non-nil, zero value otherwise.

### GetClientCertHeaderNamesOk

`func (o *PingFederateRuntimeApplication) GetClientCertHeaderNamesOk() (*[]string, bool)`

GetClientCertHeaderNamesOk returns a tuple with the ClientCertHeaderNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertHeaderNames

`func (o *PingFederateRuntimeApplication) SetClientCertHeaderNames(v []string)`

SetClientCertHeaderNames sets ClientCertHeaderNames field to given value.

### HasClientCertHeaderNames

`func (o *PingFederateRuntimeApplication) HasClientCertHeaderNames() bool`

HasClientCertHeaderNames returns a boolean if a field has been set.

### GetCaseSensitive

`func (o *PingFederateRuntimeApplication) GetCaseSensitive() bool`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *PingFederateRuntimeApplication) GetCaseSensitiveOk() (*bool, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *PingFederateRuntimeApplication) SetCaseSensitive(v bool)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *PingFederateRuntimeApplication) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


