# \DefaultsApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApplicationDefaults**](DefaultsApi.md#DeleteApplicationDefaults) | **Delete** /defaults/entities/application | 
[**GetApplicationDefaults**](DefaultsApi.md#GetApplicationDefaults) | **Get** /defaults/entities/application | 
[**UpdateApplicationDefaults**](DefaultsApi.md#UpdateApplicationDefaults) | **Put** /defaults/entities/application | 



## DeleteApplicationDefaults

> DeleteApplicationDefaults(ctx).Execute()





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
    r, err := apiClient.DefaultsApi.DeleteApplicationDefaults(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultsApi.DeleteApplicationDefaults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationDefaultsRequest struct via the builder pattern


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


## GetApplicationDefaults

> ApplicationDefaults GetApplicationDefaults(ctx).Execute()





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
    resp, r, err := apiClient.DefaultsApi.GetApplicationDefaults(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultsApi.GetApplicationDefaults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationDefaults`: ApplicationDefaults
    fmt.Fprintf(os.Stdout, "Response from `DefaultsApi.GetApplicationDefaults`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationDefaultsRequest struct via the builder pattern


### Return type

[**ApplicationDefaults**](ApplicationDefaults.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationDefaults

> ApplicationDefaults UpdateApplicationDefaults(ctx).Defaults(defaults).Execute()





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
    defaults := *openapiclient.NewApplicationDefaults() // ApplicationDefaults | Application Defaults

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultsApi.UpdateApplicationDefaults(context.Background()).Defaults(defaults).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultsApi.UpdateApplicationDefaults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApplicationDefaults`: ApplicationDefaults
    fmt.Fprintf(os.Stdout, "Response from `DefaultsApi.UpdateApplicationDefaults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationDefaultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **defaults** | [**ApplicationDefaults**](ApplicationDefaults.md) | Application Defaults | 

### Return type

[**ApplicationDefaults**](ApplicationDefaults.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

