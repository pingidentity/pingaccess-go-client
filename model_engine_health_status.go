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

// EngineHealthStatus struct for EngineHealthStatus
type EngineHealthStatus struct {
	EnginesStatus map[string]EngineInfo `json:"enginesStatus"`
	CurrentServerTime int64 `json:"currentServerTime"`
}

// NewEngineHealthStatus instantiates a new EngineHealthStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEngineHealthStatus(enginesStatus map[string]EngineInfo, currentServerTime int64) *EngineHealthStatus {
	this := EngineHealthStatus{}
	this.EnginesStatus = enginesStatus
	this.CurrentServerTime = currentServerTime
	return &this
}

// NewEngineHealthStatusWithDefaults instantiates a new EngineHealthStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEngineHealthStatusWithDefaults() *EngineHealthStatus {
	this := EngineHealthStatus{}
	return &this
}

// GetEnginesStatus returns the EnginesStatus field value
func (o *EngineHealthStatus) GetEnginesStatus() map[string]EngineInfo {
	if o == nil {
		var ret map[string]EngineInfo
		return ret
	}

	return o.EnginesStatus
}

// GetEnginesStatusOk returns a tuple with the EnginesStatus field value
// and a boolean to check if the value has been set.
func (o *EngineHealthStatus) GetEnginesStatusOk() (*map[string]EngineInfo, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EnginesStatus, true
}

// SetEnginesStatus sets field value
func (o *EngineHealthStatus) SetEnginesStatus(v map[string]EngineInfo) {
	o.EnginesStatus = v
}

// GetCurrentServerTime returns the CurrentServerTime field value
func (o *EngineHealthStatus) GetCurrentServerTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CurrentServerTime
}

// GetCurrentServerTimeOk returns a tuple with the CurrentServerTime field value
// and a boolean to check if the value has been set.
func (o *EngineHealthStatus) GetCurrentServerTimeOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CurrentServerTime, true
}

// SetCurrentServerTime sets field value
func (o *EngineHealthStatus) SetCurrentServerTime(v int64) {
	o.CurrentServerTime = v
}

func (o EngineHealthStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enginesStatus"] = o.EnginesStatus
	}
	if true {
		toSerialize["currentServerTime"] = o.CurrentServerTime
	}
	return json.Marshal(toSerialize)
}

type NullableEngineHealthStatus struct {
	value *EngineHealthStatus
	isSet bool
}

func (v NullableEngineHealthStatus) Get() *EngineHealthStatus {
	return v.value
}

func (v *NullableEngineHealthStatus) Set(val *EngineHealthStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEngineHealthStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEngineHealthStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEngineHealthStatus(val *EngineHealthStatus) *NullableEngineHealthStatus {
	return &NullableEngineHealthStatus{value: val, isSet: true}
}

func (v NullableEngineHealthStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEngineHealthStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

