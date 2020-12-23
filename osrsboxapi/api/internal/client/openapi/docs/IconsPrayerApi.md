# \IconsPrayerApi

All URIs are relative to *http://api.osrsbox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIconsPrayerItem**](IconsPrayerApi.md#GetIconsPrayerItem) | **Get** /icons_prayers/{icons_prayerId} | Retrieves a Icons_prayer document
[**GeticonsPrayers**](IconsPrayerApi.md#GeticonsPrayers) | **Get** /icons_prayers | Retrieves one or more icons_prayers



## GetIconsPrayerItem

> IconsPrayer GetIconsPrayerItem(ctx, iconsPrayerId).Execute()

Retrieves a Icons_prayer document

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
    iconsPrayerId := "iconsPrayerId_example" // string | Unique OSRS prayer ID number.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IconsPrayerApi.GetIconsPrayerItem(context.Background(), iconsPrayerId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IconsPrayerApi.GetIconsPrayerItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIconsPrayerItem`: IconsPrayer
    fmt.Fprintf(os.Stdout, "Response from `IconsPrayerApi.GetIconsPrayerItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iconsPrayerId** | **string** | Unique OSRS prayer ID number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIconsPrayerItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IconsPrayer**](Icons_prayer.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GeticonsPrayers

> InlineResponse2006 GeticonsPrayers(ctx).Where(where).Projection(projection).Sort(sort).Page(page).MaxResults(maxResults).Execute()

Retrieves one or more icons_prayers

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
    resp, r, err := api_client.IconsPrayerApi.GeticonsPrayers(context.Background()).Where(where).Projection(projection).Sort(sort).Page(page).MaxResults(maxResults).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IconsPrayerApi.GeticonsPrayers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GeticonsPrayers`: InlineResponse2006
    fmt.Fprintf(os.Stdout, "Response from `IconsPrayerApi.GeticonsPrayers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGeticonsPrayersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **where** | **string** | the filters query parameter (ex.: {\&quot;number\&quot;: 10}) | 
 **projection** | **string** | the projections query parameter (ex.: {\&quot;name\&quot;: 1}) | 
 **sort** | **string** | the sort query parameter (ex.: \&quot;city,-lastname\&quot;) | 
 **page** | **int32** | the pages query parameter | 
 **maxResults** | **int32** | the max results query parameter | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

