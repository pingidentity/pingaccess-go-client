# OIDCProviderPlugin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassName** | **string** | The class name of the OpenID Connect provider plugin. | 
**Configuration** | Pointer to **map[string]interface{}** | The OpenID Connect provider plugin&#39;s configuration data. - This value is a PingAccess plugin configuration (JSON). | [optional] 

## Methods

### NewOIDCProviderPlugin

`func NewOIDCProviderPlugin(className string, ) *OIDCProviderPlugin`

NewOIDCProviderPlugin instantiates a new OIDCProviderPlugin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOIDCProviderPluginWithDefaults

`func NewOIDCProviderPluginWithDefaults() *OIDCProviderPlugin`

NewOIDCProviderPluginWithDefaults instantiates a new OIDCProviderPlugin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassName

`func (o *OIDCProviderPlugin) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *OIDCProviderPlugin) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *OIDCProviderPlugin) SetClassName(v string)`

SetClassName sets ClassName field to given value.


### GetConfiguration

`func (o *OIDCProviderPlugin) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *OIDCProviderPlugin) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *OIDCProviderPlugin) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *OIDCProviderPlugin) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


