# \ConfigApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddConfigExportWorkflow**](ConfigApi.md#AddConfigExportWorkflow) | **Post** /config/export/workflows | 
[**AddConfigImportWorkflow**](ConfigApi.md#AddConfigImportWorkflow) | **Post** /config/import/workflows | 
[**ConfigExport**](ConfigApi.md#ConfigExport) | **Get** /config/export | 
[**ConfigImport**](ConfigApi.md#ConfigImport) | **Post** /config/import | 
[**GetConfigExportWorkflow**](ConfigApi.md#GetConfigExportWorkflow) | **Get** /config/export/workflows/{id} | 
[**GetConfigExportWorkflowData**](ConfigApi.md#GetConfigExportWorkflowData) | **Get** /config/export/workflows/{id}/data | 
[**GetConfigExportWorkflows**](ConfigApi.md#GetConfigExportWorkflows) | **Get** /config/export/workflows | 
[**GetConfigImportWorkflow**](ConfigApi.md#GetConfigImportWorkflow) | **Get** /config/import/workflows/{id} | 
[**GetConfigImportWorkflows**](ConfigApi.md#GetConfigImportWorkflows) | **Get** /config/import/workflows | 



## AddConfigExportWorkflow

> ConfigStatus AddConfigExportWorkflow(ctx).Execute()





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
    resp, r, err := apiClient.ConfigApi.AddConfigExportWorkflow(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.AddConfigExportWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddConfigExportWorkflow`: ConfigStatus
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.AddConfigExportWorkflow`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAddConfigExportWorkflowRequest struct via the builder pattern


### Return type

[**ConfigStatus**](ConfigStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddConfigImportWorkflow

> AddConfigImportWorkflow(ctx).FailFast(failFast).BackupFile(backupFile).Execute()





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
    failFast := true // bool | When set to true, stops the import as soon as any validation errors are encountered. When false,  import will continue to validate configuration after the first failure to identify all validation errors. If any validation errors are present PingAccess will revert back to the state prior to import attempt.
    backupFile := *openapiclient.NewExportData() // ExportData | JSON backup file to import

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ConfigApi.AddConfigImportWorkflow(context.Background()).FailFast(failFast).BackupFile(backupFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.AddConfigImportWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddConfigImportWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **failFast** | **bool** | When set to true, stops the import as soon as any validation errors are encountered. When false,  import will continue to validate configuration after the first failure to identify all validation errors. If any validation errors are present PingAccess will revert back to the state prior to import attempt. | 
 **backupFile** | [**ExportData**](ExportData.md) | JSON backup file to import | 

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


## ConfigExport

> ExportData ConfigExport(ctx).Execute()





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
    resp, r, err := apiClient.ConfigApi.ConfigExport(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.ConfigExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfigExport`: ExportData
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.ConfigExport`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiConfigExportRequest struct via the builder pattern


### Return type

[**ExportData**](ExportData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigImport

> ConfigImport(ctx).BackupFile(backupFile).Execute()





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
    backupFile := "backupFile_example" // string | JSON backup file to import

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ConfigApi.ConfigImport(context.Background()).BackupFile(backupFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.ConfigImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfigImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **backupFile** | **string** | JSON backup file to import | 

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


## GetConfigExportWorkflow

> ConfigStatus GetConfigExportWorkflow(ctx, id).Execute()





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
    id := "id_example" // string | Id of Export to poll

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigApi.GetConfigExportWorkflow(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.GetConfigExportWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigExportWorkflow`: ConfigStatus
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.GetConfigExportWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of Export to poll | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigExportWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConfigStatus**](ConfigStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigExportWorkflowData

> ExportData GetConfigExportWorkflowData(ctx, id).Execute()





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
    id := "id_example" // string | ID of Export to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigApi.GetConfigExportWorkflowData(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.GetConfigExportWorkflowData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigExportWorkflowData`: ExportData
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.GetConfigExportWorkflowData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Export to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigExportWorkflowDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExportData**](ExportData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigExportWorkflows

> ConfigStatuses GetConfigExportWorkflows(ctx).Execute()





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
    resp, r, err := apiClient.ConfigApi.GetConfigExportWorkflows(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.GetConfigExportWorkflows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigExportWorkflows`: ConfigStatuses
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.GetConfigExportWorkflows`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigExportWorkflowsRequest struct via the builder pattern


### Return type

[**ConfigStatuses**](ConfigStatuses.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigImportWorkflow

> ConfigStatus GetConfigImportWorkflow(ctx, id).Execute()





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
    id := "id_example" // string | Id of Import to poll

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigApi.GetConfigImportWorkflow(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.GetConfigImportWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigImportWorkflow`: ConfigStatus
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.GetConfigImportWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of Import to poll | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigImportWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConfigStatus**](ConfigStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigImportWorkflows

> ConfigStatuses GetConfigImportWorkflows(ctx).Execute()





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
    resp, r, err := apiClient.ConfigApi.GetConfigImportWorkflows(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.GetConfigImportWorkflows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigImportWorkflows`: ConfigStatuses
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.GetConfigImportWorkflows`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigImportWorkflowsRequest struct via the builder pattern


### Return type

[**ConfigStatuses**](ConfigStatuses.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

