# RuleSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | When creating a new RuleSet, this is the ID for the RuleSet. If not specified, an ID will be automatically assigned. When updating an existing RuleSet, this field is ignored and the ID is determined by the path parameter. | [optional] 
**Name** | **string** | (sortable) The rule set&#39;s name. | 
**SuccessCriteria** | Pointer to [**SuccessCriteria**](SuccessCriteria.md) |  | [optional] 
**ElementType** | Pointer to [**RuleSetElementType**](RuleSetElementType.md) |  | [optional] 
**Policy** | **[]int32** | The list of policy ids assigned to the rule set. | 

## Methods

### NewRuleSet

`func NewRuleSet(name string, policy []int32, ) *RuleSet`

NewRuleSet instantiates a new RuleSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleSetWithDefaults

`func NewRuleSetWithDefaults() *RuleSet`

NewRuleSetWithDefaults instantiates a new RuleSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RuleSet) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RuleSet) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RuleSet) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RuleSet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RuleSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleSet) SetName(v string)`

SetName sets Name field to given value.


### GetSuccessCriteria

`func (o *RuleSet) GetSuccessCriteria() SuccessCriteria`

GetSuccessCriteria returns the SuccessCriteria field if non-nil, zero value otherwise.

### GetSuccessCriteriaOk

`func (o *RuleSet) GetSuccessCriteriaOk() (*SuccessCriteria, bool)`

GetSuccessCriteriaOk returns a tuple with the SuccessCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCriteria

`func (o *RuleSet) SetSuccessCriteria(v SuccessCriteria)`

SetSuccessCriteria sets SuccessCriteria field to given value.

### HasSuccessCriteria

`func (o *RuleSet) HasSuccessCriteria() bool`

HasSuccessCriteria returns a boolean if a field has been set.

### GetElementType

`func (o *RuleSet) GetElementType() RuleSetElementType`

GetElementType returns the ElementType field if non-nil, zero value otherwise.

### GetElementTypeOk

`func (o *RuleSet) GetElementTypeOk() (*RuleSetElementType, bool)`

GetElementTypeOk returns a tuple with the ElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementType

`func (o *RuleSet) SetElementType(v RuleSetElementType)`

SetElementType sets ElementType field to given value.

### HasElementType

`func (o *RuleSet) HasElementType() bool`

HasElementType returns a boolean if a field has been set.

### GetPolicy

`func (o *RuleSet) GetPolicy() []int32`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *RuleSet) GetPolicyOk() (*[]int32, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *RuleSet) SetPolicy(v []int32)`

SetPolicy sets Policy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


