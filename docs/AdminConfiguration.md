# AdminConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostPort** | **string** | The primary host and port of the administrative console. | 
**HttpProxyId** | Pointer to **int64** | The ID of the proxy to use for HTTP requests or zero if none. | [optional] 
**HttpsProxyId** | Pointer to **int64** | The ID of the proxy to use for HTTPS requests or zero if none. | [optional] 
**StaleEngineTimeout** | Pointer to **int64** | The number of minutes since the last communication from an engine after which the engine will be deleted. Misconfiguration of this value when set less than Engine Polling Delay or a loss of connection in the admin node may result in errant engine deletions. | [optional] 
**StaleEngineTimeoutEnabled** | Pointer to **bool** | Indicates whether the Stale Engine Timeout is enabled and automatic engine pruning will occur. | [optional] 

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

`func (o *AdminConfiguration) GetHttpProxyId() int64`

GetHttpProxyId returns the HttpProxyId field if non-nil, zero value otherwise.

### GetHttpProxyIdOk

`func (o *AdminConfiguration) GetHttpProxyIdOk() (*int64, bool)`

GetHttpProxyIdOk returns a tuple with the HttpProxyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxyId

`func (o *AdminConfiguration) SetHttpProxyId(v int64)`

SetHttpProxyId sets HttpProxyId field to given value.

### HasHttpProxyId

`func (o *AdminConfiguration) HasHttpProxyId() bool`

HasHttpProxyId returns a boolean if a field has been set.

### GetHttpsProxyId

`func (o *AdminConfiguration) GetHttpsProxyId() int64`

GetHttpsProxyId returns the HttpsProxyId field if non-nil, zero value otherwise.

### GetHttpsProxyIdOk

`func (o *AdminConfiguration) GetHttpsProxyIdOk() (*int64, bool)`

GetHttpsProxyIdOk returns a tuple with the HttpsProxyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsProxyId

`func (o *AdminConfiguration) SetHttpsProxyId(v int64)`

SetHttpsProxyId sets HttpsProxyId field to given value.

### HasHttpsProxyId

`func (o *AdminConfiguration) HasHttpsProxyId() bool`

HasHttpsProxyId returns a boolean if a field has been set.

### GetStaleEngineTimeout

`func (o *AdminConfiguration) GetStaleEngineTimeout() int64`

GetStaleEngineTimeout returns the StaleEngineTimeout field if non-nil, zero value otherwise.

### GetStaleEngineTimeoutOk

`func (o *AdminConfiguration) GetStaleEngineTimeoutOk() (*int64, bool)`

GetStaleEngineTimeoutOk returns a tuple with the StaleEngineTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleEngineTimeout

`func (o *AdminConfiguration) SetStaleEngineTimeout(v int64)`

SetStaleEngineTimeout sets StaleEngineTimeout field to given value.

### HasStaleEngineTimeout

`func (o *AdminConfiguration) HasStaleEngineTimeout() bool`

HasStaleEngineTimeout returns a boolean if a field has been set.

### GetStaleEngineTimeoutEnabled

`func (o *AdminConfiguration) GetStaleEngineTimeoutEnabled() bool`

GetStaleEngineTimeoutEnabled returns the StaleEngineTimeoutEnabled field if non-nil, zero value otherwise.

### GetStaleEngineTimeoutEnabledOk

`func (o *AdminConfiguration) GetStaleEngineTimeoutEnabledOk() (*bool, bool)`

GetStaleEngineTimeoutEnabledOk returns a tuple with the StaleEngineTimeoutEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleEngineTimeoutEnabled

`func (o *AdminConfiguration) SetStaleEngineTimeoutEnabled(v bool)`

SetStaleEngineTimeoutEnabled sets StaleEngineTimeoutEnabled field to given value.

### HasStaleEngineTimeoutEnabled

`func (o *AdminConfiguration) HasStaleEngineTimeoutEnabled() bool`

HasStaleEngineTimeoutEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


