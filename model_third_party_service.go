/*
Administrative API Documentation

The PingAccess Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingAccess as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API.

API version: 7.3.0.2-SNAPSHOT
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the ThirdPartyService type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThirdPartyService{}

// ThirdPartyService A third-party service.
type ThirdPartyService struct {
	// When creating a new ThirdPartyService, this is the ID for the ThirdPartyService. If not specified, an ID will be automatically assigned. When updating an existing ThirdPartyService, this field is ignored and the ID is determined by the path parameter.
	Id *string `json:"id,omitempty"`
	// The {hostname}:{port} pairs for the hosts that make up the third-party service.
	Targets []string `json:"targets"`
	// (sortable) The name of the third-party service.
	Name string `json:"name"`
	// (sortable) This field is true if the third-party service expects HTTPS connections.
	Secure *bool `json:"secure,omitempty"`
	// The ID of the trusted certificate group associated with the third-party service.
	TrustedCertificateGroupId *int64 `json:"trustedCertificateGroupId,omitempty"`
	// (sortable) The maximum number of HTTP persistent connections you want PingAccess to have open and maintain for the third-party service. -1 indicates unlimited connections.
	MaxConnections *int64 `json:"maxConnections,omitempty"`
	// (sortable) This field is true if the hostname verification of the third-party service's certificate should be skipped.
	SkipHostnameVerification *bool `json:"skipHostnameVerification,omitempty"`
	// (sortable) The name of the host expected in the third-party service's certificate.
	ExpectedHostname *string `json:"expectedHostname,omitempty"`
	// The ID of the availability profile associated with the third-party service.
	AvailabilityProfileId *int64 `json:"availabilityProfileId,omitempty"`
	// The ID of the load balancing strategy associated with the third-party service.
	LoadBalancingStrategyId *int64 `json:"loadBalancingStrategyId,omitempty"`
	// (sortable) True if a proxy should be used for HTTP or HTTPS requests.
	UseProxy *bool `json:"useProxy,omitempty"`
	// (sortable) The Host header field value in the requests sent to a Third-Party Services. When set, PingAccess will use the hostValue as the Host header field value. Otherwise, the target value will be used.
	HostValue *string `json:"hostValue,omitempty"`
}

// NewThirdPartyService instantiates a new ThirdPartyService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartyService(targets []string, name string) *ThirdPartyService {
	this := ThirdPartyService{}
	this.Targets = targets
	this.Name = name
	return &this
}

// NewThirdPartyServiceWithDefaults instantiates a new ThirdPartyService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartyServiceWithDefaults() *ThirdPartyService {
	this := ThirdPartyService{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ThirdPartyService) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyService) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ThirdPartyService) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ThirdPartyService) SetId(v string) {
	o.Id = &v
}

// GetTargets returns the Targets field value
func (o *ThirdPartyService) GetTargets() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyService) GetTargetsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Targets, true
}

// SetTargets sets field value
func (o *ThirdPartyService) SetTargets(v []string) {
	o.Targets = v
}

// GetName returns the Name field value
func (o *ThirdPartyService) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyService) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ThirdPartyService) SetName(v string) {
	o.Name = v
}

// GetSecure returns the Secure field value if set, zero value otherwise.
func (o *ThirdPartyService) GetSecure() bool {
	if o == nil || IsNil(o.Secure) {
		var ret bool
		return ret
	}
	return *o.Secure
}

// GetSecureOk returns a tuple with the Secure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyService) GetSecureOk() (*bool, bool) {
	if o == nil || IsNil(o.Secure) {
		return nil, false
	}
	return o.Secure, true
}

// HasSecure returns a boolean if a field has been set.
func (o *ThirdPartyService) HasSecure() bool {
	if o != nil && !IsNil(o.Secure) {
		return true
	}

	return false
}

// SetSecure gets a reference to the given bool and assigns it to the Secure field.
func (o *ThirdPartyService) SetSecure(v bool) {
	o.Secure = &v
}

// GetTrustedCertificateGroupId returns the TrustedCertificateGroupId field value if set, zero value otherwise.
func (o *ThirdPartyService) GetTrustedCertificateGroupId() int64 {
	if o == nil || IsNil(o.TrustedCertificateGroupId) {
		var ret int64
		return ret
	}
	return *o.TrustedCertificateGroupId
}

// GetTrustedCertificateGroupIdOk returns a tuple with the TrustedCertificateGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyService) GetTrustedCertificateGroupIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TrustedCertificateGroupId) {
		return nil, false
	}
	return o.TrustedCertificateGroupId, true
}

// HasTrustedCertificateGroupId returns a boolean if a field has been set.
func (o *ThirdPartyService) HasTrustedCertificateGroupId() bool {
	if o != nil && !IsNil(o.TrustedCertificateGroupId) {
		return true
	}

	return false
}

// SetTrustedCertificateGroupId gets a reference to the given int64 and assigns it to the TrustedCertificateGroupId field.
func (o *ThirdPartyService) SetTrustedCertificateGroupId(v int64) {
	o.TrustedCertificateGroupId = &v
}

// GetMaxConnections returns the MaxConnections field value if set, zero value otherwise.
func (o *ThirdPartyService) GetMaxConnections() int64 {
	if o == nil || IsNil(o.MaxConnections) {
		var ret int64
		return ret
	}
	return *o.MaxConnections
}

// GetMaxConnectionsOk returns a tuple with the MaxConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyService) GetMaxConnectionsOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxConnections) {
		return nil, false
	}
	return o.MaxConnections, true
}

// HasMaxConnections returns a boolean if a field has been set.
func (o *ThirdPartyService) HasMaxConnections() bool {
	if o != nil && !IsNil(o.MaxConnections) {
		return true
	}

	return false
}

// SetMaxConnections gets a reference to the given int64 and assigns it to the MaxConnections field.
func (o *ThirdPartyService) SetMaxConnections(v int64) {
	o.MaxConnections = &v
}

// GetSkipHostnameVerification returns the SkipHostnameVerification field value if set, zero value otherwise.
func (o *ThirdPartyService) GetSkipHostnameVerification() bool {
	if o == nil || IsNil(o.SkipHostnameVerification) {
		var ret bool
		return ret
	}
	return *o.SkipHostnameVerification
}

// GetSkipHostnameVerificationOk returns a tuple with the SkipHostnameVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyService) GetSkipHostnameVerificationOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipHostnameVerification) {
		return nil, false
	}
	return o.SkipHostnameVerification, true
}

// HasSkipHostnameVerification returns a boolean if a field has been set.
func (o *ThirdPartyService) HasSkipHostnameVerification() bool {
	if o != nil && !IsNil(o.SkipHostnameVerification) {
		return true
	}

	return false
}

// SetSkipHostnameVerification gets a reference to the given bool and assigns it to the SkipHostnameVerification field.
func (o *ThirdPartyService) SetSkipHostnameVerification(v bool) {
	o.SkipHostnameVerification = &v
}

// GetExpectedHostname returns the ExpectedHostname field value if set, zero value otherwise.
func (o *ThirdPartyService) GetExpectedHostname() string {
	if o == nil || IsNil(o.ExpectedHostname) {
		var ret string
		return ret
	}
	return *o.ExpectedHostname
}

// GetExpectedHostnameOk returns a tuple with the ExpectedHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyService) GetExpectedHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.ExpectedHostname) {
		return nil, false
	}
	return o.ExpectedHostname, true
}

// HasExpectedHostname returns a boolean if a field has been set.
func (o *ThirdPartyService) HasExpectedHostname() bool {
	if o != nil && !IsNil(o.ExpectedHostname) {
		return true
	}

	return false
}

// SetExpectedHostname gets a reference to the given string and assigns it to the ExpectedHostname field.
func (o *ThirdPartyService) SetExpectedHostname(v string) {
	o.ExpectedHostname = &v
}

// GetAvailabilityProfileId returns the AvailabilityProfileId field value if set, zero value otherwise.
func (o *ThirdPartyService) GetAvailabilityProfileId() int64 {
	if o == nil || IsNil(o.AvailabilityProfileId) {
		var ret int64
		return ret
	}
	return *o.AvailabilityProfileId
}

// GetAvailabilityProfileIdOk returns a tuple with the AvailabilityProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyService) GetAvailabilityProfileIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AvailabilityProfileId) {
		return nil, false
	}
	return o.AvailabilityProfileId, true
}

// HasAvailabilityProfileId returns a boolean if a field has been set.
func (o *ThirdPartyService) HasAvailabilityProfileId() bool {
	if o != nil && !IsNil(o.AvailabilityProfileId) {
		return true
	}

	return false
}

// SetAvailabilityProfileId gets a reference to the given int64 and assigns it to the AvailabilityProfileId field.
func (o *ThirdPartyService) SetAvailabilityProfileId(v int64) {
	o.AvailabilityProfileId = &v
}

// GetLoadBalancingStrategyId returns the LoadBalancingStrategyId field value if set, zero value otherwise.
func (o *ThirdPartyService) GetLoadBalancingStrategyId() int64 {
	if o == nil || IsNil(o.LoadBalancingStrategyId) {
		var ret int64
		return ret
	}
	return *o.LoadBalancingStrategyId
}

// GetLoadBalancingStrategyIdOk returns a tuple with the LoadBalancingStrategyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyService) GetLoadBalancingStrategyIdOk() (*int64, bool) {
	if o == nil || IsNil(o.LoadBalancingStrategyId) {
		return nil, false
	}
	return o.LoadBalancingStrategyId, true
}

// HasLoadBalancingStrategyId returns a boolean if a field has been set.
func (o *ThirdPartyService) HasLoadBalancingStrategyId() bool {
	if o != nil && !IsNil(o.LoadBalancingStrategyId) {
		return true
	}

	return false
}

// SetLoadBalancingStrategyId gets a reference to the given int64 and assigns it to the LoadBalancingStrategyId field.
func (o *ThirdPartyService) SetLoadBalancingStrategyId(v int64) {
	o.LoadBalancingStrategyId = &v
}

// GetUseProxy returns the UseProxy field value if set, zero value otherwise.
func (o *ThirdPartyService) GetUseProxy() bool {
	if o == nil || IsNil(o.UseProxy) {
		var ret bool
		return ret
	}
	return *o.UseProxy
}

// GetUseProxyOk returns a tuple with the UseProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyService) GetUseProxyOk() (*bool, bool) {
	if o == nil || IsNil(o.UseProxy) {
		return nil, false
	}
	return o.UseProxy, true
}

// HasUseProxy returns a boolean if a field has been set.
func (o *ThirdPartyService) HasUseProxy() bool {
	if o != nil && !IsNil(o.UseProxy) {
		return true
	}

	return false
}

// SetUseProxy gets a reference to the given bool and assigns it to the UseProxy field.
func (o *ThirdPartyService) SetUseProxy(v bool) {
	o.UseProxy = &v
}

// GetHostValue returns the HostValue field value if set, zero value otherwise.
func (o *ThirdPartyService) GetHostValue() string {
	if o == nil || IsNil(o.HostValue) {
		var ret string
		return ret
	}
	return *o.HostValue
}

// GetHostValueOk returns a tuple with the HostValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyService) GetHostValueOk() (*string, bool) {
	if o == nil || IsNil(o.HostValue) {
		return nil, false
	}
	return o.HostValue, true
}

// HasHostValue returns a boolean if a field has been set.
func (o *ThirdPartyService) HasHostValue() bool {
	if o != nil && !IsNil(o.HostValue) {
		return true
	}

	return false
}

// SetHostValue gets a reference to the given string and assigns it to the HostValue field.
func (o *ThirdPartyService) SetHostValue(v string) {
	o.HostValue = &v
}

func (o ThirdPartyService) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThirdPartyService) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["targets"] = o.Targets
	toSerialize["name"] = o.Name
	if !IsNil(o.Secure) {
		toSerialize["secure"] = o.Secure
	}
	if !IsNil(o.TrustedCertificateGroupId) {
		toSerialize["trustedCertificateGroupId"] = o.TrustedCertificateGroupId
	}
	if !IsNil(o.MaxConnections) {
		toSerialize["maxConnections"] = o.MaxConnections
	}
	if !IsNil(o.SkipHostnameVerification) {
		toSerialize["skipHostnameVerification"] = o.SkipHostnameVerification
	}
	if !IsNil(o.ExpectedHostname) {
		toSerialize["expectedHostname"] = o.ExpectedHostname
	}
	if !IsNil(o.AvailabilityProfileId) {
		toSerialize["availabilityProfileId"] = o.AvailabilityProfileId
	}
	if !IsNil(o.LoadBalancingStrategyId) {
		toSerialize["loadBalancingStrategyId"] = o.LoadBalancingStrategyId
	}
	if !IsNil(o.UseProxy) {
		toSerialize["useProxy"] = o.UseProxy
	}
	if !IsNil(o.HostValue) {
		toSerialize["hostValue"] = o.HostValue
	}
	return toSerialize, nil
}

type NullableThirdPartyService struct {
	value *ThirdPartyService
	isSet bool
}

func (v NullableThirdPartyService) Get() *ThirdPartyService {
	return v.value
}

func (v *NullableThirdPartyService) Set(val *ThirdPartyService) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartyService) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartyService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartyService(val *ThirdPartyService) *NullableThirdPartyService {
	return &NullableThirdPartyService{value: val, isSet: true}
}

func (v NullableThirdPartyService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartyService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
