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

// checks if the TokenProviderSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenProviderSetting{}

// TokenProviderSetting Settings for a token provider.
type TokenProviderSetting struct {
	// This field is true if third-party Token Provider is enabled, and false if PingFederate is enabled. (DEPRECATED - to be removed in a future release; please use 'type' instead)
	UseThirdParty *bool              `json:"useThirdParty,omitempty"`
	Type          *TokenProviderType `json:"type,omitempty"`
}

// NewTokenProviderSetting instantiates a new TokenProviderSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenProviderSetting() *TokenProviderSetting {
	this := TokenProviderSetting{}
	return &this
}

// NewTokenProviderSettingWithDefaults instantiates a new TokenProviderSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenProviderSettingWithDefaults() *TokenProviderSetting {
	this := TokenProviderSetting{}
	return &this
}

// GetUseThirdParty returns the UseThirdParty field value if set, zero value otherwise.
func (o *TokenProviderSetting) GetUseThirdParty() bool {
	if o == nil || IsNil(o.UseThirdParty) {
		var ret bool
		return ret
	}
	return *o.UseThirdParty
}

// GetUseThirdPartyOk returns a tuple with the UseThirdParty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenProviderSetting) GetUseThirdPartyOk() (*bool, bool) {
	if o == nil || IsNil(o.UseThirdParty) {
		return nil, false
	}
	return o.UseThirdParty, true
}

// HasUseThirdParty returns a boolean if a field has been set.
func (o *TokenProviderSetting) HasUseThirdParty() bool {
	if o != nil && !IsNil(o.UseThirdParty) {
		return true
	}

	return false
}

// SetUseThirdParty gets a reference to the given bool and assigns it to the UseThirdParty field.
func (o *TokenProviderSetting) SetUseThirdParty(v bool) {
	o.UseThirdParty = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TokenProviderSetting) GetType() TokenProviderType {
	if o == nil || IsNil(o.Type) {
		var ret TokenProviderType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenProviderSetting) GetTypeOk() (*TokenProviderType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TokenProviderSetting) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given TokenProviderType and assigns it to the Type field.
func (o *TokenProviderSetting) SetType(v TokenProviderType) {
	o.Type = &v
}

func (o TokenProviderSetting) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenProviderSetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UseThirdParty) {
		toSerialize["useThirdParty"] = o.UseThirdParty
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableTokenProviderSetting struct {
	value *TokenProviderSetting
	isSet bool
}

func (v NullableTokenProviderSetting) Get() *TokenProviderSetting {
	return v.value
}

func (v *NullableTokenProviderSetting) Set(val *TokenProviderSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenProviderSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenProviderSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenProviderSetting(val *TokenProviderSetting) *NullableTokenProviderSetting {
	return &NullableTokenProviderSetting{value: val, isSet: true}
}

func (v NullableTokenProviderSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenProviderSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
