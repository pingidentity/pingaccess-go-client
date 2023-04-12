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

// checks if the NewKeyPairConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NewKeyPairConfig{}

// NewKeyPairConfig A new key pair.
type NewKeyPairConfig struct {
	// The ID for the key pair. If not specified, an ID will be automatically assigned.
	Id *int32 `json:"id,omitempty"`
	// A unique alias name to identify the key pair. Special characters and spaces are allowed.
	Alias string `json:"alias"`
	// The common name (CN) identifying the certificate.
	CommonName string `json:"commonName"`
	// The organization (O) or company name creating the certificate.
	Organization string `json:"organization"`
	// The specific unit within the organization (OU).
	OrganizationUnit string `json:"organizationUnit"`
	// The city or other primary location (L) where the company operates.
	City string `json:"city"`
	// The state (ST) or other political unit encompassing the location.
	State string `json:"state"`
	// The country (C) where the company is based, using two capital letters.
	Country string `json:"country"`
	// The number of days the certificate is valid.
	ValidDays int32 `json:"validDays"`
	// The key algorithm to use to generate a key.
	KeyAlgorithm string `json:"keyAlgorithm"`
	// The Signature Algorithm to use for the key.
	SignatureAlgorithm *string `json:"signatureAlgorithm,omitempty"`
	// Any additional DNS names or IP addresses that are valid for this certificate.
	SubjectAlternativeNames []GeneralName `json:"subjectAlternativeNames,omitempty"`
	// The number of bits used in the key.  Choices depend on selected key algorithm.
	KeySize int32 `json:"keySize"`
	// The HSM Provider ID. The default value is 0 indicating an HSM is not used for this key pair.
	HsmProviderId int32 `json:"hsmProviderId"`
}

// NewNewKeyPairConfig instantiates a new NewKeyPairConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewKeyPairConfig(alias string, commonName string, organization string, organizationUnit string, city string, state string, country string, validDays int32, keyAlgorithm string, keySize int32, hsmProviderId int32) *NewKeyPairConfig {
	this := NewKeyPairConfig{}
	this.Alias = alias
	this.CommonName = commonName
	this.Organization = organization
	this.OrganizationUnit = organizationUnit
	this.City = city
	this.State = state
	this.Country = country
	this.ValidDays = validDays
	this.KeyAlgorithm = keyAlgorithm
	this.KeySize = keySize
	this.HsmProviderId = hsmProviderId
	return &this
}

// NewNewKeyPairConfigWithDefaults instantiates a new NewKeyPairConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewKeyPairConfigWithDefaults() *NewKeyPairConfig {
	this := NewKeyPairConfig{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NewKeyPairConfig) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewKeyPairConfig) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NewKeyPairConfig) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *NewKeyPairConfig) SetId(v int32) {
	o.Id = &v
}

// GetAlias returns the Alias field value
func (o *NewKeyPairConfig) GetAlias() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Alias
}

// GetAliasOk returns a tuple with the Alias field value
// and a boolean to check if the value has been set.
func (o *NewKeyPairConfig) GetAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alias, true
}

// SetAlias sets field value
func (o *NewKeyPairConfig) SetAlias(v string) {
	o.Alias = v
}

// GetCommonName returns the CommonName field value
func (o *NewKeyPairConfig) GetCommonName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommonName
}

// GetCommonNameOk returns a tuple with the CommonName field value
// and a boolean to check if the value has been set.
func (o *NewKeyPairConfig) GetCommonNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommonName, true
}

// SetCommonName sets field value
func (o *NewKeyPairConfig) SetCommonName(v string) {
	o.CommonName = v
}

// GetOrganization returns the Organization field value
func (o *NewKeyPairConfig) GetOrganization() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value
// and a boolean to check if the value has been set.
func (o *NewKeyPairConfig) GetOrganizationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Organization, true
}

// SetOrganization sets field value
func (o *NewKeyPairConfig) SetOrganization(v string) {
	o.Organization = v
}

// GetOrganizationUnit returns the OrganizationUnit field value
func (o *NewKeyPairConfig) GetOrganizationUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationUnit
}

// GetOrganizationUnitOk returns a tuple with the OrganizationUnit field value
// and a boolean to check if the value has been set.
func (o *NewKeyPairConfig) GetOrganizationUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationUnit, true
}

// SetOrganizationUnit sets field value
func (o *NewKeyPairConfig) SetOrganizationUnit(v string) {
	o.OrganizationUnit = v
}

// GetCity returns the City field value
func (o *NewKeyPairConfig) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *NewKeyPairConfig) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *NewKeyPairConfig) SetCity(v string) {
	o.City = v
}

// GetState returns the State field value
func (o *NewKeyPairConfig) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *NewKeyPairConfig) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *NewKeyPairConfig) SetState(v string) {
	o.State = v
}

// GetCountry returns the Country field value
func (o *NewKeyPairConfig) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *NewKeyPairConfig) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *NewKeyPairConfig) SetCountry(v string) {
	o.Country = v
}

// GetValidDays returns the ValidDays field value
func (o *NewKeyPairConfig) GetValidDays() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ValidDays
}

// GetValidDaysOk returns a tuple with the ValidDays field value
// and a boolean to check if the value has been set.
func (o *NewKeyPairConfig) GetValidDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidDays, true
}

