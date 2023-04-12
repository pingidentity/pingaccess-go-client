# \PingfederateApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePingFederate**](PingfederateApi.md#DeletePingFederate) | **Delete** /pingfederate | 
[**DeletePingFederateAccessTokens**](PingfederateApi.md#DeletePingFederateAccessTokens) | **Delete** /pingfederate/accessTokens | 
[**DeletePingFederateAdmin**](PingfederateApi.md#DeletePingFederateAdmin) | **Delete** /pingfederate/admin | 
[**DeletePingFederateRuntime**](PingfederateApi.md#DeletePingFederateRuntime) | **Delete** /pingfederate/runtime | 
[**GetPingFederate**](PingfederateApi.md#GetPingFederate) | **Get** /pingfederate | 
[**GetPingFederateAccessTokens**](PingfederateApi.md#GetPingFederateAccessTokens) | **Get** /pingfederate/accessTokens | 
[**GetPingFederateAdmin**](PingfederateApi.md#GetPingFederateAdmin) | **Get** /pingfederate/admin | 
[**GetPingFederateMetadata**](PingfederateApi.md#GetPingFederateMetadata) | **Get** /pingfederate/metadata | 
[**GetPingFederateRuntime**](PingfederateApi.md#GetPingFederateRuntime) | **Get** /pingfederate/runtime | 
[**GetPingFederateRuntimeMetadata**](PingfederateApi.md#GetPingFederateRuntimeMetadata) | **Get** /pingfederate/runtime/metadata | 
[**UpdatePingFederate**](PingfederateApi.md#UpdatePingFederate) | **Put** /pingfederate | 
[**UpdatePingFederateAccessTokens**](PingfederateApi.md#UpdatePingFederateAccessTokens) | **Put** /pingfederate/accessTokens | 
[**UpdatePingFederateAdmin**](PingfederateApi.md#UpdatePingFederateAdmin) | **Put** /pingfederate/admin | 
[**UpdatePingFederateRuntime**](PingfederateApi.md#UpdatePingFederateRuntime) | **Put** /pingfederate/runtime | 



## DeletePingFederate

> DeletePingFederate(ctx).Execute()





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
    r, err := apiClient.PingfederateApi.DeletePingFederate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingfederateApi.DeletePingFederate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePingFederateRequest struct via the builder pattern


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


## DeletePingFederateAccessTokens

> DeletePingFederateAccessTokens(ctx).Execute()





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
    r, err := apiClient.PingfederateApi.DeletePingFederateAccessTokens(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingfederateApi.DeletePingFederateAccessTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePingFederateAccessTokensRequest struct via the builder pattern


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


## DeletePingFederateAdmin

> DeletePingFederateAdmin(ctx).Execute()





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
    r, err := apiClient.PingfederateApi.DeletePingFederateAdmin(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingfederateApi.DeletePingFederateAdmin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePingFederateAdminRequest struct via the builder pattern


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


## DeletePingFederateRuntime

> DeletePingFederateRuntime(ctx).Execute()





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
    r, err := apiClient.PingfederateApi.DeletePingFederateRuntime(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingfederateApi.DeletePingFederateRuntime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePingFederateRuntimeRequest struct via the builder pattern


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


## GetPingFederate

> PingFederateRuntime GetPingFederate(ctx).Execute()





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
    resp, r, err := apiClient.PingfederateApi.GetPingFederate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingfederateApi.GetPingFederate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPingFederate`: PingFederateRuntime
    fmt.Fprintf(os.Stdout, "Response from `PingfederateApi.GetPingFederate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPingFederateRequest struct via the builder pattern


### Return type

[**PingFederateRuntime**](PingFederateRuntime.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPingFederateAccessTokens

> PingFederateAccessToken GetPingFederateAccessTokens(ctx).Execute()





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
    resp, r, err := apiClient.PingfederateApi.GetPingFederateAccessTokens(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingfederateApi.GetPingFederateAccessTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPingFederateAccessTokens`: PingFederateAccessToken
    fmt.Fprintf(os.Stdout, "Response from `PingfederateApi.GetPingFederateAccessTokens`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPingFederateAccessTokensRequest struct via the builder pattern


### Return type

[**PingFederateAccessToken**](PingFederateAccessToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPingFederateAdmin

> PingFederateAdmin GetPingFederateAdmin(ctx).Execute()





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
    resp, r, err := apiClient.PingfederateApi.GetPingFederateAdmin(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingfederateApi.GetPingFederateAdmin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPingFederateAdmin`: PingFederateAdmin
    fmt.Fprintf(os.Stdout, "Response from `PingfederateApi.GetPingFederateAdmin`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPingFederateAdminRequest struct via the builder pattern


### Return type

[**PingFederateAdmin**](PingFederateAdmin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPingFederateMetadata

> OIDCProviderMetadata GetPingFederateMetadata(ctx).Execute()





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
    resp, r, err := apiClient.PingfederateApi.GetPingFederateMetadata(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingfederateApi.GetPingFederateMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPingFederateMetadata`: OIDCProviderMetadata
    fmt.Fprintf(os.Stdout, "Response from `PingfederateApi.GetPingFederateMetadata`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPingFederateMetadataRequest struct via the builder pattern


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


## GetPingFederateRuntime

> PingFederateMetadataRuntime GetPingFederateRuntime(ctx).Execute()





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
    resp, r, err := apiClient.PingfederateApi.GetPingFederateRuntime(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingfederateApi.GetPingFederateRuntime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPingFederateRuntime`: PingFederateMetadataRuntime
    fmt.Fprintf(os.Stdout, "Response from `PingfederateApi.GetPingFederateRuntime`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPingFederateRuntimeRequest struct via the builder pattern


### Return type

[**PingFederateMetadataRuntime**](PingFederateMetadataRuntime.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPingFederateRuntimeMetadata

> OIDCProviderMetadata GetPingFederateRuntimeMetadata(ctx).Execute()





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
    resp, r, err := apiClient.PingfederateApi.GetPingFederateRuntimeMetadata(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingfederateApi.GetPingFederateRuntimeMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPingFederateRuntimeMetadata`: OIDCProviderMetadata
    fmt.Fprintf(os.Stdout, "Response from `PingfederateApi.GetPingFederateRuntimeMetadata`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPingFederateRuntimeMetadataRequest struct via the builder pattern


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


## UpdatePingFederate

> PingFederateRuntime UpdatePingFederate(ctx).PingFederate(pingFederate).Execute()





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
    pingFederate := *openapiclient.NewPingFederateRuntime("Host_example", int32(123)) // PingFederateRuntime | PingFederate configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PingfederateApi.UpdatePingFederate(context.Background()).PingFederate(pingFederate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingfederateApi.UpdatePingFederate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePingFederate`: PingFederateRuntime
    fmt.Fprintf(os.Stdout, "Response from `PingfederateApi.UpdatePingFederate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePingFederateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pingFederate** | [**PingFederateRuntime**](PingFederateRuntime.md) | PingFederate configuration | 

### Return type

[**PingFederateRuntime**](PingFederateRuntime.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePingFederateAccessTokens

> PingFederateAccessToken UpdatePingFederateAccessTokens(ctx).PingFederateAccessToken(pingFederateAccessToken).Execute()





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
    pingFederateAccessToken := *openapiclient.NewPingFederateAccessToken("SubjectAttributeName_example", int32(123)) // PingFederateAccessToken | PingFederate OAuth Client configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PingfederateApi.UpdatePingFederateAccessTokens(context.Background()).PingFederateAccessToken(pingFederateAccessToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingfederateApi.UpdatePingFederateAccessTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePingFederateAccessTokens`: PingFederateAccessToken
    fmt.Fprintf(os.Stdout, "Response from `PingfederateApi.UpdatePingFederateAccessTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePingFederateAccessTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pingFederateAccessToken** | [**PingFederateAccessToken**](PingFederateAccessToken.md) | PingFederate OAuth Client configuration | 

### Return type

[**PingFederateAccessToken**](PingFederateAccessToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePingFederateAdmin

> PingFederateAdmin UpdatePingFederateAdmin(ctx).PingFederate(pingFederate).Execute()





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
    pingFederate := *openapiclient.NewPingFederateAdmin(*openapiclient.NewHiddenField(), "Host_example", int32(123)) // PingFederateAdmin | PingFederate Admin configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PingfederateApi.UpdatePingFederateAdmin(context.Background()).PingFederate(pingFederate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingfederateApi.UpdatePingFederateAdmin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePingFederateAdmin`: PingFederateAdmin
    fmt.Fprintf(os.Stdout, "Response from `PingfederateApi.UpdatePingFederateAdmin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePingFederateAdminRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pingFederate** | [**PingFederateAdmin**](PingFederateAdmin.md) | PingFederate Admin configuration | 

### Return type

[**PingFederateAdmin**](PingFederateAdmin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePingFederateRuntime

> PingFederateMetadataRuntime UpdatePingFederateRuntime(ctx).PingFederate(pingFederate).Execute()





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
    pingFederate := *openapiclient.NewPingFederateMetadataRuntime("Issuer_example") // PingFederateMetadataRuntime | PingFederate configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PingfederateApi.UpdatePingFederateRuntime(context.Background()).PingFederate(pingFederate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingfederateApi.UpdatePingFederateRuntime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePingFederateRuntime`: PingFederateMetadataRuntime
    fmt.Fprintf(os.Stdout, "Response from `PingfederateApi.UpdatePingFederateRuntime`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePingFederateRuntimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pingFederate** | [**PingFederateMetadataRuntime**](PingFederateMetadataRuntime.md) | PingFederate configuration | 

### Return type

[**PingFederateMetadataRuntime**](PingFederateMetadataRuntime.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

