# LoadBalancingStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassName** | **string** | (sortable) The class name of the load balancing strategy. | 
**Id** | Pointer to **int32** | When creating a new LoadBalancingStrategy, this is the ID for the LoadBalancingStrategy. If not specified, an ID will be automatically assigned. When updating an existing LoadBalancingStrategy, this field is ignored and the ID is determined by the path parameter. | [optional] 
**Name** | **string** | (sortable) The name of the load balancing strategy. | 
**Configuration** | Pointer to **string** | The load balancing strategy&#39;s configuration data. | [optional] 

## Methods

### NewLoadBalancingStrategy

`func NewLoadBalancingStrategy(className string, name string, ) *LoadBalancingStrategy`

NewLoadBalancingStrategy instantiates a new LoadBalancingStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancingStrategyWithDefaults

`func NewLoadBalancingStrategyWithDefaults() *LoadBalancingStrategy`

NewLoadBalancingStrategyWithDefaults instantiates a new LoadBalancingStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassName

`func (o *LoadBalancingStrategy) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *LoadBalancingStrategy) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *LoadBalancingStrategy) SetClassName(v string)`

SetClassName sets ClassName field to given value.


### GetId

`func (o *LoadBalancingStrategy) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoadBalancingStrategy) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoadBalancingStrategy) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *LoadBalancingStrategy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *LoadBalancingStrategy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoadBalancingStrategy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoadBalancingStrategy) SetName(v string)`

SetName sets Name field to given value.


### GetConfiguration

`func (o *LoadBalancingStrategy) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *LoadBalancingStrategy) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *LoadBalancingStrategy) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *LoadBalancingStrategy) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

