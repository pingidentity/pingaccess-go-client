# RequestMatcher

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassName** | **string** | The class name of the request matcher. | 
**Configuration** | Pointer to **string** | The request matcher configuration. | [optional] 

## Methods

### NewRequestMatcher

`func NewRequestMatcher(className string, ) *RequestMatcher`

NewRequestMatcher instantiates a new RequestMatcher object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestMatcherWithDefaults

`func NewRequestMatcherWithDefaults() *RequestMatcher`

NewRequestMatcherWithDefaults instantiates a new RequestMatcher object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassName

`func (o *RequestMatcher) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *RequestMatcher) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *RequestMatcher) SetClassName(v string)`

SetClassName sets ClassName field to given value.


### GetConfiguration

`func (o *RequestMatcher) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *RequestMatcher) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *RequestMatcher) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *RequestMatcher) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


