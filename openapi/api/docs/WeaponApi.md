# \WeaponApi

All URIs are relative to *https://api.osrsbox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWeaponItem**](WeaponApi.md#DeleteWeaponItem) | **Delete** /weapons/{weaponId} | Deletes a Weapon document
[**Deleteweapons**](WeaponApi.md#Deleteweapons) | **Delete** /weapons | Deletes all weapons
[**GetWeaponItem**](WeaponApi.md#GetWeaponItem) | **Get** /weapons/{weaponId} | Retrieves a Weapon document
[**Getweapons**](WeaponApi.md#Getweapons) | **Get** /weapons | Retrieves one or more weapons
[**Postweapons**](WeaponApi.md#Postweapons) | **Post** /weapons | Stores one or more weapons.
[**PutWeaponItem**](WeaponApi.md#PutWeaponItem) | **Put** /weapons/{weaponId} | Replaces a Weapon document



## DeleteWeaponItem

> DeleteWeaponItem(ctx, weaponId).IfMatch(ifMatch).Execute()

Deletes a Weapon document

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
    weaponId := "weaponId_example" // string | Unique OSRS item ID number.
    ifMatch := "ifMatch_example" // string | Current value of the _etag field

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WeaponApi.DeleteWeaponItem(context.Background(), weaponId).IfMatch(ifMatch).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `WeaponApi.DeleteWeaponItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**weaponId** | **string** | Unique OSRS item ID number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWeaponItemRequest struct via the builder pattern


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


## Deleteweapons

> Deleteweapons(ctx).Execute()

Deletes all weapons

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
    resp, r, err := api_client.WeaponApi.Deleteweapons(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `WeaponApi.Deleteweapons``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteweaponsRequest struct via the builder pattern


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


## GetWeaponItem

> Weapon GetWeaponItem(ctx, weaponId).Execute()

Retrieves a Weapon document

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
    weaponId := "weaponId_example" // string | Unique OSRS item ID number.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WeaponApi.GetWeaponItem(context.Background(), weaponId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `WeaponApi.GetWeaponItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWeaponItem`: Weapon
    fmt.Fprintf(os.Stdout, "Response from `WeaponApi.GetWeaponItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**weaponId** | **string** | Unique OSRS item ID number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWeaponItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Weapon**](Weapon.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Getweapons

> InlineResponse2001 Getweapons(ctx).Where(where).Projection(projection).Sort(sort).Page(page).MaxResults(maxResults).Execute()

Retrieves one or more weapons

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
    resp, r, err := api_client.WeaponApi.Getweapons(context.Background()).Where(where).Projection(projection).Sort(sort).Page(page).MaxResults(maxResults).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `WeaponApi.Getweapons``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Getweapons`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `WeaponApi.Getweapons`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetweaponsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **where** | **string** | the filters query parameter (ex.: {\&quot;number\&quot;: 10}) | 
 **projection** | **string** | the projections query parameter (ex.: {\&quot;name\&quot;: 1}) | 
 **sort** | **string** | the sort query parameter (ex.: \&quot;city,-lastname\&quot;) | 
 **page** | **int32** | the pages query parameter | 
 **maxResults** | **int32** | the max results query parameter | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Postweapons

> Postweapons(ctx).Weapon(weapon).Execute()

Stores one or more weapons.

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
    weapon := *openapiclient.NewWeapon("Id_example", "Name_example", "LastUpdated_example", false, false, false, false, false, int32(123), false, false, int32(123), int32(123), int32(123), false, false, false, false, int32(123), int32(123), int32(123), float32(123), int32(123), false, "ReleaseDate_example", false, "Examine_example", "Icon_example", "WikiName_example", "WikiUrl_example", *openapiclient.NewItemEquipment(int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), "Slot_example", map[string]interface{}(123)), *openapiclient.NewItemWeapon(int32(123), "WeaponType_example", []openapiclient.ItemWeaponStances{*openapiclient.NewItemWeaponStances("CombatStyle_example", "AttackType_example", "AttackStyle_example", "Experience_example", "Boosts_example")})) // Weapon | A Weapon or list of Weapon documents

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WeaponApi.Postweapons(context.Background()).Weapon(weapon).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `WeaponApi.Postweapons``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostweaponsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **weapon** | [**Weapon**](Weapon.md) | A Weapon or list of Weapon documents | 

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


## PutWeaponItem

> PutWeaponItem(ctx, weaponId).IfMatch(ifMatch).Weapon(weapon).Execute()

Replaces a Weapon document

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
    weaponId := "weaponId_example" // string | Unique OSRS item ID number.
    ifMatch := "ifMatch_example" // string | Current value of the _etag field
    weapon := *openapiclient.NewWeapon("Id_example", "Name_example", "LastUpdated_example", false, false, false, false, false, int32(123), false, false, int32(123), int32(123), int32(123), false, false, false, false, int32(123), int32(123), int32(123), float32(123), int32(123), false, "ReleaseDate_example", false, "Examine_example", "Icon_example", "WikiName_example", "WikiUrl_example", *openapiclient.NewItemEquipment(int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), "Slot_example", map[string]interface{}(123)), *openapiclient.NewItemWeapon(int32(123), "WeaponType_example", []openapiclient.ItemWeaponStances{*openapiclient.NewItemWeaponStances("CombatStyle_example", "AttackType_example", "AttackStyle_example", "Experience_example", "Boosts_example")})) // Weapon | A Weapon or list of Weapon documents

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WeaponApi.PutWeaponItem(context.Background(), weaponId).IfMatch(ifMatch).Weapon(weapon).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `WeaponApi.PutWeaponItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**weaponId** | **string** | Unique OSRS item ID number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutWeaponItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Current value of the _etag field | 
 **weapon** | [**Weapon**](Weapon.md) | A Weapon or list of Weapon documents | 

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

