# Equipment

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

### NewEquipment

`func NewEquipment(id string, name string, incomplete bool, members bool, tradeable bool, tradeableOnGe bool, stackable bool, stacked NullableInt32, noted bool, noteable bool, linkedIdItem NullableInt32, linkedIdNoted NullableInt32, linkedIdPlaceholder NullableInt32, placeholder bool, equipable bool, equipableByPlayer bool, equipableWeapon bool, cost int32, lowalch NullableInt32, highalch NullableInt32, weight NullableFloat32, buyLimit NullableInt32, questItem bool, releaseDate NullableString, duplicate bool, examine NullableString, icon string, wikiName NullableString, wikiUrl NullableString, wikiExchange NullableString, equipment NullableItemEquipment, weapon NullableItemWeapon, ) *Equipment`

NewEquipment instantiates a new Equipment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentWithDefaults

`func NewEquipmentWithDefaults() *Equipment`

NewEquipmentWithDefaults instantiates a new Equipment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Equipment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Equipment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Equipment) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Equipment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Equipment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Equipment) SetName(v string)`

SetName sets Name field to given value.


### GetIncomplete

`func (o *Equipment) GetIncomplete() bool`

GetIncomplete returns the Incomplete field if non-nil, zero value otherwise.

### GetIncompleteOk

`func (o *Equipment) GetIncompleteOk() (*bool, bool)`

GetIncompleteOk returns a tuple with the Incomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomplete

`func (o *Equipment) SetIncomplete(v bool)`

SetIncomplete sets Incomplete field to given value.


### GetMembers

`func (o *Equipment) GetMembers() bool`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Equipment) GetMembersOk() (*bool, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *Equipment) SetMembers(v bool)`

SetMembers sets Members field to given value.


### GetTradeable

`func (o *Equipment) GetTradeable() bool`

GetTradeable returns the Tradeable field if non-nil, zero value otherwise.

### GetTradeableOk

`func (o *Equipment) GetTradeableOk() (*bool, bool)`

GetTradeableOk returns a tuple with the Tradeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeable

`func (o *Equipment) SetTradeable(v bool)`

SetTradeable sets Tradeable field to given value.


### GetTradeableOnGe

`func (o *Equipment) GetTradeableOnGe() bool`

GetTradeableOnGe returns the TradeableOnGe field if non-nil, zero value otherwise.

### GetTradeableOnGeOk

`func (o *Equipment) GetTradeableOnGeOk() (*bool, bool)`

GetTradeableOnGeOk returns a tuple with the TradeableOnGe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeableOnGe

`func (o *Equipment) SetTradeableOnGe(v bool)`

SetTradeableOnGe sets TradeableOnGe field to given value.


### GetStackable

`func (o *Equipment) GetStackable() bool`

GetStackable returns the Stackable field if non-nil, zero value otherwise.

### GetStackableOk

`func (o *Equipment) GetStackableOk() (*bool, bool)`

GetStackableOk returns a tuple with the Stackable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackable

`func (o *Equipment) SetStackable(v bool)`

SetStackable sets Stackable field to given value.


### GetStacked

`func (o *Equipment) GetStacked() int32`

GetStacked returns the Stacked field if non-nil, zero value otherwise.

### GetStackedOk

`func (o *Equipment) GetStackedOk() (*int32, bool)`

GetStackedOk returns a tuple with the Stacked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStacked

`func (o *Equipment) SetStacked(v int32)`

SetStacked sets Stacked field to given value.


### SetStackedNil

`func (o *Equipment) SetStackedNil(b bool)`

 SetStackedNil sets the value for Stacked to be an explicit nil

### UnsetStacked
`func (o *Equipment) UnsetStacked()`

UnsetStacked ensures that no value is present for Stacked, not even an explicit nil
### GetNoted

`func (o *Equipment) GetNoted() bool`

GetNoted returns the Noted field if non-nil, zero value otherwise.

### GetNotedOk

`func (o *Equipment) GetNotedOk() (*bool, bool)`

GetNotedOk returns a tuple with the Noted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoted

`func (o *Equipment) SetNoted(v bool)`

SetNoted sets Noted field to given value.


### GetNoteable

`func (o *Equipment) GetNoteable() bool`

GetNoteable returns the Noteable field if non-nil, zero value otherwise.

### GetNoteableOk

`func (o *Equipment) GetNoteableOk() (*bool, bool)`

GetNoteableOk returns a tuple with the Noteable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteable

`func (o *Equipment) SetNoteable(v bool)`

