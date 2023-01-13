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

// SiteAuthenticators A collection of site authenticators.
type SiteAuthenticators struct {
	// The actual list of site authenticators.
	Items []SiteAuthenticator `json:"items"`
}

// NewSiteAuthenticators instantiates a new SiteAuthenticators object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteAuthenticators(items []SiteAuthenticator) *SiteAuthenticators {
	this := SiteAuthenticators{}
	this.Items = items
	return &this
}

// NewSiteAuthenticatorsWithDefaults instantiates a new SiteAuthenticators object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteAuthenticatorsWithDefaults() *SiteAuthenticators {
	this := SiteAuthenticators{}
	return &this
}

// GetItems returns the Items field value
func (o *SiteAuthenticators) GetItems() []SiteAuthenticator {
	if o == nil {
		var ret []SiteAuthenticator
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *SiteAuthenticators) GetItemsOk() ([]SiteAuthenticator, bool) {
	if o == nil {
    return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *SiteAuthenticators) SetItems(v []SiteAuthenticator) {
	o.Items = v
}

func (o SiteAuthenticators) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableSiteAuthenticators struct {
	value *SiteAuthenticators
	isSet bool
}

func (v NullableSiteAuthenticators) Get() *SiteAuthenticators {
	return v.value
}

func (v *NullableSiteAuthenticators) Set(val *SiteAuthenticators) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteAuthenticators) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteAuthenticators) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteAuthenticators(val *SiteAuthenticators) *NullableSiteAuthenticators {
	return &NullableSiteAuthenticators{value: val, isSet: true}
}

func (v NullableSiteAuthenticators) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteAuthenticators) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

