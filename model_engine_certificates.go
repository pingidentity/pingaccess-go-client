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

// EngineCertificates A collection of engine certificates.
type EngineCertificates struct {
	// The actual list of engine certificates.
	Items []EngineCertificate `json:"items"`
}

// NewEngineCertificates instantiates a new EngineCertificates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEngineCertificates(items []EngineCertificate) *EngineCertificates {
	this := EngineCertificates{}
	this.Items = items
	return &this
}

// NewEngineCertificatesWithDefaults instantiates a new EngineCertificates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEngineCertificatesWithDefaults() *EngineCertificates {
	this := EngineCertificates{}
	return &this
}

// GetItems returns the Items field value
func (o *EngineCertificates) GetItems() []EngineCertificate {
	if o == nil {
		var ret []EngineCertificate
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *EngineCertificates) GetItemsOk() ([]EngineCertificate, bool) {
	if o == nil {
    return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *EngineCertificates) SetItems(v []EngineCertificate) {
	o.Items = v
}

func (o EngineCertificates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableEngineCertificates struct {
	value *EngineCertificates
	isSet bool
}

func (v NullableEngineCertificates) Get() *EngineCertificates {
	return v.value
}

func (v *NullableEngineCertificates) Set(val *EngineCertificates) {
	v.value = val
	v.isSet = true
}

func (v NullableEngineCertificates) IsSet() bool {
	return v.isSet
}

func (v *NullableEngineCertificates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEngineCertificates(val *EngineCertificates) *NullableEngineCertificates {
	return &NullableEngineCertificates{value: val, isSet: true}
}

func (v NullableEngineCertificates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEngineCertificates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

