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

// ThirdPartyServices A collection of third-party service items.
type ThirdPartyServices struct {
	// The actual list of third-party services.
	Items []ThirdPartyService `json:"items"`
}

// NewThirdPartyServices instantiates a new ThirdPartyServices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartyServices(items []ThirdPartyService) *ThirdPartyServices {
	this := ThirdPartyServices{}
	this.Items = items
	return &this
}

// NewThirdPartyServicesWithDefaults instantiates a new ThirdPartyServices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartyServicesWithDefaults() *ThirdPartyServices {
	this := ThirdPartyServices{}
	return &this
}

// GetItems returns the Items field value
func (o *ThirdPartyServices) GetItems() []ThirdPartyService {
	if o == nil {
		var ret []ThirdPartyService
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyServices) GetItemsOk() ([]ThirdPartyService, bool) {
	if o == nil {
    return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ThirdPartyServices) SetItems(v []ThirdPartyService) {
	o.Items = v
}

func (o ThirdPartyServices) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableThirdPartyServices struct {
	value *ThirdPartyServices
	isSet bool
}

func (v NullableThirdPartyServices) Get() *ThirdPartyServices {
	return v.value
}

func (v *NullableThirdPartyServices) Set(val *ThirdPartyServices) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartyServices) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartyServices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartyServices(val *ThirdPartyServices) *NullableThirdPartyServices {
	return &NullableThirdPartyServices{value: val, isSet: true}
}

func (v NullableThirdPartyServices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartyServices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

