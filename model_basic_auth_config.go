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

// checks if the BasicAuthConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BasicAuthConfig{}

// BasicAuthConfig A basic authentication configuration.
type BasicAuthConfig struct {
	// This field is true if basic authentication to the Administrative API is enabled.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewBasicAuthConfig instantiates a new BasicAuthConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBasicAuthConfig() *BasicAuthConfig {
	this := BasicAuthConfig{}
	return &this
}

// NewBasicAuthConfigWithDefaults instantiates a new BasicAuthConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBasicAuthConfigWithDefaults() *BasicAuthConfig {
	this := BasicAuthConfig{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *BasicAuthConfig) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicAuthConfig) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *BasicAuthConfig) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *BasicAuthConfig) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o BasicAuthConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BasicAuthConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableBasicAuthConfig struct {
	value *BasicAuthConfig
	isSet bool
}

func (v NullableBasicAuthConfig) Get() *BasicAuthConfig {
	return v.value
}

func (v *NullableBasicAuthConfig) Set(val *BasicAuthConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableBasicAuthConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableBasicAuthConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBasicAuthConfig(val *BasicAuthConfig) *NullableBasicAuthConfig {
	return &NullableBasicAuthConfig{value: val, isSet: true}
}

func (v NullableBasicAuthConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBasicAuthConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
