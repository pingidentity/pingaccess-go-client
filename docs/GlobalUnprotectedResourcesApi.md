# \GlobalUnprotectedResourcesApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGlobalUnprotectedResource**](GlobalUnprotectedResourcesApi.md#AddGlobalUnprotectedResource) | **Post** /globalUnprotectedResources | 
[**DeleteGlobalUnprotectedResource**](GlobalUnprotectedResourcesApi.md#DeleteGlobalUnprotectedResource) | **Delete** /globalUnprotectedResources/{id} | 
[**GetGlobalUnprotectedResource**](GlobalUnprotectedResourcesApi.md#GetGlobalUnprotectedResource) | **Get** /globalUnprotectedResources/{id} | 
[**GetGlobalUnprotectedResources**](GlobalUnprotectedResourcesApi.md#GetGlobalUnprotectedResources) | **Get** /globalUnprotectedResources | 
[**UpdateGlobalUnprotectedResource**](GlobalUnprotectedResourcesApi.md#UpdateGlobalUnprotectedResource) | **Put** /globalUnprotectedResources/{id} | 



## AddGlobalUnprotectedResource

> GlobalUnprotectedResource AddGlobalUnprotectedResource(ctx).GlobalUnprotectedResource(globalUnprotectedResource).Execute()





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
    globalUnprotectedResource := *openapiclient.NewGlobalUnprotectedResource("Name_example", "WildcardPath_example") // GlobalUnprotectedResource | Global Unprotected Resource to add

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlobalUnprotectedResourcesApi.AddGlobalUnprotectedResource(context.Background()).GlobalUnprotectedResource(globalUnprotectedResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalUnprotectedResourcesApi.AddGlobalUnprotectedResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddGlobalUnprotectedResource`: GlobalUnprotectedResource
    fmt.Fprintf(os.Stdout, "Response from `GlobalUnprotectedResourcesApi.AddGlobalUnprotectedResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddGlobalUnprotectedResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **globalUnprotectedResource** | [**GlobalUnprotectedResource**](GlobalUnprotectedResource.md) | Global Unprotected Resource to add | 

### Return type

[**GlobalUnprotectedResource**](GlobalUnprotectedResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGlobalUnprotectedResource

> DeleteGlobalUnprotectedResource(ctx, id).Execute()





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
    id := "id_example" // string | ID of Global Unprotected Resource to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GlobalUnprotectedResourcesApi.DeleteGlobalUnprotectedResource(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalUnprotectedResourcesApi.DeleteGlobalUnprotectedResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Global Unprotected Resource to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGlobalUnprotectedResourceRequest struct via the builder pattern


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


## GetGlobalUnprotectedResource

> GlobalUnprotectedResource GetGlobalUnprotectedResource(ctx, id).Execute()





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
    id := "id_example" // string | ID of Global Unprotected Resource to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlobalUnprotectedResourcesApi.GetGlobalUnprotectedResource(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalUnprotectedResourcesApi.GetGlobalUnprotectedResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGlobalUnprotectedResource`: GlobalUnprotectedResource
    fmt.Fprintf(os.Stdout, "Response from `GlobalUnprotectedResourcesApi.GetGlobalUnprotectedResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Global Unprotected Resource to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalUnprotectedResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GlobalUnprotectedResource**](GlobalUnprotectedResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalUnprotectedResources

> GlobalUnprotectedResources GetGlobalUnprotectedResources(ctx).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()





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
    numberPerPage := int32(56) // int32 | Number of Global Unprotected Resources per page (optional)
    filter := "filter_example" // string | Search for Global Unprotected Resources with name matching filter text (optional)
    name := "name_example" // string | Get Global Unprotected Resources by name (optional)
    sortKey := "sortKey_example" // string | A comma separated list of Global Unprotected Resource attributes (keys) to be used to sort the results (optional)
    order := "order_example" // string | Order of the sorted results (ASC for ascending, DESC for descending) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlobalUnprotectedResourcesApi.GetGlobalUnprotectedResources(context.Background()).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalUnprotectedResourcesApi.GetGlobalUnprotectedResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGlobalUnprotectedResources`: GlobalUnprotectedResources
    fmt.Fprintf(os.Stdout, "Response from `GlobalUnprotectedResourcesApi.GetGlobalUnprotectedResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalUnprotectedResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number to retrieve | 
 **numberPerPage** | **int32** | Number of Global Unprotected Resources per page | 
 **filter** | **string** | Search for Global Unprotected Resources with name matching filter text | 
 **name** | **string** | Get Global Unprotected Resources by name | 
 **sortKey** | **string** | A comma separated list of Global Unprotected Resource attributes (keys) to be used to sort the results | 
 **order** | **string** | Order of the sorted results (ASC for ascending, DESC for descending) | 

### Return type

[**GlobalUnprotectedResources**](GlobalUnprotectedResources.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGlobalUnprotectedResource

> GlobalUnprotectedResource UpdateGlobalUnprotectedResource(ctx, id).GlobalUnprotectedResource(globalUnprotectedResource).Execute()





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
    id := "id_example" // string | ID of Global Unprotected Resource to update
    globalUnprotectedResource := *openapiclient.NewGlobalUnprotectedResource("Name_example", "WildcardPath_example") // GlobalUnprotectedResource | Global Unprotected Resource to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlobalUnprotectedResourcesApi.UpdateGlobalUnprotectedResource(context.Background(), id).GlobalUnprotectedResource(globalUnprotectedResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalUnprotectedResourcesApi.UpdateGlobalUnprotectedResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGlobalUnprotectedResource`: GlobalUnprotectedResource
    fmt.Fprintf(os.Stdout, "Response from `GlobalUnprotectedResourcesApi.UpdateGlobalUnprotectedResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Global Unprotected Resource to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGlobalUnprotectedResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **globalUnprotectedResource** | [**GlobalUnprotectedResource**](GlobalUnprotectedResource.md) | Global Unprotected Resource to update | 

### Return type

[**GlobalUnprotectedResource**](GlobalUnprotectedResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

