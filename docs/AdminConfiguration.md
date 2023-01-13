# AdminConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostPort** | **string** | The primary host and port of the administrative console. | 
**HttpProxyId** | Pointer to **int32** | The ID of the proxy to use for HTTP requests or zero if none. | [optional] 
**HttpsProxyId** | Pointer to **int32** | The ID of the proxy to use for HTTPS requests or zero if none. | [optional] 

## Methods

### NewAdminConfiguration

`func NewAdminConfiguration(hostPort string, ) *AdminConfiguration`

NewAdminConfiguration instantiates a new AdminConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminConfigurationWithDefaults

`func NewAdminConfigurationWithDefaults() *AdminConfiguration`

NewAdminConfigurationWithDefaults instantiates a new AdminConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostPort

`func (o *AdminConfiguration) GetHostPort() string`

GetHostPort returns the HostPort field if non-nil, zero value otherwise.

### GetHostPortOk

`func (o *AdminConfiguration) GetHostPortOk() (*string, bool)`

GetHostPortOk returns a tuple with the HostPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostPort

`func (o *AdminConfiguration) SetHostPort(v string)`

SetHostPort sets HostPort field to given value.


### GetHttpProxyId

`func (o *AdminConfiguration) GetHttpProxyId() int32`

GetHttpProxyId returns the HttpProxyId field if non-nil, zero value otherwise.

### GetHttpProxyIdOk

`func (o *AdminConfiguration) GetHttpProxyIdOk() (*int32, bool)`

GetHttpProxyIdOk returns a tuple with the HttpProxyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxyId

`func (o *AdminConfiguration) SetHttpProxyId(v int32)`

SetHttpProxyId sets HttpProxyId field to given value.

### HasHttpProxyId

`func (o *AdminConfiguration) HasHttpProxyId() bool`

HasHttpProxyId returns a boolean if a field has been set.

### GetHttpsProxyId

`func (o *AdminConfiguration) GetHttpsProxyId() int32`

GetHttpsProxyId returns the HttpsProxyId field if non-nil, zero value otherwise.

### GetHttpsProxyIdOk

`func (o *AdminConfiguration) GetHttpsProxyIdOk() (*int32, bool)`

GetHttpsProxyIdOk returns a tuple with the HttpsProxyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsProxyId

`func (o *AdminConfiguration) SetHttpsProxyId(v int32)`

SetHttpsProxyId sets HttpsProxyId field to given value.

### HasHttpsProxyId

`func (o *AdminConfiguration) HasHttpsProxyId() bool`

HasHttpsProxyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


