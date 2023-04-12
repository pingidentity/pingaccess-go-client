# \ThirdPartyServicesApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddThirdPartyService**](ThirdPartyServicesApi.md#AddThirdPartyService) | **Post** /thirdPartyServices | 
[**DeleteThirdPartyService**](ThirdPartyServicesApi.md#DeleteThirdPartyService) | **Delete** /thirdPartyServices/{id} | 
[**GetThirdPartyService**](ThirdPartyServicesApi.md#GetThirdPartyService) | **Get** /thirdPartyServices/{id} | 
[**GetThirdPartyServices**](ThirdPartyServicesApi.md#GetThirdPartyServices) | **Get** /thirdPartyServices | 
[**UpdateThirdPartyService**](ThirdPartyServicesApi.md#UpdateThirdPartyService) | **Put** /thirdPartyServices/{id} | 



## AddThirdPartyService

> ThirdPartyService AddThirdPartyService(ctx).ThirdPartyService(thirdPartyService).Execute()





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
    thirdPartyService := *openapiclient.NewThirdPartyService([]string{"Targets_example"}, "Name_example") // ThirdPartyService | Third-party service to add

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThirdPartyServicesApi.AddThirdPartyService(context.Background()).ThirdPartyService(thirdPartyService).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThirdPartyServicesApi.AddThirdPartyService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddThirdPartyService`: ThirdPartyService
    fmt.Fprintf(os.Stdout, "Response from `ThirdPartyServicesApi.AddThirdPartyService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddThirdPartyServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **thirdPartyService** | [**ThirdPartyService**](ThirdPartyService.md) | Third-party service to add | 

### Return type

[**ThirdPartyService**](ThirdPartyService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteThirdPartyService

> DeleteThirdPartyService(ctx, id).Execute()





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
    id := "id_example" // string | ID of Third-Party Service to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ThirdPartyServicesApi.DeleteThirdPartyService(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThirdPartyServicesApi.DeleteThirdPartyService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Third-Party Service to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteThirdPartyServiceRequest struct via the builder pattern


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


## GetThirdPartyService

> ThirdPartyService GetThirdPartyService(ctx, id).Execute()





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
    id := "id_example" // string | ID of Third-Party Service to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThirdPartyServicesApi.GetThirdPartyService(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThirdPartyServicesApi.GetThirdPartyService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetThirdPartyService`: ThirdPartyService
    fmt.Fprintf(os.Stdout, "Response from `ThirdPartyServicesApi.GetThirdPartyService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Third-Party Service to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetThirdPartyServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ThirdPartyService**](ThirdPartyService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetThirdPartyServices

> ThirdPartyServices GetThirdPartyServices(ctx).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()





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
    numberPerPage := int32(56) // int32 | Number of Third-Party Services per page (optional)
    filter := "filter_example" // string | Search for Third-Party Services with name matching filter text (optional)
    name := "name_example" // string | Get a Third-Party Service by name (case-sensitive) (optional)
    sortKey := "sortKey_example" // string | A comma separated list of Third-Party Service attributes (keys) to be used to sort the results (optional)
    order := "order_example" // string | Order of the sorted results (ASC for ascending, DESC for descending) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThirdPartyServicesApi.GetThirdPartyServices(context.Background()).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThirdPartyServicesApi.GetThirdPartyServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetThirdPartyServices`: ThirdPartyServices
    fmt.Fprintf(os.Stdout, "Response from `ThirdPartyServicesApi.GetThirdPartyServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetThirdPartyServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number to retrieve | 
 **numberPerPage** | **int32** | Number of Third-Party Services per page | 
 **filter** | **string** | Search for Third-Party Services with name matching filter text | 
 **name** | **string** | Get a Third-Party Service by name (case-sensitive) | 
 **sortKey** | **string** | A comma separated list of Third-Party Service attributes (keys) to be used to sort the results | 
 **order** | **string** | Order of the sorted results (ASC for ascending, DESC for descending) | 

### Return type

[**ThirdPartyServices**](ThirdPartyServices.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateThirdPartyService

> ThirdPartyService UpdateThirdPartyService(ctx, id).ThirdPartyService(thirdPartyService).Execute()





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
    id := "id_example" // string | ID of Third-Party Service to update
    thirdPartyService := *openapiclient.NewThirdPartyService([]string{"Targets_example"}, "Name_example") // ThirdPartyService | Third-Party Service to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThirdPartyServicesApi.UpdateThirdPartyService(context.Background(), id).ThirdPartyService(thirdPartyService).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThirdPartyServicesApi.UpdateThirdPartyService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateThirdPartyService`: ThirdPartyService
    fmt.Fprintf(os.Stdout, "Response from `ThirdPartyServicesApi.UpdateThirdPartyService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Third-Party Service to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateThirdPartyServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **thirdPartyService** | [**ThirdPartyService**](ThirdPartyService.md) | Third-Party Service to update | 

### Return type

[**ThirdPartyService**](ThirdPartyService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

