# \OidcApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOIDCProvider**](OidcApi.md#DeleteOIDCProvider) | **Delete** /oidc/provider | 
[**GetOIDCProvider**](OidcApi.md#GetOIDCProvider) | **Get** /oidc/provider | 
[**GetOIDCProviderMetadata**](OidcApi.md#GetOIDCProviderMetadata) | **Get** /oidc/provider/metadata | 
[**GetOIDCProviderPluginDescriptor**](OidcApi.md#GetOIDCProviderPluginDescriptor) | **Get** /oidc/provider/descriptors/{pluginType} | 
[**GetOIDCProviderPluginDescriptors**](OidcApi.md#GetOIDCProviderPluginDescriptors) | **Get** /oidc/provider/descriptors | 
[**UpdateOIDCProvider**](OidcApi.md#UpdateOIDCProvider) | **Put** /oidc/provider | 



## DeleteOIDCProvider

> DeleteOIDCProvider(ctx).Execute()





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
    r, err := apiClient.OidcApi.DeleteOIDCProvider(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OidcApi.DeleteOIDCProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOIDCProviderRequest struct via the builder pattern


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


## GetOIDCProvider

> OIDCProvider GetOIDCProvider(ctx).Execute()





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
    resp, r, err := apiClient.OidcApi.GetOIDCProvider(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OidcApi.GetOIDCProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOIDCProvider`: OIDCProvider
    fmt.Fprintf(os.Stdout, "Response from `OidcApi.GetOIDCProvider`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOIDCProviderRequest struct via the builder pattern


### Return type

[**OIDCProvider**](OIDCProvider.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOIDCProviderMetadata

> OIDCProviderMetadata GetOIDCProviderMetadata(ctx).Execute()





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
    resp, r, err := apiClient.OidcApi.GetOIDCProviderMetadata(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OidcApi.GetOIDCProviderMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOIDCProviderMetadata`: OIDCProviderMetadata
    fmt.Fprintf(os.Stdout, "Response from `OidcApi.GetOIDCProviderMetadata`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOIDCProviderMetadataRequest struct via the builder pattern


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


## GetOIDCProviderPluginDescriptor

> Descriptor GetOIDCProviderPluginDescriptor(ctx, pluginType).Execute()





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
    pluginType := "pluginType_example" // string | OIDC Provider plugin descriptor to retrieve

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OidcApi.GetOIDCProviderPluginDescriptor(context.Background(), pluginType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OidcApi.GetOIDCProviderPluginDescriptor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOIDCProviderPluginDescriptor`: Descriptor
    fmt.Fprintf(os.Stdout, "Response from `OidcApi.GetOIDCProviderPluginDescriptor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginType** | **string** | OIDC Provider plugin descriptor to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOIDCProviderPluginDescriptorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Descriptor**](Descriptor.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOIDCProviderPluginDescriptors

> Descriptors GetOIDCProviderPluginDescriptors(ctx).Execute()





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
    resp, r, err := apiClient.OidcApi.GetOIDCProviderPluginDescriptors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OidcApi.GetOIDCProviderPluginDescriptors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOIDCProviderPluginDescriptors`: Descriptors
    fmt.Fprintf(os.Stdout, "Response from `OidcApi.GetOIDCProviderPluginDescriptors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOIDCProviderPluginDescriptorsRequest struct via the builder pattern


### Return type

[**Descriptors**](Descriptors.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOIDCProvider

> OIDCProvider UpdateOIDCProvider(ctx).OpenIDConnectProvider(openIDConnectProvider).Execute()





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
    openIDConnectProvider := *openapiclient.NewOIDCProvider("Issuer_example") // OIDCProvider | OpenID Connect Provider configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OidcApi.UpdateOIDCProvider(context.Background()).OpenIDConnectProvider(openIDConnectProvider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OidcApi.UpdateOIDCProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOIDCProvider`: OIDCProvider
    fmt.Fprintf(os.Stdout, "Response from `OidcApi.UpdateOIDCProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOIDCProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **openIDConnectProvider** | [**OIDCProvider**](OIDCProvider.md) | OpenID Connect Provider configuration | 

### Return type

[**OIDCProvider**](OIDCProvider.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

