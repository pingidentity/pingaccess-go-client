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

// checks if the ChainCertificatesDoc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChainCertificatesDoc{}

// ChainCertificatesDoc struct for ChainCertificatesDoc
type ChainCertificatesDoc struct {
	AddChainCertificates []string `json:"addChainCertificates"`
}

// NewChainCertificatesDoc instantiates a new ChainCertificatesDoc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChainCertificatesDoc(addChainCertificates []string) *ChainCertificatesDoc {
	this := ChainCertificatesDoc{}
	this.AddChainCertificates = addChainCertificates
	return &this
}

// NewChainCertificatesDocWithDefaults instantiates a new ChainCertificatesDoc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChainCertificatesDocWithDefaults() *ChainCertificatesDoc {
	this := ChainCertificatesDoc{}
	return &this
}

// GetAddChainCertificates returns the AddChainCertificates field value
func (o *ChainCertificatesDoc) GetAddChainCertificates() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AddChainCertificates
}

// GetAddChainCertificatesOk returns a tuple with the AddChainCertificates field value
// and a boolean to check if the value has been set.
func (o *ChainCertificatesDoc) GetAddChainCertificatesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddChainCertificates, true
}

// SetAddChainCertificates sets field value
func (o *ChainCertificatesDoc) SetAddChainCertificates(v []string) {
	o.AddChainCertificates = v
}

func (o ChainCertificatesDoc) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChainCertificatesDoc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["addChainCertificates"] = o.AddChainCertificates
	return toSerialize, nil
}

type NullableChainCertificatesDoc struct {
	value *ChainCertificatesDoc
	isSet bool
}

func (v NullableChainCertificatesDoc) Get() *ChainCertificatesDoc {
	return v.value
}

func (v *NullableChainCertificatesDoc) Set(val *ChainCertificatesDoc) {
	v.value = val
	v.isSet = true
}

func (v NullableChainCertificatesDoc) IsSet() bool {
	return v.isSet
}

func (v *NullableChainCertificatesDoc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChainCertificatesDoc(val *ChainCertificatesDoc) *NullableChainCertificatesDoc {
	return &NullableChainCertificatesDoc{value: val, isSet: true}
}

func (v NullableChainCertificatesDoc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChainCertificatesDoc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
