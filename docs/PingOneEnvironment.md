# PingOneEnvironment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | When creating a new PingOneEnvironment, this is the ID for the PingOneEnvironment. If not specified, an ID will be automatically assigned. When updating an existing PingOneEnvironment, this field is ignored and the ID is determined by the path parameter. | 
**Name** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewPingOneEnvironment

`func NewPingOneEnvironment(id string, name string, type_ string, ) *PingOneEnvironment`

NewPingOneEnvironment instantiates a new PingOneEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPingOneEnvironmentWithDefaults

`func NewPingOneEnvironmentWithDefaults() *PingOneEnvironment`

NewPingOneEnvironmentWithDefaults instantiates a new PingOneEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PingOneEnvironment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PingOneEnvironment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PingOneEnvironment) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PingOneEnvironment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PingOneEnvironment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PingOneEnvironment) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *PingOneEnvironment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PingOneEnvironment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PingOneEnvironment) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


