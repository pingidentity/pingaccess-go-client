# \LicenseApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLicense**](LicenseApi.md#GetLicense) | **Get** /license | 
[**PostImportLicense**](LicenseApi.md#PostImportLicense) | **Post** /license | 
[**PutImportLicense**](LicenseApi.md#PutImportLicense) | **Put** /license | 



## GetLicense

> License GetLicense(ctx).Execute()





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
    resp, r, err := apiClient.LicenseApi.GetLicense(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseApi.GetLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLicense`: License
    fmt.Fprintf(os.Stdout, "Response from `LicenseApi.GetLicense`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseRequest struct via the builder pattern


### Return type

[**License**](License.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostImportLicense

> License PostImportLicense(ctx).LicenseFile(licenseFile).Execute()





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
    licenseFile := *openapiclient.NewLicenseImportDoc("FileData_example") // LicenseImportDoc | License to import

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicenseApi.PostImportLicense(context.Background()).LicenseFile(licenseFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseApi.PostImportLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostImportLicense`: License
    fmt.Fprintf(os.Stdout, "Response from `LicenseApi.PostImportLicense`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostImportLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **licenseFile** | [**LicenseImportDoc**](LicenseImportDoc.md) | License to import | 

### Return type

[**License**](License.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutImportLicense

> License PutImportLicense(ctx).LicenseFile(licenseFile).Execute()





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
    licenseFile := *openapiclient.NewLicenseImportDoc("FileData_example") // LicenseImportDoc | License to import

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicenseApi.PutImportLicense(context.Background()).LicenseFile(licenseFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseApi.PutImportLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutImportLicense`: License
    fmt.Fprintf(os.Stdout, "Response from `LicenseApi.PutImportLicense`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutImportLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **licenseFile** | [**LicenseImportDoc**](LicenseImportDoc.md) | License to import | 

### Return type

[**License**](License.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

