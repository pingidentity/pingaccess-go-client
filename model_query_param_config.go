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

// checks if the QueryParamConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryParamConfig{}

// QueryParamConfig Query parameter configuration settings to match requests to URLs with query parameters.
type QueryParamConfig struct {
	// Enable this setting to match requests to URLs without query parameters in addition URLs with query parameters.
	MatchesNoParams *bool `json:"matchesNoParams,omitempty"`
	// The query parameter name/value pairs.
	Params []QueryParamPair `json:"params,omitempty"`
}

// NewQueryParamConfig instantiates a new QueryParamConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryParamConfig() *QueryParamConfig {
	this := QueryParamConfig{}
	return &this
}

// NewQueryParamConfigWithDefaults instantiates a new QueryParamConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryParamConfigWithDefaults() *QueryParamConfig {
	this := QueryParamConfig{}
	return &this
}

// GetMatchesNoParams returns the MatchesNoParams field value if set, zero value otherwise.
func (o *QueryParamConfig) GetMatchesNoParams() bool {
	if o == nil || IsNil(o.MatchesNoParams) {
		var ret bool
		return ret
	}
	return *o.MatchesNoParams
}

// GetMatchesNoParamsOk returns a tuple with the MatchesNoParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryParamConfig) GetMatchesNoParamsOk() (*bool, bool) {
	if o == nil || IsNil(o.MatchesNoParams) {
		return nil, false
	}
	return o.MatchesNoParams, true
}

// HasMatchesNoParams returns a boolean if a field has been set.
func (o *QueryParamConfig) HasMatchesNoParams() bool {
	if o != nil && !IsNil(o.MatchesNoParams) {
		return true
	}

	return false
}

// SetMatchesNoParams gets a reference to the given bool and assigns it to the MatchesNoParams field.
func (o *QueryParamConfig) SetMatchesNoParams(v bool) {
	o.MatchesNoParams = &v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *QueryParamConfig) GetParams() []QueryParamPair {
	if o == nil || IsNil(o.Params) {
		var ret []QueryParamPair
		return ret
	}
	return o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryParamConfig) GetParamsOk() ([]QueryParamPair, bool) {
	if o == nil || IsNil(o.Params) {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *QueryParamConfig) HasParams() bool {
	if o != nil && !IsNil(o.Params) {
		return true
	}

	return false
}

// SetParams gets a reference to the given []QueryParamPair and assigns it to the Params field.
func (o *QueryParamConfig) SetParams(v []QueryParamPair) {
	o.Params = v
}

func (o QueryParamConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryParamConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MatchesNoParams) {
		toSerialize["matchesNoParams"] = o.MatchesNoParams
	}
	if !IsNil(o.Params) {
		toSerialize["params"] = o.Params
	}
	return toSerialize, nil
}

type NullableQueryParamConfig struct {
	value *QueryParamConfig
	isSet bool
}

func (v NullableQueryParamConfig) Get() *QueryParamConfig {
	return v.value
}

func (v *NullableQueryParamConfig) Set(val *QueryParamConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryParamConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryParamConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryParamConfig(val *QueryParamConfig) *NullableQueryParamConfig {
	return &NullableQueryParamConfig{value: val, isSet: true}
}

func (v NullableQueryParamConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryParamConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
