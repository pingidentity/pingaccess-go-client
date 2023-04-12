# Resource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | When creating a new Resource, this is the ID for the Resource. If not specified, an ID will be automatically assigned. When updating an existing Resource, this field is ignored and the ID is determined by the path parameter. | [optional] 
**Name** | **string** | (sortable) The name of the resource. | 
**Methods** | **[]string** | An array of HTTP methods configured for the resource. | 
**PathPatterns** | Pointer to [**[]PathPattern**](PathPattern.md) | A list of one or more request path-matching patterns. | [optional] 
**PathPrefixes** | Pointer to **[]string** | An array of path prefixes for the resource (DEPRECATED - to be removed in a future release; please use &#39;pathPatterns&#39; instead). | [optional] 
**AuditLevel** | Pointer to [**AuditLevel**](AuditLevel.md) |  | [optional] 
**RootResource** | Pointer to **bool** | (sortable) True if the resource is the root resource for the application. | [optional] 
**Anonymous** | Pointer to **bool** | (sortable) True if the resource is anonymous. | [optional] 
**Enabled** | Pointer to **bool** | (sortable) True if the resource is enabled. | [optional] 
**Unprotected** | Pointer to **bool** | (sortable) True if the resource is unprotected. | [optional] 
**Policy** | Pointer to [**map[string][]PolicyItem**](array.md) | A map of policy items associated with the resource.  The key is &#39;Web&#39; or &#39;API&#39; and the value is a list of PolicyItems.  Key type: String Value type: PolicyItem[] | [optional] 
**DefaultAuthTypeOverride** | [**DefaultAuthType**](DefaultAuthType.md) |  | 
**ApplicationId** | Pointer to **int32** | The id of the associated application. This field is read-only. | [optional] 
**QueryParamConfig** | Pointer to [**QueryParamConfig**](QueryParamConfig.md) |  | [optional] 
**ResourceType** | Pointer to [**ResourceType**](ResourceType.md) |  | [optional] 
**ResourceTypeConfiguration** | Pointer to [**ResourceTypeConfiguration**](ResourceTypeConfiguration.md) |  | [optional] 
**AuthenticationChallengePolicyId** | **string** | The UUID of the authentication challenge policy associated with the resource. This policy takes precedence over an application-level policy. | 
**RiskPolicyId** | Pointer to **int32** | The ID of the risk policy to use for this resource. This risk policy takes precedence over an application-level risk policy. | [optional] 

## Methods

### NewResource

`func NewResource(name string, methods []string, defaultAuthTypeOverride DefaultAuthType, authenticationChallengePolicyId string, ) *Resource`

NewResource instantiates a new Resource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceWithDefaults

`func NewResourceWithDefaults() *Resource`

NewResourceWithDefaults instantiates a new Resource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Resource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Resource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Resource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Resource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Resource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Resource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Resource) SetName(v string)`

SetName sets Name field to given value.


### GetMethods

`func (o *Resource) GetMethods() []string`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *Resource) GetMethodsOk() (*[]string, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *Resource) SetMethods(v []string)`

SetMethods sets Methods field to given value.


### GetPathPatterns

`func (o *Resource) GetPathPatterns() []PathPattern`

GetPathPatterns returns the PathPatterns field if non-nil, zero value otherwise.

### GetPathPatternsOk

`func (o *Resource) GetPathPatternsOk() (*[]PathPattern, bool)`

GetPathPatternsOk returns a tuple with the PathPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathPatterns

`func (o *Resource) SetPathPatterns(v []PathPattern)`

SetPathPatterns sets PathPatterns field to given value.

### HasPathPatterns

`func (o *Resource) HasPathPatterns() bool`

HasPathPatterns returns a boolean if a field has been set.

### GetPathPrefixes

`func (o *Resource) GetPathPrefixes() []string`

GetPathPrefixes returns the PathPrefixes field if non-nil, zero value otherwise.

### GetPathPrefixesOk

`func (o *Resource) GetPathPrefixesOk() (*[]string, bool)`

GetPathPrefixesOk returns a tuple with the PathPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathPrefixes

`func (o *Resource) SetPathPrefixes(v []string)`

SetPathPrefixes sets PathPrefixes field to given value.

### HasPathPrefixes

`func (o *Resource) HasPathPrefixes() bool`

HasPathPrefixes returns a boolean if a field has been set.

### GetAuditLevel

`func (o *Resource) GetAuditLevel() AuditLevel`

GetAuditLevel returns the AuditLevel field if non-nil, zero value otherwise.

### GetAuditLevelOk

`func (o *Resource) GetAuditLevelOk() (*AuditLevel, bool)`

GetAuditLevelOk returns a tuple with the AuditLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLevel

`func (o *Resource) SetAuditLevel(v AuditLevel)`

SetAuditLevel sets AuditLevel field to given value.

### HasAuditLevel

`func (o *Resource) HasAuditLevel() bool`

HasAuditLevel returns a boolean if a field has been set.

### GetRootResource

`func (o *Resource) GetRootResource() bool`

GetRootResource returns the RootResource field if non-nil, zero value otherwise.

### GetRootResourceOk

`func (o *Resource) GetRootResourceOk() (*bool, bool)`

GetRootResourceOk returns a tuple with the RootResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootResource

`func (o *Resource) SetRootResource(v bool)`

SetRootResource sets RootResource field to given value.

### HasRootResource

`func (o *Resource) HasRootResource() bool`

