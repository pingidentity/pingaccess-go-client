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

// checks if the RuleSetElementTypes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleSetElementTypes{}

// RuleSetElementTypes struct for RuleSetElementTypes
type RuleSetElementTypes struct {
	// The actual list of element types.
	Items []Item `json:"items"`
}

// NewRuleSetElementTypes instantiates a new RuleSetElementTypes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleSetElementTypes(items []Item) *RuleSetElementTypes {
	this := RuleSetElementTypes{}
	this.Items = items
	return &this
}

// NewRuleSetElementTypesWithDefaults instantiates a new RuleSetElementTypes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleSetElementTypesWithDefaults() *RuleSetElementTypes {
	this := RuleSetElementTypes{}
	return &this
}

// GetItems returns the Items field value
func (o *RuleSetElementTypes) GetItems() []Item {
	if o == nil {
		var ret []Item
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *RuleSetElementTypes) GetItemsOk() ([]Item, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *RuleSetElementTypes) SetItems(v []Item) {
	o.Items = v
}

func (o RuleSetElementTypes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuleSetElementTypes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableRuleSetElementTypes struct {
	value *RuleSetElementTypes
	isSet bool
}

func (v NullableRuleSetElementTypes) Get() *RuleSetElementTypes {
	return v.value
}

func (v *NullableRuleSetElementTypes) Set(val *RuleSetElementTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleSetElementTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleSetElementTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleSetElementTypes(val *RuleSetElementTypes) *NullableRuleSetElementTypes {
	return &NullableRuleSetElementTypes{value: val, isSet: true}
}

func (v NullableRuleSetElementTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleSetElementTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
