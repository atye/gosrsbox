# \MonsterApi

All URIs are relative to *https://api.osrsbox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMonsterItem**](MonsterApi.md#DeleteMonsterItem) | **Delete** /monsters/{monsterId} | Deletes a Monster document
[**Deletemonsters**](MonsterApi.md#Deletemonsters) | **Delete** /monsters | Deletes all monsters
[**GetMonsterItem**](MonsterApi.md#GetMonsterItem) | **Get** /monsters/{monsterId} | Retrieves a Monster document
[**Getmonsters**](MonsterApi.md#Getmonsters) | **Get** /monsters | Retrieves one or more monsters
[**Postmonsters**](MonsterApi.md#Postmonsters) | **Post** /monsters | Stores one or more monsters.
[**PutMonsterItem**](MonsterApi.md#PutMonsterItem) | **Put** /monsters/{monsterId} | Replaces a Monster document



## DeleteMonsterItem

> DeleteMonsterItem(ctx, monsterId).IfMatch(ifMatch).Execute()

Deletes a Monster document

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
    monsterId := "monsterId_example" // string | Unique OSRS monster ID number.
    ifMatch := "ifMatch_example" // string | Current value of the _etag field

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonsterApi.DeleteMonsterItem(context.Background(), monsterId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonsterApi.DeleteMonsterItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monsterId** | **string** | Unique OSRS monster ID number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMonsterItemRequest struct via the builder pattern


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


## Deletemonsters

> Deletemonsters(ctx).Execute()

Deletes all monsters

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
    resp, r, err := api_client.MonsterApi.Deletemonsters(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonsterApi.Deletemonsters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeletemonstersRequest struct via the builder pattern


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


## GetMonsterItem

> Monster GetMonsterItem(ctx, monsterId).Execute()

Retrieves a Monster document

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
    monsterId := "monsterId_example" // string | Unique OSRS monster ID number.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonsterApi.GetMonsterItem(context.Background(), monsterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonsterApi.GetMonsterItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonsterItem`: Monster
    fmt.Fprintf(os.Stdout, "Response from `MonsterApi.GetMonsterItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monsterId** | **string** | Unique OSRS monster ID number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonsterItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Monster**](Monster.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Getmonsters

> InlineResponse2003 Getmonsters(ctx).Where(where).Projection(projection).Sort(sort).Page(page).MaxResults(maxResults).Execute()

Retrieves one or more monsters

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
    resp, r, err := api_client.MonsterApi.Getmonsters(context.Background()).Where(where).Projection(projection).Sort(sort).Page(page).MaxResults(maxResults).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonsterApi.Getmonsters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Getmonsters`: InlineResponse2003
    fmt.Fprintf(os.Stdout, "Response from `MonsterApi.Getmonsters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetmonstersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **where** | **string** | the filters query parameter (ex.: {\&quot;number\&quot;: 10}) | 
 **projection** | **string** | the projections query parameter (ex.: {\&quot;name\&quot;: 1}) | 
 **sort** | **string** | the sort query parameter (ex.: \&quot;city,-lastname\&quot;) | 
 **page** | **int32** | the pages query parameter | 
 **maxResults** | **int32** | the max results query parameter | 

### Return type

[**InlineResponse2003**](InlineResponse2003.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Postmonsters

> Postmonsters(ctx).Monster(monster).Execute()

Stores one or more monsters.

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
    monster := *openapiclient.NewMonster("Id_example", "Name_example", "LastUpdated_example", false, false, "ReleaseDate_example", int32(123), int32(123), NullableInt32(123), NullableInt32(123), []string{"AttackType_example"}, NullableInt32(123), false, false, false, false, false, []string{"Attributes_example"}, []string{"Category_example"}, false, NullableInt32(123), NullableFloat32(123), []string{"SlayerMasters_example"}, false, "Examine_example", "WikiName_example", "WikiUrl_example", int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), []openapiclient.MonsterDrops{*openapiclient.NewMonsterDrops(int32(123), "Name_example", false, "Quantity_example", false, float32(123), int32(123))}) // Monster | A Monster or list of Monster documents

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonsterApi.Postmonsters(context.Background()).Monster(monster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonsterApi.Postmonsters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostmonstersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **monster** | [**Monster**](Monster.md) | A Monster or list of Monster documents | 

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


## PutMonsterItem

> PutMonsterItem(ctx, monsterId).IfMatch(ifMatch).Monster(monster).Execute()

Replaces a Monster document

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
    monsterId := "monsterId_example" // string | Unique OSRS monster ID number.
    ifMatch := "ifMatch_example" // string | Current value of the _etag field
    monster := *openapiclient.NewMonster("Id_example", "Name_example", "LastUpdated_example", false, false, "ReleaseDate_example", int32(123), int32(123), NullableInt32(123), NullableInt32(123), []string{"AttackType_example"}, NullableInt32(123), false, false, false, false, false, []string{"Attributes_example"}, []string{"Category_example"}, false, NullableInt32(123), NullableFloat32(123), []string{"SlayerMasters_example"}, false, "Examine_example", "WikiName_example", "WikiUrl_example", int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), []openapiclient.MonsterDrops{*openapiclient.NewMonsterDrops(int32(123), "Name_example", false, "Quantity_example", false, float32(123), int32(123))}) // Monster | A Monster or list of Monster documents

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonsterApi.PutMonsterItem(context.Background(), monsterId).IfMatch(ifMatch).Monster(monster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonsterApi.PutMonsterItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monsterId** | **string** | Unique OSRS monster ID number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutMonsterItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Current value of the _etag field | 
 **monster** | [**Monster**](Monster.md) | A Monster or list of Monster documents | 

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

