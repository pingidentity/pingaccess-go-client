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

// SiteAuthenticator A site authenticator.
type SiteAuthenticator struct {
	// (sortable) The site authenticator's class name.
	ClassName string `json:"className"`
	// When creating a new SiteAuthenticator, this is the ID for the SiteAuthenticator. If not specified, an ID will be automatically assigned. When updating an existing SiteAuthenticator, this field is ignored and the ID is determined by the path parameter.
	Id *int32 `json:"id,omitempty"`
	// (sortable) The site authenticator's name.
	Name string `json:"name"`
	// The site authenticator's configuration data.
	Configuration *string `json:"configuration,omitempty"`
}

// NewSiteAuthenticator instantiates a new SiteAuthenticator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteAuthenticator(className string, name string) *SiteAuthenticator {
	this := SiteAuthenticator{}
	this.ClassName = className
	this.Name = name
	return &this
}

// NewSiteAuthenticatorWithDefaults instantiates a new SiteAuthenticator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteAuthenticatorWithDefaults() *SiteAuthenticator {
	this := SiteAuthenticator{}
	return &this
}

// GetClassName returns the ClassName field value
func (o *SiteAuthenticator) GetClassName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassName
}

// GetClassNameOk returns a tuple with the ClassName field value
// and a boolean to check if the value has been set.
func (o *SiteAuthenticator) GetClassNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ClassName, true
}

// SetClassName sets field value
func (o *SiteAuthenticator) SetClassName(v string) {
	o.ClassName = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SiteAuthenticator) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAuthenticator) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SiteAuthenticator) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *SiteAuthenticator) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *SiteAuthenticator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SiteAuthenticator) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SiteAuthenticator) SetName(v string) {
	o.Name = v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *SiteAuthenticator) GetConfiguration() string {
	if o == nil || isNil(o.Configuration) {
		var ret string
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAuthenticator) GetConfigurationOk() (*string, bool) {
	if o == nil || isNil(o.Configuration) {
    return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *SiteAuthenticator) HasConfiguration() bool {
	if o != nil && !isNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given string and assigns it to the Configuration field.
func (o *SiteAuthenticator) SetConfiguration(v string) {
	o.Configuration = &v
}

func (o SiteAuthenticator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["className"] = o.ClassName
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	return json.Marshal(toSerialize)
}

type NullableSiteAuthenticator struct {
	value *SiteAuthenticator
	isSet bool
}

func (v NullableSiteAuthenticator) Get() *SiteAuthenticator {
	return v.value
}

func (v *NullableSiteAuthenticator) Set(val *SiteAuthenticator) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteAuthenticator) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteAuthenticator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteAuthenticator(val *SiteAuthenticator) *NullableSiteAuthenticator {
	return &NullableSiteAuthenticator{value: val, isSet: true}
}

func (v NullableSiteAuthenticator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteAuthenticator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


