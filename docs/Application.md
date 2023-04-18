# Application

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | When creating a new Application, this is the ID for the Application. If not specified, an ID will be automatically assigned. When updating an existing Application, this field is ignored and the ID is determined by the path parameter. | [optional] 
**Enabled** | Pointer to **bool** | (sortable) True if the application is enabled. | [optional] 
**AgentCacheInvalidatedResponseDuration** | Pointer to **int64** | (DEPRECATED - to be removed in a future release; this field is no longer used when processing requests for an application.) | [optional] 
**AgentCacheInvalidatedExpiration** | Pointer to **int64** | (DEPRECATED - to be removed in a future release; this field is no longer used when processing requests for an application.) | [optional] 
**LastModified** | Pointer to **int64** | The last modified time of the configuration for this application, its resources and associated policy, as the number of milliseconds since January 1, 1970, 00:00:00 GMT. This field is read-only. | [optional] 
**RequireHTTPS** | Pointer to **bool** | (sortable) True if the application requires HTTPS connections. | [optional] 
**Name** | **string** | (sortable) The application name. | 
**Description** | Pointer to **string** | (sortable) A description of the application. | [optional] 
**ManualOrderingEnabled** | Pointer to **bool** | Enable explicit, manual ordering of application resources and permit regex path patterns. | [optional] 
**ResourceOrder** | Pointer to **[]int64** | The explicit resource order defined when manual ordering is enabled. Each existing resource ID must be represented. (Required when &#39;manualOrderingEnabled&#39; is true.) | [optional] 
**ApplicationType** | Pointer to [**ApplicationType**](ApplicationType.md) |  | [optional] 
**DefaultAuthType** | [**DefaultAuthType**](DefaultAuthType.md) |  | 
**SpaSupportEnabled** | **bool** | Enable SPA support. | 
**FallbackPostEncoding** | Pointer to **string** | (sortable) Specify the name of an encoding to use when preserving POST parameters and the parameters are found to not be UTF-8 encoded. | [optional] 
**Destination** | Pointer to [**Destination**](Destination.md) |  | [optional] 
**ContextRoot** | **string** | (sortable) The context root of the application. | 
**Realm** | Pointer to **string** | (sortable) The OAuth realm associated with the application. | [optional] 
**CaseSensitivePath** | Pointer to **bool** | (sortable) True if the path is case sensitive. | [optional] 
**AllowEmptyPathSegments** | Pointer to **bool** | When true, PingAccess will not remove empty path segments from the request URL before matching a request against the resources in this application. Defaults to false. | [optional] 
**WebSessionId** | Pointer to **int64** | The ID of the web session associated with the application or zero if none. | [optional] 
**Issuer** | Pointer to **string** | Branded URL at the OpenID Connect provider to redirect unauthenticated requests to. When specified, this overrides the global token provider&#39;s issuer field. | [optional] 
**SiteId** | **int64** | The ID of the site associated with the application or zero if none. | 
**AgentId** | **int64** | The ID of the agent associated with the application or zero if none. | 
**SidebandClientId** | **string** | The ID of the sideband client associated with the application or null if none. | 
**VirtualHostIds** | **[]int64** | An array of virtual host IDs associated with the application. | 
**IdentityMappingIds** | Pointer to **map[string]int64** | A map of Identity Mappings associated with the application. The key is &#39;Web&#39; or &#39;API&#39; and the value is an Identity Mapping ID.  Key type: String Value type: Integer | [optional] 
**AccessValidatorId** | Pointer to **int64** | The ID of the access token validator for local token validation, 1 if the application is protected remotely by an Authorization Server, or zero if unprotected. Only applies to applications of type API. | [optional] 
**Policy** | Pointer to [**map[string][]PolicyItem**](array.md) | A map of policy items associated with the application.  The key is &#39;Web&#39; or &#39;API&#39; and the value is a list of PolicyItems.  Key type: String Value type: PolicyItem[] | [optional] 
**AuthenticationChallengePolicyId** | **string** | The UUID of the authentication challenge policy associated with the application. | 
**RiskPolicyId** | Pointer to **int64** | The ID of the risk policy to use by default for this application. | [optional] 
**DeployReservedResources** | Pointer to **bool** | When true, PingAccess will deploy an  instance of the reserved PA resources and runtime API endpoints using this application context root plus the  reserved application context root as the base path. Default value: false. | [optional] 

## Methods

### NewApplication

`func NewApplication(name string, defaultAuthType DefaultAuthType, spaSupportEnabled bool, contextRoot string, siteId int64, agentId int64, sidebandClientId string, virtualHostIds []int64, authenticationChallengePolicyId string, ) *Application`

NewApplication instantiates a new Application object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationWithDefaults

`func NewApplicationWithDefaults() *Application`

NewApplicationWithDefaults instantiates a new Application object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Application) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Application) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Application) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Application) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnabled

