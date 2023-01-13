/*
PingAccess Administrative API

The PingAccess Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingAccess.

API version: 7.1.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Attribute A name-value pair of user attributes.
type Attribute struct {
	// The name of the user attribute.
	AttributeName string `json:"attributeName"`
	// The value of the user attribute.
	AttributeValue string `json:"attributeValue"`
}

// NewAttribute instantiates a new Attribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttribute(attributeName string, attributeValue string) *Attribute {
	this := Attribute{}
	this.AttributeName = attributeName
	this.AttributeValue = attributeValue
	return &this
}

// NewAttributeWithDefaults instantiates a new Attribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributeWithDefaults() *Attribute {
	this := Attribute{}
	return &this
}

// GetAttributeName returns the AttributeName field value
func (o *Attribute) GetAttributeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeName
}

// GetAttributeNameOk returns a tuple with the AttributeName field value
// and a boolean to check if the value has been set.
func (o *Attribute) GetAttributeNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AttributeName, true
}

// SetAttributeName sets field value
func (o *Attribute) SetAttributeName(v string) {
	o.AttributeName = v
}

// GetAttributeValue returns the AttributeValue field value
func (o *Attribute) GetAttributeValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeValue
}

// GetAttributeValueOk returns a tuple with the AttributeValue field value
// and a boolean to check if the value has been set.
func (o *Attribute) GetAttributeValueOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AttributeValue, true
}

// SetAttributeValue sets field value
func (o *Attribute) SetAttributeValue(v string) {
	o.AttributeValue = v
}

func (o Attribute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["attributeName"] = o.AttributeName
	}
	if true {
		toSerialize["attributeValue"] = o.AttributeValue
	}
	return json.Marshal(toSerialize)
}

type NullableAttribute struct {
	value *Attribute
	isSet bool
}

func (v NullableAttribute) Get() *Attribute {
	return v.value
}

func (v *NullableAttribute) Set(val *Attribute) {
	v.value = val
	v.isSet = true
}

func (v NullableAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttribute(val *Attribute) *NullableAttribute {
	return &NullableAttribute{value: val, isSet: true}
}

func (v NullableAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


