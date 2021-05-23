# \IconsItemApi

All URIs are relative to *https://api.osrsbox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIconsItemItem**](IconsItemApi.md#DeleteIconsItemItem) | **Delete** /icons_items/{icons_itemId} | Deletes a Icons_item document
[**DeleteiconsItems**](IconsItemApi.md#DeleteiconsItems) | **Delete** /icons_items | Deletes all icons_items
[**GetIconsItemItem**](IconsItemApi.md#GetIconsItemItem) | **Get** /icons_items/{icons_itemId} | Retrieves a Icons_item document
[**GeticonsItems**](IconsItemApi.md#GeticonsItems) | **Get** /icons_items | Retrieves one or more icons_items
[**PosticonsItems**](IconsItemApi.md#PosticonsItems) | **Post** /icons_items | Stores one or more icons_items.
[**PutIconsItemItem**](IconsItemApi.md#PutIconsItemItem) | **Put** /icons_items/{icons_itemId} | Replaces a Icons_item document



## DeleteIconsItemItem

> DeleteIconsItemItem(ctx, iconsItemId).IfMatch(ifMatch).Execute()

Deletes a Icons_item document

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
    iconsItemId := "iconsItemId_example" // string | Unique OSRS item ID number.
    ifMatch := "ifMatch_example" // string | Current value of the _etag field

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IconsItemApi.DeleteIconsItemItem(context.Background(), iconsItemId).IfMatch(ifMatch).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IconsItemApi.DeleteIconsItemItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iconsItemId** | **string** | Unique OSRS item ID number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIconsItemItemRequest struct via the builder pattern


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


## DeleteiconsItems

> DeleteiconsItems(ctx).Execute()

Deletes all icons_items

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
    resp, r, err := api_client.IconsItemApi.DeleteiconsItems(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IconsItemApi.DeleteiconsItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteiconsItemsRequest struct via the builder pattern


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


## GetIconsItemItem

> IconsItem GetIconsItemItem(ctx, iconsItemId).Execute()

Retrieves a Icons_item document

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
    iconsItemId := "iconsItemId_example" // string | Unique OSRS item ID number.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IconsItemApi.GetIconsItemItem(context.Background(), iconsItemId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IconsItemApi.GetIconsItemItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIconsItemItem`: IconsItem
    fmt.Fprintf(os.Stdout, "Response from `IconsItemApi.GetIconsItemItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iconsItemId** | **string** | Unique OSRS item ID number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIconsItemItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IconsItem**](Icons_item.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GeticonsItems

> InlineResponse2005 GeticonsItems(ctx).Where(where).Projection(projection).Sort(sort).Page(page).MaxResults(maxResults).Execute()

Retrieves one or more icons_items

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
    resp, r, err := api_client.IconsItemApi.GeticonsItems(context.Background()).Where(where).Projection(projection).Sort(sort).Page(page).MaxResults(maxResults).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IconsItemApi.GeticonsItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GeticonsItems`: InlineResponse2005
    fmt.Fprintf(os.Stdout, "Response from `IconsItemApi.GeticonsItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGeticonsItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **where** | **string** | the filters query parameter (ex.: {\&quot;number\&quot;: 10}) | 
 **projection** | **string** | the projections query parameter (ex.: {\&quot;name\&quot;: 1}) | 
 **sort** | **string** | the sort query parameter (ex.: \&quot;city,-lastname\&quot;) | 
 **page** | **int32** | the pages query parameter | 
 **maxResults** | **int32** | the max results query parameter | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PosticonsItems

> PosticonsItems(ctx).IconsItem(iconsItem).Execute()

Stores one or more icons_items.

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
    iconsItem := *openapiclient.NewIconsItem("Id_example", "Icon_example") // IconsItem | A Icons_item or list of Icons_item documents

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IconsItemApi.PosticonsItems(context.Background()).IconsItem(iconsItem).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IconsItemApi.PosticonsItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPosticonsItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iconsItem** | [**IconsItem**](IconsItem.md) | A Icons_item or list of Icons_item documents | 

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


## PutIconsItemItem

> PutIconsItemItem(ctx, iconsItemId).IfMatch(ifMatch).IconsItem(iconsItem).Execute()

Replaces a Icons_item document

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
    iconsItemId := "iconsItemId_example" // string | Unique OSRS item ID number.
    ifMatch := "ifMatch_example" // string | Current value of the _etag field
    iconsItem := *openapiclient.NewIconsItem("Id_example", "Icon_example") // IconsItem | A Icons_item or list of Icons_item documents

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IconsItemApi.PutIconsItemItem(context.Background(), iconsItemId).IfMatch(ifMatch).IconsItem(iconsItem).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IconsItemApi.PutIconsItemItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iconsItemId** | **string** | Unique OSRS item ID number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutIconsItemItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Current value of the _etag field | 
 **iconsItem** | [**IconsItem**](IconsItem.md) | A Icons_item or list of Icons_item documents | 

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

