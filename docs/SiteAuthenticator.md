# SiteAuthenticator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassName** | **string** | (sortable) The site authenticator&#39;s class name. | 
**Id** | Pointer to **int64** | When creating a new SiteAuthenticator, this is the ID for the SiteAuthenticator. If not specified, an ID will be automatically assigned. When updating an existing SiteAuthenticator, this field is ignored and the ID is determined by the path parameter. | [optional] 
**Name** | **string** | (sortable) The site authenticator&#39;s name. | 
**Configuration** | Pointer to **map[string]interface{}** | The site authenticator&#39;s configuration data. - This value is a PingAccess plugin configuration (JSON). | [optional] 

## Methods

### NewSiteAuthenticator

`func NewSiteAuthenticator(className string, name string, ) *SiteAuthenticator`

NewSiteAuthenticator instantiates a new SiteAuthenticator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteAuthenticatorWithDefaults

`func NewSiteAuthenticatorWithDefaults() *SiteAuthenticator`

NewSiteAuthenticatorWithDefaults instantiates a new SiteAuthenticator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassName

`func (o *SiteAuthenticator) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *SiteAuthenticator) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *SiteAuthenticator) SetClassName(v string)`

SetClassName sets ClassName field to given value.


### GetId

`func (o *SiteAuthenticator) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SiteAuthenticator) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SiteAuthenticator) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SiteAuthenticator) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SiteAuthenticator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SiteAuthenticator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SiteAuthenticator) SetName(v string)`

SetName sets Name field to given value.


### GetConfiguration

`func (o *SiteAuthenticator) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *SiteAuthenticator) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *SiteAuthenticator) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *SiteAuthenticator) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


