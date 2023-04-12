# \AdminConfigApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddReplicaAdmin**](AdminConfigApi.md#AddReplicaAdmin) | **Post** /adminConfig/replicaAdmins | 
[**DeleteAdminConfiguration**](AdminConfigApi.md#DeleteAdminConfiguration) | **Delete** /adminConfig | 
[**DeleteReplicaAdmin**](AdminConfigApi.md#DeleteReplicaAdmin) | **Delete** /adminConfig/replicaAdmins/{id} | 
[**GetAdminConfiguration**](AdminConfigApi.md#GetAdminConfiguration) | **Get** /adminConfig | 
[**GetAdminReplicaFile**](AdminConfigApi.md#GetAdminReplicaFile) | **Post** /adminConfig/replicaAdmins/{id}/config | 
[**GetReplicaAdmin**](AdminConfigApi.md#GetReplicaAdmin) | **Get** /adminConfig/replicaAdmins/{id} | 
[**GetReplicaAdmins**](AdminConfigApi.md#GetReplicaAdmins) | **Get** /adminConfig/replicaAdmins | 
[**UpdateAdminConfiguration**](AdminConfigApi.md#UpdateAdminConfiguration) | **Put** /adminConfig | 
[**UpdateAdminReplica**](AdminConfigApi.md#UpdateAdminReplica) | **Put** /adminConfig/replicaAdmins/{id} | 



## AddReplicaAdmin

> ReplicaAdmin AddReplicaAdmin(ctx).ReplicaAdmin(replicaAdmin).Execute()





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
    replicaAdmin := *openapiclient.NewReplicaAdmin("Name_example", "HostPort_example") // ReplicaAdmin | AdminReplica to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminConfigApi.AddReplicaAdmin(context.Background()).ReplicaAdmin(replicaAdmin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminConfigApi.AddReplicaAdmin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddReplicaAdmin`: ReplicaAdmin
    fmt.Fprintf(os.Stdout, "Response from `AdminConfigApi.AddReplicaAdmin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddReplicaAdminRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **replicaAdmin** | [**ReplicaAdmin**](ReplicaAdmin.md) | AdminReplica to create | 

### Return type

[**ReplicaAdmin**](ReplicaAdmin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAdminConfiguration

> DeleteAdminConfiguration(ctx).Execute()





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
    r, err := apiClient.AdminConfigApi.DeleteAdminConfiguration(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminConfigApi.DeleteAdminConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAdminConfigurationRequest struct via the builder pattern


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


## DeleteReplicaAdmin

> DeleteReplicaAdmin(ctx, id).Execute()





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
    id := "id_example" // string | ID of the replicaAdmin to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminConfigApi.DeleteReplicaAdmin(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminConfigApi.DeleteReplicaAdmin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the replicaAdmin to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReplicaAdminRequest struct via the builder pattern


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


## GetAdminConfiguration

> AdminConfiguration GetAdminConfiguration(ctx).Execute()





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
    resp, r, err := apiClient.AdminConfigApi.GetAdminConfiguration(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminConfigApi.GetAdminConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdminConfiguration`: AdminConfiguration
    fmt.Fprintf(os.Stdout, "Response from `AdminConfigApi.GetAdminConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminConfigurationRequest struct via the builder pattern


### Return type

[**AdminConfiguration**](AdminConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminReplicaFile

> GetAdminReplicaFile(ctx, id).Execute()





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
    id := "id_example" // string | ID of replicaAdmin to get configuration for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminConfigApi.GetAdminReplicaFile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminConfigApi.GetAdminReplicaFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of replicaAdmin to get configuration for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminReplicaFileRequest struct via the builder pattern


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


## GetReplicaAdmin

> ReplicaAdmin GetReplicaAdmin(ctx, id).Execute()





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
    id := "id_example" // string | ID of replicaAdmin to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminConfigApi.GetReplicaAdmin(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminConfigApi.GetReplicaAdmin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReplicaAdmin`: ReplicaAdmin
    fmt.Fprintf(os.Stdout, "Response from `AdminConfigApi.GetReplicaAdmin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of replicaAdmin to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReplicaAdminRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReplicaAdmin**](ReplicaAdmin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReplicaAdmins

> ReplicaAdmins GetReplicaAdmins(ctx).Execute()





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
    resp, r, err := apiClient.AdminConfigApi.GetReplicaAdmins(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminConfigApi.GetReplicaAdmins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReplicaAdmins`: ReplicaAdmins
    fmt.Fprintf(os.Stdout, "Response from `AdminConfigApi.GetReplicaAdmins`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetReplicaAdminsRequest struct via the builder pattern


### Return type

[**ReplicaAdmins**](ReplicaAdmins.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAdminConfiguration

> AdminConfiguration UpdateAdminConfiguration(ctx).AdminConfiguration(adminConfiguration).Execute()





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
    adminConfiguration := *openapiclient.NewAdminConfiguration("HostPort_example") // AdminConfiguration | Updated admin configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminConfigApi.UpdateAdminConfiguration(context.Background()).AdminConfiguration(adminConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminConfigApi.UpdateAdminConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAdminConfiguration`: AdminConfiguration
    fmt.Fprintf(os.Stdout, "Response from `AdminConfigApi.UpdateAdminConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAdminConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminConfiguration** | [**AdminConfiguration**](AdminConfiguration.md) | Updated admin configuration | 

### Return type

[**AdminConfiguration**](AdminConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAdminReplica

> ReplicaAdmin UpdateAdminReplica(ctx, id).ReplicaAdmin(replicaAdmin).Execute()





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
    id := "id_example" // string | ID of ReplicaAdmin to get
    replicaAdmin := *openapiclient.NewReplicaAdmin("Name_example", "HostPort_example") // ReplicaAdmin | ReplicaAdmin to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminConfigApi.UpdateAdminReplica(context.Background(), id).ReplicaAdmin(replicaAdmin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminConfigApi.UpdateAdminReplica``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAdminReplica`: ReplicaAdmin
    fmt.Fprintf(os.Stdout, "Response from `AdminConfigApi.UpdateAdminReplica`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of ReplicaAdmin to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAdminReplicaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replicaAdmin** | [**ReplicaAdmin**](ReplicaAdmin.md) | ReplicaAdmin to update | 

### Return type

[**ReplicaAdmin**](ReplicaAdmin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

