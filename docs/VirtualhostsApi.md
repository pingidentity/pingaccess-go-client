# \VirtualhostsApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddVirtualHost**](VirtualhostsApi.md#AddVirtualHost) | **Post** /virtualhosts | 
[**DeleteVirtualHost**](VirtualhostsApi.md#DeleteVirtualHost) | **Delete** /virtualhosts/{id} | 
[**GetVirtualHost**](VirtualhostsApi.md#GetVirtualHost) | **Get** /virtualhosts/{id} | 
[**GetVirtualHosts**](VirtualhostsApi.md#GetVirtualHosts) | **Get** /virtualhosts | 
[**UpdateVirtualHost**](VirtualhostsApi.md#UpdateVirtualHost) | **Put** /virtualhosts/{id} | 



## AddVirtualHost

> VirtualHost AddVirtualHost(ctx).VirtualHost(virtualHost).Execute()





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
    virtualHost := *openapiclient.NewVirtualHost("Host_example", int64(123)) // VirtualHost | Virtual Host to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualhostsApi.AddVirtualHost(context.Background()).VirtualHost(virtualHost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualhostsApi.AddVirtualHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddVirtualHost`: VirtualHost
    fmt.Fprintf(os.Stdout, "Response from `VirtualhostsApi.AddVirtualHost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddVirtualHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **virtualHost** | [**VirtualHost**](VirtualHost.md) | Virtual Host to create | 

### Return type

[**VirtualHost**](VirtualHost.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVirtualHost

> DeleteVirtualHost(ctx, id).Execute()





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
    id := "id_example" // string | ID of Virtual Host to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.VirtualhostsApi.DeleteVirtualHost(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualhostsApi.DeleteVirtualHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Virtual Host to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVirtualHostRequest struct via the builder pattern


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


## GetVirtualHost

> VirtualHost GetVirtualHost(ctx, id).Execute()





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
    id := "id_example" // string | ID of Virtual Host to locate

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualhostsApi.GetVirtualHost(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualhostsApi.GetVirtualHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualHost`: VirtualHost
    fmt.Fprintf(os.Stdout, "Response from `VirtualhostsApi.GetVirtualHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Virtual Host to locate | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VirtualHost**](VirtualHost.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualHosts

> VirtualHosts GetVirtualHosts(ctx).Page(page).NumberPerPage(numberPerPage).Filter(filter).VirtualHost(virtualHost).SortKey(sortKey).Order(order).Execute()





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
    numberPerPage := int64(56) // int64 | Number of Virtual Hosts per page (optional)
    filter := "filter_example" // string | Search for Virtual Hosts with name matching filter text (optional)
    virtualHost := "virtualHost_example" // string | Get a Virtual Host by hostname, port or both (optional)
    sortKey := "sortKey_example" // string | A comma separated list of Virtual Host attributes (keys) to be used to sort the results (optional)
    order := "order_example" // string | Order of the sorted results (ASC for ascending, DESC for descending) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualhostsApi.GetVirtualHosts(context.Background()).Page(page).NumberPerPage(numberPerPage).Filter(filter).VirtualHost(virtualHost).SortKey(sortKey).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualhostsApi.GetVirtualHosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualHosts`: VirtualHosts
    fmt.Fprintf(os.Stdout, "Response from `VirtualhostsApi.GetVirtualHosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | Page number to retrieve | 
 **numberPerPage** | **int64** | Number of Virtual Hosts per page | 
 **filter** | **string** | Search for Virtual Hosts with name matching filter text | 
 **virtualHost** | **string** | Get a Virtual Host by hostname, port or both | 
 **sortKey** | **string** | A comma separated list of Virtual Host attributes (keys) to be used to sort the results | 
 **order** | **string** | Order of the sorted results (ASC for ascending, DESC for descending) | 

### Return type

[**VirtualHosts**](VirtualHosts.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVirtualHost

> VirtualHost UpdateVirtualHost(ctx, id).VirtualHost(virtualHost).Execute()





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
    id := "id_example" // string | ID of Virtual Host to update
    virtualHost := *openapiclient.NewVirtualHost("Host_example", int64(123)) // VirtualHost | Updated Virtual Host

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualhostsApi.UpdateVirtualHost(context.Background(), id).VirtualHost(virtualHost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualhostsApi.UpdateVirtualHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVirtualHost`: VirtualHost
    fmt.Fprintf(os.Stdout, "Response from `VirtualhostsApi.UpdateVirtualHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Virtual Host to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVirtualHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **virtualHost** | [**VirtualHost**](VirtualHost.md) | Updated Virtual Host | 

### Return type

[**VirtualHost**](VirtualHost.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

