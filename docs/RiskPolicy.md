# RiskPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassName** | **string** | (sortable) The class name of the risk policy implementation. | 
**Id** | Pointer to **int32** | When creating a new RiskPolicy, this is the ID for the RiskPolicy. If not specified, an ID will be automatically assigned. When updating an existing RiskPolicy, this field is ignored and the ID is determined by the path parameter. | [optional] 
**Name** | **string** | (sortable) The name of the risk policy. | 
**Configuration** | Pointer to **map[string]interface{}** | The configuration data  of the risk policy. - This value is a PingAccess plugin configuration (JSON). | [optional] 

## Methods

### NewRiskPolicy

`func NewRiskPolicy(className string, name string, ) *RiskPolicy`

NewRiskPolicy instantiates a new RiskPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPolicyWithDefaults

`func NewRiskPolicyWithDefaults() *RiskPolicy`

NewRiskPolicyWithDefaults instantiates a new RiskPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassName

`func (o *RiskPolicy) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *RiskPolicy) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *RiskPolicy) SetClassName(v string)`

SetClassName sets ClassName field to given value.


### GetId

`func (o *RiskPolicy) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskPolicy) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskPolicy) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RiskPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RiskPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RiskPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RiskPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetConfiguration

`func (o *RiskPolicy) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *RiskPolicy) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *RiskPolicy) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *RiskPolicy) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


