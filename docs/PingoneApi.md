# \PingoneApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPingOneConnection**](PingoneApi.md#AddPingOneConnection) | **Post** /pingone/connections | 
[**DeletePingOne4C**](PingoneApi.md#DeletePingOne4C) | **Delete** /pingone/customers | 
[**DeletePingOneConnection**](PingoneApi.md#DeletePingOneConnection) | **Delete** /pingone/connections/{id} | 
[**GetPingOne4C**](PingoneApi.md#GetPingOne4C) | **Get** /pingone/customers | 
[**GetPingOne4CMetadata**](PingoneApi.md#GetPingOne4CMetadata) | **Get** /pingone/customers/metadata | 
[**GetPingOneConnection**](PingoneApi.md#GetPingOneConnection) | **Get** /pingone/connections/{id} | 
[**GetPingOneConnectionMetadata**](PingoneApi.md#GetPingOneConnectionMetadata) | **Get** /pingone/connections/{id}/metadata | 
[**GetPingOneConnections**](PingoneApi.md#GetPingOneConnections) | **Get** /pingone/connections | 
[**UpdatePingOne4C**](PingoneApi.md#UpdatePingOne4C) | **Put** /pingone/customers | 
[**UpdatePingOneConnection**](PingoneApi.md#UpdatePingOneConnection) | **Put** /pingone/connections/{id} | 



## AddPingOneConnection

> PingOneConnection AddPingOneConnection(ctx).PingOneConnection(pingOneConnection).Execute()





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
    pingOneConnection := *openapiclient.NewPingOneConnection("Name_example", []string{"SslProtocols_example"}, []string{"SslCiphers_example"}) // PingOneConnection | PingOne Connection to add

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PingoneApi.AddPingOneConnection(context.Background()).PingOneConnection(pingOneConnection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingoneApi.AddPingOneConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPingOneConnection`: PingOneConnection
    fmt.Fprintf(os.Stdout, "Response from `PingoneApi.AddPingOneConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddPingOneConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pingOneConnection** | [**PingOneConnection**](PingOneConnection.md) | PingOne Connection to add | 

### Return type

[**PingOneConnection**](PingOneConnection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePingOne4C

> DeletePingOne4C(ctx).Execute()





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
    r, err := apiClient.PingoneApi.DeletePingOne4C(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingoneApi.DeletePingOne4C``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePingOne4CRequest struct via the builder pattern


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


## DeletePingOneConnection

> PingOneConnection DeletePingOneConnection(ctx, id).Execute()





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
    id := "id_example" // string | ID of PingOne Connection to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PingoneApi.DeletePingOneConnection(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingoneApi.DeletePingOneConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePingOneConnection`: PingOneConnection
    fmt.Fprintf(os.Stdout, "Response from `PingoneApi.DeletePingOneConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of PingOne Connection to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePingOneConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PingOneConnection**](PingOneConnection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPingOne4C

> PingOne4C GetPingOne4C(ctx).Execute()





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
    resp, r, err := apiClient.PingoneApi.GetPingOne4C(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingoneApi.GetPingOne4C``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPingOne4C`: PingOne4C
    fmt.Fprintf(os.Stdout, "Response from `PingoneApi.GetPingOne4C`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPingOne4CRequest struct via the builder pattern


### Return type

[**PingOne4C**](PingOne4C.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPingOne4CMetadata

> OIDCProviderMetadata GetPingOne4CMetadata(ctx).Execute()





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
    resp, r, err := apiClient.PingoneApi.GetPingOne4CMetadata(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingoneApi.GetPingOne4CMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPingOne4CMetadata`: OIDCProviderMetadata
    fmt.Fprintf(os.Stdout, "Response from `PingoneApi.GetPingOne4CMetadata`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPingOne4CMetadataRequest struct via the builder pattern


### Return type

[**OIDCProviderMetadata**](OIDCProviderMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPingOneConnection

> PingOneConnection GetPingOneConnection(ctx, id).Execute()





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
    id := "id_example" // string | ID of PingOne Connection to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PingoneApi.GetPingOneConnection(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingoneApi.GetPingOneConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPingOneConnection`: PingOneConnection
    fmt.Fprintf(os.Stdout, "Response from `PingoneApi.GetPingOneConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of PingOne Connection to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPingOneConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PingOneConnection**](PingOneConnection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPingOneConnectionMetadata

> PingOneConnectionMetadata GetPingOneConnectionMetadata(ctx, id).Execute()





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
    id := "id_example" // string | ID of the PingOne Connection to retrieve metadata for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PingoneApi.GetPingOneConnectionMetadata(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingoneApi.GetPingOneConnectionMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPingOneConnectionMetadata`: PingOneConnectionMetadata
    fmt.Fprintf(os.Stdout, "Response from `PingoneApi.GetPingOneConnectionMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the PingOne Connection to retrieve metadata for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPingOneConnectionMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PingOneConnectionMetadata**](PingOneConnectionMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPingOneConnections

> PingOneConnections GetPingOneConnections(ctx).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()





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
    numberPerPage := int64(56) // int64 | Number of PingOne Connections per page (optional)
    filter := "filter_example" // string | Search for PingOne Connections with name matching filter text (optional)
    name := "name_example" // string | Get a specific PingOne Connection by name (case-sensitive) (optional)
    sortKey := "sortKey_example" // string | A comma separated list of PingOne Connection attributes (keys) to be used to sort the results (optional)
    order := "order_example" // string | Order of the sorted results (ASC for ascending, DESC for descending) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PingoneApi.GetPingOneConnections(context.Background()).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingoneApi.GetPingOneConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPingOneConnections`: PingOneConnections
    fmt.Fprintf(os.Stdout, "Response from `PingoneApi.GetPingOneConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPingOneConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | Page number to retrieve | 
 **numberPerPage** | **int64** | Number of PingOne Connections per page | 
 **filter** | **string** | Search for PingOne Connections with name matching filter text | 
 **name** | **string** | Get a specific PingOne Connection by name (case-sensitive) | 
 **sortKey** | **string** | A comma separated list of PingOne Connection attributes (keys) to be used to sort the results | 
 **order** | **string** | Order of the sorted results (ASC for ascending, DESC for descending) | 

### Return type

[**PingOneConnections**](PingOneConnections.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePingOne4C

> PingOne4C UpdatePingOne4C(ctx).PingOneForCustomers(pingOneForCustomers).Execute()





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
    pingOneForCustomers := *openapiclient.NewPingOne4C("Issuer_example") // PingOne4C | PingOne For Customers configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PingoneApi.UpdatePingOne4C(context.Background()).PingOneForCustomers(pingOneForCustomers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingoneApi.UpdatePingOne4C``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePingOne4C`: PingOne4C
    fmt.Fprintf(os.Stdout, "Response from `PingoneApi.UpdatePingOne4C`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePingOne4CRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pingOneForCustomers** | [**PingOne4C**](PingOne4C.md) | PingOne For Customers configuration | 

### Return type

[**PingOne4C**](PingOne4C.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePingOneConnection

> PingOneConnection UpdatePingOneConnection(ctx, id).PingOneConnection(pingOneConnection).Execute()





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
    id := "id_example" // string | ID of the PingOne Connection to update
    pingOneConnection := *openapiclient.NewPingOneConnection("Name_example", []string{"SslProtocols_example"}, []string{"SslCiphers_example"}) // PingOneConnection | PingOne Connection to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PingoneApi.UpdatePingOneConnection(context.Background(), id).PingOneConnection(pingOneConnection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingoneApi.UpdatePingOneConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePingOneConnection`: PingOneConnection
    fmt.Fprintf(os.Stdout, "Response from `PingoneApi.UpdatePingOneConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the PingOne Connection to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePingOneConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pingOneConnection** | [**PingOneConnection**](PingOneConnection.md) | PingOne Connection to update | 

### Return type

[**PingOneConnection**](PingOneConnection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

