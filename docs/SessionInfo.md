# SessionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exp** | **int64** | Time at which session will expire due to inactivity. | 
**Iat** | **int64** | Time at which session was created. | 
**ExpWarn** | **int64** | Length of time before a timeout at which warning should appear. | 
**SesTimeout** | **int64** | Maximum length of time that a session is allowed to live, regardless of user activity, -1 indicates disabled. | 
**Sub** | **string** | The Session&#39;s subject. | 
**Flash** | **string** | The Warning message. | 
**PollIntervalSeconds** | **int32** | Session poll interval configuration in seconds. | 
**Roles** | [**[]Role**](Role.md) | The user&#39;s roles. | 
**AccessControlDirectives** | [**[]AdminAccessControlDirective**](AdminAccessControlDirective.md) | The set of access control directives. | 
**UseSlo** | **bool** | Indicates whether single log out (SLO) is enabled or not. | 
**FipsMode** | **bool** | Indicates whether FIPS mode is enabled or not. | 
**DisplayLogSettings** | **bool** |  | 
**ConfigurationExports** | [**ConfigStatuses**](ConfigStatuses.md) |  | 
**ConfigurationImports** | [**ConfigStatuses**](ConfigStatuses.md) |  | 
**SniEnabled** | **bool** | Indicates that SNI is enabled or not. | 
**MaxFileUploadSize** | **int32** | The maximum file upload size in bytes. | 
**ShowWarning** | **bool** | Indicates that a warning needs to be shown or not. | 

## Methods

### NewSessionInfo

`func NewSessionInfo(exp int64, iat int64, expWarn int64, sesTimeout int64, sub string, flash string, pollIntervalSeconds int32, roles []Role, accessControlDirectives []AdminAccessControlDirective, useSlo bool, fipsMode bool, displayLogSettings bool, configurationExports ConfigStatuses, configurationImports ConfigStatuses, sniEnabled bool, maxFileUploadSize int32, showWarning bool, ) *SessionInfo`

NewSessionInfo instantiates a new SessionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionInfoWithDefaults

`func NewSessionInfoWithDefaults() *SessionInfo`

NewSessionInfoWithDefaults instantiates a new SessionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExp

`func (o *SessionInfo) GetExp() int64`

GetExp returns the Exp field if non-nil, zero value otherwise.

### GetExpOk

`func (o *SessionInfo) GetExpOk() (*int64, bool)`

GetExpOk returns a tuple with the Exp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExp

`func (o *SessionInfo) SetExp(v int64)`

SetExp sets Exp field to given value.


### GetIat

`func (o *SessionInfo) GetIat() int64`

GetIat returns the Iat field if non-nil, zero value otherwise.

### GetIatOk

`func (o *SessionInfo) GetIatOk() (*int64, bool)`

GetIatOk returns a tuple with the Iat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIat

`func (o *SessionInfo) SetIat(v int64)`

SetIat sets Iat field to given value.


### GetExpWarn

`func (o *SessionInfo) GetExpWarn() int64`

GetExpWarn returns the ExpWarn field if non-nil, zero value otherwise.

### GetExpWarnOk

`func (o *SessionInfo) GetExpWarnOk() (*int64, bool)`

GetExpWarnOk returns a tuple with the ExpWarn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpWarn

`func (o *SessionInfo) SetExpWarn(v int64)`

SetExpWarn sets ExpWarn field to given value.


### GetSesTimeout

`func (o *SessionInfo) GetSesTimeout() int64`

GetSesTimeout returns the SesTimeout field if non-nil, zero value otherwise.

### GetSesTimeoutOk

`func (o *SessionInfo) GetSesTimeoutOk() (*int64, bool)`

GetSesTimeoutOk returns a tuple with the SesTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSesTimeout

`func (o *SessionInfo) SetSesTimeout(v int64)`

SetSesTimeout sets SesTimeout field to given value.


### GetSub

`func (o *SessionInfo) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *SessionInfo) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *SessionInfo) SetSub(v string)`

SetSub sets Sub field to given value.


### GetFlash

`func (o *SessionInfo) GetFlash() string`

GetFlash returns the Flash field if non-nil, zero value otherwise.

### GetFlashOk

`func (o *SessionInfo) GetFlashOk() (*string, bool)`

GetFlashOk returns a tuple with the Flash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlash

`func (o *SessionInfo) SetFlash(v string)`

SetFlash sets Flash field to given value.


### GetPollIntervalSeconds

`func (o *SessionInfo) GetPollIntervalSeconds() int32`

GetPollIntervalSeconds returns the PollIntervalSeconds field if non-nil, zero value otherwise.

### GetPollIntervalSecondsOk

`func (o *SessionInfo) GetPollIntervalSecondsOk() (*int32, bool)`

GetPollIntervalSecondsOk returns a tuple with the PollIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollIntervalSeconds

`func (o *SessionInfo) SetPollIntervalSeconds(v int32)`

SetPollIntervalSeconds sets PollIntervalSeconds field to given value.


### GetRoles

