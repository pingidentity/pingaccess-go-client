# \RedirectsApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRedirect**](RedirectsApi.md#AddRedirect) | **Post** /redirects | 
[**DeleteRedirect**](RedirectsApi.md#DeleteRedirect) | **Delete** /redirects/{id} | 
[**GetRedirect**](RedirectsApi.md#GetRedirect) | **Get** /redirects/{id} | 
[**GetRedirects**](RedirectsApi.md#GetRedirects) | **Get** /redirects | 
[**UpdateRedirect**](RedirectsApi.md#UpdateRedirect) | **Put** /redirects/{id} | 



## AddRedirect

> Redirect AddRedirect(ctx).Redirects(redirects).Execute()





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
    redirects := *openapiclient.NewRedirect() // Redirect | Redirect to add

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RedirectsApi.AddRedirect(context.Background()).Redirects(redirects).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RedirectsApi.AddRedirect``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRedirect`: Redirect
    fmt.Fprintf(os.Stdout, "Response from `RedirectsApi.AddRedirect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddRedirectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **redirects** | [**Redirect**](Redirect.md) | Redirect to add | 

### Return type

[**Redirect**](Redirect.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRedirect

> DeleteRedirect(ctx, id).Execute()





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
    id := "id_example" // string | ID of Redirect to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RedirectsApi.DeleteRedirect(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RedirectsApi.DeleteRedirect``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Redirect to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRedirectRequest struct via the builder pattern


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


## GetRedirect

> Redirect GetRedirect(ctx, id).Execute()





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
    id := "id_example" // string | ID of Redirect to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RedirectsApi.GetRedirect(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RedirectsApi.GetRedirect``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRedirect`: Redirect
    fmt.Fprintf(os.Stdout, "Response from `RedirectsApi.GetRedirect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Redirect to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedirectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Redirect**](Redirect.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedirects

> Redirects GetRedirects(ctx).Page(page).NumberPerPage(numberPerPage).Filter(filter).Source(source).Target(target).SortKey(sortKey).Order(order).Execute()





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
    numberPerPage := int64(56) // int64 | Number of Redirects per page (optional)
    filter := "filter_example" // string | Search for Redirects by source with host, port or both matching filter text (optional)
    source := "source_example" // string | Search for a Redirect by a source host, port or both (optional)
    target := "target_example" // string | Search for a Redirect by a target host, port or both (optional)
    sortKey := "sortKey_example" // string | A comma separated list of Redirect attributes (keys) to be used to sort the results (optional)
    order := "order_example" // string | Order of the sorted results (ASC for ascending, DESC for descending) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RedirectsApi.GetRedirects(context.Background()).Page(page).NumberPerPage(numberPerPage).Filter(filter).Source(source).Target(target).SortKey(sortKey).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RedirectsApi.GetRedirects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRedirects`: Redirects
    fmt.Fprintf(os.Stdout, "Response from `RedirectsApi.GetRedirects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRedirectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | Page number to retrieve | 
 **numberPerPage** | **int64** | Number of Redirects per page | 
 **filter** | **string** | Search for Redirects by source with host, port or both matching filter text | 
 **source** | **string** | Search for a Redirect by a source host, port or both | 
 **target** | **string** | Search for a Redirect by a target host, port or both | 
 **sortKey** | **string** | A comma separated list of Redirect attributes (keys) to be used to sort the results | 
 **order** | **string** | Order of the sorted results (ASC for ascending, DESC for descending) | 

### Return type

[**Redirects**](Redirects.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRedirect

> Redirect UpdateRedirect(ctx, id).Redirects(redirects).Execute()





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
    id := "id_example" // string | ID of Redirect to update
    redirects := *openapiclient.NewRedirect() // Redirect | Redirect to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RedirectsApi.UpdateRedirect(context.Background(), id).Redirects(redirects).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RedirectsApi.UpdateRedirect``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRedirect`: Redirect
    fmt.Fprintf(os.Stdout, "Response from `RedirectsApi.UpdateRedirect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Redirect to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRedirectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **redirects** | [**Redirect**](Redirect.md) | Redirect to update | 

### Return type

[**Redirect**](Redirect.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

