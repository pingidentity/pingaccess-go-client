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

// checks if the AuthorizationServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizationServer{}

// AuthorizationServer The third-party OAuth 2.0 Authorization Server configuration.
type AuthorizationServer struct {
	// The description of the third-party OAuth 2.0 Authorization Server.
	Description *string `json:"description,omitempty"`
	// One or more server hostname:port pairs used to access third-party OAuth 2.0 Authorization Server.
	Targets    []string    `json:"targets"`
	AuditLevel *AuditLevel `json:"auditLevel,omitempty"`
	// Enable if third-party OAuth 2.0 Authorization Server is expecting HTTPS connections.
	Secure *bool `json:"secure,omitempty"`
	// The group of certificates to use when authenticating to third-party OAuth 2.0 Authorization Server.
	TrustedCertificateGroupId int64                   `json:"trustedCertificateGroupId"`
	ClientCredentials         *OAuthClientCredentials `json:"clientCredentials,omitempty"`
	// Enable to retain token details for subsequent requests.
	CacheTokens *bool `json:"cacheTokens,omitempty"`
	// Defines the number of seconds to cache the access token. -1 means no limit. This value should be less than the PingFederate Token Lifetime.
	TokenTimeToLiveSeconds *int64 `json:"tokenTimeToLiveSeconds,omitempty"`
	// The attribute you want to use from the OAuth access token as the subject for auditing purposes.
	SubjectAttributeName string `json:"subjectAttributeName"`
	// The third-party OAuth 2.0 Authorization Server's token introspection endpoint.
	IntrospectionEndpoint string `json:"introspectionEndpoint"`
	// The third-party OAuth 2.0 Authorization Server's token endpoint.
	TokenEndpoint *string `json:"tokenEndpoint,omitempty"`
	// Enable to send the URI the user requested as the 'aud' OAuth parameter for PingAccess to the OAuth 2.0 Authorization server.
	SendAudience *bool `json:"sendAudience,omitempty"`
	// True if a proxy should be used for HTTP or HTTPS requests.
	UseProxy *bool `json:"useProxy,omitempty"`
}

// NewAuthorizationServer instantiates a new AuthorizationServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationServer(targets []string, trustedCertificateGroupId int64, subjectAttributeName string, introspectionEndpoint string) *AuthorizationServer {
	this := AuthorizationServer{}
	this.Targets = targets
	this.TrustedCertificateGroupId = trustedCertificateGroupId
	this.SubjectAttributeName = subjectAttributeName
	this.IntrospectionEndpoint = introspectionEndpoint
	return &this
}

// NewAuthorizationServerWithDefaults instantiates a new AuthorizationServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationServerWithDefaults() *AuthorizationServer {
	this := AuthorizationServer{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AuthorizationServer) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServer) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AuthorizationServer) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AuthorizationServer) SetDescription(v string) {
	o.Description = &v
}

// GetTargets returns the Targets field value
func (o *AuthorizationServer) GetTargets() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value
// and a boolean to check if the value has been set.
func (o *AuthorizationServer) GetTargetsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Targets, true
}

// SetTargets sets field value
func (o *AuthorizationServer) SetTargets(v []string) {
	o.Targets = v
}

// GetAuditLevel returns the AuditLevel field value if set, zero value otherwise.
func (o *AuthorizationServer) GetAuditLevel() AuditLevel {
	if o == nil || IsNil(o.AuditLevel) {
		var ret AuditLevel
		return ret
	}
	return *o.AuditLevel
}

// GetAuditLevelOk returns a tuple with the AuditLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServer) GetAuditLevelOk() (*AuditLevel, bool) {
	if o == nil || IsNil(o.AuditLevel) {
		return nil, false
	}
	return o.AuditLevel, true
}

// HasAuditLevel returns a boolean if a field has been set.
func (o *AuthorizationServer) HasAuditLevel() bool {
	if o != nil && !IsNil(o.AuditLevel) {
		return true
	}

	return false
}

// SetAuditLevel gets a reference to the given AuditLevel and assigns it to the AuditLevel field.
func (o *AuthorizationServer) SetAuditLevel(v AuditLevel) {
	o.AuditLevel = &v
}

// GetSecure returns the Secure field value if set, zero value otherwise.
func (o *AuthorizationServer) GetSecure() bool {
	if o == nil || IsNil(o.Secure) {
		var ret bool
		return ret
	}
	return *o.Secure
}

// GetSecureOk returns a tuple with the Secure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServer) GetSecureOk() (*bool, bool) {
	if o == nil || IsNil(o.Secure) {
		return nil, false
	}
	return o.Secure, true
}

// HasSecure returns a boolean if a field has been set.
func (o *AuthorizationServer) HasSecure() bool {
	if o != nil && !IsNil(o.Secure) {
		return true
	}

	return false
}

