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

// checks if the PingFederateMetadataRuntime type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PingFederateMetadataRuntime{}

// PingFederateMetadataRuntime A PingFederate configuration.
type PingFederateMetadataRuntime struct {
	// The description of the PingFederate Runtime token provider.
	Description *string `json:"description,omitempty"`
	// The issuer url of the PingFederate token provider.
	Issuer string `json:"issuer"`
	// The group of certificates to use when authenticating to PingFederate.
	TrustedCertificateGroupId *int64 `json:"trustedCertificateGroupId,omitempty"`
	// Set to true if a proxy should be used for HTTP or HTTPS requests.
	UseProxy *bool `json:"useProxy,omitempty"`
	// Set to true if OIDC single log out should be used on the /pa/oidc/logout on the engines.
	UseSlo *bool `json:"useSlo,omitempty"`
	// The url of the PingFederate STS token-to-token exchange endpoint that is used for token mediation. Specify if it is being served from a different host/port than the issuer is. Otherwise, it is assumed to be {{issuer}}/pf/sts.wst.
	StsTokenExchangeEndpoint *string `json:"stsTokenExchangeEndpoint,omitempty"`
	// Set to true if HTTP communications to PingFederate should not perform hostname verification of the certificate.
	SkipHostnameVerification *bool `json:"skipHostnameVerification,omitempty"`
}

// NewPingFederateMetadataRuntime instantiates a new PingFederateMetadataRuntime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPingFederateMetadataRuntime(issuer string) *PingFederateMetadataRuntime {
	this := PingFederateMetadataRuntime{}
	this.Issuer = issuer
	return &this
}

// NewPingFederateMetadataRuntimeWithDefaults instantiates a new PingFederateMetadataRuntime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPingFederateMetadataRuntimeWithDefaults() *PingFederateMetadataRuntime {
	this := PingFederateMetadataRuntime{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PingFederateMetadataRuntime) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingFederateMetadataRuntime) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PingFederateMetadataRuntime) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PingFederateMetadataRuntime) SetDescription(v string) {
	o.Description = &v
}

// GetIssuer returns the Issuer field value
func (o *PingFederateMetadataRuntime) GetIssuer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value
// and a boolean to check if the value has been set.
func (o *PingFederateMetadataRuntime) GetIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issuer, true
}

// SetIssuer sets field value
func (o *PingFederateMetadataRuntime) SetIssuer(v string) {
	o.Issuer = v
}

// GetTrustedCertificateGroupId returns the TrustedCertificateGroupId field value if set, zero value otherwise.
func (o *PingFederateMetadataRuntime) GetTrustedCertificateGroupId() int64 {
	if o == nil || IsNil(o.TrustedCertificateGroupId) {
		var ret int64
		return ret
	}
	return *o.TrustedCertificateGroupId
}

// GetTrustedCertificateGroupIdOk returns a tuple with the TrustedCertificateGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingFederateMetadataRuntime) GetTrustedCertificateGroupIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TrustedCertificateGroupId) {
		return nil, false
	}
	return o.TrustedCertificateGroupId, true
}

// HasTrustedCertificateGroupId returns a boolean if a field has been set.
func (o *PingFederateMetadataRuntime) HasTrustedCertificateGroupId() bool {
	if o != nil && !IsNil(o.TrustedCertificateGroupId) {
		return true
	}

	return false
}

// SetTrustedCertificateGroupId gets a reference to the given int64 and assigns it to the TrustedCertificateGroupId field.
func (o *PingFederateMetadataRuntime) SetTrustedCertificateGroupId(v int64) {
	o.TrustedCertificateGroupId = &v
}

// GetUseProxy returns the UseProxy field value if set, zero value otherwise.
func (o *PingFederateMetadataRuntime) GetUseProxy() bool {
	if o == nil || IsNil(o.UseProxy) {
		var ret bool
		return ret
	}
	return *o.UseProxy
}

// GetUseProxyOk returns a tuple with the UseProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingFederateMetadataRuntime) GetUseProxyOk() (*bool, bool) {
	if o == nil || IsNil(o.UseProxy) {
		return nil, false
	}
	return o.UseProxy, true
}

// HasUseProxy returns a boolean if a field has been set.
func (o *PingFederateMetadataRuntime) HasUseProxy() bool {
	if o != nil && !IsNil(o.UseProxy) {
		return true
	}

	return false
}

// SetUseProxy gets a reference to the given bool and assigns it to the UseProxy field.
func (o *PingFederateMetadataRuntime) SetUseProxy(v bool) {
	o.UseProxy = &v
}