SetNoteable sets Noteable field to given value.


### GetLinkedIdItem

`func (o *Equipment) GetLinkedIdItem() int32`

GetLinkedIdItem returns the LinkedIdItem field if non-nil, zero value otherwise.

### GetLinkedIdItemOk

`func (o *Equipment) GetLinkedIdItemOk() (*int32, bool)`

GetLinkedIdItemOk returns a tuple with the LinkedIdItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedIdItem

`func (o *Equipment) SetLinkedIdItem(v int32)`

SetLinkedIdItem sets LinkedIdItem field to given value.


### SetLinkedIdItemNil

`func (o *Equipment) SetLinkedIdItemNil(b bool)`

 SetLinkedIdItemNil sets the value for LinkedIdItem to be an explicit nil

### UnsetLinkedIdItem
`func (o *Equipment) UnsetLinkedIdItem()`

UnsetLinkedIdItem ensures that no value is present for LinkedIdItem, not even an explicit nil
### GetLinkedIdNoted

`func (o *Equipment) GetLinkedIdNoted() int32`

GetLinkedIdNoted returns the LinkedIdNoted field if non-nil, zero value otherwise.

### GetLinkedIdNotedOk

`func (o *Equipment) GetLinkedIdNotedOk() (*int32, bool)`

GetLinkedIdNotedOk returns a tuple with the LinkedIdNoted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedIdNoted

`func (o *Equipment) SetLinkedIdNoted(v int32)`

SetLinkedIdNoted sets LinkedIdNoted field to given value.


### SetLinkedIdNotedNil

`func (o *Equipment) SetLinkedIdNotedNil(b bool)`

 SetLinkedIdNotedNil sets the value for LinkedIdNoted to be an explicit nil

### UnsetLinkedIdNoted
`func (o *Equipment) UnsetLinkedIdNoted()`

UnsetLinkedIdNoted ensures that no value is present for LinkedIdNoted, not even an explicit nil
### GetLinkedIdPlaceholder

`func (o *Equipment) GetLinkedIdPlaceholder() int32`

GetLinkedIdPlaceholder returns the LinkedIdPlaceholder field if non-nil, zero value otherwise.

### GetLinkedIdPlaceholderOk

`func (o *Equipment) GetLinkedIdPlaceholderOk() (*int32, bool)`

GetLinkedIdPlaceholderOk returns a tuple with the LinkedIdPlaceholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedIdPlaceholder

`func (o *Equipment) SetLinkedIdPlaceholder(v int32)`

SetLinkedIdPlaceholder sets LinkedIdPlaceholder field to given value.


### SetLinkedIdPlaceholderNil

`func (o *Equipment) SetLinkedIdPlaceholderNil(b bool)`

 SetLinkedIdPlaceholderNil sets the value for LinkedIdPlaceholder to be an explicit nil

### UnsetLinkedIdPlaceholder
`func (o *Equipment) UnsetLinkedIdPlaceholder()`

UnsetLinkedIdPlaceholder ensures that no value is present for LinkedIdPlaceholder, not even an explicit nil
### GetPlaceholder

