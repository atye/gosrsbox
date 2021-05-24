# InlineResponse200

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]Item**](Item.md) |  | [optional] 
**Meta** | Pointer to [**ResponeMetadata**](ResponeMetadata.md) |  | [optional] 
**Links** | Pointer to [**ResponeLinks**](ResponeLinks.md) |  | [optional] 

## Methods

### NewInlineResponse200

`func NewInlineResponse200() *InlineResponse200`

NewInlineResponse200 instantiates a new InlineResponse200 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200WithDefaults

`func NewInlineResponse200WithDefaults() *InlineResponse200`

NewInlineResponse200WithDefaults instantiates a new InlineResponse200 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200) GetItems() []Item`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200) GetItemsOk() (*[]Item, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200) SetItems(v []Item)`

SetItems sets Items field to given value.

### HasItems

`func (o *InlineResponse200) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMeta

`func (o *InlineResponse200) GetMeta() ResponeMetadata`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200) GetMetaOk() (*ResponeMetadata, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200) SetMeta(v ResponeMetadata)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse200) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetLinks

`func (o *InlineResponse200) GetLinks() ResponeLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *InlineResponse200) GetLinksOk() (*ResponeLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *InlineResponse200) SetLinks(v ResponeLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *InlineResponse200) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