// SetSecure gets a reference to the given bool and assigns it to the Secure field.
func (o *AuthorizationServer) SetSecure(v bool) {
	o.Secure = &v
}

// GetTrustedCertificateGroupId returns the TrustedCertificateGroupId field value
func (o *AuthorizationServer) GetTrustedCertificateGroupId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TrustedCertificateGroupId
}

// GetTrustedCertificateGroupIdOk returns a tuple with the TrustedCertificateGroupId field value
// and a boolean to check if the value has been set.
func (o *AuthorizationServer) GetTrustedCertificateGroupIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrustedCertificateGroupId, true
}

// SetTrustedCertificateGroupId sets field value
func (o *AuthorizationServer) SetTrustedCertificateGroupId(v int64) {
	o.TrustedCertificateGroupId = v
}

// GetClientCredentials returns the ClientCredentials field value if set, zero value otherwise.
func (o *AuthorizationServer) GetClientCredentials() OAuthClientCredentials {
	if o == nil || IsNil(o.ClientCredentials) {
		var ret OAuthClientCredentials
		return ret
	}
	return *o.ClientCredentials
}

// GetClientCredentialsOk returns a tuple with the ClientCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServer) GetClientCredentialsOk() (*OAuthClientCredentials, bool) {
	if o == nil || IsNil(o.ClientCredentials) {
		return nil, false
	}
	return o.ClientCredentials, true
}

// HasClientCredentials returns a boolean if a field has been set.
func (o *AuthorizationServer) HasClientCredentials() bool {
	if o != nil && !IsNil(o.ClientCredentials) {
		return true
	}

	return false
}

// SetClientCredentials gets a reference to the given OAuthClientCredentials and assigns it to the ClientCredentials field.
func (o *AuthorizationServer) SetClientCredentials(v OAuthClientCredentials) {
	o.ClientCredentials = &v
}

// GetCacheTokens returns the CacheTokens field value if set, zero value otherwise.
func (o *AuthorizationServer) GetCacheTokens() bool {
	if o == nil || IsNil(o.CacheTokens) {
		var ret bool
		return ret
	}
	return *o.CacheTokens
}

// GetCacheTokensOk returns a tuple with the CacheTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServer) GetCacheTokensOk() (*bool, bool) {
	if o == nil || IsNil(o.CacheTokens) {
		return nil, false
	}
	return o.CacheTokens, true
}

// HasCacheTokens returns a boolean if a field has been set.
func (o *AuthorizationServer) HasCacheTokens() bool {
	if o != nil && !IsNil(o.CacheTokens) {
		return true
	}

	return false
}

// SetCacheTokens gets a reference to the given bool and assigns it to the CacheTokens field.
func (o *AuthorizationServer) SetCacheTokens(v bool) {
	o.CacheTokens = &v
}

// GetTokenTimeToLiveSeconds returns the TokenTimeToLiveSeconds field value if set, zero value otherwise.
func (o *AuthorizationServer) GetTokenTimeToLiveSeconds() int64 {
	if o == nil || IsNil(o.TokenTimeToLiveSeconds) {
		var ret int64
		return ret
	}
	return *o.TokenTimeToLiveSeconds
}

// GetTokenTimeToLiveSecondsOk returns a tuple with the TokenTimeToLiveSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServer) GetTokenTimeToLiveSecondsOk() (*int64, bool) {
	if o == nil || IsNil(o.TokenTimeToLiveSeconds) {
		return nil, false
	}
	return o.TokenTimeToLiveSeconds, true
}

// HasTokenTimeToLiveSeconds returns a boolean if a field has been set.
func (o *AuthorizationServer) HasTokenTimeToLiveSeconds() bool {
	if o != nil && !IsNil(o.TokenTimeToLiveSeconds) {
		return true
	}

	return false
}

// SetTokenTimeToLiveSeconds gets a reference to the given int64 and assigns it to the TokenTimeToLiveSeconds field.
func (o *AuthorizationServer) SetTokenTimeToLiveSeconds(v int64) {
	o.TokenTimeToLiveSeconds = &v
}

// GetSubjectAttributeName returns the SubjectAttributeName field value
func (o *AuthorizationServer) GetSubjectAttributeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubjectAttributeName
}

// GetSubjectAttributeNameOk returns a tuple with the SubjectAttributeName field value
// and a boolean to check if the value has been set.
func (o *AuthorizationServer) GetSubjectAttributeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubjectAttributeName, true
}

// SetSubjectAttributeName sets field value
func (o *AuthorizationServer) SetSubjectAttributeName(v string) {
	o.SubjectAttributeName = v
}

// GetIntrospectionEndpoint returns the IntrospectionEndpoint field value
func (o *AuthorizationServer) GetIntrospectionEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntrospectionEndpoint
}

