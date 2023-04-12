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

// checks if the X509FileImportDoc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &X509FileImportDoc{}

// X509FileImportDoc An X.509 certificate.
type X509FileImportDoc struct {
	// The alias for the certificate.
	Alias string `json:"alias"`
	// The base64-encoded data of the certificate.
	FileData string `json:"fileData"`
}

// NewX509FileImportDoc instantiates a new X509FileImportDoc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewX509FileImportDoc(alias string, fileData string) *X509FileImportDoc {
	this := X509FileImportDoc{}
	this.Alias = alias
	this.FileData = fileData
	return &this
}

// NewX509FileImportDocWithDefaults instantiates a new X509FileImportDoc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewX509FileImportDocWithDefaults() *X509FileImportDoc {
	this := X509FileImportDoc{}
	return &this
}

// GetAlias returns the Alias field value
func (o *X509FileImportDoc) GetAlias() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Alias
}

// GetAliasOk returns a tuple with the Alias field value
// and a boolean to check if the value has been set.
func (o *X509FileImportDoc) GetAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alias, true
}

// SetAlias sets field value
func (o *X509FileImportDoc) SetAlias(v string) {
	o.Alias = v
}

// GetFileData returns the FileData field value
func (o *X509FileImportDoc) GetFileData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileData
}

// GetFileDataOk returns a tuple with the FileData field value
// and a boolean to check if the value has been set.
func (o *X509FileImportDoc) GetFileDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileData, true
}

// SetFileData sets field value
func (o *X509FileImportDoc) SetFileData(v string) {
	o.FileData = v
}

func (o X509FileImportDoc) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o X509FileImportDoc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["alias"] = o.Alias
	toSerialize["fileData"] = o.FileData
	return toSerialize, nil
}

type NullableX509FileImportDoc struct {
	value *X509FileImportDoc
	isSet bool
}

func (v NullableX509FileImportDoc) Get() *X509FileImportDoc {
	return v.value
}

func (v *NullableX509FileImportDoc) Set(val *X509FileImportDoc) {
	v.value = val
	v.isSet = true
}

func (v NullableX509FileImportDoc) IsSet() bool {
	return v.isSet
}

func (v *NullableX509FileImportDoc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableX509FileImportDoc(val *X509FileImportDoc) *NullableX509FileImportDoc {
	return &NullableX509FileImportDoc{value: val, isSet: true}
}

func (v NullableX509FileImportDoc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableX509FileImportDoc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
