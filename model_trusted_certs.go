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

// TrustedCerts A collection of trusted certificates.
type TrustedCerts struct {
	// A collection of trusted certificate items.
	Items []TrustedCert `json:"items"`
}

// NewTrustedCerts instantiates a new TrustedCerts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrustedCerts(items []TrustedCert) *TrustedCerts {
	this := TrustedCerts{}
	this.Items = items
	return &this
}

// NewTrustedCertsWithDefaults instantiates a new TrustedCerts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrustedCertsWithDefaults() *TrustedCerts {
	this := TrustedCerts{}
	return &this
}

// GetItems returns the Items field value
func (o *TrustedCerts) GetItems() []TrustedCert {
	if o == nil {
		var ret []TrustedCert
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *TrustedCerts) GetItemsOk() ([]TrustedCert, bool) {
	if o == nil {
    return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *TrustedCerts) SetItems(v []TrustedCert) {
	o.Items = v
}

func (o TrustedCerts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableTrustedCerts struct {
	value *TrustedCerts
	isSet bool
}

func (v NullableTrustedCerts) Get() *TrustedCerts {
	return v.value
}

func (v *NullableTrustedCerts) Set(val *TrustedCerts) {
	v.value = val
	v.isSet = true
}

func (v NullableTrustedCerts) IsSet() bool {
	return v.isSet
}

func (v *NullableTrustedCerts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrustedCerts(val *TrustedCerts) *NullableTrustedCerts {
	return &NullableTrustedCerts{value: val, isSet: true}
}

func (v NullableTrustedCerts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrustedCerts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

