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

// Role the model 'Role'
type Role string

// List of Role
const (
	ROLE_ADMINISTRATOR  Role = "ADMINISTRATOR"
	ROLE_AUDITOR        Role = "AUDITOR"
	ROLE_PLATFORM_ADMIN Role = "PLATFORM_ADMIN"
)

// All allowed values of Role enum
var AllowedRoleEnumValues = []Role{
	"ADMINISTRATOR",
	"AUDITOR",
	"PLATFORM_ADMIN",
}

func (v *Role) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Role(value)
	for _, existing := range AllowedRoleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Role", value)
}

// NewRoleFromValue returns a pointer to a valid Role
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRoleFromValue(v string) (*Role, error) {
	ev := Role(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Role: valid values are %v", v, AllowedRoleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Role) IsValid() bool {
	for _, existing := range AllowedRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Role value
func (v Role) Ptr() *Role {
	return &v
}

type NullableRole struct {
	value *Role
	isSet bool
}

func (v NullableRole) Get() *Role {
	return v.value
}

func (v *NullableRole) Set(val *Role) {
	v.value = val
	v.isSet = true
}

func (v NullableRole) IsSet() bool {
	return v.isSet
}

func (v *NullableRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRole(val *Role) *NullableRole {
	return &NullableRole{value: val, isSet: true}
}

func (v NullableRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
