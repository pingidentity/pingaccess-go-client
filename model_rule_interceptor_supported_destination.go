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

// RuleInterceptorSupportedDestination the model 'RuleInterceptorSupportedDestination'
type RuleInterceptorSupportedDestination string

// List of RuleInterceptorSupportedDestination
const (
	RULEINTERCEPTORSUPPORTEDDESTINATION_SITE RuleInterceptorSupportedDestination = "Site"
	RULEINTERCEPTORSUPPORTEDDESTINATION_AGENT RuleInterceptorSupportedDestination = "Agent"
)

// All allowed values of RuleInterceptorSupportedDestination enum
var AllowedRuleInterceptorSupportedDestinationEnumValues = []RuleInterceptorSupportedDestination{
	"Site",
	"Agent",
}

func (v *RuleInterceptorSupportedDestination) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RuleInterceptorSupportedDestination(value)
	for _, existing := range AllowedRuleInterceptorSupportedDestinationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RuleInterceptorSupportedDestination", value)
}

// NewRuleInterceptorSupportedDestinationFromValue returns a pointer to a valid RuleInterceptorSupportedDestination
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRuleInterceptorSupportedDestinationFromValue(v string) (*RuleInterceptorSupportedDestination, error) {
	ev := RuleInterceptorSupportedDestination(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RuleInterceptorSupportedDestination: valid values are %v", v, AllowedRuleInterceptorSupportedDestinationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RuleInterceptorSupportedDestination) IsValid() bool {
	for _, existing := range AllowedRuleInterceptorSupportedDestinationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RuleInterceptorSupportedDestination value
func (v RuleInterceptorSupportedDestination) Ptr() *RuleInterceptorSupportedDestination {
	return &v
}

type NullableRuleInterceptorSupportedDestination struct {
	value *RuleInterceptorSupportedDestination
	isSet bool
}

func (v NullableRuleInterceptorSupportedDestination) Get() *RuleInterceptorSupportedDestination {
	return v.value
}

func (v *NullableRuleInterceptorSupportedDestination) Set(val *RuleInterceptorSupportedDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleInterceptorSupportedDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleInterceptorSupportedDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleInterceptorSupportedDestination(val *RuleInterceptorSupportedDestination) *NullableRuleInterceptorSupportedDestination {
	return &NullableRuleInterceptorSupportedDestination{value: val, isSet: true}
}

func (v NullableRuleInterceptorSupportedDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleInterceptorSupportedDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

