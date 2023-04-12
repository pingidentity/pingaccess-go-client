# \VersionApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Version**](VersionApi.md#Version) | **Get** /version | 



## Version

> VersionDocClass Version(ctx).Execute()





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
    resp, r, err := apiClient.VersionApi.Version(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionApi.Version``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Version`: VersionDocClass
    fmt.Fprintf(os.Stdout, "Response from `VersionApi.Version`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVersionRequest struct via the builder pattern


### Return type

[**VersionDocClass**](VersionDocClass.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

