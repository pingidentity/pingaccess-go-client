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

// RequestPreservationType the model 'RequestPreservationType'
type RequestPreservationType string

// List of RequestPreservationType
const (
	REQUESTPRESERVATIONTYPE_NONE RequestPreservationType = "None"
	REQUESTPRESERVATIONTYPE_POST RequestPreservationType = "POST"
	REQUESTPRESERVATIONTYPE_ALL  RequestPreservationType = "All"
)

// All allowed values of RequestPreservationType enum
var AllowedRequestPreservationTypeEnumValues = []RequestPreservationType{
	"None",
	"POST",
	"All",
}

func (v *RequestPreservationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RequestPreservationType(value)
	for _, existing := range AllowedRequestPreservationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RequestPreservationType", value)
}

// NewRequestPreservationTypeFromValue returns a pointer to a valid RequestPreservationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRequestPreservationTypeFromValue(v string) (*RequestPreservationType, error) {
	ev := RequestPreservationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RequestPreservationType: valid values are %v", v, AllowedRequestPreservationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RequestPreservationType) IsValid() bool {
	for _, existing := range AllowedRequestPreservationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RequestPreservationType value
func (v RequestPreservationType) Ptr() *RequestPreservationType {
	return &v
}

type NullableRequestPreservationType struct {
	value *RequestPreservationType
	isSet bool
}

func (v NullableRequestPreservationType) Get() *RequestPreservationType {
	return v.value
}

func (v *NullableRequestPreservationType) Set(val *RequestPreservationType) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestPreservationType) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestPreservationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestPreservationType(val *RequestPreservationType) *NullableRequestPreservationType {
	return &NullableRequestPreservationType{value: val, isSet: true}
}

func (v NullableRequestPreservationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestPreservationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
