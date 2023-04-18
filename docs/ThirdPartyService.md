# ThirdPartyService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | When creating a new ThirdPartyService, this is the ID for the ThirdPartyService. If not specified, an ID will be automatically assigned. When updating an existing ThirdPartyService, this field is ignored and the ID is determined by the path parameter. | [optional] 
**Targets** | **[]string** | The {hostname}:{port} pairs for the hosts that make up the third-party service. | 
**Name** | **string** | (sortable) The name of the third-party service. | 
**Secure** | Pointer to **bool** | (sortable) This field is true if the third-party service expects HTTPS connections. | [optional] 
**TrustedCertificateGroupId** | Pointer to **int64** | The ID of the trusted certificate group associated with the third-party service. | [optional] 
**MaxConnections** | Pointer to **int64** | (sortable) The maximum number of HTTP persistent connections you want PingAccess to have open and maintain for the third-party service. -1 indicates unlimited connections. | [optional] 
**SkipHostnameVerification** | Pointer to **bool** | (sortable) This field is true if the hostname verification of the third-party service&#39;s certificate should be skipped. | [optional] 
**ExpectedHostname** | Pointer to **string** | (sortable) The name of the host expected in the third-party service&#39;s certificate. | [optional] 
**AvailabilityProfileId** | Pointer to **int64** | The ID of the availability profile associated with the third-party service. | [optional] 
**LoadBalancingStrategyId** | Pointer to **int64** | The ID of the load balancing strategy associated with the third-party service. | [optional] 
**UseProxy** | Pointer to **bool** | (sortable) True if a proxy should be used for HTTP or HTTPS requests. | [optional] 
**HostValue** | Pointer to **string** | (sortable) The Host header field value in the requests sent to a Third-Party Services. When set, PingAccess will use the hostValue as the Host header field value. Otherwise, the target value will be used. | [optional] 

## Methods

### NewThirdPartyService

`func NewThirdPartyService(targets []string, name string, ) *ThirdPartyService`

NewThirdPartyService instantiates a new ThirdPartyService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartyServiceWithDefaults

`func NewThirdPartyServiceWithDefaults() *ThirdPartyService`

NewThirdPartyServiceWithDefaults instantiates a new ThirdPartyService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ThirdPartyService) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThirdPartyService) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThirdPartyService) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ThirdPartyService) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTargets

`func (o *ThirdPartyService) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *ThirdPartyService) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *ThirdPartyService) SetTargets(v []string)`

SetTargets sets Targets field to given value.


### GetName

`func (o *ThirdPartyService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ThirdPartyService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ThirdPartyService) SetName(v string)`

SetName sets Name field to given value.


### GetSecure

`func (o *ThirdPartyService) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *ThirdPartyService) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *ThirdPartyService) SetSecure(v bool)`

SetSecure sets Secure field to given value.

### HasSecure

`func (o *ThirdPartyService) HasSecure() bool`

HasSecure returns a boolean if a field has been set.

### GetTrustedCertificateGroupId

`func (o *ThirdPartyService) GetTrustedCertificateGroupId() int64`

GetTrustedCertificateGroupId returns the TrustedCertificateGroupId field if non-nil, zero value otherwise.

### GetTrustedCertificateGroupIdOk

`func (o *ThirdPartyService) GetTrustedCertificateGroupIdOk() (*int64, bool)`

GetTrustedCertificateGroupIdOk returns a tuple with the TrustedCertificateGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedCertificateGroupId

`func (o *ThirdPartyService) SetTrustedCertificateGroupId(v int64)`

SetTrustedCertificateGroupId sets TrustedCertificateGroupId field to given value.

### HasTrustedCertificateGroupId

`func (o *ThirdPartyService) HasTrustedCertificateGroupId() bool`

HasTrustedCertificateGroupId returns a boolean if a field has been set.

### GetMaxConnections

