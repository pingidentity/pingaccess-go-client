# OIDCProviderMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issuer** | Pointer to **string** | OpenID Connect provider&#39;s issuer identifier URL. | [optional] 
**AuthorizationEndpoint** | Pointer to **string** | URL of the OpenID Connect provider&#39;s authorization endpoint. | [optional] 
**TokenEndpoint** | Pointer to **string** | URL of the OpenID Connect provider&#39;s token endpoint. | [optional] 
**RevocationEndpoint** | Pointer to **string** | URL of the OpenID Connect provider&#39;s OAuth 2.0 revocation endpoint. | [optional] 
**UserinfoEndpoint** | Pointer to **string** | URL of the OpenID Connect provider&#39;s userInfo endpoint. | [optional] 
**IntrospectionEndpoint** | Pointer to **string** | URL of the OpenID Connect provider&#39;s OAuth 2.0 introspection endpoint. | [optional] 
**JwksUri** | Pointer to **string** | URL of the OpenID Connect provider&#39;s JWK Set document. | [optional] 
**ScopesSupported** | Pointer to **[]string** | JSON array containing a list of the OAuth 2.0 \&quot;scope\&quot; values that this OpenID Connect provider supports. | [optional] 
**ClaimsSupported** | Pointer to **[]string** | JSON array containing a list of the claim names of the claims that the OpenID Connect provider MAY be able to supply values for. | [optional] 
**ResponseTypesSupported** | Pointer to **[]string** | JSON array containing a list of the OAuth 2.0 \&quot;response_type\&quot; values that this OpenID Connect provider supports. | [optional] 
**ResponseModesSupported** | Pointer to **[]string** | JSON array containing a list of the OAuth 2.0 \&quot;response_mode\&quot; values that this OpenID Connect provider supports. | [optional] 
**SubjectTypesSupported** | Pointer to **[]string** | JSON array containing a list of the Subject Identifier types that this OpenID Connect provider supports. | [optional] 
**IdTokenSigningAlgValuesSupported** | Pointer to **[]string** | JSON array containing a list of the JWS signing algorithms supported by the OpenID Connect provider for the id token to encode the claims in a JWT. | [optional] 
**TokenEndpointAuthMethodsSupported** | Pointer to **[]string** | JSON array containing a list of client authentication methods supported by this token endpoint. | [optional] 
**TokenEndpointAuthSigningAlgValuesSupported** | Pointer to **[]string** | JSON array containing a list of client authentication signing algorithms supported by this token endpoint. | [optional] 
**ClaimTypesSupported** | Pointer to **[]string** | JSON array containing a list of the claim types that the OpenID Connect provider supports. | [optional] 
**CodeChallengeMethodsSupported** | Pointer to **[]string** | Proof Key for Code Exchange (PKCE) code challenge methods supported by this OpenID Connect provider. | [optional] 
**ClaimsParameterSupported** | Pointer to **bool** | boolean value specifying whether the OpenID Connect provider supports use of the claims parameter, with true indicating support. | [optional] 
**RequestParameterSupported** | Pointer to **bool** | boolean value specifying whether the OpenID Connect provider supports use of the request parameter, with true indicating support. | [optional] 
**RequestUriParameterSupported** | Pointer to **bool** | boolean value specifying whether the OpenID Connect provider supports use of the request_uri parameter, with true indicating support. | [optional] 
**EndSessionEndpoint** | Pointer to **string** | URL at the OpenID Connect provider to which a relying party can perform a redirect to request that the end-user be logged out at the OpenID Connect provider. | [optional] 
**PingRevokedSrisEndpoint** | Pointer to **string** | PingFederate session revocation endpoint. (Not applicable if PingFederate is not the OpenID Connect provider) | [optional] 
**PingEndSessionEndpoint** | Pointer to **string** | PingFederate logout endpoint. (Not applicable if PingFederate is not the OpenID Connect provider) | [optional] 
**GrantTypesSupported** | Pointer to **[]string** | JSON array containing a list of the OAuth 2.0 grant type values that this OpenID Connect provider supports. | [optional] 
**UserinfoSigningAlgValuesSupported** | Pointer to **[]string** | JSON array containing a list of the JWS signing algorithms supported by the userInfo endpoint to encode the claims in a JWT. | [optional] 
**RequestObjectSigningAlgValuesSupported** | Pointer to **[]string** | JSON array containing a list of the JWS signing algorithms supported by the OpenID Connect provider for request objects. | [optional] 
**BackchannelAuthenticationEndpoint** | Pointer to **string** | the endpoint used to initiate an out-of-band authentication. | [optional] 
**MtlsEndpointAliases** | Pointer to **map[string]string** | a map of alternative authorization server endpoints. The key is an authorization server endpoint, and the value is a preferred URL for the endpoint. | [optional] 

