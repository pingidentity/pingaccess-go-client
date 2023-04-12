# \WebSessionManagementApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWebSessionManagement**](WebSessionManagementApi.md#DeleteWebSessionManagement) | **Delete** /webSessionManagement | 
[**GetCookieTypes**](WebSessionManagementApi.md#GetCookieTypes) | **Get** /webSessionManagement/cookieTypes | 
[**GetOidcLoginTypes**](WebSessionManagementApi.md#GetOidcLoginTypes) | **Get** /webSessionManagement/oidcLoginTypes | 
[**GetOidcScopes**](WebSessionManagementApi.md#GetOidcScopes) | **Get** /webSessionManagement/oidcScopes | 
[**GetRequestPreservationTypes**](WebSessionManagementApi.md#GetRequestPreservationTypes) | **Get** /webSessionManagement/requestPreservationTypes | 
[**GetWebSessionKeySet**](WebSessionManagementApi.md#GetWebSessionKeySet) | **Get** /webSessionManagement/keySet | 
[**GetWebSessionManagement**](WebSessionManagementApi.md#GetWebSessionManagement) | **Get** /webSessionManagement | 
[**GetWebSessionSupportedEncryptionAlgorithms**](WebSessionManagementApi.md#GetWebSessionSupportedEncryptionAlgorithms) | **Get** /webSessionManagement/encryptionAlgorithms | 
[**GetWebSessionSupportedSigningAlgorithms**](WebSessionManagementApi.md#GetWebSessionSupportedSigningAlgorithms) | **Get** /webSessionManagement/signingAlgorithms | 
[**GetWebStorageTypes**](WebSessionManagementApi.md#GetWebStorageTypes) | **Get** /webSessionManagement/webStorageTypes | 
[**UpdateWebSessionKeySet**](WebSessionManagementApi.md#UpdateWebSessionKeySet) | **Put** /webSessionManagement/keySet | 
[**UpdateWebSessionManagement**](WebSessionManagementApi.md#UpdateWebSessionManagement) | **Put** /webSessionManagement | 



## DeleteWebSessionManagement

> DeleteWebSessionManagement(ctx).Execute()





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
    r, err := apiClient.WebSessionManagementApi.DeleteWebSessionManagement(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebSessionManagementApi.DeleteWebSessionManagement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebSessionManagementRequest struct via the builder pattern


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


## GetCookieTypes

> CookieTypes GetCookieTypes(ctx).Execute()





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
    resp, r, err := apiClient.WebSessionManagementApi.GetCookieTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebSessionManagementApi.GetCookieTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCookieTypes`: CookieTypes
    fmt.Fprintf(os.Stdout, "Response from `WebSessionManagementApi.GetCookieTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCookieTypesRequest struct via the builder pattern


### Return type

[**CookieTypes**](CookieTypes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOidcLoginTypes

> OidcLoginTypes GetOidcLoginTypes(ctx).Execute()





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
    resp, r, err := apiClient.WebSessionManagementApi.GetOidcLoginTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebSessionManagementApi.GetOidcLoginTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOidcLoginTypes`: OidcLoginTypes
    fmt.Fprintf(os.Stdout, "Response from `WebSessionManagementApi.GetOidcLoginTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOidcLoginTypesRequest struct via the builder pattern


### Return type

[**OidcLoginTypes**](OidcLoginTypes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOidcScopes

> SupportedScopes GetOidcScopes(ctx).ClientId(clientId).Execute()





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
    clientId := "clientId_example" // string | ID of a specific client to retrieve supported scopes for (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebSessionManagementApi.GetOidcScopes(context.Background()).ClientId(clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebSessionManagementApi.GetOidcScopes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOidcScopes`: SupportedScopes
    fmt.Fprintf(os.Stdout, "Response from `WebSessionManagementApi.GetOidcScopes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOidcScopesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** | ID of a specific client to retrieve supported scopes for | 

### Return type

[**SupportedScopes**](SupportedScopes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRequestPreservationTypes

> RequestPreservationTypes GetRequestPreservationTypes(ctx).Execute()





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
    resp, r, err := apiClient.WebSessionManagementApi.GetRequestPreservationTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebSessionManagementApi.GetRequestPreservationTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRequestPreservationTypes`: RequestPreservationTypes
    fmt.Fprintf(os.Stdout, "Response from `WebSessionManagementApi.GetRequestPreservationTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestPreservationTypesRequest struct via the builder pattern


### Return type

[**RequestPreservationTypes**](RequestPreservationTypes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebSessionKeySet

> KeySet GetWebSessionKeySet(ctx).Execute()





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
    resp, r, err := apiClient.WebSessionManagementApi.GetWebSessionKeySet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebSessionManagementApi.GetWebSessionKeySet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebSessionKeySet`: KeySet
    fmt.Fprintf(os.Stdout, "Response from `WebSessionManagementApi.GetWebSessionKeySet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebSessionKeySetRequest struct via the builder pattern


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


## GetWebSessionManagement

> WebSessionManagement GetWebSessionManagement(ctx).Execute()





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
    resp, r, err := apiClient.WebSessionManagementApi.GetWebSessionManagement(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebSessionManagementApi.GetWebSessionManagement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebSessionManagement`: WebSessionManagement
    fmt.Fprintf(os.Stdout, "Response from `WebSessionManagementApi.GetWebSessionManagement`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebSessionManagementRequest struct via the builder pattern


### Return type

[**WebSessionManagement**](WebSessionManagement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebSessionSupportedEncryptionAlgorithms

> Algorithms GetWebSessionSupportedEncryptionAlgorithms(ctx).Execute()





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
    resp, r, err := apiClient.WebSessionManagementApi.GetWebSessionSupportedEncryptionAlgorithms(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebSessionManagementApi.GetWebSessionSupportedEncryptionAlgorithms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebSessionSupportedEncryptionAlgorithms`: Algorithms
    fmt.Fprintf(os.Stdout, "Response from `WebSessionManagementApi.GetWebSessionSupportedEncryptionAlgorithms`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebSessionSupportedEncryptionAlgorithmsRequest struct via the builder pattern


### Return type

[**Algorithms**](Algorithms.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebSessionSupportedSigningAlgorithms

> SigningAlgorithms GetWebSessionSupportedSigningAlgorithms(ctx).Execute()





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
    resp, r, err := apiClient.WebSessionManagementApi.GetWebSessionSupportedSigningAlgorithms(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebSessionManagementApi.GetWebSessionSupportedSigningAlgorithms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebSessionSupportedSigningAlgorithms`: SigningAlgorithms
    fmt.Fprintf(os.Stdout, "Response from `WebSessionManagementApi.GetWebSessionSupportedSigningAlgorithms`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebSessionSupportedSigningAlgorithmsRequest struct via the builder pattern


### Return type

[**SigningAlgorithms**](SigningAlgorithms.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebStorageTypes

> WebStorageTypes GetWebStorageTypes(ctx).Execute()





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
    resp, r, err := apiClient.WebSessionManagementApi.GetWebStorageTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebSessionManagementApi.GetWebStorageTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebStorageTypes`: WebStorageTypes
    fmt.Fprintf(os.Stdout, "Response from `WebSessionManagementApi.GetWebStorageTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebStorageTypesRequest struct via the builder pattern


### Return type

[**WebStorageTypes**](WebStorageTypes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebSessionKeySet

> KeySet UpdateWebSessionKeySet(ctx).KeySet(keySet).Execute()





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
    keySet := *openapiclient.NewKeySet("Nonce_example", "KeySet_example") // KeySet | WebSession key set configuration to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebSessionManagementApi.UpdateWebSessionKeySet(context.Background()).KeySet(keySet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebSessionManagementApi.UpdateWebSessionKeySet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWebSessionKeySet`: KeySet
    fmt.Fprintf(os.Stdout, "Response from `WebSessionManagementApi.UpdateWebSessionKeySet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebSessionKeySetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keySet** | [**KeySet**](KeySet.md) | WebSession key set configuration to update | 

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


## UpdateWebSessionManagement

> WebSessionManagement UpdateWebSessionManagement(ctx).WebSessionManagement(webSessionManagement).Execute()





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
    webSessionManagement := *openapiclient.NewWebSessionManagement() // WebSessionManagement | WebSessionManagement configuration to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebSessionManagementApi.UpdateWebSessionManagement(context.Background()).WebSessionManagement(webSessionManagement).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebSessionManagementApi.UpdateWebSessionManagement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWebSessionManagement`: WebSessionManagement
    fmt.Fprintf(os.Stdout, "Response from `WebSessionManagementApi.UpdateWebSessionManagement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebSessionManagementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webSessionManagement** | [**WebSessionManagement**](WebSessionManagement.md) | WebSessionManagement configuration to update | 

### Return type

[**WebSessionManagement**](WebSessionManagement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

