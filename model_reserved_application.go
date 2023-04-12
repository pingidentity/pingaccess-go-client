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

// checks if the ReservedApplication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReservedApplication{}

// ReservedApplication The reserved application.
type ReservedApplication struct {
	// The context root for the reserved application.
	ContextRoot string `json:"contextRoot"`
}

// NewReservedApplication instantiates a new ReservedApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReservedApplication(contextRoot string) *ReservedApplication {
	this := ReservedApplication{}
	this.ContextRoot = contextRoot
	return &this
}

// NewReservedApplicationWithDefaults instantiates a new ReservedApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReservedApplicationWithDefaults() *ReservedApplication {
	this := ReservedApplication{}
	return &this
}

// GetContextRoot returns the ContextRoot field value
func (o *ReservedApplication) GetContextRoot() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContextRoot
}

// GetContextRootOk returns a tuple with the ContextRoot field value
// and a boolean to check if the value has been set.
func (o *ReservedApplication) GetContextRootOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContextRoot, true
}

// SetContextRoot sets field value
func (o *ReservedApplication) SetContextRoot(v string) {
	o.ContextRoot = v
}

func (o ReservedApplication) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReservedApplication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contextRoot"] = o.ContextRoot
	return toSerialize, nil
}

type NullableReservedApplication struct {
	value *ReservedApplication
	isSet bool
}

func (v NullableReservedApplication) Get() *ReservedApplication {
	return v.value
}

func (v *NullableReservedApplication) Set(val *ReservedApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableReservedApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableReservedApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReservedApplication(val *ReservedApplication) *NullableReservedApplication {
	return &NullableReservedApplication{value: val, isSet: true}
}

func (v NullableReservedApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReservedApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
