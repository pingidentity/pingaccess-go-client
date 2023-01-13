# Redirect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | When creating a new Redirect, this is the ID for the Redirect. If not specified, an ID will be automatically assigned. When updating an existing Redirect, this field is ignored and the ID is determined by the path parameter. | [optional] 
**Source** | Pointer to [**HostPort**](HostPort.md) |  | [optional] 
**Target** | Pointer to [**TargetHostPort**](TargetHostPort.md) |  | [optional] 
**ResponseCode** | Pointer to **int32** | (sortable) The Redirect HTTP status code used by the response. | [optional] 
**AuditLevel** | Pointer to [**AuditLevel**](AuditLevel.md) |  | [optional] 

## Methods

### NewRedirect

`func NewRedirect() *Redirect`

NewRedirect instantiates a new Redirect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedirectWithDefaults

`func NewRedirectWithDefaults() *Redirect`

NewRedirectWithDefaults instantiates a new Redirect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Redirect) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Redirect) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Redirect) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Redirect) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSource

`func (o *Redirect) GetSource() HostPort`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Redirect) GetSourceOk() (*HostPort, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Redirect) SetSource(v HostPort)`

SetSource sets Source field to given value.

### HasSource

`func (o *Redirect) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTarget

`func (o *Redirect) GetTarget() TargetHostPort`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *Redirect) GetTargetOk() (*TargetHostPort, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *Redirect) SetTarget(v TargetHostPort)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *Redirect) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetResponseCode

`func (o *Redirect) GetResponseCode() int32`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *Redirect) GetResponseCodeOk() (*int32, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *Redirect) SetResponseCode(v int32)`

SetResponseCode sets ResponseCode field to given value.

### HasResponseCode

`func (o *Redirect) HasResponseCode() bool`

HasResponseCode returns a boolean if a field has been set.

### GetAuditLevel

`func (o *Redirect) GetAuditLevel() AuditLevel`

GetAuditLevel returns the AuditLevel field if non-nil, zero value otherwise.

### GetAuditLevelOk

`func (o *Redirect) GetAuditLevelOk() (*AuditLevel, bool)`

GetAuditLevelOk returns a tuple with the AuditLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLevel

`func (o *Redirect) SetAuditLevel(v AuditLevel)`

SetAuditLevel sets AuditLevel field to given value.

### HasAuditLevel

`func (o *Redirect) HasAuditLevel() bool`

HasAuditLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


