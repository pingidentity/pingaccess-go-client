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

// checks if the OIDCProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OIDCProvider{}

// OIDCProvider The third-party OpenID Connect provider configuration.
type OIDCProvider struct {
	// The description of the third-party OpenID Connect provider.
	Description *string `json:"description,omitempty"`
	// The issuer of the third-party OpenID Connect provider.
	Issuer string `json:"issuer"`
	// The group of certificates to use when authenticating to third-party OpenID Connect provider.
	TrustedCertificateGroupId *int32 `json:"trustedCertificateGroupId,omitempty"`
	// True if a proxy should be used for HTTP or HTTPS requests.
	UseProxy   *bool       `json:"useProxy,omitempty"`
	AuditLevel *AuditLevel `json:"auditLevel,omitempty"`
	// The query parameters used by the authentication request to third-party OpenID Connect provider.
	QueryParameters []QueryParameter `json:"queryParameters,omitempty"`
	// True if single log off (SLO) should be used.
	UseSlo *bool `json:"useSlo,omitempty"`
	// Specifies whether the scopes in an access request should be limited to those advertised in the OIDC metadata.
	RequestSupportedScopesOnly *bool               `json:"requestSupportedScopesOnly,omitempty"`
	Plugin                     *OIDCProviderPlugin `json:"plugin,omitempty"`
}

// NewOIDCProvider instantiates a new OIDCProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOIDCProvider(issuer string) *OIDCProvider {
	this := OIDCProvider{}
	this.Issuer = issuer
	return &this
}

// NewOIDCProviderWithDefaults instantiates a new OIDCProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOIDCProviderWithDefaults() *OIDCProvider {
	this := OIDCProvider{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OIDCProvider) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCProvider) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OIDCProvider) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OIDCProvider) SetDescription(v string) {
	o.Description = &v
}

// GetIssuer returns the Issuer field value
func (o *OIDCProvider) GetIssuer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value
// and a boolean to check if the value has been set.
func (o *OIDCProvider) GetIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issuer, true
}

// SetIssuer sets field value
func (o *OIDCProvider) SetIssuer(v string) {
	o.Issuer = v
}

// GetTrustedCertificateGroupId returns the TrustedCertificateGroupId field value if set, zero value otherwise.
func (o *OIDCProvider) GetTrustedCertificateGroupId() int32 {
	if o == nil || IsNil(o.TrustedCertificateGroupId) {
		var ret int32
		return ret
	}
	return *o.TrustedCertificateGroupId
}

// GetTrustedCertificateGroupIdOk returns a tuple with the TrustedCertificateGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCProvider) GetTrustedCertificateGroupIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TrustedCertificateGroupId) {
		return nil, false
	}
	return o.TrustedCertificateGroupId, true
}

// HasTrustedCertificateGroupId returns a boolean if a field has been set.
func (o *OIDCProvider) HasTrustedCertificateGroupId() bool {
	if o != nil && !IsNil(o.TrustedCertificateGroupId) {
		return true
	}

	return false
}

// SetTrustedCertificateGroupId gets a reference to the given int32 and assigns it to the TrustedCertificateGroupId field.
func (o *OIDCProvider) SetTrustedCertificateGroupId(v int32) {
	o.TrustedCertificateGroupId = &v
}

// GetUseProxy returns the UseProxy field value if set, zero value otherwise.
func (o *OIDCProvider) GetUseProxy() bool {
	if o == nil || IsNil(o.UseProxy) {
		var ret bool
		return ret
	}
	return *o.UseProxy
}

// GetUseProxyOk returns a tuple with the UseProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCProvider) GetUseProxyOk() (*bool, bool) {
	if o == nil || IsNil(o.UseProxy) {
		return nil, false
	}
	return o.UseProxy, true
}

// HasUseProxy returns a boolean if a field has been set.
func (o *OIDCProvider) HasUseProxy() bool {
	if o != nil && !IsNil(o.UseProxy) {
		return true
	}

	return false
}

// SetUseProxy gets a reference to the given bool and assigns it to the UseProxy field.
func (o *OIDCProvider) SetUseProxy(v bool) {
	o.UseProxy = &v
}

// GetAuditLevel returns the AuditLevel field value if set, zero value otherwise.
func (o *OIDCProvider) GetAuditLevel() AuditLevel {
	if o == nil || IsNil(o.AuditLevel) {
		var ret AuditLevel
		return ret
	}
	return *o.AuditLevel
}

// GetAuditLevelOk returns a tuple with the AuditLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCProvider) GetAuditLevelOk() (*AuditLevel, bool) {
	if o == nil || IsNil(o.AuditLevel) {
		return nil, false
	}
	return o.AuditLevel, true
}

// HasAuditLevel returns a boolean if a field has been set.
func (o *OIDCProvider) HasAuditLevel() bool {
	if o != nil && !IsNil(o.AuditLevel) {
		return true
	}

	return false
}

// SetAuditLevel gets a reference to the given AuditLevel and assigns it to the AuditLevel field.
func (o *OIDCProvider) SetAuditLevel(v AuditLevel) {
	o.AuditLevel = &v
}

// GetQueryParameters returns the QueryParameters field value if set, zero value otherwise.
func (o *OIDCProvider) GetQueryParameters() []QueryParameter {
	if o == nil || IsNil(o.QueryParameters) {
		var ret []QueryParameter
		return ret
	}
	return o.QueryParameters
}

