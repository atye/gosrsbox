# \PrayerApi

All URIs are relative to *https://api.osrsbox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePrayerItem**](PrayerApi.md#DeletePrayerItem) | **Delete** /prayers/{prayerId} | Deletes a Prayer document
[**Deleteprayers**](PrayerApi.md#Deleteprayers) | **Delete** /prayers | Deletes all prayers
[**GetPrayerItem**](PrayerApi.md#GetPrayerItem) | **Get** /prayers/{prayerId} | Retrieves a Prayer document
[**Getprayers**](PrayerApi.md#Getprayers) | **Get** /prayers | Retrieves one or more prayers
[**Postprayers**](PrayerApi.md#Postprayers) | **Post** /prayers | Stores one or more prayers.
[**PutPrayerItem**](PrayerApi.md#PutPrayerItem) | **Put** /prayers/{prayerId} | Replaces a Prayer document



## DeletePrayerItem

> DeletePrayerItem(ctx, prayerId).IfMatch(ifMatch).Execute()

Deletes a Prayer document

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    prayerId := "prayerId_example" // string | Unique prayer ID number.
    ifMatch := "ifMatch_example" // string | Current value of the _etag field

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrayerApi.DeletePrayerItem(context.Background(), prayerId).IfMatch(ifMatch).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PrayerApi.DeletePrayerItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prayerId** | **string** | Unique prayer ID number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrayerItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Current value of the _etag field | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Deleteprayers

> Deleteprayers(ctx).Execute()

Deletes all prayers

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrayerApi.Deleteprayers(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PrayerApi.Deleteprayers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteprayersRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrayerItem

> Prayer GetPrayerItem(ctx, prayerId).Execute()

Retrieves a Prayer document

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    prayerId := "prayerId_example" // string | Unique prayer ID number.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrayerApi.GetPrayerItem(context.Background(), prayerId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PrayerApi.GetPrayerItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPrayerItem`: Prayer
    fmt.Fprintf(os.Stdout, "Response from `PrayerApi.GetPrayerItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prayerId** | **string** | Unique prayer ID number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrayerItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Prayer**](Prayer.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Getprayers

> InlineResponse2004 Getprayers(ctx).Where(where).Projection(projection).Sort(sort).Page(page).MaxResults(maxResults).Execute()

Retrieves one or more prayers

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    where := "where_example" // string | the filters query parameter (ex.: {\"number\": 10}) (optional)
    projection := "projection_example" // string | the projections query parameter (ex.: {\"name\": 1}) (optional)
    sort := "sort_example" // string | the sort query parameter (ex.: \"city,-lastname\") (optional)
    page := int32(1) // int32 | the pages query parameter (optional)
    maxResults := int32(25) // int32 | the max results query parameter (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrayerApi.Getprayers(context.Background()).Where(where).Projection(projection).Sort(sort).Page(page).MaxResults(maxResults).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PrayerApi.Getprayers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Getprayers`: InlineResponse2004
    fmt.Fprintf(os.Stdout, "Response from `PrayerApi.Getprayers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetprayersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **where** | **string** | the filters query parameter (ex.: {\&quot;number\&quot;: 10}) | 
 **projection** | **string** | the projections query parameter (ex.: {\&quot;name\&quot;: 1}) | 
 **sort** | **string** | the sort query parameter (ex.: \&quot;city,-lastname\&quot;) | 
 **page** | **int32** | the pages query parameter | 
 **maxResults** | **int32** | the max results query parameter | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Postprayers

> Postprayers(ctx).Prayer(prayer).Execute()

Stores one or more prayers.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    prayer := *openapiclient.NewPrayer("Id_example", "Name_example", false, "Description_example", float32(123), "WikiUrl_example", map[string]interface{}(123), map[string]interface{}(123), "Icon_example") // Prayer | A Prayer or list of Prayer documents

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrayerApi.Postprayers(context.Background()).Prayer(prayer).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PrayerApi.Postprayers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostprayersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **prayer** | [**Prayer**](Prayer.md) | A Prayer or list of Prayer documents | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPrayerItem

> PutPrayerItem(ctx, prayerId).IfMatch(ifMatch).Prayer(prayer).Execute()

Replaces a Prayer document

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    prayerId := "prayerId_example" // string | Unique prayer ID number.
    ifMatch := "ifMatch_example" // string | Current value of the _etag field
    prayer := *openapiclient.NewPrayer("Id_example", "Name_example", false, "Description_example", float32(123), "WikiUrl_example", map[string]interface{}(123), map[string]interface{}(123), "Icon_example") // Prayer | A Prayer or list of Prayer documents

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrayerApi.PutPrayerItem(context.Background(), prayerId).IfMatch(ifMatch).Prayer(prayer).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PrayerApi.PutPrayerItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prayerId** | **string** | Unique prayer ID number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPrayerItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Current value of the _etag field | 
 **prayer** | [**Prayer**](Prayer.md) | A Prayer or list of Prayer documents | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

