# \TokenProviderApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTokenProviderSetting**](TokenProviderApi.md#DeleteTokenProviderSetting) | **Delete** /tokenProvider/settings | 
[**GetTokenProviderSetting**](TokenProviderApi.md#GetTokenProviderSetting) | **Get** /tokenProvider/settings | 
[**UpdateTokenProviderSetting**](TokenProviderApi.md#UpdateTokenProviderSetting) | **Put** /tokenProvider/settings | 



## DeleteTokenProviderSetting

> DeleteTokenProviderSetting(ctx).Execute()





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
    r, err := apiClient.TokenProviderApi.DeleteTokenProviderSetting(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenProviderApi.DeleteTokenProviderSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTokenProviderSettingRequest struct via the builder pattern


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


## GetTokenProviderSetting

> TokenProviderSetting GetTokenProviderSetting(ctx).Execute()





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
    resp, r, err := apiClient.TokenProviderApi.GetTokenProviderSetting(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenProviderApi.GetTokenProviderSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokenProviderSetting`: TokenProviderSetting
    fmt.Fprintf(os.Stdout, "Response from `TokenProviderApi.GetTokenProviderSetting`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenProviderSettingRequest struct via the builder pattern


### Return type

[**TokenProviderSetting**](TokenProviderSetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTokenProviderSetting

> TokenProviderSetting UpdateTokenProviderSetting(ctx).TokenProviderSettings(tokenProviderSettings).Execute()





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
    tokenProviderSettings := *openapiclient.NewTokenProviderSetting() // TokenProviderSetting | Token Provider Setting

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenProviderApi.UpdateTokenProviderSetting(context.Background()).TokenProviderSettings(tokenProviderSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenProviderApi.UpdateTokenProviderSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTokenProviderSetting`: TokenProviderSetting
    fmt.Fprintf(os.Stdout, "Response from `TokenProviderApi.UpdateTokenProviderSetting`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTokenProviderSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenProviderSettings** | [**TokenProviderSetting**](TokenProviderSetting.md) | Token Provider Setting | 

### Return type

[**TokenProviderSetting**](TokenProviderSetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

