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

// HiddenField A hidden field.
type HiddenField struct {
	// The encrypted value of the field, as originally returned by the API.
	EncryptedValue *string `json:"encryptedValue,omitempty"`
	// The value of the field. This field takes precedence over the encryptedValue field, if both are specified.
	Value *string `json:"value,omitempty"`
}

// NewHiddenField instantiates a new HiddenField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHiddenField() *HiddenField {
	this := HiddenField{}
	return &this
}

// NewHiddenFieldWithDefaults instantiates a new HiddenField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHiddenFieldWithDefaults() *HiddenField {
	this := HiddenField{}
	return &this
}

// GetEncryptedValue returns the EncryptedValue field value if set, zero value otherwise.
func (o *HiddenField) GetEncryptedValue() string {
	if o == nil || isNil(o.EncryptedValue) {
		var ret string
		return ret
	}
	return *o.EncryptedValue
}

// GetEncryptedValueOk returns a tuple with the EncryptedValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HiddenField) GetEncryptedValueOk() (*string, bool) {
	if o == nil || isNil(o.EncryptedValue) {
    return nil, false
	}
	return o.EncryptedValue, true
}

// HasEncryptedValue returns a boolean if a field has been set.
func (o *HiddenField) HasEncryptedValue() bool {
	if o != nil && !isNil(o.EncryptedValue) {
		return true
	}

	return false
}

// SetEncryptedValue gets a reference to the given string and assigns it to the EncryptedValue field.
func (o *HiddenField) SetEncryptedValue(v string) {
	o.EncryptedValue = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *HiddenField) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HiddenField) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *HiddenField) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *HiddenField) SetValue(v string) {
	o.Value = &v
}

func (o HiddenField) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EncryptedValue) {
		toSerialize["encryptedValue"] = o.EncryptedValue
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableHiddenField struct {
	value *HiddenField
	isSet bool
}

func (v NullableHiddenField) Get() *HiddenField {
	return v.value
}

func (v *NullableHiddenField) Set(val *HiddenField) {
	v.value = val
	v.isSet = true
}

func (v NullableHiddenField) IsSet() bool {
	return v.isSet
}

func (v *NullableHiddenField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHiddenField(val *HiddenField) *NullableHiddenField {
	return &NullableHiddenField{value: val, isSet: true}
}

func (v NullableHiddenField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHiddenField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

