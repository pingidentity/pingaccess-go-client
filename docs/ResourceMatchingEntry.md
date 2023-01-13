# ResourceMatchingEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the associated resource. | 
**Type** | [**EntryType**](EntryType.md) |  | 
**Link** | [**Link**](Link.md) |  | 
**Methods** | **[]string** | A set of HTTP methods configured for the resource, or &#39;*&#39; to indicate any method. | 
**Pattern** | **string** | A path-matching pattern, relative to the Application context root (interpreted according to the pattern &#39;type&#39;). | 
**PatternType** | [**PathPatternType**](PathPatternType.md) |  | 

## Methods

### NewResourceMatchingEntry

`func NewResourceMatchingEntry(name string, type_ EntryType, link Link, methods []string, pattern string, patternType PathPatternType, ) *ResourceMatchingEntry`

NewResourceMatchingEntry instantiates a new ResourceMatchingEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceMatchingEntryWithDefaults

`func NewResourceMatchingEntryWithDefaults() *ResourceMatchingEntry`

NewResourceMatchingEntryWithDefaults instantiates a new ResourceMatchingEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ResourceMatchingEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceMatchingEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceMatchingEntry) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ResourceMatchingEntry) GetType() EntryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourceMatchingEntry) GetTypeOk() (*EntryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourceMatchingEntry) SetType(v EntryType)`

SetType sets Type field to given value.


### GetLink

`func (o *ResourceMatchingEntry) GetLink() Link`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *ResourceMatchingEntry) GetLinkOk() (*Link, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *ResourceMatchingEntry) SetLink(v Link)`

SetLink sets Link field to given value.


### GetMethods

`func (o *ResourceMatchingEntry) GetMethods() []string`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *ResourceMatchingEntry) GetMethodsOk() (*[]string, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *ResourceMatchingEntry) SetMethods(v []string)`

SetMethods sets Methods field to given value.


### GetPattern

`func (o *ResourceMatchingEntry) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *ResourceMatchingEntry) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *ResourceMatchingEntry) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### GetPatternType

`func (o *ResourceMatchingEntry) GetPatternType() PathPatternType`

GetPatternType returns the PatternType field if non-nil, zero value otherwise.

### GetPatternTypeOk

`func (o *ResourceMatchingEntry) GetPatternTypeOk() (*PathPatternType, bool)`

GetPatternTypeOk returns a tuple with the PatternType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatternType

`func (o *ResourceMatchingEntry) SetPatternType(v PathPatternType)`

SetPatternType sets PatternType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


