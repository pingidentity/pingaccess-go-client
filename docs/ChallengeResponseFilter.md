# ChallengeResponseFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassName** | **string** | The class name of the challenge response filter. | 
**Configuration** | Pointer to **map[string]interface{}** | The challenge response filter configuration. - This value is a PingAccess plugin configuration (JSON). | [optional] 

## Methods

### NewChallengeResponseFilter

`func NewChallengeResponseFilter(className string, ) *ChallengeResponseFilter`

NewChallengeResponseFilter instantiates a new ChallengeResponseFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChallengeResponseFilterWithDefaults

`func NewChallengeResponseFilterWithDefaults() *ChallengeResponseFilter`

NewChallengeResponseFilterWithDefaults instantiates a new ChallengeResponseFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassName

`func (o *ChallengeResponseFilter) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *ChallengeResponseFilter) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *ChallengeResponseFilter) SetClassName(v string)`

SetClassName sets ClassName field to given value.


### GetConfiguration

`func (o *ChallengeResponseFilter) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ChallengeResponseFilter) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ChallengeResponseFilter) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *ChallengeResponseFilter) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


