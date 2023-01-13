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

// ResourceMatchingEvaluationOrder Specifies an ordering of Resource Matching Entries.
type ResourceMatchingEvaluationOrder struct {
	// Resource Matching Entries.
	Entries []ResourceMatchingEntry `json:"entries"`
}

// NewResourceMatchingEvaluationOrder instantiates a new ResourceMatchingEvaluationOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceMatchingEvaluationOrder(entries []ResourceMatchingEntry) *ResourceMatchingEvaluationOrder {
	this := ResourceMatchingEvaluationOrder{}
	this.Entries = entries
	return &this
}

// NewResourceMatchingEvaluationOrderWithDefaults instantiates a new ResourceMatchingEvaluationOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceMatchingEvaluationOrderWithDefaults() *ResourceMatchingEvaluationOrder {
	this := ResourceMatchingEvaluationOrder{}
	return &this
}

// GetEntries returns the Entries field value
func (o *ResourceMatchingEvaluationOrder) GetEntries() []ResourceMatchingEntry {
	if o == nil {
		var ret []ResourceMatchingEntry
		return ret
	}

	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value
// and a boolean to check if the value has been set.
func (o *ResourceMatchingEvaluationOrder) GetEntriesOk() ([]ResourceMatchingEntry, bool) {
	if o == nil {
    return nil, false
	}
	return o.Entries, true
}

// SetEntries sets field value
func (o *ResourceMatchingEvaluationOrder) SetEntries(v []ResourceMatchingEntry) {
	o.Entries = v
}

func (o ResourceMatchingEvaluationOrder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entries"] = o.Entries
	}
	return json.Marshal(toSerialize)
}

type NullableResourceMatchingEvaluationOrder struct {
	value *ResourceMatchingEvaluationOrder
	isSet bool
}

func (v NullableResourceMatchingEvaluationOrder) Get() *ResourceMatchingEvaluationOrder {
	return v.value
}

func (v *NullableResourceMatchingEvaluationOrder) Set(val *ResourceMatchingEvaluationOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceMatchingEvaluationOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceMatchingEvaluationOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceMatchingEvaluationOrder(val *ResourceMatchingEvaluationOrder) *NullableResourceMatchingEvaluationOrder {
	return &NullableResourceMatchingEvaluationOrder{value: val, isSet: true}
}

func (v NullableResourceMatchingEvaluationOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceMatchingEvaluationOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