// SetValidDays sets field value
func (o *NewKeyPairConfig) SetValidDays(v int32) {
	o.ValidDays = v
}

// GetKeyAlgorithm returns the KeyAlgorithm field value
func (o *NewKeyPairConfig) GetKeyAlgorithm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyAlgorithm
}

// GetKeyAlgorithmOk returns a tuple with the KeyAlgorithm field value
// and a boolean to check if the value has been set.
func (o *NewKeyPairConfig) GetKeyAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyAlgorithm, true
}

// SetKeyAlgorithm sets field value
func (o *NewKeyPairConfig) SetKeyAlgorithm(v string) {
	o.KeyAlgorithm = v
}

// GetSignatureAlgorithm returns the SignatureAlgorithm field value if set, zero value otherwise.
func (o *NewKeyPairConfig) GetSignatureAlgorithm() string {
	if o == nil || IsNil(o.SignatureAlgorithm) {
		var ret string
		return ret
	}
	return *o.SignatureAlgorithm
}

// GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewKeyPairConfig) GetSignatureAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.SignatureAlgorithm) {
		return nil, false
	}
	return o.SignatureAlgorithm, true
}

// HasSignatureAlgorithm returns a boolean if a field has been set.
func (o *NewKeyPairConfig) HasSignatureAlgorithm() bool {
	if o != nil && !IsNil(o.SignatureAlgorithm) {
		return true
	}

	return false
}

// SetSignatureAlgorithm gets a reference to the given string and assigns it to the SignatureAlgorithm field.
func (o *NewKeyPairConfig) SetSignatureAlgorithm(v string) {
	o.SignatureAlgorithm = &v
}

// GetSubjectAlternativeNames returns the SubjectAlternativeNames field value if set, zero value otherwise.
func (o *NewKeyPairConfig) GetSubjectAlternativeNames() []GeneralName {
	if o == nil || IsNil(o.SubjectAlternativeNames) {
		var ret []GeneralName
		return ret
	}
	return o.SubjectAlternativeNames
}

// GetSubjectAlternativeNamesOk returns a tuple with the SubjectAlternativeNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewKeyPairConfig) GetSubjectAlternativeNamesOk() ([]GeneralName, bool) {
	if o == nil || IsNil(o.SubjectAlternativeNames) {
		return nil, false
	}
	return o.SubjectAlternativeNames, true
}

// HasSubjectAlternativeNames returns a boolean if a field has been set.
func (o *NewKeyPairConfig) HasSubjectAlternativeNames() bool {
	if o != nil && !IsNil(o.SubjectAlternativeNames) {
		return true
	}

	return false
}

// SetSubjectAlternativeNames gets a reference to the given []GeneralName and assigns it to the SubjectAlternativeNames field.
func (o *NewKeyPairConfig) SetSubjectAlternativeNames(v []GeneralName) {
	o.SubjectAlternativeNames = v
}

// GetKeySize returns the KeySize field value
func (o *NewKeyPairConfig) GetKeySize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.KeySize
}

// GetKeySizeOk returns a tuple with the KeySize field value
// and a boolean to check if the value has been set.
func (o *NewKeyPairConfig) GetKeySizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeySize, true
}

// SetKeySize sets field value
func (o *NewKeyPairConfig) SetKeySize(v int32) {
	o.KeySize = v
}

// GetHsmProviderId returns the HsmProviderId field value
func (o *NewKeyPairConfig) GetHsmProviderId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.HsmProviderId
}

// GetHsmProviderIdOk returns a tuple with the HsmProviderId field value
// and a boolean to check if the value has been set.
func (o *NewKeyPairConfig) GetHsmProviderIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HsmProviderId, true
}

// SetHsmProviderId sets field value
func (o *NewKeyPairConfig) SetHsmProviderId(v int32) {
	o.HsmProviderId = v
}

func (o NewKeyPairConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewKeyPairConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["alias"] = o.Alias
	toSerialize["commonName"] = o.CommonName
	toSerialize["organization"] = o.Organization
	toSerialize["organizationUnit"] = o.OrganizationUnit
	toSerialize["city"] = o.City
	toSerialize["state"] = o.State
	toSerialize["country"] = o.Country
	toSerialize["validDays"] = o.ValidDays
	toSerialize["keyAlgorithm"] = o.KeyAlgorithm
	if !IsNil(o.SignatureAlgorithm) {
		toSerialize["signatureAlgorithm"] = o.SignatureAlgorithm
	}
	if !IsNil(o.SubjectAlternativeNames) {
		toSerialize["subjectAlternativeNames"] = o.SubjectAlternativeNames
	}
	toSerialize["keySize"] = o.KeySize
	toSerialize["hsmProviderId"] = o.HsmProviderId
	return toSerialize, nil
}

type NullableNewKeyPairConfig struct {
	value *NewKeyPairConfig
	isSet bool
}

func (v NullableNewKeyPairConfig) Get() *NewKeyPairConfig {
	return v.value
}

func (v *NullableNewKeyPairConfig) Set(val *NewKeyPairConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableNewKeyPairConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableNewKeyPairConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewKeyPairConfig(val *NewKeyPairConfig) *NullableNewKeyPairConfig {
	return &NullableNewKeyPairConfig{value: val, isSet: true}
}

func (v NullableNewKeyPairConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewKeyPairConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
