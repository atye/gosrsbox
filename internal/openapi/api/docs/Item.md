# Item

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique OSRS item ID number. | 
**Name** | **string** | The name of the item. | 
**Incomplete** | **bool** | If the item has incomplete wiki data. | 
**Members** | **bool** | If the item is a members-only. | 
**Tradeable** | **bool** | If the item is tradeable (between players and on the GE). | 
**TradeableOnGe** | **bool** | If the item is tradeable (only on GE). | 
**Stackable** | **bool** | If the item is stackable (in inventory). | 
**Stacked** | **NullableInt32** | If the item is stacked, indicated by the stack count. | 
**Noted** | **bool** | If the item is noted. | 
**Noteable** | **bool** | If the item is noteable. | 
**LinkedIdItem** | **NullableInt32** | The linked ID of the actual item (if noted/placeholder). | 
**LinkedIdNoted** | **NullableInt32** | The linked ID of an item in noted form. | 
**LinkedIdPlaceholder** | **NullableInt32** | The linked ID of an item in placeholder form. | 
**Placeholder** | **bool** | If the item is a placeholder. | 
**Equipable** | **bool** | If the item is equipable (based on right-click menu entry). | 
**EquipableByPlayer** | **bool** | If the item is equipable in-game by a player. | 
**EquipableWeapon** | **bool** | If the item is an equipable weapon. | 
**Cost** | **int32** | The store price of an item. | 
**Lowalch** | **NullableInt32** | The low alchemy value of the item (cost * 0.4). | 
**Highalch** | **NullableInt32** | The high alchemy value of the item (cost * 0.6). | 
**Weight** | **NullableFloat32** | The weight (in kilograms) of the item. | 
**BuyLimit** | **NullableInt32** | The Grand Exchange buy limit of the item. | 
**QuestItem** | **bool** | If the item is associated with a quest. | 
**ReleaseDate** | **NullableString** | Date the item was released (in ISO8601 format). | 
**Duplicate** | **bool** | If the item is a duplicate. | 
**Examine** | **NullableString** | The examine text for the item. | 
**Icon** | **string** | The item icon (in base64 encoding). | 
**WikiName** | **NullableString** | The OSRS Wiki name for the item. | 
**WikiUrl** | **NullableString** | The OSRS Wiki URL (possibly including anchor link). | 
**WikiExchange** | **NullableString** | The OSRS Wiki Exchange URL. | 
**Equipment** | [**NullableItemEquipment**](Item_equipment.md) |  | 
**Weapon** | [**NullableItemWeapon**](Item_weapon.md) |  | 

## Methods

### NewItem

`func NewItem(id string, name string, incomplete bool, members bool, tradeable bool, tradeableOnGe bool, stackable bool, stacked NullableInt32, noted bool, noteable bool, linkedIdItem NullableInt32, linkedIdNoted NullableInt32, linkedIdPlaceholder NullableInt32, placeholder bool, equipable bool, equipableByPlayer bool, equipableWeapon bool, cost int32, lowalch NullableInt32, highalch NullableInt32, weight NullableFloat32, buyLimit NullableInt32, questItem bool, releaseDate NullableString, duplicate bool, examine NullableString, icon string, wikiName NullableString, wikiUrl NullableString, wikiExchange NullableString, equipment NullableItemEquipment, weapon NullableItemWeapon, ) *Item`

NewItem instantiates a new Item object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemWithDefaults

`func NewItemWithDefaults() *Item`

NewItemWithDefaults instantiates a new Item object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Item) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Item) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Item) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Item) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Item) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Item) SetName(v string)`

SetName sets Name field to given value.


### GetIncomplete

`func (o *Item) GetIncomplete() bool`

GetIncomplete returns the Incomplete field if non-nil, zero value otherwise.

### GetIncompleteOk

`func (o *Item) GetIncompleteOk() (*bool, bool)`

GetIncompleteOk returns a tuple with the Incomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomplete

`func (o *Item) SetIncomplete(v bool)`

SetIncomplete sets Incomplete field to given value.


### GetMembers

`func (o *Item) GetMembers() bool`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Item) GetMembersOk() (*bool, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *Item) SetMembers(v bool)`

