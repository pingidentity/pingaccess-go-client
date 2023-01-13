# Site

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | When creating a new Site, this is the ID for the Site. If not specified, an ID will be automatically assigned. When updating an existing Site, this field is ignored and the ID is determined by the path parameter. | [optional] 
**Name** | **string** | (sortable) The name of the site. | 
**Targets** | **[]string** | The {hostname}:{port} pairs for the hosts that make up the site. | 
**Secure** | Pointer to **bool** | (sortable) This field is true if the site expects HTTPS connections. | [optional] 
**TrustedCertificateGroupId** | Pointer to **int32** | The ID of the trusted certificate group associated with the site. | [optional] 
**SendPaCookie** | Pointer to **bool** | (sortable) This field is true if the PingAccess Token or OAuth Access Token should be included in the request to the site. | [optional] 
**UseTargetHostHeader** | Pointer to **bool** | (sortable) Setting this field to true causes PingAccess to adjust the Host header to the site&#39;s selected target host rather than the virtual host configured in the application. | [optional] 
**KeepAliveTimeout** | Pointer to **int32** | (sortable) The time, in milliseconds, that an HTTP persistent connection to the site can be idle before PingAccess closes the connection. | [optional] 
**MaxConnections** | Pointer to **int32** | (sortable) The maximum number of HTTP persistent connections you want PingAccess to have open and maintain for the site. -1 indicates unlimited connections. | [optional] 
**MaxWebSocketConnections** | Pointer to **int32** | (sortable) The maximum number of WebSocket connections you want PingAccess to have open and maintain for the site. -1 indicates unlimited connections. | [optional] 
**SiteAuthenticatorIds** | Pointer to **[]int32** | The IDs of the site authenticators associated with the site. | [optional] 
**SkipHostnameVerification** | Pointer to **bool** | (sortable) This field is true if the hostname verification of the site&#39;s certificate should be skipped. | [optional] 
**ExpectedHostname** | Pointer to **string** | (sortable) The name of the host expected in the site&#39;s certificate. | [optional] 
**AvailabilityProfileId** | Pointer to **int32** | The ID of the availability profile associated with the site. | [optional] 
**LoadBalancingStrategyId** | Pointer to **int32** | The ID of the load balancing strategy associated with the site. | [optional] 
**UseProxy** | Pointer to **bool** | (sortable) True if a proxy should be used for HTTP or HTTPS requests. | [optional] 

## Methods

### NewSite

`func NewSite(name string, targets []string, ) *Site`

NewSite instantiates a new Site object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteWithDefaults

`func NewSiteWithDefaults() *Site`

NewSiteWithDefaults instantiates a new Site object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Site) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Site) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Site) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Site) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Site) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Site) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Site) SetName(v string)`

SetName sets Name field to given value.


### GetTargets

`func (o *Site) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *Site) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *Site) SetTargets(v []string)`

SetTargets sets Targets field to given value.


### GetSecure

`func (o *Site) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *Site) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *Site) SetSecure(v bool)`

SetSecure sets Secure field to given value.

### HasSecure

`func (o *Site) HasSecure() bool`

HasSecure returns a boolean if a field has been set.

### GetTrustedCertificateGroupId

`func (o *Site) GetTrustedCertificateGroupId() int32`

GetTrustedCertificateGroupId returns the TrustedCertificateGroupId field if non-nil, zero value otherwise.

### GetTrustedCertificateGroupIdOk

`func (o *Site) GetTrustedCertificateGroupIdOk() (*int32, bool)`

GetTrustedCertificateGroupIdOk returns a tuple with the TrustedCertificateGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedCertificateGroupId

`func (o *Site) SetTrustedCertificateGroupId(v int32)`

SetTrustedCertificateGroupId sets TrustedCertificateGroupId field to given value.

### HasTrustedCertificateGroupId

`func (o *Site) HasTrustedCertificateGroupId() bool`

HasTrustedCertificateGroupId returns a boolean if a field has been set.

### GetSendPaCookie

`func (o *Site) GetSendPaCookie() bool`

GetSendPaCookie returns the SendPaCookie field if non-nil, zero value otherwise.

### GetSendPaCookieOk

`func (o *Site) GetSendPaCookieOk() (*bool, bool)`

GetSendPaCookieOk returns a tuple with the SendPaCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendPaCookie

`func (o *Site) SetSendPaCookie(v bool)`

SetSendPaCookie sets SendPaCookie field to given value.

### HasSendPaCookie

`func (o *Site) HasSendPaCookie() bool`

HasSendPaCookie returns a boolean if a field has been set.

### GetUseTargetHostHeader

`func (o *Site) GetUseTargetHostHeader() bool`

GetUseTargetHostHeader returns the UseTargetHostHeader field if non-nil, zero value otherwise.

### GetUseTargetHostHeaderOk

`func (o *Site) GetUseTargetHostHeaderOk() (*bool, bool)`

GetUseTargetHostHeaderOk returns a tuple with the UseTargetHostHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTargetHostHeader

`func (o *Site) SetUseTargetHostHeader(v bool)`

SetUseTargetHostHeader sets UseTargetHostHeader field to given value.

### HasUseTargetHostHeader

`func (o *Site) HasUseTargetHostHeader() bool`

HasUseTargetHostHeader returns a boolean if a field has been set.

### GetKeepAliveTimeout

`func (o *Site) GetKeepAliveTimeout() int32`

GetKeepAliveTimeout returns the KeepAliveTimeout field if non-nil, zero value otherwise.

### GetKeepAliveTimeoutOk

`func (o *Site) GetKeepAliveTimeoutOk() (*int32, bool)`

GetKeepAliveTimeoutOk returns a tuple with the KeepAliveTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepAliveTimeout

`func (o *Site) SetKeepAliveTimeout(v int32)`

SetKeepAliveTimeout sets KeepAliveTimeout field to given value.

### HasKeepAliveTimeout

`func (o *Site) HasKeepAliveTimeout() bool`

HasKeepAliveTimeout returns a boolean if a field has been set.

### GetMaxConnections

`func (o *Site) GetMaxConnections() int32`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *Site) GetMaxConnectionsOk() (*int32, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *Site) SetMaxConnections(v int32)`

