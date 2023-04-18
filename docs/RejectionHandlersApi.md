# \RejectionHandlersApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRejectionHandler**](RejectionHandlersApi.md#AddRejectionHandler) | **Post** /rejectionHandlers | 
[**DeleteRejectionHandler**](RejectionHandlersApi.md#DeleteRejectionHandler) | **Delete** /rejectionHandlers/{id} | 
[**GetRejecitonHandlerDescriptor**](RejectionHandlersApi.md#GetRejecitonHandlerDescriptor) | **Get** /rejectionHandlers/descriptors/{rejectionHandlerType} | 
[**GetRejectionHandler**](RejectionHandlersApi.md#GetRejectionHandler) | **Get** /rejectionHandlers/{id} | 
[**GetRejectionHandlerDescriptors**](RejectionHandlersApi.md#GetRejectionHandlerDescriptors) | **Get** /rejectionHandlers/descriptors | 
[**GetRejectionHandlers**](RejectionHandlersApi.md#GetRejectionHandlers) | **Get** /rejectionHandlers | 
[**UpdateRejectionHandler**](RejectionHandlersApi.md#UpdateRejectionHandler) | **Put** /rejectionHandlers/{id} | 



## AddRejectionHandler

> RejectionHandler AddRejectionHandler(ctx).RejectionHandler(rejectionHandler).Execute()





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
    rejectionHandler := *openapiclient.NewRejectionHandler("ClassName_example", "Name_example") // RejectionHandler | Rejection Handler to add

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RejectionHandlersApi.AddRejectionHandler(context.Background()).RejectionHandler(rejectionHandler).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RejectionHandlersApi.AddRejectionHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRejectionHandler`: RejectionHandler
    fmt.Fprintf(os.Stdout, "Response from `RejectionHandlersApi.AddRejectionHandler`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddRejectionHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rejectionHandler** | [**RejectionHandler**](RejectionHandler.md) | Rejection Handler to add | 

### Return type

[**RejectionHandler**](RejectionHandler.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRejectionHandler

> DeleteRejectionHandler(ctx, id).Execute()





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
    id := "id_example" // string | ID of Rejection Handler to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RejectionHandlersApi.DeleteRejectionHandler(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RejectionHandlersApi.DeleteRejectionHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Rejection Handler to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRejectionHandlerRequest struct via the builder pattern


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


## GetRejecitonHandlerDescriptor

> Descriptor GetRejecitonHandlerDescriptor(ctx, rejectionHandlerType).Execute()





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
    rejectionHandlerType := "rejectionHandlerType_example" // string | Rejection Handler descriptor to retrieve

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RejectionHandlersApi.GetRejecitonHandlerDescriptor(context.Background(), rejectionHandlerType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RejectionHandlersApi.GetRejecitonHandlerDescriptor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRejecitonHandlerDescriptor`: Descriptor
    fmt.Fprintf(os.Stdout, "Response from `RejectionHandlersApi.GetRejecitonHandlerDescriptor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rejectionHandlerType** | **string** | Rejection Handler descriptor to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRejecitonHandlerDescriptorRequest struct via the builder pattern


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


## GetRejectionHandler

> RejectionHandler GetRejectionHandler(ctx, id).Execute()





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
    id := "id_example" // string | ID of Rejection Handler to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RejectionHandlersApi.GetRejectionHandler(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RejectionHandlersApi.GetRejectionHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRejectionHandler`: RejectionHandler
    fmt.Fprintf(os.Stdout, "Response from `RejectionHandlersApi.GetRejectionHandler`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Rejection Handler to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRejectionHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RejectionHandler**](RejectionHandler.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRejectionHandlerDescriptors

> Descriptors GetRejectionHandlerDescriptors(ctx).Execute()





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
    resp, r, err := apiClient.RejectionHandlersApi.GetRejectionHandlerDescriptors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RejectionHandlersApi.GetRejectionHandlerDescriptors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRejectionHandlerDescriptors`: Descriptors
    fmt.Fprintf(os.Stdout, "Response from `RejectionHandlersApi.GetRejectionHandlerDescriptors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRejectionHandlerDescriptorsRequest struct via the builder pattern


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


## GetRejectionHandlers

> RejectionHandlers GetRejectionHandlers(ctx).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()





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
    numberPerPage := int64(56) // int64 | Number of Rejection Handlers per page (optional)
    filter := "filter_example" // string | Search for Rejection Handlers with name matching filter text (optional)
    name := "name_example" // string | Get a Rejection Handler by name (case-sensitive) (optional)
    sortKey := "sortKey_example" // string | A comma separated list of Rejection Handler attributes (keys) to be used to sort the results (optional)
    order := "order_example" // string | Order of the sorted results (ASC for ascending, DESC for descending) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RejectionHandlersApi.GetRejectionHandlers(context.Background()).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RejectionHandlersApi.GetRejectionHandlers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRejectionHandlers`: RejectionHandlers
    fmt.Fprintf(os.Stdout, "Response from `RejectionHandlersApi.GetRejectionHandlers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRejectionHandlersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | Page number to retrieve | 
 **numberPerPage** | **int64** | Number of Rejection Handlers per page | 
 **filter** | **string** | Search for Rejection Handlers with name matching filter text | 
 **name** | **string** | Get a Rejection Handler by name (case-sensitive) | 
 **sortKey** | **string** | A comma separated list of Rejection Handler attributes (keys) to be used to sort the results | 
 **order** | **string** | Order of the sorted results (ASC for ascending, DESC for descending) | 

### Return type

[**RejectionHandlers**](RejectionHandlers.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRejectionHandler

> RejectionHandler UpdateRejectionHandler(ctx, id).RejectionHandler(rejectionHandler).Execute()





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
    id := "id_example" // string | ID of Rejection Handler to update
    rejectionHandler := *openapiclient.NewRejectionHandler("ClassName_example", "Name_example") // RejectionHandler | Rejection Handler to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RejectionHandlersApi.UpdateRejectionHandler(context.Background(), id).RejectionHandler(rejectionHandler).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RejectionHandlersApi.UpdateRejectionHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRejectionHandler`: RejectionHandler
    fmt.Fprintf(os.Stdout, "Response from `RejectionHandlersApi.UpdateRejectionHandler`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Rejection Handler to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRejectionHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rejectionHandler** | [**RejectionHandler**](RejectionHandler.md) | Rejection Handler to update | 

### Return type

[**RejectionHandler**](RejectionHandler.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

