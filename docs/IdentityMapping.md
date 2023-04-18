# IdentityMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassName** | **string** | (sortable) The identity mapping&#39;s class name. | 
**Id** | Pointer to **int64** | When creating a new IdentityMapping, this is the ID for the IdentityMapping. If not specified, an ID will be automatically assigned. When updating an existing IdentityMapping, this field is ignored and the ID is determined by the path parameter. | [optional] 
**Name** | **string** | (sortable) The name of the identity mapping. | 
**Configuration** | Pointer to **map[string]interface{}** | The identity mapping&#39;s configuration data. - This value is a PingAccess plugin configuration (JSON). | [optional] 

## Methods

### NewIdentityMapping

`func NewIdentityMapping(className string, name string, ) *IdentityMapping`

NewIdentityMapping instantiates a new IdentityMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityMappingWithDefaults

`func NewIdentityMappingWithDefaults() *IdentityMapping`

NewIdentityMappingWithDefaults instantiates a new IdentityMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassName

`func (o *IdentityMapping) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *IdentityMapping) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *IdentityMapping) SetClassName(v string)`

SetClassName sets ClassName field to given value.


### GetId

`func (o *IdentityMapping) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityMapping) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityMapping) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityMapping) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IdentityMapping) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityMapping) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityMapping) SetName(v string)`

SetName sets Name field to given value.


### GetConfiguration

`func (o *IdentityMapping) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *IdentityMapping) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *IdentityMapping) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *IdentityMapping) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


