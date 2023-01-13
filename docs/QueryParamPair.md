# QueryParamPair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to [**QueryParamName**](QueryParamName.md) |  | [optional] 
**Value** | Pointer to [**QueryParamValue**](QueryParamValue.md) |  | [optional] 

## Methods

### NewQueryParamPair

`func NewQueryParamPair() *QueryParamPair`

NewQueryParamPair instantiates a new QueryParamPair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryParamPairWithDefaults

`func NewQueryParamPairWithDefaults() *QueryParamPair`

NewQueryParamPairWithDefaults instantiates a new QueryParamPair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *QueryParamPair) GetName() QueryParamName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QueryParamPair) GetNameOk() (*QueryParamName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QueryParamPair) SetName(v QueryParamName)`

SetName sets Name field to given value.

### HasName

`func (o *QueryParamPair) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *QueryParamPair) GetValue() QueryParamValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *QueryParamPair) GetValueOk() (*QueryParamValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *QueryParamPair) SetValue(v QueryParamValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *QueryParamPair) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