// GetIntrospectionEndpointOk returns a tuple with the IntrospectionEndpoint field value
// and a boolean to check if the value has been set.
func (o *AuthorizationServer) GetIntrospectionEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntrospectionEndpoint, true
}

// SetIntrospectionEndpoint sets field value
func (o *AuthorizationServer) SetIntrospectionEndpoint(v string) {
	o.IntrospectionEndpoint = v
}

// GetTokenEndpoint returns the TokenEndpoint field value if set, zero value otherwise.
func (o *AuthorizationServer) GetTokenEndpoint() string {
	if o == nil || IsNil(o.TokenEndpoint) {
		var ret string
		return ret
	}
	return *o.TokenEndpoint
}

// GetTokenEndpointOk returns a tuple with the TokenEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServer) GetTokenEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.TokenEndpoint) {
		return nil, false
	}
	return o.TokenEndpoint, true
}

// HasTokenEndpoint returns a boolean if a field has been set.
func (o *AuthorizationServer) HasTokenEndpoint() bool {
	if o != nil && !IsNil(o.TokenEndpoint) {
		return true
	}

	return false
}

// SetTokenEndpoint gets a reference to the given string and assigns it to the TokenEndpoint field.
func (o *AuthorizationServer) SetTokenEndpoint(v string) {
	o.TokenEndpoint = &v
}

// GetSendAudience returns the SendAudience field value if set, zero value otherwise.
func (o *AuthorizationServer) GetSendAudience() bool {
	if o == nil || IsNil(o.SendAudience) {
		var ret bool
		return ret
	}
	return *o.SendAudience
}

// GetSendAudienceOk returns a tuple with the SendAudience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServer) GetSendAudienceOk() (*bool, bool) {
	if o == nil || IsNil(o.SendAudience) {
		return nil, false
	}
	return o.SendAudience, true
}

// HasSendAudience returns a boolean if a field has been set.
func (o *AuthorizationServer) HasSendAudience() bool {
	if o != nil && !IsNil(o.SendAudience) {
		return true
	}

	return false
}

// SetSendAudience gets a reference to the given bool and assigns it to the SendAudience field.
func (o *AuthorizationServer) SetSendAudience(v bool) {
	o.SendAudience = &v
}

// GetUseProxy returns the UseProxy field value if set, zero value otherwise.
func (o *AuthorizationServer) GetUseProxy() bool {
	if o == nil || IsNil(o.UseProxy) {
		var ret bool
		return ret
	}
	return *o.UseProxy
}

// GetUseProxyOk returns a tuple with the UseProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServer) GetUseProxyOk() (*bool, bool) {
	if o == nil || IsNil(o.UseProxy) {
		return nil, false
	}
	return o.UseProxy, true
}

// HasUseProxy returns a boolean if a field has been set.
func (o *AuthorizationServer) HasUseProxy() bool {
	if o != nil && !IsNil(o.UseProxy) {
		return true
	}

	return false
}

// SetUseProxy gets a reference to the given bool and assigns it to the UseProxy field.
func (o *AuthorizationServer) SetUseProxy(v bool) {
	o.UseProxy = &v
}

func (o AuthorizationServer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizationServer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["targets"] = o.Targets
	if !IsNil(o.AuditLevel) {
		toSerialize["auditLevel"] = o.AuditLevel
	}
	if !IsNil(o.Secure) {
		toSerialize["secure"] = o.Secure
	}
	toSerialize["trustedCertificateGroupId"] = o.TrustedCertificateGroupId
	if !IsNil(o.ClientCredentials) {
		toSerialize["clientCredentials"] = o.ClientCredentials
	}
	if !IsNil(o.CacheTokens) {
		toSerialize["cacheTokens"] = o.CacheTokens
	}
	if !IsNil(o.TokenTimeToLiveSeconds) {
		toSerialize["tokenTimeToLiveSeconds"] = o.TokenTimeToLiveSeconds
	}
	toSerialize["subjectAttributeName"] = o.SubjectAttributeName
	toSerialize["introspectionEndpoint"] = o.IntrospectionEndpoint
	if !IsNil(o.TokenEndpoint) {
		toSerialize["tokenEndpoint"] = o.TokenEndpoint
	}
	if !IsNil(o.SendAudience) {
		toSerialize["sendAudience"] = o.SendAudience
	}
	if !IsNil(o.UseProxy) {
		toSerialize["useProxy"] = o.UseProxy
	}
	return toSerialize, nil
}

type NullableAuthorizationServer struct {
	value *AuthorizationServer
	isSet bool
}

func (v NullableAuthorizationServer) Get() *AuthorizationServer {
	return v.value
}

func (v *NullableAuthorizationServer) Set(val *AuthorizationServer) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationServer) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationServer(val *AuthorizationServer) *NullableAuthorizationServer {
	return &NullableAuthorizationServer{value: val, isSet: true}
}

func (v NullableAuthorizationServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
