# PingFederateAdmin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationType** | Pointer to [**AuthenticationType**](AuthenticationType.md) |  | [optional] 
**AdminUsername** | Pointer to **string** | The administrator username. Required when the authentication type is set to &#39;Basic&#39;. | [optional] 
**AdminPassword** | [**HiddenField**](HiddenField.md) |  | 
**SkipHostnameVerification** | Pointer to **bool** | Set to true if HTTP communications to PingFederate should not perform hostname verification of the certificate. | [optional] 
**ExpectedHostname** | Pointer to **string** | The name of the host expected in the certificate used by PingFederate. | [optional] 
**TrustedCertificateGroupId** | Pointer to **int64** | The group of certificates to use when authenticating to PingFederate Administrative API. | [optional] 
**Host** | **string** | The host name or IP address for PingFederate Administration API. | 
**Port** | **int64** | The port number for PingFederate Administration API. | 
**Secure** | Pointer to **bool** | Enable if PingFederate is expecting HTTPS connections. | [optional] 
**BasePath** | Pointer to **string** | The base path, if needed, for Administration API. | [optional] 
**UseProxy** | Pointer to **bool** | True if a proxy should be used for HTTP or HTTPS requests. | [optional] 
**AuditLevel** | Pointer to **string** | [&#39;ON&#39; or &#39;OFF&#39;]: Enable to record requests to the PingFederate Administrative API to the audit store. | [optional] 
**OAuthAuthenticationConfig** | Pointer to [**OAuthAuthenticationConfiguration**](OAuthAuthenticationConfiguration.md) |  | [optional] 

## Methods

### NewPingFederateAdmin

`func NewPingFederateAdmin(adminPassword HiddenField, host string, port int64, ) *PingFederateAdmin`

NewPingFederateAdmin instantiates a new PingFederateAdmin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPingFederateAdminWithDefaults

`func NewPingFederateAdminWithDefaults() *PingFederateAdmin`

NewPingFederateAdminWithDefaults instantiates a new PingFederateAdmin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationType

`func (o *PingFederateAdmin) GetAuthenticationType() AuthenticationType`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *PingFederateAdmin) GetAuthenticationTypeOk() (*AuthenticationType, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *PingFederateAdmin) SetAuthenticationType(v AuthenticationType)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *PingFederateAdmin) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### GetAdminUsername

`func (o *PingFederateAdmin) GetAdminUsername() string`

GetAdminUsername returns the AdminUsername field if non-nil, zero value otherwise.

### GetAdminUsernameOk

`func (o *PingFederateAdmin) GetAdminUsernameOk() (*string, bool)`

GetAdminUsernameOk returns a tuple with the AdminUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminUsername

`func (o *PingFederateAdmin) SetAdminUsername(v string)`

SetAdminUsername sets AdminUsername field to given value.

### HasAdminUsername

`func (o *PingFederateAdmin) HasAdminUsername() bool`

HasAdminUsername returns a boolean if a field has been set.

### GetAdminPassword

`func (o *PingFederateAdmin) GetAdminPassword() HiddenField`

GetAdminPassword returns the AdminPassword field if non-nil, zero value otherwise.

### GetAdminPasswordOk

`func (o *PingFederateAdmin) GetAdminPasswordOk() (*HiddenField, bool)`

GetAdminPasswordOk returns a tuple with the AdminPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPassword

`func (o *PingFederateAdmin) SetAdminPassword(v HiddenField)`

SetAdminPassword sets AdminPassword field to given value.


### GetSkipHostnameVerification

`func (o *PingFederateAdmin) GetSkipHostnameVerification() bool`

GetSkipHostnameVerification returns the SkipHostnameVerification field if non-nil, zero value otherwise.

### GetSkipHostnameVerificationOk

`func (o *PingFederateAdmin) GetSkipHostnameVerificationOk() (*bool, bool)`

GetSkipHostnameVerificationOk returns a tuple with the SkipHostnameVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipHostnameVerification

`func (o *PingFederateAdmin) SetSkipHostnameVerification(v bool)`

SetSkipHostnameVerification sets SkipHostnameVerification field to given value.

### HasSkipHostnameVerification

`func (o *PingFederateAdmin) HasSkipHostnameVerification() bool`

HasSkipHostnameVerification returns a boolean if a field has been set.

### GetExpectedHostname

`func (o *PingFederateAdmin) GetExpectedHostname() string`

GetExpectedHostname returns the ExpectedHostname field if non-nil, zero value otherwise.

### GetExpectedHostnameOk

`func (o *PingFederateAdmin) GetExpectedHostnameOk() (*string, bool)`

