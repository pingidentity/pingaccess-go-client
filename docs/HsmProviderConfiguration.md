# HsmProviderConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to **string** |  | [optional] 
**Password** | Pointer to [**HiddenField**](HiddenField.md) |  | [optional] 
**SlotId** | Pointer to **string** |  | [optional] 
**Library** | Pointer to **string** |  | [optional] 

## Methods

### NewHsmProviderConfiguration

`func NewHsmProviderConfiguration() *HsmProviderConfiguration`

NewHsmProviderConfiguration instantiates a new HsmProviderConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHsmProviderConfigurationWithDefaults

`func NewHsmProviderConfigurationWithDefaults() *HsmProviderConfiguration`

NewHsmProviderConfigurationWithDefaults instantiates a new HsmProviderConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *HsmProviderConfiguration) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *HsmProviderConfiguration) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *HsmProviderConfiguration) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *HsmProviderConfiguration) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetPassword

`func (o *HsmProviderConfiguration) GetPassword() HiddenField`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *HsmProviderConfiguration) GetPasswordOk() (*HiddenField, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *HsmProviderConfiguration) SetPassword(v HiddenField)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *HsmProviderConfiguration) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSlotId

`func (o *HsmProviderConfiguration) GetSlotId() string`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *HsmProviderConfiguration) GetSlotIdOk() (*string, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *HsmProviderConfiguration) SetSlotId(v string)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *HsmProviderConfiguration) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.

### GetLibrary

`func (o *HsmProviderConfiguration) GetLibrary() string`

GetLibrary returns the Library field if non-nil, zero value otherwise.

### GetLibraryOk

`func (o *HsmProviderConfiguration) GetLibraryOk() (*string, bool)`

GetLibraryOk returns a tuple with the Library field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibrary

`func (o *HsmProviderConfiguration) SetLibrary(v string)`

SetLibrary sets Library field to given value.

### HasLibrary

`func (o *HsmProviderConfiguration) HasLibrary() bool`

HasLibrary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


