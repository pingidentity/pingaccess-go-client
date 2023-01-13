# GeneralName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | [**GeneralNameEnum**](GeneralNameEnum.md) |  | 
**Value** | **string** |  | 

## Methods

### NewGeneralName

`func NewGeneralName(name GeneralNameEnum, value string, ) *GeneralName`

NewGeneralName instantiates a new GeneralName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneralNameWithDefaults

`func NewGeneralNameWithDefaults() *GeneralName`

NewGeneralNameWithDefaults instantiates a new GeneralName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GeneralName) GetName() GeneralNameEnum`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GeneralName) GetNameOk() (*GeneralNameEnum, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GeneralName) SetName(v GeneralNameEnum)`

SetName sets Name field to given value.


### GetValue

`func (o *GeneralName) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GeneralName) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GeneralName) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


