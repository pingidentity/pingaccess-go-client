# \AuthenticationChallengePoliciesApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAuthenticationChallengePolicy**](AuthenticationChallengePoliciesApi.md#AddAuthenticationChallengePolicy) | **Post** /authenticationChallengePolicies | 
[**DeleteAuthenticationChallengePolicy**](AuthenticationChallengePoliciesApi.md#DeleteAuthenticationChallengePolicy) | **Delete** /authenticationChallengePolicies/{id} | 
[**GetAuthenticationChallengePolicies**](AuthenticationChallengePoliciesApi.md#GetAuthenticationChallengePolicies) | **Get** /authenticationChallengePolicies | 
[**GetAuthenticationChallengePolicy**](AuthenticationChallengePoliciesApi.md#GetAuthenticationChallengePolicy) | **Get** /authenticationChallengePolicies/{id} | 
[**GetChallengeResponseFilterDescriptor**](AuthenticationChallengePoliciesApi.md#GetChallengeResponseFilterDescriptor) | **Get** /authenticationChallengePolicies/responseFilters/descriptors/{responseFilterType} | 
[**GetChallengeResponseFilterDescriptors**](AuthenticationChallengePoliciesApi.md#GetChallengeResponseFilterDescriptors) | **Get** /authenticationChallengePolicies/responseFilters/descriptors | 
[**GetChallengeResponseGeneratorDescriptor**](AuthenticationChallengePoliciesApi.md#GetChallengeResponseGeneratorDescriptor) | **Get** /authenticationChallengePolicies/responseGenerators/descriptors/{responseGeneratorType} | 
[**GetChallengeResponseGeneratorDescriptors**](AuthenticationChallengePoliciesApi.md#GetChallengeResponseGeneratorDescriptors) | **Get** /authenticationChallengePolicies/responseGenerators/descriptors | 
[**GetRequestMatcherDescriptor**](AuthenticationChallengePoliciesApi.md#GetRequestMatcherDescriptor) | **Get** /authenticationChallengePolicies/requestMatchers/descriptors/{requestMatcherType} | 
[**GetRequestMatcherDescriptors**](AuthenticationChallengePoliciesApi.md#GetRequestMatcherDescriptors) | **Get** /authenticationChallengePolicies/requestMatchers/descriptors | 
[**UpdateAuthenticationChallengePolicy**](AuthenticationChallengePoliciesApi.md#UpdateAuthenticationChallengePolicy) | **Put** /authenticationChallengePolicies/{id} | 



## AddAuthenticationChallengePolicy

> AuthenticationChallengePolicy AddAuthenticationChallengePolicy(ctx).AuthenticationChallengePolicy(authenticationChallengePolicy).Execute()





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
    authenticationChallengePolicy := *openapiclient.NewAuthenticationChallengePolicy("Name_example", []openapiclient.ChallengeResponseMapping{*openapiclient.NewChallengeResponseMapping(*openapiclient.NewRequestMatcher("ClassName_example"), *openapiclient.NewChallengeResponse(*openapiclient.NewChallengeResponseGenerator("ClassName_example")))}, *openapiclient.NewChallengeResponse(*openapiclient.NewChallengeResponseGenerator("ClassName_example"))) // AuthenticationChallengePolicy | Authentication Challenge Policy to add

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationChallengePoliciesApi.AddAuthenticationChallengePolicy(context.Background()).AuthenticationChallengePolicy(authenticationChallengePolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationChallengePoliciesApi.AddAuthenticationChallengePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAuthenticationChallengePolicy`: AuthenticationChallengePolicy
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationChallengePoliciesApi.AddAuthenticationChallengePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAuthenticationChallengePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authenticationChallengePolicy** | [**AuthenticationChallengePolicy**](AuthenticationChallengePolicy.md) | Authentication Challenge Policy to add | 

### Return type

[**AuthenticationChallengePolicy**](AuthenticationChallengePolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAuthenticationChallengePolicy

> DeleteAuthenticationChallengePolicy(ctx, id).Execute()





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
    id := "id_example" // string | ID of Authentication Challenge Policy to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthenticationChallengePoliciesApi.DeleteAuthenticationChallengePolicy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationChallengePoliciesApi.DeleteAuthenticationChallengePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Authentication Challenge Policy to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthenticationChallengePolicyRequest struct via the builder pattern


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


## GetAuthenticationChallengePolicies

> AuthenticationChallengePolicies GetAuthenticationChallengePolicies(ctx).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()





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
    numberPerPage := int64(56) // int64 | Number of Authentication Challenge Policies per page (optional)
    filter := "filter_example" // string | Search for an Authentication Challenge Policy with name matching filter text (optional)
    name := "name_example" // string | Get a specific Authentication Challenge Policy by name (case-sensitive) (optional)
    sortKey := "sortKey_example" // string | A comma separated list of Authentication Challenge Policy attributes (keys) to be used to sort the results (optional)
    order := "order_example" // string | Order of the sorted results (ASC for ascending, DESC for descending) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationChallengePoliciesApi.GetAuthenticationChallengePolicies(context.Background()).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationChallengePoliciesApi.GetAuthenticationChallengePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthenticationChallengePolicies`: AuthenticationChallengePolicies
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationChallengePoliciesApi.GetAuthenticationChallengePolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthenticationChallengePoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | Page number to retrieve | 
 **numberPerPage** | **int64** | Number of Authentication Challenge Policies per page | 
 **filter** | **string** | Search for an Authentication Challenge Policy with name matching filter text | 
 **name** | **string** | Get a specific Authentication Challenge Policy by name (case-sensitive) | 
 **sortKey** | **string** | A comma separated list of Authentication Challenge Policy attributes (keys) to be used to sort the results | 
 **order** | **string** | Order of the sorted results (ASC for ascending, DESC for descending) | 

### Return type

[**AuthenticationChallengePolicies**](AuthenticationChallengePolicies.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthenticationChallengePolicy

> AuthenticationChallengePolicy GetAuthenticationChallengePolicy(ctx, id).Execute()





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
    id := "id_example" // string | ID of Authentication Challenge Policy to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationChallengePoliciesApi.GetAuthenticationChallengePolicy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationChallengePoliciesApi.GetAuthenticationChallengePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthenticationChallengePolicy`: AuthenticationChallengePolicy
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationChallengePoliciesApi.GetAuthenticationChallengePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Authentication Challenge Policy to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthenticationChallengePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthenticationChallengePolicy**](AuthenticationChallengePolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChallengeResponseFilterDescriptor

> Descriptor GetChallengeResponseFilterDescriptor(ctx, responseFilterType).Execute()





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
    responseFilterType := "responseFilterType_example" // string | Response Filter descriptor to retrieve

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationChallengePoliciesApi.GetChallengeResponseFilterDescriptor(context.Background(), responseFilterType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationChallengePoliciesApi.GetChallengeResponseFilterDescriptor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChallengeResponseFilterDescriptor`: Descriptor
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationChallengePoliciesApi.GetChallengeResponseFilterDescriptor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**responseFilterType** | **string** | Response Filter descriptor to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChallengeResponseFilterDescriptorRequest struct via the builder pattern


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


## GetChallengeResponseFilterDescriptors

> Descriptors GetChallengeResponseFilterDescriptors(ctx).Execute()





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
    resp, r, err := apiClient.AuthenticationChallengePoliciesApi.GetChallengeResponseFilterDescriptors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationChallengePoliciesApi.GetChallengeResponseFilterDescriptors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChallengeResponseFilterDescriptors`: Descriptors
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationChallengePoliciesApi.GetChallengeResponseFilterDescriptors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetChallengeResponseFilterDescriptorsRequest struct via the builder pattern


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


## GetChallengeResponseGeneratorDescriptor

> Descriptor GetChallengeResponseGeneratorDescriptor(ctx, responseGeneratorType).Execute()





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
    responseGeneratorType := "responseGeneratorType_example" // string | Response Generator descriptor to retrieve

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationChallengePoliciesApi.GetChallengeResponseGeneratorDescriptor(context.Background(), responseGeneratorType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationChallengePoliciesApi.GetChallengeResponseGeneratorDescriptor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChallengeResponseGeneratorDescriptor`: Descriptor
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationChallengePoliciesApi.GetChallengeResponseGeneratorDescriptor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**responseGeneratorType** | **string** | Response Generator descriptor to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChallengeResponseGeneratorDescriptorRequest struct via the builder pattern


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


## GetChallengeResponseGeneratorDescriptors

> Descriptors GetChallengeResponseGeneratorDescriptors(ctx).Execute()





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
    resp, r, err := apiClient.AuthenticationChallengePoliciesApi.GetChallengeResponseGeneratorDescriptors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationChallengePoliciesApi.GetChallengeResponseGeneratorDescriptors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChallengeResponseGeneratorDescriptors`: Descriptors
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationChallengePoliciesApi.GetChallengeResponseGeneratorDescriptors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetChallengeResponseGeneratorDescriptorsRequest struct via the builder pattern


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


## GetRequestMatcherDescriptor

> Descriptor GetRequestMatcherDescriptor(ctx, requestMatcherType).Execute()





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
    requestMatcherType := "requestMatcherType_example" // string | Request Matcher descriptor to retrieve

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationChallengePoliciesApi.GetRequestMatcherDescriptor(context.Background(), requestMatcherType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationChallengePoliciesApi.GetRequestMatcherDescriptor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRequestMatcherDescriptor`: Descriptor
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationChallengePoliciesApi.GetRequestMatcherDescriptor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestMatcherType** | **string** | Request Matcher descriptor to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestMatcherDescriptorRequest struct via the builder pattern


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


## GetRequestMatcherDescriptors

> Descriptors GetRequestMatcherDescriptors(ctx).Execute()





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
    resp, r, err := apiClient.AuthenticationChallengePoliciesApi.GetRequestMatcherDescriptors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationChallengePoliciesApi.GetRequestMatcherDescriptors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRequestMatcherDescriptors`: Descriptors
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationChallengePoliciesApi.GetRequestMatcherDescriptors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestMatcherDescriptorsRequest struct via the builder pattern


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


## UpdateAuthenticationChallengePolicy

> AuthenticationChallengePolicy UpdateAuthenticationChallengePolicy(ctx, id).AuthenticationChallengePolicy(authenticationChallengePolicy).Execute()





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
    id := "id_example" // string | ID of the Authentication Challenge Policy to update
    authenticationChallengePolicy := *openapiclient.NewAuthenticationChallengePolicy("Name_example", []openapiclient.ChallengeResponseMapping{*openapiclient.NewChallengeResponseMapping(*openapiclient.NewRequestMatcher("ClassName_example"), *openapiclient.NewChallengeResponse(*openapiclient.NewChallengeResponseGenerator("ClassName_example")))}, *openapiclient.NewChallengeResponse(*openapiclient.NewChallengeResponseGenerator("ClassName_example"))) // AuthenticationChallengePolicy | Authentication Challenge Policy to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationChallengePoliciesApi.UpdateAuthenticationChallengePolicy(context.Background(), id).AuthenticationChallengePolicy(authenticationChallengePolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationChallengePoliciesApi.UpdateAuthenticationChallengePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAuthenticationChallengePolicy`: AuthenticationChallengePolicy
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationChallengePoliciesApi.UpdateAuthenticationChallengePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Authentication Challenge Policy to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAuthenticationChallengePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authenticationChallengePolicy** | [**AuthenticationChallengePolicy**](AuthenticationChallengePolicy.md) | Authentication Challenge Policy to update | 

### Return type

[**AuthenticationChallengePolicy**](AuthenticationChallengePolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