## Methods

### NewOIDCProviderMetadata

`func NewOIDCProviderMetadata() *OIDCProviderMetadata`

NewOIDCProviderMetadata instantiates a new OIDCProviderMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOIDCProviderMetadataWithDefaults

`func NewOIDCProviderMetadataWithDefaults() *OIDCProviderMetadata`

NewOIDCProviderMetadataWithDefaults instantiates a new OIDCProviderMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuer

`func (o *OIDCProviderMetadata) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *OIDCProviderMetadata) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *OIDCProviderMetadata) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *OIDCProviderMetadata) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetAuthorizationEndpoint

`func (o *OIDCProviderMetadata) GetAuthorizationEndpoint() string`

GetAuthorizationEndpoint returns the AuthorizationEndpoint field if non-nil, zero value otherwise.

### GetAuthorizationEndpointOk

`func (o *OIDCProviderMetadata) GetAuthorizationEndpointOk() (*string, bool)`

GetAuthorizationEndpointOk returns a tuple with the AuthorizationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationEndpoint

`func (o *OIDCProviderMetadata) SetAuthorizationEndpoint(v string)`

SetAuthorizationEndpoint sets AuthorizationEndpoint field to given value.

### HasAuthorizationEndpoint

`func (o *OIDCProviderMetadata) HasAuthorizationEndpoint() bool`

HasAuthorizationEndpoint returns a boolean if a field has been set.

### GetTokenEndpoint

`func (o *OIDCProviderMetadata) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *OIDCProviderMetadata) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *OIDCProviderMetadata) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.

### HasTokenEndpoint

`func (o *OIDCProviderMetadata) HasTokenEndpoint() bool`

HasTokenEndpoint returns a boolean if a field has been set.

### GetRevocationEndpoint

`func (o *OIDCProviderMetadata) GetRevocationEndpoint() string`

GetRevocationEndpoint returns the RevocationEndpoint field if non-nil, zero value otherwise.

### GetRevocationEndpointOk

`func (o *OIDCProviderMetadata) GetRevocationEndpointOk() (*string, bool)`

GetRevocationEndpointOk returns a tuple with the RevocationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationEndpoint

`func (o *OIDCProviderMetadata) SetRevocationEndpoint(v string)`

SetRevocationEndpoint sets RevocationEndpoint field to given value.

### HasRevocationEndpoint

`func (o *OIDCProviderMetadata) HasRevocationEndpoint() bool`

HasRevocationEndpoint returns a boolean if a field has been set.

### GetUserinfoEndpoint

`func (o *OIDCProviderMetadata) GetUserinfoEndpoint() string`

GetUserinfoEndpoint returns the UserinfoEndpoint field if non-nil, zero value otherwise.

### GetUserinfoEndpointOk

`func (o *OIDCProviderMetadata) GetUserinfoEndpointOk() (*string, bool)`

GetUserinfoEndpointOk returns a tuple with the UserinfoEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserinfoEndpoint

`func (o *OIDCProviderMetadata) SetUserinfoEndpoint(v string)`

SetUserinfoEndpoint sets UserinfoEndpoint field to given value.

### HasUserinfoEndpoint

`func (o *OIDCProviderMetadata) HasUserinfoEndpoint() bool`

HasUserinfoEndpoint returns a boolean if a field has been set.

### GetIntrospectionEndpoint

`func (o *OIDCProviderMetadata) GetIntrospectionEndpoint() string`

