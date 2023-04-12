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

// checks if the QueryParameter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryParameter{}

// QueryParameter A name-value pair of custom query parameters.
type QueryParameter struct {
	// The name of the query parameter.
	Name string `json:"name"`
	// The value of the query parameter.
	Value string `json:"value"`
}

// NewQueryParameter instantiates a new QueryParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryParameter(name string, value string) *QueryParameter {
	this := QueryParameter{}
	this.Name = name
	this.Value = value
	return &this
}

// NewQueryParameterWithDefaults instantiates a new QueryParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryParameterWithDefaults() *QueryParameter {
	this := QueryParameter{}
	return &this
}

// GetName returns the Name field value
func (o *QueryParameter) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *QueryParameter) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *QueryParameter) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *QueryParameter) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *QueryParameter) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *QueryParameter) SetValue(v string) {
	o.Value = v
}

func (o QueryParameter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryParameter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableQueryParameter struct {
	value *QueryParameter
	isSet bool
}

func (v NullableQueryParameter) Get() *QueryParameter {
	return v.value
}

func (v *NullableQueryParameter) Set(val *QueryParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryParameter(val *QueryParameter) *NullableQueryParameter {
	return &NullableQueryParameter{value: val, isSet: true}
}

func (v NullableQueryParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
