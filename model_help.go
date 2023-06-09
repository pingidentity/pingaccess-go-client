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

// checks if the Help type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Help{}

// Help struct for Help
type Help struct {
	// The help title of the configuration field.
	Title string `json:"title"`
	// The help content of the configuration field.
	Content string `json:"content"`
	// The help URL for the configuration field.
	Url string `json:"url"`
}

// NewHelp instantiates a new Help object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHelp(title string, content string, url string) *Help {
	this := Help{}
	this.Title = title
	this.Content = content
	this.Url = url
	return &this
}

// NewHelpWithDefaults instantiates a new Help object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHelpWithDefaults() *Help {
	this := Help{}
	return &this
}

// GetTitle returns the Title field value
func (o *Help) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *Help) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *Help) SetTitle(v string) {
	o.Title = v
}

// GetContent returns the Content field value
func (o *Help) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *Help) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *Help) SetContent(v string) {
	o.Content = v
}

// GetUrl returns the Url field value
func (o *Help) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Help) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Help) SetUrl(v string) {
	o.Url = v
}

func (o Help) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Help) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	toSerialize["content"] = o.Content
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

type NullableHelp struct {
	value *Help
	isSet bool
}

func (v NullableHelp) Get() *Help {
	return v.value
}

func (v *NullableHelp) Set(val *Help) {
	v.value = val
	v.isSet = true
}

func (v NullableHelp) IsSet() bool {
	return v.isSet
}

func (v *NullableHelp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHelp(val *Help) *NullableHelp {
	return &NullableHelp{value: val, isSet: true}
}

func (v NullableHelp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHelp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
