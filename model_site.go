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

// checks if the Site type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Site{}

// Site A site.
type Site struct {
	// When creating a new Site, this is the ID for the Site. If not specified, an ID will be automatically assigned. When updating an existing Site, this field is ignored and the ID is determined by the path parameter.
	Id *int64 `json:"id,omitempty"`
	// (sortable) The name of the site.
	Name string `json:"name"`
	// The {hostname}:{port} pairs for the hosts that make up the site.
	Targets []string `json:"targets"`
	// (sortable) This field is true if the site expects HTTPS connections.
	Secure *bool `json:"secure,omitempty"`
	// The ID of the trusted certificate group associated with the site.
	TrustedCertificateGroupId *int64 `json:"trustedCertificateGroupId,omitempty"`
	// (sortable) This field is true if the PingAccess Token or OAuth Access Token should be included in the request to the site.
	SendPaCookie *bool `json:"sendPaCookie,omitempty"`
	// (sortable) Setting this field to true causes PingAccess to adjust the Host header to the site's selected target host rather than the virtual host configured in the application.
	UseTargetHostHeader *bool `json:"useTargetHostHeader,omitempty"`
	// (sortable) The time, in milliseconds, that an HTTP persistent connection to the site can be idle before PingAccess closes the connection.
	KeepAliveTimeout *int64 `json:"keepAliveTimeout,omitempty"`
	// (sortable) The maximum number of HTTP persistent connections you want PingAccess to have open and maintain for the site. -1 indicates unlimited connections.
	MaxConnections *int64 `json:"maxConnections,omitempty"`
	// (sortable) The maximum number of WebSocket connections you want PingAccess to have open and maintain for the site. -1 indicates unlimited connections.
	MaxWebSocketConnections *int64 `json:"maxWebSocketConnections,omitempty"`
	// The IDs of the site authenticators associated with the site.
	SiteAuthenticatorIds []int64 `json:"siteAuthenticatorIds,omitempty"`
	// (sortable) This field is true if the hostname verification of the site's certificate should be skipped.
	SkipHostnameVerification *bool `json:"skipHostnameVerification,omitempty"`
	// (sortable) The name of the host expected in the site's certificate.
	ExpectedHostname *string `json:"expectedHostname,omitempty"`
	// The ID of the availability profile associated with the site.
	AvailabilityProfileId *int64 `json:"availabilityProfileId,omitempty"`
	// The ID of the load balancing strategy associated with the site.
	LoadBalancingStrategyId *int64 `json:"loadBalancingStrategyId,omitempty"`
	// (sortable) True if a proxy should be used for HTTP or HTTPS requests.
	UseProxy *bool `json:"useProxy,omitempty"`
}

// NewSite instantiates a new Site object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSite(name string, targets []string) *Site {
	this := Site{}
	this.Name = name
	this.Targets = targets
	return &this
}

// NewSiteWithDefaults instantiates a new Site object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteWithDefaults() *Site {
	this := Site{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Site) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Site) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Site) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *Site) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Site) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Site) SetName(v string) {
	o.Name = v
}

// GetTargets returns the Targets field value
func (o *Site) GetTargets() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value
// and a boolean to check if the value has been set.
func (o *Site) GetTargetsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Targets, true
}

// SetTargets sets field value
func (o *Site) SetTargets(v []string) {
	o.Targets = v
}

// GetSecure returns the Secure field value if set, zero value otherwise.
func (o *Site) GetSecure() bool {
	if o == nil || IsNil(o.Secure) {
		var ret bool
		return ret
	}
	return *o.Secure
}

// GetSecureOk returns a tuple with the Secure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetSecureOk() (*bool, bool) {
	if o == nil || IsNil(o.Secure) {
		return nil, false
	}
	return o.Secure, true
}

// HasSecure returns a boolean if a field has been set.
func (o *Site) HasSecure() bool {
	if o != nil && !IsNil(o.Secure) {
		return true
	}

	return false
}

// SetSecure gets a reference to the given bool and assigns it to the Secure field.
func (o *Site) SetSecure(v bool) {
	o.Secure = &v
}

// GetTrustedCertificateGroupId returns the TrustedCertificateGroupId field value if set, zero value otherwise.
func (o *Site) GetTrustedCertificateGroupId() int64 {
	if o == nil || IsNil(o.TrustedCertificateGroupId) {
		var ret int64
		return ret
	}
	return *o.TrustedCertificateGroupId
}

