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

// ConfiguredAuthorizationServerType the model 'ConfiguredAuthorizationServerType'
type ConfiguredAuthorizationServerType string

// List of ConfiguredAuthorizationServerType
const (
	CONFIGUREDAUTHORIZATIONSERVERTYPE_PINGFEDERATE_RUNTIME ConfiguredAuthorizationServerType = "PINGFEDERATE_RUNTIME"
	CONFIGUREDAUTHORIZATIONSERVERTYPE_PINGONE ConfiguredAuthorizationServerType = "PINGONE"
	CONFIGUREDAUTHORIZATIONSERVERTYPE_COMMON_AUTHZSERVER ConfiguredAuthorizationServerType = "COMMON_AUTHZSERVER"
	CONFIGUREDAUTHORIZATIONSERVERTYPE_ADMIN_TOKENPROVIDER ConfiguredAuthorizationServerType = "ADMIN_TOKENPROVIDER"
)

// All allowed values of ConfiguredAuthorizationServerType enum
var AllowedConfiguredAuthorizationServerTypeEnumValues = []ConfiguredAuthorizationServerType{
	"PINGFEDERATE_RUNTIME",
	"PINGONE",
	"COMMON_AUTHZSERVER",
	"ADMIN_TOKENPROVIDER",
}

func (v *ConfiguredAuthorizationServerType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConfiguredAuthorizationServerType(value)
	for _, existing := range AllowedConfiguredAuthorizationServerTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConfiguredAuthorizationServerType", value)
}

// NewConfiguredAuthorizationServerTypeFromValue returns a pointer to a valid ConfiguredAuthorizationServerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConfiguredAuthorizationServerTypeFromValue(v string) (*ConfiguredAuthorizationServerType, error) {
	ev := ConfiguredAuthorizationServerType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConfiguredAuthorizationServerType: valid values are %v", v, AllowedConfiguredAuthorizationServerTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConfiguredAuthorizationServerType) IsValid() bool {
	for _, existing := range AllowedConfiguredAuthorizationServerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConfiguredAuthorizationServerType value
func (v ConfiguredAuthorizationServerType) Ptr() *ConfiguredAuthorizationServerType {
	return &v
}

type NullableConfiguredAuthorizationServerType struct {
	value *ConfiguredAuthorizationServerType
	isSet bool
}

func (v NullableConfiguredAuthorizationServerType) Get() *ConfiguredAuthorizationServerType {
	return v.value
}

func (v *NullableConfiguredAuthorizationServerType) Set(val *ConfiguredAuthorizationServerType) {
	v.value = val
	v.isSet = true
}

func (v NullableConfiguredAuthorizationServerType) IsSet() bool {
	return v.isSet
}

func (v *NullableConfiguredAuthorizationServerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfiguredAuthorizationServerType(val *ConfiguredAuthorizationServerType) *NullableConfiguredAuthorizationServerType {
	return &NullableConfiguredAuthorizationServerType{value: val, isSet: true}
}

func (v NullableConfiguredAuthorizationServerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfiguredAuthorizationServerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

