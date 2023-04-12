# EngineInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PollingDelay** | **int64** | The polling delay of the engine. | 
**CreatedAt** | **int64** | The date the engine was created. | 
**LastUpdated** | **int64** | The last updated date. | 
**Description** | **string** | The description of the engine. | 
**Name** | **string** | The name of the engine. | 

## Methods

### NewEngineInfo

`func NewEngineInfo(pollingDelay int64, createdAt int64, lastUpdated int64, description string, name string, ) *EngineInfo`

NewEngineInfo instantiates a new EngineInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineInfoWithDefaults

`func NewEngineInfoWithDefaults() *EngineInfo`

NewEngineInfoWithDefaults instantiates a new EngineInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPollingDelay

`func (o *EngineInfo) GetPollingDelay() int64`

GetPollingDelay returns the PollingDelay field if non-nil, zero value otherwise.

### GetPollingDelayOk

`func (o *EngineInfo) GetPollingDelayOk() (*int64, bool)`

GetPollingDelayOk returns a tuple with the PollingDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingDelay

`func (o *EngineInfo) SetPollingDelay(v int64)`

SetPollingDelay sets PollingDelay field to given value.


### GetCreatedAt

`func (o *EngineInfo) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EngineInfo) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EngineInfo) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetLastUpdated

`func (o *EngineInfo) GetLastUpdated() int64`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *EngineInfo) GetLastUpdatedOk() (*int64, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *EngineInfo) SetLastUpdated(v int64)`

SetLastUpdated sets LastUpdated field to given value.


### GetDescription

`func (o *EngineInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EngineInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EngineInfo) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *EngineInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EngineInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EngineInfo) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


