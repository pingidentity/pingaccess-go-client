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

// PingFederateAdmin A PingFederate Admin configuration.
type PingFederateAdmin struct {
	AuthenticationType *AuthenticationType `json:"authenticationType,omitempty"`
	// The administrator username. Required when the authentication type is set to 'Basic'.
	AdminUsername *string `json:"adminUsername,omitempty"`
	AdminPassword HiddenField `json:"adminPassword"`
	// Set to true if HTTP communications to PingFederate should not perform hostname verification of the certificate.
	SkipHostnameVerification *bool `json:"skipHostnameVerification,omitempty"`
	// The name of the host expected in the certificate used by PingFederate.
	ExpectedHostname *string `json:"expectedHostname,omitempty"`
	// The group of certificates to use when authenticating to PingFederate Administrative API.
	TrustedCertificateGroupId *int32 `json:"trustedCertificateGroupId,omitempty"`
	// Enable if PingFederate is expecting HTTPS connections.
	Secure *bool `json:"secure,omitempty"`
	// The base path, if needed, for Administration API.
	BasePath *string `json:"basePath,omitempty"`
	// True if a proxy should be used for HTTP or HTTPS requests.
	UseProxy *bool `json:"useProxy,omitempty"`
	// The host name or IP address for PingFederate Administration API.
	Host string `json:"host"`
	// The port number for PingFederate Administration API.
	Port int32 `json:"port"`
	// ['ON' or 'OFF']: Enable to record requests to the PingFederate Administrative API to the audit store.
	AuditLevel *string `json:"auditLevel,omitempty"`
	OAuthAuthenticationConfig *OAuthAuthenticationConfiguration `json:"oAuthAuthenticationConfig,omitempty"`
}

// NewPingFederateAdmin instantiates a new PingFederateAdmin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPingFederateAdmin(adminPassword HiddenField, host string, port int32) *PingFederateAdmin {
	this := PingFederateAdmin{}
	this.AdminPassword = adminPassword
	this.Host = host
	this.Port = port
	return &this
}

// NewPingFederateAdminWithDefaults instantiates a new PingFederateAdmin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPingFederateAdminWithDefaults() *PingFederateAdmin {
	this := PingFederateAdmin{}
	return &this
}

// GetAuthenticationType returns the AuthenticationType field value if set, zero value otherwise.
func (o *PingFederateAdmin) GetAuthenticationType() AuthenticationType {
	if o == nil || isNil(o.AuthenticationType) {
		var ret AuthenticationType
		return ret
	}
	return *o.AuthenticationType
}

// GetAuthenticationTypeOk returns a tuple with the AuthenticationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingFederateAdmin) GetAuthenticationTypeOk() (*AuthenticationType, bool) {
	if o == nil || isNil(o.AuthenticationType) {
    return nil, false
	}
	return o.AuthenticationType, true
}

// HasAuthenticationType returns a boolean if a field has been set.
func (o *PingFederateAdmin) HasAuthenticationType() bool {
	if o != nil && !isNil(o.AuthenticationType) {
		return true
	}

	return false
}

// SetAuthenticationType gets a reference to the given AuthenticationType and assigns it to the AuthenticationType field.
func (o *PingFederateAdmin) SetAuthenticationType(v AuthenticationType) {
	o.AuthenticationType = &v
}

// GetAdminUsername returns the AdminUsername field value if set, zero value otherwise.
func (o *PingFederateAdmin) GetAdminUsername() string {
	if o == nil || isNil(o.AdminUsername) {
		var ret string
		return ret
	}
	return *o.AdminUsername
}

// GetAdminUsernameOk returns a tuple with the AdminUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingFederateAdmin) GetAdminUsernameOk() (*string, bool) {
	if o == nil || isNil(o.AdminUsername) {
    return nil, false
	}
	return o.AdminUsername, true
}

