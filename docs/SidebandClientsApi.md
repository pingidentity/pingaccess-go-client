# \SidebandClientsApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSidebandClient**](SidebandClientsApi.md#AddSidebandClient) | **Post** /sidebandClients | 
[**DeleteSidebandClient**](SidebandClientsApi.md#DeleteSidebandClient) | **Delete** /sidebandClients/{id} | 
[**GetSidebandClient**](SidebandClientsApi.md#GetSidebandClient) | **Get** /sidebandClients/{id} | 
[**GetSidebandClients**](SidebandClientsApi.md#GetSidebandClients) | **Get** /sidebandClients | 
[**UpdateSidebandClient**](SidebandClientsApi.md#UpdateSidebandClient) | **Put** /sidebandClients/{id} | 



## AddSidebandClient

> SidebandClient AddSidebandClient(ctx).SidebandClient(sidebandClient).Execute()





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
    sidebandClient := *openapiclient.NewSidebandClient("Name_example", []openapiclient.SidebandClientCredentials{*openapiclient.NewSidebandClientCredentials(*openapiclient.NewSidebandSharedSecretConfig())}) // SidebandClient | Sideband Client to add

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SidebandClientsApi.AddSidebandClient(context.Background()).SidebandClient(sidebandClient).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SidebandClientsApi.AddSidebandClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSidebandClient`: SidebandClient
    fmt.Fprintf(os.Stdout, "Response from `SidebandClientsApi.AddSidebandClient`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddSidebandClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sidebandClient** | [**SidebandClient**](SidebandClient.md) | Sideband Client to add | 

### Return type

[**SidebandClient**](SidebandClient.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSidebandClient

> DeleteSidebandClient(ctx, id).Execute()





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
    id := "id_example" // string | ID of Sideband Client to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SidebandClientsApi.DeleteSidebandClient(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SidebandClientsApi.DeleteSidebandClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Sideband Client to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSidebandClientRequest struct via the builder pattern


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


## GetSidebandClient

> SidebandClient GetSidebandClient(ctx, id).Execute()





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
    id := "id_example" // string | ID of Sideband client to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SidebandClientsApi.GetSidebandClient(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SidebandClientsApi.GetSidebandClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSidebandClient`: SidebandClient
    fmt.Fprintf(os.Stdout, "Response from `SidebandClientsApi.GetSidebandClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Sideband client to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSidebandClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SidebandClient**](SidebandClient.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSidebandClients

> SidebandClients GetSidebandClients(ctx).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()





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
    numberPerPage := int32(56) // int32 | Number of Sideband Clients per page (optional)
    filter := "filter_example" // string | Search for an Sideband Clients with name matching filter text (optional)
    name := "name_example" // string | Get a specific Sideband Clients by name (case-sensitive) (optional)
    sortKey := "sortKey_example" // string | A comma separated list of Sideband Client attributes (keys) to be used to sort the results (optional)
    order := "order_example" // string | Order of the sorted results (ASC for ascending, DESC for descending) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SidebandClientsApi.GetSidebandClients(context.Background()).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SidebandClientsApi.GetSidebandClients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSidebandClients`: SidebandClients
    fmt.Fprintf(os.Stdout, "Response from `SidebandClientsApi.GetSidebandClients`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSidebandClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number to retrieve | 
 **numberPerPage** | **int32** | Number of Sideband Clients per page | 
 **filter** | **string** | Search for an Sideband Clients with name matching filter text | 
 **name** | **string** | Get a specific Sideband Clients by name (case-sensitive) | 
 **sortKey** | **string** | A comma separated list of Sideband Client attributes (keys) to be used to sort the results | 
 **order** | **string** | Order of the sorted results (ASC for ascending, DESC for descending) | 

### Return type

[**SidebandClients**](SidebandClients.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSidebandClient

> SidebandClient UpdateSidebandClient(ctx, id).SidebandClient(sidebandClient).Execute()





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
    id := "id_example" // string | ID of the Sideband Client to update
    sidebandClient := *openapiclient.NewSidebandClient("Name_example", []openapiclient.SidebandClientCredentials{*openapiclient.NewSidebandClientCredentials(*openapiclient.NewSidebandSharedSecretConfig())}) // SidebandClient | Sideband Client to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SidebandClientsApi.UpdateSidebandClient(context.Background(), id).SidebandClient(sidebandClient).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SidebandClientsApi.UpdateSidebandClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSidebandClient`: SidebandClient
    fmt.Fprintf(os.Stdout, "Response from `SidebandClientsApi.UpdateSidebandClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Sideband Client to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSidebandClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sidebandClient** | [**SidebandClient**](SidebandClient.md) | Sideband Client to update | 

### Return type

[**SidebandClient**](SidebandClient.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

