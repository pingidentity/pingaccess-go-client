# ApiError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flash** | **[]string** | The flash message giving a general error message. | 
**Form** | **map[string][]string** | The errors on specific configuration fields defined by the API request.  Key type: String Value type: String[] | 

## Methods

### NewApiError

`func NewApiError(flash []string, form map[string][]string, ) *ApiError`

NewApiError instantiates a new ApiError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiErrorWithDefaults

`func NewApiErrorWithDefaults() *ApiError`

NewApiErrorWithDefaults instantiates a new ApiError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlash

`func (o *ApiError) GetFlash() []string`

GetFlash returns the Flash field if non-nil, zero value otherwise.

### GetFlashOk

`func (o *ApiError) GetFlashOk() (*[]string, bool)`

GetFlashOk returns a tuple with the Flash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlash

`func (o *ApiError) SetFlash(v []string)`

SetFlash sets Flash field to given value.


### GetForm

`func (o *ApiError) GetForm() map[string][]string`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *ApiError) GetFormOk() (*map[string][]string, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *ApiError) SetForm(v map[string][]string)`

SetForm sets Form field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


