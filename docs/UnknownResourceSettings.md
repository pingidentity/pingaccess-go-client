# UnknownResourceSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorStatusCode** | **int32** | The HTTP error response status code | 
**ErrorTemplateFile** | **string** | The name of the velocity template file to use for generating the error response body | 
**ErrorContentType** | [**ContentType**](ContentType.md) |  | 
**AgentDefaultMode** | [**UnknownResourceMode**](UnknownResourceMode.md) |  | 
**AgentDefaultCacheTTL** | **int32** | The default agent resource cache TTL (in seconds) to be used for unknown resources when a request cannot be mapped to a known virtual host. | 
**AuditLevel** | Pointer to [**AuditLevel**](AuditLevel.md) |  | [optional] 

## Methods

### NewUnknownResourceSettings

`func NewUnknownResourceSettings(errorStatusCode int32, errorTemplateFile string, errorContentType ContentType, agentDefaultMode UnknownResourceMode, agentDefaultCacheTTL int32, ) *UnknownResourceSettings`

NewUnknownResourceSettings instantiates a new UnknownResourceSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnknownResourceSettingsWithDefaults

`func NewUnknownResourceSettingsWithDefaults() *UnknownResourceSettings`

NewUnknownResourceSettingsWithDefaults instantiates a new UnknownResourceSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorStatusCode

`func (o *UnknownResourceSettings) GetErrorStatusCode() int32`

GetErrorStatusCode returns the ErrorStatusCode field if non-nil, zero value otherwise.

### GetErrorStatusCodeOk

`func (o *UnknownResourceSettings) GetErrorStatusCodeOk() (*int32, bool)`

GetErrorStatusCodeOk returns a tuple with the ErrorStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorStatusCode

`func (o *UnknownResourceSettings) SetErrorStatusCode(v int32)`

SetErrorStatusCode sets ErrorStatusCode field to given value.


### GetErrorTemplateFile

`func (o *UnknownResourceSettings) GetErrorTemplateFile() string`

GetErrorTemplateFile returns the ErrorTemplateFile field if non-nil, zero value otherwise.

### GetErrorTemplateFileOk

`func (o *UnknownResourceSettings) GetErrorTemplateFileOk() (*string, bool)`

GetErrorTemplateFileOk returns a tuple with the ErrorTemplateFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorTemplateFile

`func (o *UnknownResourceSettings) SetErrorTemplateFile(v string)`

SetErrorTemplateFile sets ErrorTemplateFile field to given value.


### GetErrorContentType

`func (o *UnknownResourceSettings) GetErrorContentType() ContentType`

GetErrorContentType returns the ErrorContentType field if non-nil, zero value otherwise.

### GetErrorContentTypeOk

`func (o *UnknownResourceSettings) GetErrorContentTypeOk() (*ContentType, bool)`

GetErrorContentTypeOk returns a tuple with the ErrorContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorContentType

`func (o *UnknownResourceSettings) SetErrorContentType(v ContentType)`

SetErrorContentType sets ErrorContentType field to given value.


### GetAgentDefaultMode

`func (o *UnknownResourceSettings) GetAgentDefaultMode() UnknownResourceMode`

GetAgentDefaultMode returns the AgentDefaultMode field if non-nil, zero value otherwise.

### GetAgentDefaultModeOk

`func (o *UnknownResourceSettings) GetAgentDefaultModeOk() (*UnknownResourceMode, bool)`

GetAgentDefaultModeOk returns a tuple with the AgentDefaultMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentDefaultMode

`func (o *UnknownResourceSettings) SetAgentDefaultMode(v UnknownResourceMode)`

SetAgentDefaultMode sets AgentDefaultMode field to given value.


### GetAgentDefaultCacheTTL

`func (o *UnknownResourceSettings) GetAgentDefaultCacheTTL() int32`

GetAgentDefaultCacheTTL returns the AgentDefaultCacheTTL field if non-nil, zero value otherwise.

### GetAgentDefaultCacheTTLOk

`func (o *UnknownResourceSettings) GetAgentDefaultCacheTTLOk() (*int32, bool)`

GetAgentDefaultCacheTTLOk returns a tuple with the AgentDefaultCacheTTL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentDefaultCacheTTL

`func (o *UnknownResourceSettings) SetAgentDefaultCacheTTL(v int32)`

SetAgentDefaultCacheTTL sets AgentDefaultCacheTTL field to given value.


### GetAuditLevel

`func (o *UnknownResourceSettings) GetAuditLevel() AuditLevel`

GetAuditLevel returns the AuditLevel field if non-nil, zero value otherwise.

### GetAuditLevelOk

`func (o *UnknownResourceSettings) GetAuditLevelOk() (*AuditLevel, bool)`

GetAuditLevelOk returns a tuple with the AuditLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLevel

`func (o *UnknownResourceSettings) SetAuditLevel(v AuditLevel)`

SetAuditLevel sets AuditLevel field to given value.

### HasAuditLevel

`func (o *UnknownResourceSettings) HasAuditLevel() bool`

HasAuditLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