`func (o *ThirdPartyService) GetMaxConnections() int64`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *ThirdPartyService) GetMaxConnectionsOk() (*int64, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *ThirdPartyService) SetMaxConnections(v int64)`

SetMaxConnections sets MaxConnections field to given value.

### HasMaxConnections

`func (o *ThirdPartyService) HasMaxConnections() bool`

HasMaxConnections returns a boolean if a field has been set.

### GetSkipHostnameVerification

`func (o *ThirdPartyService) GetSkipHostnameVerification() bool`

GetSkipHostnameVerification returns the SkipHostnameVerification field if non-nil, zero value otherwise.

### GetSkipHostnameVerificationOk

`func (o *ThirdPartyService) GetSkipHostnameVerificationOk() (*bool, bool)`

GetSkipHostnameVerificationOk returns a tuple with the SkipHostnameVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipHostnameVerification

`func (o *ThirdPartyService) SetSkipHostnameVerification(v bool)`

SetSkipHostnameVerification sets SkipHostnameVerification field to given value.

### HasSkipHostnameVerification

`func (o *ThirdPartyService) HasSkipHostnameVerification() bool`

HasSkipHostnameVerification returns a boolean if a field has been set.

### GetExpectedHostname

`func (o *ThirdPartyService) GetExpectedHostname() string`

GetExpectedHostname returns the ExpectedHostname field if non-nil, zero value otherwise.

### GetExpectedHostnameOk

`func (o *ThirdPartyService) GetExpectedHostnameOk() (*string, bool)`

GetExpectedHostnameOk returns a tuple with the ExpectedHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedHostname

`func (o *ThirdPartyService) SetExpectedHostname(v string)`

SetExpectedHostname sets ExpectedHostname field to given value.

### HasExpectedHostname

`func (o *ThirdPartyService) HasExpectedHostname() bool`

HasExpectedHostname returns a boolean if a field has been set.

### GetAvailabilityProfileId

`func (o *ThirdPartyService) GetAvailabilityProfileId() int64`

GetAvailabilityProfileId returns the AvailabilityProfileId field if non-nil, zero value otherwise.

### GetAvailabilityProfileIdOk

`func (o *ThirdPartyService) GetAvailabilityProfileIdOk() (*int64, bool)`

GetAvailabilityProfileIdOk returns a tuple with the AvailabilityProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityProfileId

`func (o *ThirdPartyService) SetAvailabilityProfileId(v int64)`

SetAvailabilityProfileId sets AvailabilityProfileId field to given value.

### HasAvailabilityProfileId

`func (o *ThirdPartyService) HasAvailabilityProfileId() bool`

HasAvailabilityProfileId returns a boolean if a field has been set.

### GetLoadBalancingStrategyId

`func (o *ThirdPartyService) GetLoadBalancingStrategyId() int64`

GetLoadBalancingStrategyId returns the LoadBalancingStrategyId field if non-nil, zero value otherwise.

### GetLoadBalancingStrategyIdOk

`func (o *ThirdPartyService) GetLoadBalancingStrategyIdOk() (*int64, bool)`

GetLoadBalancingStrategyIdOk returns a tuple with the LoadBalancingStrategyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancingStrategyId

`func (o *ThirdPartyService) SetLoadBalancingStrategyId(v int64)`

SetLoadBalancingStrategyId sets LoadBalancingStrategyId field to given value.

### HasLoadBalancingStrategyId

`func (o *ThirdPartyService) HasLoadBalancingStrategyId() bool`

HasLoadBalancingStrategyId returns a boolean if a field has been set.

### GetUseProxy

`func (o *ThirdPartyService) GetUseProxy() bool`

GetUseProxy returns the UseProxy field if non-nil, zero value otherwise.

### GetUseProxyOk

`func (o *ThirdPartyService) GetUseProxyOk() (*bool, bool)`

GetUseProxyOk returns a tuple with the UseProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseProxy

`func (o *ThirdPartyService) SetUseProxy(v bool)`

SetUseProxy sets UseProxy field to given value.

### HasUseProxy

`func (o *ThirdPartyService) HasUseProxy() bool`

HasUseProxy returns a boolean if a field has been set.

### GetHostValue

`func (o *ThirdPartyService) GetHostValue() string`

GetHostValue returns the HostValue field if non-nil, zero value otherwise.

### GetHostValueOk

`func (o *ThirdPartyService) GetHostValueOk() (*string, bool)`

GetHostValueOk returns a tuple with the HostValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostValue

`func (o *ThirdPartyService) SetHostValue(v string)`

SetHostValue sets HostValue field to given value.

### HasHostValue

`func (o *ThirdPartyService) HasHostValue() bool`

HasHostValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