`func (o *Equipment) GetPlaceholder() bool`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *Equipment) GetPlaceholderOk() (*bool, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *Equipment) SetPlaceholder(v bool)`

SetPlaceholder sets Placeholder field to given value.


### GetEquipable

`func (o *Equipment) GetEquipable() bool`

GetEquipable returns the Equipable field if non-nil, zero value otherwise.

### GetEquipableOk

`func (o *Equipment) GetEquipableOk() (*bool, bool)`

GetEquipableOk returns a tuple with the Equipable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipable

`func (o *Equipment) SetEquipable(v bool)`

SetEquipable sets Equipable field to given value.


### GetEquipableByPlayer

`func (o *Equipment) GetEquipableByPlayer() bool`

GetEquipableByPlayer returns the EquipableByPlayer field if non-nil, zero value otherwise.

### GetEquipableByPlayerOk

`func (o *Equipment) GetEquipableByPlayerOk() (*bool, bool)`

GetEquipableByPlayerOk returns a tuple with the EquipableByPlayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipableByPlayer

`func (o *Equipment) SetEquipableByPlayer(v bool)`

SetEquipableByPlayer sets EquipableByPlayer field to given value.


### GetEquipableWeapon

`func (o *Equipment) GetEquipableWeapon() bool`

GetEquipableWeapon returns the EquipableWeapon field if non-nil, zero value otherwise.

### GetEquipableWeaponOk

`func (o *Equipment) GetEquipableWeaponOk() (*bool, bool)`

GetEquipableWeaponOk returns a tuple with the EquipableWeapon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipableWeapon

`func (o *Equipment) SetEquipableWeapon(v bool)`

SetEquipableWeapon sets EquipableWeapon field to given value.


### GetCost

`func (o *Equipment) GetCost() int32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *Equipment) GetCostOk() (*int32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *Equipment) SetCost(v int32)`

SetCost sets Cost field to given value.


### GetLowalch

`func (o *Equipment) GetLowalch() int32`

GetLowalch returns the Lowalch field if non-nil, zero value otherwise.

### GetLowalchOk

`func (o *Equipment) GetLowalchOk() (*int32, bool)`

GetLowalchOk returns a tuple with the Lowalch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowalch

`func (o *Equipment) SetLowalch(v int32)`

SetLowalch sets Lowalch field to given value.


### SetLowalchNil

`func (o *Equipment) SetLowalchNil(b bool)`

 SetLowalchNil sets the value for Lowalch to be an explicit nil

### UnsetLowalch
`func (o *Equipment) UnsetLowalch()`

UnsetLowalch ensures that no value is present for Lowalch, not even an explicit nil
### GetHighalch

`func (o *Equipment) GetHighalch() int32`

GetHighalch returns the Highalch field if non-nil, zero value otherwise.

### GetHighalchOk

`func (o *Equipment) GetHighalchOk() (*int32, bool)`

GetHighalchOk returns a tuple with the Highalch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighalch

`func (o *Equipment) SetHighalch(v int32)`

SetHighalch sets Highalch field to given value.


### SetHighalchNil

`func (o *Equipment) SetHighalchNil(b bool)`

 SetHighalchNil sets the value for Highalch to be an explicit nil

### UnsetHighalch
`func (o *Equipment) UnsetHighalch()`

UnsetHighalch ensures that no value is present for Highalch, not even an explicit nil
### GetWeight

`func (o *Equipment) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *Equipment) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *Equipment) SetWeight(v float32)`

SetWeight sets Weight field to given value.


### SetWeightNil

`func (o *Equipment) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *Equipment) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetBuyLimit

`func (o *Equipment) GetBuyLimit() int32`

GetBuyLimit returns the BuyLimit field if non-nil, zero value otherwise.

### GetBuyLimitOk

`func (o *Equipment) GetBuyLimitOk() (*int32, bool)`

GetBuyLimitOk returns a tuple with the BuyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyLimit

`func (o *Equipment) SetBuyLimit(v int32)`

SetBuyLimit sets BuyLimit field to given value.


### SetBuyLimitNil

`func (o *Equipment) SetBuyLimitNil(b bool)`

 SetBuyLimitNil sets the value for BuyLimit to be an explicit nil

### UnsetBuyLimit
`func (o *Equipment) UnsetBuyLimit()`

UnsetBuyLimit ensures that no value is present for BuyLimit, not even an explicit nil
### GetQuestItem

`func (o *Equipment) GetQuestItem() bool`

GetQuestItem returns the QuestItem field if non-nil, zero value otherwise.

### GetQuestItemOk

`func (o *Equipment) GetQuestItemOk() (*bool, bool)`

GetQuestItemOk returns a tuple with the QuestItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestItem

`func (o *Equipment) SetQuestItem(v bool)`

SetQuestItem sets QuestItem field to given value.


### GetReleaseDate

