# \SiteAuthenticatorsApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSiteAuthenticator**](SiteAuthenticatorsApi.md#AddSiteAuthenticator) | **Post** /siteAuthenticators | 
[**DeleteSiteAuthenticator**](SiteAuthenticatorsApi.md#DeleteSiteAuthenticator) | **Delete** /siteAuthenticators/{id} | 
[**GetSiteAuthenticator**](SiteAuthenticatorsApi.md#GetSiteAuthenticator) | **Get** /siteAuthenticators/{id} | 
[**GetSiteAuthenticatorDescriptor**](SiteAuthenticatorsApi.md#GetSiteAuthenticatorDescriptor) | **Get** /siteAuthenticators/descriptors/{siteAuthenticatorType} | 
[**GetSiteAuthenticatorDescriptors**](SiteAuthenticatorsApi.md#GetSiteAuthenticatorDescriptors) | **Get** /siteAuthenticators/descriptors | 
[**GetSiteAuthenticators**](SiteAuthenticatorsApi.md#GetSiteAuthenticators) | **Get** /siteAuthenticators | 
[**UpdateSiteAuthenticator**](SiteAuthenticatorsApi.md#UpdateSiteAuthenticator) | **Put** /siteAuthenticators/{id} | 



## AddSiteAuthenticator

> SiteAuthenticator AddSiteAuthenticator(ctx).SiteAuthenticator(siteAuthenticator).Execute()





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
    siteAuthenticator := *openapiclient.NewSiteAuthenticator("ClassName_example", "Name_example") // SiteAuthenticator | Site Authenticator to add

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SiteAuthenticatorsApi.AddSiteAuthenticator(context.Background()).SiteAuthenticator(siteAuthenticator).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SiteAuthenticatorsApi.AddSiteAuthenticator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSiteAuthenticator`: SiteAuthenticator
    fmt.Fprintf(os.Stdout, "Response from `SiteAuthenticatorsApi.AddSiteAuthenticator`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddSiteAuthenticatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **siteAuthenticator** | [**SiteAuthenticator**](SiteAuthenticator.md) | Site Authenticator to add | 

### Return type

[**SiteAuthenticator**](SiteAuthenticator.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSiteAuthenticator

> DeleteSiteAuthenticator(ctx, id).Execute()





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
    id := "id_example" // string | ID of Site Authenticator to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SiteAuthenticatorsApi.DeleteSiteAuthenticator(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SiteAuthenticatorsApi.DeleteSiteAuthenticator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Site Authenticator to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSiteAuthenticatorRequest struct via the builder pattern


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


## GetSiteAuthenticator

> SiteAuthenticator GetSiteAuthenticator(ctx, id).Execute()





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
    id := "id_example" // string | ID of Site Authenticator to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SiteAuthenticatorsApi.GetSiteAuthenticator(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SiteAuthenticatorsApi.GetSiteAuthenticator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSiteAuthenticator`: SiteAuthenticator
    fmt.Fprintf(os.Stdout, "Response from `SiteAuthenticatorsApi.GetSiteAuthenticator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Site Authenticator to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteAuthenticatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SiteAuthenticator**](SiteAuthenticator.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteAuthenticatorDescriptor

> Descriptor GetSiteAuthenticatorDescriptor(ctx, siteAuthenticatorType).Execute()





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
    siteAuthenticatorType := "siteAuthenticatorType_example" // string | Site Authenticator descriptor to retrieve

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SiteAuthenticatorsApi.GetSiteAuthenticatorDescriptor(context.Background(), siteAuthenticatorType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SiteAuthenticatorsApi.GetSiteAuthenticatorDescriptor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSiteAuthenticatorDescriptor`: Descriptor
    fmt.Fprintf(os.Stdout, "Response from `SiteAuthenticatorsApi.GetSiteAuthenticatorDescriptor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteAuthenticatorType** | **string** | Site Authenticator descriptor to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteAuthenticatorDescriptorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Descriptor**](Descriptor.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteAuthenticatorDescriptors

> Descriptors GetSiteAuthenticatorDescriptors(ctx).Execute()





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
    resp, r, err := apiClient.SiteAuthenticatorsApi.GetSiteAuthenticatorDescriptors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SiteAuthenticatorsApi.GetSiteAuthenticatorDescriptors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSiteAuthenticatorDescriptors`: Descriptors
    fmt.Fprintf(os.Stdout, "Response from `SiteAuthenticatorsApi.GetSiteAuthenticatorDescriptors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteAuthenticatorDescriptorsRequest struct via the builder pattern


### Return type

[**Descriptors**](Descriptors.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteAuthenticators

> SiteAuthenticators GetSiteAuthenticators(ctx).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()





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
    numberPerPage := int32(56) // int32 | Number of Site Authenticators per page (optional)
    filter := "filter_example" // string | Search for Site Authenticators with name matching filter text (optional)
    name := "name_example" // string | Get a Site Authenticator by name (case-sensitive) (optional)
    sortKey := "sortKey_example" // string | A comma separated list of Site Authenticator attributes (keys) to be used to sort the results (optional)
    order := "order_example" // string | Order of the sorted results (ASC for ascending, DESC for descending) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SiteAuthenticatorsApi.GetSiteAuthenticators(context.Background()).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SiteAuthenticatorsApi.GetSiteAuthenticators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSiteAuthenticators`: SiteAuthenticators
    fmt.Fprintf(os.Stdout, "Response from `SiteAuthenticatorsApi.GetSiteAuthenticators`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteAuthenticatorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number to retrieve | 
 **numberPerPage** | **int32** | Number of Site Authenticators per page | 
 **filter** | **string** | Search for Site Authenticators with name matching filter text | 
 **name** | **string** | Get a Site Authenticator by name (case-sensitive) | 
 **sortKey** | **string** | A comma separated list of Site Authenticator attributes (keys) to be used to sort the results | 
 **order** | **string** | Order of the sorted results (ASC for ascending, DESC for descending) | 

### Return type

[**SiteAuthenticators**](SiteAuthenticators.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSiteAuthenticator

> SiteAuthenticator UpdateSiteAuthenticator(ctx, id).SiteAuthenticator(siteAuthenticator).Execute()





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
    id := "id_example" // string | ID of Site Authenticator to update
    siteAuthenticator := *openapiclient.NewSiteAuthenticator("ClassName_example", "Name_example") // SiteAuthenticator | Site Authenticator to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SiteAuthenticatorsApi.UpdateSiteAuthenticator(context.Background(), id).SiteAuthenticator(siteAuthenticator).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SiteAuthenticatorsApi.UpdateSiteAuthenticator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSiteAuthenticator`: SiteAuthenticator
    fmt.Fprintf(os.Stdout, "Response from `SiteAuthenticatorsApi.UpdateSiteAuthenticator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Site Authenticator to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteAuthenticatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteAuthenticator** | [**SiteAuthenticator**](SiteAuthenticator.md) | Site Authenticator to update | 

### Return type

[**SiteAuthenticator**](SiteAuthenticator.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