GetIntrospectionEndpoint returns the IntrospectionEndpoint field if non-nil, zero value otherwise.

### GetIntrospectionEndpointOk

`func (o *OIDCProviderMetadata) GetIntrospectionEndpointOk() (*string, bool)`

GetIntrospectionEndpointOk returns a tuple with the IntrospectionEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntrospectionEndpoint

`func (o *OIDCProviderMetadata) SetIntrospectionEndpoint(v string)`

SetIntrospectionEndpoint sets IntrospectionEndpoint field to given value.

### HasIntrospectionEndpoint

`func (o *OIDCProviderMetadata) HasIntrospectionEndpoint() bool`

HasIntrospectionEndpoint returns a boolean if a field has been set.

### GetJwksUri

`func (o *OIDCProviderMetadata) GetJwksUri() string`

GetJwksUri returns the JwksUri field if non-nil, zero value otherwise.

### GetJwksUriOk

`func (o *OIDCProviderMetadata) GetJwksUriOk() (*string, bool)`

GetJwksUriOk returns a tuple with the JwksUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUri

`func (o *OIDCProviderMetadata) SetJwksUri(v string)`

SetJwksUri sets JwksUri field to given value.

### HasJwksUri

`func (o *OIDCProviderMetadata) HasJwksUri() bool`

HasJwksUri returns a boolean if a field has been set.

### GetScopesSupported

`func (o *OIDCProviderMetadata) GetScopesSupported() []string`

GetScopesSupported returns the ScopesSupported field if non-nil, zero value otherwise.

### GetScopesSupportedOk

`func (o *OIDCProviderMetadata) GetScopesSupportedOk() (*[]string, bool)`

GetScopesSupportedOk returns a tuple with the ScopesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopesSupported

`func (o *OIDCProviderMetadata) SetScopesSupported(v []string)`

SetScopesSupported sets ScopesSupported field to given value.

### HasScopesSupported

`func (o *OIDCProviderMetadata) HasScopesSupported() bool`

HasScopesSupported returns a boolean if a field has been set.

### GetClaimsSupported

`func (o *OIDCProviderMetadata) GetClaimsSupported() []string`

GetClaimsSupported returns the ClaimsSupported field if non-nil, zero value otherwise.

### GetClaimsSupportedOk

`func (o *OIDCProviderMetadata) GetClaimsSupportedOk() (*[]string, bool)`

GetClaimsSupportedOk returns a tuple with the ClaimsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimsSupported

`func (o *OIDCProviderMetadata) SetClaimsSupported(v []string)`

SetClaimsSupported sets ClaimsSupported field to given value.

### HasClaimsSupported

`func (o *OIDCProviderMetadata) HasClaimsSupported() bool`

HasClaimsSupported returns a boolean if a field has been set.

### GetResponseTypesSupported

`func (o *OIDCProviderMetadata) GetResponseTypesSupported() []string`

GetResponseTypesSupported returns the ResponseTypesSupported field if non-nil, zero value otherwise.

### GetResponseTypesSupportedOk

`func (o *OIDCProviderMetadata) GetResponseTypesSupportedOk() (*[]string, bool)`

GetResponseTypesSupportedOk returns a tuple with the ResponseTypesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTypesSupported

`func (o *OIDCProviderMetadata) SetResponseTypesSupported(v []string)`

SetResponseTypesSupported sets ResponseTypesSupported field to given value.

### HasResponseTypesSupported

`func (o *OIDCProviderMetadata) HasResponseTypesSupported() bool`

HasResponseTypesSupported returns a boolean if a field has been set.

### GetResponseModesSupported

`func (o *OIDCProviderMetadata) GetResponseModesSupported() []string`

GetResponseModesSupported returns the ResponseModesSupported field if non-nil, zero value otherwise.

### GetResponseModesSupportedOk

`func (o *OIDCProviderMetadata) GetResponseModesSupportedOk() (*[]string, bool)`

GetResponseModesSupportedOk returns a tuple with the ResponseModesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseModesSupported

`func (o *OIDCProviderMetadata) SetResponseModesSupported(v []string)`