// GetTrustedCertificateGroupIdOk returns a tuple with the TrustedCertificateGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetTrustedCertificateGroupIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TrustedCertificateGroupId) {
		return nil, false
	}
	return o.TrustedCertificateGroupId, true
}

// HasTrustedCertificateGroupId returns a boolean if a field has been set.
func (o *Site) HasTrustedCertificateGroupId() bool {
	if o != nil && !IsNil(o.TrustedCertificateGroupId) {
		return true
	}

	return false
}

// SetTrustedCertificateGroupId gets a reference to the given int64 and assigns it to the TrustedCertificateGroupId field.
func (o *Site) SetTrustedCertificateGroupId(v int64) {
	o.TrustedCertificateGroupId = &v
}

// GetSendPaCookie returns the SendPaCookie field value if set, zero value otherwise.
func (o *Site) GetSendPaCookie() bool {
	if o == nil || IsNil(o.SendPaCookie) {
		var ret bool
		return ret
	}
	return *o.SendPaCookie
}

// GetSendPaCookieOk returns a tuple with the SendPaCookie field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetSendPaCookieOk() (*bool, bool) {
	if o == nil || IsNil(o.SendPaCookie) {
		return nil, false
	}
	return o.SendPaCookie, true
}

// HasSendPaCookie returns a boolean if a field has been set.
func (o *Site) HasSendPaCookie() bool {
	if o != nil && !IsNil(o.SendPaCookie) {
		return true
	}

	return false
}

// SetSendPaCookie gets a reference to the given bool and assigns it to the SendPaCookie field.
func (o *Site) SetSendPaCookie(v bool) {
	o.SendPaCookie = &v
}

// GetUseTargetHostHeader returns the UseTargetHostHeader field value if set, zero value otherwise.
func (o *Site) GetUseTargetHostHeader() bool {
	if o == nil || IsNil(o.UseTargetHostHeader) {
		var ret bool
		return ret
	}
	return *o.UseTargetHostHeader
}

// GetUseTargetHostHeaderOk returns a tuple with the UseTargetHostHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetUseTargetHostHeaderOk() (*bool, bool) {
	if o == nil || IsNil(o.UseTargetHostHeader) {
		return nil, false
	}
	return o.UseTargetHostHeader, true
}

// HasUseTargetHostHeader returns a boolean if a field has been set.
func (o *Site) HasUseTargetHostHeader() bool {
	if o != nil && !IsNil(o.UseTargetHostHeader) {
		return true
	}

	return false
}

// SetUseTargetHostHeader gets a reference to the given bool and assigns it to the UseTargetHostHeader field.
func (o *Site) SetUseTargetHostHeader(v bool) {
	o.UseTargetHostHeader = &v
}

// GetKeepAliveTimeout returns the KeepAliveTimeout field value if set, zero value otherwise.
func (o *Site) GetKeepAliveTimeout() int64 {
	if o == nil || IsNil(o.KeepAliveTimeout) {
		var ret int64
		return ret
	}
	return *o.KeepAliveTimeout
}

// GetKeepAliveTimeoutOk returns a tuple with the KeepAliveTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetKeepAliveTimeoutOk() (*int64, bool) {
	if o == nil || IsNil(o.KeepAliveTimeout) {
		return nil, false
	}
	return o.KeepAliveTimeout, true
}

// HasKeepAliveTimeout returns a boolean if a field has been set.
func (o *Site) HasKeepAliveTimeout() bool {
	if o != nil && !IsNil(o.KeepAliveTimeout) {
		return true
	}

	return false
}

// SetKeepAliveTimeout gets a reference to the given int64 and assigns it to the KeepAliveTimeout field.
func (o *Site) SetKeepAliveTimeout(v int64) {
	o.KeepAliveTimeout = &v
}

// GetMaxConnections returns the MaxConnections field value if set, zero value otherwise.
func (o *Site) GetMaxConnections() int64 {
	if o == nil || IsNil(o.MaxConnections) {
		var ret int64
		return ret
	}
	return *o.MaxConnections
}

// GetMaxConnectionsOk returns a tuple with the MaxConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetMaxConnectionsOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxConnections) {
		return nil, false
	}
	return o.MaxConnections, true
}

// HasMaxConnections returns a boolean if a field has been set.
func (o *Site) HasMaxConnections() bool {
	if o != nil && !IsNil(o.MaxConnections) {
		return true
	}

	return false
}

