# MonsterDrops

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID number of the item drop. | 
**Name** | **string** | The name of the item drop. | 
**Members** | **bool** | If the drop is a members-only item. | 
**Quantity** | **NullableString** | The quantity of the item drop (integer, comma-separated or range). | 
**Noted** | **bool** | If the item drop is noted, or not. | 
**Rarity** | **NullableFloat32** | The rarity of the item drop (as a float out of 1.0). | 
**DropRequirements** | **NullableString** | If there are any requirements to getting the specific drop. | 

## Methods

### NewMonsterDrops

`func NewMonsterDrops(id int32, name string, members bool, quantity NullableString, noted bool, rarity NullableFloat32, dropRequirements NullableString, ) *MonsterDrops`

NewMonsterDrops instantiates a new MonsterDrops object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonsterDropsWithDefaults

`func NewMonsterDropsWithDefaults() *MonsterDrops`

NewMonsterDropsWithDefaults instantiates a new MonsterDrops object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MonsterDrops) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MonsterDrops) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MonsterDrops) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *MonsterDrops) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MonsterDrops) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MonsterDrops) SetName(v string)`

SetName sets Name field to given value.


### GetMembers

`func (o *MonsterDrops) GetMembers() bool`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *MonsterDrops) GetMembersOk() (*bool, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *MonsterDrops) SetMembers(v bool)`

SetMembers sets Members field to given value.


### GetQuantity

`func (o *MonsterDrops) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *MonsterDrops) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *MonsterDrops) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.


### SetQuantityNil

`func (o *MonsterDrops) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *MonsterDrops) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetNoted

`func (o *MonsterDrops) GetNoted() bool`

GetNoted returns the Noted field if non-nil, zero value otherwise.

### GetNotedOk

`func (o *MonsterDrops) GetNotedOk() (*bool, bool)`

GetNotedOk returns a tuple with the Noted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoted

`func (o *MonsterDrops) SetNoted(v bool)`

SetNoted sets Noted field to given value.


### GetRarity

`func (o *MonsterDrops) GetRarity() float32`

GetRarity returns the Rarity field if non-nil, zero value otherwise.

### GetRarityOk

`func (o *MonsterDrops) GetRarityOk() (*float32, bool)`

GetRarityOk returns a tuple with the Rarity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRarity

`func (o *MonsterDrops) SetRarity(v float32)`

SetRarity sets Rarity field to given value.


### SetRarityNil

`func (o *MonsterDrops) SetRarityNil(b bool)`

 SetRarityNil sets the value for Rarity to be an explicit nil

### UnsetRarity
`func (o *MonsterDrops) UnsetRarity()`

UnsetRarity ensures that no value is present for Rarity, not even an explicit nil
### GetDropRequirements

`func (o *MonsterDrops) GetDropRequirements() string`

GetDropRequirements returns the DropRequirements field if non-nil, zero value otherwise.

### GetDropRequirementsOk

`func (o *MonsterDrops) GetDropRequirementsOk() (*string, bool)`

GetDropRequirementsOk returns a tuple with the DropRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropRequirements

`func (o *MonsterDrops) SetDropRequirements(v string)`

SetDropRequirements sets DropRequirements field to given value.


### SetDropRequirementsNil

`func (o *MonsterDrops) SetDropRequirementsNil(b bool)`

 SetDropRequirementsNil sets the value for DropRequirements to be an explicit nil

### UnsetDropRequirements
`func (o *MonsterDrops) UnsetDropRequirements()`

UnsetDropRequirements ensures that no value is present for DropRequirements, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