HasRootResource returns a boolean if a field has been set.

### GetAnonymous

`func (o *Resource) GetAnonymous() bool`

GetAnonymous returns the Anonymous field if non-nil, zero value otherwise.

### GetAnonymousOk

`func (o *Resource) GetAnonymousOk() (*bool, bool)`

GetAnonymousOk returns a tuple with the Anonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymous

`func (o *Resource) SetAnonymous(v bool)`

SetAnonymous sets Anonymous field to given value.

### HasAnonymous

`func (o *Resource) HasAnonymous() bool`

HasAnonymous returns a boolean if a field has been set.

### GetEnabled

`func (o *Resource) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Resource) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Resource) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Resource) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUnprotected

`func (o *Resource) GetUnprotected() bool`

GetUnprotected returns the Unprotected field if non-nil, zero value otherwise.

### GetUnprotectedOk

`func (o *Resource) GetUnprotectedOk() (*bool, bool)`

GetUnprotectedOk returns a tuple with the Unprotected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnprotected

`func (o *Resource) SetUnprotected(v bool)`

SetUnprotected sets Unprotected field to given value.

### HasUnprotected

`func (o *Resource) HasUnprotected() bool`

HasUnprotected returns a boolean if a field has been set.

### GetPolicy

`func (o *Resource) GetPolicy() map[string][]PolicyItem`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *Resource) GetPolicyOk() (*map[string][]PolicyItem, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *Resource) SetPolicy(v map[string][]PolicyItem)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *Resource) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetDefaultAuthTypeOverride

`func (o *Resource) GetDefaultAuthTypeOverride() DefaultAuthType`

GetDefaultAuthTypeOverride returns the DefaultAuthTypeOverride field if non-nil, zero value otherwise.

### GetDefaultAuthTypeOverrideOk

`func (o *Resource) GetDefaultAuthTypeOverrideOk() (*DefaultAuthType, bool)`

GetDefaultAuthTypeOverrideOk returns a tuple with the DefaultAuthTypeOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAuthTypeOverride

`func (o *Resource) SetDefaultAuthTypeOverride(v DefaultAuthType)`

SetDefaultAuthTypeOverride sets DefaultAuthTypeOverride field to given value.


### GetApplicationId

`func (o *Resource) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *Resource) GetApplicationIdOk() (*int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *Resource) SetApplicationId(v int32)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *Resource) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetQueryParamConfig

`func (o *Resource) GetQueryParamConfig() QueryParamConfig`

GetQueryParamConfig returns the QueryParamConfig field if non-nil, zero value otherwise.

### GetQueryParamConfigOk

`func (o *Resource) GetQueryParamConfigOk() (*QueryParamConfig, bool)`

GetQueryParamConfigOk returns a tuple with the QueryParamConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryParamConfig

`func (o *Resource) SetQueryParamConfig(v QueryParamConfig)`

SetQueryParamConfig sets QueryParamConfig field to given value.

### HasQueryParamConfig

`func (o *Resource) HasQueryParamConfig() bool`

HasQueryParamConfig returns a boolean if a field has been set.

### GetResourceType

`func (o *Resource) GetResourceType() ResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *Resource) GetResourceTypeOk() (*ResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *Resource) SetResourceType(v ResourceType)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *Resource) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetResourceTypeConfiguration

`func (o *Resource) GetResourceTypeConfiguration() ResourceTypeConfiguration`

GetResourceTypeConfiguration returns the ResourceTypeConfiguration field if non-nil, zero value otherwise.

### GetResourceTypeConfigurationOk

`func (o *Resource) GetResourceTypeConfigurationOk() (*ResourceTypeConfiguration, bool)`

GetResourceTypeConfigurationOk returns a tuple with the ResourceTypeConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTypeConfiguration

`func (o *Resource) SetResourceTypeConfiguration(v ResourceTypeConfiguration)`

SetResourceTypeConfiguration sets ResourceTypeConfiguration field to given value.

### HasResourceTypeConfiguration

`func (o *Resource) HasResourceTypeConfiguration() bool`

HasResourceTypeConfiguration returns a boolean if a field has been set.

### GetAuthenticationChallengePolicyId

`func (o *Resource) GetAuthenticationChallengePolicyId() string`

GetAuthenticationChallengePolicyId returns the AuthenticationChallengePolicyId field if non-nil, zero value otherwise.

### GetAuthenticationChallengePolicyIdOk

`func (o *Resource) GetAuthenticationChallengePolicyIdOk() (*string, bool)`

GetAuthenticationChallengePolicyIdOk returns a tuple with the AuthenticationChallengePolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationChallengePolicyId

`func (o *Resource) SetAuthenticationChallengePolicyId(v string)`

SetAuthenticationChallengePolicyId sets AuthenticationChallengePolicyId field to given value.


### GetRiskPolicyId

`func (o *Resource) GetRiskPolicyId() int32`

GetRiskPolicyId returns the RiskPolicyId field if non-nil, zero value otherwise.

### GetRiskPolicyIdOk

`func (o *Resource) GetRiskPolicyIdOk() (*int32, bool)`

GetRiskPolicyIdOk returns a tuple with the RiskPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskPolicyId

`func (o *Resource) SetRiskPolicyId(v int32)`

SetRiskPolicyId sets RiskPolicyId field to given value.

### HasRiskPolicyId

`func (o *Resource) HasRiskPolicyId() bool`

HasRiskPolicyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


