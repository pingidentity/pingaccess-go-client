# SupportedScopes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scopes** | Pointer to **[]string** | The list of supported scopes (excluding &#39;openid&#39;). | [optional] 
**ClientId** | Pointer to **string** | The ID of the client that the scopes are associated with. If not specified, the list of scopes represents all scopes supported by the provider. | [optional] 

## Methods

### NewSupportedScopes

`func NewSupportedScopes() *SupportedScopes`

NewSupportedScopes instantiates a new SupportedScopes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportedScopesWithDefaults

`func NewSupportedScopesWithDefaults() *SupportedScopes`

NewSupportedScopesWithDefaults instantiates a new SupportedScopes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScopes

`func (o *SupportedScopes) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *SupportedScopes) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *SupportedScopes) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *SupportedScopes) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetClientId

`func (o *SupportedScopes) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *SupportedScopes) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *SupportedScopes) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *SupportedScopes) HasClientId() bool`

HasClientId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