`func (o *Application) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Application) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Application) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Application) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAgentCacheInvalidatedResponseDuration

`func (o *Application) GetAgentCacheInvalidatedResponseDuration() int64`

GetAgentCacheInvalidatedResponseDuration returns the AgentCacheInvalidatedResponseDuration field if non-nil, zero value otherwise.

### GetAgentCacheInvalidatedResponseDurationOk

`func (o *Application) GetAgentCacheInvalidatedResponseDurationOk() (*int64, bool)`

GetAgentCacheInvalidatedResponseDurationOk returns a tuple with the AgentCacheInvalidatedResponseDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentCacheInvalidatedResponseDuration

`func (o *Application) SetAgentCacheInvalidatedResponseDuration(v int64)`

SetAgentCacheInvalidatedResponseDuration sets AgentCacheInvalidatedResponseDuration field to given value.

### HasAgentCacheInvalidatedResponseDuration

`func (o *Application) HasAgentCacheInvalidatedResponseDuration() bool`

HasAgentCacheInvalidatedResponseDuration returns a boolean if a field has been set.

### GetAgentCacheInvalidatedExpiration

`func (o *Application) GetAgentCacheInvalidatedExpiration() int64`

GetAgentCacheInvalidatedExpiration returns the AgentCacheInvalidatedExpiration field if non-nil, zero value otherwise.

### GetAgentCacheInvalidatedExpirationOk

`func (o *Application) GetAgentCacheInvalidatedExpirationOk() (*int64, bool)`

GetAgentCacheInvalidatedExpirationOk returns a tuple with the AgentCacheInvalidatedExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentCacheInvalidatedExpiration

`func (o *Application) SetAgentCacheInvalidatedExpiration(v int64)`

SetAgentCacheInvalidatedExpiration sets AgentCacheInvalidatedExpiration field to given value.

### HasAgentCacheInvalidatedExpiration

`func (o *Application) HasAgentCacheInvalidatedExpiration() bool`

HasAgentCacheInvalidatedExpiration returns a boolean if a field has been set.

### GetLastModified

`func (o *Application) GetLastModified() int64`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *Application) GetLastModifiedOk() (*int64, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *Application) SetLastModified(v int64)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *Application) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetRequireHTTPS

`func (o *Application) GetRequireHTTPS() bool`

GetRequireHTTPS returns the RequireHTTPS field if non-nil, zero value otherwise.

### GetRequireHTTPSOk

`func (o *Application) GetRequireHTTPSOk() (*bool, bool)`

GetRequireHTTPSOk returns a tuple with the RequireHTTPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireHTTPS

`func (o *Application) SetRequireHTTPS(v bool)`

SetRequireHTTPS sets RequireHTTPS field to given value.

### HasRequireHTTPS

`func (o *Application) HasRequireHTTPS() bool`

HasRequireHTTPS returns a boolean if a field has been set.

### GetName

`func (o *Application) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Application) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Application) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Application) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Application) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Application) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Application) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetManualOrderingEnabled

`func (o *Application) GetManualOrderingEnabled() bool`

GetManualOrderingEnabled returns the ManualOrderingEnabled field if non-nil, zero value otherwise.

### GetManualOrderingEnabledOk

`func (o *Application) GetManualOrderingEnabledOk() (*bool, bool)`

GetManualOrderingEnabledOk returns a tuple with the ManualOrderingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualOrderingEnabled

`func (o *Application) SetManualOrderingEnabled(v bool)`

SetManualOrderingEnabled sets ManualOrderingEnabled field to given value.

### HasManualOrderingEnabled

`func (o *Application) HasManualOrderingEnabled() bool`

HasManualOrderingEnabled returns a boolean if a field has been set.

### GetResourceOrder

`func (o *Application) GetResourceOrder() []int64`

GetResourceOrder returns the ResourceOrder field if non-nil, zero value otherwise.

### GetResourceOrderOk

`func (o *Application) GetResourceOrderOk() (*[]int64, bool)`

GetResourceOrderOk returns a tuple with the ResourceOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOrder

`func (o *Application) SetResourceOrder(v []int64)`

SetResourceOrder sets ResourceOrder field to given value.

### HasResourceOrder

`func (o *Application) HasResourceOrder() bool`

HasResourceOrder returns a boolean if a field has been set.

### GetApplicationType

