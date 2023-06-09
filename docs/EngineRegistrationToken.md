# EngineRegistrationToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpProxyId** | Pointer to **int64** | The ID of the proxy to use for HTTP requests or zero if none. | [optional] 
**HttpsProxyId** | Pointer to **int64** | The ID of the proxy to use for HTTPS requests or zero if none. | [optional] 
**ExpirationSeconds** | **int64** | The number of seconds after which this token will expire and be unavailable for use to register engines. | 
**SelectedCertificateId** | Pointer to **int64** | The ID of the certificate the engine will use to contact PingAccess via SSL/TLS. | [optional] 

## Methods

### NewEngineRegistrationToken

`func NewEngineRegistrationToken(expirationSeconds int64, ) *EngineRegistrationToken`

NewEngineRegistrationToken instantiates a new EngineRegistrationToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineRegistrationTokenWithDefaults

`func NewEngineRegistrationTokenWithDefaults() *EngineRegistrationToken`

NewEngineRegistrationTokenWithDefaults instantiates a new EngineRegistrationToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpProxyId

`func (o *EngineRegistrationToken) GetHttpProxyId() int64`

GetHttpProxyId returns the HttpProxyId field if non-nil, zero value otherwise.

### GetHttpProxyIdOk

`func (o *EngineRegistrationToken) GetHttpProxyIdOk() (*int64, bool)`

GetHttpProxyIdOk returns a tuple with the HttpProxyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxyId

`func (o *EngineRegistrationToken) SetHttpProxyId(v int64)`

SetHttpProxyId sets HttpProxyId field to given value.

### HasHttpProxyId

`func (o *EngineRegistrationToken) HasHttpProxyId() bool`

HasHttpProxyId returns a boolean if a field has been set.

### GetHttpsProxyId

`func (o *EngineRegistrationToken) GetHttpsProxyId() int64`

GetHttpsProxyId returns the HttpsProxyId field if non-nil, zero value otherwise.

### GetHttpsProxyIdOk

`func (o *EngineRegistrationToken) GetHttpsProxyIdOk() (*int64, bool)`

GetHttpsProxyIdOk returns a tuple with the HttpsProxyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsProxyId

`func (o *EngineRegistrationToken) SetHttpsProxyId(v int64)`

SetHttpsProxyId sets HttpsProxyId field to given value.

### HasHttpsProxyId

`func (o *EngineRegistrationToken) HasHttpsProxyId() bool`

HasHttpsProxyId returns a boolean if a field has been set.

### GetExpirationSeconds

`func (o *EngineRegistrationToken) GetExpirationSeconds() int64`

GetExpirationSeconds returns the ExpirationSeconds field if non-nil, zero value otherwise.

### GetExpirationSecondsOk

`func (o *EngineRegistrationToken) GetExpirationSecondsOk() (*int64, bool)`

GetExpirationSecondsOk returns a tuple with the ExpirationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationSeconds

`func (o *EngineRegistrationToken) SetExpirationSeconds(v int64)`

SetExpirationSeconds sets ExpirationSeconds field to given value.


### GetSelectedCertificateId

`func (o *EngineRegistrationToken) GetSelectedCertificateId() int64`

GetSelectedCertificateId returns the SelectedCertificateId field if non-nil, zero value otherwise.

### GetSelectedCertificateIdOk

`func (o *EngineRegistrationToken) GetSelectedCertificateIdOk() (*int64, bool)`

GetSelectedCertificateIdOk returns a tuple with the SelectedCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedCertificateId

`func (o *EngineRegistrationToken) SetSelectedCertificateId(v int64)`

SetSelectedCertificateId sets SelectedCertificateId field to given value.

### HasSelectedCertificateId

`func (o *EngineRegistrationToken) HasSelectedCertificateId() bool`

HasSelectedCertificateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


