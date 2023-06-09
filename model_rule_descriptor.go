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

// checks if the RuleDescriptor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleDescriptor{}

// RuleDescriptor A rule descriptor.
type RuleDescriptor struct {
	// The list of configuration fields of the rule
	ConfigurationFields []ConfigurationField `json:"configurationFields"`
	// The type of the rule.
	Type string `json:"type"`
	// The label of the rule.
	Label string `json:"label"`
	// The class name of the rule.
	ClassName string                  `json:"className"`
	Category  RuleInterceptorCategory `json:"category"`
	// The modes that the rules can be used in. Site mode implies Sideband and Site destination support.
	Modes []RuleInterceptorSupportedDestination `json:"modes"`
	// Indicates that agent caching is disabled or not.
	AgentCachingDisabled bool `json:"agentCachingDisabled"`
}

// NewRuleDescriptor instantiates a new RuleDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleDescriptor(configurationFields []ConfigurationField, type_ string, label string, className string, category RuleInterceptorCategory, modes []RuleInterceptorSupportedDestination, agentCachingDisabled bool) *RuleDescriptor {
	this := RuleDescriptor{}
	this.ConfigurationFields = configurationFields
	this.Type = type_
	this.Label = label
	this.ClassName = className
	this.Category = category
	this.Modes = modes
	this.AgentCachingDisabled = agentCachingDisabled
	return &this
}

// NewRuleDescriptorWithDefaults instantiates a new RuleDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleDescriptorWithDefaults() *RuleDescriptor {
	this := RuleDescriptor{}
	return &this
}

// GetConfigurationFields returns the ConfigurationFields field value
func (o *RuleDescriptor) GetConfigurationFields() []ConfigurationField {
	if o == nil {
		var ret []ConfigurationField
		return ret
	}

	return o.ConfigurationFields
}

// GetConfigurationFieldsOk returns a tuple with the ConfigurationFields field value
// and a boolean to check if the value has been set.
func (o *RuleDescriptor) GetConfigurationFieldsOk() ([]ConfigurationField, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigurationFields, true
}

// SetConfigurationFields sets field value
func (o *RuleDescriptor) SetConfigurationFields(v []ConfigurationField) {
	o.ConfigurationFields = v
}

// GetType returns the Type field value
func (o *RuleDescriptor) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RuleDescriptor) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RuleDescriptor) SetType(v string) {
	o.Type = v
}

// GetLabel returns the Label field value
func (o *RuleDescriptor) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *RuleDescriptor) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *RuleDescriptor) SetLabel(v string) {
	o.Label = v
}

// GetClassName returns the ClassName field value
func (o *RuleDescriptor) GetClassName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassName
}

// GetClassNameOk returns a tuple with the ClassName field value
// and a boolean to check if the value has been set.
func (o *RuleDescriptor) GetClassNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassName, true
}

// SetClassName sets field value
func (o *RuleDescriptor) SetClassName(v string) {
	o.ClassName = v
}

// GetCategory returns the Category field value
func (o *RuleDescriptor) GetCategory() RuleInterceptorCategory {
	if o == nil {
		var ret RuleInterceptorCategory
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *RuleDescriptor) GetCategoryOk() (*RuleInterceptorCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *RuleDescriptor) SetCategory(v RuleInterceptorCategory) {
	o.Category = v
}

// GetModes returns the Modes field value
func (o *RuleDescriptor) GetModes() []RuleInterceptorSupportedDestination {
	if o == nil {
		var ret []RuleInterceptorSupportedDestination
		return ret
	}

	return o.Modes
}

// GetModesOk returns a tuple with the Modes field value
// and a boolean to check if the value has been set.
func (o *RuleDescriptor) GetModesOk() ([]RuleInterceptorSupportedDestination, bool) {
	if o == nil {
		return nil, false
	}
	return o.Modes, true
}

// SetModes sets field value
func (o *RuleDescriptor) SetModes(v []RuleInterceptorSupportedDestination) {
	o.Modes = v
}

// GetAgentCachingDisabled returns the AgentCachingDisabled field value
func (o *RuleDescriptor) GetAgentCachingDisabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AgentCachingDisabled
}

// GetAgentCachingDisabledOk returns a tuple with the AgentCachingDisabled field value
// and a boolean to check if the value has been set.
func (o *RuleDescriptor) GetAgentCachingDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentCachingDisabled, true
}

// SetAgentCachingDisabled sets field value
func (o *RuleDescriptor) SetAgentCachingDisabled(v bool) {
	o.AgentCachingDisabled = v
}

func (o RuleDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuleDescriptor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["configurationFields"] = o.ConfigurationFields
	toSerialize["type"] = o.Type
	toSerialize["label"] = o.Label
	toSerialize["className"] = o.ClassName
	toSerialize["category"] = o.Category
	toSerialize["modes"] = o.Modes
	toSerialize["agentCachingDisabled"] = o.AgentCachingDisabled
	return toSerialize, nil
}

type NullableRuleDescriptor struct {
	value *RuleDescriptor
	isSet bool
}

func (v NullableRuleDescriptor) Get() *RuleDescriptor {
	return v.value
}

func (v *NullableRuleDescriptor) Set(val *RuleDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleDescriptor(val *RuleDescriptor) *NullableRuleDescriptor {
	return &NullableRuleDescriptor{value: val, isSet: true}
}

func (v NullableRuleDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
