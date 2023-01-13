# HostMultiValueSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HeaderNameList** | **[]string** | An array of header names used to identify the host source name. | 
**ListValueLocation** | [**ListValueLocation**](ListValueLocation.md) |  | 

## Methods

### NewHostMultiValueSource

`func NewHostMultiValueSource(headerNameList []string, listValueLocation ListValueLocation, ) *HostMultiValueSource`

NewHostMultiValueSource instantiates a new HostMultiValueSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostMultiValueSourceWithDefaults

`func NewHostMultiValueSourceWithDefaults() *HostMultiValueSource`

NewHostMultiValueSourceWithDefaults instantiates a new HostMultiValueSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaderNameList

`func (o *HostMultiValueSource) GetHeaderNameList() []string`

GetHeaderNameList returns the HeaderNameList field if non-nil, zero value otherwise.

### GetHeaderNameListOk

`func (o *HostMultiValueSource) GetHeaderNameListOk() (*[]string, bool)`

GetHeaderNameListOk returns a tuple with the HeaderNameList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderNameList

`func (o *HostMultiValueSource) SetHeaderNameList(v []string)`

SetHeaderNameList sets HeaderNameList field to given value.


### GetListValueLocation

`func (o *HostMultiValueSource) GetListValueLocation() ListValueLocation`

GetListValueLocation returns the ListValueLocation field if non-nil, zero value otherwise.

### GetListValueLocationOk

`func (o *HostMultiValueSource) GetListValueLocationOk() (*ListValueLocation, bool)`

GetListValueLocationOk returns a tuple with the ListValueLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListValueLocation

`func (o *HostMultiValueSource) SetListValueLocation(v ListValueLocation)`

SetListValueLocation sets ListValueLocation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


