# Descriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationFields** | [**[]ConfigurationField**](ConfigurationField.md) | The list of configuration fields for the class. | 
**Type** | **string** | The type for the entity. | 
**Label** | **string** | The label for the entity. | 
**ClassName** | **string** | The name of the implementation class. | 

## Methods

### NewDescriptor

`func NewDescriptor(configurationFields []ConfigurationField, type_ string, label string, className string, ) *Descriptor`

NewDescriptor instantiates a new Descriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescriptorWithDefaults

`func NewDescriptorWithDefaults() *Descriptor`

NewDescriptorWithDefaults instantiates a new Descriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationFields

`func (o *Descriptor) GetConfigurationFields() []ConfigurationField`

GetConfigurationFields returns the ConfigurationFields field if non-nil, zero value otherwise.

### GetConfigurationFieldsOk

`func (o *Descriptor) GetConfigurationFieldsOk() (*[]ConfigurationField, bool)`

GetConfigurationFieldsOk returns a tuple with the ConfigurationFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationFields

`func (o *Descriptor) SetConfigurationFields(v []ConfigurationField)`

SetConfigurationFields sets ConfigurationFields field to given value.


### GetType

`func (o *Descriptor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Descriptor) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Descriptor) SetType(v string)`

SetType sets Type field to given value.


### GetLabel

`func (o *Descriptor) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Descriptor) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Descriptor) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetClassName

`func (o *Descriptor) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *Descriptor) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *Descriptor) SetClassName(v string)`

SetClassName sets ClassName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


