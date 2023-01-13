# SidebandClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | When creating a new SidebandClient, this is the ID for the SidebandClient. If not specified, an ID will be automatically assigned. When updating an existing SidebandClient, this field is ignored and the ID is determined by the path parameter. | [optional] 
**Name** | **string** | (sortable) Name of the sideband client. | 
**Description** | Pointer to **string** | (sortable) Description of the sideband client. | [optional] 
**ClientCredentials** | [**[]SidebandClientCredentials**](SidebandClientCredentials.md) | The authentication configuration for the sideband client. | 

## Methods

### NewSidebandClient

`func NewSidebandClient(name string, clientCredentials []SidebandClientCredentials, ) *SidebandClient`

NewSidebandClient instantiates a new SidebandClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSidebandClientWithDefaults

`func NewSidebandClientWithDefaults() *SidebandClient`

NewSidebandClientWithDefaults instantiates a new SidebandClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SidebandClient) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SidebandClient) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SidebandClient) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SidebandClient) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SidebandClient) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SidebandClient) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SidebandClient) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SidebandClient) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SidebandClient) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SidebandClient) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SidebandClient) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetClientCredentials

`func (o *SidebandClient) GetClientCredentials() []SidebandClientCredentials`

GetClientCredentials returns the ClientCredentials field if non-nil, zero value otherwise.

### GetClientCredentialsOk

`func (o *SidebandClient) GetClientCredentialsOk() (*[]SidebandClientCredentials, bool)`

GetClientCredentialsOk returns a tuple with the ClientCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCredentials

`func (o *SidebandClient) SetClientCredentials(v []SidebandClientCredentials)`

SetClientCredentials sets ClientCredentials field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


