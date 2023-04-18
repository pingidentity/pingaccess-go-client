# \EnginesApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddEngine**](EnginesApi.md#AddEngine) | **Post** /engines | 
[**DeleteEngine**](EnginesApi.md#DeleteEngine) | **Delete** /engines/{id} | 
[**GetEngine**](EnginesApi.md#GetEngine) | **Get** /engines/{id} | 
[**GetEngineCertificate**](EnginesApi.md#GetEngineCertificate) | **Get** /engines/certificates/{id} | 
[**GetEngineCertificates**](EnginesApi.md#GetEngineCertificates) | **Get** /engines/certificates | 
[**GetEngineConfigFile**](EnginesApi.md#GetEngineConfigFile) | **Post** /engines/{id}/config | 
[**GetEngineRegistrationJwt**](EnginesApi.md#GetEngineRegistrationJwt) | **Post** /engines/registration/token | 
[**GetEngineStatus**](EnginesApi.md#GetEngineStatus) | **Get** /engines/status | 
[**GetEngines**](EnginesApi.md#GetEngines) | **Get** /engines | 
[**UpdateEngine**](EnginesApi.md#UpdateEngine) | **Put** /engines/{id} | 



## AddEngine

> Engine AddEngine(ctx).Engine(engine).Execute()





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
    engine := *openapiclient.NewEngine("Name_example") // Engine | Engine to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnginesApi.AddEngine(context.Background()).Engine(engine).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnginesApi.AddEngine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddEngine`: Engine
    fmt.Fprintf(os.Stdout, "Response from `EnginesApi.AddEngine`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddEngineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **engine** | [**Engine**](Engine.md) | Engine to create | 

### Return type

[**Engine**](Engine.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEngine

> DeleteEngine(ctx, id).Execute()





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
    id := "id_example" // string | ID of the engine to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EnginesApi.DeleteEngine(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnginesApi.DeleteEngine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the engine to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEngineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetEngine

> Engine GetEngine(ctx, id).Execute()





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
    id := "id_example" // string | ID of engine to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnginesApi.GetEngine(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnginesApi.GetEngine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEngine`: Engine
    fmt.Fprintf(os.Stdout, "Response from `EnginesApi.GetEngine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of engine to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEngineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Engine**](Engine.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEngineCertificate

> EngineCertificate GetEngineCertificate(ctx, id).Execute()





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
    id := "id_example" // string | ID of Certificate to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnginesApi.GetEngineCertificate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnginesApi.GetEngineCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEngineCertificate`: EngineCertificate
    fmt.Fprintf(os.Stdout, "Response from `EnginesApi.GetEngineCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Certificate to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEngineCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EngineCertificate**](EngineCertificate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEngineCertificates

> EngineCertificates GetEngineCertificates(ctx).Page(page).NumberPerPage(numberPerPage).Filter(filter).Alias(alias).SortKey(sortKey).Order(order).Execute()





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
    page := int64(56) // int64 | Page number to retrieve (optional)
    numberPerPage := int64(56) // int64 | Number of Engine Certificates per page (optional)
    filter := "filter_example" // string | Search for Engine Certificates with alias matching filter text (optional)
    alias := "alias_example" // string | Get Engine Certificates by alias (optional)
    sortKey := "sortKey_example" // string | A comma separated list of Engine Certificate attributes (keys) to be used to sort the results (optional)
    order := "order_example" // string | Order of the sorted results (ASC for ascending, DESC for descending) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnginesApi.GetEngineCertificates(context.Background()).Page(page).NumberPerPage(numberPerPage).Filter(filter).Alias(alias).SortKey(sortKey).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnginesApi.GetEngineCertificates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEngineCertificates`: EngineCertificates
    fmt.Fprintf(os.Stdout, "Response from `EnginesApi.GetEngineCertificates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEngineCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | Page number to retrieve | 
 **numberPerPage** | **int64** | Number of Engine Certificates per page | 
 **filter** | **string** | Search for Engine Certificates with alias matching filter text | 
 **alias** | **string** | Get Engine Certificates by alias | 
 **sortKey** | **string** | A comma separated list of Engine Certificate attributes (keys) to be used to sort the results | 
 **order** | **string** | Order of the sorted results (ASC for ascending, DESC for descending) | 

### Return type

[**EngineCertificates**](EngineCertificates.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEngineConfigFile

> GetEngineConfigFile(ctx, id).Execute()





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
    id := "id_example" // string | ID of engine to get configuration for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EnginesApi.GetEngineConfigFile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnginesApi.GetEngineConfigFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of engine to get configuration for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEngineConfigFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetEngineRegistrationJwt

> GetEngineRegistrationJwt(ctx).EngineRegistrationToken(engineRegistrationToken).Execute()





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
    engineRegistrationToken := *openapiclient.NewEngineRegistrationToken(int64(123)) // EngineRegistrationToken | The required data to generate an engine self-registration token

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EnginesApi.GetEngineRegistrationJwt(context.Background()).EngineRegistrationToken(engineRegistrationToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnginesApi.GetEngineRegistrationJwt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEngineRegistrationJwtRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **engineRegistrationToken** | [**EngineRegistrationToken**](EngineRegistrationToken.md) | The required data to generate an engine self-registration token | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEngineStatus

> EngineHealthStatus GetEngineStatus(ctx).Execute()





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
    resp, r, err := apiClient.EnginesApi.GetEngineStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnginesApi.GetEngineStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEngineStatus`: EngineHealthStatus
    fmt.Fprintf(os.Stdout, "Response from `EnginesApi.GetEngineStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEngineStatusRequest struct via the builder pattern


### Return type

[**EngineHealthStatus**](EngineHealthStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEngines

> Engines GetEngines(ctx).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()





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
    page := int64(56) // int64 | Page number to retrieve (optional)
    numberPerPage := int64(56) // int64 | Number of Engines per page (optional)
    filter := "filter_example" // string | Search for Engines with name matching filter text (optional)
    name := "name_example" // string | Get a Engines by name (optional)
    sortKey := "sortKey_example" // string | A comma separated list of Engine attributes (keys) to be used to sort the results (optional)
    order := "order_example" // string | Order of the sorted results (ASC for ascending, DESC for descending) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnginesApi.GetEngines(context.Background()).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnginesApi.GetEngines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEngines`: Engines
    fmt.Fprintf(os.Stdout, "Response from `EnginesApi.GetEngines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEnginesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | Page number to retrieve | 
 **numberPerPage** | **int64** | Number of Engines per page | 
 **filter** | **string** | Search for Engines with name matching filter text | 
 **name** | **string** | Get a Engines by name | 
 **sortKey** | **string** | A comma separated list of Engine attributes (keys) to be used to sort the results | 
 **order** | **string** | Order of the sorted results (ASC for ascending, DESC for descending) | 

### Return type

[**Engines**](Engines.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEngine

> Engine UpdateEngine(ctx, id).Engine(engine).Execute()





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
    id := "id_example" // string | ID of engine to get
    engine := *openapiclient.NewEngine("Name_example") // Engine | Engine to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnginesApi.UpdateEngine(context.Background(), id).Engine(engine).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnginesApi.UpdateEngine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEngine`: Engine
    fmt.Fprintf(os.Stdout, "Response from `EnginesApi.UpdateEngine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of engine to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEngineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **engine** | [**Engine**](Engine.md) | Engine to update | 

### Return type

[**Engine**](Engine.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

