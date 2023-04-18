# OidcConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthnReqListId** | Pointer to **int64** | The ID of the authentication requirement list for Administrative SSO login to PingAccess. | [optional] 
**Enabled** | Pointer to **bool** | This field is true to enable Administrator SSO Authentication. | [optional] 
**UseSlo** | Pointer to **bool** | Enable if OIDC single log out should be used on the /pa/oidc/logout for admin console. | [optional] 
**UsernameAttributeName** | Pointer to **string** | Attribute to display as the logged in user. If not set, the sub attribute will be used. | [optional] 
**RoleMapping** | Pointer to [**RoleMappingConfiguration**](RoleMappingConfiguration.md) |  | [optional] 
**OidcConfiguration** | [**AdminWebSessionOidcConfiguration**](AdminWebSessionOidcConfiguration.md) |  | 

## Methods

### NewOidcConfig

`func NewOidcConfig(oidcConfiguration AdminWebSessionOidcConfiguration, ) *OidcConfig`

NewOidcConfig instantiates a new OidcConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOidcConfigWithDefaults

`func NewOidcConfigWithDefaults() *OidcConfig`

NewOidcConfigWithDefaults instantiates a new OidcConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthnReqListId

`func (o *OidcConfig) GetAuthnReqListId() int64`

GetAuthnReqListId returns the AuthnReqListId field if non-nil, zero value otherwise.

### GetAuthnReqListIdOk

`func (o *OidcConfig) GetAuthnReqListIdOk() (*int64, bool)`

GetAuthnReqListIdOk returns a tuple with the AuthnReqListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthnReqListId

`func (o *OidcConfig) SetAuthnReqListId(v int64)`

SetAuthnReqListId sets AuthnReqListId field to given value.

### HasAuthnReqListId

`func (o *OidcConfig) HasAuthnReqListId() bool`

HasAuthnReqListId returns a boolean if a field has been set.

### GetEnabled

`func (o *OidcConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OidcConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OidcConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OidcConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUseSlo

`func (o *OidcConfig) GetUseSlo() bool`

GetUseSlo returns the UseSlo field if non-nil, zero value otherwise.

### GetUseSloOk

`func (o *OidcConfig) GetUseSloOk() (*bool, bool)`

GetUseSloOk returns a tuple with the UseSlo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSlo

`func (o *OidcConfig) SetUseSlo(v bool)`

SetUseSlo sets UseSlo field to given value.

### HasUseSlo

`func (o *OidcConfig) HasUseSlo() bool`

HasUseSlo returns a boolean if a field has been set.

### GetUsernameAttributeName

`func (o *OidcConfig) GetUsernameAttributeName() string`

GetUsernameAttributeName returns the UsernameAttributeName field if non-nil, zero value otherwise.

### GetUsernameAttributeNameOk

`func (o *OidcConfig) GetUsernameAttributeNameOk() (*string, bool)`

GetUsernameAttributeNameOk returns a tuple with the UsernameAttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameAttributeName

`func (o *OidcConfig) SetUsernameAttributeName(v string)`

SetUsernameAttributeName sets UsernameAttributeName field to given value.

### HasUsernameAttributeName

`func (o *OidcConfig) HasUsernameAttributeName() bool`

HasUsernameAttributeName returns a boolean if a field has been set.

### GetRoleMapping

`func (o *OidcConfig) GetRoleMapping() RoleMappingConfiguration`

GetRoleMapping returns the RoleMapping field if non-nil, zero value otherwise.

### GetRoleMappingOk

`func (o *OidcConfig) GetRoleMappingOk() (*RoleMappingConfiguration, bool)`

GetRoleMappingOk returns a tuple with the RoleMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleMapping

`func (o *OidcConfig) SetRoleMapping(v RoleMappingConfiguration)`

SetRoleMapping sets RoleMapping field to given value.

### HasRoleMapping

`func (o *OidcConfig) HasRoleMapping() bool`

HasRoleMapping returns a boolean if a field has been set.

### GetOidcConfiguration

`func (o *OidcConfig) GetOidcConfiguration() AdminWebSessionOidcConfiguration`

GetOidcConfiguration returns the OidcConfiguration field if non-nil, zero value otherwise.

### GetOidcConfigurationOk

`func (o *OidcConfig) GetOidcConfigurationOk() (*AdminWebSessionOidcConfiguration, bool)`

GetOidcConfigurationOk returns a tuple with the OidcConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcConfiguration

`func (o *OidcConfig) SetOidcConfiguration(v AdminWebSessionOidcConfiguration)`

SetOidcConfiguration sets OidcConfiguration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


