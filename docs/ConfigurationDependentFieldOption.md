# ConfigurationDependentFieldOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** | The dependent field option value. | 
**Options** | [**[]ConfigurationOption**](ConfigurationOption.md) | The dependent field options. | 

## Methods

### NewConfigurationDependentFieldOption

`func NewConfigurationDependentFieldOption(value string, options []ConfigurationOption, ) *ConfigurationDependentFieldOption`

NewConfigurationDependentFieldOption instantiates a new ConfigurationDependentFieldOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationDependentFieldOptionWithDefaults

`func NewConfigurationDependentFieldOptionWithDefaults() *ConfigurationDependentFieldOption`

NewConfigurationDependentFieldOptionWithDefaults instantiates a new ConfigurationDependentFieldOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ConfigurationDependentFieldOption) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConfigurationDependentFieldOption) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConfigurationDependentFieldOption) SetValue(v string)`

SetValue sets Value field to given value.


### GetOptions

`func (o *ConfigurationDependentFieldOption) GetOptions() []ConfigurationOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ConfigurationDependentFieldOption) GetOptionsOk() (*[]ConfigurationOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ConfigurationDependentFieldOption) SetOptions(v []ConfigurationOption)`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