// HasAdminUsername returns a boolean if a field has been set.
func (o *PingFederateAdmin) HasAdminUsername() bool {
	if o != nil && !isNil(o.AdminUsername) {
		return true
	}

	return false
}

// SetAdminUsername gets a reference to the given string and assigns it to the AdminUsername field.
func (o *PingFederateAdmin) SetAdminUsername(v string) {
	o.AdminUsername = &v
}

// GetAdminPassword returns the AdminPassword field value
func (o *PingFederateAdmin) GetAdminPassword() HiddenField {
	if o == nil {
		var ret HiddenField
		return ret
	}

	return o.AdminPassword
}

// GetAdminPasswordOk returns a tuple with the AdminPassword field value
// and a boolean to check if the value has been set.
func (o *PingFederateAdmin) GetAdminPasswordOk() (*HiddenField, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AdminPassword, true
}

// SetAdminPassword sets field value
func (o *PingFederateAdmin) SetAdminPassword(v HiddenField) {
	o.AdminPassword = v
}

// GetSkipHostnameVerification returns the SkipHostnameVerification field value if set, zero value otherwise.
func (o *PingFederateAdmin) GetSkipHostnameVerification() bool {
	if o == nil || isNil(o.SkipHostnameVerification) {
		var ret bool
		return ret
	}
	return *o.SkipHostnameVerification
}

// GetSkipHostnameVerificationOk returns a tuple with the SkipHostnameVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingFederateAdmin) GetSkipHostnameVerificationOk() (*bool, bool) {
	if o == nil || isNil(o.SkipHostnameVerification) {
    return nil, false
	}
	return o.SkipHostnameVerification, true
}

// HasSkipHostnameVerification returns a boolean if a field has been set.
func (o *PingFederateAdmin) HasSkipHostnameVerification() bool {
	if o != nil && !isNil(o.SkipHostnameVerification) {
		return true
	}

	return false
}

// SetSkipHostnameVerification gets a reference to the given bool and assigns it to the SkipHostnameVerification field.
func (o *PingFederateAdmin) SetSkipHostnameVerification(v bool) {
	o.SkipHostnameVerification = &v
}

// GetExpectedHostname returns the ExpectedHostname field value if set, zero value otherwise.
func (o *PingFederateAdmin) GetExpectedHostname() string {
	if o == nil || isNil(o.ExpectedHostname) {
		var ret string
		return ret
	}
	return *o.ExpectedHostname
}

// GetExpectedHostnameOk returns a tuple with the ExpectedHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingFederateAdmin) GetExpectedHostnameOk() (*string, bool) {
	if o == nil || isNil(o.ExpectedHostname) {
    return nil, false
	}
	return o.ExpectedHostname, true
}

// HasExpectedHostname returns a boolean if a field has been set.
func (o *PingFederateAdmin) HasExpectedHostname() bool {
	if o != nil && !isNil(o.ExpectedHostname) {
		return true
	}

	return false
}

// SetExpectedHostname gets a reference to the given string and assigns it to the ExpectedHostname field.
func (o *PingFederateAdmin) SetExpectedHostname(v string) {
	o.ExpectedHostname = &v
}

// GetTrustedCertificateGroupId returns the TrustedCertificateGroupId field value if set, zero value otherwise.
func (o *PingFederateAdmin) GetTrustedCertificateGroupId() int32 {
	if o == nil || isNil(o.TrustedCertificateGroupId) {
		var ret int32
		return ret
	}
	return *o.TrustedCertificateGroupId
}

// GetTrustedCertificateGroupIdOk returns a tuple with the TrustedCertificateGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingFederateAdmin) GetTrustedCertificateGroupIdOk() (*int32, bool) {
	if o == nil || isNil(o.TrustedCertificateGroupId) {
    return nil, false
	}
	return o.TrustedCertificateGroupId, true
}

// HasTrustedCertificateGroupId returns a boolean if a field has been set.
func (o *PingFederateAdmin) HasTrustedCertificateGroupId() bool {
	if o != nil && !isNil(o.TrustedCertificateGroupId) {
		return true
	}

	return false
}

