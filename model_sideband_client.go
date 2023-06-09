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

// checks if the SidebandClient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SidebandClient{}

// SidebandClient A sideband client.
type SidebandClient struct {
	// When creating a new SidebandClient, this is the ID for the SidebandClient. If not specified, an ID will be automatically assigned. When updating an existing SidebandClient, this field is ignored and the ID is determined by the path parameter.
	Id *string `json:"id,omitempty"`
	// (sortable) Name of the sideband client.
	Name string `json:"name"`
	// (sortable) Description of the sideband client.
	Description *string `json:"description,omitempty"`
	// The authentication configuration for the sideband client.
	ClientCredentials []SidebandClientCredentials `json:"clientCredentials"`
}

// NewSidebandClient instantiates a new SidebandClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSidebandClient(name string, clientCredentials []SidebandClientCredentials) *SidebandClient {
	this := SidebandClient{}
	this.Name = name
	this.ClientCredentials = clientCredentials
	return &this
}

// NewSidebandClientWithDefaults instantiates a new SidebandClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSidebandClientWithDefaults() *SidebandClient {
	this := SidebandClient{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SidebandClient) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SidebandClient) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SidebandClient) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SidebandClient) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *SidebandClient) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SidebandClient) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SidebandClient) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SidebandClient) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SidebandClient) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SidebandClient) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SidebandClient) SetDescription(v string) {
	o.Description = &v
}

// GetClientCredentials returns the ClientCredentials field value
func (o *SidebandClient) GetClientCredentials() []SidebandClientCredentials {
	if o == nil {
		var ret []SidebandClientCredentials
		return ret
	}

	return o.ClientCredentials
}

// GetClientCredentialsOk returns a tuple with the ClientCredentials field value
// and a boolean to check if the value has been set.
func (o *SidebandClient) GetClientCredentialsOk() ([]SidebandClientCredentials, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientCredentials, true
}

// SetClientCredentials sets field value
func (o *SidebandClient) SetClientCredentials(v []SidebandClientCredentials) {
	o.ClientCredentials = v
}

func (o SidebandClient) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SidebandClient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["clientCredentials"] = o.ClientCredentials
	return toSerialize, nil
}

type NullableSidebandClient struct {
	value *SidebandClient
	isSet bool
}

func (v NullableSidebandClient) Get() *SidebandClient {
	return v.value
}

func (v *NullableSidebandClient) Set(val *SidebandClient) {
	v.value = val
	v.isSet = true
}

func (v NullableSidebandClient) IsSet() bool {
	return v.isSet
}

func (v *NullableSidebandClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSidebandClient(val *SidebandClient) *NullableSidebandClient {
	return &NullableSidebandClient{value: val, isSet: true}
}

func (v NullableSidebandClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSidebandClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
