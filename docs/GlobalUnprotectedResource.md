# GlobalUnprotectedResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | When creating a new GlobalUnprotectedResource, this is the ID for the GlobalUnprotectedResource. If not specified, an ID will be automatically assigned. When updating an existing GlobalUnprotectedResource, this field is ignored and the ID is determined by the path parameter. | [optional] 
**Name** | **string** | (sortable) The name of the global unprotected resource. | 
**WildcardPath** | **string** | A path for the global unprotected resource. | 
**AuditLevel** | Pointer to [**AuditLevel**](AuditLevel.md) |  | [optional] 
**Enabled** | Pointer to **bool** | (sortable) True if the global resource is enabled. | [optional] 
**Description** | Pointer to **string** | (sortable) A description of the global unprotected resource. | [optional] 

## Methods

### NewGlobalUnprotectedResource

`func NewGlobalUnprotectedResource(name string, wildcardPath string, ) *GlobalUnprotectedResource`

NewGlobalUnprotectedResource instantiates a new GlobalUnprotectedResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalUnprotectedResourceWithDefaults

`func NewGlobalUnprotectedResourceWithDefaults() *GlobalUnprotectedResource`

NewGlobalUnprotectedResourceWithDefaults instantiates a new GlobalUnprotectedResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GlobalUnprotectedResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GlobalUnprotectedResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GlobalUnprotectedResource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GlobalUnprotectedResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GlobalUnprotectedResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GlobalUnprotectedResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GlobalUnprotectedResource) SetName(v string)`

SetName sets Name field to given value.


### GetWildcardPath

`func (o *GlobalUnprotectedResource) GetWildcardPath() string`

GetWildcardPath returns the WildcardPath field if non-nil, zero value otherwise.

### GetWildcardPathOk

`func (o *GlobalUnprotectedResource) GetWildcardPathOk() (*string, bool)`

GetWildcardPathOk returns a tuple with the WildcardPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWildcardPath

`func (o *GlobalUnprotectedResource) SetWildcardPath(v string)`

SetWildcardPath sets WildcardPath field to given value.


### GetAuditLevel

`func (o *GlobalUnprotectedResource) GetAuditLevel() AuditLevel`

GetAuditLevel returns the AuditLevel field if non-nil, zero value otherwise.

### GetAuditLevelOk

`func (o *GlobalUnprotectedResource) GetAuditLevelOk() (*AuditLevel, bool)`

GetAuditLevelOk returns a tuple with the AuditLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLevel

`func (o *GlobalUnprotectedResource) SetAuditLevel(v AuditLevel)`

SetAuditLevel sets AuditLevel field to given value.

### HasAuditLevel

`func (o *GlobalUnprotectedResource) HasAuditLevel() bool`

HasAuditLevel returns a boolean if a field has been set.

### GetEnabled

`func (o *GlobalUnprotectedResource) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GlobalUnprotectedResource) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GlobalUnprotectedResource) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GlobalUnprotectedResource) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *GlobalUnprotectedResource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GlobalUnprotectedResource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GlobalUnprotectedResource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GlobalUnprotectedResource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