// SetMaxConnections gets a reference to the given int64 and assigns it to the MaxConnections field.
func (o *Site) SetMaxConnections(v int64) {
	o.MaxConnections = &v
}

// GetMaxWebSocketConnections returns the MaxWebSocketConnections field value if set, zero value otherwise.
func (o *Site) GetMaxWebSocketConnections() int64 {
	if o == nil || IsNil(o.MaxWebSocketConnections) {
		var ret int64
		return ret
	}
	return *o.MaxWebSocketConnections
}

// GetMaxWebSocketConnectionsOk returns a tuple with the MaxWebSocketConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetMaxWebSocketConnectionsOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxWebSocketConnections) {
		return nil, false
	}
	return o.MaxWebSocketConnections, true
}

// HasMaxWebSocketConnections returns a boolean if a field has been set.
func (o *Site) HasMaxWebSocketConnections() bool {
	if o != nil && !IsNil(o.MaxWebSocketConnections) {
		return true
	}

	return false
}

// SetMaxWebSocketConnections gets a reference to the given int64 and assigns it to the MaxWebSocketConnections field.
func (o *Site) SetMaxWebSocketConnections(v int64) {
	o.MaxWebSocketConnections = &v
}

// GetSiteAuthenticatorIds returns the SiteAuthenticatorIds field value if set, zero value otherwise.
func (o *Site) GetSiteAuthenticatorIds() []int64 {
	if o == nil || IsNil(o.SiteAuthenticatorIds) {
		var ret []int64
		return ret
	}
	return o.SiteAuthenticatorIds
}

// GetSiteAuthenticatorIdsOk returns a tuple with the SiteAuthenticatorIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetSiteAuthenticatorIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.SiteAuthenticatorIds) {
		return nil, false
	}
	return o.SiteAuthenticatorIds, true
}

// HasSiteAuthenticatorIds returns a boolean if a field has been set.
func (o *Site) HasSiteAuthenticatorIds() bool {
	if o != nil && !IsNil(o.SiteAuthenticatorIds) {
		return true
	}

	return false
}

// SetSiteAuthenticatorIds gets a reference to the given []int64 and assigns it to the SiteAuthenticatorIds field.
func (o *Site) SetSiteAuthenticatorIds(v []int64) {
	o.SiteAuthenticatorIds = v
}

// GetSkipHostnameVerification returns the SkipHostnameVerification field value if set, zero value otherwise.
func (o *Site) GetSkipHostnameVerification() bool {
	if o == nil || IsNil(o.SkipHostnameVerification) {
		var ret bool
		return ret
	}
	return *o.SkipHostnameVerification
}

// GetSkipHostnameVerificationOk returns a tuple with the SkipHostnameVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetSkipHostnameVerificationOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipHostnameVerification) {
		return nil, false
	}
	return o.SkipHostnameVerification, true
}

// HasSkipHostnameVerification returns a boolean if a field has been set.
func (o *Site) HasSkipHostnameVerification() bool {
	if o != nil && !IsNil(o.SkipHostnameVerification) {
		return true
	}

	return false
}

// SetSkipHostnameVerification gets a reference to the given bool and assigns it to the SkipHostnameVerification field.
func (o *Site) SetSkipHostnameVerification(v bool) {
	o.SkipHostnameVerification = &v
}

// GetExpectedHostname returns the ExpectedHostname field value if set, zero value otherwise.
func (o *Site) GetExpectedHostname() string {
	if o == nil || IsNil(o.ExpectedHostname) {
		var ret string
		return ret
	}
	return *o.ExpectedHostname
}

// GetExpectedHostnameOk returns a tuple with the ExpectedHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetExpectedHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.ExpectedHostname) {
		return nil, false
	}
	return o.ExpectedHostname, true
}

// HasExpectedHostname returns a boolean if a field has been set.
func (o *Site) HasExpectedHostname() bool {
	if o != nil && !IsNil(o.ExpectedHostname) {
		return true
	}

	return false
}

// SetExpectedHostname gets a reference to the given string and assigns it to the ExpectedHostname field.
func (o *Site) SetExpectedHostname(v string) {
	o.ExpectedHostname = &v
}

// GetAvailabilityProfileId returns the AvailabilityProfileId field value if set, zero value otherwise.
func (o *Site) GetAvailabilityProfileId() int64 {
	if o == nil || IsNil(o.AvailabilityProfileId) {
		var ret int64
		return ret
	}
	return *o.AvailabilityProfileId
}

