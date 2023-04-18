# \EngineListenersApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddEngineListener**](EngineListenersApi.md#AddEngineListener) | **Post** /engineListeners | 
[**DeleteEngineListener**](EngineListenersApi.md#DeleteEngineListener) | **Delete** /engineListeners/{id} | 
[**GetEngineListener**](EngineListenersApi.md#GetEngineListener) | **Get** /engineListeners/{id} | 
[**GetEngineListeners**](EngineListenersApi.md#GetEngineListeners) | **Get** /engineListeners | 
[**UpdateEngineListener**](EngineListenersApi.md#UpdateEngineListener) | **Put** /engineListeners/{id} | 



## AddEngineListener

> EngineListener AddEngineListener(ctx).EngineListener(engineListener).Execute()





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
    engineListener := *openapiclient.NewEngineListener("Name_example", int64(123)) // EngineListener | Engine Listener to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EngineListenersApi.AddEngineListener(context.Background()).EngineListener(engineListener).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EngineListenersApi.AddEngineListener``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddEngineListener`: EngineListener
    fmt.Fprintf(os.Stdout, "Response from `EngineListenersApi.AddEngineListener`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddEngineListenerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **engineListener** | [**EngineListener**](EngineListener.md) | Engine Listener to create | 

### Return type

[**EngineListener**](EngineListener.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEngineListener

> DeleteEngineListener(ctx, id).Execute()





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
    id := "id_example" // string | ID of the Engine Listener to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EngineListenersApi.DeleteEngineListener(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EngineListenersApi.DeleteEngineListener``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Engine Listener to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEngineListenerRequest struct via the builder pattern


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


## GetEngineListener

> EngineListener GetEngineListener(ctx, id).Execute()





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
    id := "id_example" // string | ID of Engine Listener to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EngineListenersApi.GetEngineListener(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EngineListenersApi.GetEngineListener``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEngineListener`: EngineListener
    fmt.Fprintf(os.Stdout, "Response from `EngineListenersApi.GetEngineListener`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Engine Listener to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEngineListenerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EngineListener**](EngineListener.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEngineListeners

> EngineListeners GetEngineListeners(ctx).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()





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
    numberPerPage := int64(56) // int64 | Number of Engine Listeners per page (optional)
    filter := "filter_example" // string | Search for Engine Listeners with name matching filter text (optional)
    name := "name_example" // string | Get an Engine Listener by name (optional)
    sortKey := "sortKey_example" // string | A comma separated list of Engine Listener attributes (keys) to be used to sort the results (optional)
    order := "order_example" // string | Order of the sorted results (ASC for ascending, DESC for descending) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EngineListenersApi.GetEngineListeners(context.Background()).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EngineListenersApi.GetEngineListeners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEngineListeners`: EngineListeners
    fmt.Fprintf(os.Stdout, "Response from `EngineListenersApi.GetEngineListeners`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEngineListenersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | Page number to retrieve | 
 **numberPerPage** | **int64** | Number of Engine Listeners per page | 
 **filter** | **string** | Search for Engine Listeners with name matching filter text | 
 **name** | **string** | Get an Engine Listener by name | 
 **sortKey** | **string** | A comma separated list of Engine Listener attributes (keys) to be used to sort the results | 
 **order** | **string** | Order of the sorted results (ASC for ascending, DESC for descending) | 

### Return type

[**EngineListeners**](EngineListeners.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEngineListener

> EngineListener UpdateEngineListener(ctx, id).EngineListener(engineListener).Execute()





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
    id := "id_example" // string | ID of Engine Listener to update
    engineListener := *openapiclient.NewEngineListener("Name_example", int64(123)) // EngineListener | Engine Listener to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EngineListenersApi.UpdateEngineListener(context.Background(), id).EngineListener(engineListener).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EngineListenersApi.UpdateEngineListener``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEngineListener`: EngineListener
    fmt.Fprintf(os.Stdout, "Response from `EngineListenersApi.UpdateEngineListener`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Engine Listener to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEngineListenerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **engineListener** | [**EngineListener**](EngineListener.md) | Engine Listener to update | 

### Return type

[**EngineListener**](EngineListener.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

