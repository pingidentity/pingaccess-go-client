# AccessTokenValidatorConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Path** | **string** |  | 
**SubjectAttributeName** | Pointer to **string** |  | [optional] 
**Issuer** | Pointer to **string** |  | [optional] 
**Audience** | Pointer to **string** |  | [optional] 

## Methods

### NewAccessTokenValidatorConfiguration

`func NewAccessTokenValidatorConfiguration(path string, ) *AccessTokenValidatorConfiguration`

NewAccessTokenValidatorConfiguration instantiates a new AccessTokenValidatorConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenValidatorConfigurationWithDefaults

`func NewAccessTokenValidatorConfigurationWithDefaults() *AccessTokenValidatorConfiguration`

NewAccessTokenValidatorConfigurationWithDefaults instantiates a new AccessTokenValidatorConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AccessTokenValidatorConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccessTokenValidatorConfiguration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccessTokenValidatorConfiguration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccessTokenValidatorConfiguration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPath

`func (o *AccessTokenValidatorConfiguration) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AccessTokenValidatorConfiguration) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AccessTokenValidatorConfiguration) SetPath(v string)`

SetPath sets Path field to given value.


### GetSubjectAttributeName

`func (o *AccessTokenValidatorConfiguration) GetSubjectAttributeName() string`

GetSubjectAttributeName returns the SubjectAttributeName field if non-nil, zero value otherwise.

### GetSubjectAttributeNameOk

`func (o *AccessTokenValidatorConfiguration) GetSubjectAttributeNameOk() (*string, bool)`

GetSubjectAttributeNameOk returns a tuple with the SubjectAttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAttributeName

`func (o *AccessTokenValidatorConfiguration) SetSubjectAttributeName(v string)`

SetSubjectAttributeName sets SubjectAttributeName field to given value.

### HasSubjectAttributeName

`func (o *AccessTokenValidatorConfiguration) HasSubjectAttributeName() bool`

HasSubjectAttributeName returns a boolean if a field has been set.

### GetIssuer

`func (o *AccessTokenValidatorConfiguration) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *AccessTokenValidatorConfiguration) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *AccessTokenValidatorConfiguration) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *AccessTokenValidatorConfiguration) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetAudience

`func (o *AccessTokenValidatorConfiguration) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *AccessTokenValidatorConfiguration) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *AccessTokenValidatorConfiguration) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *AccessTokenValidatorConfiguration) HasAudience() bool`

HasAudience returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


