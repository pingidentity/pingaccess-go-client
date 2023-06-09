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

// checks if the HttpClientProxy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HttpClientProxy{}

// HttpClientProxy A proxy.
type HttpClientProxy struct {
	// When creating a new HttpClientProxy, this is the ID for the HttpClientProxy. If not specified, an ID will be automatically assigned. When updating an existing HttpClientProxy, this field is ignored and the ID is determined by the path parameter.
	Id *int64 `json:"id,omitempty"`
	// (sortable) The name of the proxy.
	Name string `json:"name"`
	// (sortable) The proxy host.
	Host string `json:"host"`
	// (sortable) The proxy port.
	Port int64 `json:"port"`
	// (sortable) A description of the proxy.
	Description *string `json:"description,omitempty"`
	// (sortable) True if the proxy requires authentication.
	RequiresAuthentication *bool `json:"requiresAuthentication,omitempty"`
	// (sortable) The username used for proxy authentication.
	Username *string      `json:"username,omitempty"`
	Password *HiddenField `json:"password,omitempty"`
}

// NewHttpClientProxy instantiates a new HttpClientProxy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpClientProxy(name string, host string, port int64) *HttpClientProxy {
	this := HttpClientProxy{}
	this.Name = name
	this.Host = host
	this.Port = port
	return &this
}

// NewHttpClientProxyWithDefaults instantiates a new HttpClientProxy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpClientProxyWithDefaults() *HttpClientProxy {
	this := HttpClientProxy{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HttpClientProxy) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpClientProxy) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HttpClientProxy) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *HttpClientProxy) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *HttpClientProxy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *HttpClientProxy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *HttpClientProxy) SetName(v string) {
	o.Name = v
}

// GetHost returns the Host field value
func (o *HttpClientProxy) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *HttpClientProxy) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *HttpClientProxy) SetHost(v string) {
	o.Host = v
}

// GetPort returns the Port field value
func (o *HttpClientProxy) GetPort() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *HttpClientProxy) GetPortOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *HttpClientProxy) SetPort(v int64) {
	o.Port = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *HttpClientProxy) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpClientProxy) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *HttpClientProxy) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *HttpClientProxy) SetDescription(v string) {
	o.Description = &v
}

// GetRequiresAuthentication returns the RequiresAuthentication field value if set, zero value otherwise.
func (o *HttpClientProxy) GetRequiresAuthentication() bool {
	if o == nil || IsNil(o.RequiresAuthentication) {
		var ret bool
		return ret
	}
	return *o.RequiresAuthentication
}

// GetRequiresAuthenticationOk returns a tuple with the RequiresAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpClientProxy) GetRequiresAuthenticationOk() (*bool, bool) {
	if o == nil || IsNil(o.RequiresAuthentication) {
		return nil, false
	}
	return o.RequiresAuthentication, true
}

// HasRequiresAuthentication returns a boolean if a field has been set.
func (o *HttpClientProxy) HasRequiresAuthentication() bool {
	if o != nil && !IsNil(o.RequiresAuthentication) {
		return true
	}

	return false
}

// SetRequiresAuthentication gets a reference to the given bool and assigns it to the RequiresAuthentication field.
func (o *HttpClientProxy) SetRequiresAuthentication(v bool) {
	o.RequiresAuthentication = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *HttpClientProxy) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpClientProxy) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *HttpClientProxy) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *HttpClientProxy) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *HttpClientProxy) GetPassword() HiddenField {
	if o == nil || IsNil(o.Password) {
		var ret HiddenField
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpClientProxy) GetPasswordOk() (*HiddenField, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *HttpClientProxy) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given HiddenField and assigns it to the Password field.
func (o *HttpClientProxy) SetPassword(v HiddenField) {
	o.Password = &v
}

func (o HttpClientProxy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HttpClientProxy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	toSerialize["host"] = o.Host
	toSerialize["port"] = o.Port
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.RequiresAuthentication) {
		toSerialize["requiresAuthentication"] = o.RequiresAuthentication
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	return toSerialize, nil
}

type NullableHttpClientProxy struct {
	value *HttpClientProxy
	isSet bool
}

func (v NullableHttpClientProxy) Get() *HttpClientProxy {
	return v.value
}

func (v *NullableHttpClientProxy) Set(val *HttpClientProxy) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpClientProxy) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpClientProxy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpClientProxy(val *HttpClientProxy) *NullableHttpClientProxy {
	return &NullableHttpClientProxy{value: val, isSet: true}
}

func (v NullableHttpClientProxy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpClientProxy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