SetResponseModesSupported sets ResponseModesSupported field to given value.

### HasResponseModesSupported

`func (o *OIDCProviderMetadata) HasResponseModesSupported() bool`

HasResponseModesSupported returns a boolean if a field has been set.

### GetSubjectTypesSupported

`func (o *OIDCProviderMetadata) GetSubjectTypesSupported() []string`

GetSubjectTypesSupported returns the SubjectTypesSupported field if non-nil, zero value otherwise.

### GetSubjectTypesSupportedOk

`func (o *OIDCProviderMetadata) GetSubjectTypesSupportedOk() (*[]string, bool)`

GetSubjectTypesSupportedOk returns a tuple with the SubjectTypesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectTypesSupported

`func (o *OIDCProviderMetadata) SetSubjectTypesSupported(v []string)`

SetSubjectTypesSupported sets SubjectTypesSupported field to given value.

### HasSubjectTypesSupported

`func (o *OIDCProviderMetadata) HasSubjectTypesSupported() bool`

HasSubjectTypesSupported returns a boolean if a field has been set.

### GetIdTokenSigningAlgValuesSupported

`func (o *OIDCProviderMetadata) GetIdTokenSigningAlgValuesSupported() []string`

GetIdTokenSigningAlgValuesSupported returns the IdTokenSigningAlgValuesSupported field if non-nil, zero value otherwise.

### GetIdTokenSigningAlgValuesSupportedOk

`func (o *OIDCProviderMetadata) GetIdTokenSigningAlgValuesSupportedOk() (*[]string, bool)`

GetIdTokenSigningAlgValuesSupportedOk returns a tuple with the IdTokenSigningAlgValuesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenSigningAlgValuesSupported

`func (o *OIDCProviderMetadata) SetIdTokenSigningAlgValuesSupported(v []string)`

SetIdTokenSigningAlgValuesSupported sets IdTokenSigningAlgValuesSupported field to given value.

### HasIdTokenSigningAlgValuesSupported

`func (o *OIDCProviderMetadata) HasIdTokenSigningAlgValuesSupported() bool`

HasIdTokenSigningAlgValuesSupported returns a boolean if a field has been set.

### GetTokenEndpointAuthMethodsSupported

`func (o *OIDCProviderMetadata) GetTokenEndpointAuthMethodsSupported() []string`

GetTokenEndpointAuthMethodsSupported returns the TokenEndpointAuthMethodsSupported field if non-nil, zero value otherwise.

### GetTokenEndpointAuthMethodsSupportedOk

`func (o *OIDCProviderMetadata) GetTokenEndpointAuthMethodsSupportedOk() (*[]string, bool)`

GetTokenEndpointAuthMethodsSupportedOk returns a tuple with the TokenEndpointAuthMethodsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointAuthMethodsSupported

`func (o *OIDCProviderMetadata) SetTokenEndpointAuthMethodsSupported(v []string)`

SetTokenEndpointAuthMethodsSupported sets TokenEndpointAuthMethodsSupported field to given value.

### HasTokenEndpointAuthMethodsSupported

`func (o *OIDCProviderMetadata) HasTokenEndpointAuthMethodsSupported() bool`

HasTokenEndpointAuthMethodsSupported returns a boolean if a field has been set.

### GetTokenEndpointAuthSigningAlgValuesSupported

`func (o *OIDCProviderMetadata) GetTokenEndpointAuthSigningAlgValuesSupported() []string`

GetTokenEndpointAuthSigningAlgValuesSupported returns the TokenEndpointAuthSigningAlgValuesSupported field if non-nil, zero value otherwise.

### GetTokenEndpointAuthSigningAlgValuesSupportedOk

`func (o *OIDCProviderMetadata) GetTokenEndpointAuthSigningAlgValuesSupportedOk() (*[]string, bool)`

GetTokenEndpointAuthSigningAlgValuesSupportedOk returns a tuple with the TokenEndpointAuthSigningAlgValuesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointAuthSigningAlgValuesSupported

`func (o *OIDCProviderMetadata) SetTokenEndpointAuthSigningAlgValuesSupported(v []string)`

