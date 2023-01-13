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

// Applications A collection of applications.
type Applications struct {
	// An array of applications.
	Items []Application `json:"items"`
}

// NewApplications instantiates a new Applications object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplications(items []Application) *Applications {
	this := Applications{}
	this.Items = items
	return &this
}

// NewApplicationsWithDefaults instantiates a new Applications object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationsWithDefaults() *Applications {
	this := Applications{}
	return &this
}

// GetItems returns the Items field value
func (o *Applications) GetItems() []Application {
	if o == nil {
		var ret []Application
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *Applications) GetItemsOk() ([]Application, bool) {
	if o == nil {
    return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *Applications) SetItems(v []Application) {
	o.Items = v
}

func (o Applications) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableApplications struct {
	value *Applications
	isSet bool
}

func (v NullableApplications) Get() *Applications {
	return v.value
}

func (v *NullableApplications) Set(val *Applications) {
	v.value = val
	v.isSet = true
}

func (v NullableApplications) IsSet() bool {
	return v.isSet
}

func (v *NullableApplications) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplications(val *Applications) *NullableApplications {
	return &NullableApplications{value: val, isSet: true}
}

func (v NullableApplications) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplications) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

