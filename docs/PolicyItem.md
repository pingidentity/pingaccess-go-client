# PolicyItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**PolicyItemType**](PolicyItemType.md) |  | 
**Id** | **int64** | The ID of the rule or rule set. | 

## Methods

### NewPolicyItem

`func NewPolicyItem(type_ PolicyItemType, id int64, ) *PolicyItem`

NewPolicyItem instantiates a new PolicyItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyItemWithDefaults

`func NewPolicyItemWithDefaults() *PolicyItem`

NewPolicyItemWithDefaults instantiates a new PolicyItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PolicyItem) GetType() PolicyItemType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PolicyItem) GetTypeOk() (*PolicyItemType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PolicyItem) SetType(v PolicyItemType)`

SetType sets Type field to given value.


### GetId

`func (o *PolicyItem) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyItem) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyItem) SetId(v int64)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


