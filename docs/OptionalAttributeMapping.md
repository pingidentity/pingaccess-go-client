# OptionalAttributeMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Set to true to enable the role in the system. | [optional] 
**Attributes** | [**[]Attribute**](Attribute.md) | The attributes that define the role. | 

## Methods

### NewOptionalAttributeMapping

`func NewOptionalAttributeMapping(attributes []Attribute, ) *OptionalAttributeMapping`

NewOptionalAttributeMapping instantiates a new OptionalAttributeMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionalAttributeMappingWithDefaults

`func NewOptionalAttributeMappingWithDefaults() *OptionalAttributeMapping`

NewOptionalAttributeMappingWithDefaults instantiates a new OptionalAttributeMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *OptionalAttributeMapping) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OptionalAttributeMapping) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OptionalAttributeMapping) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OptionalAttributeMapping) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAttributes

`func (o *OptionalAttributeMapping) GetAttributes() []Attribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OptionalAttributeMapping) GetAttributesOk() (*[]Attribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OptionalAttributeMapping) SetAttributes(v []Attribute)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


