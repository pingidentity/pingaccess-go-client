/*
Administrative API Documentation

The PingAccess Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingAccess as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API.

API version: 7.3.0.2-SNAPSHOT
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the Rule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Rule{}

// Rule A rule.
type Rule struct {
	// (sortable) The rule's class name.
	ClassName string `json:"className"`
	// When creating a new Rule, this is the ID for the Rule. If not specified, an ID will be automatically assigned. When updating an existing Rule, this field is ignored and the ID is determined by the path parameter.
	Id *int64 `json:"id,omitempty"`
	// (sortable) The rule's name.
	Name string `json:"name"`
	// The supported destinations for this rule. This field is read-only and meant to provide contextual metadata on where the rule can be applied.
	SupportedDestinations []string `json:"supportedDestinations,omitempty"`
	// The rule's configuration data. - This value is a PingAccess plugin configuration (JSON).
	Configuration map[string]interface{} `json:"configuration,omitempty"`
}

// NewRule instantiates a new Rule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRule(className string, name string) *Rule {
	this := Rule{}
	this.ClassName = className
	this.Name = name
	return &this
}

// NewRuleWithDefaults instantiates a new Rule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleWithDefaults() *Rule {
	this := Rule{}
	return &this
}

// GetClassName returns the ClassName field value
func (o *Rule) GetClassName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassName
}

// GetClassNameOk returns a tuple with the ClassName field value
// and a boolean to check if the value has been set.
func (o *Rule) GetClassNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassName, true
}

// SetClassName sets field value
func (o *Rule) SetClassName(v string) {
	o.ClassName = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Rule) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rule) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Rule) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Rule) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *Rule) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Rule) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Rule) SetName(v string) {
	o.Name = v
}

// GetSupportedDestinations returns the SupportedDestinations field value if set, zero value otherwise.
func (o *Rule) GetSupportedDestinations() []string {
	if o == nil || IsNil(o.SupportedDestinations) {
		var ret []string
		return ret
	}
	return o.SupportedDestinations
}

// GetSupportedDestinationsOk returns a tuple with the SupportedDestinations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rule) GetSupportedDestinationsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedDestinations) {
		return nil, false
	}
	return o.SupportedDestinations, true
}

// HasSupportedDestinations returns a boolean if a field has been set.
func (o *Rule) HasSupportedDestinations() bool {
	if o != nil && !IsNil(o.SupportedDestinations) {
		return true
	}

	return false
}

// SetSupportedDestinations gets a reference to the given []string and assigns it to the SupportedDestinations field.
func (o *Rule) SetSupportedDestinations(v []string) {
	o.SupportedDestinations = v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *Rule) GetConfiguration() map[string]interface{} {
	if o == nil || IsNil(o.Configuration) {
		var ret map[string]interface{}
		return ret
	}
	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rule) GetConfigurationOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Configuration) {
		return map[string]interface{}{}, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *Rule) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given map[string]interface{} and assigns it to the Configuration field.
func (o *Rule) SetConfiguration(v map[string]interface{}) {
	o.Configuration = v
}

func (o Rule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Rule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["className"] = o.ClassName
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.SupportedDestinations) {
		toSerialize["supportedDestinations"] = o.SupportedDestinations
	}
	if !IsNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	return toSerialize, nil
}

type NullableRule struct {
	value *Rule
	isSet bool
}

func (v NullableRule) Get() *Rule {
	return v.value
}

func (v *NullableRule) Set(val *Rule) {
	v.value = val
	v.isSet = true
}

func (v NullableRule) IsSet() bool {
	return v.isSet
}

func (v *NullableRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRule(val *Rule) *NullableRule {
	return &NullableRule{value: val, isSet: true}
}

func (v NullableRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