SetTokenEndpointAuthSigningAlgValuesSupported sets TokenEndpointAuthSigningAlgValuesSupported field to given value.

### HasTokenEndpointAuthSigningAlgValuesSupported

`func (o *OIDCProviderMetadata) HasTokenEndpointAuthSigningAlgValuesSupported() bool`

HasTokenEndpointAuthSigningAlgValuesSupported returns a boolean if a field has been set.

### GetClaimTypesSupported

`func (o *OIDCProviderMetadata) GetClaimTypesSupported() []string`

GetClaimTypesSupported returns the ClaimTypesSupported field if non-nil, zero value otherwise.

### GetClaimTypesSupportedOk

`func (o *OIDCProviderMetadata) GetClaimTypesSupportedOk() (*[]string, bool)`

GetClaimTypesSupportedOk returns a tuple with the ClaimTypesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimTypesSupported

`func (o *OIDCProviderMetadata) SetClaimTypesSupported(v []string)`

SetClaimTypesSupported sets ClaimTypesSupported field to given value.

### HasClaimTypesSupported

`func (o *OIDCProviderMetadata) HasClaimTypesSupported() bool`

HasClaimTypesSupported returns a boolean if a field has been set.

### GetCodeChallengeMethodsSupported

`func (o *OIDCProviderMetadata) GetCodeChallengeMethodsSupported() []string`

GetCodeChallengeMethodsSupported returns the CodeChallengeMethodsSupported field if non-nil, zero value otherwise.

### GetCodeChallengeMethodsSupportedOk

`func (o *OIDCProviderMetadata) GetCodeChallengeMethodsSupportedOk() (*[]string, bool)`

GetCodeChallengeMethodsSupportedOk returns a tuple with the CodeChallengeMethodsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeChallengeMethodsSupported

`func (o *OIDCProviderMetadata) SetCodeChallengeMethodsSupported(v []string)`

SetCodeChallengeMethodsSupported sets CodeChallengeMethodsSupported field to given value.

### HasCodeChallengeMethodsSupported

`func (o *OIDCProviderMetadata) HasCodeChallengeMethodsSupported() bool`

HasCodeChallengeMethodsSupported returns a boolean if a field has been set.

### GetClaimsParameterSupported

`func (o *OIDCProviderMetadata) GetClaimsParameterSupported() bool`

GetClaimsParameterSupported returns the ClaimsParameterSupported field if non-nil, zero value otherwise.

### GetClaimsParameterSupportedOk

`func (o *OIDCProviderMetadata) GetClaimsParameterSupportedOk() (*bool, bool)`

GetClaimsParameterSupportedOk returns a tuple with the ClaimsParameterSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimsParameterSupported

`func (o *OIDCProviderMetadata) SetClaimsParameterSupported(v bool)`

SetClaimsParameterSupported sets ClaimsParameterSupported field to given value.

### HasClaimsParameterSupported

`func (o *OIDCProviderMetadata) HasClaimsParameterSupported() bool`

HasClaimsParameterSupported returns a boolean if a field has been set.

### GetRequestParameterSupported

`func (o *OIDCProviderMetadata) GetRequestParameterSupported() bool`

GetRequestParameterSupported returns the RequestParameterSupported field if non-nil, zero value otherwise.

### GetRequestParameterSupportedOk

`func (o *OIDCProviderMetadata) GetRequestParameterSupportedOk() (*bool, bool)`

GetRequestParameterSupportedOk returns a tuple with the RequestParameterSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestParameterSupported

`func (o *OIDCProviderMetadata) SetRequestParameterSupported(v bool)`

SetRequestParameterSupported sets RequestParameterSupported field to given value.

### HasRequestParameterSupported

`func (o *OIDCProviderMetadata) HasRequestParameterSupported() bool`

HasRequestParameterSupported returns a boolean if a field has been set.

### GetRequestUriParameterSupported

`func (o *OIDCProviderMetadata) GetRequestUriParameterSupported() bool`

GetRequestUriParameterSupported returns the RequestUriParameterSupported field if non-nil, zero value otherwise.

