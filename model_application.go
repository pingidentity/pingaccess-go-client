/*
Administrative API Documentation

The PingAccess Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingAccess as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API.

API version: 7.3.0.2-SNAPSHOT
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the Application type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Application{}

// Application An application.
type Application struct {
	// When creating a new Application, this is the ID for the Application. If not specified, an ID will be automatically assigned. When updating an existing Application, this field is ignored and the ID is determined by the path parameter.
	Id *int32 `json:"id,omitempty"`
	// (sortable) True if the application is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// (DEPRECATED - to be removed in a future release; this field is no longer used when processing requests for an application.)
	AgentCacheInvalidatedResponseDuration *int32 `json:"agentCacheInvalidatedResponseDuration,omitempty"`
	// (DEPRECATED - to be removed in a future release; this field is no longer used when processing requests for an application.)
	AgentCacheInvalidatedExpiration *int64 `json:"agentCacheInvalidatedExpiration,omitempty"`
	// The last modified time of the configuration for this application, its resources and associated policy, as the number of milliseconds since January 1, 1970, 00:00:00 GMT. This field is read-only.
	LastModified *int64 `json:"lastModified,omitempty"`
	// (sortable) True if the application requires HTTPS connections.
	RequireHTTPS *bool `json:"requireHTTPS,omitempty"`
	// (sortable) The application name.
	Name string `json:"name"`
	// (sortable) A description of the application.
	Description *string `json:"description,omitempty"`
	// Enable explicit, manual ordering of application resources and permit regex path patterns.
	ManualOrderingEnabled *bool `json:"manualOrderingEnabled,omitempty"`
	// The explicit resource order defined when manual ordering is enabled. Each existing resource ID must be represented. (Required when 'manualOrderingEnabled' is true.)
	ResourceOrder   []int32          `json:"resourceOrder,omitempty"`
	ApplicationType *ApplicationType `json:"applicationType,omitempty"`
	DefaultAuthType DefaultAuthType  `json:"defaultAuthType"`
	// Enable SPA support.
	SpaSupportEnabled bool `json:"spaSupportEnabled"`
	// (sortable) Specify the name of an encoding to use when preserving POST parameters and the parameters are found to not be UTF-8 encoded.
	FallbackPostEncoding *string      `json:"fallbackPostEncoding,omitempty"`
	Destination          *Destination `json:"destination,omitempty"`
	// (sortable) The context root of the application.
	ContextRoot string `json:"contextRoot"`
	// (sortable) The OAuth realm associated with the application.
	Realm *string `json:"realm,omitempty"`
	// (sortable) True if the path is case sensitive.
	CaseSensitivePath *bool `json:"caseSensitivePath,omitempty"`
	// When true, PingAccess will not remove empty path segments from the request URL before matching a request against the resources in this application. Defaults to false.
	AllowEmptyPathSegments *bool `json:"allowEmptyPathSegments,omitempty"`
	// The ID of the web session associated with the application or zero if none.
	WebSessionId *int32 `json:"webSessionId,omitempty"`
	// Branded URL at the OpenID Connect provider to redirect unauthenticated requests to. When specified, this overrides the global token provider's issuer field.
	Issuer *string `json:"issuer,omitempty"`
	// The ID of the site associated with the application or zero if none.
	SiteId int32 `json:"siteId"`
	// The ID of the agent associated with the application or zero if none.
	AgentId int32 `json:"agentId"`
	// The ID of the sideband client associated with the application or null if none.
	SidebandClientId string `json:"sidebandClientId"`
	// An array of virtual host IDs associated with the application.
	VirtualHostIds []int32 `json:"virtualHostIds"`
	// A map of Identity Mappings associated with the application. The key is 'Web' or 'API' and the value is an Identity Mapping ID.  Key type: String Value type: Integer
	IdentityMappingIds *map[string]int32 `json:"identityMappingIds,omitempty"`
	// The ID of the access token validator for local token validation, 1 if the application is protected remotely by an Authorization Server, or zero if unprotected. Only applies to applications of type API.
	AccessValidatorId *int32 `json:"accessValidatorId,omitempty"`
	// A map of policy items associated with the application.  The key is 'Web' or 'API' and the value is a list of PolicyItems.  Key type: String Value type: PolicyItem[]
	Policy *map[string][]PolicyItem `json:"policy,omitempty"`
	// The UUID of the authentication challenge policy associated with the application.
	AuthenticationChallengePolicyId string `json:"authenticationChallengePolicyId"`
	// The ID of the risk policy to use by default for this application.
	RiskPolicyId *int32 `json:"riskPolicyId,omitempty"`
	// When true, PingAccess will deploy an  instance of the reserved PA resources and runtime API endpoints using this application context root plus the  reserved application context root as the base path. Default value: false.
	DeployReservedResources *bool `json:"deployReservedResources,omitempty"`
}

// NewApplication instantiates a new Application object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplication(name string, defaultAuthType DefaultAuthType, spaSupportEnabled bool, contextRoot string, siteId int32, agentId int32, sidebandClientId string, virtualHostIds []int32, authenticationChallengePolicyId string) *Application {
	this := Application{}
	this.Name = name
	this.DefaultAuthType = defaultAuthType
	this.SpaSupportEnabled = spaSupportEnabled
	this.ContextRoot = contextRoot
	this.SiteId = siteId
	this.AgentId = agentId
	this.SidebandClientId = sidebandClientId
	this.VirtualHostIds = virtualHostIds
	this.AuthenticationChallengePolicyId = authenticationChallengePolicyId
	return &this
}

// NewApplicationWithDefaults instantiates a new Application object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationWithDefaults() *Application {
	this := Application{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Application) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Application) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Application) SetId(v int32) {
	o.Id = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *Application) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *Application) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *Application) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAgentCacheInvalidatedResponseDuration returns the AgentCacheInvalidatedResponseDuration field value if set, zero value otherwise.
func (o *Application) GetAgentCacheInvalidatedResponseDuration() int32 {
	if o == nil || IsNil(o.AgentCacheInvalidatedResponseDuration) {
		var ret int32
		return ret
	}
	return *o.AgentCacheInvalidatedResponseDuration
}

// GetAgentCacheInvalidatedResponseDurationOk returns a tuple with the AgentCacheInvalidatedResponseDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetAgentCacheInvalidatedResponseDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.AgentCacheInvalidatedResponseDuration) {
		return nil, false
	}
	return o.AgentCacheInvalidatedResponseDuration, true
}

// HasAgentCacheInvalidatedResponseDuration returns a boolean if a field has been set.
func (o *Application) HasAgentCacheInvalidatedResponseDuration() bool {
	if o != nil && !IsNil(o.AgentCacheInvalidatedResponseDuration) {
		return true
	}

	return false
}

// SetAgentCacheInvalidatedResponseDuration gets a reference to the given int32 and assigns it to the AgentCacheInvalidatedResponseDuration field.
func (o *Application) SetAgentCacheInvalidatedResponseDuration(v int32) {
	o.AgentCacheInvalidatedResponseDuration = &v
}

// GetAgentCacheInvalidatedExpiration returns the AgentCacheInvalidatedExpiration field value if set, zero value otherwise.
func (o *Application) GetAgentCacheInvalidatedExpiration() int64 {
	if o == nil || IsNil(o.AgentCacheInvalidatedExpiration) {
		var ret int64
		return ret
	}
	return *o.AgentCacheInvalidatedExpiration
}

// GetAgentCacheInvalidatedExpirationOk returns a tuple with the AgentCacheInvalidatedExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetAgentCacheInvalidatedExpirationOk() (*int64, bool) {
	if o == nil || IsNil(o.AgentCacheInvalidatedExpiration) {
		return nil, false
	}
	return o.AgentCacheInvalidatedExpiration, true
}

// HasAgentCacheInvalidatedExpiration returns a boolean if a field has been set.
func (o *Application) HasAgentCacheInvalidatedExpiration() bool {
	if o != nil && !IsNil(o.AgentCacheInvalidatedExpiration) {
		return true
	}

	return false
}

// SetAgentCacheInvalidatedExpiration gets a reference to the given int64 and assigns it to the AgentCacheInvalidatedExpiration field.
func (o *Application) SetAgentCacheInvalidatedExpiration(v int64) {
	o.AgentCacheInvalidatedExpiration = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *Application) GetLastModified() int64 {
	if o == nil || IsNil(o.LastModified) {
		var ret int64
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetLastModifiedOk() (*int64, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *Application) HasLastModified() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given int64 and assigns it to the LastModified field.
func (o *Application) SetLastModified(v int64) {
	o.LastModified = &v
}

// GetRequireHTTPS returns the RequireHTTPS field value if set, zero value otherwise.
func (o *Application) GetRequireHTTPS() bool {
	if o == nil || IsNil(o.RequireHTTPS) {
		var ret bool
		return ret
	}
	return *o.RequireHTTPS
}

// GetRequireHTTPSOk returns a tuple with the RequireHTTPS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetRequireHTTPSOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireHTTPS) {
		return nil, false
	}
	return o.RequireHTTPS, true
}

// HasRequireHTTPS returns a boolean if a field has been set.
func (o *Application) HasRequireHTTPS() bool {
	if o != nil && !IsNil(o.RequireHTTPS) {
		return true
	}

	return false
}

// SetRequireHTTPS gets a reference to the given bool and assigns it to the RequireHTTPS field.
func (o *Application) SetRequireHTTPS(v bool) {
	o.RequireHTTPS = &v
}

// GetName returns the Name field value
func (o *Application) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Application) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Application) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Application) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Application) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Application) SetDescription(v string) {
	o.Description = &v
}

// GetManualOrderingEnabled returns the ManualOrderingEnabled field value if set, zero value otherwise.
func (o *Application) GetManualOrderingEnabled() bool {
	if o == nil || IsNil(o.ManualOrderingEnabled) {
		var ret bool
		return ret
	}
	return *o.ManualOrderingEnabled
}

// GetManualOrderingEnabledOk returns a tuple with the ManualOrderingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetManualOrderingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ManualOrderingEnabled) {
		return nil, false
	}
	return o.ManualOrderingEnabled, true
}

// HasManualOrderingEnabled returns a boolean if a field has been set.
func (o *Application) HasManualOrderingEnabled() bool {
	if o != nil && !IsNil(o.ManualOrderingEnabled) {
		return true
	}

	return false
}

// SetManualOrderingEnabled gets a reference to the given bool and assigns it to the ManualOrderingEnabled field.
func (o *Application) SetManualOrderingEnabled(v bool) {
	o.ManualOrderingEnabled = &v
}

// GetResourceOrder returns the ResourceOrder field value if set, zero value otherwise.
func (o *Application) GetResourceOrder() []int32 {
	if o == nil || IsNil(o.ResourceOrder) {
		var ret []int32
		return ret
	}
	return o.ResourceOrder
}

// GetResourceOrderOk returns a tuple with the ResourceOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetResourceOrderOk() ([]int32, bool) {
	if o == nil || IsNil(o.ResourceOrder) {
		return nil, false
	}
	return o.ResourceOrder, true
}

// HasResourceOrder returns a boolean if a field has been set.
func (o *Application) HasResourceOrder() bool {
	if o != nil && !IsNil(o.ResourceOrder) {
		return true
	}

	return false
}

// SetResourceOrder gets a reference to the given []int32 and assigns it to the ResourceOrder field.
func (o *Application) SetResourceOrder(v []int32) {
	o.ResourceOrder = v
}

// GetApplicationType returns the ApplicationType field value if set, zero value otherwise.
func (o *Application) GetApplicationType() ApplicationType {
	if o == nil || IsNil(o.ApplicationType) {
		var ret ApplicationType
		return ret
	}
	return *o.ApplicationType
}

// GetApplicationTypeOk returns a tuple with the ApplicationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetApplicationTypeOk() (*ApplicationType, bool) {
	if o == nil || IsNil(o.ApplicationType) {
		return nil, false
	}
	return o.ApplicationType, true
}

// HasApplicationType returns a boolean if a field has been set.
func (o *Application) HasApplicationType() bool {
	if o != nil && !IsNil(o.ApplicationType) {
		return true
	}

	return false
}

// SetApplicationType gets a reference to the given ApplicationType and assigns it to the ApplicationType field.
func (o *Application) SetApplicationType(v ApplicationType) {
	o.ApplicationType = &v
}

// GetDefaultAuthType returns the DefaultAuthType field value
func (o *Application) GetDefaultAuthType() DefaultAuthType {
	if o == nil {
		var ret DefaultAuthType
		return ret
	}

	return o.DefaultAuthType
}

// GetDefaultAuthTypeOk returns a tuple with the DefaultAuthType field value
// and a boolean to check if the value has been set.
func (o *Application) GetDefaultAuthTypeOk() (*DefaultAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultAuthType, true
}

// SetDefaultAuthType sets field value
func (o *Application) SetDefaultAuthType(v DefaultAuthType) {
	o.DefaultAuthType = v
}

// GetSpaSupportEnabled returns the SpaSupportEnabled field value
func (o *Application) GetSpaSupportEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SpaSupportEnabled
}

// GetSpaSupportEnabledOk returns a tuple with the SpaSupportEnabled field value
// and a boolean to check if the value has been set.
func (o *Application) GetSpaSupportEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpaSupportEnabled, true
}

// SetSpaSupportEnabled sets field value
func (o *Application) SetSpaSupportEnabled(v bool) {
	o.SpaSupportEnabled = v
}

// GetFallbackPostEncoding returns the FallbackPostEncoding field value if set, zero value otherwise.
func (o *Application) GetFallbackPostEncoding() string {
	if o == nil || IsNil(o.FallbackPostEncoding) {
		var ret string
		return ret
	}
	return *o.FallbackPostEncoding
}

// GetFallbackPostEncodingOk returns a tuple with the FallbackPostEncoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetFallbackPostEncodingOk() (*string, bool) {
	if o == nil || IsNil(o.FallbackPostEncoding) {
		return nil, false
	}
	return o.FallbackPostEncoding, true
}

// HasFallbackPostEncoding returns a boolean if a field has been set.
func (o *Application) HasFallbackPostEncoding() bool {
	if o != nil && !IsNil(o.FallbackPostEncoding) {
		return true
	}

	return false
}

// SetFallbackPostEncoding gets a reference to the given string and assigns it to the FallbackPostEncoding field.
func (o *Application) SetFallbackPostEncoding(v string) {
	o.FallbackPostEncoding = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *Application) GetDestination() Destination {
	if o == nil || IsNil(o.Destination) {
		var ret Destination
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetDestinationOk() (*Destination, bool) {
	if o == nil || IsNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *Application) HasDestination() bool {
	if o != nil && !IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given Destination and assigns it to the Destination field.
func (o *Application) SetDestination(v Destination) {
	o.Destination = &v
}

// GetContextRoot returns the ContextRoot field value
func (o *Application) GetContextRoot() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContextRoot
}

// GetContextRootOk returns a tuple with the ContextRoot field value
// and a boolean to check if the value has been set.
func (o *Application) GetContextRootOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContextRoot, true
}

// SetContextRoot sets field value
func (o *Application) SetContextRoot(v string) {
	o.ContextRoot = v
}

// GetRealm returns the Realm field value if set, zero value otherwise.
func (o *Application) GetRealm() string {
	if o == nil || IsNil(o.Realm) {
		var ret string
		return ret
	}
	return *o.Realm
}

// GetRealmOk returns a tuple with the Realm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetRealmOk() (*string, bool) {
	if o == nil || IsNil(o.Realm) {
		return nil, false
	}
	return o.Realm, true
}

// HasRealm returns a boolean if a field has been set.
func (o *Application) HasRealm() bool {
	if o != nil && !IsNil(o.Realm) {
		return true
	}

	return false
}

// SetRealm gets a reference to the given string and assigns it to the Realm field.
func (o *Application) SetRealm(v string) {
	o.Realm = &v
}

// GetCaseSensitivePath returns the CaseSensitivePath field value if set, zero value otherwise.
func (o *Application) GetCaseSensitivePath() bool {
	if o == nil || IsNil(o.CaseSensitivePath) {
		var ret bool
		return ret
	}
	return *o.CaseSensitivePath
}

// GetCaseSensitivePathOk returns a tuple with the CaseSensitivePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetCaseSensitivePathOk() (*bool, bool) {
	if o == nil || IsNil(o.CaseSensitivePath) {
		return nil, false
	}
	return o.CaseSensitivePath, true
}

// HasCaseSensitivePath returns a boolean if a field has been set.
func (o *Application) HasCaseSensitivePath() bool {
	if o != nil && !IsNil(o.CaseSensitivePath) {
		return true
	}

	return false
}

// SetCaseSensitivePath gets a reference to the given bool and assigns it to the CaseSensitivePath field.
func (o *Application) SetCaseSensitivePath(v bool) {
	o.CaseSensitivePath = &v
}

// GetAllowEmptyPathSegments returns the AllowEmptyPathSegments field value if set, zero value otherwise.
func (o *Application) GetAllowEmptyPathSegments() bool {
	if o == nil || IsNil(o.AllowEmptyPathSegments) {
		var ret bool
		return ret
	}
	return *o.AllowEmptyPathSegments
}

// GetAllowEmptyPathSegmentsOk returns a tuple with the AllowEmptyPathSegments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetAllowEmptyPathSegmentsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowEmptyPathSegments) {
		return nil, false
	}
	return o.AllowEmptyPathSegments, true
}

// HasAllowEmptyPathSegments returns a boolean if a field has been set.
func (o *Application) HasAllowEmptyPathSegments() bool {
	if o != nil && !IsNil(o.AllowEmptyPathSegments) {
		return true
	}

	return false
}

// SetAllowEmptyPathSegments gets a reference to the given bool and assigns it to the AllowEmptyPathSegments field.
func (o *Application) SetAllowEmptyPathSegments(v bool) {
	o.AllowEmptyPathSegments = &v
}

// GetWebSessionId returns the WebSessionId field value if set, zero value otherwise.
func (o *Application) GetWebSessionId() int32 {
	if o == nil || IsNil(o.WebSessionId) {
		var ret int32
		return ret
	}
	return *o.WebSessionId
}

// GetWebSessionIdOk returns a tuple with the WebSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetWebSessionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.WebSessionId) {
		return nil, false
	}
	return o.WebSessionId, true
}

// HasWebSessionId returns a boolean if a field has been set.
func (o *Application) HasWebSessionId() bool {
	if o != nil && !IsNil(o.WebSessionId) {
		return true
	}

	return false
}

// SetWebSessionId gets a reference to the given int32 and assigns it to the WebSessionId field.
func (o *Application) SetWebSessionId(v int32) {
	o.WebSessionId = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *Application) GetIssuer() string {
	if o == nil || IsNil(o.Issuer) {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetIssuerOk() (*string, bool) {
	if o == nil || IsNil(o.Issuer) {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *Application) HasIssuer() bool {
	if o != nil && !IsNil(o.Issuer) {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *Application) SetIssuer(v string) {
	o.Issuer = &v
}

// GetSiteId returns the SiteId field value
func (o *Application) GetSiteId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value
// and a boolean to check if the value has been set.
func (o *Application) GetSiteIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteId, true
}

// SetSiteId sets field value
func (o *Application) SetSiteId(v int32) {
	o.SiteId = v
}

// GetAgentId returns the AgentId field value
func (o *Application) GetAgentId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AgentId
}

// GetAgentIdOk returns a tuple with the AgentId field value
// and a boolean to check if the value has been set.
func (o *Application) GetAgentIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentId, true
}

// SetAgentId sets field value
func (o *Application) SetAgentId(v int32) {
	o.AgentId = v
}

// GetSidebandClientId returns the SidebandClientId field value
func (o *Application) GetSidebandClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SidebandClientId
}

// GetSidebandClientIdOk returns a tuple with the SidebandClientId field value
// and a boolean to check if the value has been set.
func (o *Application) GetSidebandClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SidebandClientId, true
}

// SetSidebandClientId sets field value
func (o *Application) SetSidebandClientId(v string) {
	o.SidebandClientId = v
}

// GetVirtualHostIds returns the VirtualHostIds field value
func (o *Application) GetVirtualHostIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.VirtualHostIds
}

// GetVirtualHostIdsOk returns a tuple with the VirtualHostIds field value
// and a boolean to check if the value has been set.
func (o *Application) GetVirtualHostIdsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.VirtualHostIds, true
}

// SetVirtualHostIds sets field value
func (o *Application) SetVirtualHostIds(v []int32) {
	o.VirtualHostIds = v
}

// GetIdentityMappingIds returns the IdentityMappingIds field value if set, zero value otherwise.
func (o *Application) GetIdentityMappingIds() map[string]int32 {
	if o == nil || IsNil(o.IdentityMappingIds) {
		var ret map[string]int32
		return ret
	}
	return *o.IdentityMappingIds
}

// GetIdentityMappingIdsOk returns a tuple with the IdentityMappingIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetIdentityMappingIdsOk() (*map[string]int32, bool) {
	if o == nil || IsNil(o.IdentityMappingIds) {
		return nil, false
	}
	return o.IdentityMappingIds, true
}

// HasIdentityMappingIds returns a boolean if a field has been set.
func (o *Application) HasIdentityMappingIds() bool {
	if o != nil && !IsNil(o.IdentityMappingIds) {
		return true
	}

	return false
}

// SetIdentityMappingIds gets a reference to the given map[string]int32 and assigns it to the IdentityMappingIds field.
func (o *Application) SetIdentityMappingIds(v map[string]int32) {
	o.IdentityMappingIds = &v
}

// GetAccessValidatorId returns the AccessValidatorId field value if set, zero value otherwise.
func (o *Application) GetAccessValidatorId() int32 {
	if o == nil || IsNil(o.AccessValidatorId) {
		var ret int32
		return ret
	}
	return *o.AccessValidatorId
}

// GetAccessValidatorIdOk returns a tuple with the AccessValidatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetAccessValidatorIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AccessValidatorId) {
		return nil, false
	}
	return o.AccessValidatorId, true
}

// HasAccessValidatorId returns a boolean if a field has been set.
func (o *Application) HasAccessValidatorId() bool {
	if o != nil && !IsNil(o.AccessValidatorId) {
		return true
	}

	return false
}

// SetAccessValidatorId gets a reference to the given int32 and assigns it to the AccessValidatorId field.
func (o *Application) SetAccessValidatorId(v int32) {
	o.AccessValidatorId = &v
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *Application) GetPolicy() map[string][]PolicyItem {
	if o == nil || IsNil(o.Policy) {
		var ret map[string][]PolicyItem
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetPolicyOk() (*map[string][]PolicyItem, bool) {
	if o == nil || IsNil(o.Policy) {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *Application) HasPolicy() bool {
	if o != nil && !IsNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given map[string][]PolicyItem and assigns it to the Policy field.
func (o *Application) SetPolicy(v map[string][]PolicyItem) {
	o.Policy = &v
}

// GetAuthenticationChallengePolicyId returns the AuthenticationChallengePolicyId field value
func (o *Application) GetAuthenticationChallengePolicyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthenticationChallengePolicyId
}

// GetAuthenticationChallengePolicyIdOk returns a tuple with the AuthenticationChallengePolicyId field value
// and a boolean to check if the value has been set.
func (o *Application) GetAuthenticationChallengePolicyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationChallengePolicyId, true
}

// SetAuthenticationChallengePolicyId sets field value
func (o *Application) SetAuthenticationChallengePolicyId(v string) {
	o.AuthenticationChallengePolicyId = v
}

// GetRiskPolicyId returns the RiskPolicyId field value if set, zero value otherwise.
func (o *Application) GetRiskPolicyId() int32 {
	if o == nil || IsNil(o.RiskPolicyId) {
		var ret int32
		return ret
	}
	return *o.RiskPolicyId
}

// GetRiskPolicyIdOk returns a tuple with the RiskPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetRiskPolicyIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RiskPolicyId) {
		return nil, false
	}
	return o.RiskPolicyId, true
}

// HasRiskPolicyId returns a boolean if a field has been set.
func (o *Application) HasRiskPolicyId() bool {
	if o != nil && !IsNil(o.RiskPolicyId) {
		return true
	}

	return false
}

// SetRiskPolicyId gets a reference to the given int32 and assigns it to the RiskPolicyId field.
func (o *Application) SetRiskPolicyId(v int32) {
	o.RiskPolicyId = &v
}

// GetDeployReservedResources returns the DeployReservedResources field value if set, zero value otherwise.
func (o *Application) GetDeployReservedResources() bool {
	if o == nil || IsNil(o.DeployReservedResources) {
		var ret bool
		return ret
	}
	return *o.DeployReservedResources
}

// GetDeployReservedResourcesOk returns a tuple with the DeployReservedResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetDeployReservedResourcesOk() (*bool, bool) {
	if o == nil || IsNil(o.DeployReservedResources) {
		return nil, false
	}
	return o.DeployReservedResources, true
}

// HasDeployReservedResources returns a boolean if a field has been set.
func (o *Application) HasDeployReservedResources() bool {
	if o != nil && !IsNil(o.DeployReservedResources) {
		return true
	}

	return false
}

// SetDeployReservedResources gets a reference to the given bool and assigns it to the DeployReservedResources field.
func (o *Application) SetDeployReservedResources(v bool) {
	o.DeployReservedResources = &v
}

func (o Application) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Application) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.AgentCacheInvalidatedResponseDuration) {
		toSerialize["agentCacheInvalidatedResponseDuration"] = o.AgentCacheInvalidatedResponseDuration
	}
	if !IsNil(o.AgentCacheInvalidatedExpiration) {
		toSerialize["agentCacheInvalidatedExpiration"] = o.AgentCacheInvalidatedExpiration
	}
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	if !IsNil(o.RequireHTTPS) {
		toSerialize["requireHTTPS"] = o.RequireHTTPS
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ManualOrderingEnabled) {
		toSerialize["manualOrderingEnabled"] = o.ManualOrderingEnabled
	}
	if !IsNil(o.ResourceOrder) {
		toSerialize["resourceOrder"] = o.ResourceOrder
	}
	if !IsNil(o.ApplicationType) {
		toSerialize["applicationType"] = o.ApplicationType
	}
	toSerialize["defaultAuthType"] = o.DefaultAuthType
	toSerialize["spaSupportEnabled"] = o.SpaSupportEnabled
	if !IsNil(o.FallbackPostEncoding) {
		toSerialize["fallbackPostEncoding"] = o.FallbackPostEncoding
	}
	if !IsNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	toSerialize["contextRoot"] = o.ContextRoot
	if !IsNil(o.Realm) {
		toSerialize["realm"] = o.Realm
	}
	if !IsNil(o.CaseSensitivePath) {
		toSerialize["caseSensitivePath"] = o.CaseSensitivePath
	}
	if !IsNil(o.AllowEmptyPathSegments) {
		toSerialize["allowEmptyPathSegments"] = o.AllowEmptyPathSegments
	}
	if !IsNil(o.WebSessionId) {
		toSerialize["webSessionId"] = o.WebSessionId
	}
	if !IsNil(o.Issuer) {
		toSerialize["issuer"] = o.Issuer
	}
	toSerialize["siteId"] = o.SiteId
	toSerialize["agentId"] = o.AgentId
	toSerialize["sidebandClientId"] = o.SidebandClientId
	toSerialize["virtualHostIds"] = o.VirtualHostIds
	if !IsNil(o.IdentityMappingIds) {
		toSerialize["identityMappingIds"] = o.IdentityMappingIds
	}
	if !IsNil(o.AccessValidatorId) {
		toSerialize["accessValidatorId"] = o.AccessValidatorId
	}
	if !IsNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	toSerialize["authenticationChallengePolicyId"] = o.AuthenticationChallengePolicyId
	if !IsNil(o.RiskPolicyId) {
		toSerialize["riskPolicyId"] = o.RiskPolicyId
	}
	if !IsNil(o.DeployReservedResources) {
		toSerialize["deployReservedResources"] = o.DeployReservedResources
	}
	return toSerialize, nil
}

type NullableApplication struct {
	value *Application
	isSet bool
}

func (v NullableApplication) Get() *Application {
	return v.value
}

func (v *NullableApplication) Set(val *Application) {
	v.value = val
	v.isSet = true
}

func (v NullableApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplication(val *Application) *NullableApplication {
	return &NullableApplication{value: val, isSet: true}
}

func (v NullableApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
