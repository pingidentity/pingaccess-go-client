# \HttpConfigApi

All URIs are relative to *http://localhost/pa-admin-api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteHostSource**](HttpConfigApi.md#DeleteHostSource) | **Delete** /httpConfig/request/hostSource | 
[**DeleteHttpMonitoring**](HttpConfigApi.md#DeleteHttpMonitoring) | **Delete** /httpConfig/monitoring | 
[**DeleteIpSource**](HttpConfigApi.md#DeleteIpSource) | **Delete** /httpConfig/request/ipSource | 
[**DeleteProtoSource**](HttpConfigApi.md#DeleteProtoSource) | **Delete** /httpConfig/request/protocolSource | 
[**GetHostSource**](HttpConfigApi.md#GetHostSource) | **Get** /httpConfig/request/hostSource | 
[**GetHttpMonitoring**](HttpConfigApi.md#GetHttpMonitoring) | **Get** /httpConfig/monitoring | 
[**GetIpSource**](HttpConfigApi.md#GetIpSource) | **Get** /httpConfig/request/ipSource | 
[**GetProtoSource**](HttpConfigApi.md#GetProtoSource) | **Get** /httpConfig/request/protocolSource | 
[**UpdateHostSource**](HttpConfigApi.md#UpdateHostSource) | **Put** /httpConfig/request/hostSource | 
[**UpdateHttpMonitoring**](HttpConfigApi.md#UpdateHttpMonitoring) | **Put** /httpConfig/monitoring | 
[**UpdateIpSource**](HttpConfigApi.md#UpdateIpSource) | **Put** /httpConfig/request/ipSource | 
[**UpdateProtocolSource**](HttpConfigApi.md#UpdateProtocolSource) | **Put** /httpConfig/request/protocolSource | 



## DeleteHostSource

> DeleteHostSource(ctx).Execute()





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
    r, err := apiClient.HttpConfigApi.DeleteHostSource(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HttpConfigApi.DeleteHostSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHostSourceRequest struct via the builder pattern


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


## DeleteHttpMonitoring

> DeleteHttpMonitoring(ctx).Execute()





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
    r, err := apiClient.HttpConfigApi.DeleteHttpMonitoring(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HttpConfigApi.DeleteHttpMonitoring``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHttpMonitoringRequest struct via the builder pattern


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


## DeleteIpSource

> DeleteIpSource(ctx).Execute()





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
    r, err := apiClient.HttpConfigApi.DeleteIpSource(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HttpConfigApi.DeleteIpSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIpSourceRequest struct via the builder pattern


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


## DeleteProtoSource

> DeleteProtoSource(ctx).Execute()





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
    r, err := apiClient.HttpConfigApi.DeleteProtoSource(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HttpConfigApi.DeleteProtoSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProtoSourceRequest struct via the builder pattern


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


## GetHostSource

> HostMultiValueSource GetHostSource(ctx).Execute()





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
    resp, r, err := apiClient.HttpConfigApi.GetHostSource(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HttpConfigApi.GetHostSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHostSource`: HostMultiValueSource
    fmt.Fprintf(os.Stdout, "Response from `HttpConfigApi.GetHostSource`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostSourceRequest struct via the builder pattern


### Return type

[**HostMultiValueSource**](HostMultiValueSource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHttpMonitoring

> HttpMonitoring GetHttpMonitoring(ctx).Execute()





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
    resp, r, err := apiClient.HttpConfigApi.GetHttpMonitoring(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HttpConfigApi.GetHttpMonitoring``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHttpMonitoring`: HttpMonitoring
    fmt.Fprintf(os.Stdout, "Response from `HttpConfigApi.GetHttpMonitoring`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHttpMonitoringRequest struct via the builder pattern


### Return type

[**HttpMonitoring**](HttpMonitoring.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIpSource

> IpMultiValueSource GetIpSource(ctx).Execute()





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
    resp, r, err := apiClient.HttpConfigApi.GetIpSource(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HttpConfigApi.GetIpSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIpSource`: IpMultiValueSource
    fmt.Fprintf(os.Stdout, "Response from `HttpConfigApi.GetIpSource`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIpSourceRequest struct via the builder pattern


### Return type

[**IpMultiValueSource**](IpMultiValueSource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProtoSource

> ProtocolSource GetProtoSource(ctx).Execute()





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
    resp, r, err := apiClient.HttpConfigApi.GetProtoSource(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HttpConfigApi.GetProtoSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProtoSource`: ProtocolSource
    fmt.Fprintf(os.Stdout, "Response from `HttpConfigApi.GetProtoSource`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProtoSourceRequest struct via the builder pattern


### Return type

[**ProtocolSource**](ProtocolSource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHostSource

> HostMultiValueSource UpdateHostSource(ctx).HttpConfiguration(httpConfiguration).Execute()





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
    httpConfiguration := *openapiclient.NewHostMultiValueSource([]string{"HeaderNameList_example"}, openapiclient.ListValueLocation("FIRST")) // HostMultiValueSource | Host source type to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HttpConfigApi.UpdateHostSource(context.Background()).HttpConfiguration(httpConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HttpConfigApi.UpdateHostSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateHostSource`: HostMultiValueSource
    fmt.Fprintf(os.Stdout, "Response from `HttpConfigApi.UpdateHostSource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHostSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **httpConfiguration** | [**HostMultiValueSource**](HostMultiValueSource.md) | Host source type to update | 

### Return type

[**HostMultiValueSource**](HostMultiValueSource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHttpMonitoring

> HttpMonitoring UpdateHttpMonitoring(ctx).Monitoring(monitoring).Execute()





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
    monitoring := *openapiclient.NewHttpMonitoring() // HttpMonitoring | Http Monitoring

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HttpConfigApi.UpdateHttpMonitoring(context.Background()).Monitoring(monitoring).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HttpConfigApi.UpdateHttpMonitoring``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateHttpMonitoring`: HttpMonitoring
    fmt.Fprintf(os.Stdout, "Response from `HttpConfigApi.UpdateHttpMonitoring`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHttpMonitoringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **monitoring** | [**HttpMonitoring**](HttpMonitoring.md) | Http Monitoring | 

### Return type

[**HttpMonitoring**](HttpMonitoring.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIpSource

> IpMultiValueSource UpdateIpSource(ctx).HttpConfiguration(httpConfiguration).Execute()





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
    httpConfiguration := *openapiclient.NewIpMultiValueSource([]string{"HeaderNameList_example"}, openapiclient.ListValueLocation("FIRST")) // IpMultiValueSource | IP source type to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HttpConfigApi.UpdateIpSource(context.Background()).HttpConfiguration(httpConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HttpConfigApi.UpdateIpSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIpSource`: IpMultiValueSource
    fmt.Fprintf(os.Stdout, "Response from `HttpConfigApi.UpdateIpSource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIpSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **httpConfiguration** | [**IpMultiValueSource**](IpMultiValueSource.md) | IP source type to update | 

### Return type

[**IpMultiValueSource**](IpMultiValueSource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProtocolSource

> ProtocolSource UpdateProtocolSource(ctx).HttpConfiguration(httpConfiguration).Execute()





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
    httpConfiguration := *openapiclient.NewProtocolSource("HeaderName_example") // ProtocolSource | Protocol source type to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HttpConfigApi.UpdateProtocolSource(context.Background()).HttpConfiguration(httpConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HttpConfigApi.UpdateProtocolSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProtocolSource`: ProtocolSource
    fmt.Fprintf(os.Stdout, "Response from `HttpConfigApi.UpdateProtocolSource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProtocolSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **httpConfiguration** | [**ProtocolSource**](ProtocolSource.md) | Protocol source type to update | 

### Return type

[**ProtocolSource**](ProtocolSource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

