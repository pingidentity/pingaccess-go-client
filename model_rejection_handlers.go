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

// checks if the RejectionHandlers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RejectionHandlers{}

// RejectionHandlers A collection of rejection handlers.
type RejectionHandlers struct {
	// The actual list of rejection handlers.
	Items []RejectionHandler `json:"items"`
}

// NewRejectionHandlers instantiates a new RejectionHandlers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRejectionHandlers(items []RejectionHandler) *RejectionHandlers {
	this := RejectionHandlers{}
	this.Items = items
	return &this
}

// NewRejectionHandlersWithDefaults instantiates a new RejectionHandlers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRejectionHandlersWithDefaults() *RejectionHandlers {
	this := RejectionHandlers{}
	return &this
}

// GetItems returns the Items field value
func (o *RejectionHandlers) GetItems() []RejectionHandler {
	if o == nil {
		var ret []RejectionHandler
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *RejectionHandlers) GetItemsOk() ([]RejectionHandler, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *RejectionHandlers) SetItems(v []RejectionHandler) {
	o.Items = v
}

func (o RejectionHandlers) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RejectionHandlers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableRejectionHandlers struct {
	value *RejectionHandlers
	isSet bool
}

func (v NullableRejectionHandlers) Get() *RejectionHandlers {
	return v.value
}

func (v *NullableRejectionHandlers) Set(val *RejectionHandlers) {
	v.value = val
	v.isSet = true
}

func (v NullableRejectionHandlers) IsSet() bool {
	return v.isSet
}

func (v *NullableRejectionHandlers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRejectionHandlers(val *RejectionHandlers) *NullableRejectionHandlers {
	return &NullableRejectionHandlers{value: val, isSet: true}
}

func (v NullableRejectionHandlers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRejectionHandlers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