// GetQueryParametersOk returns a tuple with the QueryParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCProvider) GetQueryParametersOk() ([]QueryParameter, bool) {
	if o == nil || IsNil(o.QueryParameters) {
		return nil, false
	}
	return o.QueryParameters, true
}

// HasQueryParameters returns a boolean if a field has been set.
func (o *OIDCProvider) HasQueryParameters() bool {
	if o != nil && !IsNil(o.QueryParameters) {
		return true
	}

	return false
}

// SetQueryParameters gets a reference to the given []QueryParameter and assigns it to the QueryParameters field.
func (o *OIDCProvider) SetQueryParameters(v []QueryParameter) {
	o.QueryParameters = v
}

// GetUseSlo returns the UseSlo field value if set, zero value otherwise.
func (o *OIDCProvider) GetUseSlo() bool {
	if o == nil || IsNil(o.UseSlo) {
		var ret bool
		return ret
	}
	return *o.UseSlo
}

// GetUseSloOk returns a tuple with the UseSlo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCProvider) GetUseSloOk() (*bool, bool) {
	if o == nil || IsNil(o.UseSlo) {
		return nil, false
	}
	return o.UseSlo, true
}

// HasUseSlo returns a boolean if a field has been set.
func (o *OIDCProvider) HasUseSlo() bool {
	if o != nil && !IsNil(o.UseSlo) {
		return true
	}

	return false
}

// SetUseSlo gets a reference to the given bool and assigns it to the UseSlo field.
func (o *OIDCProvider) SetUseSlo(v bool) {
	o.UseSlo = &v
}

// GetRequestSupportedScopesOnly returns the RequestSupportedScopesOnly field value if set, zero value otherwise.
func (o *OIDCProvider) GetRequestSupportedScopesOnly() bool {
	if o == nil || IsNil(o.RequestSupportedScopesOnly) {
		var ret bool
		return ret
	}
	return *o.RequestSupportedScopesOnly
}

// GetRequestSupportedScopesOnlyOk returns a tuple with the RequestSupportedScopesOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCProvider) GetRequestSupportedScopesOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.RequestSupportedScopesOnly) {
		return nil, false
	}
	return o.RequestSupportedScopesOnly, true
}

// HasRequestSupportedScopesOnly returns a boolean if a field has been set.
func (o *OIDCProvider) HasRequestSupportedScopesOnly() bool {
	if o != nil && !IsNil(o.RequestSupportedScopesOnly) {
		return true
	}

	return false
}

// SetRequestSupportedScopesOnly gets a reference to the given bool and assigns it to the RequestSupportedScopesOnly field.
func (o *OIDCProvider) SetRequestSupportedScopesOnly(v bool) {
	o.RequestSupportedScopesOnly = &v
}

// GetPlugin returns the Plugin field value if set, zero value otherwise.
func (o *OIDCProvider) GetPlugin() OIDCProviderPlugin {
	if o == nil || IsNil(o.Plugin) {
		var ret OIDCProviderPlugin
		return ret
	}
	return *o.Plugin
}

// GetPluginOk returns a tuple with the Plugin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCProvider) GetPluginOk() (*OIDCProviderPlugin, bool) {
	if o == nil || IsNil(o.Plugin) {
		return nil, false
	}
	return o.Plugin, true
}

// HasPlugin returns a boolean if a field has been set.
func (o *OIDCProvider) HasPlugin() bool {
	if o != nil && !IsNil(o.Plugin) {
		return true
	}

	return false
}

// SetPlugin gets a reference to the given OIDCProviderPlugin and assigns it to the Plugin field.
func (o *OIDCProvider) SetPlugin(v OIDCProviderPlugin) {
	o.Plugin = &v
}

func (o OIDCProvider) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OIDCProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["issuer"] = o.Issuer
	if !IsNil(o.TrustedCertificateGroupId) {
		toSerialize["trustedCertificateGroupId"] = o.TrustedCertificateGroupId
	}
	if !IsNil(o.UseProxy) {
		toSerialize["useProxy"] = o.UseProxy
	}
	if !IsNil(o.AuditLevel) {
		toSerialize["auditLevel"] = o.AuditLevel
	}
	if !IsNil(o.QueryParameters) {
		toSerialize["queryParameters"] = o.QueryParameters
	}
	if !IsNil(o.UseSlo) {
		toSerialize["useSlo"] = o.UseSlo
	}
	if !IsNil(o.RequestSupportedScopesOnly) {
		toSerialize["requestSupportedScopesOnly"] = o.RequestSupportedScopesOnly
	}
	if !IsNil(o.Plugin) {
		toSerialize["plugin"] = o.Plugin
	}
	return toSerialize, nil
}

type NullableOIDCProvider struct {
	value *OIDCProvider
	isSet bool
}

func (v NullableOIDCProvider) Get() *OIDCProvider {
	return v.value
}

func (v *NullableOIDCProvider) Set(val *OIDCProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableOIDCProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableOIDCProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOIDCProvider(val *OIDCProvider) *NullableOIDCProvider {
	return &NullableOIDCProvider{value: val, isSet: true}
}

func (v NullableOIDCProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOIDCProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
