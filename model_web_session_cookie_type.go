/*
PingAccess Administrative API

The PingAccess Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingAccess.

API version: 7.1.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// WebSessionCookieType the model 'WebSessionCookieType'
type WebSessionCookieType string

// List of WebSessionCookieType
const (
	WEBSESSIONCOOKIETYPE_ENCRYPTED WebSessionCookieType = "Encrypted"
	WEBSESSIONCOOKIETYPE_SIGNED WebSessionCookieType = "Signed"
)

// All allowed values of WebSessionCookieType enum
var AllowedWebSessionCookieTypeEnumValues = []WebSessionCookieType{
	"Encrypted",
	"Signed",
}

func (v *WebSessionCookieType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WebSessionCookieType(value)
	for _, existing := range AllowedWebSessionCookieTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WebSessionCookieType", value)
}

// NewWebSessionCookieTypeFromValue returns a pointer to a valid WebSessionCookieType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWebSessionCookieTypeFromValue(v string) (*WebSessionCookieType, error) {
	ev := WebSessionCookieType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WebSessionCookieType: valid values are %v", v, AllowedWebSessionCookieTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WebSessionCookieType) IsValid() bool {
	for _, existing := range AllowedWebSessionCookieTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WebSessionCookieType value
func (v WebSessionCookieType) Ptr() *WebSessionCookieType {
	return &v
}

type NullableWebSessionCookieType struct {
	value *WebSessionCookieType
	isSet bool
}

func (v NullableWebSessionCookieType) Get() *WebSessionCookieType {
	return v.value
}

func (v *NullableWebSessionCookieType) Set(val *WebSessionCookieType) {
	v.value = val
	v.isSet = true
}

func (v NullableWebSessionCookieType) IsSet() bool {
	return v.isSet
}

func (v *NullableWebSessionCookieType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebSessionCookieType(val *WebSessionCookieType) *NullableWebSessionCookieType {
	return &NullableWebSessionCookieType{value: val, isSet: true}
}

func (v NullableWebSessionCookieType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebSessionCookieType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
