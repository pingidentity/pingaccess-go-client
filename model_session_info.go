/*
PingAccess Administrative API

The PingAccess Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingAccess.

API version: 7.1.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// SessionInfo A session.
type SessionInfo struct {
	// Time at which session will expire due to inactivity.
	Exp int64 `json:"exp"`
	// Time at which session was created.
	Iat int64 `json:"iat"`
	// Length of time before a timeout at which warning should appear.
	ExpWarn int64 `json:"expWarn"`
	// Maximum length of time that a session is allowed to live, regardless of user activity, -1 indicates disabled.
	SesTimeout int64 `json:"sesTimeout"`
	// The Session's subject.
	Sub string `json:"sub"`
	// The Warning message.
	Flash string `json:"flash"`
	// Session poll interval configuration in seconds.
	PollIntervalSeconds int32 `json:"pollIntervalSeconds"`
	// The user's roles.
	Roles []Role `json:"roles"`
	// The set of access control directives.
	AccessControlDirectives []AdminAccessControlDirective `json:"accessControlDirectives"`
	// Indicates whether single log out (SLO) is enabled or not.
	UseSlo bool `json:"useSlo"`
	// Indicates whether FIPS mode is enabled or not.
	FipsMode bool `json:"fipsMode"`
	ConfigurationExports ConfigStatuses `json:"configurationExports"`
	ConfigurationImports ConfigStatuses `json:"configurationImports"`
	// Indicates that SNI is enabled or not.
	SniEnabled bool `json:"sniEnabled"`
	// The maximum file upload size in bytes.
	MaxFileUploadSize int32 `json:"maxFileUploadSize"`
	// Indicates that a warning needs to be shown or not.
	ShowWarning bool `json:"showWarning"`
}

// NewSessionInfo instantiates a new SessionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionInfo(exp int64, iat int64, expWarn int64, sesTimeout int64, sub string, flash string, pollIntervalSeconds int32, roles []Role, accessControlDirectives []AdminAccessControlDirective, useSlo bool, fipsMode bool, configurationExports ConfigStatuses, configurationImports ConfigStatuses, sniEnabled bool, maxFileUploadSize int32, showWarning bool) *SessionInfo {
	this := SessionInfo{}
	this.Exp = exp
	this.Iat = iat
	this.ExpWarn = expWarn
	this.SesTimeout = sesTimeout
	this.Sub = sub
	this.Flash = flash
	this.PollIntervalSeconds = pollIntervalSeconds
	this.Roles = roles
	this.AccessControlDirectives = accessControlDirectives
	this.UseSlo = useSlo
	this.FipsMode = fipsMode
	this.ConfigurationExports = configurationExports
	this.ConfigurationImports = configurationImports
	this.SniEnabled = sniEnabled
	this.MaxFileUploadSize = maxFileUploadSize
	this.ShowWarning = showWarning
	return &this
}

// NewSessionInfoWithDefaults instantiates a new SessionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionInfoWithDefaults() *SessionInfo {
	this := SessionInfo{}
	return &this
}

// GetExp returns the Exp field value
func (o *SessionInfo) GetExp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Exp
}

// GetExpOk returns a tuple with the Exp field value
// and a boolean to check if the value has been set.
func (o *SessionInfo) GetExpOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Exp, true
}

// SetExp sets field value
func (o *SessionInfo) SetExp(v int64) {
	o.Exp = v
}

// GetIat returns the Iat field value
func (o *SessionInfo) GetIat() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Iat
}

// GetIatOk returns a tuple with the Iat field value
// and a boolean to check if the value has been set.
func (o *SessionInfo) GetIatOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Iat, true
}

// SetIat sets field value
func (o *SessionInfo) SetIat(v int64) {
	o.Iat = v
}

// GetExpWarn returns the ExpWarn field value
func (o *SessionInfo) GetExpWarn() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ExpWarn
}

// GetExpWarnOk returns a tuple with the ExpWarn field value
// and a boolean to check if the value has been set.
func (o *SessionInfo) GetExpWarnOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ExpWarn, true
}

// SetExpWarn sets field value
func (o *SessionInfo) SetExpWarn(v int64) {
	o.ExpWarn = v
}

// GetSesTimeout returns the SesTimeout field value
func (o *SessionInfo) GetSesTimeout() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SesTimeout
}

// GetSesTimeoutOk returns a tuple with the SesTimeout field value
// and a boolean to check if the value has been set.
func (o *SessionInfo) GetSesTimeoutOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SesTimeout, true
}

// SetSesTimeout sets field value
func (o *SessionInfo) SetSesTimeout(v int64) {
	o.SesTimeout = v
}

// GetSub returns the Sub field value
func (o *SessionInfo) GetSub() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sub
}

// GetSubOk returns a tuple with the Sub field value
// and a boolean to check if the value has been set.
func (o *SessionInfo) GetSubOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Sub, true
}

// SetSub sets field value
func (o *SessionInfo) SetSub(v string) {
	o.Sub = v
}

// GetFlash returns the Flash field value
func (o *SessionInfo) GetFlash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Flash
}

// GetFlashOk returns a tuple with the Flash field value
// and a boolean to check if the value has been set.
func (o *SessionInfo) GetFlashOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Flash, true
}

// SetFlash sets field value
func (o *SessionInfo) SetFlash(v string) {
	o.Flash = v
}

// GetPollIntervalSeconds returns the PollIntervalSeconds field value
func (o *SessionInfo) GetPollIntervalSeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PollIntervalSeconds
}

// GetPollIntervalSecondsOk returns a tuple with the PollIntervalSeconds field value
// and a boolean to check if the value has been set.
func (o *SessionInfo) GetPollIntervalSecondsOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PollIntervalSeconds, true
}

// SetPollIntervalSeconds sets field value
func (o *SessionInfo) SetPollIntervalSeconds(v int32) {
	o.PollIntervalSeconds = v
}

// GetRoles returns the Roles field value
func (o *SessionInfo) GetRoles() []Role {
	if o == nil {
		var ret []Role
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *SessionInfo) GetRolesOk() ([]Role, bool) {
	if o == nil {
    return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *SessionInfo) SetRoles(v []Role) {
	o.Roles = v
}

// GetAccessControlDirectives returns the AccessControlDirectives field value
func (o *SessionInfo) GetAccessControlDirectives() []AdminAccessControlDirective {
	if o == nil {
		var ret []AdminAccessControlDirective
		return ret
	}

	return o.AccessControlDirectives
}

// GetAccessControlDirectivesOk returns a tuple with the AccessControlDirectives field value
// and a boolean to check if the value has been set.
func (o *SessionInfo) GetAccessControlDirectivesOk() ([]AdminAccessControlDirective, bool) {
	if o == nil {
    return nil, false
	}
	return o.AccessControlDirectives, true
}

// SetAccessControlDirectives sets field value
func (o *SessionInfo) SetAccessControlDirectives(v []AdminAccessControlDirective) {
	o.AccessControlDirectives = v
}

// GetUseSlo returns the UseSlo field value
func (o *SessionInfo) GetUseSlo() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseSlo
}

// GetUseSloOk returns a tuple with the UseSlo field value
// and a boolean to check if the value has been set.
func (o *SessionInfo) GetUseSloOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UseSlo, true
}

// SetUseSlo sets field value
func (o *SessionInfo) SetUseSlo(v bool) {
	o.UseSlo = v
}

// GetFipsMode returns the FipsMode field value
func (o *SessionInfo) GetFipsMode() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FipsMode
}

// GetFipsModeOk returns a tuple with the FipsMode field value
// and a boolean to check if the value has been set.
func (o *SessionInfo) GetFipsModeOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.FipsMode, true
}

// SetFipsMode sets field value
func (o *SessionInfo) SetFipsMode(v bool) {
	o.FipsMode = v
}

// GetConfigurationExports returns the ConfigurationExports field value
func (o *SessionInfo) GetConfigurationExports() ConfigStatuses {
	if o == nil {
		var ret ConfigStatuses
		return ret
	}

	return o.ConfigurationExports
}

// GetConfigurationExportsOk returns a tuple with the ConfigurationExports field value
// and a boolean to check if the value has been set.
func (o *SessionInfo) GetConfigurationExportsOk() (*ConfigStatuses, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ConfigurationExports, true
}

// SetConfigurationExports sets field value
func (o *SessionInfo) SetConfigurationExports(v ConfigStatuses) {
	o.ConfigurationExports = v
}

// GetConfigurationImports returns the ConfigurationImports field value
func (o *SessionInfo) GetConfigurationImports() ConfigStatuses {
	if o == nil {
		var ret ConfigStatuses
		return ret
	}

	return o.ConfigurationImports
}

// GetConfigurationImportsOk returns a tuple with the ConfigurationImports field value
// and a boolean to check if the value has been set.
func (o *SessionInfo) GetConfigurationImportsOk() (*ConfigStatuses, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ConfigurationImports, true
}

// SetConfigurationImports sets field value
func (o *SessionInfo) SetConfigurationImports(v ConfigStatuses) {
	o.ConfigurationImports = v
}

// GetSniEnabled returns the SniEnabled field value
func (o *SessionInfo) GetSniEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SniEnabled
}

// GetSniEnabledOk returns a tuple with the SniEnabled field value
// and a boolean to check if the value has been set.
func (o *SessionInfo) GetSniEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SniEnabled, true
}

// SetSniEnabled sets field value
func (o *SessionInfo) SetSniEnabled(v bool) {
	o.SniEnabled = v
}

// GetMaxFileUploadSize returns the MaxFileUploadSize field value
func (o *SessionInfo) GetMaxFileUploadSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxFileUploadSize
}

// GetMaxFileUploadSizeOk returns a tuple with the MaxFileUploadSize field value
// and a boolean to check if the value has been set.
func (o *SessionInfo) GetMaxFileUploadSizeOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MaxFileUploadSize, true
}

// SetMaxFileUploadSize sets field value
func (o *SessionInfo) SetMaxFileUploadSize(v int32) {
	o.MaxFileUploadSize = v
}

// GetShowWarning returns the ShowWarning field value
func (o *SessionInfo) GetShowWarning() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ShowWarning
}

// GetShowWarningOk returns a tuple with the ShowWarning field value
// and a boolean to check if the value has been set.
func (o *SessionInfo) GetShowWarningOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ShowWarning, true
}

// SetShowWarning sets field value
func (o *SessionInfo) SetShowWarning(v bool) {
	o.ShowWarning = v
}

func (o SessionInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["exp"] = o.Exp
	}
	if true {
		toSerialize["iat"] = o.Iat
	}
	if true {
		toSerialize["expWarn"] = o.ExpWarn
	}
	if true {
		toSerialize["sesTimeout"] = o.SesTimeout
	}
	if true {
		toSerialize["sub"] = o.Sub
	}
	if true {
		toSerialize["flash"] = o.Flash
	}
	if true {
		toSerialize["pollIntervalSeconds"] = o.PollIntervalSeconds
	}
	if true {
		toSerialize["roles"] = o.Roles
	}
	if true {
		toSerialize["accessControlDirectives"] = o.AccessControlDirectives
	}
	if true {
		toSerialize["useSlo"] = o.UseSlo
	}
	if true {
		toSerialize["fipsMode"] = o.FipsMode
	}
	if true {
		toSerialize["configurationExports"] = o.ConfigurationExports
	}
	if true {
		toSerialize["configurationImports"] = o.ConfigurationImports
	}
	if true {
		toSerialize["sniEnabled"] = o.SniEnabled
	}
	if true {
		toSerialize["maxFileUploadSize"] = o.MaxFileUploadSize
	}
	if true {
		toSerialize["showWarning"] = o.ShowWarning
	}
	return json.Marshal(toSerialize)
}

type NullableSessionInfo struct {
	value *SessionInfo
	isSet bool
}

func (v NullableSessionInfo) Get() *SessionInfo {
	return v.value
}

func (v *NullableSessionInfo) Set(val *SessionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionInfo(val *SessionInfo) *NullableSessionInfo {
	return &NullableSessionInfo{value: val, isSet: true}
}

func (v NullableSessionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


