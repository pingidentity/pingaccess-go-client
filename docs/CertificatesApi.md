# \CertificatesApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTrustedCert**](CertificatesApi.md#DeleteTrustedCert) | **Delete** /certificates/{id} | 
[**ExportTrustedCert**](CertificatesApi.md#ExportTrustedCert) | **Get** /certificates/{id}/file | 
[**GetTrustedCert**](CertificatesApi.md#GetTrustedCert) | **Get** /certificates/{id} | 
[**GetTrustedCerts**](CertificatesApi.md#GetTrustedCerts) | **Get** /certificates | 
[**ImportTrustedCert**](CertificatesApi.md#ImportTrustedCert) | **Post** /certificates | 
[**UpdateTrustedCert**](CertificatesApi.md#UpdateTrustedCert) | **Put** /certificates/{id} | 



## DeleteTrustedCert

> DeleteTrustedCert(ctx, id).Execute()





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
    id := "id_example" // string | ID of the Certificate to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CertificatesApi.DeleteTrustedCert(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesApi.DeleteTrustedCert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Certificate to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTrustedCertRequest struct via the builder pattern


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


## ExportTrustedCert

> ExportTrustedCert(ctx, id).Execute()





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
    id := "id_example" // string | ID of the Certificate to export

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CertificatesApi.ExportTrustedCert(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesApi.ExportTrustedCert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Certificate to export | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportTrustedCertRequest struct via the builder pattern


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


## GetTrustedCert

> TrustedCert GetTrustedCert(ctx, id).Execute()





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
    id := "id_example" // string | ID of Certificate to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificatesApi.GetTrustedCert(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesApi.GetTrustedCert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrustedCert`: TrustedCert
    fmt.Fprintf(os.Stdout, "Response from `CertificatesApi.GetTrustedCert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Certificate to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrustedCertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TrustedCert**](TrustedCert.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrustedCerts

> TrustedCerts GetTrustedCerts(ctx).Page(page).NumberPerPage(numberPerPage).Filter(filter).Alias(alias).SortKey(sortKey).Order(order).Execute()





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
    numberPerPage := int32(56) // int32 | Number of Trusted Certificates per page (optional)
    filter := "filter_example" // string | Search for Trusted Certificates with alias matching filter text (optional)
    alias := "alias_example" // string | Get a Trusted Certificate by alias (optional)
    sortKey := "sortKey_example" // string | A comma separated list of Trusted Certificate attributes (keys) to be used to sort the results (optional)
    order := "order_example" // string | Order of the sorted results (ASC for ascending, DESC for descending) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificatesApi.GetTrustedCerts(context.Background()).Page(page).NumberPerPage(numberPerPage).Filter(filter).Alias(alias).SortKey(sortKey).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesApi.GetTrustedCerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrustedCerts`: TrustedCerts
    fmt.Fprintf(os.Stdout, "Response from `CertificatesApi.GetTrustedCerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrustedCertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number to retrieve | 
 **numberPerPage** | **int32** | Number of Trusted Certificates per page | 
 **filter** | **string** | Search for Trusted Certificates with alias matching filter text | 
 **alias** | **string** | Get a Trusted Certificate by alias | 
 **sortKey** | **string** | A comma separated list of Trusted Certificate attributes (keys) to be used to sort the results | 
 **order** | **string** | Order of the sorted results (ASC for ascending, DESC for descending) | 

### Return type

[**TrustedCerts**](TrustedCerts.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportTrustedCert

> TrustedCert ImportTrustedCert(ctx).X509File(x509File).Execute()





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
    x509File := *openapiclient.NewX509FileImportDoc("Alias_example", "FileData_example") // X509FileImportDoc | Certificate to import

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificatesApi.ImportTrustedCert(context.Background()).X509File(x509File).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesApi.ImportTrustedCert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportTrustedCert`: TrustedCert
    fmt.Fprintf(os.Stdout, "Response from `CertificatesApi.ImportTrustedCert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportTrustedCertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **x509File** | [**X509FileImportDoc**](X509FileImportDoc.md) | Certificate to import | 

### Return type

[**TrustedCert**](TrustedCert.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTrustedCert

> TrustedCert UpdateTrustedCert(ctx, id).X509File(x509File).Execute()





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
    id := "id_example" // string | ID of Certificate to update
    x509File := *openapiclient.NewX509FileImportDoc("Alias_example", "FileData_example") // X509FileImportDoc | Certificate to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificatesApi.UpdateTrustedCert(context.Background(), id).X509File(x509File).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesApi.UpdateTrustedCert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTrustedCert`: TrustedCert
    fmt.Fprintf(os.Stdout, "Response from `CertificatesApi.UpdateTrustedCert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Certificate to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTrustedCertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **x509File** | [**X509FileImportDoc**](X509FileImportDoc.md) | Certificate to update | 

### Return type

[**TrustedCert**](TrustedCert.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

