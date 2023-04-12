# \RulesetsApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRuleSet**](RulesetsApi.md#AddRuleSet) | **Post** /rulesets | 
[**DeleteRuleSet**](RulesetsApi.md#DeleteRuleSet) | **Delete** /rulesets/{id} | 
[**GetRuleSet**](RulesetsApi.md#GetRuleSet) | **Get** /rulesets/{id} | 
[**GetRuleSetElementTypes**](RulesetsApi.md#GetRuleSetElementTypes) | **Get** /rulesets/elementTypes | 
[**GetRuleSetSuccessCriteria**](RulesetsApi.md#GetRuleSetSuccessCriteria) | **Get** /rulesets/successCriteria | 
[**GetRuleSets**](RulesetsApi.md#GetRuleSets) | **Get** /rulesets | 
[**UpdateRuleSet**](RulesetsApi.md#UpdateRuleSet) | **Put** /rulesets/{id} | 



## AddRuleSet

> RuleSet AddRuleSet(ctx).RuleSet(ruleSet).Execute()





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
    ruleSet := *openapiclient.NewRuleSet("Name_example", []int32{int32(123)}) // RuleSet | Rule Set to add

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesetsApi.AddRuleSet(context.Background()).RuleSet(ruleSet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesetsApi.AddRuleSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRuleSet`: RuleSet
    fmt.Fprintf(os.Stdout, "Response from `RulesetsApi.AddRuleSet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddRuleSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ruleSet** | [**RuleSet**](RuleSet.md) | Rule Set to add | 

### Return type

[**RuleSet**](RuleSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRuleSet

> DeleteRuleSet(ctx, id).Execute()





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
    id := "id_example" // string | ID of Rule Set to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RulesetsApi.DeleteRuleSet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesetsApi.DeleteRuleSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Rule Set to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRuleSetRequest struct via the builder pattern


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


## GetRuleSet

> RuleSet GetRuleSet(ctx, id).Execute()





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
    id := "id_example" // string | ID of Rule Set to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesetsApi.GetRuleSet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesetsApi.GetRuleSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRuleSet`: RuleSet
    fmt.Fprintf(os.Stdout, "Response from `RulesetsApi.GetRuleSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Rule Set to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RuleSet**](RuleSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleSetElementTypes

> RuleSetElementTypes GetRuleSetElementTypes(ctx).Execute()





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
    resp, r, err := apiClient.RulesetsApi.GetRuleSetElementTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesetsApi.GetRuleSetElementTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRuleSetElementTypes`: RuleSetElementTypes
    fmt.Fprintf(os.Stdout, "Response from `RulesetsApi.GetRuleSetElementTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleSetElementTypesRequest struct via the builder pattern


### Return type

[**RuleSetElementTypes**](RuleSetElementTypes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleSetSuccessCriteria

> RuleSetSuccessCriteria GetRuleSetSuccessCriteria(ctx).Execute()





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
    resp, r, err := apiClient.RulesetsApi.GetRuleSetSuccessCriteria(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesetsApi.GetRuleSetSuccessCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRuleSetSuccessCriteria`: RuleSetSuccessCriteria
    fmt.Fprintf(os.Stdout, "Response from `RulesetsApi.GetRuleSetSuccessCriteria`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleSetSuccessCriteriaRequest struct via the builder pattern


### Return type

[**RuleSetSuccessCriteria**](RuleSetSuccessCriteria.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleSets

> RuleSets GetRuleSets(ctx).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()





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
    numberPerPage := int32(56) // int32 | Number of Rule Sets per page (optional)
    filter := "filter_example" // string | Search for Rule Set with name matching filter text (optional)
    name := "name_example" // string | Get a specific Rule Set by name (case-sensitive) (optional)
    sortKey := "sortKey_example" // string | A comma separated list of Rule Set attributes (keys) to be used to sort the results (optional)
    order := "order_example" // string | Order of the sorted results (ASC for ascending, DESC for descending) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesetsApi.GetRuleSets(context.Background()).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesetsApi.GetRuleSets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRuleSets`: RuleSets
    fmt.Fprintf(os.Stdout, "Response from `RulesetsApi.GetRuleSets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleSetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number to retrieve | 
 **numberPerPage** | **int32** | Number of Rule Sets per page | 
 **filter** | **string** | Search for Rule Set with name matching filter text | 
 **name** | **string** | Get a specific Rule Set by name (case-sensitive) | 
 **sortKey** | **string** | A comma separated list of Rule Set attributes (keys) to be used to sort the results | 
 **order** | **string** | Order of the sorted results (ASC for ascending, DESC for descending) | 

### Return type

[**RuleSets**](RuleSets.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRuleSet

> RuleSet UpdateRuleSet(ctx, id).RuleSet(ruleSet).Execute()





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
    id := "id_example" // string | ID of Rule Set to update
    ruleSet := *openapiclient.NewRuleSet("Name_example", []int32{int32(123)}) // RuleSet | Rule Set to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesetsApi.UpdateRuleSet(context.Background(), id).RuleSet(ruleSet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesetsApi.UpdateRuleSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRuleSet`: RuleSet
    fmt.Fprintf(os.Stdout, "Response from `RulesetsApi.UpdateRuleSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Rule Set to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRuleSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ruleSet** | [**RuleSet**](RuleSet.md) | Rule Set to update | 

### Return type

[**RuleSet**](RuleSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

