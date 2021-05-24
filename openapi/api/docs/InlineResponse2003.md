# InlineResponse2003

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]Monster**](Monster.md) |  | [optional] 
**Meta** | Pointer to [**ResponeMetadata**](ResponeMetadata.md) |  | [optional] 
**Links** | Pointer to [**ResponeLinks**](ResponeLinks.md) |  | [optional] 

## Methods

### NewInlineResponse2003

`func NewInlineResponse2003() *InlineResponse2003`

NewInlineResponse2003 instantiates a new InlineResponse2003 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2003WithDefaults

`func NewInlineResponse2003WithDefaults() *InlineResponse2003`

NewInlineResponse2003WithDefaults instantiates a new InlineResponse2003 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse2003) GetItems() []Monster`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse2003) GetItemsOk() (*[]Monster, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse2003) SetItems(v []Monster)`

SetItems sets Items field to given value.

### HasItems

`func (o *InlineResponse2003) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMeta

`func (o *InlineResponse2003) GetMeta() ResponeMetadata`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse2003) GetMetaOk() (*ResponeMetadata, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse2003) SetMeta(v ResponeMetadata)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse2003) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetLinks

`func (o *InlineResponse2003) GetLinks() ResponeLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *InlineResponse2003) GetLinksOk() (*ResponeLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *InlineResponse2003) SetLinks(v ResponeLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *InlineResponse2003) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


