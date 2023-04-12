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

// checks if the Hash type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Hash{}

// Hash struct for Hash
type Hash struct {
	Algorithm *HashAlgorithm `json:"algorithm,omitempty"`
	HexValue  *string        `json:"hexValue,omitempty"`
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
	if o == nil || IsNil(o.Algorithm) {
		var ret HashAlgorithm
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hash) GetAlgorithmOk() (*HashAlgorithm, bool) {
	if o == nil || IsNil(o.Algorithm) {
		return nil, false
	}
	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *Hash) HasAlgorithm() bool {
	if o != nil && !IsNil(o.Algorithm) {
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
	if o == nil || IsNil(o.HexValue) {
		var ret string
		return ret
	}
	return *o.HexValue
}

// GetHexValueOk returns a tuple with the HexValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hash) GetHexValueOk() (*string, bool) {
	if o == nil || IsNil(o.HexValue) {
		return nil, false
	}
	return o.HexValue, true
}

// HasHexValue returns a boolean if a field has been set.
func (o *Hash) HasHexValue() bool {
	if o != nil && !IsNil(o.HexValue) {
		return true
	}

	return false
}

// SetHexValue gets a reference to the given string and assigns it to the HexValue field.
func (o *Hash) SetHexValue(v string) {
	o.HexValue = &v
}

func (o Hash) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Hash) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Algorithm) {
		toSerialize["algorithm"] = o.Algorithm
	}
	if !IsNil(o.HexValue) {
		toSerialize["hexValue"] = o.HexValue
	}
	return toSerialize, nil
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
