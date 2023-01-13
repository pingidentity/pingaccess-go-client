# AcmeServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | When creating a new AcmeServer, this is the ID for the AcmeServer. If not specified, an ID will be automatically assigned. When updating an existing AcmeServer, this field is ignored and the ID is determined by the path parameter. | [optional] 
**Name** | **string** | (sortable) A user-friendly name for the ACME server. | 
**Url** | **string** | The URL of the ACME directory resource on the ACME server. | 
**AcmeAccounts** | Pointer to [**[]Link**](Link.md) | An array of references to accounts. This array is read-only. | [optional] 

## Methods

### NewAcmeServer

`func NewAcmeServer(name string, url string, ) *AcmeServer`

NewAcmeServer instantiates a new AcmeServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcmeServerWithDefaults

`func NewAcmeServerWithDefaults() *AcmeServer`

NewAcmeServerWithDefaults instantiates a new AcmeServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AcmeServer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AcmeServer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AcmeServer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AcmeServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AcmeServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AcmeServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AcmeServer) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *AcmeServer) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AcmeServer) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AcmeServer) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetAcmeAccounts

`func (o *AcmeServer) GetAcmeAccounts() []Link`

GetAcmeAccounts returns the AcmeAccounts field if non-nil, zero value otherwise.

### GetAcmeAccountsOk

`func (o *AcmeServer) GetAcmeAccountsOk() (*[]Link, bool)`

GetAcmeAccountsOk returns a tuple with the AcmeAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeAccounts

`func (o *AcmeServer) SetAcmeAccounts(v []Link)`

SetAcmeAccounts sets AcmeAccounts field to given value.

### HasAcmeAccounts

`func (o *AcmeServer) HasAcmeAccounts() bool`

HasAcmeAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


