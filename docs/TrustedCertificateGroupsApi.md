# \TrustedCertificateGroupsApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTrustedCertificateGroup**](TrustedCertificateGroupsApi.md#AddTrustedCertificateGroup) | **Post** /trustedCertificateGroups | 
[**DeleteTrustedCertificateGroup**](TrustedCertificateGroupsApi.md#DeleteTrustedCertificateGroup) | **Delete** /trustedCertificateGroups/{id} | 
[**GetTrustedCertificateGroup**](TrustedCertificateGroupsApi.md#GetTrustedCertificateGroup) | **Get** /trustedCertificateGroups/{id} | 
[**GetTrustedCertificateGroups**](TrustedCertificateGroupsApi.md#GetTrustedCertificateGroups) | **Get** /trustedCertificateGroups | 
[**UpdateTrustedCertificateGroup**](TrustedCertificateGroupsApi.md#UpdateTrustedCertificateGroup) | **Put** /trustedCertificateGroups/{id} | 



## AddTrustedCertificateGroup

> TrustedCertificateGroup AddTrustedCertificateGroup(ctx).TrustedCertificateGroup(trustedCertificateGroup).Execute()





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
    trustedCertificateGroup := *openapiclient.NewTrustedCertificateGroup("Name_example") // TrustedCertificateGroup | Trusted Certificate Group to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustedCertificateGroupsApi.AddTrustedCertificateGroup(context.Background()).TrustedCertificateGroup(trustedCertificateGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedCertificateGroupsApi.AddTrustedCertificateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTrustedCertificateGroup`: TrustedCertificateGroup
    fmt.Fprintf(os.Stdout, "Response from `TrustedCertificateGroupsApi.AddTrustedCertificateGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddTrustedCertificateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trustedCertificateGroup** | [**TrustedCertificateGroup**](TrustedCertificateGroup.md) | Trusted Certificate Group to create | 

### Return type

[**TrustedCertificateGroup**](TrustedCertificateGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTrustedCertificateGroup

> DeleteTrustedCertificateGroup(ctx, id).Execute()





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
    id := "id_example" // string | ID of the Trusted Certificate Group to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TrustedCertificateGroupsApi.DeleteTrustedCertificateGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedCertificateGroupsApi.DeleteTrustedCertificateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Trusted Certificate Group to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTrustedCertificateGroupRequest struct via the builder pattern


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


## GetTrustedCertificateGroup

> TrustedCertificateGroup GetTrustedCertificateGroup(ctx, id).Execute()





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
    id := "id_example" // string | ID of TrustedCertificateGroup to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustedCertificateGroupsApi.GetTrustedCertificateGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedCertificateGroupsApi.GetTrustedCertificateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrustedCertificateGroup`: TrustedCertificateGroup
    fmt.Fprintf(os.Stdout, "Response from `TrustedCertificateGroupsApi.GetTrustedCertificateGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of TrustedCertificateGroup to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrustedCertificateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TrustedCertificateGroup**](TrustedCertificateGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrustedCertificateGroups

> TrustedCertificateGroups GetTrustedCertificateGroups(ctx).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()





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
    numberPerPage := int32(56) // int32 | Number of Trusted Certificate Groups per page (optional)
    filter := "filter_example" // string | Search for Trusted Certificate Groups with name matching filter text (optional)
    name := "name_example" // string | Get a Trusted Certificate Group by name (case-sensitive) (optional)
    sortKey := "sortKey_example" // string | A comma separated list of Trusted Certificate Group attributes (keys) to be used to sort the results (optional)
    order := "order_example" // string | Order of the sorted results (ASC for ascending, DESC for descending) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustedCertificateGroupsApi.GetTrustedCertificateGroups(context.Background()).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedCertificateGroupsApi.GetTrustedCertificateGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrustedCertificateGroups`: TrustedCertificateGroups
    fmt.Fprintf(os.Stdout, "Response from `TrustedCertificateGroupsApi.GetTrustedCertificateGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrustedCertificateGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number to retrieve | 
 **numberPerPage** | **int32** | Number of Trusted Certificate Groups per page | 
 **filter** | **string** | Search for Trusted Certificate Groups with name matching filter text | 
 **name** | **string** | Get a Trusted Certificate Group by name (case-sensitive) | 
 **sortKey** | **string** | A comma separated list of Trusted Certificate Group attributes (keys) to be used to sort the results | 
 **order** | **string** | Order of the sorted results (ASC for ascending, DESC for descending) | 

### Return type

[**TrustedCertificateGroups**](TrustedCertificateGroups.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTrustedCertificateGroup

> TrustedCertificateGroup UpdateTrustedCertificateGroup(ctx, id).TrustedCertificateGroup(trustedCertificateGroup).Execute()





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
    id := "id_example" // string | ID of TrustedCertificateGroup to update
    trustedCertificateGroup := *openapiclient.NewTrustedCertificateGroup("Name_example") // TrustedCertificateGroup | TrustedCertificateGroup to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustedCertificateGroupsApi.UpdateTrustedCertificateGroup(context.Background(), id).TrustedCertificateGroup(trustedCertificateGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedCertificateGroupsApi.UpdateTrustedCertificateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTrustedCertificateGroup`: TrustedCertificateGroup
    fmt.Fprintf(os.Stdout, "Response from `TrustedCertificateGroupsApi.UpdateTrustedCertificateGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of TrustedCertificateGroup to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTrustedCertificateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trustedCertificateGroup** | [**TrustedCertificateGroup**](TrustedCertificateGroup.md) | TrustedCertificateGroup to update | 

### Return type

[**TrustedCertificateGroup**](TrustedCertificateGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

