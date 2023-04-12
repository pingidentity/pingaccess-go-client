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

// checks if the CSRResponseImportDoc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CSRResponseImportDoc{}

// CSRResponseImportDoc A CSR response.
type CSRResponseImportDoc struct {
	// The HSM Provider ID. The default value is 0 indicating an HSM is not used for this key pair.
	HsmProviderId int32 `json:"hsmProviderId"`
	// The CSR response data.
	FileData string `json:"fileData"`
	// A list of base64-encoded certificates to add to the key pair certificate chain.
	ChainCertificates []string `json:"chainCertificates"`
	// The ID of the trusted certificate group associated with the CSR response.
	TrustedCertGroupId int32 `json:"trustedCertGroupId"`
}

// NewCSRResponseImportDoc instantiates a new CSRResponseImportDoc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCSRResponseImportDoc(hsmProviderId int32, fileData string, chainCertificates []string, trustedCertGroupId int32) *CSRResponseImportDoc {
	this := CSRResponseImportDoc{}
	this.HsmProviderId = hsmProviderId
	this.FileData = fileData
	this.ChainCertificates = chainCertificates
	this.TrustedCertGroupId = trustedCertGroupId
	return &this
}

// NewCSRResponseImportDocWithDefaults instantiates a new CSRResponseImportDoc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCSRResponseImportDocWithDefaults() *CSRResponseImportDoc {
	this := CSRResponseImportDoc{}
	return &this
}

// GetHsmProviderId returns the HsmProviderId field value
func (o *CSRResponseImportDoc) GetHsmProviderId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.HsmProviderId
}

// GetHsmProviderIdOk returns a tuple with the HsmProviderId field value
// and a boolean to check if the value has been set.
func (o *CSRResponseImportDoc) GetHsmProviderIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HsmProviderId, true
}

// SetHsmProviderId sets field value
func (o *CSRResponseImportDoc) SetHsmProviderId(v int32) {
	o.HsmProviderId = v
}

// GetFileData returns the FileData field value
func (o *CSRResponseImportDoc) GetFileData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileData
}

// GetFileDataOk returns a tuple with the FileData field value
// and a boolean to check if the value has been set.
func (o *CSRResponseImportDoc) GetFileDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileData, true
}

// SetFileData sets field value
func (o *CSRResponseImportDoc) SetFileData(v string) {
	o.FileData = v
}

// GetChainCertificates returns the ChainCertificates field value
func (o *CSRResponseImportDoc) GetChainCertificates() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ChainCertificates
}

// GetChainCertificatesOk returns a tuple with the ChainCertificates field value
// and a boolean to check if the value has been set.
func (o *CSRResponseImportDoc) GetChainCertificatesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChainCertificates, true
}

// SetChainCertificates sets field value
func (o *CSRResponseImportDoc) SetChainCertificates(v []string) {
	o.ChainCertificates = v
}

// GetTrustedCertGroupId returns the TrustedCertGroupId field value
func (o *CSRResponseImportDoc) GetTrustedCertGroupId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TrustedCertGroupId
}

// GetTrustedCertGroupIdOk returns a tuple with the TrustedCertGroupId field value
// and a boolean to check if the value has been set.
func (o *CSRResponseImportDoc) GetTrustedCertGroupIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrustedCertGroupId, true
}

// SetTrustedCertGroupId sets field value
func (o *CSRResponseImportDoc) SetTrustedCertGroupId(v int32) {
	o.TrustedCertGroupId = v
}

func (o CSRResponseImportDoc) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CSRResponseImportDoc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hsmProviderId"] = o.HsmProviderId
	toSerialize["fileData"] = o.FileData
	toSerialize["chainCertificates"] = o.ChainCertificates
	toSerialize["trustedCertGroupId"] = o.TrustedCertGroupId
	return toSerialize, nil
}

type NullableCSRResponseImportDoc struct {
	value *CSRResponseImportDoc
	isSet bool
}

func (v NullableCSRResponseImportDoc) Get() *CSRResponseImportDoc {
	return v.value
}

func (v *NullableCSRResponseImportDoc) Set(val *CSRResponseImportDoc) {
	v.value = val
	v.isSet = true
}

func (v NullableCSRResponseImportDoc) IsSet() bool {
	return v.isSet
}

func (v *NullableCSRResponseImportDoc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCSRResponseImportDoc(val *CSRResponseImportDoc) *NullableCSRResponseImportDoc {
	return &NullableCSRResponseImportDoc{value: val, isSet: true}
}

func (v NullableCSRResponseImportDoc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCSRResponseImportDoc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
