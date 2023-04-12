# \WebSessionsApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddWebSession**](WebSessionsApi.md#AddWebSession) | **Post** /webSessions | 
[**DeleteWebSession**](WebSessionsApi.md#DeleteWebSession) | **Delete** /webSessions/{id} | 
[**GetWebSession**](WebSessionsApi.md#GetWebSession) | **Get** /webSessions/{id} | 
[**GetWebSessions**](WebSessionsApi.md#GetWebSessions) | **Get** /webSessions | 
[**UpdateWebSession**](WebSessionsApi.md#UpdateWebSession) | **Put** /webSessions/{id} | 



## AddWebSession

> WebSession AddWebSession(ctx).WebSessions(webSessions).Execute()





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
    webSessions := *openapiclient.NewWebSession("Name_example", "Audience_example", *openapiclient.NewOAuthClientCredentials("ClientId_example"), false) // WebSession | WebSession to add

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebSessionsApi.AddWebSession(context.Background()).WebSessions(webSessions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebSessionsApi.AddWebSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddWebSession`: WebSession
    fmt.Fprintf(os.Stdout, "Response from `WebSessionsApi.AddWebSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddWebSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webSessions** | [**WebSession**](WebSession.md) | WebSession to add | 

### Return type

[**WebSession**](WebSession.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebSession

> DeleteWebSession(ctx, id).Execute()





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
    id := "id_example" // string | ID of WebSession to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WebSessionsApi.DeleteWebSession(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebSessionsApi.DeleteWebSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of WebSession to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebSessionRequest struct via the builder pattern


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


## GetWebSession

> WebSession GetWebSession(ctx, id).Execute()





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
    id := int32(56) // int32 | ID of WebSession to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebSessionsApi.GetWebSession(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebSessionsApi.GetWebSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebSession`: WebSession
    fmt.Fprintf(os.Stdout, "Response from `WebSessionsApi.GetWebSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of WebSession to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebSession**](WebSession.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebSessions

> WebSessions GetWebSessions(ctx).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()





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
    numberPerPage := int32(56) // int32 | Number of Web Sessions per page (optional)
    filter := "filter_example" // string | Search for Web Sessions with name matching filter text (optional)
    name := "name_example" // string | Get a Web Sessions by name (optional)
    sortKey := "sortKey_example" // string | A comma separated list of Web Session attributes (keys) to be used to sort the results (optional)
    order := "order_example" // string | Order of the sorted results (ASC for ascending, DESC for descending) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebSessionsApi.GetWebSessions(context.Background()).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebSessionsApi.GetWebSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebSessions`: WebSessions
    fmt.Fprintf(os.Stdout, "Response from `WebSessionsApi.GetWebSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWebSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number to retrieve | 
 **numberPerPage** | **int32** | Number of Web Sessions per page | 
 **filter** | **string** | Search for Web Sessions with name matching filter text | 
 **name** | **string** | Get a Web Sessions by name | 
 **sortKey** | **string** | A comma separated list of Web Session attributes (keys) to be used to sort the results | 
 **order** | **string** | Order of the sorted results (ASC for ascending, DESC for descending) | 

### Return type

[**WebSessions**](WebSessions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebSession

> WebSession UpdateWebSession(ctx, id).WebSessions(webSessions).Execute()





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
    id := "id_example" // string | ID of WebSession to update
    webSessions := *openapiclient.NewWebSession("Name_example", "Audience_example", *openapiclient.NewOAuthClientCredentials("ClientId_example"), false) // WebSession | WebSession to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebSessionsApi.UpdateWebSession(context.Background(), id).WebSessions(webSessions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebSessionsApi.UpdateWebSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWebSession`: WebSession
    fmt.Fprintf(os.Stdout, "Response from `WebSessionsApi.UpdateWebSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of WebSession to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webSessions** | [**WebSession**](WebSession.md) | WebSession to update | 

### Return type

[**WebSession**](WebSession.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

