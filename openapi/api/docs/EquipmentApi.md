# \EquipmentApi

All URIs are relative to *https://api.osrsbox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEquipmentItem**](EquipmentApi.md#DeleteEquipmentItem) | **Delete** /equipment/{equipmentId} | Deletes a Equipment document
[**Deleteequipment**](EquipmentApi.md#Deleteequipment) | **Delete** /equipment | Deletes all equipment
[**GetEquipmentItem**](EquipmentApi.md#GetEquipmentItem) | **Get** /equipment/{equipmentId} | Retrieves a Equipment document
[**Getequipment**](EquipmentApi.md#Getequipment) | **Get** /equipment | Retrieves one or more equipment
[**Postequipment**](EquipmentApi.md#Postequipment) | **Post** /equipment | Stores one or more equipment.
[**PutEquipmentItem**](EquipmentApi.md#PutEquipmentItem) | **Put** /equipment/{equipmentId} | Replaces a Equipment document



## DeleteEquipmentItem

> DeleteEquipmentItem(ctx, equipmentId).IfMatch(ifMatch).Execute()

Deletes a Equipment document

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
    equipmentId := "equipmentId_example" // string | Unique OSRS item ID number.
    ifMatch := "ifMatch_example" // string | Current value of the _etag field

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EquipmentApi.DeleteEquipmentItem(context.Background(), equipmentId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EquipmentApi.DeleteEquipmentItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**equipmentId** | **string** | Unique OSRS item ID number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEquipmentItemRequest struct via the builder pattern


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


## Deleteequipment

> Deleteequipment(ctx).Execute()

Deletes all equipment

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
    resp, r, err := api_client.EquipmentApi.Deleteequipment(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EquipmentApi.Deleteequipment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteequipmentRequest struct via the builder pattern


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


## GetEquipmentItem

> Equipment GetEquipmentItem(ctx, equipmentId).Execute()

Retrieves a Equipment document

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
    equipmentId := "equipmentId_example" // string | Unique OSRS item ID number.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EquipmentApi.GetEquipmentItem(context.Background(), equipmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EquipmentApi.GetEquipmentItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEquipmentItem`: Equipment
    fmt.Fprintf(os.Stdout, "Response from `EquipmentApi.GetEquipmentItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**equipmentId** | **string** | Unique OSRS item ID number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEquipmentItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Equipment**](Equipment.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Getequipment

> InlineResponse2002 Getequipment(ctx).Where(where).Projection(projection).Sort(sort).Page(page).MaxResults(maxResults).Execute()

Retrieves one or more equipment

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
    resp, r, err := api_client.EquipmentApi.Getequipment(context.Background()).Where(where).Projection(projection).Sort(sort).Page(page).MaxResults(maxResults).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EquipmentApi.Getequipment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Getequipment`: InlineResponse2002
    fmt.Fprintf(os.Stdout, "Response from `EquipmentApi.Getequipment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetequipmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **where** | **string** | the filters query parameter (ex.: {\&quot;number\&quot;: 10}) | 
 **projection** | **string** | the projections query parameter (ex.: {\&quot;name\&quot;: 1}) | 
 **sort** | **string** | the sort query parameter (ex.: \&quot;city,-lastname\&quot;) | 
 **page** | **int32** | the pages query parameter | 
 **maxResults** | **int32** | the max results query parameter | 

### Return type

[**InlineResponse2002**](InlineResponse2002.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Postequipment

> Postequipment(ctx).Equipment(equipment).Execute()

Stores one or more equipment.

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
    equipment := *openapiclient.NewEquipment("Id_example", "Name_example", "LastUpdated_example", false, false, false, false, false, NullableInt32(123), false, false, NullableInt32(123), NullableInt32(123), NullableInt32(123), false, false, false, false, int32(123), NullableInt32(123), NullableInt32(123), NullableFloat32(123), NullableInt32(123), false, "ReleaseDate_example", false, "Examine_example", "Icon_example", "WikiName_example", "WikiUrl_example", "TODO", "TODO") // Equipment | A Equipment or list of Equipment documents

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EquipmentApi.Postequipment(context.Background()).Equipment(equipment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EquipmentApi.Postequipment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostequipmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **equipment** | [**Equipment**](Equipment.md) | A Equipment or list of Equipment documents | 

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


## PutEquipmentItem

> PutEquipmentItem(ctx, equipmentId).IfMatch(ifMatch).Equipment(equipment).Execute()

Replaces a Equipment document

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
    equipmentId := "equipmentId_example" // string | Unique OSRS item ID number.
    ifMatch := "ifMatch_example" // string | Current value of the _etag field
    equipment := *openapiclient.NewEquipment("Id_example", "Name_example", "LastUpdated_example", false, false, false, false, false, NullableInt32(123), false, false, NullableInt32(123), NullableInt32(123), NullableInt32(123), false, false, false, false, int32(123), NullableInt32(123), NullableInt32(123), NullableFloat32(123), NullableInt32(123), false, "ReleaseDate_example", false, "Examine_example", "Icon_example", "WikiName_example", "WikiUrl_example", "TODO", "TODO") // Equipment | A Equipment or list of Equipment documents

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EquipmentApi.PutEquipmentItem(context.Background(), equipmentId).IfMatch(ifMatch).Equipment(equipment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EquipmentApi.PutEquipmentItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**equipmentId** | **string** | Unique OSRS item ID number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutEquipmentItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Current value of the _etag field | 
 **equipment** | [**Equipment**](Equipment.md) | A Equipment or list of Equipment documents | 

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

