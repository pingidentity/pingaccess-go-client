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

// ChallengeResponseFilter An authentication challenge response filter.
type ChallengeResponseFilter struct {
	// The class name of the challenge response filter.
	ClassName string `json:"className"`
	// The challenge response filter configuration.
	Configuration *string `json:"configuration,omitempty"`
}

// NewChallengeResponseFilter instantiates a new ChallengeResponseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChallengeResponseFilter(className string) *ChallengeResponseFilter {
	this := ChallengeResponseFilter{}
	this.ClassName = className
	return &this
}

// NewChallengeResponseFilterWithDefaults instantiates a new ChallengeResponseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChallengeResponseFilterWithDefaults() *ChallengeResponseFilter {
	this := ChallengeResponseFilter{}
	return &this
}

// GetClassName returns the ClassName field value
func (o *ChallengeResponseFilter) GetClassName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassName
}

// GetClassNameOk returns a tuple with the ClassName field value
// and a boolean to check if the value has been set.
func (o *ChallengeResponseFilter) GetClassNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ClassName, true
}

// SetClassName sets field value
func (o *ChallengeResponseFilter) SetClassName(v string) {
	o.ClassName = v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *ChallengeResponseFilter) GetConfiguration() string {
	if o == nil || isNil(o.Configuration) {
		var ret string
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChallengeResponseFilter) GetConfigurationOk() (*string, bool) {
	if o == nil || isNil(o.Configuration) {
    return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *ChallengeResponseFilter) HasConfiguration() bool {
	if o != nil && !isNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given string and assigns it to the Configuration field.
func (o *ChallengeResponseFilter) SetConfiguration(v string) {
	o.Configuration = &v
}

func (o ChallengeResponseFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["className"] = o.ClassName
	}
	if !isNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	return json.Marshal(toSerialize)
}

type NullableChallengeResponseFilter struct {
	value *ChallengeResponseFilter
	isSet bool
}

func (v NullableChallengeResponseFilter) Get() *ChallengeResponseFilter {
	return v.value
}

func (v *NullableChallengeResponseFilter) Set(val *ChallengeResponseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableChallengeResponseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableChallengeResponseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChallengeResponseFilter(val *ChallengeResponseFilter) *NullableChallengeResponseFilter {
	return &NullableChallengeResponseFilter{value: val, isSet: true}
}

func (v NullableChallengeResponseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChallengeResponseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


