# ConfigStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The id of the configuration workflow. | [optional] 
**Status** | Pointer to **string** | The status of the configuration import or export. | [optional] 
**CurrentEntity** | Pointer to **map[string]interface{}** | The current entity being imported or exported. | [optional] 
**TotalEntities** | Pointer to **int64** | The total number of entities being imported or exported. | [optional] 
**ApiErrorView** | Pointer to [**ApiError**](ApiError.md) |  | [optional] 
**Warnings** | **[]string** | The API warnings for import or export, if there are any warnings. | 

## Methods

### NewConfigStatus

`func NewConfigStatus(warnings []string, ) *ConfigStatus`

NewConfigStatus instantiates a new ConfigStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigStatusWithDefaults

`func NewConfigStatusWithDefaults() *ConfigStatus`

NewConfigStatusWithDefaults instantiates a new ConfigStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConfigStatus) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigStatus) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigStatus) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *ConfigStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConfigStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConfigStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConfigStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCurrentEntity

`func (o *ConfigStatus) GetCurrentEntity() map[string]interface{}`

GetCurrentEntity returns the CurrentEntity field if non-nil, zero value otherwise.

### GetCurrentEntityOk

`func (o *ConfigStatus) GetCurrentEntityOk() (*map[string]interface{}, bool)`

GetCurrentEntityOk returns a tuple with the CurrentEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentEntity

`func (o *ConfigStatus) SetCurrentEntity(v map[string]interface{})`

SetCurrentEntity sets CurrentEntity field to given value.

### HasCurrentEntity

`func (o *ConfigStatus) HasCurrentEntity() bool`

HasCurrentEntity returns a boolean if a field has been set.

### GetTotalEntities

`func (o *ConfigStatus) GetTotalEntities() int64`

GetTotalEntities returns the TotalEntities field if non-nil, zero value otherwise.

### GetTotalEntitiesOk

`func (o *ConfigStatus) GetTotalEntitiesOk() (*int64, bool)`

GetTotalEntitiesOk returns a tuple with the TotalEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEntities

`func (o *ConfigStatus) SetTotalEntities(v int64)`

SetTotalEntities sets TotalEntities field to given value.

### HasTotalEntities

`func (o *ConfigStatus) HasTotalEntities() bool`

HasTotalEntities returns a boolean if a field has been set.

### GetApiErrorView

`func (o *ConfigStatus) GetApiErrorView() ApiError`

GetApiErrorView returns the ApiErrorView field if non-nil, zero value otherwise.

### GetApiErrorViewOk

`func (o *ConfigStatus) GetApiErrorViewOk() (*ApiError, bool)`

GetApiErrorViewOk returns a tuple with the ApiErrorView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiErrorView

`func (o *ConfigStatus) SetApiErrorView(v ApiError)`

SetApiErrorView sets ApiErrorView field to given value.

### HasApiErrorView

`func (o *ConfigStatus) HasApiErrorView() bool`

HasApiErrorView returns a boolean if a field has been set.

### GetWarnings

`func (o *ConfigStatus) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *ConfigStatus) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *ConfigStatus) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


