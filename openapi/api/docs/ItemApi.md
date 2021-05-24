# \ItemApi

All URIs are relative to *https://api.osrsbox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteItemItem**](ItemApi.md#DeleteItemItem) | **Delete** /items/{itemId} | Deletes a Item document
[**Deleteitems**](ItemApi.md#Deleteitems) | **Delete** /items | Deletes all items
[**GetItemItem**](ItemApi.md#GetItemItem) | **Get** /items/{itemId} | Retrieves a Item document
[**Getitems**](ItemApi.md#Getitems) | **Get** /items | Retrieves one or more items
[**Postitems**](ItemApi.md#Postitems) | **Post** /items | Stores one or more items.
[**PutItemItem**](ItemApi.md#PutItemItem) | **Put** /items/{itemId} | Replaces a Item document



## DeleteItemItem

> DeleteItemItem(ctx, itemId).IfMatch(ifMatch).Execute()

Deletes a Item document

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
    itemId := "itemId_example" // string | Unique OSRS item ID number.
    ifMatch := "ifMatch_example" // string | Current value of the _etag field

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ItemApi.DeleteItemItem(context.Background(), itemId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemApi.DeleteItemItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | Unique OSRS item ID number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteItemItemRequest struct via the builder pattern


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


## Deleteitems

> Deleteitems(ctx).Execute()

Deletes all items

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
    resp, r, err := api_client.ItemApi.Deleteitems(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemApi.Deleteitems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteitemsRequest struct via the builder pattern


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


## GetItemItem

> Item GetItemItem(ctx, itemId).Execute()

Retrieves a Item document

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
    itemId := "itemId_example" // string | Unique OSRS item ID number.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ItemApi.GetItemItem(context.Background(), itemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemApi.GetItemItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetItemItem`: Item
    fmt.Fprintf(os.Stdout, "Response from `ItemApi.GetItemItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | Unique OSRS item ID number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Item**](Item.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Getitems

> InlineResponse200 Getitems(ctx).Where(where).Projection(projection).Sort(sort).Page(page).MaxResults(maxResults).Execute()

Retrieves one or more items

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
    resp, r, err := api_client.ItemApi.Getitems(context.Background()).Where(where).Projection(projection).Sort(sort).Page(page).MaxResults(maxResults).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemApi.Getitems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Getitems`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `ItemApi.Getitems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetitemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **where** | **string** | the filters query parameter (ex.: {\&quot;number\&quot;: 10}) | 
 **projection** | **string** | the projections query parameter (ex.: {\&quot;name\&quot;: 1}) | 
 **sort** | **string** | the sort query parameter (ex.: \&quot;city,-lastname\&quot;) | 
 **page** | **int32** | the pages query parameter | 
 **maxResults** | **int32** | the max results query parameter | 

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Postitems

> Postitems(ctx).Item(item).Execute()

Stores one or more items.

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
    item := *openapiclient.NewItem("Id_example", "Name_example", "LastUpdated_example", false, false, false, false, false, NullableInt32(123), false, false, NullableInt32(123), NullableInt32(123), NullableInt32(123), false, false, false, false, int32(123), NullableInt32(123), NullableInt32(123), NullableFloat32(123), NullableInt32(123), false, "ReleaseDate_example", false, "Examine_example", "Icon_example", "WikiName_example", "WikiUrl_example", "TODO", "TODO") // Item | A Item or list of Item documents

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ItemApi.Postitems(context.Background()).Item(item).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemApi.Postitems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostitemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **item** | [**Item**](Item.md) | A Item or list of Item documents | 

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


## PutItemItem

> PutItemItem(ctx, itemId).IfMatch(ifMatch).Item(item).Execute()

Replaces a Item document

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
    itemId := "itemId_example" // string | Unique OSRS item ID number.
    ifMatch := "ifMatch_example" // string | Current value of the _etag field
    item := *openapiclient.NewItem("Id_example", "Name_example", "LastUpdated_example", false, false, false, false, false, NullableInt32(123), false, false, NullableInt32(123), NullableInt32(123), NullableInt32(123), false, false, false, false, int32(123), NullableInt32(123), NullableInt32(123), NullableFloat32(123), NullableInt32(123), false, "ReleaseDate_example", false, "Examine_example", "Icon_example", "WikiName_example", "WikiUrl_example", "TODO", "TODO") // Item | A Item or list of Item documents

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ItemApi.PutItemItem(context.Background(), itemId).IfMatch(ifMatch).Item(item).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemApi.PutItemItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | Unique OSRS item ID number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutItemItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Current value of the _etag field | 
 **item** | [**Item**](Item.md) | A Item or list of Item documents | 

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

