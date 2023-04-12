# \AcmeApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAcmeAccount**](AcmeApi.md#AddAcmeAccount) | **Post** /acme/servers/{acmeServerId}/accounts | 
[**AddAcmeCertificateRequest**](AcmeApi.md#AddAcmeCertificateRequest) | **Post** /acme/servers/{acmeServerId}/accounts/{acmeAccountId}/certificateRequests | 
[**AddAcmeServer**](AcmeApi.md#AddAcmeServer) | **Post** /acme/servers | 
[**DeleteAcmeAccount**](AcmeApi.md#DeleteAcmeAccount) | **Delete** /acme/servers/{acmeServerId}/accounts/{acmeAccountId} | 
[**DeleteAcmeCertificateRequest**](AcmeApi.md#DeleteAcmeCertificateRequest) | **Delete** /acme/servers/{acmeServerId}/accounts/{acmeAccountId}/certificateRequests/{acmeCertificateRequestId} | 
[**DeleteAcmeServer**](AcmeApi.md#DeleteAcmeServer) | **Delete** /acme/servers/{acmeServerId} | 
[**GetAcmeAccount**](AcmeApi.md#GetAcmeAccount) | **Get** /acme/servers/{acmeServerId}/accounts/{acmeAccountId} | 
[**GetAcmeAccounts**](AcmeApi.md#GetAcmeAccounts) | **Get** /acme/servers/{acmeServerId}/accounts | 
[**GetAcmeCertificateRequest**](AcmeApi.md#GetAcmeCertificateRequest) | **Get** /acme/servers/{acmeServerId}/accounts/{acmeAccountId}/certificateRequests/{acmeCertificateRequestId} | 
[**GetAcmeCertificateRequests**](AcmeApi.md#GetAcmeCertificateRequests) | **Get** /acme/servers/{acmeServerId}/accounts/{acmeAccountId}/certificateRequests | 
[**GetAcmeServer**](AcmeApi.md#GetAcmeServer) | **Get** /acme/servers/{acmeServerId} | 
[**GetAcmeServers**](AcmeApi.md#GetAcmeServers) | **Get** /acme/servers | 
[**GetDefaultAcmeServer**](AcmeApi.md#GetDefaultAcmeServer) | **Get** /acme/servers/default | 
[**UpdateDefaultAcmeServer**](AcmeApi.md#UpdateDefaultAcmeServer) | **Put** /acme/servers/default | 



## AddAcmeAccount

> AcmeAccount AddAcmeAccount(ctx, acmeServerId).AcmeAccount(acmeAccount).Execute()





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
    acmeServerId := "acmeServerId_example" // string | ACME Server to add account to
    acmeAccount := *openapiclient.NewAcmeAccount() // AcmeAccount | ACME Account to add (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcmeApi.AddAcmeAccount(context.Background(), acmeServerId).AcmeAccount(acmeAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcmeApi.AddAcmeAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAcmeAccount`: AcmeAccount
    fmt.Fprintf(os.Stdout, "Response from `AcmeApi.AddAcmeAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**acmeServerId** | **string** | ACME Server to add account to | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAcmeAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acmeAccount** | [**AcmeAccount**](AcmeAccount.md) | ACME Account to add | 

### Return type

[**AcmeAccount**](AcmeAccount.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddAcmeCertificateRequest

> AcmeCertificateRequest AddAcmeCertificateRequest(ctx, acmeServerId, acmeAccountId).AcmeCertificateRequest(acmeCertificateRequest).Execute()





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
    acmeServerId := "acmeServerId_example" // string | ACME Server that the account belongs to
    acmeAccountId := "acmeAccountId_example" // string | ACME Account to add certificate to
    acmeCertificateRequest := *openapiclient.NewAcmeCertificateRequest(int32(123)) // AcmeCertificateRequest | ACME Certificate Request data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcmeApi.AddAcmeCertificateRequest(context.Background(), acmeServerId, acmeAccountId).AcmeCertificateRequest(acmeCertificateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcmeApi.AddAcmeCertificateRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAcmeCertificateRequest`: AcmeCertificateRequest
    fmt.Fprintf(os.Stdout, "Response from `AcmeApi.AddAcmeCertificateRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**acmeServerId** | **string** | ACME Server that the account belongs to | 
**acmeAccountId** | **string** | ACME Account to add certificate to | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAcmeCertificateRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **acmeCertificateRequest** | [**AcmeCertificateRequest**](AcmeCertificateRequest.md) | ACME Certificate Request data | 

### Return type

[**AcmeCertificateRequest**](AcmeCertificateRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddAcmeServer

> AcmeServer AddAcmeServer(ctx).AcmeServer(acmeServer).Execute()





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
    acmeServer := *openapiclient.NewAcmeServer("Name_example", "Url_example") // AcmeServer | ACME Server to add

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcmeApi.AddAcmeServer(context.Background()).AcmeServer(acmeServer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcmeApi.AddAcmeServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAcmeServer`: AcmeServer
    fmt.Fprintf(os.Stdout, "Response from `AcmeApi.AddAcmeServer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAcmeServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acmeServer** | [**AcmeServer**](AcmeServer.md) | ACME Server to add | 

### Return type

[**AcmeServer**](AcmeServer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAcmeAccount

> AcmeAccount DeleteAcmeAccount(ctx, acmeServerId, acmeAccountId).Execute()





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
    acmeServerId := "acmeServerId_example" // string | ID of ACME Server that holds account
    acmeAccountId := "acmeAccountId_example" // string | ID of ACME Account to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcmeApi.DeleteAcmeAccount(context.Background(), acmeServerId, acmeAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcmeApi.DeleteAcmeAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAcmeAccount`: AcmeAccount
    fmt.Fprintf(os.Stdout, "Response from `AcmeApi.DeleteAcmeAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**acmeServerId** | **string** | ID of ACME Server that holds account | 
**acmeAccountId** | **string** | ID of ACME Account to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAcmeAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AcmeAccount**](AcmeAccount.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAcmeCertificateRequest

> AcmeCertificateRequest DeleteAcmeCertificateRequest(ctx, acmeServerId, acmeAccountId, acmeCertificateRequestId).Execute()





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
    acmeServerId := "acmeServerId_example" // string | ID of ACME Server that holds account
    acmeAccountId := "acmeAccountId_example" // string | ID of an ACME Account
    acmeCertificateRequestId := "acmeCertificateRequestId_example" // string | ID of ACME Certificate Request to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcmeApi.DeleteAcmeCertificateRequest(context.Background(), acmeServerId, acmeAccountId, acmeCertificateRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcmeApi.DeleteAcmeCertificateRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAcmeCertificateRequest`: AcmeCertificateRequest
    fmt.Fprintf(os.Stdout, "Response from `AcmeApi.DeleteAcmeCertificateRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**acmeServerId** | **string** | ID of ACME Server that holds account | 
**acmeAccountId** | **string** | ID of an ACME Account | 
**acmeCertificateRequestId** | **string** | ID of ACME Certificate Request to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAcmeCertificateRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**AcmeCertificateRequest**](AcmeCertificateRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAcmeServer

> AcmeServer DeleteAcmeServer(ctx, acmeServerId).Execute()





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
    acmeServerId := "acmeServerId_example" // string | ID of ACME Server to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcmeApi.DeleteAcmeServer(context.Background(), acmeServerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcmeApi.DeleteAcmeServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAcmeServer`: AcmeServer
    fmt.Fprintf(os.Stdout, "Response from `AcmeApi.DeleteAcmeServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**acmeServerId** | **string** | ID of ACME Server to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAcmeServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AcmeServer**](AcmeServer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAcmeAccount

> AcmeAccount GetAcmeAccount(ctx, acmeServerId, acmeAccountId).Execute()





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
    acmeServerId := "acmeServerId_example" // string | ID of ACME Server that holds account
    acmeAccountId := "acmeAccountId_example" // string | ID of ACME Account to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcmeApi.GetAcmeAccount(context.Background(), acmeServerId, acmeAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcmeApi.GetAcmeAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAcmeAccount`: AcmeAccount
    fmt.Fprintf(os.Stdout, "Response from `AcmeApi.GetAcmeAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**acmeServerId** | **string** | ID of ACME Server that holds account | 
**acmeAccountId** | **string** | ID of ACME Account to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAcmeAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AcmeAccount**](AcmeAccount.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAcmeAccounts

> AcmeAccount GetAcmeAccounts(ctx, acmeServerId).Page(page).NumberPerPage(numberPerPage).SortKey(sortKey).Order(order).Execute()





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
    acmeServerId := "acmeServerId_example" // string | ID of ACME Server that holds accounts
    page := int32(56) // int32 | Page number to retrieve (optional)
    numberPerPage := int32(56) // int32 | Number of ACME Accounts per page (optional)
    sortKey := "sortKey_example" // string | A comma separated list of ACME Account attributes (keys) to be used to sort the results (optional)
    order := "order_example" // string | Order of the sorted results (ASC for ascending, DESC for descending) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcmeApi.GetAcmeAccounts(context.Background(), acmeServerId).Page(page).NumberPerPage(numberPerPage).SortKey(sortKey).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcmeApi.GetAcmeAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAcmeAccounts`: AcmeAccount
    fmt.Fprintf(os.Stdout, "Response from `AcmeApi.GetAcmeAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**acmeServerId** | **string** | ID of ACME Server that holds accounts | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAcmeAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number to retrieve | 
 **numberPerPage** | **int32** | Number of ACME Accounts per page | 
 **sortKey** | **string** | A comma separated list of ACME Account attributes (keys) to be used to sort the results | 
 **order** | **string** | Order of the sorted results (ASC for ascending, DESC for descending) | 

### Return type

[**AcmeAccount**](AcmeAccount.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAcmeCertificateRequest

> AcmeCertificateRequest GetAcmeCertificateRequest(ctx, acmeServerId, acmeAccountId, acmeCertificateRequestId).Execute()





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
    acmeServerId := "acmeServerId_example" // string | ID of ACME Server that holds account
    acmeAccountId := "acmeAccountId_example" // string | ID of the ACME Account
    acmeCertificateRequestId := "acmeCertificateRequestId_example" // string | ID of ACME Certificate Request to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcmeApi.GetAcmeCertificateRequest(context.Background(), acmeServerId, acmeAccountId, acmeCertificateRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcmeApi.GetAcmeCertificateRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAcmeCertificateRequest`: AcmeCertificateRequest
    fmt.Fprintf(os.Stdout, "Response from `AcmeApi.GetAcmeCertificateRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**acmeServerId** | **string** | ID of ACME Server that holds account | 
**acmeAccountId** | **string** | ID of the ACME Account | 
**acmeCertificateRequestId** | **string** | ID of ACME Certificate Request to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAcmeCertificateRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**AcmeCertificateRequest**](AcmeCertificateRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAcmeCertificateRequests

> AcmeCertificateRequest GetAcmeCertificateRequests(ctx, acmeServerId, acmeAccountId).KeyPairId(keyPairId).Page(page).NumberPerPage(numberPerPage).SortKey(sortKey).Order(order).Execute()





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
    acmeServerId := "acmeServerId_example" // string | ID of the ACME Server
    acmeAccountId := "acmeAccountId_example" // string | ID of the ACME Account
    keyPairId := "keyPairId_example" // string | ID of the Key Pair (optional)
    page := int32(56) // int32 | Page number to retrieve (optional)
    numberPerPage := int32(56) // int32 | Number of ACME Certificate Requests per page (optional)
    sortKey := "sortKey_example" // string | A comma separated list of ACME Certificate Request attributes (keys) to be used to sort the results (optional)
    order := "order_example" // string | Order of the sorted results (ASC for ascending, DESC for descending) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcmeApi.GetAcmeCertificateRequests(context.Background(), acmeServerId, acmeAccountId).KeyPairId(keyPairId).Page(page).NumberPerPage(numberPerPage).SortKey(sortKey).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcmeApi.GetAcmeCertificateRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAcmeCertificateRequests`: AcmeCertificateRequest
    fmt.Fprintf(os.Stdout, "Response from `AcmeApi.GetAcmeCertificateRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**acmeServerId** | **string** | ID of the ACME Server | 
**acmeAccountId** | **string** | ID of the ACME Account | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAcmeCertificateRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **keyPairId** | **string** | ID of the Key Pair | 
 **page** | **int32** | Page number to retrieve | 
 **numberPerPage** | **int32** | Number of ACME Certificate Requests per page | 
 **sortKey** | **string** | A comma separated list of ACME Certificate Request attributes (keys) to be used to sort the results | 
 **order** | **string** | Order of the sorted results (ASC for ascending, DESC for descending) | 

### Return type

[**AcmeCertificateRequest**](AcmeCertificateRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAcmeServer

> AcmeServer GetAcmeServer(ctx, acmeServerId).Execute()





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
    acmeServerId := "acmeServerId_example" // string | ID of ACME Server to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcmeApi.GetAcmeServer(context.Background(), acmeServerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcmeApi.GetAcmeServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAcmeServer`: AcmeServer
    fmt.Fprintf(os.Stdout, "Response from `AcmeApi.GetAcmeServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**acmeServerId** | **string** | ID of ACME Server to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAcmeServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AcmeServer**](AcmeServer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAcmeServers

> AcmeServers GetAcmeServers(ctx).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()





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
    numberPerPage := int32(56) // int32 | Number of ACME Servers per page (optional)
    filter := "filter_example" // string | Search for ACME Server with name matching filter text (optional)
    name := "name_example" // string | Get a specific ACME Server by name (case-sensitive) (optional)
    sortKey := "sortKey_example" // string | A comma separated list of ACME Server attributes (keys) to be used to sort the results (optional)
    order := "order_example" // string | Order of the sorted results (ASC for ascending, DESC for descending) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcmeApi.GetAcmeServers(context.Background()).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcmeApi.GetAcmeServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAcmeServers`: AcmeServers
    fmt.Fprintf(os.Stdout, "Response from `AcmeApi.GetAcmeServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAcmeServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number to retrieve | 
 **numberPerPage** | **int32** | Number of ACME Servers per page | 
 **filter** | **string** | Search for ACME Server with name matching filter text | 
 **name** | **string** | Get a specific ACME Server by name (case-sensitive) | 
 **sortKey** | **string** | A comma separated list of ACME Server attributes (keys) to be used to sort the results | 
 **order** | **string** | Order of the sorted results (ASC for ascending, DESC for descending) | 

### Return type

[**AcmeServers**](AcmeServers.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultAcmeServer

> Link GetDefaultAcmeServer(ctx).Execute()





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
    resp, r, err := apiClient.AcmeApi.GetDefaultAcmeServer(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcmeApi.GetDefaultAcmeServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefaultAcmeServer`: Link
    fmt.Fprintf(os.Stdout, "Response from `AcmeApi.GetDefaultAcmeServer`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultAcmeServerRequest struct via the builder pattern


### Return type

[**Link**](Link.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDefaultAcmeServer

> Link UpdateDefaultAcmeServer(ctx).AcmeServer(acmeServer).Execute()





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
    acmeServer := *openapiclient.NewLink() // Link | ACME Server to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcmeApi.UpdateDefaultAcmeServer(context.Background()).AcmeServer(acmeServer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcmeApi.UpdateDefaultAcmeServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDefaultAcmeServer`: Link
    fmt.Fprintf(os.Stdout, "Response from `AcmeApi.UpdateDefaultAcmeServer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDefaultAcmeServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acmeServer** | [**Link**](Link.md) | ACME Server to update | 

### Return type

[**Link**](Link.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

