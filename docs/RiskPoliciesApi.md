# \RiskPoliciesApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRiskPolicy**](RiskPoliciesApi.md#AddRiskPolicy) | **Post** /riskPolicies | 
[**DeleteRiskPolicy**](RiskPoliciesApi.md#DeleteRiskPolicy) | **Delete** /riskPolicies/{id} | 
[**GetRiskPolicies**](RiskPoliciesApi.md#GetRiskPolicies) | **Get** /riskPolicies | 
[**GetRiskPolicy**](RiskPoliciesApi.md#GetRiskPolicy) | **Get** /riskPolicies/{id} | 
[**GetRiskPolicyDescriptor**](RiskPoliciesApi.md#GetRiskPolicyDescriptor) | **Get** /riskPolicies/descriptors/{riskPolicyType} | 
[**GetRiskPolicyDescriptors**](RiskPoliciesApi.md#GetRiskPolicyDescriptors) | **Get** /riskPolicies/descriptors | 
[**UpdateRiskPolicy**](RiskPoliciesApi.md#UpdateRiskPolicy) | **Put** /riskPolicies/{id} | 



## AddRiskPolicy

> RiskPolicy AddRiskPolicy(ctx).RiskPolicy(riskPolicy).Execute()





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
    riskPolicy := *openapiclient.NewRiskPolicy("ClassName_example", "Name_example") // RiskPolicy | Risk Policy to add

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RiskPoliciesApi.AddRiskPolicy(context.Background()).RiskPolicy(riskPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskPoliciesApi.AddRiskPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRiskPolicy`: RiskPolicy
    fmt.Fprintf(os.Stdout, "Response from `RiskPoliciesApi.AddRiskPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddRiskPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **riskPolicy** | [**RiskPolicy**](RiskPolicy.md) | Risk Policy to add | 

### Return type

[**RiskPolicy**](RiskPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRiskPolicy

> DeleteRiskPolicy(ctx, id).Execute()





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
    id := "id_example" // string | ID of Risk Policy to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RiskPoliciesApi.DeleteRiskPolicy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskPoliciesApi.DeleteRiskPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Risk Policy to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRiskPolicyRequest struct via the builder pattern


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


## GetRiskPolicies

> RiskPolicy GetRiskPolicies(ctx).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()





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
    numberPerPage := int64(56) // int64 | Number of Risk Policies per page (optional)
    filter := "filter_example" // string | Search for Risk Policy with name matching filter text (optional)
    name := "name_example" // string | Get a specific Risk Policy by name (case-sensitive) (optional)
    sortKey := "sortKey_example" // string | A comma separated list of Risk Policy attributes (keys) to be used to sort the results (optional)
    order := "order_example" // string | Order of the sorted results (ASC for ascending, DESC for descending) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RiskPoliciesApi.GetRiskPolicies(context.Background()).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskPoliciesApi.GetRiskPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRiskPolicies`: RiskPolicy
    fmt.Fprintf(os.Stdout, "Response from `RiskPoliciesApi.GetRiskPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRiskPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | Page number to retrieve | 
 **numberPerPage** | **int64** | Number of Risk Policies per page | 
 **filter** | **string** | Search for Risk Policy with name matching filter text | 
 **name** | **string** | Get a specific Risk Policy by name (case-sensitive) | 
 **sortKey** | **string** | A comma separated list of Risk Policy attributes (keys) to be used to sort the results | 
 **order** | **string** | Order of the sorted results (ASC for ascending, DESC for descending) | 

### Return type

[**RiskPolicy**](RiskPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRiskPolicy

> RiskPolicy GetRiskPolicy(ctx, id).Execute()





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
    id := "id_example" // string | ID of Risk Policy to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RiskPoliciesApi.GetRiskPolicy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskPoliciesApi.GetRiskPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRiskPolicy`: RiskPolicy
    fmt.Fprintf(os.Stdout, "Response from `RiskPoliciesApi.GetRiskPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Risk Policy to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRiskPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RiskPolicy**](RiskPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRiskPolicyDescriptor

> Descriptor GetRiskPolicyDescriptor(ctx, riskPolicyType).Execute()





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
    riskPolicyType := "riskPolicyType_example" // string | Risk Policy descriptor to retrieve

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RiskPoliciesApi.GetRiskPolicyDescriptor(context.Background(), riskPolicyType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskPoliciesApi.GetRiskPolicyDescriptor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRiskPolicyDescriptor`: Descriptor
    fmt.Fprintf(os.Stdout, "Response from `RiskPoliciesApi.GetRiskPolicyDescriptor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**riskPolicyType** | **string** | Risk Policy descriptor to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRiskPolicyDescriptorRequest struct via the builder pattern


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


## GetRiskPolicyDescriptors

> Descriptors GetRiskPolicyDescriptors(ctx).Execute()





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
    resp, r, err := apiClient.RiskPoliciesApi.GetRiskPolicyDescriptors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskPoliciesApi.GetRiskPolicyDescriptors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRiskPolicyDescriptors`: Descriptors
    fmt.Fprintf(os.Stdout, "Response from `RiskPoliciesApi.GetRiskPolicyDescriptors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRiskPolicyDescriptorsRequest struct via the builder pattern


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


## UpdateRiskPolicy

> RiskPolicy UpdateRiskPolicy(ctx, id).RiskPolicy(riskPolicy).Execute()





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
    id := "id_example" // string | ID of the Risk Policy to update
    riskPolicy := *openapiclient.NewRiskPolicy("ClassName_example", "Name_example") // RiskPolicy | Risk Policy to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RiskPoliciesApi.UpdateRiskPolicy(context.Background(), id).RiskPolicy(riskPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskPoliciesApi.UpdateRiskPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRiskPolicy`: RiskPolicy
    fmt.Fprintf(os.Stdout, "Response from `RiskPoliciesApi.UpdateRiskPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Risk Policy to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRiskPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **riskPolicy** | [**RiskPolicy**](RiskPolicy.md) | Risk Policy to update | 

### Return type

[**RiskPolicy**](RiskPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

