# \AuthApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAdminBasicWebSession**](AuthApi.md#DeleteAdminBasicWebSession) | **Delete** /auth/webSession | 
[**DeleteAdminTokenProvider**](AuthApi.md#DeleteAdminTokenProvider) | **Delete** /auth/tokenProvider | 
[**DeleteBasicAuth**](AuthApi.md#DeleteBasicAuth) | **Delete** /auth/basic | 
[**DeleteOAuthAuth**](AuthApi.md#DeleteOAuthAuth) | **Delete** /auth/oauth | 
[**DeleteOidcAuth**](AuthApi.md#DeleteOidcAuth) | **Delete** /auth/oidc | 
[**GetAdminBasicWebSession**](AuthApi.md#GetAdminBasicWebSession) | **Get** /auth/webSession | 
[**GetAdminTokenProvider**](AuthApi.md#GetAdminTokenProvider) | **Get** /auth/tokenProvider | 
[**GetAdminTokenProviderMetadata**](AuthApi.md#GetAdminTokenProviderMetadata) | **Get** /auth/tokenProvider/metadata | 
[**GetAuthOidcScopes**](AuthApi.md#GetAuthOidcScopes) | **Get** /auth/oidc/scopes | 
[**GetBasicAuth**](AuthApi.md#GetBasicAuth) | **Get** /auth/basic | 
[**GetOAuthAuth**](AuthApi.md#GetOAuthAuth) | **Get** /auth/oauth | 
[**GetOidcAuth**](AuthApi.md#GetOidcAuth) | **Get** /auth/oidc | 
[**UpdateAdminBasicWebSession**](AuthApi.md#UpdateAdminBasicWebSession) | **Put** /auth/webSession | 
[**UpdateAdminTokenProvider**](AuthApi.md#UpdateAdminTokenProvider) | **Put** /auth/tokenProvider | 
[**UpdateBasicAuth**](AuthApi.md#UpdateBasicAuth) | **Put** /auth/basic | 
[**UpdateOAuthAuth**](AuthApi.md#UpdateOAuthAuth) | **Put** /auth/oauth | 
[**UpdateOidcAuth**](AuthApi.md#UpdateOidcAuth) | **Put** /auth/oidc | 



## DeleteAdminBasicWebSession

> DeleteAdminBasicWebSession(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingaccess-go-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthApi.DeleteAdminBasicWebSession(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.DeleteAdminBasicWebSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAdminBasicWebSessionRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAdminTokenProvider

> DeleteAdminTokenProvider(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingaccess-go-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthApi.DeleteAdminTokenProvider(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.DeleteAdminTokenProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAdminTokenProviderRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBasicAuth

> DeleteBasicAuth(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingaccess-go-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthApi.DeleteBasicAuth(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.DeleteBasicAuth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBasicAuthRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOAuthAuth

> DeleteOAuthAuth(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingaccess-go-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthApi.DeleteOAuthAuth(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.DeleteOAuthAuth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOAuthAuthRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOidcAuth

> DeleteOidcAuth(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingaccess-go-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthApi.DeleteOidcAuth(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.DeleteOidcAuth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOidcAuthRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminBasicWebSession

> AdminBasicWebSession GetAdminBasicWebSession(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingaccess-go-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.GetAdminBasicWebSession(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.GetAdminBasicWebSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdminBasicWebSession`: AdminBasicWebSession
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.GetAdminBasicWebSession`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminBasicWebSessionRequest struct via the builder pattern


### Return type

[**AdminBasicWebSession**](AdminBasicWebSession.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminTokenProvider

> AdminTokenProvider GetAdminTokenProvider(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingaccess-go-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.GetAdminTokenProvider(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.GetAdminTokenProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdminTokenProvider`: AdminTokenProvider
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.GetAdminTokenProvider`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminTokenProviderRequest struct via the builder pattern


### Return type

[**AdminTokenProvider**](AdminTokenProvider.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminTokenProviderMetadata

> OIDCProviderMetadata GetAdminTokenProviderMetadata(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingaccess-go-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.GetAdminTokenProviderMetadata(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.GetAdminTokenProviderMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdminTokenProviderMetadata`: OIDCProviderMetadata
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.GetAdminTokenProviderMetadata`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminTokenProviderMetadataRequest struct via the builder pattern


### Return type

[**OIDCProviderMetadata**](OIDCProviderMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthOidcScopes

> SupportedScopes GetAuthOidcScopes(ctx).ClientId(clientId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingaccess-go-client"
)

func main() {
    clientId := "clientId_example" // string | ID of a specific client to retrieve supported scopes for (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.GetAuthOidcScopes(context.Background()).ClientId(clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.GetAuthOidcScopes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthOidcScopes`: SupportedScopes
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.GetAuthOidcScopes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthOidcScopesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** | ID of a specific client to retrieve supported scopes for | 

### Return type

[**SupportedScopes**](SupportedScopes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBasicAuth

> BasicConfig GetBasicAuth(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingaccess-go-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.GetBasicAuth(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.GetBasicAuth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBasicAuth`: BasicConfig
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.GetBasicAuth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBasicAuthRequest struct via the builder pattern


### Return type

[**BasicConfig**](BasicConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOAuthAuth

> OAuthConfig GetOAuthAuth(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingaccess-go-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.GetOAuthAuth(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.GetOAuthAuth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOAuthAuth`: OAuthConfig
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.GetOAuthAuth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOAuthAuthRequest struct via the builder pattern


### Return type

[**OAuthConfig**](OAuthConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOidcAuth

> OidcConfig GetOidcAuth(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingaccess-go-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.GetOidcAuth(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.GetOidcAuth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOidcAuth`: OidcConfig
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.GetOidcAuth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOidcAuthRequest struct via the builder pattern


### Return type

[**OidcConfig**](OidcConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAdminBasicWebSession

> AdminBasicWebSession UpdateAdminBasicWebSession(ctx).AdminWebSession(adminWebSession).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingaccess-go-client"
)

func main() {
    adminWebSession := *openapiclient.NewAdminBasicWebSession(int32(123), int32(123), int32(123), int32(123), "Audience_example", openapiclient.WebSessionCookieType("Encrypted"), "TimeoutGroovyScript_example") // AdminBasicWebSession | admin web session configuration to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.UpdateAdminBasicWebSession(context.Background()).AdminWebSession(adminWebSession).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.UpdateAdminBasicWebSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAdminBasicWebSession`: AdminBasicWebSession
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.UpdateAdminBasicWebSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAdminBasicWebSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminWebSession** | [**AdminBasicWebSession**](AdminBasicWebSession.md) | admin web session configuration to update | 

### Return type

[**AdminBasicWebSession**](AdminBasicWebSession.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAdminTokenProvider

> AdminTokenProvider UpdateAdminTokenProvider(ctx).AdminTokenProvider(adminTokenProvider).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingaccess-go-client"
)

func main() {
    adminTokenProvider := *openapiclient.NewAdminTokenProvider("Issuer_example", []string{"SslProtocols_example"}, []string{"SslCiphers_example"}) // AdminTokenProvider | Admin Token Provider configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.UpdateAdminTokenProvider(context.Background()).AdminTokenProvider(adminTokenProvider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.UpdateAdminTokenProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAdminTokenProvider`: AdminTokenProvider
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.UpdateAdminTokenProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAdminTokenProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminTokenProvider** | [**AdminTokenProvider**](AdminTokenProvider.md) | Admin Token Provider configuration | 

### Return type

[**AdminTokenProvider**](AdminTokenProvider.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBasicAuth

> BasicAuthConfig UpdateBasicAuth(ctx).BasicConfig(basicConfig).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingaccess-go-client"
)

func main() {
    basicConfig := *openapiclient.NewBasicAuthConfig() // BasicAuthConfig | Basic configuration to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.UpdateBasicAuth(context.Background()).BasicConfig(basicConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.UpdateBasicAuth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBasicAuth`: BasicAuthConfig
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.UpdateBasicAuth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBasicAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **basicConfig** | [**BasicAuthConfig**](BasicAuthConfig.md) | Basic configuration to update | 

### Return type

[**BasicAuthConfig**](BasicAuthConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOAuthAuth

> OAuthConfig UpdateOAuthAuth(ctx).OAuthConfig(oAuthConfig).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingaccess-go-client"
)

func main() {
    oAuthConfig := *openapiclient.NewOAuthConfig("Scope_example") // OAuthConfig | OAuth configuration to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.UpdateOAuthAuth(context.Background()).OAuthConfig(oAuthConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.UpdateOAuthAuth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOAuthAuth`: OAuthConfig
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.UpdateOAuthAuth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOAuthAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oAuthConfig** | [**OAuthConfig**](OAuthConfig.md) | OAuth configuration to update | 

### Return type

[**OAuthConfig**](OAuthConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOidcAuth

> OidcConfig UpdateOidcAuth(ctx).OidcConfig(oidcConfig).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingaccess-go-client"
)

func main() {
    oidcConfig := *openapiclient.NewOidcConfig(*openapiclient.NewAdminWebSessionOidcConfiguration(*openapiclient.NewOAuthClientCredentials("ClientId_example"))) // OidcConfig | OIDC configuration to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.UpdateOidcAuth(context.Background()).OidcConfig(oidcConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.UpdateOidcAuth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOidcAuth`: OidcConfig
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.UpdateOidcAuth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOidcAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oidcConfig** | [**OidcConfig**](OidcConfig.md) | OIDC configuration to update | 

### Return type

[**OidcConfig**](OidcConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