SetMembers sets Members field to given value.


### GetTradeable

`func (o *Item) GetTradeable() bool`

GetTradeable returns the Tradeable field if non-nil, zero value otherwise.

### GetTradeableOk

`func (o *Item) GetTradeableOk() (*bool, bool)`

GetTradeableOk returns a tuple with the Tradeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeable

`func (o *Item) SetTradeable(v bool)`

SetTradeable sets Tradeable field to given value.


### GetTradeableOnGe

`func (o *Item) GetTradeableOnGe() bool`

GetTradeableOnGe returns the TradeableOnGe field if non-nil, zero value otherwise.

### GetTradeableOnGeOk

`func (o *Item) GetTradeableOnGeOk() (*bool, bool)`

GetTradeableOnGeOk returns a tuple with the TradeableOnGe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeableOnGe

`func (o *Item) SetTradeableOnGe(v bool)`

SetTradeableOnGe sets TradeableOnGe field to given value.


### GetStackable

`func (o *Item) GetStackable() bool`

GetStackable returns the Stackable field if non-nil, zero value otherwise.

### GetStackableOk

`func (o *Item) GetStackableOk() (*bool, bool)`

GetStackableOk returns a tuple with the Stackable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackable

`func (o *Item) SetStackable(v bool)`

SetStackable sets Stackable field to given value.


### GetStacked

`func (o *Item) GetStacked() int32`

GetStacked returns the Stacked field if non-nil, zero value otherwise.

### GetStackedOk

`func (o *Item) GetStackedOk() (*int32, bool)`

GetStackedOk returns a tuple with the Stacked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStacked

`func (o *Item) SetStacked(v int32)`

SetStacked sets Stacked field to given value.


### SetStackedNil

`func (o *Item) SetStackedNil(b bool)`

 SetStackedNil sets the value for Stacked to be an explicit nil

### UnsetStacked
`func (o *Item) UnsetStacked()`

UnsetStacked ensures that no value is present for Stacked, not even an explicit nil
### GetNoted

`func (o *Item) GetNoted() bool`

GetNoted returns the Noted field if non-nil, zero value otherwise.

### GetNotedOk

`func (o *Item) GetNotedOk() (*bool, bool)`

GetNotedOk returns a tuple with the Noted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoted

`func (o *Item) SetNoted(v bool)`

SetNoted sets Noted field to given value.


### GetNoteable

`func (o *Item) GetNoteable() bool`

GetNoteable returns the Noteable field if non-nil, zero value otherwise.

### GetNoteableOk

`func (o *Item) GetNoteableOk() (*bool, bool)`

GetNoteableOk returns a tuple with the Noteable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteable

`func (o *Item) SetNoteable(v bool)`

SetNoteable sets Noteable field to given value.


### GetLinkedIdItem

`func (o *Item) GetLinkedIdItem() int32`

GetLinkedIdItem returns the LinkedIdItem field if non-nil, zero value otherwise.

### GetLinkedIdItemOk

`func (o *Item) GetLinkedIdItemOk() (*int32, bool)`

GetLinkedIdItemOk returns a tuple with the LinkedIdItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedIdItem

`func (o *Item) SetLinkedIdItem(v int32)`

SetLinkedIdItem sets LinkedIdItem field to given value.


### SetLinkedIdItemNil

`func (o *Item) SetLinkedIdItemNil(b bool)`

 SetLinkedIdItemNil sets the value for LinkedIdItem to be an explicit nil

### UnsetLinkedIdItem
`func (o *Item) UnsetLinkedIdItem()`

UnsetLinkedIdItem ensures that no value is present for LinkedIdItem, not even an explicit nil
### GetLinkedIdNoted

`func (o *Item) GetLinkedIdNoted() int32`

GetLinkedIdNoted returns the LinkedIdNoted field if non-nil, zero value otherwise.

