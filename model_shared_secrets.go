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

// SharedSecrets A collection of shared secrets.
type SharedSecrets struct {
	// The actual list of shared secrets.
	Items []SharedSecret `json:"items"`
}

// NewSharedSecrets instantiates a new SharedSecrets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSharedSecrets(items []SharedSecret) *SharedSecrets {
	this := SharedSecrets{}
	this.Items = items
	return &this
}

// NewSharedSecretsWithDefaults instantiates a new SharedSecrets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharedSecretsWithDefaults() *SharedSecrets {
	this := SharedSecrets{}
	return &this
}

// GetItems returns the Items field value
func (o *SharedSecrets) GetItems() []SharedSecret {
	if o == nil {
		var ret []SharedSecret
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *SharedSecrets) GetItemsOk() ([]SharedSecret, bool) {
	if o == nil {
    return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *SharedSecrets) SetItems(v []SharedSecret) {
	o.Items = v
}

func (o SharedSecrets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableSharedSecrets struct {
	value *SharedSecrets
	isSet bool
}

func (v NullableSharedSecrets) Get() *SharedSecrets {
	return v.value
}

func (v *NullableSharedSecrets) Set(val *SharedSecrets) {
	v.value = val
	v.isSet = true
}

func (v NullableSharedSecrets) IsSet() bool {
	return v.isSet
}

func (v *NullableSharedSecrets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharedSecrets(val *SharedSecrets) *NullableSharedSecrets {
	return &NullableSharedSecrets{value: val, isSet: true}
}

func (v NullableSharedSecrets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharedSecrets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