`func (o *Equipment) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *Equipment) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *Equipment) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.


### SetReleaseDateNil

`func (o *Equipment) SetReleaseDateNil(b bool)`

 SetReleaseDateNil sets the value for ReleaseDate to be an explicit nil

### UnsetReleaseDate
`func (o *Equipment) UnsetReleaseDate()`

UnsetReleaseDate ensures that no value is present for ReleaseDate, not even an explicit nil
### GetDuplicate

`func (o *Equipment) GetDuplicate() bool`

GetDuplicate returns the Duplicate field if non-nil, zero value otherwise.

### GetDuplicateOk

`func (o *Equipment) GetDuplicateOk() (*bool, bool)`

GetDuplicateOk returns a tuple with the Duplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicate

`func (o *Equipment) SetDuplicate(v bool)`

SetDuplicate sets Duplicate field to given value.


### GetExamine

`func (o *Equipment) GetExamine() string`

GetExamine returns the Examine field if non-nil, zero value otherwise.

### GetExamineOk

`func (o *Equipment) GetExamineOk() (*string, bool)`

GetExamineOk returns a tuple with the Examine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExamine

`func (o *Equipment) SetExamine(v string)`

SetExamine sets Examine field to given value.


### SetExamineNil

`func (o *Equipment) SetExamineNil(b bool)`

 SetExamineNil sets the value for Examine to be an explicit nil

### UnsetExamine
`func (o *Equipment) UnsetExamine()`

UnsetExamine ensures that no value is present for Examine, not even an explicit nil
### GetIcon

`func (o *Equipment) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *Equipment) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *Equipment) SetIcon(v string)`

SetIcon sets Icon field to given value.


### GetWikiName

`func (o *Equipment) GetWikiName() string`

GetWikiName returns the WikiName field if non-nil, zero value otherwise.

### GetWikiNameOk

`func (o *Equipment) GetWikiNameOk() (*string, bool)`

GetWikiNameOk returns a tuple with the WikiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWikiName

`func (o *Equipment) SetWikiName(v string)`

SetWikiName sets WikiName field to given value.


### SetWikiNameNil

`func (o *Equipment) SetWikiNameNil(b bool)`

 SetWikiNameNil sets the value for WikiName to be an explicit nil

### UnsetWikiName
`func (o *Equipment) UnsetWikiName()`

UnsetWikiName ensures that no value is present for WikiName, not even an explicit nil
### GetWikiUrl

`func (o *Equipment) GetWikiUrl() string`

GetWikiUrl returns the WikiUrl field if non-nil, zero value otherwise.

### GetWikiUrlOk

`func (o *Equipment) GetWikiUrlOk() (*string, bool)`

GetWikiUrlOk returns a tuple with the WikiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWikiUrl

`func (o *Equipment) SetWikiUrl(v string)`

SetWikiUrl sets WikiUrl field to given value.


### SetWikiUrlNil

`func (o *Equipment) SetWikiUrlNil(b bool)`

 SetWikiUrlNil sets the value for WikiUrl to be an explicit nil

### UnsetWikiUrl
`func (o *Equipment) UnsetWikiUrl()`

UnsetWikiUrl ensures that no value is present for WikiUrl, not even an explicit nil
### GetWikiExchange

`func (o *Equipment) GetWikiExchange() string`

GetWikiExchange returns the WikiExchange field if non-nil, zero value otherwise.

### GetWikiExchangeOk

`func (o *Equipment) GetWikiExchangeOk() (*string, bool)`

GetWikiExchangeOk returns a tuple with the WikiExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWikiExchange

`func (o *Equipment) SetWikiExchange(v string)`

SetWikiExchange sets WikiExchange field to given value.


### SetWikiExchangeNil

`func (o *Equipment) SetWikiExchangeNil(b bool)`

 SetWikiExchangeNil sets the value for WikiExchange to be an explicit nil

### UnsetWikiExchange
`func (o *Equipment) UnsetWikiExchange()`

UnsetWikiExchange ensures that no value is present for WikiExchange, not even an explicit nil
### GetEquipment

`func (o *Equipment) GetEquipment() ItemEquipment`

GetEquipment returns the Equipment field if non-nil, zero value otherwise.

### GetEquipmentOk

`func (o *Equipment) GetEquipmentOk() (*ItemEquipment, bool)`

GetEquipmentOk returns a tuple with the Equipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipment

`func (o *Equipment) SetEquipment(v ItemEquipment)`

SetEquipment sets Equipment field to given value.


### SetEquipmentNil

`func (o *Equipment) SetEquipmentNil(b bool)`

 SetEquipmentNil sets the value for Equipment to be an explicit nil

### UnsetEquipment
`func (o *Equipment) UnsetEquipment()`

UnsetEquipment ensures that no value is present for Equipment, not even an explicit nil
### GetWeapon

`func (o *Equipment) GetWeapon() ItemWeapon`

GetWeapon returns the Weapon field if non-nil, zero value otherwise.

### GetWeaponOk

`func (o *Equipment) GetWeaponOk() (*ItemWeapon, bool)`

GetWeaponOk returns a tuple with the Weapon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeapon

`func (o *Equipment) SetWeapon(v ItemWeapon)`

SetWeapon sets Weapon field to given value.


### SetWeaponNil

`func (o *Equipment) SetWeaponNil(b bool)`

 SetWeaponNil sets the value for Weapon to be an explicit nil

### UnsetWeapon
`func (o *Equipment) UnsetWeapon()`

UnsetWeapon ensures that no value is present for Weapon, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


