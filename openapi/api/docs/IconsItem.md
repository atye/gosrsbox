# IconsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique OSRS item ID number. | 
**Icon** | **string** | The icon image (in base64 encoding) of the item. | 

## Methods

### NewIconsItem

`func NewIconsItem(id string, icon string, ) *IconsItem`

NewIconsItem instantiates a new IconsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIconsItemWithDefaults

`func NewIconsItemWithDefaults() *IconsItem`

NewIconsItemWithDefaults instantiates a new IconsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IconsItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IconsItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IconsItem) SetId(v string)`

SetId sets Id field to given value.


### GetIcon

`func (o *IconsItem) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *IconsItem) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *IconsItem) SetIcon(v string)`

SetIcon sets Icon field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


