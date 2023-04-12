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

// checks if the LoadBalancingStrategies type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancingStrategies{}

// LoadBalancingStrategies A collection of load balancing strategies.
type LoadBalancingStrategies struct {
	// An array of load balancing strategies.
	Items []LoadBalancingStrategy `json:"items"`
}

// NewLoadBalancingStrategies instantiates a new LoadBalancingStrategies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancingStrategies(items []LoadBalancingStrategy) *LoadBalancingStrategies {
	this := LoadBalancingStrategies{}
	this.Items = items
	return &this
}

// NewLoadBalancingStrategiesWithDefaults instantiates a new LoadBalancingStrategies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancingStrategiesWithDefaults() *LoadBalancingStrategies {
	this := LoadBalancingStrategies{}
	return &this
}

// GetItems returns the Items field value
func (o *LoadBalancingStrategies) GetItems() []LoadBalancingStrategy {
	if o == nil {
		var ret []LoadBalancingStrategy
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *LoadBalancingStrategies) GetItemsOk() ([]LoadBalancingStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *LoadBalancingStrategies) SetItems(v []LoadBalancingStrategy) {
	o.Items = v
}

func (o LoadBalancingStrategies) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadBalancingStrategies) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableLoadBalancingStrategies struct {
	value *LoadBalancingStrategies
	isSet bool
}

func (v NullableLoadBalancingStrategies) Get() *LoadBalancingStrategies {
	return v.value
}

func (v *NullableLoadBalancingStrategies) Set(val *LoadBalancingStrategies) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancingStrategies) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancingStrategies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancingStrategies(val *LoadBalancingStrategies) *NullableLoadBalancingStrategies {
	return &NullableLoadBalancingStrategies{value: val, isSet: true}
}

func (v NullableLoadBalancingStrategies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancingStrategies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
