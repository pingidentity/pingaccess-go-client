# \HsmProvidersApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddHsmProvider**](HsmProvidersApi.md#AddHsmProvider) | **Post** /hsmProviders | 
[**DeleteHsmProvider**](HsmProvidersApi.md#DeleteHsmProvider) | **Delete** /hsmProviders/{id} | 
[**GetHsmProvider**](HsmProvidersApi.md#GetHsmProvider) | **Get** /hsmProviders/{id} | 
[**GetHsmProviderDescriptors**](HsmProvidersApi.md#GetHsmProviderDescriptors) | **Get** /hsmProviders/descriptors | 
[**GetHsmProviders**](HsmProvidersApi.md#GetHsmProviders) | **Get** /hsmProviders | 
[**UpdateHsmProvider**](HsmProvidersApi.md#UpdateHsmProvider) | **Put** /hsmProviders/{id} | 



## AddHsmProvider

> HsmProvider AddHsmProvider(ctx).HsmProvider(hsmProvider).Execute()





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
    hsmProvider := *openapiclient.NewHsmProvider("ClassName_example", "Name_example") // HsmProvider | HSM Provider to add

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HsmProvidersApi.AddHsmProvider(context.Background()).HsmProvider(hsmProvider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HsmProvidersApi.AddHsmProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddHsmProvider`: HsmProvider
    fmt.Fprintf(os.Stdout, "Response from `HsmProvidersApi.AddHsmProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddHsmProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hsmProvider** | [**HsmProvider**](HsmProvider.md) | HSM Provider to add | 

### Return type

[**HsmProvider**](HsmProvider.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHsmProvider

> DeleteHsmProvider(ctx, id).Execute()





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
    id := "id_example" // string | ID of HSM Provider to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HsmProvidersApi.DeleteHsmProvider(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HsmProvidersApi.DeleteHsmProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of HSM Provider to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHsmProviderRequest struct via the builder pattern


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


## GetHsmProvider

> HsmProvider GetHsmProvider(ctx, id).Execute()





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
    id := "id_example" // string | ID of HSM Provider to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HsmProvidersApi.GetHsmProvider(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HsmProvidersApi.GetHsmProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHsmProvider`: HsmProvider
    fmt.Fprintf(os.Stdout, "Response from `HsmProvidersApi.GetHsmProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of HSM Provider to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHsmProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HsmProvider**](HsmProvider.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHsmProviderDescriptors

> Descriptors GetHsmProviderDescriptors(ctx).Execute()





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
    resp, r, err := apiClient.HsmProvidersApi.GetHsmProviderDescriptors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HsmProvidersApi.GetHsmProviderDescriptors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHsmProviderDescriptors`: Descriptors
    fmt.Fprintf(os.Stdout, "Response from `HsmProvidersApi.GetHsmProviderDescriptors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHsmProviderDescriptorsRequest struct via the builder pattern


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


## GetHsmProviders

> HsmProvider GetHsmProviders(ctx).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()





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
    numberPerPage := int64(56) // int64 | Number of HSM Providers per page (optional)
    filter := "filter_example" // string | Search for HSM Provider with name matching filter text (optional)
    name := "name_example" // string | Get a specific HSM Provider by name (case-sensitive) (optional)
    sortKey := "sortKey_example" // string | A comma separated list of HSM Provider attributes (keys) to be used to sort the results (optional)
    order := "order_example" // string | Order of the sorted results (ASC for ascending, DESC for descending) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HsmProvidersApi.GetHsmProviders(context.Background()).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HsmProvidersApi.GetHsmProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHsmProviders`: HsmProvider
    fmt.Fprintf(os.Stdout, "Response from `HsmProvidersApi.GetHsmProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHsmProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | Page number to retrieve | 
 **numberPerPage** | **int64** | Number of HSM Providers per page | 
 **filter** | **string** | Search for HSM Provider with name matching filter text | 
 **name** | **string** | Get a specific HSM Provider by name (case-sensitive) | 
 **sortKey** | **string** | A comma separated list of HSM Provider attributes (keys) to be used to sort the results | 
 **order** | **string** | Order of the sorted results (ASC for ascending, DESC for descending) | 

### Return type

[**HsmProvider**](HsmProvider.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHsmProvider

> HsmProvider UpdateHsmProvider(ctx, id).HsmProvider(hsmProvider).Execute()





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
    id := "id_example" // string | ID of HSM Provider to update
    hsmProvider := *openapiclient.NewHsmProvider("ClassName_example", "Name_example") // HsmProvider | HSM Provider to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HsmProvidersApi.UpdateHsmProvider(context.Background(), id).HsmProvider(hsmProvider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HsmProvidersApi.UpdateHsmProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateHsmProvider`: HsmProvider
    fmt.Fprintf(os.Stdout, "Response from `HsmProvidersApi.UpdateHsmProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of HSM Provider to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHsmProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hsmProvider** | [**HsmProvider**](HsmProvider.md) | HSM Provider to update | 

### Return type

[**HsmProvider**](HsmProvider.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

