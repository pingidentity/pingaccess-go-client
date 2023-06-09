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

// checks if the VirtualHosts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualHosts{}

// VirtualHosts A collection of virtual hosts.
type VirtualHosts struct {
	// The actual list of virtual hosts.
	Items []VirtualHost `json:"items"`
}

// NewVirtualHosts instantiates a new VirtualHosts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualHosts(items []VirtualHost) *VirtualHosts {
	this := VirtualHosts{}
	this.Items = items
	return &this
}

// NewVirtualHostsWithDefaults instantiates a new VirtualHosts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualHostsWithDefaults() *VirtualHosts {
	this := VirtualHosts{}
	return &this
}

// GetItems returns the Items field value
func (o *VirtualHosts) GetItems() []VirtualHost {
	if o == nil {
		var ret []VirtualHost
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *VirtualHosts) GetItemsOk() ([]VirtualHost, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *VirtualHosts) SetItems(v []VirtualHost) {
	o.Items = v
}

func (o VirtualHosts) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualHosts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableVirtualHosts struct {
	value *VirtualHosts
	isSet bool
}

func (v NullableVirtualHosts) Get() *VirtualHosts {
	return v.value
}

func (v *NullableVirtualHosts) Set(val *VirtualHosts) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualHosts) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualHosts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualHosts(val *VirtualHosts) *NullableVirtualHosts {
	return &NullableVirtualHosts{value: val, isSet: true}
}

func (v NullableVirtualHosts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualHosts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
