# AuthenticationChallengePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The UUID for the authentication challenge policy. If not specified during creation, an ID will be automatically be assigned. When updating an existing authentication challenge policy, this field is ignored and the ID is determined from the URL path parameter. | [optional] 
**Name** | **string** | (sortable) The name of this authentication challenge policy. The number of characters in the name is limited to 64. | 
**Description** | Pointer to **string** | (sortable) A description of the authentication challenge policy. The number of characters in the description is limited to 1000. | [optional] 
**System** | Pointer to **bool** | (sortable) This field is read-only and indicates this authentication challenge policy cannot be modified. | [optional] 
**ChallengeResponseChain** | [**[]ChallengeResponseMapping**](ChallengeResponseMapping.md) | An array of challenge response mappings, ordered by the precedence of each mapping. When no challengeResponseChain is needed for the policy, this field must be set to an empty array. | 
**DefaultChallengeResponse** | [**ChallengeResponse**](ChallengeResponse.md) |  | 

## Methods

### NewAuthenticationChallengePolicy

`func NewAuthenticationChallengePolicy(name string, challengeResponseChain []ChallengeResponseMapping, defaultChallengeResponse ChallengeResponse, ) *AuthenticationChallengePolicy`

NewAuthenticationChallengePolicy instantiates a new AuthenticationChallengePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationChallengePolicyWithDefaults

`func NewAuthenticationChallengePolicyWithDefaults() *AuthenticationChallengePolicy`

NewAuthenticationChallengePolicyWithDefaults instantiates a new AuthenticationChallengePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthenticationChallengePolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthenticationChallengePolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthenticationChallengePolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthenticationChallengePolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AuthenticationChallengePolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthenticationChallengePolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthenticationChallengePolicy) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AuthenticationChallengePolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthenticationChallengePolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthenticationChallengePolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthenticationChallengePolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSystem

`func (o *AuthenticationChallengePolicy) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *AuthenticationChallengePolicy) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *AuthenticationChallengePolicy) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *AuthenticationChallengePolicy) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetChallengeResponseChain

`func (o *AuthenticationChallengePolicy) GetChallengeResponseChain() []ChallengeResponseMapping`

GetChallengeResponseChain returns the ChallengeResponseChain field if non-nil, zero value otherwise.

### GetChallengeResponseChainOk

`func (o *AuthenticationChallengePolicy) GetChallengeResponseChainOk() (*[]ChallengeResponseMapping, bool)`

GetChallengeResponseChainOk returns a tuple with the ChallengeResponseChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeResponseChain

`func (o *AuthenticationChallengePolicy) SetChallengeResponseChain(v []ChallengeResponseMapping)`

SetChallengeResponseChain sets ChallengeResponseChain field to given value.


### GetDefaultChallengeResponse

`func (o *AuthenticationChallengePolicy) GetDefaultChallengeResponse() ChallengeResponse`

GetDefaultChallengeResponse returns the DefaultChallengeResponse field if non-nil, zero value otherwise.

### GetDefaultChallengeResponseOk

`func (o *AuthenticationChallengePolicy) GetDefaultChallengeResponseOk() (*ChallengeResponse, bool)`

GetDefaultChallengeResponseOk returns a tuple with the DefaultChallengeResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultChallengeResponse

`func (o *AuthenticationChallengePolicy) SetDefaultChallengeResponse(v ChallengeResponse)`

SetDefaultChallengeResponse sets DefaultChallengeResponse field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


