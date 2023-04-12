# LogSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogCategories** | [**[]Category**](Category.md) | A list of logging categories. | 

## Methods

### NewLogSettings

`func NewLogSettings(logCategories []Category, ) *LogSettings`

NewLogSettings instantiates a new LogSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogSettingsWithDefaults

`func NewLogSettingsWithDefaults() *LogSettings`

NewLogSettingsWithDefaults instantiates a new LogSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogCategories

`func (o *LogSettings) GetLogCategories() []Category`

GetLogCategories returns the LogCategories field if non-nil, zero value otherwise.

### GetLogCategoriesOk

`func (o *LogSettings) GetLogCategoriesOk() (*[]Category, bool)`

GetLogCategoriesOk returns a tuple with the LogCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogCategories

`func (o *LogSettings) SetLogCategories(v []Category)`

SetLogCategories sets LogCategories field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


