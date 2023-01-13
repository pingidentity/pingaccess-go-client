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

// Methods struct for Methods
type Methods struct {
	Items []Method `json:"items"`
}

// NewMethods instantiates a new Methods object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMethods(items []Method) *Methods {
	this := Methods{}
	this.Items = items
	return &this
}

// NewMethodsWithDefaults instantiates a new Methods object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMethodsWithDefaults() *Methods {
	this := Methods{}
	return &this
}

// GetItems returns the Items field value
func (o *Methods) GetItems() []Method {
	if o == nil {
		var ret []Method
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *Methods) GetItemsOk() ([]Method, bool) {
	if o == nil {
    return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *Methods) SetItems(v []Method) {
	o.Items = v
}

func (o Methods) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableMethods struct {
	value *Methods
	isSet bool
}

func (v NullableMethods) Get() *Methods {
	return v.value
}

func (v *NullableMethods) Set(val *Methods) {
	v.value = val
	v.isSet = true
}

func (v NullableMethods) IsSet() bool {
	return v.isSet
}

func (v *NullableMethods) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMethods(val *Methods) *NullableMethods {
	return &NullableMethods{value: val, isSet: true}
}

func (v NullableMethods) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMethods) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

