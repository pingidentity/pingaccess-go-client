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

// ConfigurationType the model 'ConfigurationType'
type ConfigurationType string

// List of ConfigurationType
const (
	CONFIGURATIONTYPE_TEXT ConfigurationType = "TEXT"
	CONFIGURATIONTYPE_TEXTAREA ConfigurationType = "TEXTAREA"
	CONFIGURATIONTYPE_TIME ConfigurationType = "TIME"
	CONFIGURATIONTYPE_SELECT ConfigurationType = "SELECT"
	CONFIGURATIONTYPE_GROOVY ConfigurationType = "GROOVY"
	CONFIGURATIONTYPE_CONCEALED ConfigurationType = "CONCEALED"
	CONFIGURATIONTYPE_LIST ConfigurationType = "LIST"
	CONFIGURATIONTYPE_TABLE ConfigurationType = "TABLE"
	CONFIGURATIONTYPE_CHECKBOX ConfigurationType = "CHECKBOX"
	CONFIGURATIONTYPE_AUTOCOMPLETEOPEN ConfigurationType = "AUTOCOMPLETEOPEN"
	CONFIGURATIONTYPE_AUTOCOMPLETECLOSED ConfigurationType = "AUTOCOMPLETECLOSED"
	CONFIGURATIONTYPE_COMPOSITE ConfigurationType = "COMPOSITE"
	CONFIGURATIONTYPE_RADIO_BUTTON ConfigurationType = "RADIO_BUTTON"
)

// All allowed values of ConfigurationType enum
var AllowedConfigurationTypeEnumValues = []ConfigurationType{
	"TEXT",
	"TEXTAREA",
	"TIME",
	"SELECT",
	"GROOVY",
	"CONCEALED",
	"LIST",
	"TABLE",
	"CHECKBOX",
	"AUTOCOMPLETEOPEN",
	"AUTOCOMPLETECLOSED",
	"COMPOSITE",
	"RADIO_BUTTON",
}

func (v *ConfigurationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConfigurationType(value)
	for _, existing := range AllowedConfigurationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConfigurationType", value)
}

// NewConfigurationTypeFromValue returns a pointer to a valid ConfigurationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConfigurationTypeFromValue(v string) (*ConfigurationType, error) {
	ev := ConfigurationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConfigurationType: valid values are %v", v, AllowedConfigurationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConfigurationType) IsValid() bool {
	for _, existing := range AllowedConfigurationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConfigurationType value
func (v ConfigurationType) Ptr() *ConfigurationType {
	return &v
}

type NullableConfigurationType struct {
	value *ConfigurationType
	isSet bool
}

func (v NullableConfigurationType) Get() *ConfigurationType {
	return v.value
}

func (v *NullableConfigurationType) Set(val *ConfigurationType) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationType) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationType(val *ConfigurationType) *NullableConfigurationType {
	return &NullableConfigurationType{value: val, isSet: true}
}

func (v NullableConfigurationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