### GetRequestUriParameterSupportedOk

`func (o *OIDCProviderMetadata) GetRequestUriParameterSupportedOk() (*bool, bool)`

GetRequestUriParameterSupportedOk returns a tuple with the RequestUriParameterSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUriParameterSupported

`func (o *OIDCProviderMetadata) SetRequestUriParameterSupported(v bool)`

SetRequestUriParameterSupported sets RequestUriParameterSupported field to given value.

### HasRequestUriParameterSupported

`func (o *OIDCProviderMetadata) HasRequestUriParameterSupported() bool`

HasRequestUriParameterSupported returns a boolean if a field has been set.

### GetEndSessionEndpoint

`func (o *OIDCProviderMetadata) GetEndSessionEndpoint() string`

GetEndSessionEndpoint returns the EndSessionEndpoint field if non-nil, zero value otherwise.

### GetEndSessionEndpointOk

`func (o *OIDCProviderMetadata) GetEndSessionEndpointOk() (*string, bool)`

GetEndSessionEndpointOk returns a tuple with the EndSessionEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndSessionEndpoint

`func (o *OIDCProviderMetadata) SetEndSessionEndpoint(v string)`

SetEndSessionEndpoint sets EndSessionEndpoint field to given value.

### HasEndSessionEndpoint

`func (o *OIDCProviderMetadata) HasEndSessionEndpoint() bool`

HasEndSessionEndpoint returns a boolean if a field has been set.

### GetPingRevokedSrisEndpoint

`func (o *OIDCProviderMetadata) GetPingRevokedSrisEndpoint() string`

GetPingRevokedSrisEndpoint returns the PingRevokedSrisEndpoint field if non-nil, zero value otherwise.

### GetPingRevokedSrisEndpointOk

`func (o *OIDCProviderMetadata) GetPingRevokedSrisEndpointOk() (*string, bool)`

GetPingRevokedSrisEndpointOk returns a tuple with the PingRevokedSrisEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingRevokedSrisEndpoint

`func (o *OIDCProviderMetadata) SetPingRevokedSrisEndpoint(v string)`

SetPingRevokedSrisEndpoint sets PingRevokedSrisEndpoint field to given value.

### HasPingRevokedSrisEndpoint

`func (o *OIDCProviderMetadata) HasPingRevokedSrisEndpoint() bool`

HasPingRevokedSrisEndpoint returns a boolean if a field has been set.

### GetPingEndSessionEndpoint

`func (o *OIDCProviderMetadata) GetPingEndSessionEndpoint() string`

GetPingEndSessionEndpoint returns the PingEndSessionEndpoint field if non-nil, zero value otherwise.

### GetPingEndSessionEndpointOk

`func (o *OIDCProviderMetadata) GetPingEndSessionEndpointOk() (*string, bool)`

GetPingEndSessionEndpointOk returns a tuple with the PingEndSessionEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingEndSessionEndpoint

`func (o *OIDCProviderMetadata) SetPingEndSessionEndpoint(v string)`

SetPingEndSessionEndpoint sets PingEndSessionEndpoint field to given value.

### HasPingEndSessionEndpoint

`func (o *OIDCProviderMetadata) HasPingEndSessionEndpoint() bool`

HasPingEndSessionEndpoint returns a boolean if a field has been set.

### GetGrantTypesSupported

`func (o *OIDCProviderMetadata) GetGrantTypesSupported() []string`

GetGrantTypesSupported returns the GrantTypesSupported field if non-nil, zero value otherwise.

### GetGrantTypesSupportedOk

`func (o *OIDCProviderMetadata) GetGrantTypesSupportedOk() (*[]string, bool)`

GetGrantTypesSupportedOk returns a tuple with the GrantTypesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypesSupported

`func (o *OIDCProviderMetadata) SetGrantTypesSupported(v []string)`

SetGrantTypesSupported sets GrantTypesSupported field to given value.

### HasGrantTypesSupported

`func (o *OIDCProviderMetadata) HasGrantTypesSupported() bool`

HasGrantTypesSupported returns a boolean if a field has been set.

