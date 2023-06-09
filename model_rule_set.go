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

// checks if the RuleSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleSet{}

// RuleSet A rule set.
type RuleSet struct {
	// When creating a new RuleSet, this is the ID for the RuleSet. If not specified, an ID will be automatically assigned. When updating an existing RuleSet, this field is ignored and the ID is determined by the path parameter.
	Id *int64 `json:"id,omitempty"`
	// (sortable) The rule set's name.
	Name            string              `json:"name"`
	SuccessCriteria *SuccessCriteria    `json:"successCriteria,omitempty"`
	ElementType     *RuleSetElementType `json:"elementType,omitempty"`
	// The list of policy ids assigned to the rule set.
	Policy []int64 `json:"policy"`
}

// NewRuleSet instantiates a new RuleSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleSet(name string, policy []int64) *RuleSet {
	this := RuleSet{}
	this.Name = name
	this.Policy = policy
	return &this
}

// NewRuleSetWithDefaults instantiates a new RuleSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleSetWithDefaults() *RuleSet {
	this := RuleSet{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RuleSet) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleSet) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RuleSet) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *RuleSet) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *RuleSet) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RuleSet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RuleSet) SetName(v string) {
	o.Name = v
}

// GetSuccessCriteria returns the SuccessCriteria field value if set, zero value otherwise.
func (o *RuleSet) GetSuccessCriteria() SuccessCriteria {
	if o == nil || IsNil(o.SuccessCriteria) {
		var ret SuccessCriteria
		return ret
	}
	return *o.SuccessCriteria
}

// GetSuccessCriteriaOk returns a tuple with the SuccessCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleSet) GetSuccessCriteriaOk() (*SuccessCriteria, bool) {
	if o == nil || IsNil(o.SuccessCriteria) {
		return nil, false
	}
	return o.SuccessCriteria, true
}

// HasSuccessCriteria returns a boolean if a field has been set.
func (o *RuleSet) HasSuccessCriteria() bool {
	if o != nil && !IsNil(o.SuccessCriteria) {
		return true
	}

	return false
}

// SetSuccessCriteria gets a reference to the given SuccessCriteria and assigns it to the SuccessCriteria field.
func (o *RuleSet) SetSuccessCriteria(v SuccessCriteria) {
	o.SuccessCriteria = &v
}

// GetElementType returns the ElementType field value if set, zero value otherwise.
func (o *RuleSet) GetElementType() RuleSetElementType {
	if o == nil || IsNil(o.ElementType) {
		var ret RuleSetElementType
		return ret
	}
	return *o.ElementType
}

// GetElementTypeOk returns a tuple with the ElementType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleSet) GetElementTypeOk() (*RuleSetElementType, bool) {
	if o == nil || IsNil(o.ElementType) {
		return nil, false
	}
	return o.ElementType, true
}

// HasElementType returns a boolean if a field has been set.
func (o *RuleSet) HasElementType() bool {
	if o != nil && !IsNil(o.ElementType) {
		return true
	}

	return false
}

// SetElementType gets a reference to the given RuleSetElementType and assigns it to the ElementType field.
func (o *RuleSet) SetElementType(v RuleSetElementType) {
	o.ElementType = &v
}

// GetPolicy returns the Policy field value
func (o *RuleSet) GetPolicy() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value
// and a boolean to check if the value has been set.
func (o *RuleSet) GetPolicyOk() ([]int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Policy, true
}

// SetPolicy sets field value
func (o *RuleSet) SetPolicy(v []int64) {
	o.Policy = v
}

func (o RuleSet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuleSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.SuccessCriteria) {
		toSerialize["successCriteria"] = o.SuccessCriteria
	}
	if !IsNil(o.ElementType) {
		toSerialize["elementType"] = o.ElementType
	}
	toSerialize["policy"] = o.Policy
	return toSerialize, nil
}

type NullableRuleSet struct {
	value *RuleSet
	isSet bool
}

func (v NullableRuleSet) Get() *RuleSet {
	return v.value
}

func (v *NullableRuleSet) Set(val *RuleSet) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleSet) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleSet(val *RuleSet) *NullableRuleSet {
	return &NullableRuleSet{value: val, isSet: true}
}

func (v NullableRuleSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
