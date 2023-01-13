# IpMultiValueSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HeaderNameList** | **[]string** | An array of header names used to identify the source IP address. | 
**ListValueLocation** | [**ListValueLocation**](ListValueLocation.md) |  | 
**FallbackToLastHopIp** | Pointer to **bool** | Indicator if the upstream IP address should be used as the client address if none of the expected headers is returned. | [optional] 

## Methods

### NewIpMultiValueSource

`func NewIpMultiValueSource(headerNameList []string, listValueLocation ListValueLocation, ) *IpMultiValueSource`

NewIpMultiValueSource instantiates a new IpMultiValueSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpMultiValueSourceWithDefaults

`func NewIpMultiValueSourceWithDefaults() *IpMultiValueSource`

NewIpMultiValueSourceWithDefaults instantiates a new IpMultiValueSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaderNameList

`func (o *IpMultiValueSource) GetHeaderNameList() []string`

GetHeaderNameList returns the HeaderNameList field if non-nil, zero value otherwise.

### GetHeaderNameListOk

`func (o *IpMultiValueSource) GetHeaderNameListOk() (*[]string, bool)`

GetHeaderNameListOk returns a tuple with the HeaderNameList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderNameList

`func (o *IpMultiValueSource) SetHeaderNameList(v []string)`

SetHeaderNameList sets HeaderNameList field to given value.


### GetListValueLocation

`func (o *IpMultiValueSource) GetListValueLocation() ListValueLocation`

GetListValueLocation returns the ListValueLocation field if non-nil, zero value otherwise.

### GetListValueLocationOk

`func (o *IpMultiValueSource) GetListValueLocationOk() (*ListValueLocation, bool)`

GetListValueLocationOk returns a tuple with the ListValueLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListValueLocation

`func (o *IpMultiValueSource) SetListValueLocation(v ListValueLocation)`

SetListValueLocation sets ListValueLocation field to given value.


### GetFallbackToLastHopIp

`func (o *IpMultiValueSource) GetFallbackToLastHopIp() bool`

GetFallbackToLastHopIp returns the FallbackToLastHopIp field if non-nil, zero value otherwise.

### GetFallbackToLastHopIpOk

`func (o *IpMultiValueSource) GetFallbackToLastHopIpOk() (*bool, bool)`

GetFallbackToLastHopIpOk returns a tuple with the FallbackToLastHopIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackToLastHopIp

`func (o *IpMultiValueSource) SetFallbackToLastHopIp(v bool)`

SetFallbackToLastHopIp sets FallbackToLastHopIp field to given value.

### HasFallbackToLastHopIp

`func (o *IpMultiValueSource) HasFallbackToLastHopIp() bool`

HasFallbackToLastHopIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