// GetAvailabilityProfileIdOk returns a tuple with the AvailabilityProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetAvailabilityProfileIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AvailabilityProfileId) {
		return nil, false
	}
	return o.AvailabilityProfileId, true
}

// HasAvailabilityProfileId returns a boolean if a field has been set.
func (o *Site) HasAvailabilityProfileId() bool {
	if o != nil && !IsNil(o.AvailabilityProfileId) {
		return true
	}

	return false
}

// SetAvailabilityProfileId gets a reference to the given int64 and assigns it to the AvailabilityProfileId field.
func (o *Site) SetAvailabilityProfileId(v int64) {
	o.AvailabilityProfileId = &v
}

// GetLoadBalancingStrategyId returns the LoadBalancingStrategyId field value if set, zero value otherwise.
func (o *Site) GetLoadBalancingStrategyId() int64 {
	if o == nil || IsNil(o.LoadBalancingStrategyId) {
		var ret int64
		return ret
	}
	return *o.LoadBalancingStrategyId
}

// GetLoadBalancingStrategyIdOk returns a tuple with the LoadBalancingStrategyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetLoadBalancingStrategyIdOk() (*int64, bool) {
	if o == nil || IsNil(o.LoadBalancingStrategyId) {
		return nil, false
	}
	return o.LoadBalancingStrategyId, true
}

// HasLoadBalancingStrategyId returns a boolean if a field has been set.
func (o *Site) HasLoadBalancingStrategyId() bool {
	if o != nil && !IsNil(o.LoadBalancingStrategyId) {
		return true
	}

	return false
}

// SetLoadBalancingStrategyId gets a reference to the given int64 and assigns it to the LoadBalancingStrategyId field.
func (o *Site) SetLoadBalancingStrategyId(v int64) {
	o.LoadBalancingStrategyId = &v
}

// GetUseProxy returns the UseProxy field value if set, zero value otherwise.
func (o *Site) GetUseProxy() bool {
	if o == nil || IsNil(o.UseProxy) {
		var ret bool
		return ret
	}
	return *o.UseProxy
}

// GetUseProxyOk returns a tuple with the UseProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetUseProxyOk() (*bool, bool) {
	if o == nil || IsNil(o.UseProxy) {
		return nil, false
	}
	return o.UseProxy, true
}

// HasUseProxy returns a boolean if a field has been set.
func (o *Site) HasUseProxy() bool {
	if o != nil && !IsNil(o.UseProxy) {
		return true
	}

	return false
}

// SetUseProxy gets a reference to the given bool and assigns it to the UseProxy field.
func (o *Site) SetUseProxy(v bool) {
	o.UseProxy = &v
}

func (o Site) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Site) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	toSerialize["targets"] = o.Targets
	if !IsNil(o.Secure) {
		toSerialize["secure"] = o.Secure
	}
	if !IsNil(o.TrustedCertificateGroupId) {
		toSerialize["trustedCertificateGroupId"] = o.TrustedCertificateGroupId
	}
	if !IsNil(o.SendPaCookie) {
		toSerialize["sendPaCookie"] = o.SendPaCookie
	}
	if !IsNil(o.UseTargetHostHeader) {
		toSerialize["useTargetHostHeader"] = o.UseTargetHostHeader
	}
	if !IsNil(o.KeepAliveTimeout) {
		toSerialize["keepAliveTimeout"] = o.KeepAliveTimeout
	}
	if !IsNil(o.MaxConnections) {
		toSerialize["maxConnections"] = o.MaxConnections
	}
	if !IsNil(o.MaxWebSocketConnections) {
		toSerialize["maxWebSocketConnections"] = o.MaxWebSocketConnections
	}
	if !IsNil(o.SiteAuthenticatorIds) {
		toSerialize["siteAuthenticatorIds"] = o.SiteAuthenticatorIds
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
	return toSerialize, nil
}

type NullableSite struct {
	value *Site
	isSet bool
}

func (v NullableSite) Get() *Site {
	return v.value
}

func (v *NullableSite) Set(val *Site) {
	v.value = val
	v.isSet = true
}

func (v NullableSite) IsSet() bool {
	return v.isSet
}

func (v *NullableSite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSite(val *Site) *NullableSite {
	return &NullableSite{value: val, isSet: true}
}

func (v NullableSite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
