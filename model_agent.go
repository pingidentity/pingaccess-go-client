/*
PingAccess Administrative API

The PingAccess Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingAccess.

API version: 7.1.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Agent An agent.
type Agent struct {
	// When creating a new Agent, this is the ID for the Agent. If not specified, an ID will be automatically assigned. When updating an existing Agent, this field is ignored and the ID is determined by the path parameter.
	Id *int32 `json:"id,omitempty"`
	// (sortable) The name of the agent.
	Name string `json:"name"`
	// (sortable) A description of the agent.
	Description *string `json:"description,omitempty"`
	// (sortable) The hostname where the agent is listening.
	Hostname string `json:"hostname"`
	// (sortable) The port the agent is listening on.
	Port int32 `json:"port"`
	// An array of shared secret ids associated with this agent.
	SharedSecretIds []int32 `json:"sharedSecretIds"`
	// (sortable) Indicates whether the default IP source is overridden for this agent.
	OverrideIpSource *bool `json:"overrideIpSource,omitempty"`
	IpSource *IpMultiValueSource `json:"ipSource,omitempty"`
	// A list of hostname:port strings for the backup PingAccess policy servers.
	FailoverHosts []string `json:"failoverHosts,omitempty"`
	// The number of seconds to wait before an agent should retry an unavailable policy server.
	FailedRetryTimeout *int32 `json:"failedRetryTimeout,omitempty"`
	// The max number of times an agent request will be attempted before failing over to a backup policy server and marking the current policy server as unavailable.
	MaxRetries *int32 `json:"maxRetries,omitempty"`
	UnknownResourceMode *UnknownResourceMode `json:"unknownResourceMode,omitempty"`
	// The ID of the certificate the agent will use to contact PingAccess via SSL/TLS.
	SelectedCertificateId *int32 `json:"selectedCertificateId,omitempty"`
	CertificateHash *Hash `json:"certificateHash,omitempty"`
}

// NewAgent instantiates a new Agent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgent(name string, hostname string, port int32, sharedSecretIds []int32) *Agent {
	this := Agent{}
	this.Name = name
	this.Hostname = hostname
	this.Port = port
	this.SharedSecretIds = sharedSecretIds
	return &this
}

// NewAgentWithDefaults instantiates a new Agent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentWithDefaults() *Agent {
	this := Agent{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Agent) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Agent) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Agent) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *Agent) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Agent) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Agent) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Agent) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Agent) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Agent) SetDescription(v string) {
	o.Description = &v
}

// GetHostname returns the Hostname field value
func (o *Agent) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *Agent) GetHostnameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *Agent) SetHostname(v string) {
	o.Hostname = v
}

// GetPort returns the Port field value
func (o *Agent) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *Agent) GetPortOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *Agent) SetPort(v int32) {
	o.Port = v
}

// GetSharedSecretIds returns the SharedSecretIds field value
func (o *Agent) GetSharedSecretIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.SharedSecretIds
}

// GetSharedSecretIdsOk returns a tuple with the SharedSecretIds field value
// and a boolean to check if the value has been set.
func (o *Agent) GetSharedSecretIdsOk() ([]int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.SharedSecretIds, true
}

// SetSharedSecretIds sets field value
func (o *Agent) SetSharedSecretIds(v []int32) {
	o.SharedSecretIds = v
}

// GetOverrideIpSource returns the OverrideIpSource field value if set, zero value otherwise.
func (o *Agent) GetOverrideIpSource() bool {
	if o == nil || isNil(o.OverrideIpSource) {
		var ret bool
		return ret
	}
	return *o.OverrideIpSource
}

// GetOverrideIpSourceOk returns a tuple with the OverrideIpSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetOverrideIpSourceOk() (*bool, bool) {
	if o == nil || isNil(o.OverrideIpSource) {
    return nil, false
	}
	return o.OverrideIpSource, true
}

// HasOverrideIpSource returns a boolean if a field has been set.
func (o *Agent) HasOverrideIpSource() bool {
	if o != nil && !isNil(o.OverrideIpSource) {
		return true
	}

	return false
}

// SetOverrideIpSource gets a reference to the given bool and assigns it to the OverrideIpSource field.
func (o *Agent) SetOverrideIpSource(v bool) {
	o.OverrideIpSource = &v
}

// GetIpSource returns the IpSource field value if set, zero value otherwise.
func (o *Agent) GetIpSource() IpMultiValueSource {
	if o == nil || isNil(o.IpSource) {
		var ret IpMultiValueSource
		return ret
	}
	return *o.IpSource
}

// GetIpSourceOk returns a tuple with the IpSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetIpSourceOk() (*IpMultiValueSource, bool) {
	if o == nil || isNil(o.IpSource) {
    return nil, false
	}
	return o.IpSource, true
}

// HasIpSource returns a boolean if a field has been set.
func (o *Agent) HasIpSource() bool {
	if o != nil && !isNil(o.IpSource) {
		return true
	}

	return false
}

// SetIpSource gets a reference to the given IpMultiValueSource and assigns it to the IpSource field.
func (o *Agent) SetIpSource(v IpMultiValueSource) {
	o.IpSource = &v
}

// GetFailoverHosts returns the FailoverHosts field value if set, zero value otherwise.
func (o *Agent) GetFailoverHosts() []string {
	if o == nil || isNil(o.FailoverHosts) {
		var ret []string
		return ret
	}
	return o.FailoverHosts
}

// GetFailoverHostsOk returns a tuple with the FailoverHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetFailoverHostsOk() ([]string, bool) {
	if o == nil || isNil(o.FailoverHosts) {
    return nil, false
	}
	return o.FailoverHosts, true
}

// HasFailoverHosts returns a boolean if a field has been set.
func (o *Agent) HasFailoverHosts() bool {
	if o != nil && !isNil(o.FailoverHosts) {
		return true
	}

	return false
}

// SetFailoverHosts gets a reference to the given []string and assigns it to the FailoverHosts field.
func (o *Agent) SetFailoverHosts(v []string) {
	o.FailoverHosts = v
}

// GetFailedRetryTimeout returns the FailedRetryTimeout field value if set, zero value otherwise.
func (o *Agent) GetFailedRetryTimeout() int32 {
	if o == nil || isNil(o.FailedRetryTimeout) {
		var ret int32
		return ret
	}
	return *o.FailedRetryTimeout
}

// GetFailedRetryTimeoutOk returns a tuple with the FailedRetryTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetFailedRetryTimeoutOk() (*int32, bool) {
	if o == nil || isNil(o.FailedRetryTimeout) {
    return nil, false
	}
	return o.FailedRetryTimeout, true
}

// HasFailedRetryTimeout returns a boolean if a field has been set.
func (o *Agent) HasFailedRetryTimeout() bool {
	if o != nil && !isNil(o.FailedRetryTimeout) {
		return true
	}

	return false
}

// SetFailedRetryTimeout gets a reference to the given int32 and assigns it to the FailedRetryTimeout field.
func (o *Agent) SetFailedRetryTimeout(v int32) {
	o.FailedRetryTimeout = &v
}

// GetMaxRetries returns the MaxRetries field value if set, zero value otherwise.
func (o *Agent) GetMaxRetries() int32 {
	if o == nil || isNil(o.MaxRetries) {
		var ret int32
		return ret
	}
	return *o.MaxRetries
}

// GetMaxRetriesOk returns a tuple with the MaxRetries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetMaxRetriesOk() (*int32, bool) {
	if o == nil || isNil(o.MaxRetries) {
    return nil, false
	}
	return o.MaxRetries, true
}

// HasMaxRetries returns a boolean if a field has been set.
func (o *Agent) HasMaxRetries() bool {
	if o != nil && !isNil(o.MaxRetries) {
		return true
	}

	return false
}

// SetMaxRetries gets a reference to the given int32 and assigns it to the MaxRetries field.
func (o *Agent) SetMaxRetries(v int32) {
	o.MaxRetries = &v
}

// GetUnknownResourceMode returns the UnknownResourceMode field value if set, zero value otherwise.
func (o *Agent) GetUnknownResourceMode() UnknownResourceMode {
	if o == nil || isNil(o.UnknownResourceMode) {
		var ret UnknownResourceMode
		return ret
	}
	return *o.UnknownResourceMode
}

// GetUnknownResourceModeOk returns a tuple with the UnknownResourceMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetUnknownResourceModeOk() (*UnknownResourceMode, bool) {
	if o == nil || isNil(o.UnknownResourceMode) {
    return nil, false
	}
	return o.UnknownResourceMode, true
}

// HasUnknownResourceMode returns a boolean if a field has been set.
func (o *Agent) HasUnknownResourceMode() bool {
	if o != nil && !isNil(o.UnknownResourceMode) {
		return true
	}

	return false
}

// SetUnknownResourceMode gets a reference to the given UnknownResourceMode and assigns it to the UnknownResourceMode field.
func (o *Agent) SetUnknownResourceMode(v UnknownResourceMode) {
	o.UnknownResourceMode = &v
}

// GetSelectedCertificateId returns the SelectedCertificateId field value if set, zero value otherwise.
func (o *Agent) GetSelectedCertificateId() int32 {
	if o == nil || isNil(o.SelectedCertificateId) {
		var ret int32
		return ret
	}
	return *o.SelectedCertificateId
}

// GetSelectedCertificateIdOk returns a tuple with the SelectedCertificateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetSelectedCertificateIdOk() (*int32, bool) {
	if o == nil || isNil(o.SelectedCertificateId) {
    return nil, false
	}
	return o.SelectedCertificateId, true
}

// HasSelectedCertificateId returns a boolean if a field has been set.
func (o *Agent) HasSelectedCertificateId() bool {
	if o != nil && !isNil(o.SelectedCertificateId) {
		return true
	}

	return false
}

// SetSelectedCertificateId gets a reference to the given int32 and assigns it to the SelectedCertificateId field.
func (o *Agent) SetSelectedCertificateId(v int32) {
	o.SelectedCertificateId = &v
}

// GetCertificateHash returns the CertificateHash field value if set, zero value otherwise.
func (o *Agent) GetCertificateHash() Hash {
	if o == nil || isNil(o.CertificateHash) {
		var ret Hash
		return ret
	}
	return *o.CertificateHash
}

// GetCertificateHashOk returns a tuple with the CertificateHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetCertificateHashOk() (*Hash, bool) {
	if o == nil || isNil(o.CertificateHash) {
    return nil, false
	}
	return o.CertificateHash, true
}

// HasCertificateHash returns a boolean if a field has been set.
func (o *Agent) HasCertificateHash() bool {
	if o != nil && !isNil(o.CertificateHash) {
		return true
	}

	return false
}

// SetCertificateHash gets a reference to the given Hash and assigns it to the CertificateHash field.
func (o *Agent) SetCertificateHash(v Hash) {
	o.CertificateHash = &v
}

func (o Agent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["hostname"] = o.Hostname
	}
	if true {
		toSerialize["port"] = o.Port
	}
	if true {
		toSerialize["sharedSecretIds"] = o.SharedSecretIds
	}
	if !isNil(o.OverrideIpSource) {
		toSerialize["overrideIpSource"] = o.OverrideIpSource
	}
	if !isNil(o.IpSource) {
		toSerialize["ipSource"] = o.IpSource
	}
	if !isNil(o.FailoverHosts) {
		toSerialize["failoverHosts"] = o.FailoverHosts
	}
	if !isNil(o.FailedRetryTimeout) {
		toSerialize["failedRetryTimeout"] = o.FailedRetryTimeout
	}
	if !isNil(o.MaxRetries) {
		toSerialize["maxRetries"] = o.MaxRetries
	}
	if !isNil(o.UnknownResourceMode) {
		toSerialize["unknownResourceMode"] = o.UnknownResourceMode
	}
	if !isNil(o.SelectedCertificateId) {
		toSerialize["selectedCertificateId"] = o.SelectedCertificateId
	}
	if !isNil(o.CertificateHash) {
		toSerialize["certificateHash"] = o.CertificateHash
	}
	return json.Marshal(toSerialize)
}

type NullableAgent struct {
	value *Agent
	isSet bool
}

func (v NullableAgent) Get() *Agent {
	return v.value
}

func (v *NullableAgent) Set(val *Agent) {
	v.value = val
	v.isSet = true
}

func (v NullableAgent) IsSet() bool {
	return v.isSet
}

func (v *NullableAgent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgent(val *Agent) *NullableAgent {
	return &NullableAgent{value: val, isSet: true}
}

func (v NullableAgent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

