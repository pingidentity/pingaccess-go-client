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

// ApiError An API error used by the PingAccess Administrative UI.
type ApiError struct {
	// The flash message displayed in the UI.
	Flash []string `json:"flash"`
	// The specific errors related to the current form.
	Form map[string][]string `json:"form"`
}

// NewApiError instantiates a new ApiError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiError(flash []string, form map[string][]string) *ApiError {
	this := ApiError{}
	this.Flash = flash
	this.Form = form
	return &this
}

// NewApiErrorWithDefaults instantiates a new ApiError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiErrorWithDefaults() *ApiError {
	this := ApiError{}
	return &this
}

// GetFlash returns the Flash field value
func (o *ApiError) GetFlash() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Flash
}

// GetFlashOk returns a tuple with the Flash field value
// and a boolean to check if the value has been set.
func (o *ApiError) GetFlashOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Flash, true
}

// SetFlash sets field value
func (o *ApiError) SetFlash(v []string) {
	o.Flash = v
}

// GetForm returns the Form field value
func (o *ApiError) GetForm() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}

	return o.Form
}

// GetFormOk returns a tuple with the Form field value
// and a boolean to check if the value has been set.
func (o *ApiError) GetFormOk() (*map[string][]string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Form, true
}

// SetForm sets field value
func (o *ApiError) SetForm(v map[string][]string) {
	o.Form = v
}

func (o ApiError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["flash"] = o.Flash
	}
	if true {
		toSerialize["form"] = o.Form
	}
	return json.Marshal(toSerialize)
}

type NullableApiError struct {
	value *ApiError
	isSet bool
}

func (v NullableApiError) Get() *ApiError {
	return v.value
}

func (v *NullableApiError) Set(val *ApiError) {
	v.value = val
	v.isSet = true
}

func (v NullableApiError) IsSet() bool {
	return v.isSet
}

func (v *NullableApiError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiError(val *ApiError) *NullableApiError {
	return &NullableApiError{value: val, isSet: true}
}

func (v NullableApiError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


