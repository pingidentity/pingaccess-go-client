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

// checks if the AuthenticationChallengePolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticationChallengePolicy{}

// AuthenticationChallengePolicy An authentication challenge policy.
type AuthenticationChallengePolicy struct {
	// The UUID for the authentication challenge policy. If not specified during creation, an ID will be automatically be assigned. When updating an existing authentication challenge policy, this field is ignored and the ID is determined from the URL path parameter.
	Id *string `json:"id,omitempty"`
	// (sortable) The name of this authentication challenge policy. The number of characters in the name is limited to 64.
	Name string `json:"name"`
	// (sortable) A description of the authentication challenge policy. The number of characters in the description is limited to 1000.
	Description *string `json:"description,omitempty"`
	// (sortable) This field is read-only and indicates this authentication challenge policy cannot be modified.
	System *bool `json:"system,omitempty"`
	// An array of challenge response mappings, ordered by the precedence of each mapping. When no challengeResponseChain is needed for the policy, this field must be set to an empty array.
	ChallengeResponseChain   []ChallengeResponseMapping `json:"challengeResponseChain"`
	DefaultChallengeResponse ChallengeResponse          `json:"defaultChallengeResponse"`
}

// NewAuthenticationChallengePolicy instantiates a new AuthenticationChallengePolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationChallengePolicy(name string, challengeResponseChain []ChallengeResponseMapping, defaultChallengeResponse ChallengeResponse) *AuthenticationChallengePolicy {
	this := AuthenticationChallengePolicy{}
	this.Name = name
	this.ChallengeResponseChain = challengeResponseChain
	this.DefaultChallengeResponse = defaultChallengeResponse
	return &this
}

// NewAuthenticationChallengePolicyWithDefaults instantiates a new AuthenticationChallengePolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationChallengePolicyWithDefaults() *AuthenticationChallengePolicy {
	this := AuthenticationChallengePolicy{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuthenticationChallengePolicy) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationChallengePolicy) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthenticationChallengePolicy) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AuthenticationChallengePolicy) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *AuthenticationChallengePolicy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AuthenticationChallengePolicy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AuthenticationChallengePolicy) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AuthenticationChallengePolicy) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationChallengePolicy) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AuthenticationChallengePolicy) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AuthenticationChallengePolicy) SetDescription(v string) {
	o.Description = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *AuthenticationChallengePolicy) GetSystem() bool {
	if o == nil || IsNil(o.System) {
		var ret bool
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationChallengePolicy) GetSystemOk() (*bool, bool) {
	if o == nil || IsNil(o.System) {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *AuthenticationChallengePolicy) HasSystem() bool {
	if o != nil && !IsNil(o.System) {
		return true
	}

	return false
}

// SetSystem gets a reference to the given bool and assigns it to the System field.
func (o *AuthenticationChallengePolicy) SetSystem(v bool) {
	o.System = &v
}

// GetChallengeResponseChain returns the ChallengeResponseChain field value
func (o *AuthenticationChallengePolicy) GetChallengeResponseChain() []ChallengeResponseMapping {
	if o == nil {
		var ret []ChallengeResponseMapping
		return ret
	}

	return o.ChallengeResponseChain
}

// GetChallengeResponseChainOk returns a tuple with the ChallengeResponseChain field value
// and a boolean to check if the value has been set.
func (o *AuthenticationChallengePolicy) GetChallengeResponseChainOk() ([]ChallengeResponseMapping, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChallengeResponseChain, true
}

// SetChallengeResponseChain sets field value
func (o *AuthenticationChallengePolicy) SetChallengeResponseChain(v []ChallengeResponseMapping) {
	o.ChallengeResponseChain = v
}

// GetDefaultChallengeResponse returns the DefaultChallengeResponse field value
func (o *AuthenticationChallengePolicy) GetDefaultChallengeResponse() ChallengeResponse {
	if o == nil {
		var ret ChallengeResponse
		return ret
	}

	return o.DefaultChallengeResponse
}

// GetDefaultChallengeResponseOk returns a tuple with the DefaultChallengeResponse field value
// and a boolean to check if the value has been set.
func (o *AuthenticationChallengePolicy) GetDefaultChallengeResponseOk() (*ChallengeResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultChallengeResponse, true
}

// SetDefaultChallengeResponse sets field value
func (o *AuthenticationChallengePolicy) SetDefaultChallengeResponse(v ChallengeResponse) {
	o.DefaultChallengeResponse = v
}

func (o AuthenticationChallengePolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticationChallengePolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.System) {
		toSerialize["system"] = o.System
	}
	toSerialize["challengeResponseChain"] = o.ChallengeResponseChain
	toSerialize["defaultChallengeResponse"] = o.DefaultChallengeResponse
	return toSerialize, nil
}

type NullableAuthenticationChallengePolicy struct {
	value *AuthenticationChallengePolicy
	isSet bool
}

func (v NullableAuthenticationChallengePolicy) Get() *AuthenticationChallengePolicy {
	return v.value
}

func (v *NullableAuthenticationChallengePolicy) Set(val *AuthenticationChallengePolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationChallengePolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationChallengePolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationChallengePolicy(val *AuthenticationChallengePolicy) *NullableAuthenticationChallengePolicy {
	return &NullableAuthenticationChallengePolicy{value: val, isSet: true}
}

func (v NullableAuthenticationChallengePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationChallengePolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
