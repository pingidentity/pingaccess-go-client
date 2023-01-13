# Agent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | When creating a new Agent, this is the ID for the Agent. If not specified, an ID will be automatically assigned. When updating an existing Agent, this field is ignored and the ID is determined by the path parameter. | [optional] 
**Name** | **string** | (sortable) The name of the agent. | 
**Description** | Pointer to **string** | (sortable) A description of the agent. | [optional] 
**Hostname** | **string** | (sortable) The hostname where the agent is listening. | 
**Port** | **int32** | (sortable) The port the agent is listening on. | 
**SharedSecretIds** | **[]int32** | An array of shared secret ids associated with this agent. | 
**OverrideIpSource** | Pointer to **bool** | (sortable) Indicates whether the default IP source is overridden for this agent. | [optional] 
**IpSource** | Pointer to [**IpMultiValueSource**](IpMultiValueSource.md) |  | [optional] 
**FailoverHosts** | Pointer to **[]string** | A list of hostname:port strings for the backup PingAccess policy servers. | [optional] 
**FailedRetryTimeout** | Pointer to **int32** | The number of seconds to wait before an agent should retry an unavailable policy server. | [optional] 
**MaxRetries** | Pointer to **int32** | The max number of times an agent request will be attempted before failing over to a backup policy server and marking the current policy server as unavailable. | [optional] 
**UnknownResourceMode** | Pointer to [**UnknownResourceMode**](UnknownResourceMode.md) |  | [optional] 
**SelectedCertificateId** | Pointer to **int32** | The ID of the certificate the agent will use to contact PingAccess via SSL/TLS. | [optional] 
**CertificateHash** | Pointer to [**Hash**](Hash.md) |  | [optional] 

## Methods

### NewAgent

`func NewAgent(name string, hostname string, port int32, sharedSecretIds []int32, ) *Agent`

NewAgent instantiates a new Agent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentWithDefaults

`func NewAgentWithDefaults() *Agent`

NewAgentWithDefaults instantiates a new Agent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Agent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Agent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Agent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Agent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Agent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Agent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Agent) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Agent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Agent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Agent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Agent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHostname

`func (o *Agent) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *Agent) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *Agent) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetPort

`func (o *Agent) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Agent) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Agent) SetPort(v int32)`

SetPort sets Port field to given value.


### GetSharedSecretIds

`func (o *Agent) GetSharedSecretIds() []int32`

GetSharedSecretIds returns the SharedSecretIds field if non-nil, zero value otherwise.

### GetSharedSecretIdsOk

`func (o *Agent) GetSharedSecretIdsOk() (*[]int32, bool)`

GetSharedSecretIdsOk returns a tuple with the SharedSecretIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecretIds

`func (o *Agent) SetSharedSecretIds(v []int32)`

SetSharedSecretIds sets SharedSecretIds field to given value.


### GetOverrideIpSource

`func (o *Agent) GetOverrideIpSource() bool`

GetOverrideIpSource returns the OverrideIpSource field if non-nil, zero value otherwise.

### GetOverrideIpSourceOk

`func (o *Agent) GetOverrideIpSourceOk() (*bool, bool)`

GetOverrideIpSourceOk returns a tuple with the OverrideIpSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideIpSource

`func (o *Agent) SetOverrideIpSource(v bool)`

SetOverrideIpSource sets OverrideIpSource field to given value.

### HasOverrideIpSource

`func (o *Agent) HasOverrideIpSource() bool`

HasOverrideIpSource returns a boolean if a field has been set.

### GetIpSource

`func (o *Agent) GetIpSource() IpMultiValueSource`

GetIpSource returns the IpSource field if non-nil, zero value otherwise.

### GetIpSourceOk

`func (o *Agent) GetIpSourceOk() (*IpMultiValueSource, bool)`

GetIpSourceOk returns a tuple with the IpSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSource

`func (o *Agent) SetIpSource(v IpMultiValueSource)`

SetIpSource sets IpSource field to given value.

### HasIpSource

`func (o *Agent) HasIpSource() bool`

HasIpSource returns a boolean if a field has been set.

### GetFailoverHosts

`func (o *Agent) GetFailoverHosts() []string`

GetFailoverHosts returns the FailoverHosts field if non-nil, zero value otherwise.

### GetFailoverHostsOk

`func (o *Agent) GetFailoverHostsOk() (*[]string, bool)`

GetFailoverHostsOk returns a tuple with the FailoverHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverHosts

`func (o *Agent) SetFailoverHosts(v []string)`

SetFailoverHosts sets FailoverHosts field to given value.

### HasFailoverHosts

`func (o *Agent) HasFailoverHosts() bool`

HasFailoverHosts returns a boolean if a field has been set.

### GetFailedRetryTimeout

`func (o *Agent) GetFailedRetryTimeout() int32`

GetFailedRetryTimeout returns the FailedRetryTimeout field if non-nil, zero value otherwise.

### GetFailedRetryTimeoutOk

`func (o *Agent) GetFailedRetryTimeoutOk() (*int32, bool)`

GetFailedRetryTimeoutOk returns a tuple with the FailedRetryTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedRetryTimeout

`func (o *Agent) SetFailedRetryTimeout(v int32)`

SetFailedRetryTimeout sets FailedRetryTimeout field to given value.

### HasFailedRetryTimeout

`func (o *Agent) HasFailedRetryTimeout() bool`

HasFailedRetryTimeout returns a boolean if a field has been set.

### GetMaxRetries

`func (o *Agent) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *Agent) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *Agent) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *Agent) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### GetUnknownResourceMode

`func (o *Agent) GetUnknownResourceMode() UnknownResourceMode`

GetUnknownResourceMode returns the UnknownResourceMode field if non-nil, zero value otherwise.

### GetUnknownResourceModeOk

`func (o *Agent) GetUnknownResourceModeOk() (*UnknownResourceMode, bool)`

GetUnknownResourceModeOk returns a tuple with the UnknownResourceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownResourceMode

`func (o *Agent) SetUnknownResourceMode(v UnknownResourceMode)`

SetUnknownResourceMode sets UnknownResourceMode field to given value.

### HasUnknownResourceMode

`func (o *Agent) HasUnknownResourceMode() bool`

HasUnknownResourceMode returns a boolean if a field has been set.

### GetSelectedCertificateId

`func (o *Agent) GetSelectedCertificateId() int32`

GetSelectedCertificateId returns the SelectedCertificateId field if non-nil, zero value otherwise.

### GetSelectedCertificateIdOk

`func (o *Agent) GetSelectedCertificateIdOk() (*int32, bool)`

GetSelectedCertificateIdOk returns a tuple with the SelectedCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedCertificateId

`func (o *Agent) SetSelectedCertificateId(v int32)`

SetSelectedCertificateId sets SelectedCertificateId field to given value.

### HasSelectedCertificateId

`func (o *Agent) HasSelectedCertificateId() bool`

HasSelectedCertificateId returns a boolean if a field has been set.

### GetCertificateHash

`func (o *Agent) GetCertificateHash() Hash`

GetCertificateHash returns the CertificateHash field if non-nil, zero value otherwise.

### GetCertificateHashOk

`func (o *Agent) GetCertificateHashOk() (*Hash, bool)`

GetCertificateHashOk returns a tuple with the CertificateHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateHash

`func (o *Agent) SetCertificateHash(v Hash)`

SetCertificateHash sets CertificateHash field to given value.

### HasCertificateHash

`func (o *Agent) HasCertificateHash() bool`

HasCertificateHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


