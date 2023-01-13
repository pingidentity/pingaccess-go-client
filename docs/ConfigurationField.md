# ConfigurationField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the configuration field. | 
**Label** | **string** | The label of the configuration field. | 
**Type** | [**ConfigurationType**](ConfigurationType.md) |  | 
**Advanced** | **bool** | Indicates that the configuration field is an advanced field or not. | 
**Required** | **bool** | Indicates that the configuration field value is required or not. | 
**Help** | [**Help**](Help.md) |  | 
**Fields** | [**[]ConfigurationField**](ConfigurationField.md) | The list of configuration fields that the current configuration field is the parent of. | 
**Options** | [**[]ConfigurationOption**](ConfigurationOption.md) | The available options for the select based configuration fields. | 
**ParentField** | [**ConfigurationParentField**](ConfigurationParentField.md) |  | 
**ButtonGroup** | **string** | The name of group radio buttons that cooperate in a single selection. | 
**Deselectable** | **bool** | Indicates that a radio button is able to be deselected. | 
**Default** | Pointer to **string** | The default value of the configuration field. | [optional] 

## Methods

### NewConfigurationField

`func NewConfigurationField(name string, label string, type_ ConfigurationType, advanced bool, required bool, help Help, fields []ConfigurationField, options []ConfigurationOption, parentField ConfigurationParentField, buttonGroup string, deselectable bool, ) *ConfigurationField`

NewConfigurationField instantiates a new ConfigurationField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationFieldWithDefaults

`func NewConfigurationFieldWithDefaults() *ConfigurationField`

NewConfigurationFieldWithDefaults instantiates a new ConfigurationField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConfigurationField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigurationField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigurationField) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *ConfigurationField) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ConfigurationField) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ConfigurationField) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetType

`func (o *ConfigurationField) GetType() ConfigurationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConfigurationField) GetTypeOk() (*ConfigurationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConfigurationField) SetType(v ConfigurationType)`

SetType sets Type field to given value.


### GetAdvanced

`func (o *ConfigurationField) GetAdvanced() bool`

GetAdvanced returns the Advanced field if non-nil, zero value otherwise.

### GetAdvancedOk

`func (o *ConfigurationField) GetAdvancedOk() (*bool, bool)`

GetAdvancedOk returns a tuple with the Advanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvanced

`func (o *ConfigurationField) SetAdvanced(v bool)`

SetAdvanced sets Advanced field to given value.


### GetRequired

`func (o *ConfigurationField) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ConfigurationField) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ConfigurationField) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetHelp

`func (o *ConfigurationField) GetHelp() Help`

GetHelp returns the Help field if non-nil, zero value otherwise.

### GetHelpOk

`func (o *ConfigurationField) GetHelpOk() (*Help, bool)`

GetHelpOk returns a tuple with the Help field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelp

`func (o *ConfigurationField) SetHelp(v Help)`

SetHelp sets Help field to given value.


### GetFields

`func (o *ConfigurationField) GetFields() []ConfigurationField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ConfigurationField) GetFieldsOk() (*[]ConfigurationField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ConfigurationField) SetFields(v []ConfigurationField)`

SetFields sets Fields field to given value.


### GetOptions

`func (o *ConfigurationField) GetOptions() []ConfigurationOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ConfigurationField) GetOptionsOk() (*[]ConfigurationOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ConfigurationField) SetOptions(v []ConfigurationOption)`

SetOptions sets Options field to given value.


### GetParentField

`func (o *ConfigurationField) GetParentField() ConfigurationParentField`

GetParentField returns the ParentField field if non-nil, zero value otherwise.

### GetParentFieldOk

`func (o *ConfigurationField) GetParentFieldOk() (*ConfigurationParentField, bool)`

GetParentFieldOk returns a tuple with the ParentField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentField

`func (o *ConfigurationField) SetParentField(v ConfigurationParentField)`

SetParentField sets ParentField field to given value.


### GetButtonGroup

`func (o *ConfigurationField) GetButtonGroup() string`

GetButtonGroup returns the ButtonGroup field if non-nil, zero value otherwise.

### GetButtonGroupOk

`func (o *ConfigurationField) GetButtonGroupOk() (*string, bool)`

GetButtonGroupOk returns a tuple with the ButtonGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonGroup

`func (o *ConfigurationField) SetButtonGroup(v string)`

SetButtonGroup sets ButtonGroup field to given value.


### GetDeselectable

`func (o *ConfigurationField) GetDeselectable() bool`

GetDeselectable returns the Deselectable field if non-nil, zero value otherwise.

### GetDeselectableOk

`func (o *ConfigurationField) GetDeselectableOk() (*bool, bool)`

GetDeselectableOk returns a tuple with the Deselectable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeselectable

`func (o *ConfigurationField) SetDeselectable(v bool)`

SetDeselectable sets Deselectable field to given value.


### GetDefault

`func (o *ConfigurationField) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *ConfigurationField) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *ConfigurationField) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *ConfigurationField) HasDefault() bool`

HasDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


