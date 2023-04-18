# \HighAvailabilityApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAvailabilityProfile**](HighAvailabilityApi.md#AddAvailabilityProfile) | **Post** /highAvailability/availabilityProfiles | 
[**AddLoadBalancingStrategy**](HighAvailabilityApi.md#AddLoadBalancingStrategy) | **Post** /highAvailability/loadBalancingStrategies | 
[**DeleteAvailabilityProfile**](HighAvailabilityApi.md#DeleteAvailabilityProfile) | **Delete** /highAvailability/availabilityProfiles/{id} | 
[**DeleteLoadBalancingStrategy**](HighAvailabilityApi.md#DeleteLoadBalancingStrategy) | **Delete** /highAvailability/loadBalancingStrategies/{id} | 
[**GetAvailabilityProfile**](HighAvailabilityApi.md#GetAvailabilityProfile) | **Get** /highAvailability/availabilityProfiles/{id} | 
[**GetAvailabilityProfileDescriptor**](HighAvailabilityApi.md#GetAvailabilityProfileDescriptor) | **Get** /highAvailability/availabilityProfiles/descriptors/{availabilityProfileType} | 
[**GetAvailabilityProfileDescriptors**](HighAvailabilityApi.md#GetAvailabilityProfileDescriptors) | **Get** /highAvailability/availabilityProfiles/descriptors | 
[**GetAvailabilityProfiles**](HighAvailabilityApi.md#GetAvailabilityProfiles) | **Get** /highAvailability/availabilityProfiles | 
[**GetLoadBalancingStrategies**](HighAvailabilityApi.md#GetLoadBalancingStrategies) | **Get** /highAvailability/loadBalancingStrategies | 
[**GetLoadBalancingStrategy**](HighAvailabilityApi.md#GetLoadBalancingStrategy) | **Get** /highAvailability/loadBalancingStrategies/{id} | 
[**GetLoadBalancingStrategyDescriptor**](HighAvailabilityApi.md#GetLoadBalancingStrategyDescriptor) | **Get** /highAvailability/loadBalancingStrategies/descriptors/{loadBalancingStrategyType} | 
[**GetLoadBalancingStrategyDescriptors**](HighAvailabilityApi.md#GetLoadBalancingStrategyDescriptors) | **Get** /highAvailability/loadBalancingStrategies/descriptors | 
[**UpdateAvailabilityProfile**](HighAvailabilityApi.md#UpdateAvailabilityProfile) | **Put** /highAvailability/availabilityProfiles/{id} | 
[**UpdateLoadBalancingStrategy**](HighAvailabilityApi.md#UpdateLoadBalancingStrategy) | **Put** /highAvailability/loadBalancingStrategies/{id} | 



## AddAvailabilityProfile

> AvailabilityProfile AddAvailabilityProfile(ctx).AvailabilityProfile(availabilityProfile).Execute()





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
    availabilityProfile := *openapiclient.NewAvailabilityProfile("ClassName_example", "Name_example") // AvailabilityProfile | Availability Profile to add

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HighAvailabilityApi.AddAvailabilityProfile(context.Background()).AvailabilityProfile(availabilityProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HighAvailabilityApi.AddAvailabilityProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAvailabilityProfile`: AvailabilityProfile
    fmt.Fprintf(os.Stdout, "Response from `HighAvailabilityApi.AddAvailabilityProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAvailabilityProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **availabilityProfile** | [**AvailabilityProfile**](AvailabilityProfile.md) | Availability Profile to add | 

### Return type

[**AvailabilityProfile**](AvailabilityProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddLoadBalancingStrategy

> LoadBalancingStrategy AddLoadBalancingStrategy(ctx).LoadBalancingStrategy(loadBalancingStrategy).Execute()





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
    loadBalancingStrategy := *openapiclient.NewLoadBalancingStrategy("ClassName_example", "Name_example") // LoadBalancingStrategy | Load Balancing Strategy to add

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HighAvailabilityApi.AddLoadBalancingStrategy(context.Background()).LoadBalancingStrategy(loadBalancingStrategy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HighAvailabilityApi.AddLoadBalancingStrategy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddLoadBalancingStrategy`: LoadBalancingStrategy
    fmt.Fprintf(os.Stdout, "Response from `HighAvailabilityApi.AddLoadBalancingStrategy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddLoadBalancingStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loadBalancingStrategy** | [**LoadBalancingStrategy**](LoadBalancingStrategy.md) | Load Balancing Strategy to add | 

### Return type

[**LoadBalancingStrategy**](LoadBalancingStrategy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAvailabilityProfile

> DeleteAvailabilityProfile(ctx, id).Execute()





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
    id := "id_example" // string | ID of the Availability Profile to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HighAvailabilityApi.DeleteAvailabilityProfile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HighAvailabilityApi.DeleteAvailabilityProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Availability Profile to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAvailabilityProfileRequest struct via the builder pattern


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


## DeleteLoadBalancingStrategy

> DeleteLoadBalancingStrategy(ctx, id).Execute()





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
    id := "id_example" // string | ID of the load balancing strategy to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HighAvailabilityApi.DeleteLoadBalancingStrategy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HighAvailabilityApi.DeleteLoadBalancingStrategy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the load balancing strategy to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLoadBalancingStrategyRequest struct via the builder pattern


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


## GetAvailabilityProfile

> AvailabilityProfile GetAvailabilityProfile(ctx, id).Execute()





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
    id := "id_example" // string | ID of Availability Profile to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HighAvailabilityApi.GetAvailabilityProfile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HighAvailabilityApi.GetAvailabilityProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAvailabilityProfile`: AvailabilityProfile
    fmt.Fprintf(os.Stdout, "Response from `HighAvailabilityApi.GetAvailabilityProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Availability Profile to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailabilityProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AvailabilityProfile**](AvailabilityProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvailabilityProfileDescriptor

> Descriptor GetAvailabilityProfileDescriptor(ctx, availabilityProfileType).Execute()





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
    availabilityProfileType := "availabilityProfileType_example" // string | Availability Profile descriptor to retrieve

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HighAvailabilityApi.GetAvailabilityProfileDescriptor(context.Background(), availabilityProfileType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HighAvailabilityApi.GetAvailabilityProfileDescriptor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAvailabilityProfileDescriptor`: Descriptor
    fmt.Fprintf(os.Stdout, "Response from `HighAvailabilityApi.GetAvailabilityProfileDescriptor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**availabilityProfileType** | **string** | Availability Profile descriptor to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailabilityProfileDescriptorRequest struct via the builder pattern


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


## GetAvailabilityProfileDescriptors

> Descriptors GetAvailabilityProfileDescriptors(ctx).Execute()





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
    resp, r, err := apiClient.HighAvailabilityApi.GetAvailabilityProfileDescriptors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HighAvailabilityApi.GetAvailabilityProfileDescriptors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAvailabilityProfileDescriptors`: Descriptors
    fmt.Fprintf(os.Stdout, "Response from `HighAvailabilityApi.GetAvailabilityProfileDescriptors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailabilityProfileDescriptorsRequest struct via the builder pattern


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


## GetAvailabilityProfiles

> AvailabilityProfiles GetAvailabilityProfiles(ctx).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()





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
    numberPerPage := int64(56) // int64 | Number of Availability Profiles per page (optional)
    filter := "filter_example" // string | Search for Availability Profile with name matching filter text (optional)
    name := "name_example" // string | Get a specific Availability Profile by name (case-sensitive) (optional)
    sortKey := "sortKey_example" // string | A comma separated list of Availability Profile attributes (keys) to be used to sort the results (optional)
    order := "order_example" // string | Order of the sorted results (ASC for ascending, DESC for descending) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HighAvailabilityApi.GetAvailabilityProfiles(context.Background()).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HighAvailabilityApi.GetAvailabilityProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAvailabilityProfiles`: AvailabilityProfiles
    fmt.Fprintf(os.Stdout, "Response from `HighAvailabilityApi.GetAvailabilityProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailabilityProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | Page number to retrieve | 
 **numberPerPage** | **int64** | Number of Availability Profiles per page | 
 **filter** | **string** | Search for Availability Profile with name matching filter text | 
 **name** | **string** | Get a specific Availability Profile by name (case-sensitive) | 
 **sortKey** | **string** | A comma separated list of Availability Profile attributes (keys) to be used to sort the results | 
 **order** | **string** | Order of the sorted results (ASC for ascending, DESC for descending) | 

### Return type

[**AvailabilityProfiles**](AvailabilityProfiles.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoadBalancingStrategies

> LoadBalancingStrategies GetLoadBalancingStrategies(ctx).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()





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
    numberPerPage := int64(56) // int64 | Number of Load Balancing Strategies per page (optional)
    filter := "filter_example" // string | Search for Load Balancing Strategies with name matching filter text (optional)
    name := "name_example" // string | Get a specific Load Balancing Strategy by name (case-sensitive) (optional)
    sortKey := "sortKey_example" // string | A comma separated list of Load Balancing Strategy attributes (keys) to be used to sort the results (optional)
    order := "order_example" // string | Order of the sorted results (ASC for ascending, DESC for descending) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HighAvailabilityApi.GetLoadBalancingStrategies(context.Background()).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HighAvailabilityApi.GetLoadBalancingStrategies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLoadBalancingStrategies`: LoadBalancingStrategies
    fmt.Fprintf(os.Stdout, "Response from `HighAvailabilityApi.GetLoadBalancingStrategies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLoadBalancingStrategiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | Page number to retrieve | 
 **numberPerPage** | **int64** | Number of Load Balancing Strategies per page | 
 **filter** | **string** | Search for Load Balancing Strategies with name matching filter text | 
 **name** | **string** | Get a specific Load Balancing Strategy by name (case-sensitive) | 
 **sortKey** | **string** | A comma separated list of Load Balancing Strategy attributes (keys) to be used to sort the results | 
 **order** | **string** | Order of the sorted results (ASC for ascending, DESC for descending) | 

### Return type

[**LoadBalancingStrategies**](LoadBalancingStrategies.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoadBalancingStrategy

> LoadBalancingStrategy GetLoadBalancingStrategy(ctx, id).Execute()





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
    id := "id_example" // string | ID of Load Balancing Strategy to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HighAvailabilityApi.GetLoadBalancingStrategy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HighAvailabilityApi.GetLoadBalancingStrategy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLoadBalancingStrategy`: LoadBalancingStrategy
    fmt.Fprintf(os.Stdout, "Response from `HighAvailabilityApi.GetLoadBalancingStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Load Balancing Strategy to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoadBalancingStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoadBalancingStrategy**](LoadBalancingStrategy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoadBalancingStrategyDescriptor

> Descriptor GetLoadBalancingStrategyDescriptor(ctx, loadBalancingStrategyType).Execute()





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
    loadBalancingStrategyType := "loadBalancingStrategyType_example" // string | Load Balancing Strategy descriptor to retrieve

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HighAvailabilityApi.GetLoadBalancingStrategyDescriptor(context.Background(), loadBalancingStrategyType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HighAvailabilityApi.GetLoadBalancingStrategyDescriptor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLoadBalancingStrategyDescriptor`: Descriptor
    fmt.Fprintf(os.Stdout, "Response from `HighAvailabilityApi.GetLoadBalancingStrategyDescriptor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancingStrategyType** | **string** | Load Balancing Strategy descriptor to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoadBalancingStrategyDescriptorRequest struct via the builder pattern


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


## GetLoadBalancingStrategyDescriptors

> Descriptors GetLoadBalancingStrategyDescriptors(ctx).Execute()





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
    resp, r, err := apiClient.HighAvailabilityApi.GetLoadBalancingStrategyDescriptors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HighAvailabilityApi.GetLoadBalancingStrategyDescriptors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLoadBalancingStrategyDescriptors`: Descriptors
    fmt.Fprintf(os.Stdout, "Response from `HighAvailabilityApi.GetLoadBalancingStrategyDescriptors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoadBalancingStrategyDescriptorsRequest struct via the builder pattern


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


## UpdateAvailabilityProfile

> AvailabilityProfile UpdateAvailabilityProfile(ctx, id).AvailabilityProfile(availabilityProfile).Execute()





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
    id := "id_example" // string | ID of Availability Profile to update
    availabilityProfile := *openapiclient.NewAvailabilityProfile("ClassName_example", "Name_example") // AvailabilityProfile | Availability Profile to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HighAvailabilityApi.UpdateAvailabilityProfile(context.Background(), id).AvailabilityProfile(availabilityProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HighAvailabilityApi.UpdateAvailabilityProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAvailabilityProfile`: AvailabilityProfile
    fmt.Fprintf(os.Stdout, "Response from `HighAvailabilityApi.UpdateAvailabilityProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Availability Profile to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAvailabilityProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **availabilityProfile** | [**AvailabilityProfile**](AvailabilityProfile.md) | Availability Profile to update | 

### Return type

[**AvailabilityProfile**](AvailabilityProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLoadBalancingStrategy

> LoadBalancingStrategy UpdateLoadBalancingStrategy(ctx, id).LoadBalancingStrategy(loadBalancingStrategy).Execute()





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
    id := "id_example" // string | ID of Load Balancing Strategy to update
    loadBalancingStrategy := *openapiclient.NewLoadBalancingStrategy("ClassName_example", "Name_example") // LoadBalancingStrategy | Load Balancing Strategy to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HighAvailabilityApi.UpdateLoadBalancingStrategy(context.Background(), id).LoadBalancingStrategy(loadBalancingStrategy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HighAvailabilityApi.UpdateLoadBalancingStrategy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLoadBalancingStrategy`: LoadBalancingStrategy
    fmt.Fprintf(os.Stdout, "Response from `HighAvailabilityApi.UpdateLoadBalancingStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Load Balancing Strategy to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLoadBalancingStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **loadBalancingStrategy** | [**LoadBalancingStrategy**](LoadBalancingStrategy.md) | Load Balancing Strategy to update | 

### Return type

[**LoadBalancingStrategy**](LoadBalancingStrategy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

