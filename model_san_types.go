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

// checks if the SanTypes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SanTypes{}

// SanTypes struct for SanTypes
type SanTypes struct {
	// The actual list of available general names.
	Items []SanType `json:"items"`
}

// NewSanTypes instantiates a new SanTypes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSanTypes(items []SanType) *SanTypes {
	this := SanTypes{}
	this.Items = items
	return &this
}

// NewSanTypesWithDefaults instantiates a new SanTypes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSanTypesWithDefaults() *SanTypes {
	this := SanTypes{}
	return &this
}

// GetItems returns the Items field value
func (o *SanTypes) GetItems() []SanType {
	if o == nil {
		var ret []SanType
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *SanTypes) GetItemsOk() ([]SanType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *SanTypes) SetItems(v []SanType) {
	o.Items = v
}

func (o SanTypes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SanTypes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableSanTypes struct {
	value *SanTypes
	isSet bool
}

func (v NullableSanTypes) Get() *SanTypes {
	return v.value
}

func (v *NullableSanTypes) Set(val *SanTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableSanTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableSanTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSanTypes(val *SanTypes) *NullableSanTypes {
	return &NullableSanTypes{value: val, isSet: true}
}

func (v NullableSanTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSanTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
