# \HttpsListenersApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHttpsListener**](HttpsListenersApi.md#GetHttpsListener) | **Get** /httpsListeners/{id} | 
[**GetHttpsListeners**](HttpsListenersApi.md#GetHttpsListeners) | **Get** /httpsListeners | 
[**UpdateHttpsListener**](HttpsListenersApi.md#UpdateHttpsListener) | **Put** /httpsListeners/{id} | 



## GetHttpsListener

> HttpsListener GetHttpsListener(ctx, id).Execute()





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
    id := "id_example" // string | ID of HttpsListener to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HttpsListenersApi.GetHttpsListener(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HttpsListenersApi.GetHttpsListener``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHttpsListener`: HttpsListener
    fmt.Fprintf(os.Stdout, "Response from `HttpsListenersApi.GetHttpsListener`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of HttpsListener to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHttpsListenerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HttpsListener**](HttpsListener.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHttpsListeners

> HttpsListeners GetHttpsListeners(ctx).SortKey(sortKey).Order(order).Execute()





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
    sortKey := "sortKey_example" // string | A comma separated list of HTTPS Listeners attributes (keys) to be used to sort the results (optional)
    order := "order_example" // string | Order of the sorted results (ASC for ascending, DESC for descending) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HttpsListenersApi.GetHttpsListeners(context.Background()).SortKey(sortKey).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HttpsListenersApi.GetHttpsListeners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHttpsListeners`: HttpsListeners
    fmt.Fprintf(os.Stdout, "Response from `HttpsListenersApi.GetHttpsListeners`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHttpsListenersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sortKey** | **string** | A comma separated list of HTTPS Listeners attributes (keys) to be used to sort the results | 
 **order** | **string** | Order of the sorted results (ASC for ascending, DESC for descending) | 

### Return type

[**HttpsListeners**](HttpsListeners.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHttpsListener

> HttpsListener UpdateHttpsListener(ctx, id).HttpsListeners(httpsListeners).Execute()





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
    id := "id_example" // string | ID of HttpsListener to update
    httpsListeners := *openapiclient.NewHttpsListener("Name_example", int64(123), false, false) // HttpsListener | HttpsListener to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HttpsListenersApi.UpdateHttpsListener(context.Background(), id).HttpsListeners(httpsListeners).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HttpsListenersApi.UpdateHttpsListener``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateHttpsListener`: HttpsListener
    fmt.Fprintf(os.Stdout, "Response from `HttpsListenersApi.UpdateHttpsListener`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of HttpsListener to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHttpsListenerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **httpsListeners** | [**HttpsListener**](HttpsListener.md) | HttpsListener to update | 

### Return type

[**HttpsListener**](HttpsListener.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

