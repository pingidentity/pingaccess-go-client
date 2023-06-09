# TargetHostPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | The hostname. | 
**Port** | **int64** | The port number. | 
**Secure** | **bool** | True if the target host expects HTTPS connections. | 

## Methods

### NewTargetHostPort

`func NewTargetHostPort(host string, port int64, secure bool, ) *TargetHostPort`

NewTargetHostPort instantiates a new TargetHostPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetHostPortWithDefaults

`func NewTargetHostPortWithDefaults() *TargetHostPort`

NewTargetHostPortWithDefaults instantiates a new TargetHostPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *TargetHostPort) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *TargetHostPort) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *TargetHostPort) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *TargetHostPort) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TargetHostPort) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TargetHostPort) SetPort(v int64)`

SetPort sets Port field to given value.


### GetSecure

`func (o *TargetHostPort) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *TargetHostPort) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *TargetHostPort) SetSecure(v bool)`

SetSecure sets Secure field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


