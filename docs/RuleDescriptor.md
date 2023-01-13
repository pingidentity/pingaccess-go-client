# RuleDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationFields** | [**[]ConfigurationField**](ConfigurationField.md) | The list of configuration fields of the rule | 
**Type** | **string** | The type of the rule. | 
**Label** | **string** | The label of the rule. | 
**ClassName** | **string** | The class name of the rule. | 
**Category** | [**RuleInterceptorCategory**](RuleInterceptorCategory.md) |  | 
**Modes** | [**[]RuleInterceptorSupportedDestination**](RuleInterceptorSupportedDestination.md) | The modes that the rules can be used in. Site mode implies Sideband and Site destination support. | 
**AgentCachingDisabled** | **bool** | Indicates that agent caching is disabled or not. | 

## Methods

### NewRuleDescriptor

`func NewRuleDescriptor(configurationFields []ConfigurationField, type_ string, label string, className string, category RuleInterceptorCategory, modes []RuleInterceptorSupportedDestination, agentCachingDisabled bool, ) *RuleDescriptor`

NewRuleDescriptor instantiates a new RuleDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleDescriptorWithDefaults

`func NewRuleDescriptorWithDefaults() *RuleDescriptor`

NewRuleDescriptorWithDefaults instantiates a new RuleDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationFields

`func (o *RuleDescriptor) GetConfigurationFields() []ConfigurationField`

GetConfigurationFields returns the ConfigurationFields field if non-nil, zero value otherwise.

### GetConfigurationFieldsOk

`func (o *RuleDescriptor) GetConfigurationFieldsOk() (*[]ConfigurationField, bool)`

GetConfigurationFieldsOk returns a tuple with the ConfigurationFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationFields

`func (o *RuleDescriptor) SetConfigurationFields(v []ConfigurationField)`

SetConfigurationFields sets ConfigurationFields field to given value.


### GetType

`func (o *RuleDescriptor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RuleDescriptor) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RuleDescriptor) SetType(v string)`

SetType sets Type field to given value.


### GetLabel

`func (o *RuleDescriptor) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *RuleDescriptor) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *RuleDescriptor) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetClassName

`func (o *RuleDescriptor) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *RuleDescriptor) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *RuleDescriptor) SetClassName(v string)`

SetClassName sets ClassName field to given value.


### GetCategory

`func (o *RuleDescriptor) GetCategory() RuleInterceptorCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *RuleDescriptor) GetCategoryOk() (*RuleInterceptorCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *RuleDescriptor) SetCategory(v RuleInterceptorCategory)`

SetCategory sets Category field to given value.


### GetModes

`func (o *RuleDescriptor) GetModes() []RuleInterceptorSupportedDestination`

GetModes returns the Modes field if non-nil, zero value otherwise.

### GetModesOk

`func (o *RuleDescriptor) GetModesOk() (*[]RuleInterceptorSupportedDestination, bool)`

GetModesOk returns a tuple with the Modes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModes

`func (o *RuleDescriptor) SetModes(v []RuleInterceptorSupportedDestination)`

SetModes sets Modes field to given value.


### GetAgentCachingDisabled

`func (o *RuleDescriptor) GetAgentCachingDisabled() bool`

GetAgentCachingDisabled returns the AgentCachingDisabled field if non-nil, zero value otherwise.

### GetAgentCachingDisabledOk

`func (o *RuleDescriptor) GetAgentCachingDisabledOk() (*bool, bool)`

GetAgentCachingDisabledOk returns a tuple with the AgentCachingDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentCachingDisabled

`func (o *RuleDescriptor) SetAgentCachingDisabled(v bool)`

SetAgentCachingDisabled sets AgentCachingDisabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


