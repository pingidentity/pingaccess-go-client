# \AccessTokenValidatorsApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAccessTokenValidator**](AccessTokenValidatorsApi.md#AddAccessTokenValidator) | **Post** /accessTokenValidators | 
[**DeleteAccessTokenValidator**](AccessTokenValidatorsApi.md#DeleteAccessTokenValidator) | **Delete** /accessTokenValidators/{id} | 
[**GetAccessTokenValidator**](AccessTokenValidatorsApi.md#GetAccessTokenValidator) | **Get** /accessTokenValidators/{id} | 
[**GetAccessTokenValidatorDescriptors**](AccessTokenValidatorsApi.md#GetAccessTokenValidatorDescriptors) | **Get** /accessTokenValidators/descriptors | 
[**GetAccessTokenValidators**](AccessTokenValidatorsApi.md#GetAccessTokenValidators) | **Get** /accessTokenValidators | 
[**UpdateAccessTokenValidator**](AccessTokenValidatorsApi.md#UpdateAccessTokenValidator) | **Put** /accessTokenValidators/{id} | 



## AddAccessTokenValidator

> AccessTokenValidator AddAccessTokenValidator(ctx).AccessTokenValidator(accessTokenValidator).Execute()





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
    accessTokenValidator := *openapiclient.NewAccessTokenValidator("ClassName_example", "Name_example") // AccessTokenValidator | Access Token Validator to add

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessTokenValidatorsApi.AddAccessTokenValidator(context.Background()).AccessTokenValidator(accessTokenValidator).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessTokenValidatorsApi.AddAccessTokenValidator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAccessTokenValidator`: AccessTokenValidator
    fmt.Fprintf(os.Stdout, "Response from `AccessTokenValidatorsApi.AddAccessTokenValidator`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAccessTokenValidatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessTokenValidator** | [**AccessTokenValidator**](AccessTokenValidator.md) | Access Token Validator to add | 

### Return type

[**AccessTokenValidator**](AccessTokenValidator.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccessTokenValidator

> DeleteAccessTokenValidator(ctx, id).Execute()





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
    id := "id_example" // string | ID of Access Token Validator to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AccessTokenValidatorsApi.DeleteAccessTokenValidator(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessTokenValidatorsApi.DeleteAccessTokenValidator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Access Token Validator to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccessTokenValidatorRequest struct via the builder pattern


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


## GetAccessTokenValidator

> AccessTokenValidator GetAccessTokenValidator(ctx, id).Execute()





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
    id := "id_example" // string | ID of Access Token Validator to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessTokenValidatorsApi.GetAccessTokenValidator(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessTokenValidatorsApi.GetAccessTokenValidator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccessTokenValidator`: AccessTokenValidator
    fmt.Fprintf(os.Stdout, "Response from `AccessTokenValidatorsApi.GetAccessTokenValidator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Access Token Validator to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessTokenValidatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccessTokenValidator**](AccessTokenValidator.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessTokenValidatorDescriptors

> Descriptors GetAccessTokenValidatorDescriptors(ctx).Execute()





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
    resp, r, err := apiClient.AccessTokenValidatorsApi.GetAccessTokenValidatorDescriptors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessTokenValidatorsApi.GetAccessTokenValidatorDescriptors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccessTokenValidatorDescriptors`: Descriptors
    fmt.Fprintf(os.Stdout, "Response from `AccessTokenValidatorsApi.GetAccessTokenValidatorDescriptors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessTokenValidatorDescriptorsRequest struct via the builder pattern


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


## GetAccessTokenValidators

> AccessTokenValidators GetAccessTokenValidators(ctx).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()





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
    numberPerPage := int64(56) // int64 | Number of Access Token Validators per page (optional)
    filter := "filter_example" // string | Search for Access Token Validator with name matching filter text (optional)
    name := "name_example" // string | Get a specific Access Token Validator by name (case-sensitive) (optional)
    sortKey := "sortKey_example" // string | A comma separated list of Access Token Validator attributes (keys) to be used to sort the results (optional)
    order := "order_example" // string | Order of the sorted results (ASC for ascending, DESC for descending) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessTokenValidatorsApi.GetAccessTokenValidators(context.Background()).Page(page).NumberPerPage(numberPerPage).Filter(filter).Name(name).SortKey(sortKey).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessTokenValidatorsApi.GetAccessTokenValidators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccessTokenValidators`: AccessTokenValidators
    fmt.Fprintf(os.Stdout, "Response from `AccessTokenValidatorsApi.GetAccessTokenValidators`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessTokenValidatorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | Page number to retrieve | 
 **numberPerPage** | **int64** | Number of Access Token Validators per page | 
 **filter** | **string** | Search for Access Token Validator with name matching filter text | 
 **name** | **string** | Get a specific Access Token Validator by name (case-sensitive) | 
 **sortKey** | **string** | A comma separated list of Access Token Validator attributes (keys) to be used to sort the results | 
 **order** | **string** | Order of the sorted results (ASC for ascending, DESC for descending) | 

### Return type

[**AccessTokenValidators**](AccessTokenValidators.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccessTokenValidator

> AccessTokenValidator UpdateAccessTokenValidator(ctx, id).AccessTokenValidator(accessTokenValidator).Execute()





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
    id := "id_example" // string | ID of Access Token Validator to update
    accessTokenValidator := *openapiclient.NewAccessTokenValidator("ClassName_example", "Name_example") // AccessTokenValidator | Access Token Validator to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessTokenValidatorsApi.UpdateAccessTokenValidator(context.Background(), id).AccessTokenValidator(accessTokenValidator).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessTokenValidatorsApi.UpdateAccessTokenValidator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccessTokenValidator`: AccessTokenValidator
    fmt.Fprintf(os.Stdout, "Response from `AccessTokenValidatorsApi.UpdateAccessTokenValidator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Access Token Validator to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccessTokenValidatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessTokenValidator** | [**AccessTokenValidator**](AccessTokenValidator.md) | Access Token Validator to update | 

### Return type

[**AccessTokenValidator**](AccessTokenValidator.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

