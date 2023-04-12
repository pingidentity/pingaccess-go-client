# \OauthKeyManagementApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOAuthKeyManagement**](OauthKeyManagementApi.md#DeleteOAuthKeyManagement) | **Delete** /oauthKeyManagement | 
[**GetOAuthKeyManagement**](OauthKeyManagementApi.md#GetOAuthKeyManagement) | **Get** /oauthKeyManagement | 
[**GetOAuthKeySet**](OauthKeyManagementApi.md#GetOAuthKeySet) | **Get** /oauthKeyManagement/keySet | 
[**UpdateOAuthKeyManagement**](OauthKeyManagementApi.md#UpdateOAuthKeyManagement) | **Put** /oauthKeyManagement | 
[**UpdateOAuthKeySet**](OauthKeyManagementApi.md#UpdateOAuthKeySet) | **Put** /oauthKeyManagement/keySet | 



## DeleteOAuthKeyManagement

> DeleteOAuthKeyManagement(ctx).Execute()





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
    r, err := apiClient.OauthKeyManagementApi.DeleteOAuthKeyManagement(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthKeyManagementApi.DeleteOAuthKeyManagement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOAuthKeyManagementRequest struct via the builder pattern


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


## GetOAuthKeyManagement

> OAuthKeyManagement GetOAuthKeyManagement(ctx).Execute()





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
    resp, r, err := apiClient.OauthKeyManagementApi.GetOAuthKeyManagement(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthKeyManagementApi.GetOAuthKeyManagement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOAuthKeyManagement`: OAuthKeyManagement
    fmt.Fprintf(os.Stdout, "Response from `OauthKeyManagementApi.GetOAuthKeyManagement`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOAuthKeyManagementRequest struct via the builder pattern


### Return type

[**OAuthKeyManagement**](OAuthKeyManagement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOAuthKeySet

> KeySet GetOAuthKeySet(ctx).Execute()





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
    resp, r, err := apiClient.OauthKeyManagementApi.GetOAuthKeySet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthKeyManagementApi.GetOAuthKeySet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOAuthKeySet`: KeySet
    fmt.Fprintf(os.Stdout, "Response from `OauthKeyManagementApi.GetOAuthKeySet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOAuthKeySetRequest struct via the builder pattern


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


## UpdateOAuthKeyManagement

> OAuthKeyManagement UpdateOAuthKeyManagement(ctx).OAuthKeyManagement(oAuthKeyManagement).Execute()





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
    oAuthKeyManagement := *openapiclient.NewOAuthKeyManagement() // OAuthKeyManagement | OAuth Key Management configuration to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthKeyManagementApi.UpdateOAuthKeyManagement(context.Background()).OAuthKeyManagement(oAuthKeyManagement).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthKeyManagementApi.UpdateOAuthKeyManagement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOAuthKeyManagement`: OAuthKeyManagement
    fmt.Fprintf(os.Stdout, "Response from `OauthKeyManagementApi.UpdateOAuthKeyManagement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOAuthKeyManagementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oAuthKeyManagement** | [**OAuthKeyManagement**](OAuthKeyManagement.md) | OAuth Key Management configuration to update | 

### Return type

[**OAuthKeyManagement**](OAuthKeyManagement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOAuthKeySet

> KeySet UpdateOAuthKeySet(ctx).KeySet(keySet).Execute()





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
    keySet := *openapiclient.NewKeySet("Nonce_example", "KeySet_example") // KeySet | OAuth key set configuration to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthKeyManagementApi.UpdateOAuthKeySet(context.Background()).KeySet(keySet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthKeyManagementApi.UpdateOAuthKeySet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOAuthKeySet`: KeySet
    fmt.Fprintf(os.Stdout, "Response from `OauthKeyManagementApi.UpdateOAuthKeySet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOAuthKeySetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keySet** | [**KeySet**](KeySet.md) | OAuth key set configuration to update | 

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

