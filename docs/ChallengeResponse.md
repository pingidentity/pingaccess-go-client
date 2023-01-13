# ChallengeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Generator** | [**ChallengeResponseGenerator**](ChallengeResponseGenerator.md) |  | 
**Filter** | Pointer to [**ChallengeResponseFilter**](ChallengeResponseFilter.md) |  | [optional] 

## Methods

### NewChallengeResponse

`func NewChallengeResponse(generator ChallengeResponseGenerator, ) *ChallengeResponse`

NewChallengeResponse instantiates a new ChallengeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChallengeResponseWithDefaults

`func NewChallengeResponseWithDefaults() *ChallengeResponse`

NewChallengeResponseWithDefaults instantiates a new ChallengeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGenerator

`func (o *ChallengeResponse) GetGenerator() ChallengeResponseGenerator`

GetGenerator returns the Generator field if non-nil, zero value otherwise.

### GetGeneratorOk

`func (o *ChallengeResponse) GetGeneratorOk() (*ChallengeResponseGenerator, bool)`

GetGeneratorOk returns a tuple with the Generator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerator

`func (o *ChallengeResponse) SetGenerator(v ChallengeResponseGenerator)`

SetGenerator sets Generator field to given value.


### GetFilter

`func (o *ChallengeResponse) GetFilter() ChallengeResponseFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ChallengeResponse) GetFilterOk() (*ChallengeResponseFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ChallengeResponse) SetFilter(v ChallengeResponseFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ChallengeResponse) HasFilter() bool`

HasFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


