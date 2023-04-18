# \AuthnReqListsApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAuthnReqList**](AuthnReqListsApi.md#AddAuthnReqList) | **Post** /authnReqLists | 
[**DeleteAuthnReqList**](AuthnReqListsApi.md#DeleteAuthnReqList) | **Delete** /authnReqLists/{id} | 
[**GetAuthnReqList**](AuthnReqListsApi.md#GetAuthnReqList) | **Get** /authnReqLists/{id} | 
[**GetAuthnReqLists**](AuthnReqListsApi.md#GetAuthnReqLists) | **Get** /authnReqLists | 
[**UpdateAuthnReqList**](AuthnReqListsApi.md#UpdateAuthnReqList) | **Put** /authnReqLists/{id} | 



## AddAuthnReqList

> AuthnReqList AddAuthnReqList(ctx).AuthnReqList(authnReqList).Execute()





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
    authnReqList := *openapiclient.NewAuthnReqList("Name_example", []string{"AuthnReqs_example"}) // AuthnReqList | Authentication Requirement List to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthnReqListsApi.AddAuthnReqList(context.Background()).AuthnReqList(authnReqList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthnReqListsApi.AddAuthnReqList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAuthnReqList`: AuthnReqList
    fmt.Fprintf(os.Stdout, "Response from `AuthnReqListsApi.AddAuthnReqList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAuthnReqListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authnReqList** | [**AuthnReqList**](AuthnReqList.md) | Authentication Requirement List to create | 

### Return type

[**AuthnReqList**](AuthnReqList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAuthnReqList

> DeleteAuthnReqList(ctx, id).Execute()





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
    id := "id_example" // string | ID of the Authentication Requirement List to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthnReqListsApi.DeleteAuthnReqList(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthnReqListsApi.DeleteAuthnReqList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Authentication Requirement List to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthnReqListRequest struct via the builder pattern


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


## GetAuthnReqList

> AuthnReqList GetAuthnReqList(ctx, id).Execute()





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
    id := "id_example" // string | ID of Authentication Requirement List to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthnReqListsApi.GetAuthnReqList(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthnReqListsApi.GetAuthnReqList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthnReqList`: AuthnReqList
    fmt.Fprintf(os.Stdout, "Response from `AuthnReqListsApi.GetAuthnReqList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Authentication Requirement List to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthnReqListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthnReqList**](AuthnReqList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthnReqLists

> AuthnReqLists GetAuthnReqLists(ctx).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()





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
    numberPerPage := int64(56) // int64 | Number of Authentication Requirements per page (optional)
    filter := "filter_example" // string | Search for Authentication Requirements with name matching filter text (optional)
    name := "name_example" // string | Get specific Authentication Requirements by name (case-sensitive) (optional)
    sortKey := "sortKey_example" // string | A comma separated list of Authentication Requirements attributes (keys) to be used to sort the results (optional)
    order := "order_example" // string | Order of the sorted results (ASC for ascending, DESC for descending) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthnReqListsApi.GetAuthnReqLists(context.Background()).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthnReqListsApi.GetAuthnReqLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthnReqLists`: AuthnReqLists
    fmt.Fprintf(os.Stdout, "Response from `AuthnReqListsApi.GetAuthnReqLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthnReqListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | Page number to retrieve | 
 **numberPerPage** | **int64** | Number of Authentication Requirements per page | 
 **filter** | **string** | Search for Authentication Requirements with name matching filter text | 
 **name** | **string** | Get specific Authentication Requirements by name (case-sensitive) | 
 **sortKey** | **string** | A comma separated list of Authentication Requirements attributes (keys) to be used to sort the results | 
 **order** | **string** | Order of the sorted results (ASC for ascending, DESC for descending) | 

### Return type

[**AuthnReqLists**](AuthnReqLists.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAuthnReqList

> AuthnReqList UpdateAuthnReqList(ctx, id).AuthnReqList(authnReqList).Execute()





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
    id := "id_example" // string | ID of Authentication Requirement List to update
    authnReqList := *openapiclient.NewAuthnReqList("Name_example", []string{"AuthnReqs_example"}) // AuthnReqList | Authentication Requirement List to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthnReqListsApi.UpdateAuthnReqList(context.Background(), id).AuthnReqList(authnReqList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthnReqListsApi.UpdateAuthnReqList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAuthnReqList`: AuthnReqList
    fmt.Fprintf(os.Stdout, "Response from `AuthnReqListsApi.UpdateAuthnReqList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Authentication Requirement List to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAuthnReqListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authnReqList** | [**AuthnReqList**](AuthnReqList.md) | Authentication Requirement List to update | 

### Return type

[**AuthnReqList**](AuthnReqList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

