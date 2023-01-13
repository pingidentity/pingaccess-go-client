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

// Hash struct for Hash
type Hash struct {
	Algorithm *HashAlgorithm `json:"algorithm,omitempty"`
	HexValue *string `json:"hexValue,omitempty"`
}

// NewHash instantiates a new Hash object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHash() *Hash {
	this := Hash{}
	return &this
}

// NewHashWithDefaults instantiates a new Hash object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHashWithDefaults() *Hash {
	this := Hash{}
	return &this
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise.
func (o *Hash) GetAlgorithm() HashAlgorithm {
	if o == nil || isNil(o.Algorithm) {
		var ret HashAlgorithm
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hash) GetAlgorithmOk() (*HashAlgorithm, bool) {
	if o == nil || isNil(o.Algorithm) {
    return nil, false
	}
	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *Hash) HasAlgorithm() bool {
	if o != nil && !isNil(o.Algorithm) {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given HashAlgorithm and assigns it to the Algorithm field.
func (o *Hash) SetAlgorithm(v HashAlgorithm) {
	o.Algorithm = &v
}

// GetHexValue returns the HexValue field value if set, zero value otherwise.
func (o *Hash) GetHexValue() string {
	if o == nil || isNil(o.HexValue) {
		var ret string
		return ret
	}
	return *o.HexValue
}

// GetHexValueOk returns a tuple with the HexValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hash) GetHexValueOk() (*string, bool) {
	if o == nil || isNil(o.HexValue) {
    return nil, false
	}
	return o.HexValue, true
}

// HasHexValue returns a boolean if a field has been set.
func (o *Hash) HasHexValue() bool {
	if o != nil && !isNil(o.HexValue) {
		return true
	}

	return false
}

// SetHexValue gets a reference to the given string and assigns it to the HexValue field.
func (o *Hash) SetHexValue(v string) {
	o.HexValue = &v
}

func (o Hash) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Algorithm) {
		toSerialize["algorithm"] = o.Algorithm
	}
	if !isNil(o.HexValue) {
		toSerialize["hexValue"] = o.HexValue
	}
	return json.Marshal(toSerialize)
}

type NullableHash struct {
	value *Hash
	isSet bool
}

func (v NullableHash) Get() *Hash {
	return v.value
}

func (v *NullableHash) Set(val *Hash) {
	v.value = val
	v.isSet = true
}

func (v NullableHash) IsSet() bool {
	return v.isSet
}

func (v *NullableHash) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHash(val *Hash) *NullableHash {
	return &NullableHash{value: val, isSet: true}
}

func (v NullableHash) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHash) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

