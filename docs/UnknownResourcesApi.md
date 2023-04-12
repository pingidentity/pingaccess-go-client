# \UnknownResourcesApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUnknownResourceSettings**](UnknownResourcesApi.md#DeleteUnknownResourceSettings) | **Delete** /unknownResources/settings | 
[**GetUnknownResourceSettings**](UnknownResourcesApi.md#GetUnknownResourceSettings) | **Get** /unknownResources/settings | 
[**UpdateUnknownResourceSettings**](UnknownResourcesApi.md#UpdateUnknownResourceSettings) | **Put** /unknownResources/settings | 



## DeleteUnknownResourceSettings

> DeleteUnknownResourceSettings(ctx).Execute()





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
    r, err := apiClient.UnknownResourcesApi.DeleteUnknownResourceSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnknownResourcesApi.DeleteUnknownResourceSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUnknownResourceSettingsRequest struct via the builder pattern


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


## GetUnknownResourceSettings

> UnknownResourceSettings GetUnknownResourceSettings(ctx).Execute()





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
    resp, r, err := apiClient.UnknownResourcesApi.GetUnknownResourceSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnknownResourcesApi.GetUnknownResourceSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUnknownResourceSettings`: UnknownResourceSettings
    fmt.Fprintf(os.Stdout, "Response from `UnknownResourcesApi.GetUnknownResourceSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUnknownResourceSettingsRequest struct via the builder pattern


### Return type

[**UnknownResourceSettings**](UnknownResourceSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUnknownResourceSettings

> UnknownResourceSettings UpdateUnknownResourceSettings(ctx).Settings(settings).Execute()





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
    settings := *openapiclient.NewUnknownResourceSettings(int32(123), "ErrorTemplateFile_example", openapiclient.ContentType("JSON"), openapiclient.UnknownResourceMode("DENY"), int32(123)) // UnknownResourceSettings | The new settings

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnknownResourcesApi.UpdateUnknownResourceSettings(context.Background()).Settings(settings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnknownResourcesApi.UpdateUnknownResourceSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUnknownResourceSettings`: UnknownResourceSettings
    fmt.Fprintf(os.Stdout, "Response from `UnknownResourcesApi.UpdateUnknownResourceSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUnknownResourceSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **settings** | [**UnknownResourceSettings**](UnknownResourceSettings.md) | The new settings | 

### Return type

[**UnknownResourceSettings**](UnknownResourceSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

