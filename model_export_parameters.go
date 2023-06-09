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

// checks if the ExportParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExportParameters{}

// ExportParameters The export parameters for a key pair. In the exported PEM file, the private key is protected with PBES2 encryption and AES.
type ExportParameters struct {
	Password HiddenField `json:"password"`
}

// NewExportParameters instantiates a new ExportParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportParameters(password HiddenField) *ExportParameters {
	this := ExportParameters{}
	this.Password = password
	return &this
}

// NewExportParametersWithDefaults instantiates a new ExportParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportParametersWithDefaults() *ExportParameters {
	this := ExportParameters{}
	return &this
}

// GetPassword returns the Password field value
func (o *ExportParameters) GetPassword() HiddenField {
	if o == nil {
		var ret HiddenField
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *ExportParameters) GetPasswordOk() (*HiddenField, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *ExportParameters) SetPassword(v HiddenField) {
	o.Password = v
}

func (o ExportParameters) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExportParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["password"] = o.Password
	return toSerialize, nil
}

type NullableExportParameters struct {
	value *ExportParameters
	isSet bool
}

func (v NullableExportParameters) Get() *ExportParameters {
	return v.value
}

func (v *NullableExportParameters) Set(val *ExportParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableExportParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableExportParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportParameters(val *ExportParameters) *NullableExportParameters {
	return &NullableExportParameters{value: val, isSet: true}
}

func (v NullableExportParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
