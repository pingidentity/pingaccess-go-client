# QueryParamConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchesNoParams** | Pointer to **bool** | Enable this setting to match requests to URLs without query parameters in addition URLs with query parameters. | [optional] 
**Params** | Pointer to [**[]QueryParamPair**](QueryParamPair.md) | The query parameter name/value pairs. | [optional] 

## Methods

### NewQueryParamConfig

`func NewQueryParamConfig() *QueryParamConfig`

NewQueryParamConfig instantiates a new QueryParamConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryParamConfigWithDefaults

`func NewQueryParamConfigWithDefaults() *QueryParamConfig`

NewQueryParamConfigWithDefaults instantiates a new QueryParamConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchesNoParams

`func (o *QueryParamConfig) GetMatchesNoParams() bool`

GetMatchesNoParams returns the MatchesNoParams field if non-nil, zero value otherwise.

### GetMatchesNoParamsOk

`func (o *QueryParamConfig) GetMatchesNoParamsOk() (*bool, bool)`

GetMatchesNoParamsOk returns a tuple with the MatchesNoParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchesNoParams

`func (o *QueryParamConfig) SetMatchesNoParams(v bool)`

SetMatchesNoParams sets MatchesNoParams field to given value.

### HasMatchesNoParams

`func (o *QueryParamConfig) HasMatchesNoParams() bool`

HasMatchesNoParams returns a boolean if a field has been set.

### GetParams

`func (o *QueryParamConfig) GetParams() []QueryParamPair`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *QueryParamConfig) GetParamsOk() (*[]QueryParamPair, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *QueryParamConfig) SetParams(v []QueryParamPair)`

SetParams sets Params field to given value.

### HasParams

`func (o *QueryParamConfig) HasParams() bool`

HasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