`func (o *Application) GetApplicationType() ApplicationType`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *Application) GetApplicationTypeOk() (*ApplicationType, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationType

`func (o *Application) SetApplicationType(v ApplicationType)`

SetApplicationType sets ApplicationType field to given value.

### HasApplicationType

`func (o *Application) HasApplicationType() bool`

HasApplicationType returns a boolean if a field has been set.

### GetDefaultAuthType

`func (o *Application) GetDefaultAuthType() DefaultAuthType`

GetDefaultAuthType returns the DefaultAuthType field if non-nil, zero value otherwise.

### GetDefaultAuthTypeOk

`func (o *Application) GetDefaultAuthTypeOk() (*DefaultAuthType, bool)`

GetDefaultAuthTypeOk returns a tuple with the DefaultAuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAuthType

`func (o *Application) SetDefaultAuthType(v DefaultAuthType)`

SetDefaultAuthType sets DefaultAuthType field to given value.


### GetSpaSupportEnabled

`func (o *Application) GetSpaSupportEnabled() bool`

GetSpaSupportEnabled returns the SpaSupportEnabled field if non-nil, zero value otherwise.

### GetSpaSupportEnabledOk

`func (o *Application) GetSpaSupportEnabledOk() (*bool, bool)`

GetSpaSupportEnabledOk returns a tuple with the SpaSupportEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaSupportEnabled

`func (o *Application) SetSpaSupportEnabled(v bool)`

SetSpaSupportEnabled sets SpaSupportEnabled field to given value.


### GetFallbackPostEncoding

`func (o *Application) GetFallbackPostEncoding() string`

GetFallbackPostEncoding returns the FallbackPostEncoding field if non-nil, zero value otherwise.

### GetFallbackPostEncodingOk

`func (o *Application) GetFallbackPostEncodingOk() (*string, bool)`

GetFallbackPostEncodingOk returns a tuple with the FallbackPostEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackPostEncoding

`func (o *Application) SetFallbackPostEncoding(v string)`

SetFallbackPostEncoding sets FallbackPostEncoding field to given value.

### HasFallbackPostEncoding

`func (o *Application) HasFallbackPostEncoding() bool`

HasFallbackPostEncoding returns a boolean if a field has been set.

### GetDestination

`func (o *Application) GetDestination() Destination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *Application) GetDestinationOk() (*Destination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *Application) SetDestination(v Destination)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *Application) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetContextRoot

`func (o *Application) GetContextRoot() string`

GetContextRoot returns the ContextRoot field if non-nil, zero value otherwise.

### GetContextRootOk

`func (o *Application) GetContextRootOk() (*string, bool)`

GetContextRootOk returns a tuple with the ContextRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextRoot

`func (o *Application) SetContextRoot(v string)`

SetContextRoot sets ContextRoot field to given value.


### GetRealm

`func (o *Application) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *Application) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *Application) SetRealm(v string)`

SetRealm sets Realm field to given value.

### HasRealm

`func (o *Application) HasRealm() bool`

HasRealm returns a boolean if a field has been set.

### GetCaseSensitivePath

`func (o *Application) GetCaseSensitivePath() bool`

GetCaseSensitivePath returns the CaseSensitivePath field if non-nil, zero value otherwise.

### GetCaseSensitivePathOk

`func (o *Application) GetCaseSensitivePathOk() (*bool, bool)`

GetCaseSensitivePathOk returns a tuple with the CaseSensitivePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitivePath

`func (o *Application) SetCaseSensitivePath(v bool)`

SetCaseSensitivePath sets CaseSensitivePath field to given value.

### HasCaseSensitivePath

`func (o *Application) HasCaseSensitivePath() bool`

HasCaseSensitivePath returns a boolean if a field has been set.

### GetAllowEmptyPathSegments

`func (o *Application) GetAllowEmptyPathSegments() bool`

GetAllowEmptyPathSegments returns the AllowEmptyPathSegments field if non-nil, zero value otherwise.

### GetAllowEmptyPathSegmentsOk

`func (o *Application) GetAllowEmptyPathSegmentsOk() (*bool, bool)`

GetAllowEmptyPathSegmentsOk returns a tuple with the AllowEmptyPathSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowEmptyPathSegments

`func (o *Application) SetAllowEmptyPathSegments(v bool)`

SetAllowEmptyPathSegments sets AllowEmptyPathSegments field to given value.

### HasAllowEmptyPathSegments

`func (o *Application) HasAllowEmptyPathSegments() bool`

HasAllowEmptyPathSegments returns a boolean if a field has been set.

### GetWebSessionId

`func (o *Application) GetWebSessionId() int64`

GetWebSessionId returns the WebSessionId field if non-nil, zero value otherwise.

### GetWebSessionIdOk

`func (o *Application) GetWebSessionIdOk() (*int64, bool)`

GetWebSessionIdOk returns a tuple with the WebSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebSessionId

`func (o *Application) SetWebSessionId(v int64)`

SetWebSessionId sets WebSessionId field to given value.

### HasWebSessionId

`func (o *Application) HasWebSessionId() bool`

HasWebSessionId returns a boolean if a field has been set.

### GetIssuer

`func (o *Application) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *Application) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *Application) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *Application) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetSiteId

`func (o *Application) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Application) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *Application) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.


### GetAgentId

`func (o *Application) GetAgentId() int64`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *Application) GetAgentIdOk() (*int64, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *Application) SetAgentId(v int64)`

SetAgentId sets AgentId field to given value.


### GetSidebandClientId

`func (o *Application) GetSidebandClientId() string`

GetSidebandClientId returns the SidebandClientId field if non-nil, zero value otherwise.

### GetSidebandClientIdOk

`func (o *Application) GetSidebandClientIdOk() (*string, bool)`

GetSidebandClientIdOk returns a tuple with the SidebandClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSidebandClientId

`func (o *Application) SetSidebandClientId(v string)`

SetSidebandClientId sets SidebandClientId field to given value.


### GetVirtualHostIds

`func (o *Application) GetVirtualHostIds() []int64`

GetVirtualHostIds returns the VirtualHostIds field if non-nil, zero value otherwise.

### GetVirtualHostIdsOk

`func (o *Application) GetVirtualHostIdsOk() (*[]int64, bool)`

GetVirtualHostIdsOk returns a tuple with the VirtualHostIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualHostIds

`func (o *Application) SetVirtualHostIds(v []int64)`

SetVirtualHostIds sets VirtualHostIds field to given value.


### GetIdentityMappingIds

`func (o *Application) GetIdentityMappingIds() map[string]int64`

GetIdentityMappingIds returns the IdentityMappingIds field if non-nil, zero value otherwise.

### GetIdentityMappingIdsOk

`func (o *Application) GetIdentityMappingIdsOk() (*map[string]int64, bool)`

GetIdentityMappingIdsOk returns a tuple with the IdentityMappingIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMappingIds

`func (o *Application) SetIdentityMappingIds(v map[string]int64)`

SetIdentityMappingIds sets IdentityMappingIds field to given value.

### HasIdentityMappingIds

`func (o *Application) HasIdentityMappingIds() bool`

HasIdentityMappingIds returns a boolean if a field has been set.

### GetAccessValidatorId

`func (o *Application) GetAccessValidatorId() int64`

GetAccessValidatorId returns the AccessValidatorId field if non-nil, zero value otherwise.

### GetAccessValidatorIdOk

`func (o *Application) GetAccessValidatorIdOk() (*int64, bool)`

GetAccessValidatorIdOk returns a tuple with the AccessValidatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessValidatorId

`func (o *Application) SetAccessValidatorId(v int64)`

SetAccessValidatorId sets AccessValidatorId field to given value.

### HasAccessValidatorId

`func (o *Application) HasAccessValidatorId() bool`

HasAccessValidatorId returns a boolean if a field has been set.

### GetPolicy

`func (o *Application) GetPolicy() map[string][]PolicyItem`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *Application) GetPolicyOk() (*map[string][]PolicyItem, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *Application) SetPolicy(v map[string][]PolicyItem)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *Application) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetAuthenticationChallengePolicyId

`func (o *Application) GetAuthenticationChallengePolicyId() string`

GetAuthenticationChallengePolicyId returns the AuthenticationChallengePolicyId field if non-nil, zero value otherwise.

### GetAuthenticationChallengePolicyIdOk

`func (o *Application) GetAuthenticationChallengePolicyIdOk() (*string, bool)`

GetAuthenticationChallengePolicyIdOk returns a tuple with the AuthenticationChallengePolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationChallengePolicyId

`func (o *Application) SetAuthenticationChallengePolicyId(v string)`

SetAuthenticationChallengePolicyId sets AuthenticationChallengePolicyId field to given value.


### GetRiskPolicyId

`func (o *Application) GetRiskPolicyId() int64`

GetRiskPolicyId returns the RiskPolicyId field if non-nil, zero value otherwise.

### GetRiskPolicyIdOk

`func (o *Application) GetRiskPolicyIdOk() (*int64, bool)`

GetRiskPolicyIdOk returns a tuple with the RiskPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskPolicyId

`func (o *Application) SetRiskPolicyId(v int64)`

SetRiskPolicyId sets RiskPolicyId field to given value.

### HasRiskPolicyId

`func (o *Application) HasRiskPolicyId() bool`

HasRiskPolicyId returns a boolean if a field has been set.

### GetDeployReservedResources

`func (o *Application) GetDeployReservedResources() bool`

GetDeployReservedResources returns the DeployReservedResources field if non-nil, zero value otherwise.

### GetDeployReservedResourcesOk

`func (o *Application) GetDeployReservedResourcesOk() (*bool, bool)`

GetDeployReservedResourcesOk returns a tuple with the DeployReservedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployReservedResources

`func (o *Application) SetDeployReservedResources(v bool)`

SetDeployReservedResources sets DeployReservedResources field to given value.

### HasDeployReservedResources

`func (o *Application) HasDeployReservedResources() bool`

HasDeployReservedResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