// SetTrustedCertificateGroupId gets a reference to the given int32 and assigns it to the TrustedCertificateGroupId field.
func (o *PingFederateAdmin) SetTrustedCertificateGroupId(v int32) {
	o.TrustedCertificateGroupId = &v
}

// GetSecure returns the Secure field value if set, zero value otherwise.
func (o *PingFederateAdmin) GetSecure() bool {
	if o == nil || isNil(o.Secure) {
		var ret bool
		return ret
	}
	return *o.Secure
}

// GetSecureOk returns a tuple with the Secure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingFederateAdmin) GetSecureOk() (*bool, bool) {
	if o == nil || isNil(o.Secure) {
    return nil, false
	}
	return o.Secure, true
}

// HasSecure returns a boolean if a field has been set.
func (o *PingFederateAdmin) HasSecure() bool {
	if o != nil && !isNil(o.Secure) {
		return true
	}

	return false
}

// SetSecure gets a reference to the given bool and assigns it to the Secure field.
func (o *PingFederateAdmin) SetSecure(v bool) {
	o.Secure = &v
}

// GetBasePath returns the BasePath field value if set, zero value otherwise.
func (o *PingFederateAdmin) GetBasePath() string {
	if o == nil || isNil(o.BasePath) {
		var ret string
		return ret
	}
	return *o.BasePath
}

// GetBasePathOk returns a tuple with the BasePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingFederateAdmin) GetBasePathOk() (*string, bool) {
	if o == nil || isNil(o.BasePath) {
    return nil, false
	}
	return o.BasePath, true
}

// HasBasePath returns a boolean if a field has been set.
func (o *PingFederateAdmin) HasBasePath() bool {
	if o != nil && !isNil(o.BasePath) {
		return true
	}

	return false
}

// SetBasePath gets a reference to the given string and assigns it to the BasePath field.
func (o *PingFederateAdmin) SetBasePath(v string) {
	o.BasePath = &v
}

// GetUseProxy returns the UseProxy field value if set, zero value otherwise.
func (o *PingFederateAdmin) GetUseProxy() bool {
	if o == nil || isNil(o.UseProxy) {
		var ret bool
		return ret
	}
	return *o.UseProxy
}

// GetUseProxyOk returns a tuple with the UseProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingFederateAdmin) GetUseProxyOk() (*bool, bool) {
	if o == nil || isNil(o.UseProxy) {
    return nil, false
	}
	return o.UseProxy, true
}

// HasUseProxy returns a boolean if a field has been set.
func (o *PingFederateAdmin) HasUseProxy() bool {
	if o != nil && !isNil(o.UseProxy) {
		return true
	}

	return false
}

// SetUseProxy gets a reference to the given bool and assigns it to the UseProxy field.
func (o *PingFederateAdmin) SetUseProxy(v bool) {
	o.UseProxy = &v
}

// GetHost returns the Host field value
func (o *PingFederateAdmin) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *PingFederateAdmin) GetHostOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *PingFederateAdmin) SetHost(v string) {
	o.Host = v
}

// GetPort returns the Port field value
func (o *PingFederateAdmin) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *PingFederateAdmin) GetPortOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *PingFederateAdmin) SetPort(v int32) {
	o.Port = v
}

// GetAuditLevel returns the AuditLevel field value if set, zero value otherwise.
func (o *PingFederateAdmin) GetAuditLevel() string {
	if o == nil || isNil(o.AuditLevel) {
		var ret string
		return ret
	}
	return *o.AuditLevel
}

// GetAuditLevelOk returns a tuple with the AuditLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingFederateAdmin) GetAuditLevelOk() (*string, bool) {
	if o == nil || isNil(o.AuditLevel) {
    return nil, false
	}
	return o.AuditLevel, true
}

