/*
PingAccess Administrative API

The PingAccess Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingAccess.

API version: 7.1.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the AccessTokenValidatorConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessTokenValidatorConfiguration{}

// AccessTokenValidatorConfiguration The access token validator's configuration data.
type AccessTokenValidatorConfiguration struct {
	Description          *string `json:"description,omitempty"`
	Path                 string  `json:"path"`
	SubjectAttributeName *string `json:"subjectAttributeName,omitempty"`
	Issuer               *string `json:"issuer,omitempty"`
	Audience             *string `json:"audience,omitempty"`
}

// NewAccessTokenValidatorConfiguration instantiates a new AccessTokenValidatorConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessTokenValidatorConfiguration(path string) *AccessTokenValidatorConfiguration {
	this := AccessTokenValidatorConfiguration{}
	this.Path = path
	return &this
}

// NewAccessTokenValidatorConfigurationWithDefaults instantiates a new AccessTokenValidatorConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessTokenValidatorConfigurationWithDefaults() *AccessTokenValidatorConfiguration {
	this := AccessTokenValidatorConfiguration{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AccessTokenValidatorConfiguration) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenValidatorConfiguration) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AccessTokenValidatorConfiguration) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AccessTokenValidatorConfiguration) SetDescription(v string) {
	o.Description = &v
}

// GetPath returns the Path field value
func (o *AccessTokenValidatorConfiguration) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *AccessTokenValidatorConfiguration) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *AccessTokenValidatorConfiguration) SetPath(v string) {
	o.Path = v
}

// GetSubjectAttributeName returns the SubjectAttributeName field value if set, zero value otherwise.
func (o *AccessTokenValidatorConfiguration) GetSubjectAttributeName() string {
	if o == nil || IsNil(o.SubjectAttributeName) {
		var ret string
		return ret
	}
	return *o.SubjectAttributeName
}

// GetSubjectAttributeNameOk returns a tuple with the SubjectAttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenValidatorConfiguration) GetSubjectAttributeNameOk() (*string, bool) {
	if o == nil || IsNil(o.SubjectAttributeName) {
		return nil, false
	}
	return o.SubjectAttributeName, true
}

// HasSubjectAttributeName returns a boolean if a field has been set.
func (o *AccessTokenValidatorConfiguration) HasSubjectAttributeName() bool {
	if o != nil && !IsNil(o.SubjectAttributeName) {
		return true
	}

	return false
}

// SetSubjectAttributeName gets a reference to the given string and assigns it to the SubjectAttributeName field.
func (o *AccessTokenValidatorConfiguration) SetSubjectAttributeName(v string) {
	o.SubjectAttributeName = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *AccessTokenValidatorConfiguration) GetIssuer() string {
	if o == nil || IsNil(o.Issuer) {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenValidatorConfiguration) GetIssuerOk() (*string, bool) {
	if o == nil || IsNil(o.Issuer) {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *AccessTokenValidatorConfiguration) HasIssuer() bool {
	if o != nil && !IsNil(o.Issuer) {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *AccessTokenValidatorConfiguration) SetIssuer(v string) {
	o.Issuer = &v
}

// GetAudience returns the Audience field value if set, zero value otherwise.
func (o *AccessTokenValidatorConfiguration) GetAudience() string {
	if o == nil || IsNil(o.Audience) {
		var ret string
		return ret
	}
	return *o.Audience
}

// GetAudienceOk returns a tuple with the Audience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenValidatorConfiguration) GetAudienceOk() (*string, bool) {
	if o == nil || IsNil(o.Audience) {
		return nil, false
	}
	return o.Audience, true
}

// HasAudience returns a boolean if a field has been set.
func (o *AccessTokenValidatorConfiguration) HasAudience() bool {
	if o != nil && !IsNil(o.Audience) {
		return true
	}

	return false
}

// SetAudience gets a reference to the given string and assigns it to the Audience field.
func (o *AccessTokenValidatorConfiguration) SetAudience(v string) {
	o.Audience = &v
}

func (o AccessTokenValidatorConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessTokenValidatorConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["path"] = o.Path
	if !IsNil(o.SubjectAttributeName) {
		toSerialize["subjectAttributeName"] = o.SubjectAttributeName
	}
	if !IsNil(o.Issuer) {
		toSerialize["issuer"] = o.Issuer
	}
	if !IsNil(o.Audience) {
		toSerialize["audience"] = o.Audience
	}
	return toSerialize, nil
}

type NullableAccessTokenValidatorConfiguration struct {
	value *AccessTokenValidatorConfiguration
	isSet bool
}

func (v NullableAccessTokenValidatorConfiguration) Get() *AccessTokenValidatorConfiguration {
	return v.value
}

func (v *NullableAccessTokenValidatorConfiguration) Set(val *AccessTokenValidatorConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessTokenValidatorConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessTokenValidatorConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessTokenValidatorConfiguration(val *AccessTokenValidatorConfiguration) *NullableAccessTokenValidatorConfiguration {
	return &NullableAccessTokenValidatorConfiguration{value: val, isSet: true}
}

func (v NullableAccessTokenValidatorConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessTokenValidatorConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}