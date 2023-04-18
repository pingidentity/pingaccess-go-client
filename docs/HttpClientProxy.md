# HttpClientProxy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | When creating a new HttpClientProxy, this is the ID for the HttpClientProxy. If not specified, an ID will be automatically assigned. When updating an existing HttpClientProxy, this field is ignored and the ID is determined by the path parameter. | [optional] 
**Name** | **string** | (sortable) The name of the proxy. | 
**Host** | **string** | (sortable) The proxy host. | 
**Port** | **int64** | (sortable) The proxy port. | 
**Description** | Pointer to **string** | (sortable) A description of the proxy. | [optional] 
**RequiresAuthentication** | Pointer to **bool** | (sortable) True if the proxy requires authentication. | [optional] 
**Username** | Pointer to **string** | (sortable) The username used for proxy authentication. | [optional] 
**Password** | Pointer to [**HiddenField**](HiddenField.md) |  | [optional] 

## Methods

### NewHttpClientProxy

`func NewHttpClientProxy(name string, host string, port int64, ) *HttpClientProxy`

NewHttpClientProxy instantiates a new HttpClientProxy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpClientProxyWithDefaults

`func NewHttpClientProxyWithDefaults() *HttpClientProxy`

NewHttpClientProxyWithDefaults instantiates a new HttpClientProxy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HttpClientProxy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HttpClientProxy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HttpClientProxy) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *HttpClientProxy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *HttpClientProxy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HttpClientProxy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HttpClientProxy) SetName(v string)`

SetName sets Name field to given value.


### GetHost

`func (o *HttpClientProxy) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HttpClientProxy) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HttpClientProxy) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *HttpClientProxy) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HttpClientProxy) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HttpClientProxy) SetPort(v int64)`

SetPort sets Port field to given value.


### GetDescription

`func (o *HttpClientProxy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HttpClientProxy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HttpClientProxy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HttpClientProxy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRequiresAuthentication

`func (o *HttpClientProxy) GetRequiresAuthentication() bool`

GetRequiresAuthentication returns the RequiresAuthentication field if non-nil, zero value otherwise.

### GetRequiresAuthenticationOk

`func (o *HttpClientProxy) GetRequiresAuthenticationOk() (*bool, bool)`

GetRequiresAuthenticationOk returns a tuple with the RequiresAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresAuthentication

`func (o *HttpClientProxy) SetRequiresAuthentication(v bool)`

SetRequiresAuthentication sets RequiresAuthentication field to given value.

### HasRequiresAuthentication

`func (o *HttpClientProxy) HasRequiresAuthentication() bool`

HasRequiresAuthentication returns a boolean if a field has been set.

### GetUsername

`func (o *HttpClientProxy) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *HttpClientProxy) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *HttpClientProxy) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *HttpClientProxy) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *HttpClientProxy) GetPassword() HiddenField`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *HttpClientProxy) GetPasswordOk() (*HiddenField, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *HttpClientProxy) SetPassword(v HiddenField)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *HttpClientProxy) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