### GetUserinfoSigningAlgValuesSupported

`func (o *OIDCProviderMetadata) GetUserinfoSigningAlgValuesSupported() []string`

GetUserinfoSigningAlgValuesSupported returns the UserinfoSigningAlgValuesSupported field if non-nil, zero value otherwise.

### GetUserinfoSigningAlgValuesSupportedOk

`func (o *OIDCProviderMetadata) GetUserinfoSigningAlgValuesSupportedOk() (*[]string, bool)`

GetUserinfoSigningAlgValuesSupportedOk returns a tuple with the UserinfoSigningAlgValuesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserinfoSigningAlgValuesSupported

`func (o *OIDCProviderMetadata) SetUserinfoSigningAlgValuesSupported(v []string)`

SetUserinfoSigningAlgValuesSupported sets UserinfoSigningAlgValuesSupported field to given value.

### HasUserinfoSigningAlgValuesSupported

`func (o *OIDCProviderMetadata) HasUserinfoSigningAlgValuesSupported() bool`

HasUserinfoSigningAlgValuesSupported returns a boolean if a field has been set.

### GetRequestObjectSigningAlgValuesSupported

`func (o *OIDCProviderMetadata) GetRequestObjectSigningAlgValuesSupported() []string`

GetRequestObjectSigningAlgValuesSupported returns the RequestObjectSigningAlgValuesSupported field if non-nil, zero value otherwise.

### GetRequestObjectSigningAlgValuesSupportedOk

`func (o *OIDCProviderMetadata) GetRequestObjectSigningAlgValuesSupportedOk() (*[]string, bool)`

GetRequestObjectSigningAlgValuesSupportedOk returns a tuple with the RequestObjectSigningAlgValuesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestObjectSigningAlgValuesSupported

`func (o *OIDCProviderMetadata) SetRequestObjectSigningAlgValuesSupported(v []string)`

SetRequestObjectSigningAlgValuesSupported sets RequestObjectSigningAlgValuesSupported field to given value.

### HasRequestObjectSigningAlgValuesSupported

`func (o *OIDCProviderMetadata) HasRequestObjectSigningAlgValuesSupported() bool`

HasRequestObjectSigningAlgValuesSupported returns a boolean if a field has been set.

### GetBackchannelAuthenticationEndpoint

`func (o *OIDCProviderMetadata) GetBackchannelAuthenticationEndpoint() string`

GetBackchannelAuthenticationEndpoint returns the BackchannelAuthenticationEndpoint field if non-nil, zero value otherwise.

### GetBackchannelAuthenticationEndpointOk

`func (o *OIDCProviderMetadata) GetBackchannelAuthenticationEndpointOk() (*string, bool)`

GetBackchannelAuthenticationEndpointOk returns a tuple with the BackchannelAuthenticationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackchannelAuthenticationEndpoint

`func (o *OIDCProviderMetadata) SetBackchannelAuthenticationEndpoint(v string)`

SetBackchannelAuthenticationEndpoint sets BackchannelAuthenticationEndpoint field to given value.

### HasBackchannelAuthenticationEndpoint

`func (o *OIDCProviderMetadata) HasBackchannelAuthenticationEndpoint() bool`

HasBackchannelAuthenticationEndpoint returns a boolean if a field has been set.

### GetMtlsEndpointAliases

`func (o *OIDCProviderMetadata) GetMtlsEndpointAliases() map[string]string`

GetMtlsEndpointAliases returns the MtlsEndpointAliases field if non-nil, zero value otherwise.

### GetMtlsEndpointAliasesOk

`func (o *OIDCProviderMetadata) GetMtlsEndpointAliasesOk() (*map[string]string, bool)`

GetMtlsEndpointAliasesOk returns a tuple with the MtlsEndpointAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtlsEndpointAliases

`func (o *OIDCProviderMetadata) SetMtlsEndpointAliases(v map[string]string)`

SetMtlsEndpointAliases sets MtlsEndpointAliases field to given value.

### HasMtlsEndpointAliases

`func (o *OIDCProviderMetadata) HasMtlsEndpointAliases() bool`

HasMtlsEndpointAliases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


