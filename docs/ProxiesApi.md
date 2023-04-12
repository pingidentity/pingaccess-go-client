# \ProxiesApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddProxy**](ProxiesApi.md#AddProxy) | **Post** /proxies | 
[**DeleteProxy**](ProxiesApi.md#DeleteProxy) | **Delete** /proxies/{id} | 
[**GetProxies**](ProxiesApi.md#GetProxies) | **Get** /proxies | 
[**GetProxy**](ProxiesApi.md#GetProxy) | **Get** /proxies/{id} | 
[**UpdateProxy**](ProxiesApi.md#UpdateProxy) | **Put** /proxies/{id} | 



## AddProxy

> HttpClientProxy AddProxy(ctx).Proxy(proxy).Execute()





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
    proxy := *openapiclient.NewHttpClientProxy("Name_example", "Host_example", int32(123)) // HttpClientProxy | Proxy to add

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProxiesApi.AddProxy(context.Background()).Proxy(proxy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProxiesApi.AddProxy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddProxy`: HttpClientProxy
    fmt.Fprintf(os.Stdout, "Response from `ProxiesApi.AddProxy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **proxy** | [**HttpClientProxy**](HttpClientProxy.md) | Proxy to add | 

### Return type

[**HttpClientProxy**](HttpClientProxy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProxy

> DeleteProxy(ctx, id).Execute()





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
    id := "id_example" // string | ID of Proxy to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProxiesApi.DeleteProxy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProxiesApi.DeleteProxy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Proxy to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProxyRequest struct via the builder pattern


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


## GetProxies

> HttpClientProxy GetProxies(ctx).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()





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
    page := int32(56) // int32 | Page number to retrieve (optional)
    numberPerPage := int32(56) // int32 | Number of Proxies per page (optional)
    filter := "filter_example" // string | Search for Proxies with name matching filter text (optional)
    name := "name_example" // string | Get a Proxy by name (case-sensitive) (optional)
    sortKey := "sortKey_example" // string | A comma separated list of Proxy attributes (keys) to be used to sort the results (optional)
    order := "order_example" // string | Order of the sorted results (ASC for ascending, DESC for descending) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProxiesApi.GetProxies(context.Background()).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProxiesApi.GetProxies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProxies`: HttpClientProxy
    fmt.Fprintf(os.Stdout, "Response from `ProxiesApi.GetProxies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProxiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number to retrieve | 
 **numberPerPage** | **int32** | Number of Proxies per page | 
 **filter** | **string** | Search for Proxies with name matching filter text | 
 **name** | **string** | Get a Proxy by name (case-sensitive) | 
 **sortKey** | **string** | A comma separated list of Proxy attributes (keys) to be used to sort the results | 
 **order** | **string** | Order of the sorted results (ASC for ascending, DESC for descending) | 

### Return type

[**HttpClientProxy**](HttpClientProxy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProxy

> HttpClientProxy GetProxy(ctx, id).Execute()





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
    id := "id_example" // string | ID of Proxy to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProxiesApi.GetProxy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProxiesApi.GetProxy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProxy`: HttpClientProxy
    fmt.Fprintf(os.Stdout, "Response from `ProxiesApi.GetProxy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Proxy to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HttpClientProxy**](HttpClientProxy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProxy

> HttpClientProxy UpdateProxy(ctx, id).Proxy(proxy).Execute()





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
    id := "id_example" // string | ID of Proxy to update
    proxy := *openapiclient.NewHttpClientProxy("Name_example", "Host_example", int32(123)) // HttpClientProxy | Proxy to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProxiesApi.UpdateProxy(context.Background(), id).Proxy(proxy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProxiesApi.UpdateProxy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProxy`: HttpClientProxy
    fmt.Fprintf(os.Stdout, "Response from `ProxiesApi.UpdateProxy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Proxy to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **proxy** | [**HttpClientProxy**](HttpClientProxy.md) | Proxy to update | 

### Return type

[**HttpClientProxy**](HttpClientProxy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