GetExpectedHostnameOk returns a tuple with the ExpectedHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedHostname

`func (o *PingFederateAdmin) SetExpectedHostname(v string)`

SetExpectedHostname sets ExpectedHostname field to given value.

### HasExpectedHostname

`func (o *PingFederateAdmin) HasExpectedHostname() bool`

HasExpectedHostname returns a boolean if a field has been set.

### GetTrustedCertificateGroupId

`func (o *PingFederateAdmin) GetTrustedCertificateGroupId() int64`

GetTrustedCertificateGroupId returns the TrustedCertificateGroupId field if non-nil, zero value otherwise.

### GetTrustedCertificateGroupIdOk

`func (o *PingFederateAdmin) GetTrustedCertificateGroupIdOk() (*int64, bool)`

GetTrustedCertificateGroupIdOk returns a tuple with the TrustedCertificateGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedCertificateGroupId

`func (o *PingFederateAdmin) SetTrustedCertificateGroupId(v int64)`

SetTrustedCertificateGroupId sets TrustedCertificateGroupId field to given value.

### HasTrustedCertificateGroupId

`func (o *PingFederateAdmin) HasTrustedCertificateGroupId() bool`

HasTrustedCertificateGroupId returns a boolean if a field has been set.

### GetHost

`func (o *PingFederateAdmin) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *PingFederateAdmin) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *PingFederateAdmin) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *PingFederateAdmin) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PingFederateAdmin) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PingFederateAdmin) SetPort(v int64)`

SetPort sets Port field to given value.


### GetSecure

`func (o *PingFederateAdmin) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *PingFederateAdmin) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *PingFederateAdmin) SetSecure(v bool)`

SetSecure sets Secure field to given value.

### HasSecure

`func (o *PingFederateAdmin) HasSecure() bool`

HasSecure returns a boolean if a field has been set.

### GetBasePath

`func (o *PingFederateAdmin) GetBasePath() string`

GetBasePath returns the BasePath field if non-nil, zero value otherwise.

### GetBasePathOk

`func (o *PingFederateAdmin) GetBasePathOk() (*string, bool)`

GetBasePathOk returns a tuple with the BasePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePath

`func (o *PingFederateAdmin) SetBasePath(v string)`

SetBasePath sets BasePath field to given value.

### HasBasePath

`func (o *PingFederateAdmin) HasBasePath() bool`

HasBasePath returns a boolean if a field has been set.

### GetUseProxy

`func (o *PingFederateAdmin) GetUseProxy() bool`

GetUseProxy returns the UseProxy field if non-nil, zero value otherwise.

### GetUseProxyOk

`func (o *PingFederateAdmin) GetUseProxyOk() (*bool, bool)`

GetUseProxyOk returns a tuple with the UseProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseProxy

`func (o *PingFederateAdmin) SetUseProxy(v bool)`

SetUseProxy sets UseProxy field to given value.

### HasUseProxy

`func (o *PingFederateAdmin) HasUseProxy() bool`

HasUseProxy returns a boolean if a field has been set.

### GetAuditLevel

`func (o *PingFederateAdmin) GetAuditLevel() string`

GetAuditLevel returns the AuditLevel field if non-nil, zero value otherwise.

### GetAuditLevelOk

`func (o *PingFederateAdmin) GetAuditLevelOk() (*string, bool)`

GetAuditLevelOk returns a tuple with the AuditLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLevel

`func (o *PingFederateAdmin) SetAuditLevel(v string)`

SetAuditLevel sets AuditLevel field to given value.

### HasAuditLevel

`func (o *PingFederateAdmin) HasAuditLevel() bool`

HasAuditLevel returns a boolean if a field has been set.

### GetOAuthAuthenticationConfig

`func (o *PingFederateAdmin) GetOAuthAuthenticationConfig() OAuthAuthenticationConfiguration`

GetOAuthAuthenticationConfig returns the OAuthAuthenticationConfig field if non-nil, zero value otherwise.

### GetOAuthAuthenticationConfigOk

`func (o *PingFederateAdmin) GetOAuthAuthenticationConfigOk() (*OAuthAuthenticationConfiguration, bool)`

GetOAuthAuthenticationConfigOk returns a tuple with the OAuthAuthenticationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuthAuthenticationConfig

`func (o *PingFederateAdmin) SetOAuthAuthenticationConfig(v OAuthAuthenticationConfiguration)`

SetOAuthAuthenticationConfig sets OAuthAuthenticationConfig field to given value.

### HasOAuthAuthenticationConfig

`func (o *PingFederateAdmin) HasOAuthAuthenticationConfig() bool`

HasOAuthAuthenticationConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


