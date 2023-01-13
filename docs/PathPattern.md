# PathPattern

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pattern** | Pointer to **string** | The path-matching pattern, relative to the Application context root (interpreted according to the pattern &#39;type&#39;). | [optional] 
**Type** | Pointer to [**PathPatternType**](PathPatternType.md) |  | [optional] 

## Methods

### NewPathPattern

`func NewPathPattern() *PathPattern`

NewPathPattern instantiates a new PathPattern object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPathPatternWithDefaults

`func NewPathPatternWithDefaults() *PathPattern`

NewPathPatternWithDefaults instantiates a new PathPattern object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPattern

`func (o *PathPattern) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *PathPattern) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *PathPattern) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *PathPattern) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetType

`func (o *PathPattern) GetType() PathPatternType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PathPattern) GetTypeOk() (*PathPatternType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PathPattern) SetType(v PathPatternType)`

SetType sets Type field to given value.

### HasType

`func (o *PathPattern) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


