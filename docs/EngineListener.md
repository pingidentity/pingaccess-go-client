# EngineListener

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | When creating a new EngineListener, this is the ID for the EngineListener. If not specified, an ID will be automatically assigned. When updating an existing EngineListener, this field is ignored and the ID is determined by the path parameter. | [optional] 
**Name** | **string** | (sortable) The name of the engine listener. | 
**Port** | **int64** | (sortable) The port the engine listener listens on. | 
**Secure** | Pointer to **bool** | (sortable) Indicator if the engine listener should listen to HTTPS connections. | [optional] 
**TrustedCertificateGroupId** | Pointer to **int64** | Trusted Certificate Group assigned to engine listener for client certificate authentication. | [optional] 

## Methods

### NewEngineListener

`func NewEngineListener(name string, port int64, ) *EngineListener`

NewEngineListener instantiates a new EngineListener object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineListenerWithDefaults

`func NewEngineListenerWithDefaults() *EngineListener`

NewEngineListenerWithDefaults instantiates a new EngineListener object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EngineListener) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EngineListener) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EngineListener) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *EngineListener) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EngineListener) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EngineListener) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EngineListener) SetName(v string)`

SetName sets Name field to given value.


### GetPort

`func (o *EngineListener) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *EngineListener) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *EngineListener) SetPort(v int64)`

SetPort sets Port field to given value.


### GetSecure

`func (o *EngineListener) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *EngineListener) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *EngineListener) SetSecure(v bool)`

SetSecure sets Secure field to given value.

### HasSecure

`func (o *EngineListener) HasSecure() bool`

HasSecure returns a boolean if a field has been set.

### GetTrustedCertificateGroupId

`func (o *EngineListener) GetTrustedCertificateGroupId() int64`

GetTrustedCertificateGroupId returns the TrustedCertificateGroupId field if non-nil, zero value otherwise.

### GetTrustedCertificateGroupIdOk

`func (o *EngineListener) GetTrustedCertificateGroupIdOk() (*int64, bool)`

GetTrustedCertificateGroupIdOk returns a tuple with the TrustedCertificateGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedCertificateGroupId

`func (o *EngineListener) SetTrustedCertificateGroupId(v int64)`

SetTrustedCertificateGroupId sets TrustedCertificateGroupId field to given value.

### HasTrustedCertificateGroupId

`func (o *EngineListener) HasTrustedCertificateGroupId() bool`

HasTrustedCertificateGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


