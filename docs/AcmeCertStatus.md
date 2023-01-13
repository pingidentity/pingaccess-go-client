# AcmeCertStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | [**AcmeCertState**](AcmeCertState.md) |  | 
**Problems** | Pointer to [**map[string]ProblemDocument**](ProblemDocument.md) | A map of problem documents for requested domains. The key is the domain and the value is a ProblemDocumentView. | [optional] 

## Methods

### NewAcmeCertStatus

`func NewAcmeCertStatus(state AcmeCertState, ) *AcmeCertStatus`

NewAcmeCertStatus instantiates a new AcmeCertStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcmeCertStatusWithDefaults

`func NewAcmeCertStatusWithDefaults() *AcmeCertStatus`

NewAcmeCertStatusWithDefaults instantiates a new AcmeCertStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *AcmeCertStatus) GetState() AcmeCertState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AcmeCertStatus) GetStateOk() (*AcmeCertState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AcmeCertStatus) SetState(v AcmeCertState)`

SetState sets State field to given value.


### GetProblems

`func (o *AcmeCertStatus) GetProblems() map[string]ProblemDocument`

GetProblems returns the Problems field if non-nil, zero value otherwise.

### GetProblemsOk

`func (o *AcmeCertStatus) GetProblemsOk() (*map[string]ProblemDocument, bool)`

GetProblemsOk returns a tuple with the Problems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblems

`func (o *AcmeCertStatus) SetProblems(v map[string]ProblemDocument)`

SetProblems sets Problems field to given value.

### HasProblems

`func (o *AcmeCertStatus) HasProblems() bool`

HasProblems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