### GetLinkedIdNotedOk

`func (o *Item) GetLinkedIdNotedOk() (*int32, bool)`

GetLinkedIdNotedOk returns a tuple with the LinkedIdNoted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedIdNoted

`func (o *Item) SetLinkedIdNoted(v int32)`

SetLinkedIdNoted sets LinkedIdNoted field to given value.


### SetLinkedIdNotedNil

`func (o *Item) SetLinkedIdNotedNil(b bool)`

 SetLinkedIdNotedNil sets the value for LinkedIdNoted to be an explicit nil

### UnsetLinkedIdNoted
`func (o *Item) UnsetLinkedIdNoted()`

UnsetLinkedIdNoted ensures that no value is present for LinkedIdNoted, not even an explicit nil
### GetLinkedIdPlaceholder

`func (o *Item) GetLinkedIdPlaceholder() int32`

GetLinkedIdPlaceholder returns the LinkedIdPlaceholder field if non-nil, zero value otherwise.

### GetLinkedIdPlaceholderOk

`func (o *Item) GetLinkedIdPlaceholderOk() (*int32, bool)`

GetLinkedIdPlaceholderOk returns a tuple with the LinkedIdPlaceholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedIdPlaceholder

`func (o *Item) SetLinkedIdPlaceholder(v int32)`

SetLinkedIdPlaceholder sets LinkedIdPlaceholder field to given value.


### SetLinkedIdPlaceholderNil

`func (o *Item) SetLinkedIdPlaceholderNil(b bool)`

 SetLinkedIdPlaceholderNil sets the value for LinkedIdPlaceholder to be an explicit nil

### UnsetLinkedIdPlaceholder
`func (o *Item) UnsetLinkedIdPlaceholder()`

UnsetLinkedIdPlaceholder ensures that no value is present for LinkedIdPlaceholder, not even an explicit nil
### GetPlaceholder

`func (o *Item) GetPlaceholder() bool`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *Item) GetPlaceholderOk() (*bool, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *Item) SetPlaceholder(v bool)`

SetPlaceholder sets Placeholder field to given value.


### GetEquipable

`func (o *Item) GetEquipable() bool`

GetEquipable returns the Equipable field if non-nil, zero value otherwise.

### GetEquipableOk

`func (o *Item) GetEquipableOk() (*bool, bool)`

GetEquipableOk returns a tuple with the Equipable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipable

`func (o *Item) SetEquipable(v bool)`

SetEquipable sets Equipable field to given value.


### GetEquipableByPlayer

`func (o *Item) GetEquipableByPlayer() bool`

GetEquipableByPlayer returns the EquipableByPlayer field if non-nil, zero value otherwise.

### GetEquipableByPlayerOk

`func (o *Item) GetEquipableByPlayerOk() (*bool, bool)`

GetEquipableByPlayerOk returns a tuple with the EquipableByPlayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipableByPlayer

`func (o *Item) SetEquipableByPlayer(v bool)`

SetEquipableByPlayer sets EquipableByPlayer field to given value.


### GetEquipableWeapon

`func (o *Item) GetEquipableWeapon() bool`

GetEquipableWeapon returns the EquipableWeapon field if non-nil, zero value otherwise.

### GetEquipableWeaponOk

`func (o *Item) GetEquipableWeaponOk() (*bool, bool)`

GetEquipableWeaponOk returns a tuple with the EquipableWeapon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipableWeapon

`func (o *Item) SetEquipableWeapon(v bool)`

SetEquipableWeapon sets EquipableWeapon field to given value.


### GetCost

`func (o *Item) GetCost() int32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *Item) GetCostOk() (*int32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *Item) SetCost(v int32)`

SetCost sets Cost field to given value.


### GetLowalch

`func (o *Item) GetLowalch() int32`

GetLowalch returns the Lowalch field if non-nil, zero value otherwise.

### GetLowalchOk

`func (o *Item) GetLowalchOk() (*int32, bool)`

GetLowalchOk returns a tuple with the Lowalch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowalch

`func (o *Item) SetLowalch(v int32)`

SetLowalch sets Lowalch field to given value.


### SetLowalchNil

`func (o *Item) SetLowalchNil(b bool)`

 SetLowalchNil sets the value for Lowalch to be an explicit nil

### UnsetLowalch
`func (o *Item) UnsetLowalch()`

UnsetLowalch ensures that no value is present for Lowalch, not even an explicit nil
### GetHighalch

`func (o *Item) GetHighalch() int32`

GetHighalch returns the Highalch field if non-nil, zero value otherwise.

### GetHighalchOk

`func (o *Item) GetHighalchOk() (*int32, bool)`

GetHighalchOk returns a tuple with the Highalch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighalch

`func (o *Item) SetHighalch(v int32)`

SetHighalch sets Highalch field to given value.


### SetHighalchNil

`func (o *Item) SetHighalchNil(b bool)`

 SetHighalchNil sets the value for Highalch to be an explicit nil

### UnsetHighalch
`func (o *Item) UnsetHighalch()`

UnsetHighalch ensures that no value is present for Highalch, not even an explicit nil
### GetWeight

`func (o *Item) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *Item) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *Item) SetWeight(v float32)`

SetWeight sets Weight field to given value.


### SetWeightNil

`func (o *Item) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *Item) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetBuyLimit

