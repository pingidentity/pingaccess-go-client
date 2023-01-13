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

// Destination the model 'Destination'
type Destination string

// List of Destination
const (
	DESTINATION_SITE Destination = "Site"
	DESTINATION_AGENT Destination = "Agent"
	DESTINATION_SIDEBAND Destination = "Sideband"
)

// All allowed values of Destination enum
var AllowedDestinationEnumValues = []Destination{
	"Site",
	"Agent",
	"Sideband",
}

func (v *Destination) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Destination(value)
	for _, existing := range AllowedDestinationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Destination", value)
}

// NewDestinationFromValue returns a pointer to a valid Destination
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDestinationFromValue(v string) (*Destination, error) {
	ev := Destination(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Destination: valid values are %v", v, AllowedDestinationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Destination) IsValid() bool {
	for _, existing := range AllowedDestinationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Destination value
func (v Destination) Ptr() *Destination {
	return &v
}

type NullableDestination struct {
	value *Destination
	isSet bool
}

func (v NullableDestination) Get() *Destination {
	return v.value
}

func (v *NullableDestination) Set(val *Destination) {
	v.value = val
	v.isSet = true
}

func (v NullableDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestination(val *Destination) *NullableDestination {
	return &NullableDestination{value: val, isSet: true}
}

func (v NullableDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

