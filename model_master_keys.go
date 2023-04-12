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

// checks if the MasterKeys type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MasterKeys{}

// MasterKeys An encrypted master key.
type MasterKeys struct {
	// The encrypted master key.
	EncryptedValue []string `json:"encryptedValue,omitempty"`
	// The key identifier.
	KeyId *string `json:"keyId,omitempty"`
}

// NewMasterKeys instantiates a new MasterKeys object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMasterKeys() *MasterKeys {
	this := MasterKeys{}
	return &this
}

// NewMasterKeysWithDefaults instantiates a new MasterKeys object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMasterKeysWithDefaults() *MasterKeys {
	this := MasterKeys{}
	return &this
}

// GetEncryptedValue returns the EncryptedValue field value if set, zero value otherwise.
func (o *MasterKeys) GetEncryptedValue() []string {
	if o == nil || IsNil(o.EncryptedValue) {
		var ret []string
		return ret
	}
	return o.EncryptedValue
}

// GetEncryptedValueOk returns a tuple with the EncryptedValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterKeys) GetEncryptedValueOk() ([]string, bool) {
	if o == nil || IsNil(o.EncryptedValue) {
		return nil, false
	}
	return o.EncryptedValue, true
}

// HasEncryptedValue returns a boolean if a field has been set.
func (o *MasterKeys) HasEncryptedValue() bool {
	if o != nil && !IsNil(o.EncryptedValue) {
		return true
	}

	return false
}

// SetEncryptedValue gets a reference to the given []string and assigns it to the EncryptedValue field.
func (o *MasterKeys) SetEncryptedValue(v []string) {
	o.EncryptedValue = v
}

// GetKeyId returns the KeyId field value if set, zero value otherwise.
func (o *MasterKeys) GetKeyId() string {
	if o == nil || IsNil(o.KeyId) {
		var ret string
		return ret
	}
	return *o.KeyId
}

// GetKeyIdOk returns a tuple with the KeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterKeys) GetKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.KeyId) {
		return nil, false
	}
	return o.KeyId, true
}

// HasKeyId returns a boolean if a field has been set.
func (o *MasterKeys) HasKeyId() bool {
	if o != nil && !IsNil(o.KeyId) {
		return true
	}

	return false
}

// SetKeyId gets a reference to the given string and assigns it to the KeyId field.
func (o *MasterKeys) SetKeyId(v string) {
	o.KeyId = &v
}

func (o MasterKeys) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MasterKeys) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EncryptedValue) {
		toSerialize["encryptedValue"] = o.EncryptedValue
	}
	if !IsNil(o.KeyId) {
		toSerialize["keyId"] = o.KeyId
	}
	return toSerialize, nil
}

type NullableMasterKeys struct {
	value *MasterKeys
	isSet bool
}

func (v NullableMasterKeys) Get() *MasterKeys {
	return v.value
}

func (v *NullableMasterKeys) Set(val *MasterKeys) {
	v.value = val
	v.isSet = true
}

func (v NullableMasterKeys) IsSet() bool {
	return v.isSet
}

func (v *NullableMasterKeys) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMasterKeys(val *MasterKeys) *NullableMasterKeys {
	return &NullableMasterKeys{value: val, isSet: true}
}

func (v NullableMasterKeys) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMasterKeys) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
