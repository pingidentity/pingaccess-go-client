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

// GeneralNameEnum the model 'GeneralNameEnum'
type GeneralNameEnum string

// List of GeneralNameEnum
const (
	GENERALNAMEENUM_OTHER_NAME GeneralNameEnum = "otherName"
	GENERALNAMEENUM_RFC822_NAME GeneralNameEnum = "rfc822Name"
	GENERALNAMEENUM_D_NS_NAME GeneralNameEnum = "dNSName"
	GENERALNAMEENUM_X400_ADDRESS GeneralNameEnum = "x400Address"
	GENERALNAMEENUM_DIRECTORY_NAME GeneralNameEnum = "directoryName"
	GENERALNAMEENUM_EDI_PARTY_NAME GeneralNameEnum = "ediPartyName"
	GENERALNAMEENUM_UNIFORM_RESOURCE_IDENTIFIER GeneralNameEnum = "uniformResourceIdentifier"
	GENERALNAMEENUM_I_P_ADDRESS GeneralNameEnum = "iPAddress"
	GENERALNAMEENUM_REGISTERED_ID GeneralNameEnum = "registeredID"
)

// All allowed values of GeneralNameEnum enum
var AllowedGeneralNameEnumEnumValues = []GeneralNameEnum{
	"otherName",
	"rfc822Name",
	"dNSName",
	"x400Address",
	"directoryName",
	"ediPartyName",
	"uniformResourceIdentifier",
	"iPAddress",
	"registeredID",
}

func (v *GeneralNameEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GeneralNameEnum(value)
	for _, existing := range AllowedGeneralNameEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GeneralNameEnum", value)
}

// NewGeneralNameEnumFromValue returns a pointer to a valid GeneralNameEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGeneralNameEnumFromValue(v string) (*GeneralNameEnum, error) {
	ev := GeneralNameEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GeneralNameEnum: valid values are %v", v, AllowedGeneralNameEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GeneralNameEnum) IsValid() bool {
	for _, existing := range AllowedGeneralNameEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GeneralNameEnum value
func (v GeneralNameEnum) Ptr() *GeneralNameEnum {
	return &v
}

type NullableGeneralNameEnum struct {
	value *GeneralNameEnum
	isSet bool
}

func (v NullableGeneralNameEnum) Get() *GeneralNameEnum {
	return v.value
}

func (v *NullableGeneralNameEnum) Set(val *GeneralNameEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableGeneralNameEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableGeneralNameEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeneralNameEnum(val *GeneralNameEnum) *NullableGeneralNameEnum {
	return &NullableGeneralNameEnum{value: val, isSet: true}
}

func (v NullableGeneralNameEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeneralNameEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