`func (o *Item) GetBuyLimit() int32`

GetBuyLimit returns the BuyLimit field if non-nil, zero value otherwise.

### GetBuyLimitOk

`func (o *Item) GetBuyLimitOk() (*int32, bool)`

GetBuyLimitOk returns a tuple with the BuyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyLimit

`func (o *Item) SetBuyLimit(v int32)`

SetBuyLimit sets BuyLimit field to given value.


### SetBuyLimitNil

`func (o *Item) SetBuyLimitNil(b bool)`

 SetBuyLimitNil sets the value for BuyLimit to be an explicit nil

### UnsetBuyLimit
`func (o *Item) UnsetBuyLimit()`

UnsetBuyLimit ensures that no value is present for BuyLimit, not even an explicit nil
### GetQuestItem

`func (o *Item) GetQuestItem() bool`

GetQuestItem returns the QuestItem field if non-nil, zero value otherwise.

### GetQuestItemOk

`func (o *Item) GetQuestItemOk() (*bool, bool)`

GetQuestItemOk returns a tuple with the QuestItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestItem

`func (o *Item) SetQuestItem(v bool)`

SetQuestItem sets QuestItem field to given value.


### GetReleaseDate

`func (o *Item) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *Item) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *Item) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.


### SetReleaseDateNil

`func (o *Item) SetReleaseDateNil(b bool)`

 SetReleaseDateNil sets the value for ReleaseDate to be an explicit nil

### UnsetReleaseDate
`func (o *Item) UnsetReleaseDate()`

UnsetReleaseDate ensures that no value is present for ReleaseDate, not even an explicit nil
### GetDuplicate

`func (o *Item) GetDuplicate() bool`

GetDuplicate returns the Duplicate field if non-nil, zero value otherwise.

### GetDuplicateOk

`func (o *Item) GetDuplicateOk() (*bool, bool)`

GetDuplicateOk returns a tuple with the Duplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicate

`func (o *Item) SetDuplicate(v bool)`

SetDuplicate sets Duplicate field to given value.


### GetExamine

`func (o *Item) GetExamine() string`

GetExamine returns the Examine field if non-nil, zero value otherwise.

### GetExamineOk

`func (o *Item) GetExamineOk() (*string, bool)`

GetExamineOk returns a tuple with the Examine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExamine

`func (o *Item) SetExamine(v string)`

SetExamine sets Examine field to given value.


### SetExamineNil

`func (o *Item) SetExamineNil(b bool)`

 SetExamineNil sets the value for Examine to be an explicit nil

### UnsetExamine
`func (o *Item) UnsetExamine()`

UnsetExamine ensures that no value is present for Examine, not even an explicit nil
### GetIcon

`func (o *Item) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *Item) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *Item) SetIcon(v string)`

SetIcon sets Icon field to given value.


### GetWikiName

`func (o *Item) GetWikiName() string`

GetWikiName returns the WikiName field if non-nil, zero value otherwise.

### GetWikiNameOk

`func (o *Item) GetWikiNameOk() (*string, bool)`

GetWikiNameOk returns a tuple with the WikiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWikiName

`func (o *Item) SetWikiName(v string)`

SetWikiName sets WikiName field to given value.


### SetWikiNameNil

`func (o *Item) SetWikiNameNil(b bool)`

 SetWikiNameNil sets the value for WikiName to be an explicit nil

### UnsetWikiName
`func (o *Item) UnsetWikiName()`

UnsetWikiName ensures that no value is present for WikiName, not even an explicit nil
### GetWikiUrl

`func (o *Item) GetWikiUrl() string`

GetWikiUrl returns the WikiUrl field if non-nil, zero value otherwise.

### GetWikiUrlOk

`func (o *Item) GetWikiUrlOk() (*string, bool)`

GetWikiUrlOk returns a tuple with the WikiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWikiUrl

`func (o *Item) SetWikiUrl(v string)`

SetWikiUrl sets WikiUrl field to given value.


### SetWikiUrlNil

`func (o *Item) SetWikiUrlNil(b bool)`

 SetWikiUrlNil sets the value for WikiUrl to be an explicit nil

### UnsetWikiUrl
`func (o *Item) UnsetWikiUrl()`

UnsetWikiUrl ensures that no value is present for WikiUrl, not even an explicit nil
### GetWikiExchange

`func (o *Item) GetWikiExchange() string`

GetWikiExchange returns the WikiExchange field if non-nil, zero value otherwise.

### GetWikiExchangeOk

`func (o *Item) GetWikiExchangeOk() (*string, bool)`

GetWikiExchangeOk returns a tuple with the WikiExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWikiExchange

`func (o *Item) SetWikiExchange(v string)`

SetWikiExchange sets WikiExchange field to given value.


### SetWikiExchangeNil

`func (o *Item) SetWikiExchangeNil(b bool)`

 SetWikiExchangeNil sets the value for WikiExchange to be an explicit nil

### UnsetWikiExchange
`func (o *Item) UnsetWikiExchange()`

UnsetWikiExchange ensures that no value is present for WikiExchange, not even an explicit nil
### GetEquipment

`func (o *Item) GetEquipment() ItemEquipment`

GetEquipment returns the Equipment field if non-nil, zero value otherwise.

### GetEquipmentOk

`func (o *Item) GetEquipmentOk() (*ItemEquipment, bool)`

GetEquipmentOk returns a tuple with the Equipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipment

`func (o *Item) SetEquipment(v ItemEquipment)`

SetEquipment sets Equipment field to given value.


### SetEquipmentNil

`func (o *Item) SetEquipmentNil(b bool)`

 SetEquipmentNil sets the value for Equipment to be an explicit nil

### UnsetEquipment
`func (o *Item) UnsetEquipment()`

UnsetEquipment ensures that no value is present for Equipment, not even an explicit nil
### GetWeapon

`func (o *Item) GetWeapon() ItemWeapon`

GetWeapon returns the Weapon field if non-nil, zero value otherwise.

### GetWeaponOk

`func (o *Item) GetWeaponOk() (*ItemWeapon, bool)`

GetWeaponOk returns a tuple with the Weapon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeapon

`func (o *Item) SetWeapon(v ItemWeapon)`

SetWeapon sets Weapon field to given value.


### SetWeaponNil

`func (o *Item) SetWeaponNil(b bool)`

 SetWeaponNil sets the value for Weapon to be an explicit nil

### UnsetWeapon
`func (o *Item) UnsetWeapon()`

UnsetWeapon ensures that no value is present for Weapon, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


