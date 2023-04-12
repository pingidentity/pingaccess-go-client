# \AuthTokenManagementApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAuthTokenManagement**](AuthTokenManagementApi.md#DeleteAuthTokenManagement) | **Delete** /authTokenManagement | 
[**GetAuthTokenKeySet**](AuthTokenManagementApi.md#GetAuthTokenKeySet) | **Get** /authTokenManagement/keySet | 
[**GetAuthTokenManagement**](AuthTokenManagementApi.md#GetAuthTokenManagement) | **Get** /authTokenManagement | 
[**UpdateAuthTokenKeySet**](AuthTokenManagementApi.md#UpdateAuthTokenKeySet) | **Put** /authTokenManagement/keySet | 
[**UpdateAuthTokenManagement**](AuthTokenManagementApi.md#UpdateAuthTokenManagement) | **Put** /authTokenManagement | 



## DeleteAuthTokenManagement

> DeleteAuthTokenManagement(ctx).Execute()





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
    r, err := apiClient.AuthTokenManagementApi.DeleteAuthTokenManagement(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthTokenManagementApi.DeleteAuthTokenManagement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthTokenManagementRequest struct via the builder pattern


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


## GetAuthTokenKeySet

> KeySet GetAuthTokenKeySet(ctx).Execute()





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
    resp, r, err := apiClient.AuthTokenManagementApi.GetAuthTokenKeySet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthTokenManagementApi.GetAuthTokenKeySet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthTokenKeySet`: KeySet
    fmt.Fprintf(os.Stdout, "Response from `AuthTokenManagementApi.GetAuthTokenKeySet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthTokenKeySetRequest struct via the builder pattern


### Return type

[**KeySet**](KeySet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthTokenManagement

> AuthTokenManagement GetAuthTokenManagement(ctx).Execute()





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
    resp, r, err := apiClient.AuthTokenManagementApi.GetAuthTokenManagement(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthTokenManagementApi.GetAuthTokenManagement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthTokenManagement`: AuthTokenManagement
    fmt.Fprintf(os.Stdout, "Response from `AuthTokenManagementApi.GetAuthTokenManagement`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthTokenManagementRequest struct via the builder pattern


### Return type

[**AuthTokenManagement**](AuthTokenManagement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAuthTokenKeySet

> KeySet UpdateAuthTokenKeySet(ctx).KeySet(keySet).Execute()





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
    keySet := *openapiclient.NewKeySet("Nonce_example", "KeySet_example") // KeySet | AuthToken key set configuration to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthTokenManagementApi.UpdateAuthTokenKeySet(context.Background()).KeySet(keySet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthTokenManagementApi.UpdateAuthTokenKeySet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAuthTokenKeySet`: KeySet
    fmt.Fprintf(os.Stdout, "Response from `AuthTokenManagementApi.UpdateAuthTokenKeySet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAuthTokenKeySetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keySet** | [**KeySet**](KeySet.md) | AuthToken key set configuration to update | 

### Return type

[**KeySet**](KeySet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAuthTokenManagement

> AuthTokenManagement UpdateAuthTokenManagement(ctx).AuthTokenManagement(authTokenManagement).Execute()





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
    authTokenManagement := *openapiclient.NewAuthTokenManagement() // AuthTokenManagement | AuthTokenManagement configuration to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthTokenManagementApi.UpdateAuthTokenManagement(context.Background()).AuthTokenManagement(authTokenManagement).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthTokenManagementApi.UpdateAuthTokenManagement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAuthTokenManagement`: AuthTokenManagement
    fmt.Fprintf(os.Stdout, "Response from `AuthTokenManagementApi.UpdateAuthTokenManagement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAuthTokenManagementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authTokenManagement** | [**AuthTokenManagement**](AuthTokenManagement.md) | AuthTokenManagement configuration to update | 

### Return type

[**AuthTokenManagement**](AuthTokenManagement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

