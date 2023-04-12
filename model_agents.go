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

// checks if the Agents type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Agents{}

// Agents A collection of agents.
type Agents struct {
	// The actual list of agents.
	Items []Agent `json:"items"`
}

// NewAgents instantiates a new Agents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgents(items []Agent) *Agents {
	this := Agents{}
	this.Items = items
	return &this
}

// NewAgentsWithDefaults instantiates a new Agents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentsWithDefaults() *Agents {
	this := Agents{}
	return &this
}

// GetItems returns the Items field value
func (o *Agents) GetItems() []Agent {
	if o == nil {
		var ret []Agent
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *Agents) GetItemsOk() ([]Agent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *Agents) SetItems(v []Agent) {
	o.Items = v
}

func (o Agents) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Agents) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableAgents struct {
	value *Agents
	isSet bool
}

func (v NullableAgents) Get() *Agents {
	return v.value
}

func (v *NullableAgents) Set(val *Agents) {
	v.value = val
	v.isSet = true
}

func (v NullableAgents) IsSet() bool {
	return v.isSet
}

func (v *NullableAgents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgents(val *Agents) *NullableAgents {
	return &NullableAgents{value: val, isSet: true}
}

func (v NullableAgents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
