# \IdentityMappingsApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddIdentityMapping**](IdentityMappingsApi.md#AddIdentityMapping) | **Post** /identityMappings | 
[**DeleteIdentityMapping**](IdentityMappingsApi.md#DeleteIdentityMapping) | **Delete** /identityMappings/{id} | 
[**GetIdentityMapping**](IdentityMappingsApi.md#GetIdentityMapping) | **Get** /identityMappings/{id} | 
[**GetIdentityMappingDescriptor**](IdentityMappingsApi.md#GetIdentityMappingDescriptor) | **Get** /identityMappings/descriptors/{identityMappingType} | 
[**GetIdentityMappingDescriptors**](IdentityMappingsApi.md#GetIdentityMappingDescriptors) | **Get** /identityMappings/descriptors | 
[**GetIdentityMappings**](IdentityMappingsApi.md#GetIdentityMappings) | **Get** /identityMappings | 
[**UpdateIdentityMapping**](IdentityMappingsApi.md#UpdateIdentityMapping) | **Put** /identityMappings/{id} | 



## AddIdentityMapping

> IdentityMapping AddIdentityMapping(ctx).IdentityMappings(identityMappings).Execute()





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
    identityMappings := *openapiclient.NewIdentityMapping("ClassName_example", "Name_example") // IdentityMapping | IdentityMapping to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityMappingsApi.AddIdentityMapping(context.Background()).IdentityMappings(identityMappings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityMappingsApi.AddIdentityMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddIdentityMapping`: IdentityMapping
    fmt.Fprintf(os.Stdout, "Response from `IdentityMappingsApi.AddIdentityMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddIdentityMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identityMappings** | [**IdentityMapping**](IdentityMapping.md) | IdentityMapping to create | 

### Return type

[**IdentityMapping**](IdentityMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentityMapping

> DeleteIdentityMapping(ctx, id).Execute()





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
    id := "id_example" // string | ID of IdentityMapping to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IdentityMappingsApi.DeleteIdentityMapping(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityMappingsApi.DeleteIdentityMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of IdentityMapping to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityMappingRequest struct via the builder pattern


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


## GetIdentityMapping

> IdentityMapping GetIdentityMapping(ctx, id).Execute()





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
    id := "id_example" // string | ID of identityMapping to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityMappingsApi.GetIdentityMapping(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityMappingsApi.GetIdentityMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityMapping`: IdentityMapping
    fmt.Fprintf(os.Stdout, "Response from `IdentityMappingsApi.GetIdentityMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of identityMapping to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdentityMapping**](IdentityMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityMappingDescriptor

> Descriptor GetIdentityMappingDescriptor(ctx, identityMappingType).Execute()





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
    identityMappingType := "identityMappingType_example" // string | Identity Mapping descriptor to retrieve

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityMappingsApi.GetIdentityMappingDescriptor(context.Background(), identityMappingType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityMappingsApi.GetIdentityMappingDescriptor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityMappingDescriptor`: Descriptor
    fmt.Fprintf(os.Stdout, "Response from `IdentityMappingsApi.GetIdentityMappingDescriptor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityMappingType** | **string** | Identity Mapping descriptor to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityMappingDescriptorRequest struct via the builder pattern


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


## GetIdentityMappingDescriptors

> Descriptors GetIdentityMappingDescriptors(ctx).Execute()





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
    resp, r, err := apiClient.IdentityMappingsApi.GetIdentityMappingDescriptors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityMappingsApi.GetIdentityMappingDescriptors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityMappingDescriptors`: Descriptors
    fmt.Fprintf(os.Stdout, "Response from `IdentityMappingsApi.GetIdentityMappingDescriptors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityMappingDescriptorsRequest struct via the builder pattern


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


## GetIdentityMappings

> IdentityMappings GetIdentityMappings(ctx).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()





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
    numberPerPage := int32(56) // int32 | Number of Identity Mappings per page (optional)
    filter := "filter_example" // string | Search for Identity Mapping with name matching filter text (optional)
    name := "name_example" // string | Get a specific Identity Mapping by name (case-sensitive) (optional)
    sortKey := "sortKey_example" // string | A comma separated list of Identity Mapping attributes (keys) to be used to sort the results (optional)
    order := "order_example" // string | Order of the sorted results (ASC for ascending, DESC for descending) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityMappingsApi.GetIdentityMappings(context.Background()).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityMappingsApi.GetIdentityMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityMappings`: IdentityMappings
    fmt.Fprintf(os.Stdout, "Response from `IdentityMappingsApi.GetIdentityMappings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number to retrieve | 
 **numberPerPage** | **int32** | Number of Identity Mappings per page | 
 **filter** | **string** | Search for Identity Mapping with name matching filter text | 
 **name** | **string** | Get a specific Identity Mapping by name (case-sensitive) | 
 **sortKey** | **string** | A comma separated list of Identity Mapping attributes (keys) to be used to sort the results | 
 **order** | **string** | Order of the sorted results (ASC for ascending, DESC for descending) | 

### Return type

[**IdentityMappings**](IdentityMappings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdentityMapping

> IdentityMapping UpdateIdentityMapping(ctx, id).IdentityMappings(identityMappings).Execute()





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
    id := "id_example" // string | ID of IdentityMapping to update
    identityMappings := *openapiclient.NewIdentityMapping("ClassName_example", "Name_example") // IdentityMapping | IdentityMapping to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityMappingsApi.UpdateIdentityMapping(context.Background(), id).IdentityMappings(identityMappings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityMappingsApi.UpdateIdentityMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIdentityMapping`: IdentityMapping
    fmt.Fprintf(os.Stdout, "Response from `IdentityMappingsApi.UpdateIdentityMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of IdentityMapping to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdentityMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identityMappings** | [**IdentityMapping**](IdentityMapping.md) | IdentityMapping to update | 

### Return type

[**IdentityMapping**](IdentityMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

