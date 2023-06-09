/*
Administrative API Documentation

The PingAccess Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingAccess as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API.

API version: 7.3.0.2-SNAPSHOT
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
	"fmt"
)

// UnknownResourceMode the model 'UnknownResourceMode'
type UnknownResourceMode string

// List of UnknownResourceMode
const (
	UNKNOWNRESOURCEMODE_DENY        UnknownResourceMode = "DENY"
	UNKNOWNRESOURCEMODE_PASSTHROUGH UnknownResourceMode = "PASSTHROUGH"
)

// All allowed values of UnknownResourceMode enum
var AllowedUnknownResourceModeEnumValues = []UnknownResourceMode{
	"DENY",
	"PASSTHROUGH",
}

func (v *UnknownResourceMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UnknownResourceMode(value)
	for _, existing := range AllowedUnknownResourceModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UnknownResourceMode", value)
}

// NewUnknownResourceModeFromValue returns a pointer to a valid UnknownResourceMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUnknownResourceModeFromValue(v string) (*UnknownResourceMode, error) {
	ev := UnknownResourceMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UnknownResourceMode: valid values are %v", v, AllowedUnknownResourceModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UnknownResourceMode) IsValid() bool {
	for _, existing := range AllowedUnknownResourceModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UnknownResourceMode value
func (v UnknownResourceMode) Ptr() *UnknownResourceMode {
	return &v
}

type NullableUnknownResourceMode struct {
	value *UnknownResourceMode
	isSet bool
}

func (v NullableUnknownResourceMode) Get() *UnknownResourceMode {
	return v.value
}

func (v *NullableUnknownResourceMode) Set(val *UnknownResourceMode) {
	v.value = val
	v.isSet = true
}

func (v NullableUnknownResourceMode) IsSet() bool {
	return v.isSet
}

func (v *NullableUnknownResourceMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnknownResourceMode(val *UnknownResourceMode) *NullableUnknownResourceMode {
	return &NullableUnknownResourceMode{value: val, isSet: true}
}

func (v NullableUnknownResourceMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnknownResourceMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
