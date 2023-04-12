# EngineHealthStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentServerTime** | **int64** | The current epoch time of the server in milliseconds. | 
**EnginesStatus** | [**map[string]EngineInfo**](EngineInfo.md) | Map of engines with their details. Details include name, description, last updated time, and polling delay.  Key type: String Value type: EngineInfo | 

## Methods

### NewEngineHealthStatus

`func NewEngineHealthStatus(currentServerTime int64, enginesStatus map[string]EngineInfo, ) *EngineHealthStatus`

NewEngineHealthStatus instantiates a new EngineHealthStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineHealthStatusWithDefaults

`func NewEngineHealthStatusWithDefaults() *EngineHealthStatus`

NewEngineHealthStatusWithDefaults instantiates a new EngineHealthStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentServerTime

`func (o *EngineHealthStatus) GetCurrentServerTime() int64`

GetCurrentServerTime returns the CurrentServerTime field if non-nil, zero value otherwise.

### GetCurrentServerTimeOk

`func (o *EngineHealthStatus) GetCurrentServerTimeOk() (*int64, bool)`

GetCurrentServerTimeOk returns a tuple with the CurrentServerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentServerTime

`func (o *EngineHealthStatus) SetCurrentServerTime(v int64)`

SetCurrentServerTime sets CurrentServerTime field to given value.


### GetEnginesStatus

`func (o *EngineHealthStatus) GetEnginesStatus() map[string]EngineInfo`

GetEnginesStatus returns the EnginesStatus field if non-nil, zero value otherwise.

### GetEnginesStatusOk

`func (o *EngineHealthStatus) GetEnginesStatusOk() (*map[string]EngineInfo, bool)`

GetEnginesStatusOk returns a tuple with the EnginesStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnginesStatus

`func (o *EngineHealthStatus) SetEnginesStatus(v map[string]EngineInfo)`

SetEnginesStatus sets EnginesStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


