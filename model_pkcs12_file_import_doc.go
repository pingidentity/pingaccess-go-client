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

// PKCS12FileImportDoc A PKCS#12 or PEM file import.
type PKCS12FileImportDoc struct {
	// The HSM Provider ID. The default value is 0 indicating an HSM is not used for this key pair.
	HsmProviderId int32 `json:"hsmProviderId"`
	Password HiddenField `json:"password"`
	// A unique alias name to identify the key pair. Special characters and spaces are allowed.
	Alias string `json:"alias"`
	// An array of chain certificates.
	ChainCertificates []string `json:"chainCertificates"`
	// Base-64 encoded PKCS12 or PEM file data. In BCFIPS mode, only PEM with PBES2 and AES or Triple DES encryption is accepted and 128-bit salt is required.
	FileData string `json:"fileData"`
}

// NewPKCS12FileImportDoc instantiates a new PKCS12FileImportDoc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPKCS12FileImportDoc(hsmProviderId int32, password HiddenField, alias string, chainCertificates []string, fileData string) *PKCS12FileImportDoc {
	this := PKCS12FileImportDoc{}
	this.HsmProviderId = hsmProviderId
	this.Password = password
	this.Alias = alias
	this.ChainCertificates = chainCertificates
	this.FileData = fileData
	return &this
}

// NewPKCS12FileImportDocWithDefaults instantiates a new PKCS12FileImportDoc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPKCS12FileImportDocWithDefaults() *PKCS12FileImportDoc {
	this := PKCS12FileImportDoc{}
	return &this
}

// GetHsmProviderId returns the HsmProviderId field value
func (o *PKCS12FileImportDoc) GetHsmProviderId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.HsmProviderId
}

// GetHsmProviderIdOk returns a tuple with the HsmProviderId field value
// and a boolean to check if the value has been set.
func (o *PKCS12FileImportDoc) GetHsmProviderIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.HsmProviderId, true
}

// SetHsmProviderId sets field value
func (o *PKCS12FileImportDoc) SetHsmProviderId(v int32) {
	o.HsmProviderId = v
}

// GetPassword returns the Password field value
func (o *PKCS12FileImportDoc) GetPassword() HiddenField {
	if o == nil {
		var ret HiddenField
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *PKCS12FileImportDoc) GetPasswordOk() (*HiddenField, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *PKCS12FileImportDoc) SetPassword(v HiddenField) {
	o.Password = v
}

// GetAlias returns the Alias field value
func (o *PKCS12FileImportDoc) GetAlias() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Alias
}

// GetAliasOk returns a tuple with the Alias field value
// and a boolean to check if the value has been set.
func (o *PKCS12FileImportDoc) GetAliasOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Alias, true
}

// SetAlias sets field value
func (o *PKCS12FileImportDoc) SetAlias(v string) {
	o.Alias = v
}

// GetChainCertificates returns the ChainCertificates field value
func (o *PKCS12FileImportDoc) GetChainCertificates() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ChainCertificates
}

// GetChainCertificatesOk returns a tuple with the ChainCertificates field value
// and a boolean to check if the value has been set.
func (o *PKCS12FileImportDoc) GetChainCertificatesOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ChainCertificates, true
}

// SetChainCertificates sets field value
func (o *PKCS12FileImportDoc) SetChainCertificates(v []string) {
	o.ChainCertificates = v
}

// GetFileData returns the FileData field value
func (o *PKCS12FileImportDoc) GetFileData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileData
}

// GetFileDataOk returns a tuple with the FileData field value
// and a boolean to check if the value has been set.
func (o *PKCS12FileImportDoc) GetFileDataOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.FileData, true
}

// SetFileData sets field value
func (o *PKCS12FileImportDoc) SetFileData(v string) {
	o.FileData = v
}

func (o PKCS12FileImportDoc) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["hsmProviderId"] = o.HsmProviderId
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["alias"] = o.Alias
	}
	if true {
		toSerialize["chainCertificates"] = o.ChainCertificates
	}
	if true {
		toSerialize["fileData"] = o.FileData
	}
	return json.Marshal(toSerialize)
}

type NullablePKCS12FileImportDoc struct {
	value *PKCS12FileImportDoc
	isSet bool
}

func (v NullablePKCS12FileImportDoc) Get() *PKCS12FileImportDoc {
	return v.value
}

func (v *NullablePKCS12FileImportDoc) Set(val *PKCS12FileImportDoc) {
	v.value = val
	v.isSet = true
}

func (v NullablePKCS12FileImportDoc) IsSet() bool {
	return v.isSet
}

func (v *NullablePKCS12FileImportDoc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePKCS12FileImportDoc(val *PKCS12FileImportDoc) *NullablePKCS12FileImportDoc {
	return &NullablePKCS12FileImportDoc{value: val, isSet: true}
}

func (v NullablePKCS12FileImportDoc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePKCS12FileImportDoc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

