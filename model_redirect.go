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

// checks if the Redirect type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Redirect{}

// Redirect A Redirect.
type Redirect struct {
	// When creating a new Redirect, this is the ID for the Redirect. If not specified, an ID will be automatically assigned. When updating an existing Redirect, this field is ignored and the ID is determined by the path parameter.
	Id     *string         `json:"id,omitempty"`
	Source *HostPort       `json:"source,omitempty"`
	Target *TargetHostPort `json:"target,omitempty"`
	// (sortable) The Redirect HTTP status code used by the response.
	ResponseCode *int64      `json:"responseCode,omitempty"`
	AuditLevel   *AuditLevel `json:"auditLevel,omitempty"`
}

// NewRedirect instantiates a new Redirect object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedirect() *Redirect {
	this := Redirect{}
	return &this
}

// NewRedirectWithDefaults instantiates a new Redirect object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedirectWithDefaults() *Redirect {
	this := Redirect{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Redirect) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Redirect) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Redirect) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Redirect) SetId(v string) {
	o.Id = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *Redirect) GetSource() HostPort {
	if o == nil || IsNil(o.Source) {
		var ret HostPort
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Redirect) GetSourceOk() (*HostPort, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *Redirect) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given HostPort and assigns it to the Source field.
func (o *Redirect) SetSource(v HostPort) {
	o.Source = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *Redirect) GetTarget() TargetHostPort {
	if o == nil || IsNil(o.Target) {
		var ret TargetHostPort
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Redirect) GetTargetOk() (*TargetHostPort, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *Redirect) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given TargetHostPort and assigns it to the Target field.
func (o *Redirect) SetTarget(v TargetHostPort) {
	o.Target = &v
}

// GetResponseCode returns the ResponseCode field value if set, zero value otherwise.
func (o *Redirect) GetResponseCode() int64 {
	if o == nil || IsNil(o.ResponseCode) {
		var ret int64
		return ret
	}
	return *o.ResponseCode
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Redirect) GetResponseCodeOk() (*int64, bool) {
	if o == nil || IsNil(o.ResponseCode) {
		return nil, false
	}
	return o.ResponseCode, true
}

// HasResponseCode returns a boolean if a field has been set.
func (o *Redirect) HasResponseCode() bool {
	if o != nil && !IsNil(o.ResponseCode) {
		return true
	}

	return false
}

// SetResponseCode gets a reference to the given int64 and assigns it to the ResponseCode field.
func (o *Redirect) SetResponseCode(v int64) {
	o.ResponseCode = &v
}

// GetAuditLevel returns the AuditLevel field value if set, zero value otherwise.
func (o *Redirect) GetAuditLevel() AuditLevel {
	if o == nil || IsNil(o.AuditLevel) {
		var ret AuditLevel
		return ret
	}
	return *o.AuditLevel
}

// GetAuditLevelOk returns a tuple with the AuditLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Redirect) GetAuditLevelOk() (*AuditLevel, bool) {
	if o == nil || IsNil(o.AuditLevel) {
		return nil, false
	}
	return o.AuditLevel, true
}

// HasAuditLevel returns a boolean if a field has been set.
func (o *Redirect) HasAuditLevel() bool {
	if o != nil && !IsNil(o.AuditLevel) {
		return true
	}

	return false
}

// SetAuditLevel gets a reference to the given AuditLevel and assigns it to the AuditLevel field.
func (o *Redirect) SetAuditLevel(v AuditLevel) {
	o.AuditLevel = &v
}

func (o Redirect) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Redirect) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	if !IsNil(o.ResponseCode) {
		toSerialize["responseCode"] = o.ResponseCode
	}
	if !IsNil(o.AuditLevel) {
		toSerialize["auditLevel"] = o.AuditLevel
	}
	return toSerialize, nil
}

type NullableRedirect struct {
	value *Redirect
	isSet bool
}

func (v NullableRedirect) Get() *Redirect {
	return v.value
}

func (v *NullableRedirect) Set(val *Redirect) {
	v.value = val
	v.isSet = true
}

func (v NullableRedirect) IsSet() bool {
	return v.isSet
}

func (v *NullableRedirect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedirect(val *Redirect) *NullableRedirect {
	return &NullableRedirect{value: val, isSet: true}
}

func (v NullableRedirect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedirect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