SetMaxConnections sets MaxConnections field to given value.

### HasMaxConnections

`func (o *Site) HasMaxConnections() bool`

HasMaxConnections returns a boolean if a field has been set.

### GetMaxWebSocketConnections

`func (o *Site) GetMaxWebSocketConnections() int32`

GetMaxWebSocketConnections returns the MaxWebSocketConnections field if non-nil, zero value otherwise.

### GetMaxWebSocketConnectionsOk

`func (o *Site) GetMaxWebSocketConnectionsOk() (*int32, bool)`

GetMaxWebSocketConnectionsOk returns a tuple with the MaxWebSocketConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWebSocketConnections

`func (o *Site) SetMaxWebSocketConnections(v int32)`

SetMaxWebSocketConnections sets MaxWebSocketConnections field to given value.

### HasMaxWebSocketConnections

`func (o *Site) HasMaxWebSocketConnections() bool`

HasMaxWebSocketConnections returns a boolean if a field has been set.

### GetSiteAuthenticatorIds

`func (o *Site) GetSiteAuthenticatorIds() []int32`

GetSiteAuthenticatorIds returns the SiteAuthenticatorIds field if non-nil, zero value otherwise.

### GetSiteAuthenticatorIdsOk

`func (o *Site) GetSiteAuthenticatorIdsOk() (*[]int32, bool)`

GetSiteAuthenticatorIdsOk returns a tuple with the SiteAuthenticatorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteAuthenticatorIds

`func (o *Site) SetSiteAuthenticatorIds(v []int32)`

SetSiteAuthenticatorIds sets SiteAuthenticatorIds field to given value.

### HasSiteAuthenticatorIds

`func (o *Site) HasSiteAuthenticatorIds() bool`

HasSiteAuthenticatorIds returns a boolean if a field has been set.

### GetSkipHostnameVerification

`func (o *Site) GetSkipHostnameVerification() bool`

GetSkipHostnameVerification returns the SkipHostnameVerification field if non-nil, zero value otherwise.

### GetSkipHostnameVerificationOk

`func (o *Site) GetSkipHostnameVerificationOk() (*bool, bool)`

GetSkipHostnameVerificationOk returns a tuple with the SkipHostnameVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipHostnameVerification

`func (o *Site) SetSkipHostnameVerification(v bool)`

SetSkipHostnameVerification sets SkipHostnameVerification field to given value.

### HasSkipHostnameVerification

`func (o *Site) HasSkipHostnameVerification() bool`

HasSkipHostnameVerification returns a boolean if a field has been set.

### GetExpectedHostname

`func (o *Site) GetExpectedHostname() string`

GetExpectedHostname returns the ExpectedHostname field if non-nil, zero value otherwise.

### GetExpectedHostnameOk

`func (o *Site) GetExpectedHostnameOk() (*string, bool)`

GetExpectedHostnameOk returns a tuple with the ExpectedHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedHostname

`func (o *Site) SetExpectedHostname(v string)`

SetExpectedHostname sets ExpectedHostname field to given value.

### HasExpectedHostname

`func (o *Site) HasExpectedHostname() bool`

HasExpectedHostname returns a boolean if a field has been set.

### GetAvailabilityProfileId

`func (o *Site) GetAvailabilityProfileId() int32`

GetAvailabilityProfileId returns the AvailabilityProfileId field if non-nil, zero value otherwise.

### GetAvailabilityProfileIdOk

`func (o *Site) GetAvailabilityProfileIdOk() (*int32, bool)`

GetAvailabilityProfileIdOk returns a tuple with the AvailabilityProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityProfileId

`func (o *Site) SetAvailabilityProfileId(v int32)`

SetAvailabilityProfileId sets AvailabilityProfileId field to given value.

### HasAvailabilityProfileId

`func (o *Site) HasAvailabilityProfileId() bool`

HasAvailabilityProfileId returns a boolean if a field has been set.

### GetLoadBalancingStrategyId

`func (o *Site) GetLoadBalancingStrategyId() int32`

GetLoadBalancingStrategyId returns the LoadBalancingStrategyId field if non-nil, zero value otherwise.

### GetLoadBalancingStrategyIdOk

`func (o *Site) GetLoadBalancingStrategyIdOk() (*int32, bool)`

GetLoadBalancingStrategyIdOk returns a tuple with the LoadBalancingStrategyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancingStrategyId

`func (o *Site) SetLoadBalancingStrategyId(v int32)`

SetLoadBalancingStrategyId sets LoadBalancingStrategyId field to given value.

### HasLoadBalancingStrategyId

`func (o *Site) HasLoadBalancingStrategyId() bool`

HasLoadBalancingStrategyId returns a boolean if a field has been set.

### GetUseProxy

`func (o *Site) GetUseProxy() bool`

GetUseProxy returns the UseProxy field if non-nil, zero value otherwise.

### GetUseProxyOk

`func (o *Site) GetUseProxyOk() (*bool, bool)`

GetUseProxyOk returns a tuple with the UseProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseProxy

`func (o *Site) SetUseProxy(v bool)`

SetUseProxy sets UseProxy field to given value.

### HasUseProxy

`func (o *Site) HasUseProxy() bool`

HasUseProxy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


