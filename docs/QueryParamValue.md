# QueryParamValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pattern** | Pointer to **string** | The string value of the query parameter value. | [optional] 
**Type** | Pointer to [**QueryParamPatternType**](QueryParamPatternType.md) |  | [optional] 
**MatchAny** | Pointer to **bool** | Matching any value. Ignoring the specified value. | [optional] 

## Methods

### NewQueryParamValue

`func NewQueryParamValue() *QueryParamValue`

NewQueryParamValue instantiates a new QueryParamValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryParamValueWithDefaults

`func NewQueryParamValueWithDefaults() *QueryParamValue`

NewQueryParamValueWithDefaults instantiates a new QueryParamValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPattern

`func (o *QueryParamValue) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *QueryParamValue) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *QueryParamValue) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *QueryParamValue) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetType

`func (o *QueryParamValue) GetType() QueryParamPatternType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QueryParamValue) GetTypeOk() (*QueryParamPatternType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QueryParamValue) SetType(v QueryParamPatternType)`

SetType sets Type field to given value.

### HasType

`func (o *QueryParamValue) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMatchAny

`func (o *QueryParamValue) GetMatchAny() bool`

GetMatchAny returns the MatchAny field if non-nil, zero value otherwise.

### GetMatchAnyOk

`func (o *QueryParamValue) GetMatchAnyOk() (*bool, bool)`

GetMatchAnyOk returns a tuple with the MatchAny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchAny

`func (o *QueryParamValue) SetMatchAny(v bool)`

SetMatchAny sets MatchAny field to given value.

### HasMatchAny

`func (o *QueryParamValue) HasMatchAny() bool`

HasMatchAny returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


