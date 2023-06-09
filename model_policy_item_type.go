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

// PolicyItemType the model 'PolicyItemType'
type PolicyItemType string

// List of PolicyItemType
const (
	POLICYITEMTYPE_RULE     PolicyItemType = "Rule"
	POLICYITEMTYPE_RULE_SET PolicyItemType = "RuleSet"
)

// All allowed values of PolicyItemType enum
var AllowedPolicyItemTypeEnumValues = []PolicyItemType{
	"Rule",
	"RuleSet",
}

func (v *PolicyItemType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PolicyItemType(value)
	for _, existing := range AllowedPolicyItemTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PolicyItemType", value)
}

// NewPolicyItemTypeFromValue returns a pointer to a valid PolicyItemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPolicyItemTypeFromValue(v string) (*PolicyItemType, error) {
	ev := PolicyItemType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PolicyItemType: valid values are %v", v, AllowedPolicyItemTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PolicyItemType) IsValid() bool {
	for _, existing := range AllowedPolicyItemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PolicyItemType value
func (v PolicyItemType) Ptr() *PolicyItemType {
	return &v
}

type NullablePolicyItemType struct {
	value *PolicyItemType
	isSet bool
}

func (v NullablePolicyItemType) Get() *PolicyItemType {
	return v.value
}

func (v *NullablePolicyItemType) Set(val *PolicyItemType) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyItemType) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyItemType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyItemType(val *PolicyItemType) *NullablePolicyItemType {
	return &NullablePolicyItemType{value: val, isSet: true}
}

func (v NullablePolicyItemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyItemType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
