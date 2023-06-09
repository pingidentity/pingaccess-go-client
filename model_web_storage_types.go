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

// checks if the WebStorageTypes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebStorageTypes{}

// WebStorageTypes struct for WebStorageTypes
type WebStorageTypes struct {
	// The actual list of valid web storage types.
	Items []Item `json:"items"`
}

// NewWebStorageTypes instantiates a new WebStorageTypes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebStorageTypes(items []Item) *WebStorageTypes {
	this := WebStorageTypes{}
	this.Items = items
	return &this
}

// NewWebStorageTypesWithDefaults instantiates a new WebStorageTypes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebStorageTypesWithDefaults() *WebStorageTypes {
	this := WebStorageTypes{}
	return &this
}

// GetItems returns the Items field value
func (o *WebStorageTypes) GetItems() []Item {
	if o == nil {
		var ret []Item
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *WebStorageTypes) GetItemsOk() ([]Item, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *WebStorageTypes) SetItems(v []Item) {
	o.Items = v
}

func (o WebStorageTypes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebStorageTypes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableWebStorageTypes struct {
	value *WebStorageTypes
	isSet bool
}

func (v NullableWebStorageTypes) Get() *WebStorageTypes {
	return v.value
}

func (v *NullableWebStorageTypes) Set(val *WebStorageTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableWebStorageTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableWebStorageTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebStorageTypes(val *WebStorageTypes) *NullableWebStorageTypes {
	return &NullableWebStorageTypes{value: val, isSet: true}
}

func (v NullableWebStorageTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebStorageTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
