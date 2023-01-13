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

// Engines A collection of engines.
type Engines struct {
	// An array of engines.
	Items []Engine `json:"items"`
}

// NewEngines instantiates a new Engines object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEngines(items []Engine) *Engines {
	this := Engines{}
	this.Items = items
	return &this
}

// NewEnginesWithDefaults instantiates a new Engines object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnginesWithDefaults() *Engines {
	this := Engines{}
	return &this
}

// GetItems returns the Items field value
func (o *Engines) GetItems() []Engine {
	if o == nil {
		var ret []Engine
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *Engines) GetItemsOk() ([]Engine, bool) {
	if o == nil {
    return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *Engines) SetItems(v []Engine) {
	o.Items = v
}

func (o Engines) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableEngines struct {
	value *Engines
	isSet bool
}

func (v NullableEngines) Get() *Engines {
	return v.value
}

func (v *NullableEngines) Set(val *Engines) {
	v.value = val
	v.isSet = true
}

func (v NullableEngines) IsSet() bool {
	return v.isSet
}

func (v *NullableEngines) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEngines(val *Engines) *NullableEngines {
	return &NullableEngines{value: val, isSet: true}
}

func (v NullableEngines) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEngines) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


