# PingFederateRuntime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Targets** | Pointer to **[]string** | One or more server hostname:port pairs used to access the PingFederate server from inside a private network. Required when the PingFederate application is configured. | [optional] 
**SkipHostnameVerification** | Pointer to **bool** | Enable if the Back Channel servers should not perform hostname verification of the certificate. | [optional] 
**ExpectedHostname** | Pointer to **string** | The name of the host expected in the certificate. | [optional] 
**BackChannelBasePath** | Pointer to **string** | The base path, if needed, for the PingFederate Runtime accessed using a Back Channel hostname. This field is ignored when the PingFederate application is configured. | [optional] 
**BackChannelSecure** | Pointer to **bool** | Enable if PingFederate is expecting HTTPS connections for calls via the Back Channel hostnames. | [optional] 
**UseSlo** | Pointer to **bool** | Enable if OIDC single log out should be used on the /pa/oidc/logout on the engines. | [optional] 
**Application** | Pointer to [**PingFederateRuntimeApplication**](PingFederateRuntimeApplication.md) |  | [optional] 
**AvailabilityProfileId** | Pointer to **int32** | The ID of the availability profile to use for the PingFederate runtime. When set to 0, an availability profile defined by the pa.default.availability.ondemand properties in run.properties will be used for back end communication to PingFederate. | [optional] 
**LoadBalancingStrategyId** | Pointer to **int32** | The ID of the load balancing strategy to use for requests to the PingFederate targets. | [optional] 
**TrustedCertificateGroupId** | Pointer to **int32** | The group of certificates to use when authenticating to PingFederate. | [optional] 
**Secure** | Pointer to **bool** | Enable if PingFederate is expecting HTTPS connections. This field is ignored when the PingFederate application is configured. In this case, use backChannelSecure instead. | [optional] 
**BasePath** | Pointer to **string** | The base path, if needed, for PingFederate Runtime. This field is ignored when the PingFederate application is configured. | [optional] 
**UseProxy** | Pointer to **bool** | True if a proxy should be used for HTTP or HTTPS requests. | [optional] 
**Host** | **string** | The host name or IP address for PingFederate Runtime. This field is ignored and can be an empty string when the PingFederate application is configured. | 
**Port** | **int32** | The port number for PingFederate Runtime. This field is ignored when the PingFederate application is configured. | 
**AuditLevel** | Pointer to **string** | [&#39;ON&#39; or &#39;OFF&#39;]: Enable to record requests to PingFederate to the audit store. | [optional] 

## Methods

### NewPingFederateRuntime

`func NewPingFederateRuntime(host string, port int32, ) *PingFederateRuntime`

NewPingFederateRuntime instantiates a new PingFederateRuntime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPingFederateRuntimeWithDefaults

`func NewPingFederateRuntimeWithDefaults() *PingFederateRuntime`

NewPingFederateRuntimeWithDefaults instantiates a new PingFederateRuntime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargets

`func (o *PingFederateRuntime) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *PingFederateRuntime) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *PingFederateRuntime) SetTargets(v []string)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *PingFederateRuntime) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetSkipHostnameVerification

`func (o *PingFederateRuntime) GetSkipHostnameVerification() bool`

GetSkipHostnameVerification returns the SkipHostnameVerification field if non-nil, zero value otherwise.

### GetSkipHostnameVerificationOk

`func (o *PingFederateRuntime) GetSkipHostnameVerificationOk() (*bool, bool)`

GetSkipHostnameVerificationOk returns a tuple with the SkipHostnameVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipHostnameVerification

`func (o *PingFederateRuntime) SetSkipHostnameVerification(v bool)`

SetSkipHostnameVerification sets SkipHostnameVerification field to given value.

### HasSkipHostnameVerification

`func (o *PingFederateRuntime) HasSkipHostnameVerification() bool`

HasSkipHostnameVerification returns a boolean if a field has been set.

### GetExpectedHostname

`func (o *PingFederateRuntime) GetExpectedHostname() string`

GetExpectedHostname returns the ExpectedHostname field if non-nil, zero value otherwise.

### GetExpectedHostnameOk

`func (o *PingFederateRuntime) GetExpectedHostnameOk() (*string, bool)`

GetExpectedHostnameOk returns a tuple with the ExpectedHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedHostname

`func (o *PingFederateRuntime) SetExpectedHostname(v string)`

SetExpectedHostname sets ExpectedHostname field to given value.

### HasExpectedHostname

`func (o *PingFederateRuntime) HasExpectedHostname() bool`

HasExpectedHostname returns a boolean if a field has been set.

### GetBackChannelBasePath

`func (o *PingFederateRuntime) GetBackChannelBasePath() string`

GetBackChannelBasePath returns the BackChannelBasePath field if non-nil, zero value otherwise.

### GetBackChannelBasePathOk

`func (o *PingFederateRuntime) GetBackChannelBasePathOk() (*string, bool)`

GetBackChannelBasePathOk returns a tuple with the BackChannelBasePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackChannelBasePath

`func (o *PingFederateRuntime) SetBackChannelBasePath(v string)`

SetBackChannelBasePath sets BackChannelBasePath field to given value.

### HasBackChannelBasePath

`func (o *PingFederateRuntime) HasBackChannelBasePath() bool`

HasBackChannelBasePath returns a boolean if a field has been set.

### GetBackChannelSecure

`func (o *PingFederateRuntime) GetBackChannelSecure() bool`

GetBackChannelSecure returns the BackChannelSecure field if non-nil, zero value otherwise.

### GetBackChannelSecureOk

`func (o *PingFederateRuntime) GetBackChannelSecureOk() (*bool, bool)`

GetBackChannelSecureOk returns a tuple with the BackChannelSecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackChannelSecure

`func (o *PingFederateRuntime) SetBackChannelSecure(v bool)`

SetBackChannelSecure sets BackChannelSecure field to given value.

### HasBackChannelSecure

`func (o *PingFederateRuntime) HasBackChannelSecure() bool`

HasBackChannelSecure returns a boolean if a field has been set.

### GetUseSlo

`func (o *PingFederateRuntime) GetUseSlo() bool`

GetUseSlo returns the UseSlo field if non-nil, zero value otherwise.

### GetUseSloOk

`func (o *PingFederateRuntime) GetUseSloOk() (*bool, bool)`

GetUseSloOk returns a tuple with the UseSlo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSlo

`func (o *PingFederateRuntime) SetUseSlo(v bool)`

SetUseSlo sets UseSlo field to given value.

### HasUseSlo

`func (o *PingFederateRuntime) HasUseSlo() bool`

HasUseSlo returns a boolean if a field has been set.

### GetApplication

`func (o *PingFederateRuntime) GetApplication() PingFederateRuntimeApplication`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *PingFederateRuntime) GetApplicationOk() (*PingFederateRuntimeApplication, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *PingFederateRuntime) SetApplication(v PingFederateRuntimeApplication)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *PingFederateRuntime) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetAvailabilityProfileId

`func (o *PingFederateRuntime) GetAvailabilityProfileId() int32`

GetAvailabilityProfileId returns the AvailabilityProfileId field if non-nil, zero value otherwise.

### GetAvailabilityProfileIdOk

`func (o *PingFederateRuntime) GetAvailabilityProfileIdOk() (*int32, bool)`

GetAvailabilityProfileIdOk returns a tuple with the AvailabilityProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityProfileId

`func (o *PingFederateRuntime) SetAvailabilityProfileId(v int32)`

SetAvailabilityProfileId sets AvailabilityProfileId field to given value.

### HasAvailabilityProfileId

`func (o *PingFederateRuntime) HasAvailabilityProfileId() bool`

HasAvailabilityProfileId returns a boolean if a field has been set.

### GetLoadBalancingStrategyId

`func (o *PingFederateRuntime) GetLoadBalancingStrategyId() int32`

GetLoadBalancingStrategyId returns the LoadBalancingStrategyId field if non-nil, zero value otherwise.

### GetLoadBalancingStrategyIdOk

`func (o *PingFederateRuntime) GetLoadBalancingStrategyIdOk() (*int32, bool)`

GetLoadBalancingStrategyIdOk returns a tuple with the LoadBalancingStrategyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancingStrategyId

`func (o *PingFederateRuntime) SetLoadBalancingStrategyId(v int32)`

SetLoadBalancingStrategyId sets LoadBalancingStrategyId field to given value.

### HasLoadBalancingStrategyId

`func (o *PingFederateRuntime) HasLoadBalancingStrategyId() bool`

HasLoadBalancingStrategyId returns a boolean if a field has been set.

### GetTrustedCertificateGroupId

`func (o *PingFederateRuntime) GetTrustedCertificateGroupId() int32`

GetTrustedCertificateGroupId returns the TrustedCertificateGroupId field if non-nil, zero value otherwise.

### GetTrustedCertificateGroupIdOk

`func (o *PingFederateRuntime) GetTrustedCertificateGroupIdOk() (*int32, bool)`

GetTrustedCertificateGroupIdOk returns a tuple with the TrustedCertificateGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedCertificateGroupId

`func (o *PingFederateRuntime) SetTrustedCertificateGroupId(v int32)`

SetTrustedCertificateGroupId sets TrustedCertificateGroupId field to given value.

### HasTrustedCertificateGroupId

`func (o *PingFederateRuntime) HasTrustedCertificateGroupId() bool`

HasTrustedCertificateGroupId returns a boolean if a field has been set.

### GetSecure

`func (o *PingFederateRuntime) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *PingFederateRuntime) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *PingFederateRuntime) SetSecure(v bool)`

SetSecure sets Secure field to given value.

### HasSecure

`func (o *PingFederateRuntime) HasSecure() bool`

HasSecure returns a boolean if a field has been set.

### GetBasePath

`func (o *PingFederateRuntime) GetBasePath() string`

GetBasePath returns the BasePath field if non-nil, zero value otherwise.

### GetBasePathOk

`func (o *PingFederateRuntime) GetBasePathOk() (*string, bool)`

GetBasePathOk returns a tuple with the BasePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePath

`func (o *PingFederateRuntime) SetBasePath(v string)`

SetBasePath sets BasePath field to given value.

### HasBasePath

`func (o *PingFederateRuntime) HasBasePath() bool`

HasBasePath returns a boolean if a field has been set.

### GetUseProxy

`func (o *PingFederateRuntime) GetUseProxy() bool`

GetUseProxy returns the UseProxy field if non-nil, zero value otherwise.

### GetUseProxyOk

`func (o *PingFederateRuntime) GetUseProxyOk() (*bool, bool)`

GetUseProxyOk returns a tuple with the UseProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseProxy

`func (o *PingFederateRuntime) SetUseProxy(v bool)`

SetUseProxy sets UseProxy field to given value.

### HasUseProxy

`func (o *PingFederateRuntime) HasUseProxy() bool`

HasUseProxy returns a boolean if a field has been set.

### GetHost

`func (o *PingFederateRuntime) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *PingFederateRuntime) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *PingFederateRuntime) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *PingFederateRuntime) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PingFederateRuntime) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PingFederateRuntime) SetPort(v int32)`

SetPort sets Port field to given value.


### GetAuditLevel

`func (o *PingFederateRuntime) GetAuditLevel() string`

GetAuditLevel returns the AuditLevel field if non-nil, zero value otherwise.

### GetAuditLevelOk

`func (o *PingFederateRuntime) GetAuditLevelOk() (*string, bool)`

GetAuditLevelOk returns a tuple with the AuditLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLevel

`func (o *PingFederateRuntime) SetAuditLevel(v string)`

SetAuditLevel sets AuditLevel field to given value.

### HasAuditLevel

`func (o *PingFederateRuntime) HasAuditLevel() bool`

HasAuditLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


