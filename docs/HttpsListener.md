# HttpsListener

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | (sortable) The ID of the HTTPS listener. | [optional] 
**Name** | **string** | (sortable) The name of the HTTPS listener. | 
**KeyPairId** | **int32** | The ID of the default key pair used by the HTTPS listener. | 
**UseServerCipherSuiteOrder** | **bool** | (sortable) Enable server cipher suite ordering for the HTTPS listener. | 
**RestartRequired** | **bool** | Indicates whether an admin or engine restart is required to update the HTTPS listener. Cannot be True for the ENGINE listener. | 

## Methods

### NewHttpsListener

`func NewHttpsListener(name string, keyPairId int32, useServerCipherSuiteOrder bool, restartRequired bool, ) *HttpsListener`

NewHttpsListener instantiates a new HttpsListener object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpsListenerWithDefaults

`func NewHttpsListenerWithDefaults() *HttpsListener`

NewHttpsListenerWithDefaults instantiates a new HttpsListener object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HttpsListener) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HttpsListener) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HttpsListener) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *HttpsListener) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *HttpsListener) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HttpsListener) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HttpsListener) SetName(v string)`

SetName sets Name field to given value.


### GetKeyPairId

`func (o *HttpsListener) GetKeyPairId() int32`

GetKeyPairId returns the KeyPairId field if non-nil, zero value otherwise.

### GetKeyPairIdOk

`func (o *HttpsListener) GetKeyPairIdOk() (*int32, bool)`

GetKeyPairIdOk returns a tuple with the KeyPairId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPairId

`func (o *HttpsListener) SetKeyPairId(v int32)`

SetKeyPairId sets KeyPairId field to given value.


### GetUseServerCipherSuiteOrder

`func (o *HttpsListener) GetUseServerCipherSuiteOrder() bool`

GetUseServerCipherSuiteOrder returns the UseServerCipherSuiteOrder field if non-nil, zero value otherwise.

### GetUseServerCipherSuiteOrderOk

`func (o *HttpsListener) GetUseServerCipherSuiteOrderOk() (*bool, bool)`

GetUseServerCipherSuiteOrderOk returns a tuple with the UseServerCipherSuiteOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseServerCipherSuiteOrder

`func (o *HttpsListener) SetUseServerCipherSuiteOrder(v bool)`

SetUseServerCipherSuiteOrder sets UseServerCipherSuiteOrder field to given value.


### GetRestartRequired

`func (o *HttpsListener) GetRestartRequired() bool`

GetRestartRequired returns the RestartRequired field if non-nil, zero value otherwise.

### GetRestartRequiredOk

`func (o *HttpsListener) GetRestartRequiredOk() (*bool, bool)`

GetRestartRequiredOk returns a tuple with the RestartRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartRequired

`func (o *HttpsListener) SetRestartRequired(v bool)`

SetRestartRequired sets RestartRequired field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


