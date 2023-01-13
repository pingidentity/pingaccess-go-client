# QueryParamName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pattern** | Pointer to **string** | The string value of the query parameter name. | [optional] 
**Type** | Pointer to [**QueryParamPatternType**](QueryParamPatternType.md) |  | [optional] 

## Methods

### NewQueryParamName

`func NewQueryParamName() *QueryParamName`

NewQueryParamName instantiates a new QueryParamName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryParamNameWithDefaults

`func NewQueryParamNameWithDefaults() *QueryParamName`

NewQueryParamNameWithDefaults instantiates a new QueryParamName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPattern

`func (o *QueryParamName) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *QueryParamName) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *QueryParamName) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *QueryParamName) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetType

`func (o *QueryParamName) GetType() QueryParamPatternType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QueryParamName) GetTypeOk() (*QueryParamPatternType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QueryParamName) SetType(v QueryParamPatternType)`

SetType sets Type field to given value.

### HasType

`func (o *QueryParamName) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


