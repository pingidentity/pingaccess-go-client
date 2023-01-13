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

// RuleSetSuccessCriteria struct for RuleSetSuccessCriteria
type RuleSetSuccessCriteria struct {
	Items []Item `json:"items"`
}

// NewRuleSetSuccessCriteria instantiates a new RuleSetSuccessCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleSetSuccessCriteria(items []Item) *RuleSetSuccessCriteria {
	this := RuleSetSuccessCriteria{}
	this.Items = items
	return &this
}

// NewRuleSetSuccessCriteriaWithDefaults instantiates a new RuleSetSuccessCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleSetSuccessCriteriaWithDefaults() *RuleSetSuccessCriteria {
	this := RuleSetSuccessCriteria{}
	return &this
}

// GetItems returns the Items field value
func (o *RuleSetSuccessCriteria) GetItems() []Item {
	if o == nil {
		var ret []Item
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *RuleSetSuccessCriteria) GetItemsOk() ([]Item, bool) {
	if o == nil {
    return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *RuleSetSuccessCriteria) SetItems(v []Item) {
	o.Items = v
}

func (o RuleSetSuccessCriteria) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableRuleSetSuccessCriteria struct {
	value *RuleSetSuccessCriteria
	isSet bool
}

func (v NullableRuleSetSuccessCriteria) Get() *RuleSetSuccessCriteria {
	return v.value
}

func (v *NullableRuleSetSuccessCriteria) Set(val *RuleSetSuccessCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleSetSuccessCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleSetSuccessCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleSetSuccessCriteria(val *RuleSetSuccessCriteria) *NullableRuleSetSuccessCriteria {
	return &NullableRuleSetSuccessCriteria{value: val, isSet: true}
}

func (v NullableRuleSetSuccessCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleSetSuccessCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

