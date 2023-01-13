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

// Resources A collection of resources.
type Resources struct {
	// The actual list of resources.
	Items []Resource `json:"items"`
}

// NewResources instantiates a new Resources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResources(items []Resource) *Resources {
	this := Resources{}
	this.Items = items
	return &this
}

// NewResourcesWithDefaults instantiates a new Resources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourcesWithDefaults() *Resources {
	this := Resources{}
	return &this
}

// GetItems returns the Items field value
func (o *Resources) GetItems() []Resource {
	if o == nil {
		var ret []Resource
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *Resources) GetItemsOk() ([]Resource, bool) {
	if o == nil {
    return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *Resources) SetItems(v []Resource) {
	o.Items = v
}

func (o Resources) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableResources struct {
	value *Resources
	isSet bool
}

func (v NullableResources) Get() *Resources {
	return v.value
}

func (v *NullableResources) Set(val *Resources) {
	v.value = val
	v.isSet = true
}

func (v NullableResources) IsSet() bool {
	return v.isSet
}

func (v *NullableResources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResources(val *Resources) *NullableResources {
	return &NullableResources{value: val, isSet: true}
}

func (v NullableResources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