`func (o *SessionInfo) GetRoles() []Role`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *SessionInfo) GetRolesOk() (*[]Role, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *SessionInfo) SetRoles(v []Role)`

SetRoles sets Roles field to given value.


### GetAccessControlDirectives

`func (o *SessionInfo) GetAccessControlDirectives() []AdminAccessControlDirective`

GetAccessControlDirectives returns the AccessControlDirectives field if non-nil, zero value otherwise.

### GetAccessControlDirectivesOk

`func (o *SessionInfo) GetAccessControlDirectivesOk() (*[]AdminAccessControlDirective, bool)`

GetAccessControlDirectivesOk returns a tuple with the AccessControlDirectives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControlDirectives

`func (o *SessionInfo) SetAccessControlDirectives(v []AdminAccessControlDirective)`

SetAccessControlDirectives sets AccessControlDirectives field to given value.


### GetUseSlo

`func (o *SessionInfo) GetUseSlo() bool`

GetUseSlo returns the UseSlo field if non-nil, zero value otherwise.

### GetUseSloOk

`func (o *SessionInfo) GetUseSloOk() (*bool, bool)`

GetUseSloOk returns a tuple with the UseSlo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSlo

`func (o *SessionInfo) SetUseSlo(v bool)`

SetUseSlo sets UseSlo field to given value.


### GetFipsMode

`func (o *SessionInfo) GetFipsMode() bool`

GetFipsMode returns the FipsMode field if non-nil, zero value otherwise.

### GetFipsModeOk

`func (o *SessionInfo) GetFipsModeOk() (*bool, bool)`

GetFipsModeOk returns a tuple with the FipsMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFipsMode

`func (o *SessionInfo) SetFipsMode(v bool)`

SetFipsMode sets FipsMode field to given value.


### GetDisplayLogSettings

`func (o *SessionInfo) GetDisplayLogSettings() bool`

GetDisplayLogSettings returns the DisplayLogSettings field if non-nil, zero value otherwise.

### GetDisplayLogSettingsOk

`func (o *SessionInfo) GetDisplayLogSettingsOk() (*bool, bool)`

GetDisplayLogSettingsOk returns a tuple with the DisplayLogSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLogSettings

`func (o *SessionInfo) SetDisplayLogSettings(v bool)`

SetDisplayLogSettings sets DisplayLogSettings field to given value.


### GetConfigurationExports

`func (o *SessionInfo) GetConfigurationExports() ConfigStatuses`

GetConfigurationExports returns the ConfigurationExports field if non-nil, zero value otherwise.

### GetConfigurationExportsOk

`func (o *SessionInfo) GetConfigurationExportsOk() (*ConfigStatuses, bool)`

GetConfigurationExportsOk returns a tuple with the ConfigurationExports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationExports

`func (o *SessionInfo) SetConfigurationExports(v ConfigStatuses)`

SetConfigurationExports sets ConfigurationExports field to given value.


### GetConfigurationImports

`func (o *SessionInfo) GetConfigurationImports() ConfigStatuses`

GetConfigurationImports returns the ConfigurationImports field if non-nil, zero value otherwise.

### GetConfigurationImportsOk

`func (o *SessionInfo) GetConfigurationImportsOk() (*ConfigStatuses, bool)`

GetConfigurationImportsOk returns a tuple with the ConfigurationImports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationImports

`func (o *SessionInfo) SetConfigurationImports(v ConfigStatuses)`

SetConfigurationImports sets ConfigurationImports field to given value.


### GetSniEnabled

`func (o *SessionInfo) GetSniEnabled() bool`

GetSniEnabled returns the SniEnabled field if non-nil, zero value otherwise.

### GetSniEnabledOk

`func (o *SessionInfo) GetSniEnabledOk() (*bool, bool)`

GetSniEnabledOk returns a tuple with the SniEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSniEnabled

`func (o *SessionInfo) SetSniEnabled(v bool)`

SetSniEnabled sets SniEnabled field to given value.


### GetMaxFileUploadSize

`func (o *SessionInfo) GetMaxFileUploadSize() int32`

GetMaxFileUploadSize returns the MaxFileUploadSize field if non-nil, zero value otherwise.

### GetMaxFileUploadSizeOk

`func (o *SessionInfo) GetMaxFileUploadSizeOk() (*int32, bool)`

GetMaxFileUploadSizeOk returns a tuple with the MaxFileUploadSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFileUploadSize

`func (o *SessionInfo) SetMaxFileUploadSize(v int32)`

SetMaxFileUploadSize sets MaxFileUploadSize field to given value.


### GetShowWarning

`func (o *SessionInfo) GetShowWarning() bool`

GetShowWarning returns the ShowWarning field if non-nil, zero value otherwise.

### GetShowWarningOk

`func (o *SessionInfo) GetShowWarningOk() (*bool, bool)`

GetShowWarningOk returns a tuple with the ShowWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowWarning

`func (o *SessionInfo) SetShowWarning(v bool)`

SetShowWarning sets ShowWarning field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


