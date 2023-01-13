# ConfigurationParentField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldName** | **string** | The configuration parent field name. | 
**DependentFieldOptions** | [**[]ConfigurationDependentFieldOption**](ConfigurationDependentFieldOption.md) | The dependent field options attributes. | 

## Methods

### NewConfigurationParentField

`func NewConfigurationParentField(fieldName string, dependentFieldOptions []ConfigurationDependentFieldOption, ) *ConfigurationParentField`

NewConfigurationParentField instantiates a new ConfigurationParentField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationParentFieldWithDefaults

`func NewConfigurationParentFieldWithDefaults() *ConfigurationParentField`

NewConfigurationParentFieldWithDefaults instantiates a new ConfigurationParentField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldName

`func (o *ConfigurationParentField) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *ConfigurationParentField) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *ConfigurationParentField) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.


### GetDependentFieldOptions

`func (o *ConfigurationParentField) GetDependentFieldOptions() []ConfigurationDependentFieldOption`

GetDependentFieldOptions returns the DependentFieldOptions field if non-nil, zero value otherwise.

### GetDependentFieldOptionsOk

`func (o *ConfigurationParentField) GetDependentFieldOptionsOk() (*[]ConfigurationDependentFieldOption, bool)`

GetDependentFieldOptionsOk returns a tuple with the DependentFieldOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentFieldOptions

`func (o *ConfigurationParentField) SetDependentFieldOptions(v []ConfigurationDependentFieldOption)`

SetDependentFieldOptions sets DependentFieldOptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


