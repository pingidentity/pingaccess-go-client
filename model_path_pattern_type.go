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

// PathPatternType the model 'PathPatternType'
type PathPatternType string

// List of PathPatternType
const (
	PATHPATTERNTYPE_WILDCARD PathPatternType = "WILDCARD"
	PATHPATTERNTYPE_REGEX    PathPatternType = "REGEX"
)

// All allowed values of PathPatternType enum
var AllowedPathPatternTypeEnumValues = []PathPatternType{
	"WILDCARD",
	"REGEX",
}

func (v *PathPatternType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PathPatternType(value)
	for _, existing := range AllowedPathPatternTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PathPatternType", value)
}

// NewPathPatternTypeFromValue returns a pointer to a valid PathPatternType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPathPatternTypeFromValue(v string) (*PathPatternType, error) {
	ev := PathPatternType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PathPatternType: valid values are %v", v, AllowedPathPatternTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PathPatternType) IsValid() bool {
	for _, existing := range AllowedPathPatternTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PathPatternType value
func (v PathPatternType) Ptr() *PathPatternType {
	return &v
}

type NullablePathPatternType struct {
	value *PathPatternType
	isSet bool
}

func (v NullablePathPatternType) Get() *PathPatternType {
	return v.value
}

func (v *NullablePathPatternType) Set(val *PathPatternType) {
	v.value = val
	v.isSet = true
}

func (v NullablePathPatternType) IsSet() bool {
	return v.isSet
}

func (v *NullablePathPatternType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePathPatternType(val *PathPatternType) *NullablePathPatternType {
	return &NullablePathPatternType{value: val, isSet: true}
}

func (v NullablePathPatternType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePathPatternType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
