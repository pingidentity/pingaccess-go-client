# AcmeServers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]AcmeServer**](AcmeServer.md) | An array of ACME servers. | 

## Methods

### NewAcmeServers

`func NewAcmeServers(items []AcmeServer, ) *AcmeServers`

NewAcmeServers instantiates a new AcmeServers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcmeServersWithDefaults

`func NewAcmeServersWithDefaults() *AcmeServers`

NewAcmeServersWithDefaults instantiates a new AcmeServers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *AcmeServers) GetItems() []AcmeServer`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AcmeServers) GetItemsOk() (*[]AcmeServer, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AcmeServers) SetItems(v []AcmeServer)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


