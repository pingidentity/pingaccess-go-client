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

// checks if the PingOneConnections type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PingOneConnections{}

// PingOneConnections A collection of PingOne Connections.
type PingOneConnections struct {
	// An array of PingOne Connections.
	Items []PingOneConnection `json:"items"`
}

// NewPingOneConnections instantiates a new PingOneConnections object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPingOneConnections(items []PingOneConnection) *PingOneConnections {
	this := PingOneConnections{}
	this.Items = items
	return &this
}

// NewPingOneConnectionsWithDefaults instantiates a new PingOneConnections object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPingOneConnectionsWithDefaults() *PingOneConnections {
	this := PingOneConnections{}
	return &this
}

// GetItems returns the Items field value
func (o *PingOneConnections) GetItems() []PingOneConnection {
	if o == nil {
		var ret []PingOneConnection
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *PingOneConnections) GetItemsOk() ([]PingOneConnection, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *PingOneConnections) SetItems(v []PingOneConnection) {
	o.Items = v
}

func (o PingOneConnections) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PingOneConnections) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullablePingOneConnections struct {
	value *PingOneConnections
	isSet bool
}

func (v NullablePingOneConnections) Get() *PingOneConnections {
	return v.value
}

func (v *NullablePingOneConnections) Set(val *PingOneConnections) {
	v.value = val
	v.isSet = true
}

func (v NullablePingOneConnections) IsSet() bool {
	return v.isSet
}

func (v *NullablePingOneConnections) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePingOneConnections(val *PingOneConnections) *NullablePingOneConnections {
	return &NullablePingOneConnections{value: val, isSet: true}
}

func (v NullablePingOneConnections) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePingOneConnections) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
