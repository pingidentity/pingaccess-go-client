# Rule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassName** | **string** | (sortable) The rule&#39;s class name. | 
**Id** | Pointer to **int64** | When creating a new Rule, this is the ID for the Rule. If not specified, an ID will be automatically assigned. When updating an existing Rule, this field is ignored and the ID is determined by the path parameter. | [optional] 
**Name** | **string** | (sortable) The rule&#39;s name. | 
**SupportedDestinations** | Pointer to **[]string** | The supported destinations for this rule. This field is read-only and meant to provide contextual metadata on where the rule can be applied. | [optional] 
**Configuration** | Pointer to **map[string]interface{}** | The rule&#39;s configuration data. - This value is a PingAccess plugin configuration (JSON). | [optional] 

## Methods

### NewRule

`func NewRule(className string, name string, ) *Rule`

NewRule instantiates a new Rule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleWithDefaults

`func NewRuleWithDefaults() *Rule`

NewRuleWithDefaults instantiates a new Rule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassName

`func (o *Rule) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *Rule) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *Rule) SetClassName(v string)`

SetClassName sets ClassName field to given value.


### GetId

`func (o *Rule) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Rule) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Rule) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Rule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Rule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Rule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Rule) SetName(v string)`

SetName sets Name field to given value.


### GetSupportedDestinations

`func (o *Rule) GetSupportedDestinations() []string`

GetSupportedDestinations returns the SupportedDestinations field if non-nil, zero value otherwise.

### GetSupportedDestinationsOk

`func (o *Rule) GetSupportedDestinationsOk() (*[]string, bool)`

GetSupportedDestinationsOk returns a tuple with the SupportedDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedDestinations

`func (o *Rule) SetSupportedDestinations(v []string)`

SetSupportedDestinations sets SupportedDestinations field to given value.

### HasSupportedDestinations

`func (o *Rule) HasSupportedDestinations() bool`

HasSupportedDestinations returns a boolean if a field has been set.

### GetConfiguration

`func (o *Rule) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *Rule) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *Rule) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *Rule) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


