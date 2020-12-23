# InlineResponse2002

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]Equipment**](Equipment.md) |  | [optional] 
**Meta** | Pointer to [**ResponeMetadata**](respone_metadata.md) |  | [optional] 
**Links** | Pointer to [**ResponeLinks**](respone_links.md) |  | [optional] 

## Methods

### NewInlineResponse2002

`func NewInlineResponse2002() *InlineResponse2002`

NewInlineResponse2002 instantiates a new InlineResponse2002 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2002WithDefaults

`func NewInlineResponse2002WithDefaults() *InlineResponse2002`

NewInlineResponse2002WithDefaults instantiates a new InlineResponse2002 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse2002) GetItems() []Equipment`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse2002) GetItemsOk() (*[]Equipment, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse2002) SetItems(v []Equipment)`

SetItems sets Items field to given value.

### HasItems

`func (o *InlineResponse2002) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMeta

`func (o *InlineResponse2002) GetMeta() ResponeMetadata`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse2002) GetMetaOk() (*ResponeMetadata, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse2002) SetMeta(v ResponeMetadata)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse2002) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetLinks

`func (o *InlineResponse2002) GetLinks() ResponeLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *InlineResponse2002) GetLinksOk() (*ResponeLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *InlineResponse2002) SetLinks(v ResponeLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *InlineResponse2002) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


