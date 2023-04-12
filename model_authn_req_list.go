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

// checks if the AuthnReqList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthnReqList{}

// AuthnReqList An authentication requirements list.
type AuthnReqList struct {
	// When creating a new AuthnReqList, this is the ID for the AuthnReqList. If not specified, an ID will be automatically assigned. When updating an existing AuthnReqList, this field is ignored and the ID is determined by the path parameter.
	Id *int32 `json:"id,omitempty"`
	// (sortable) The name of the authentication requirements list.
	Name string `json:"name"`
	// The ordered list of authentication requirements, or identifiers, which define how PingFederate will authenticate a user during the OIDC login flow.
	AuthnReqs []string `json:"authnReqs"`
}

// NewAuthnReqList instantiates a new AuthnReqList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthnReqList(name string, authnReqs []string) *AuthnReqList {
	this := AuthnReqList{}
	this.Name = name
	this.AuthnReqs = authnReqs
	return &this
}

// NewAuthnReqListWithDefaults instantiates a new AuthnReqList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthnReqListWithDefaults() *AuthnReqList {
	this := AuthnReqList{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuthnReqList) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthnReqList) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthnReqList) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AuthnReqList) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *AuthnReqList) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AuthnReqList) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AuthnReqList) SetName(v string) {
	o.Name = v
}

// GetAuthnReqs returns the AuthnReqs field value
func (o *AuthnReqList) GetAuthnReqs() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AuthnReqs
}

// GetAuthnReqsOk returns a tuple with the AuthnReqs field value
// and a boolean to check if the value has been set.
func (o *AuthnReqList) GetAuthnReqsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthnReqs, true
}

// SetAuthnReqs sets field value
func (o *AuthnReqList) SetAuthnReqs(v []string) {
	o.AuthnReqs = v
}

func (o AuthnReqList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthnReqList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	toSerialize["authnReqs"] = o.AuthnReqs
	return toSerialize, nil
}

type NullableAuthnReqList struct {
	value *AuthnReqList
	isSet bool
}

func (v NullableAuthnReqList) Get() *AuthnReqList {
	return v.value
}

func (v *NullableAuthnReqList) Set(val *AuthnReqList) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthnReqList) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthnReqList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthnReqList(val *AuthnReqList) *NullableAuthnReqList {
	return &NullableAuthnReqList{value: val, isSet: true}
}

func (v NullableAuthnReqList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthnReqList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
