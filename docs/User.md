# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the Administrative user. This value is read-only. | [optional] 
**Username** | **string** | (sortable) The Administrative users&#39;s username. | 
**Email** | Pointer to **string** | (sortable) The Administrative account&#39;s email. | [optional] 
**SlaAccepted** | Pointer to **bool** | (sortable) The Administrative account&#39;s sla acceptance indicator. | [optional] 
**FirstLogin** | Pointer to **bool** | (sortable) The Administrative account&#39;s first login indicator. | [optional] 
**ShowTutorial** | Pointer to **bool** | (sortable) The Administrative account&#39;s show tutorial indicator. | [optional] 

## Methods

### NewUser

`func NewUser(username string, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *User) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *User) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *User) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *User) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetSlaAccepted

`func (o *User) GetSlaAccepted() bool`

GetSlaAccepted returns the SlaAccepted field if non-nil, zero value otherwise.

### GetSlaAcceptedOk

`func (o *User) GetSlaAcceptedOk() (*bool, bool)`

GetSlaAcceptedOk returns a tuple with the SlaAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaAccepted

`func (o *User) SetSlaAccepted(v bool)`

SetSlaAccepted sets SlaAccepted field to given value.

### HasSlaAccepted

`func (o *User) HasSlaAccepted() bool`

HasSlaAccepted returns a boolean if a field has been set.

### GetFirstLogin

`func (o *User) GetFirstLogin() bool`

GetFirstLogin returns the FirstLogin field if non-nil, zero value otherwise.

### GetFirstLoginOk

`func (o *User) GetFirstLoginOk() (*bool, bool)`

GetFirstLoginOk returns a tuple with the FirstLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstLogin

`func (o *User) SetFirstLogin(v bool)`

SetFirstLogin sets FirstLogin field to given value.

### HasFirstLogin

`func (o *User) HasFirstLogin() bool`

HasFirstLogin returns a boolean if a field has been set.

### GetShowTutorial

`func (o *User) GetShowTutorial() bool`

GetShowTutorial returns the ShowTutorial field if non-nil, zero value otherwise.

### GetShowTutorialOk

`func (o *User) GetShowTutorialOk() (*bool, bool)`

GetShowTutorialOk returns a tuple with the ShowTutorial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTutorial

`func (o *User) SetShowTutorial(v bool)`

SetShowTutorial sets ShowTutorial field to given value.

### HasShowTutorial

`func (o *User) HasShowTutorial() bool`

HasShowTutorial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


