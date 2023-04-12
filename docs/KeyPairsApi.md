# \KeyPairsApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteChainCertificate**](KeyPairsApi.md#DeleteChainCertificate) | **Delete** /keyPairs/{keyPairId}/chainCertificates/{chainCertificateId} | 
[**DeleteKeyPair**](KeyPairsApi.md#DeleteKeyPair) | **Delete** /keyPairs/{id} | 
[**ExportKeyPairCert**](KeyPairsApi.md#ExportKeyPairCert) | **Get** /keyPairs/{id}/certificate | 
[**ExportKeyPairPem**](KeyPairsApi.md#ExportKeyPairPem) | **Post** /keyPairs/{id}/pem | 
[**ExportKeyPairPkcs12**](KeyPairsApi.md#ExportKeyPairPkcs12) | **Post** /keyPairs/{id}/pkcs12 | 
[**GenerateCsr**](KeyPairsApi.md#GenerateCsr) | **Get** /keyPairs/{id}/csr | 
[**GenerateKeyPair**](KeyPairsApi.md#GenerateKeyPair) | **Post** /keyPairs/generate | 
[**GetKeyPair**](KeyPairsApi.md#GetKeyPair) | **Get** /keyPairs/{id} | 
[**GetKeyPairs**](KeyPairsApi.md#GetKeyPairs) | **Get** /keyPairs | 
[**GetKeypairsCreatableGeneralNames**](KeyPairsApi.md#GetKeypairsCreatableGeneralNames) | **Get** /keyPairs/subjectAlternativeTypes | 
[**ImportKeyPair**](KeyPairsApi.md#ImportKeyPair) | **Post** /keyPairs/import | 
[**KeyAlgorithms**](KeyPairsApi.md#KeyAlgorithms) | **Get** /keyPairs/keyAlgorithms | 
[**PatchKeyPair**](KeyPairsApi.md#PatchKeyPair) | **Patch** /keyPairs/{id} | 
[**PostImportCSRResponse**](KeyPairsApi.md#PostImportCSRResponse) | **Post** /keyPairs/{id}/csr | 
[**PutImportCSRResponse**](KeyPairsApi.md#PutImportCSRResponse) | **Put** /keyPairs/{id}/csr | 
[**UpdateKeyPair**](KeyPairsApi.md#UpdateKeyPair) | **Put** /keyPairs/{id} | 



## DeleteChainCertificate

> DeleteChainCertificate(ctx, keyPairId, chainCertificateId).Execute()





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
    keyPairId := "keyPairId_example" // string | ID of Key Pair
    chainCertificateId := "chainCertificateId_example" // string | ID of the Chain Certificate to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KeyPairsApi.DeleteChainCertificate(context.Background(), keyPairId, chainCertificateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsApi.DeleteChainCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyPairId** | **string** | ID of Key Pair | 
**chainCertificateId** | **string** | ID of the Chain Certificate to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteChainCertificateRequest struct via the builder pattern


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


## DeleteKeyPair

> DeleteKeyPair(ctx, id).Execute()





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
    id := "id_example" // string | ID of the Key Pair to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KeyPairsApi.DeleteKeyPair(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsApi.DeleteKeyPair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Key Pair to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKeyPairRequest struct via the builder pattern


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


## ExportKeyPairCert

> ExportKeyPairCert(ctx, id).Execute()





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
    id := "id_example" // string | ID of the Key Pair Certificate to export

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KeyPairsApi.ExportKeyPairCert(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsApi.ExportKeyPairCert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Key Pair Certificate to export | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportKeyPairCertRequest struct via the builder pattern


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


## ExportKeyPairPem

> ExportKeyPairPem(ctx, id).ExportParameters(exportParameters).Execute()





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
    id := "id_example" // string | ID of the Key Pair to export
    exportParameters := *openapiclient.NewExportParameters(*openapiclient.NewHiddenField()) // ExportParameters | Parameters for the export

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KeyPairsApi.ExportKeyPairPem(context.Background(), id).ExportParameters(exportParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsApi.ExportKeyPairPem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Key Pair to export | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportKeyPairPemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exportParameters** | [**ExportParameters**](ExportParameters.md) | Parameters for the export | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportKeyPairPkcs12

> ExportKeyPairPkcs12(ctx, id).ExportParameters(exportParameters).Execute()





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
    id := "id_example" // string | ID of the Key Pair to export
    exportParameters := *openapiclient.NewExportParameters(*openapiclient.NewHiddenField()) // ExportParameters | Parameters for the export

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KeyPairsApi.ExportKeyPairPkcs12(context.Background(), id).ExportParameters(exportParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsApi.ExportKeyPairPkcs12``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Key Pair to export | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportKeyPairPkcs12Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exportParameters** | [**ExportParameters**](ExportParameters.md) | Parameters for the export | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateCsr

> GenerateCsr(ctx, id).Execute()





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
    id := "id_example" // string | ID of the Key Pair to sign

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KeyPairsApi.GenerateCsr(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsApi.GenerateCsr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Key Pair to sign | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateCsrRequest struct via the builder pattern


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


## GenerateKeyPair

> KeyPair GenerateKeyPair(ctx).NewKeyPairConfigView(newKeyPairConfigView).Execute()





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
    newKeyPairConfigView := *openapiclient.NewNewKeyPairConfig("Alias_example", "CommonName_example", "Organization_example", "OrganizationUnit_example", "City_example", "State_example", "Country_example", int32(123), "KeyAlgorithm_example", int32(123), int32(123)) // NewKeyPairConfig | Configuration for the new Key Pair

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyPairsApi.GenerateKeyPair(context.Background()).NewKeyPairConfigView(newKeyPairConfigView).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsApi.GenerateKeyPair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateKeyPair`: KeyPair
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsApi.GenerateKeyPair`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateKeyPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newKeyPairConfigView** | [**NewKeyPairConfig**](NewKeyPairConfig.md) | Configuration for the new Key Pair | 

### Return type

[**KeyPair**](KeyPair.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKeyPair

> KeyPair GetKeyPair(ctx, id).Execute()





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
    id := "id_example" // string | ID of Key Pair to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyPairsApi.GetKeyPair(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsApi.GetKeyPair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKeyPair`: KeyPair
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsApi.GetKeyPair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Key Pair to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeyPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KeyPair**](KeyPair.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKeyPairs

> KeyPairs GetKeyPairs(ctx).Page(page).NumberPerPage(numberPerPage).Filter(filter).Alias(alias).SortKey(sortKey).Order(order).Execute()





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
    numberPerPage := int32(56) // int32 | Number of Key Pairs per page (optional)
    filter := "filter_example" // string | Search for Key Pairs with alias matching filter text (optional)
    alias := "alias_example" // string | Get a Key Pair by alias (optional)
    sortKey := "sortKey_example" // string | A comma separated list of Key Pair attributes (keys) to be used to sort the results (optional)
    order := "order_example" // string | Order of the sorted results (ASC for ascending, DESC for descending) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyPairsApi.GetKeyPairs(context.Background()).Page(page).NumberPerPage(numberPerPage).Filter(filter).Alias(alias).SortKey(sortKey).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsApi.GetKeyPairs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKeyPairs`: KeyPairs
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsApi.GetKeyPairs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKeyPairsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number to retrieve | 
 **numberPerPage** | **int32** | Number of Key Pairs per page | 
 **filter** | **string** | Search for Key Pairs with alias matching filter text | 
 **alias** | **string** | Get a Key Pair by alias | 
 **sortKey** | **string** | A comma separated list of Key Pair attributes (keys) to be used to sort the results | 
 **order** | **string** | Order of the sorted results (ASC for ascending, DESC for descending) | 

### Return type

[**KeyPairs**](KeyPairs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKeypairsCreatableGeneralNames

> SanTypes GetKeypairsCreatableGeneralNames(ctx).Execute()





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
    resp, r, err := apiClient.KeyPairsApi.GetKeypairsCreatableGeneralNames(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsApi.GetKeypairsCreatableGeneralNames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKeypairsCreatableGeneralNames`: SanTypes
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsApi.GetKeypairsCreatableGeneralNames`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeypairsCreatableGeneralNamesRequest struct via the builder pattern


### Return type

[**SanTypes**](SanTypes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportKeyPair

> KeyPair ImportKeyPair(ctx).KeyPairFile(keyPairFile).Execute()





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
    keyPairFile := *openapiclient.NewPKCS12FileImportDoc("Alias_example", int32(123), "FileData_example", *openapiclient.NewHiddenField(), []string{"ChainCertificates_example"}) // PKCS12FileImportDoc | Key Pair to import

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyPairsApi.ImportKeyPair(context.Background()).KeyPairFile(keyPairFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsApi.ImportKeyPair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportKeyPair`: KeyPair
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsApi.ImportKeyPair`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportKeyPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keyPairFile** | [**PKCS12FileImportDoc**](PKCS12FileImportDoc.md) | Key Pair to import | 

### Return type

[**KeyPair**](KeyPair.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeyAlgorithms

> KeyAlgorithms KeyAlgorithms(ctx).Execute()





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
    resp, r, err := apiClient.KeyPairsApi.KeyAlgorithms(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsApi.KeyAlgorithms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KeyAlgorithms`: KeyAlgorithms
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsApi.KeyAlgorithms`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiKeyAlgorithmsRequest struct via the builder pattern


### Return type

[**KeyAlgorithms**](KeyAlgorithms.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchKeyPair

> KeyPair PatchKeyPair(ctx, id).AddChainCertificates(addChainCertificates).Execute()





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
    id := "id_example" // string | ID of Key Pair to update
    addChainCertificates := *openapiclient.NewChainCertificatesDoc([]string{"AddChainCertificates_example"}) // ChainCertificatesDoc | A list of base64-encoded Chain Certificates

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyPairsApi.PatchKeyPair(context.Background(), id).AddChainCertificates(addChainCertificates).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsApi.PatchKeyPair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchKeyPair`: KeyPair
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsApi.PatchKeyPair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Key Pair to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchKeyPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addChainCertificates** | [**ChainCertificatesDoc**](ChainCertificatesDoc.md) | A list of base64-encoded Chain Certificates | 

### Return type

[**KeyPair**](KeyPair.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostImportCSRResponse

> KeyPair PostImportCSRResponse(ctx, id).CSRResponse(cSRResponse).Execute()





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
    id := "id_example" // string | ID of the Key Pair to update
    cSRResponse := *openapiclient.NewCSRResponseImportDoc(int32(123), "FileData_example", []string{"ChainCertificates_example"}, int32(123)) // CSRResponseImportDoc | The Certificate Signing Request (CSR) response

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyPairsApi.PostImportCSRResponse(context.Background(), id).CSRResponse(cSRResponse).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsApi.PostImportCSRResponse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostImportCSRResponse`: KeyPair
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsApi.PostImportCSRResponse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Key Pair to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostImportCSRResponseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cSRResponse** | [**CSRResponseImportDoc**](CSRResponseImportDoc.md) | The Certificate Signing Request (CSR) response | 

### Return type

[**KeyPair**](KeyPair.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutImportCSRResponse

> KeyPair PutImportCSRResponse(ctx, id).CSRResponse(cSRResponse).Execute()





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
    id := "id_example" // string | ID of the Key Pair to update
    cSRResponse := *openapiclient.NewCSRResponseImportDoc(int32(123), "FileData_example", []string{"ChainCertificates_example"}, int32(123)) // CSRResponseImportDoc | The Certificate Signing Request (CSR) response

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyPairsApi.PutImportCSRResponse(context.Background(), id).CSRResponse(cSRResponse).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsApi.PutImportCSRResponse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutImportCSRResponse`: KeyPair
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsApi.PutImportCSRResponse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Key Pair to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutImportCSRResponseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cSRResponse** | [**CSRResponseImportDoc**](CSRResponseImportDoc.md) | The Certificate Signing Request (CSR) response | 

### Return type

[**KeyPair**](KeyPair.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKeyPair

> KeyPair UpdateKeyPair(ctx, id).PKCS12File(pKCS12File).Execute()





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
    id := "id_example" // string | ID of Key Pair to update
    pKCS12File := *openapiclient.NewPKCS12FileImportDoc("Alias_example", int32(123), "FileData_example", *openapiclient.NewHiddenField(), []string{"ChainCertificates_example"}) // PKCS12FileImportDoc | PKCS12 file containing the Key Pair

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyPairsApi.UpdateKeyPair(context.Background(), id).PKCS12File(pKCS12File).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsApi.UpdateKeyPair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateKeyPair`: KeyPair
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsApi.UpdateKeyPair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Key Pair to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKeyPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pKCS12File** | [**PKCS12FileImportDoc**](PKCS12FileImportDoc.md) | PKCS12 file containing the Key Pair | 

### Return type

[**KeyPair**](KeyPair.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

