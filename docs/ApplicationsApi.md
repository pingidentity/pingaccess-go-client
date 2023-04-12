# \ApplicationsApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddApplication**](ApplicationsApi.md#AddApplication) | **Post** /applications | 
[**AddApplicationResource**](ApplicationsApi.md#AddApplicationResource) | **Post** /applications/{id}/resources | 
[**DeleteApplication**](ApplicationsApi.md#DeleteApplication) | **Delete** /applications/{id} | 
[**DeleteApplicationResource**](ApplicationsApi.md#DeleteApplicationResource) | **Delete** /applications/{applicationId}/resources/{resourceId} | 
[**DeleteReservedApplication**](ApplicationsApi.md#DeleteReservedApplication) | **Delete** /applications/reserved | 
[**GetApplication**](ApplicationsApi.md#GetApplication) | **Get** /applications/{id} | 
[**GetApplicationResource**](ApplicationsApi.md#GetApplicationResource) | **Get** /applications/{applicationId}/resources/{resourceId} | 
[**GetApplicationResourceResponseGeneratorDescriptor**](ApplicationsApi.md#GetApplicationResourceResponseGeneratorDescriptor) | **Get** /applications/resources/responseGenerators/descriptors/{responseGeneratorType} | 
[**GetApplicationResourceResponseGeneratorDescriptors**](ApplicationsApi.md#GetApplicationResourceResponseGeneratorDescriptors) | **Get** /applications/resources/responseGenerators/descriptors | 
[**GetApplicationResources**](ApplicationsApi.md#GetApplicationResources) | **Get** /applications/{id}/resources | 
[**GetApplications**](ApplicationsApi.md#GetApplications) | **Get** /applications | 
[**GetApplicationsResourcesMethods**](ApplicationsApi.md#GetApplicationsResourcesMethods) | **Get** /applications/resources/methods | 
[**GetReservedApplication**](ApplicationsApi.md#GetReservedApplication) | **Get** /applications/reserved | 
[**GetResourceAutoOrder**](ApplicationsApi.md#GetResourceAutoOrder) | **Get** /applications/{id}/resources/autoOrder | 
[**GetResourceMatchingEvaluationOrder**](ApplicationsApi.md#GetResourceMatchingEvaluationOrder) | **Get** /applications/{id}/resourceMatchingEvaluationOrder | 
[**GetResources**](ApplicationsApi.md#GetResources) | **Get** /applications/resources | 
[**UpdateApplication**](ApplicationsApi.md#UpdateApplication) | **Put** /applications/{id} | 
[**UpdateApplicationResource**](ApplicationsApi.md#UpdateApplicationResource) | **Put** /applications/{applicationId}/resources/{resourceId} | 
[**UpdateReservedApplication**](ApplicationsApi.md#UpdateReservedApplication) | **Put** /applications/reserved | 



## AddApplication

> Application AddApplication(ctx).Application(application).Execute()





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
    application := *openapiclient.NewApplication("Name_example", openapiclient.DefaultAuthType("Web"), false, "ContextRoot_example", int32(123), int32(123), "SidebandClientId_example", []int32{int32(123)}, "AuthenticationChallengePolicyId_example") // Application | Application to add

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApi.AddApplication(context.Background()).Application(application).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.AddApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddApplication`: Application
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsApi.AddApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **application** | [**Application**](Application.md) | Application to add | 

### Return type

[**Application**](Application.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddApplicationResource

> Resource AddApplicationResource(ctx, id).Resource(resource).Execute()





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
    id := "id_example" // string | ID of Application to add Resource to
    resource := *openapiclient.NewResource("Name_example", []string{"Methods_example"}, openapiclient.DefaultAuthType("Web"), "AuthenticationChallengePolicyId_example") // Resource | Resource to add

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApi.AddApplicationResource(context.Background(), id).Resource(resource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.AddApplicationResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddApplicationResource`: Resource
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsApi.AddApplicationResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Application to add Resource to | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddApplicationResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resource** | [**Resource**](Resource.md) | Resource to add | 

### Return type

[**Resource**](Resource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplication

> DeleteApplication(ctx, id).Execute()





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
    id := "id_example" // string | ID of Application to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationsApi.DeleteApplication(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.DeleteApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Application to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationRequest struct via the builder pattern


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


## DeleteApplicationResource

> DeleteApplicationResource(ctx, applicationId, resourceId).Execute()





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
    applicationId := "applicationId_example" // string | ID of Application
    resourceId := "resourceId_example" // string | ID of the Resource to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationsApi.DeleteApplicationResource(context.Background(), applicationId, resourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.DeleteApplicationResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | ID of Application | 
**resourceId** | **string** | ID of the Resource to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationResourceRequest struct via the builder pattern


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


## DeleteReservedApplication

> DeleteReservedApplication(ctx).Execute()





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
    r, err := apiClient.ApplicationsApi.DeleteReservedApplication(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.DeleteReservedApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReservedApplicationRequest struct via the builder pattern


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


## GetApplication

> Application GetApplication(ctx, id).Execute()





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
    id := "id_example" // string | ID of Application to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApi.GetApplication(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.GetApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplication`: Application
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsApi.GetApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Application to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Application**](Application.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationResource

> Resource GetApplicationResource(ctx, applicationId, resourceId).Execute()





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
    applicationId := "applicationId_example" // string | ID of Application
    resourceId := "resourceId_example" // string | ID of the Resource to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApi.GetApplicationResource(context.Background(), applicationId, resourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.GetApplicationResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationResource`: Resource
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsApi.GetApplicationResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | ID of Application | 
**resourceId** | **string** | ID of the Resource to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Resource**](Resource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationResourceResponseGeneratorDescriptor

> Descriptor GetApplicationResourceResponseGeneratorDescriptor(ctx, responseGeneratorType).Execute()





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
    resp, r, err := apiClient.ApplicationsApi.GetApplicationResourceResponseGeneratorDescriptor(context.Background(), responseGeneratorType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.GetApplicationResourceResponseGeneratorDescriptor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationResourceResponseGeneratorDescriptor`: Descriptor
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsApi.GetApplicationResourceResponseGeneratorDescriptor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**responseGeneratorType** | **string** | Response Generator descriptor to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationResourceResponseGeneratorDescriptorRequest struct via the builder pattern


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


## GetApplicationResourceResponseGeneratorDescriptors

> Descriptors GetApplicationResourceResponseGeneratorDescriptors(ctx).Execute()





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
    resp, r, err := apiClient.ApplicationsApi.GetApplicationResourceResponseGeneratorDescriptors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.GetApplicationResourceResponseGeneratorDescriptors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationResourceResponseGeneratorDescriptors`: Descriptors
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsApi.GetApplicationResourceResponseGeneratorDescriptors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationResourceResponseGeneratorDescriptorsRequest struct via the builder pattern


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


## GetApplicationResources

> Resources GetApplicationResources(ctx, id).Page(page).NumberPerPage(numberPerPage).Name(name).Filter(filter).SortKey(sortKey).Order(order).Execute()





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
    id := "id_example" // string | ID of Application to get Resources for
    page := int32(56) // int32 | Page number to retrieve (optional)
    numberPerPage := int32(56) // int32 | Number of Resources per page (optional)
    name := "name_example" // string | Get a Resource by name and application Id (optional)
    filter := "filter_example" // string | Search for Resources with name matching filter text (optional)
    sortKey := "sortKey_example" // string | A comma separated list of Resource attributes (keys) to be used to sort the results (optional)
    order := "order_example" // string | Order of the sorted results (ASC for ascending, DESC for descending) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApi.GetApplicationResources(context.Background(), id).Page(page).NumberPerPage(numberPerPage).Name(name).Filter(filter).SortKey(sortKey).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.GetApplicationResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationResources`: Resources
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsApi.GetApplicationResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Application to get Resources for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number to retrieve | 
 **numberPerPage** | **int32** | Number of Resources per page | 
 **name** | **string** | Get a Resource by name and application Id | 
 **filter** | **string** | Search for Resources with name matching filter text | 
 **sortKey** | **string** | A comma separated list of Resource attributes (keys) to be used to sort the results | 
 **order** | **string** | Order of the sorted results (ASC for ascending, DESC for descending) | 

### Return type

[**Resources**](Resources.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplications

> Applications GetApplications(ctx).Page(page).SiteId(siteId).NumberPerPage(numberPerPage).AgentId(agentId).SidebandClientId(sidebandClientId).VirtualHostId(virtualHostId).RuleId(ruleId).RulesetId(rulesetId).RiskPolicyId(riskPolicyId).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()





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
    siteId := int32(56) // int32 | Search for Applications with Site (optional)
    numberPerPage := int32(56) // int32 | Number of Applications per page (optional)
    agentId := int32(56) // int32 | Search for Applications with Agent (optional)
    sidebandClientId := int32(56) // int32 | Search for Applications with Sideband Client (optional)
    virtualHostId := int32(56) // int32 | Search for Applications with Virtual Host (optional)
    ruleId := int32(56) // int32 | Search for Applications with Rule attached to an Application or one of its Resources (optional)
    rulesetId := int32(56) // int32 | Search for Applications with Rule Set attached to an Application or one of its Resources (optional)
    riskPolicyId := int32(56) // int32 | Search for Applications with Risk Policy attached to an Application or one of its Resources (optional)
    filter := "filter_example" // string | Search for Applications with name matching filter text (optional)
    name := "name_example" // string | Get an Application by name (optional)
    sortKey := "sortKey_example" // string | A comma separated list of Application attributes (keys) to be used to sort the results (optional)
    order := "order_example" // string | Order of the sorted results (ASC for ascending, DESC for descending) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApi.GetApplications(context.Background()).Page(page).SiteId(siteId).NumberPerPage(numberPerPage).AgentId(agentId).SidebandClientId(sidebandClientId).VirtualHostId(virtualHostId).RuleId(ruleId).RulesetId(rulesetId).RiskPolicyId(riskPolicyId).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.GetApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplications`: Applications
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsApi.GetApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number to retrieve | 
 **siteId** | **int32** | Search for Applications with Site | 
 **numberPerPage** | **int32** | Number of Applications per page | 
 **agentId** | **int32** | Search for Applications with Agent | 
 **sidebandClientId** | **int32** | Search for Applications with Sideband Client | 
 **virtualHostId** | **int32** | Search for Applications with Virtual Host | 
 **ruleId** | **int32** | Search for Applications with Rule attached to an Application or one of its Resources | 
 **rulesetId** | **int32** | Search for Applications with Rule Set attached to an Application or one of its Resources | 
 **riskPolicyId** | **int32** | Search for Applications with Risk Policy attached to an Application or one of its Resources | 
 **filter** | **string** | Search for Applications with name matching filter text | 
 **name** | **string** | Get an Application by name | 
 **sortKey** | **string** | A comma separated list of Application attributes (keys) to be used to sort the results | 
 **order** | **string** | Order of the sorted results (ASC for ascending, DESC for descending) | 

### Return type

[**Applications**](Applications.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationsResourcesMethods

> Methods GetApplicationsResourcesMethods(ctx).Execute()





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
    resp, r, err := apiClient.ApplicationsApi.GetApplicationsResourcesMethods(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.GetApplicationsResourcesMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationsResourcesMethods`: Methods
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsApi.GetApplicationsResourcesMethods`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationsResourcesMethodsRequest struct via the builder pattern


### Return type

[**Methods**](Methods.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReservedApplication

> ReservedApplication GetReservedApplication(ctx).Execute()





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
    resp, r, err := apiClient.ApplicationsApi.GetReservedApplication(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.GetReservedApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReservedApplication`: ReservedApplication
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsApi.GetReservedApplication`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetReservedApplicationRequest struct via the builder pattern


### Return type

[**ReservedApplication**](ReservedApplication.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResourceAutoOrder

> ResourceOrder GetResourceAutoOrder(ctx, id).Execute()





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
    id := "id_example" // string | ID of Application to compute resource ordering for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApi.GetResourceAutoOrder(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.GetResourceAutoOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResourceAutoOrder`: ResourceOrder
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsApi.GetResourceAutoOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Application to compute resource ordering for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceAutoOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResourceOrder**](ResourceOrder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResourceMatchingEvaluationOrder

> ResourceMatchingEvaluationOrder GetResourceMatchingEvaluationOrder(ctx, id).Execute()





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
    id := "id_example" // string | ID of Application to compute resource ordering for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApi.GetResourceMatchingEvaluationOrder(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.GetResourceMatchingEvaluationOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResourceMatchingEvaluationOrder`: ResourceMatchingEvaluationOrder
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsApi.GetResourceMatchingEvaluationOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Application to compute resource ordering for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceMatchingEvaluationOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResourceMatchingEvaluationOrder**](ResourceMatchingEvaluationOrder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResources

> Resources GetResources(ctx).Page(page).NumberPerPage(numberPerPage).RuleId(ruleId).RulesetId(rulesetId).RiskPolicyId(riskPolicyId).Name(name).Filter(filter).SortKey(sortKey).Order(order).Execute()





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
    numberPerPage := int32(56) // int32 | Number of Resources per page (optional)
    ruleId := int32(56) // int32 | Get Resources with Rule (optional)
    rulesetId := int32(56) // int32 | Get Resources with Rule Set (optional)
    riskPolicyId := int32(56) // int32 | Get Resources with Risk Policy (optional)
    name := "name_example" // string | Get Resources by name (optional)
    filter := "filter_example" // string | Search for Resources with name matching filter text (optional)
    sortKey := "sortKey_example" // string | A comma separated list of Resource attributes (keys) to be used to sort the results (optional)
    order := "order_example" // string | Order of the sorted results (ASC for ascending, DESC for descending) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApi.GetResources(context.Background()).Page(page).NumberPerPage(numberPerPage).RuleId(ruleId).RulesetId(rulesetId).RiskPolicyId(riskPolicyId).Name(name).Filter(filter).SortKey(sortKey).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.GetResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResources`: Resources
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsApi.GetResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number to retrieve | 
 **numberPerPage** | **int32** | Number of Resources per page | 
 **ruleId** | **int32** | Get Resources with Rule | 
 **rulesetId** | **int32** | Get Resources with Rule Set | 
 **riskPolicyId** | **int32** | Get Resources with Risk Policy | 
 **name** | **string** | Get Resources by name | 
 **filter** | **string** | Search for Resources with name matching filter text | 
 **sortKey** | **string** | A comma separated list of Resource attributes (keys) to be used to sort the results | 
 **order** | **string** | Order of the sorted results (ASC for ascending, DESC for descending) | 

### Return type

[**Resources**](Resources.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplication

> Application UpdateApplication(ctx, id).Application(application).Execute()





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
    id := "id_example" // string | ID of Application to update
    application := *openapiclient.NewApplication("Name_example", openapiclient.DefaultAuthType("Web"), false, "ContextRoot_example", int32(123), int32(123), "SidebandClientId_example", []int32{int32(123)}, "AuthenticationChallengePolicyId_example") // Application | Application to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApi.UpdateApplication(context.Background(), id).Application(application).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.UpdateApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApplication`: Application
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsApi.UpdateApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Application to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **application** | [**Application**](Application.md) | Application to update | 

### Return type

[**Application**](Application.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationResource

> Resource UpdateApplicationResource(ctx, applicationId, resourceId).Resource(resource).Execute()





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
    applicationId := "applicationId_example" // string | ID of Application
    resourceId := "resourceId_example" // string | ID of Resource to update
    resource := *openapiclient.NewResource("Name_example", []string{"Methods_example"}, openapiclient.DefaultAuthType("Web"), "AuthenticationChallengePolicyId_example") // Resource | Resource to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApi.UpdateApplicationResource(context.Background(), applicationId, resourceId).Resource(resource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.UpdateApplicationResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApplicationResource`: Resource
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsApi.UpdateApplicationResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | ID of Application | 
**resourceId** | **string** | ID of Resource to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resource** | [**Resource**](Resource.md) | Resource to update | 

### Return type

[**Resource**](Resource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReservedApplication

> ReservedApplication UpdateReservedApplication(ctx).ReservedApplication(reservedApplication).Execute()





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
    reservedApplication := *openapiclient.NewReservedApplication("ContextRoot_example") // ReservedApplication | Reserved Application configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApi.UpdateReservedApplication(context.Background()).ReservedApplication(reservedApplication).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.UpdateReservedApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateReservedApplication`: ReservedApplication
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsApi.UpdateReservedApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReservedApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reservedApplication** | [**ReservedApplication**](ReservedApplication.md) | Reserved Application configuration | 

### Return type

[**ReservedApplication**](ReservedApplication.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

