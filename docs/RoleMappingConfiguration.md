# RoleMappingConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Set to true to enable mapping of user attributes to roles. | [optional] 
**Administrator** | Pointer to [**RequiredAttributeMapping**](RequiredAttributeMapping.md) |  | [optional] 
**PlatformAdmin** | Pointer to [**OptionalAttributeMapping**](OptionalAttributeMapping.md) |  | [optional] 
**Auditor** | Pointer to [**OptionalAttributeMapping**](OptionalAttributeMapping.md) |  | [optional] 

## Methods

### NewRoleMappingConfiguration

`func NewRoleMappingConfiguration() *RoleMappingConfiguration`

NewRoleMappingConfiguration instantiates a new RoleMappingConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleMappingConfigurationWithDefaults

`func NewRoleMappingConfigurationWithDefaults() *RoleMappingConfiguration`

NewRoleMappingConfigurationWithDefaults instantiates a new RoleMappingConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *RoleMappingConfiguration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RoleMappingConfiguration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RoleMappingConfiguration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RoleMappingConfiguration) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAdministrator

`func (o *RoleMappingConfiguration) GetAdministrator() RequiredAttributeMapping`

GetAdministrator returns the Administrator field if non-nil, zero value otherwise.

### GetAdministratorOk

`func (o *RoleMappingConfiguration) GetAdministratorOk() (*RequiredAttributeMapping, bool)`

GetAdministratorOk returns a tuple with the Administrator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrator

`func (o *RoleMappingConfiguration) SetAdministrator(v RequiredAttributeMapping)`

SetAdministrator sets Administrator field to given value.

### HasAdministrator

`func (o *RoleMappingConfiguration) HasAdministrator() bool`

HasAdministrator returns a boolean if a field has been set.

### GetPlatformAdmin

`func (o *RoleMappingConfiguration) GetPlatformAdmin() OptionalAttributeMapping`

GetPlatformAdmin returns the PlatformAdmin field if non-nil, zero value otherwise.

### GetPlatformAdminOk

`func (o *RoleMappingConfiguration) GetPlatformAdminOk() (*OptionalAttributeMapping, bool)`

GetPlatformAdminOk returns a tuple with the PlatformAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAdmin

`func (o *RoleMappingConfiguration) SetPlatformAdmin(v OptionalAttributeMapping)`

SetPlatformAdmin sets PlatformAdmin field to given value.

### HasPlatformAdmin

`func (o *RoleMappingConfiguration) HasPlatformAdmin() bool`

HasPlatformAdmin returns a boolean if a field has been set.

### GetAuditor

`func (o *RoleMappingConfiguration) GetAuditor() OptionalAttributeMapping`

GetAuditor returns the Auditor field if non-nil, zero value otherwise.

### GetAuditorOk

`func (o *RoleMappingConfiguration) GetAuditorOk() (*OptionalAttributeMapping, bool)`

GetAuditorOk returns a tuple with the Auditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditor

`func (o *RoleMappingConfiguration) SetAuditor(v OptionalAttributeMapping)`

SetAuditor sets Auditor field to given value.

### HasAuditor

`func (o *RoleMappingConfiguration) HasAuditor() bool`

HasAuditor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


