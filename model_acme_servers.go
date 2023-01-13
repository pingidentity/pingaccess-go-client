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

// AcmeServers A collection of ACME servers.
type AcmeServers struct {
	// An array of ACME servers.
	Items []AcmeServer `json:"items"`
}

// NewAcmeServers instantiates a new AcmeServers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcmeServers(items []AcmeServer) *AcmeServers {
	this := AcmeServers{}
	this.Items = items
	return &this
}

// NewAcmeServersWithDefaults instantiates a new AcmeServers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcmeServersWithDefaults() *AcmeServers {
	this := AcmeServers{}
	return &this
}

// GetItems returns the Items field value
func (o *AcmeServers) GetItems() []AcmeServer {
	if o == nil {
		var ret []AcmeServer
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *AcmeServers) GetItemsOk() ([]AcmeServer, bool) {
	if o == nil {
    return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *AcmeServers) SetItems(v []AcmeServer) {
	o.Items = v
}

func (o AcmeServers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableAcmeServers struct {
	value *AcmeServers
	isSet bool
}

func (v NullableAcmeServers) Get() *AcmeServers {
	return v.value
}

func (v *NullableAcmeServers) Set(val *AcmeServers) {
	v.value = val
	v.isSet = true
}

func (v NullableAcmeServers) IsSet() bool {
	return v.isSet
}

func (v *NullableAcmeServers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcmeServers(val *AcmeServers) *NullableAcmeServers {
	return &NullableAcmeServers{value: val, isSet: true}
}

func (v NullableAcmeServers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcmeServers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

