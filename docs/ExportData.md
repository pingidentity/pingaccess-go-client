# ExportData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | The version of PingAccess that was exported. | [optional] 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**EncryptionKey** | Pointer to [**JsonWebKey**](JsonWebKey.md) |  | [optional] 
**MasterKeys** | Pointer to [**MasterKeys**](MasterKeys.md) |  | [optional] 

## Methods

### NewExportData

`func NewExportData() *ExportData`

NewExportData instantiates a new ExportData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportDataWithDefaults

`func NewExportDataWithDefaults() *ExportData`

NewExportDataWithDefaults instantiates a new ExportData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ExportData) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExportData) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExportData) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ExportData) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetData

`func (o *ExportData) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ExportData) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ExportData) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ExportData) HasData() bool`

HasData returns a boolean if a field has been set.

### GetEncryptionKey

`func (o *ExportData) GetEncryptionKey() JsonWebKey`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *ExportData) GetEncryptionKeyOk() (*JsonWebKey, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *ExportData) SetEncryptionKey(v JsonWebKey)`

SetEncryptionKey sets EncryptionKey field to given value.

### HasEncryptionKey

`func (o *ExportData) HasEncryptionKey() bool`

HasEncryptionKey returns a boolean if a field has been set.

### GetMasterKeys

`func (o *ExportData) GetMasterKeys() MasterKeys`

GetMasterKeys returns the MasterKeys field if non-nil, zero value otherwise.

### GetMasterKeysOk

`func (o *ExportData) GetMasterKeysOk() (*MasterKeys, bool)`

GetMasterKeysOk returns a tuple with the MasterKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterKeys

`func (o *ExportData) SetMasterKeys(v MasterKeys)`

SetMasterKeys sets MasterKeys field to given value.

### HasMasterKeys

`func (o *ExportData) HasMasterKeys() bool`

HasMasterKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


