# AuthnReqList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | When creating a new AuthnReqList, this is the ID for the AuthnReqList. If not specified, an ID will be automatically assigned. When updating an existing AuthnReqList, this field is ignored and the ID is determined by the path parameter. | [optional] 
**Name** | **string** | (sortable) The name of the authentication requirements list. | 
**AuthnReqs** | **[]string** | The ordered list of authentication requirements, or identifiers, which define how PingFederate will authenticate a user during the OIDC login flow. | 

## Methods

### NewAuthnReqList

`func NewAuthnReqList(name string, authnReqs []string, ) *AuthnReqList`

NewAuthnReqList instantiates a new AuthnReqList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthnReqListWithDefaults

`func NewAuthnReqListWithDefaults() *AuthnReqList`

NewAuthnReqListWithDefaults instantiates a new AuthnReqList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthnReqList) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthnReqList) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthnReqList) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AuthnReqList) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AuthnReqList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthnReqList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthnReqList) SetName(v string)`

SetName sets Name field to given value.


### GetAuthnReqs

`func (o *AuthnReqList) GetAuthnReqs() []string`

GetAuthnReqs returns the AuthnReqs field if non-nil, zero value otherwise.

### GetAuthnReqsOk

`func (o *AuthnReqList) GetAuthnReqsOk() (*[]string, bool)`

GetAuthnReqsOk returns a tuple with the AuthnReqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthnReqs

`func (o *AuthnReqList) SetAuthnReqs(v []string)`

SetAuthnReqs sets AuthnReqs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