// GetUseSlo returns the UseSlo field value if set, zero value otherwise.
func (o *PingFederateMetadataRuntime) GetUseSlo() bool {
	if o == nil || IsNil(o.UseSlo) {
		var ret bool
		return ret
	}
	return *o.UseSlo
}

// GetUseSloOk returns a tuple with the UseSlo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingFederateMetadataRuntime) GetUseSloOk() (*bool, bool) {
	if o == nil || IsNil(o.UseSlo) {
		return nil, false
	}
	return o.UseSlo, true
}

// HasUseSlo returns a boolean if a field has been set.
func (o *PingFederateMetadataRuntime) HasUseSlo() bool {
	if o != nil && !IsNil(o.UseSlo) {
		return true
	}

	return false
}

// SetUseSlo gets a reference to the given bool and assigns it to the UseSlo field.
func (o *PingFederateMetadataRuntime) SetUseSlo(v bool) {
	o.UseSlo = &v
}

// GetStsTokenExchangeEndpoint returns the StsTokenExchangeEndpoint field value if set, zero value otherwise.
func (o *PingFederateMetadataRuntime) GetStsTokenExchangeEndpoint() string {
	if o == nil || IsNil(o.StsTokenExchangeEndpoint) {
		var ret string
		return ret
	}
	return *o.StsTokenExchangeEndpoint
}

// GetStsTokenExchangeEndpointOk returns a tuple with the StsTokenExchangeEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingFederateMetadataRuntime) GetStsTokenExchangeEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.StsTokenExchangeEndpoint) {
		return nil, false
	}
	return o.StsTokenExchangeEndpoint, true
}

// HasStsTokenExchangeEndpoint returns a boolean if a field has been set.
func (o *PingFederateMetadataRuntime) HasStsTokenExchangeEndpoint() bool {
	if o != nil && !IsNil(o.StsTokenExchangeEndpoint) {
		return true
	}

	return false
}

// SetStsTokenExchangeEndpoint gets a reference to the given string and assigns it to the StsTokenExchangeEndpoint field.
func (o *PingFederateMetadataRuntime) SetStsTokenExchangeEndpoint(v string) {
	o.StsTokenExchangeEndpoint = &v
}

// GetSkipHostnameVerification returns the SkipHostnameVerification field value if set, zero value otherwise.
func (o *PingFederateMetadataRuntime) GetSkipHostnameVerification() bool {
	if o == nil || IsNil(o.SkipHostnameVerification) {
		var ret bool
		return ret
	}
	return *o.SkipHostnameVerification
}

// GetSkipHostnameVerificationOk returns a tuple with the SkipHostnameVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingFederateMetadataRuntime) GetSkipHostnameVerificationOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipHostnameVerification) {
		return nil, false
	}
	return o.SkipHostnameVerification, true
}

// HasSkipHostnameVerification returns a boolean if a field has been set.
func (o *PingFederateMetadataRuntime) HasSkipHostnameVerification() bool {
	if o != nil && !IsNil(o.SkipHostnameVerification) {
		return true
	}

	return false
}

// SetSkipHostnameVerification gets a reference to the given bool and assigns it to the SkipHostnameVerification field.
func (o *PingFederateMetadataRuntime) SetSkipHostnameVerification(v bool) {
	o.SkipHostnameVerification = &v
}

func (o PingFederateMetadataRuntime) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PingFederateMetadataRuntime) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.UseSlo) {
		toSerialize["useSlo"] = o.UseSlo
	}
	if !IsNil(o.StsTokenExchangeEndpoint) {
		toSerialize["stsTokenExchangeEndpoint"] = o.StsTokenExchangeEndpoint
	}
	if !IsNil(o.SkipHostnameVerification) {
		toSerialize["skipHostnameVerification"] = o.SkipHostnameVerification
	}
	return toSerialize, nil
}

type NullablePingFederateMetadataRuntime struct {
	value *PingFederateMetadataRuntime
	isSet bool
}

func (v NullablePingFederateMetadataRuntime) Get() *PingFederateMetadataRuntime {
	return v.value
}

func (v *NullablePingFederateMetadataRuntime) Set(val *PingFederateMetadataRuntime) {
	v.value = val
	v.isSet = true
}

func (v NullablePingFederateMetadataRuntime) IsSet() bool {
	return v.isSet
}

func (v *NullablePingFederateMetadataRuntime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePingFederateMetadataRuntime(val *PingFederateMetadataRuntime) *NullablePingFederateMetadataRuntime {
	return &NullablePingFederateMetadataRuntime{value: val, isSet: true}
}

func (v NullablePingFederateMetadataRuntime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePingFederateMetadataRuntime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