// HasAuditLevel returns a boolean if a field has been set.
func (o *PingFederateAdmin) HasAuditLevel() bool {
	if o != nil && !isNil(o.AuditLevel) {
		return true
	}

	return false
}

// SetAuditLevel gets a reference to the given string and assigns it to the AuditLevel field.
func (o *PingFederateAdmin) SetAuditLevel(v string) {
	o.AuditLevel = &v
}

// GetOAuthAuthenticationConfig returns the OAuthAuthenticationConfig field value if set, zero value otherwise.
func (o *PingFederateAdmin) GetOAuthAuthenticationConfig() OAuthAuthenticationConfiguration {
	if o == nil || isNil(o.OAuthAuthenticationConfig) {
		var ret OAuthAuthenticationConfiguration
		return ret
	}
	return *o.OAuthAuthenticationConfig
}

// GetOAuthAuthenticationConfigOk returns a tuple with the OAuthAuthenticationConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingFederateAdmin) GetOAuthAuthenticationConfigOk() (*OAuthAuthenticationConfiguration, bool) {
	if o == nil || isNil(o.OAuthAuthenticationConfig) {
    return nil, false
	}
	return o.OAuthAuthenticationConfig, true
}

// HasOAuthAuthenticationConfig returns a boolean if a field has been set.
func (o *PingFederateAdmin) HasOAuthAuthenticationConfig() bool {
	if o != nil && !isNil(o.OAuthAuthenticationConfig) {
		return true
	}

	return false
}

// SetOAuthAuthenticationConfig gets a reference to the given OAuthAuthenticationConfiguration and assigns it to the OAuthAuthenticationConfig field.
func (o *PingFederateAdmin) SetOAuthAuthenticationConfig(v OAuthAuthenticationConfiguration) {
	o.OAuthAuthenticationConfig = &v
}

func (o PingFederateAdmin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AuthenticationType) {
		toSerialize["authenticationType"] = o.AuthenticationType
	}
	if !isNil(o.AdminUsername) {
		toSerialize["adminUsername"] = o.AdminUsername
	}
	if true {
		toSerialize["adminPassword"] = o.AdminPassword
	}
	if !isNil(o.SkipHostnameVerification) {
		toSerialize["skipHostnameVerification"] = o.SkipHostnameVerification
	}
	if !isNil(o.ExpectedHostname) {
		toSerialize["expectedHostname"] = o.ExpectedHostname
	}
	if !isNil(o.TrustedCertificateGroupId) {
		toSerialize["trustedCertificateGroupId"] = o.TrustedCertificateGroupId
	}
	if !isNil(o.Secure) {
		toSerialize["secure"] = o.Secure
	}
	if !isNil(o.BasePath) {
		toSerialize["basePath"] = o.BasePath
	}
	if !isNil(o.UseProxy) {
		toSerialize["useProxy"] = o.UseProxy
	}
	if true {
		toSerialize["host"] = o.Host
	}
	if true {
		toSerialize["port"] = o.Port
	}
	if !isNil(o.AuditLevel) {
		toSerialize["auditLevel"] = o.AuditLevel
	}
	if !isNil(o.OAuthAuthenticationConfig) {
		toSerialize["oAuthAuthenticationConfig"] = o.OAuthAuthenticationConfig
	}
	return json.Marshal(toSerialize)
}

type NullablePingFederateAdmin struct {
	value *PingFederateAdmin
	isSet bool
}

func (v NullablePingFederateAdmin) Get() *PingFederateAdmin {
	return v.value
}

func (v *NullablePingFederateAdmin) Set(val *PingFederateAdmin) {
	v.value = val
	v.isSet = true
}

func (v NullablePingFederateAdmin) IsSet() bool {
	return v.isSet
}

func (v *NullablePingFederateAdmin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePingFederateAdmin(val *PingFederateAdmin) *NullablePingFederateAdmin {
	return &NullablePingFederateAdmin{value: val, isSet: true}
}

func (v NullablePingFederateAdmin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePingFederateAdmin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

