# AvailabilityProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassName** | **string** | (sortable) The class name of the availability profile. | 
**Id** | Pointer to **int32** | When creating a new AvailabilityProfile, this is the ID for the AvailabilityProfile. If not specified, an ID will be automatically assigned. When updating an existing AvailabilityProfile, this field is ignored and the ID is determined by the path parameter. | [optional] 
**Name** | **string** | (sortable) The name of the availability profile. | 
**Configuration** | Pointer to **map[string]interface{}** | The availability profile&#39;s configuration data. - This value is a PingAccess plugin configuration (JSON). | [optional] 

## Methods

### NewAvailabilityProfile

`func NewAvailabilityProfile(className string, name string, ) *AvailabilityProfile`

NewAvailabilityProfile instantiates a new AvailabilityProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailabilityProfileWithDefaults

`func NewAvailabilityProfileWithDefaults() *AvailabilityProfile`

NewAvailabilityProfileWithDefaults instantiates a new AvailabilityProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassName

`func (o *AvailabilityProfile) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *AvailabilityProfile) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *AvailabilityProfile) SetClassName(v string)`

SetClassName sets ClassName field to given value.


### GetId

`func (o *AvailabilityProfile) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AvailabilityProfile) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AvailabilityProfile) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AvailabilityProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AvailabilityProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AvailabilityProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AvailabilityProfile) SetName(v string)`

SetName sets Name field to given value.


### GetConfiguration

`func (o *AvailabilityProfile) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *AvailabilityProfile) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *AvailabilityProfile) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *AvailabilityProfile) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


