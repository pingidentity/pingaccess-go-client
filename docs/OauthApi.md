# \OauthApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAuthorizationServer**](OauthApi.md#DeleteAuthorizationServer) | **Delete** /oauth/authServer | 
[**GetAuthorizationServer**](OauthApi.md#GetAuthorizationServer) | **Get** /oauth/authServer | 
[**UpdateAuthorizationServer**](OauthApi.md#UpdateAuthorizationServer) | **Put** /oauth/authServer | 



## DeleteAuthorizationServer

> DeleteAuthorizationServer(ctx).Execute()





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
    r, err := apiClient.OauthApi.DeleteAuthorizationServer(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthApi.DeleteAuthorizationServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthorizationServerRequest struct via the builder pattern


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


## GetAuthorizationServer

> AuthorizationServer GetAuthorizationServer(ctx).Execute()





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
    resp, r, err := apiClient.OauthApi.GetAuthorizationServer(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthApi.GetAuthorizationServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthorizationServer`: AuthorizationServer
    fmt.Fprintf(os.Stdout, "Response from `OauthApi.GetAuthorizationServer`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthorizationServerRequest struct via the builder pattern


### Return type

[**AuthorizationServer**](AuthorizationServer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAuthorizationServer

> AuthorizationServer UpdateAuthorizationServer(ctx).OpenIDConnectProvider(openIDConnectProvider).Execute()





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
    openIDConnectProvider := *openapiclient.NewAuthorizationServer([]string{"Targets_example"}, int64(123), "SubjectAttributeName_example", "IntrospectionEndpoint_example") // AuthorizationServer | OAuth 2.0 Authorization Server configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthApi.UpdateAuthorizationServer(context.Background()).OpenIDConnectProvider(openIDConnectProvider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthApi.UpdateAuthorizationServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAuthorizationServer`: AuthorizationServer
    fmt.Fprintf(os.Stdout, "Response from `OauthApi.UpdateAuthorizationServer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAuthorizationServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **openIDConnectProvider** | [**AuthorizationServer**](AuthorizationServer.md) | OAuth 2.0 Authorization Server configuration | 

### Return type

[**AuthorizationServer**](AuthorizationServer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

