/*
 * osrsbox-api
 *
 * An open, free, complete and up-to-date RESTful API for Old School RuneScape (OSRS) items, monsters and prayers.
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Equipment struct for Equipment
type Equipment struct {
	// Unique OSRS item ID number.
	Id string `json:"id"`
	// The name of the item.
	Name string `json:"name"`
	// If the item has incomplete wiki data.
	Incomplete bool `json:"incomplete"`
	// If the item is a members-only.
	Members bool `json:"members"`
	// If the item is tradeable (between players and on the GE).
	Tradeable bool `json:"tradeable"`
	// If the item is tradeable (only on GE).
	TradeableOnGe bool `json:"tradeable_on_ge"`
	// If the item is stackable (in inventory).
	Stackable bool `json:"stackable"`
	// If the item is stacked, indicated by the stack count.
	Stacked NullableInt32 `json:"stacked"`
	// If the item is noted.
	Noted bool `json:"noted"`
	// If the item is noteable.
	Noteable bool `json:"noteable"`
	// The linked ID of the actual item (if noted/placeholder).
	LinkedIdItem NullableInt32 `json:"linked_id_item"`
	// The linked ID of an item in noted form.
	LinkedIdNoted NullableInt32 `json:"linked_id_noted"`
	// The linked ID of an item in placeholder form.
	LinkedIdPlaceholder NullableInt32 `json:"linked_id_placeholder"`
	// If the item is a placeholder.
	Placeholder bool `json:"placeholder"`
	// If the item is equipable (based on right-click menu entry).
	Equipable bool `json:"equipable"`
	// If the item is equipable in-game by a player.
	EquipableByPlayer bool `json:"equipable_by_player"`
	// If the item is an equipable weapon.
	EquipableWeapon bool `json:"equipable_weapon"`
	// The store price of an item.
	Cost int32 `json:"cost"`
	// The low alchemy value of the item (cost * 0.4).
	Lowalch NullableInt32 `json:"lowalch"`
	// The high alchemy value of the item (cost * 0.6).
	Highalch NullableInt32 `json:"highalch"`
	// The weight (in kilograms) of the item.
	Weight NullableFloat32 `json:"weight"`
	// The Grand Exchange buy limit of the item.
	BuyLimit NullableInt32 `json:"buy_limit"`
	// If the item is associated with a quest.
	QuestItem bool `json:"quest_item"`
	// Date the item was released (in ISO8601 format).
	ReleaseDate NullableString `json:"release_date"`
	// If the item is a duplicate.
	Duplicate bool `json:"duplicate"`
	// The examine text for the item.
	Examine NullableString `json:"examine"`
	// The item icon (in base64 encoding).
	Icon string `json:"icon"`
	// The OSRS Wiki name for the item.
	WikiName NullableString `json:"wiki_name"`
	// The OSRS Wiki URL (possibly including anchor link).
	WikiUrl NullableString `json:"wiki_url"`
	// The OSRS Wiki Exchange URL.
	WikiExchange NullableString `json:"wiki_exchange"`
	Equipment NullableItemEquipment `json:"equipment"`
	Weapon NullableItemWeapon `json:"weapon"`
}

// NewEquipment instantiates a new Equipment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipment(id string, name string, incomplete bool, members bool, tradeable bool, tradeableOnGe bool, stackable bool, stacked NullableInt32, noted bool, noteable bool, linkedIdItem NullableInt32, linkedIdNoted NullableInt32, linkedIdPlaceholder NullableInt32, placeholder bool, equipable bool, equipableByPlayer bool, equipableWeapon bool, cost int32, lowalch NullableInt32, highalch NullableInt32, weight NullableFloat32, buyLimit NullableInt32, questItem bool, releaseDate NullableString, duplicate bool, examine NullableString, icon string, wikiName NullableString, wikiUrl NullableString, wikiExchange NullableString, equipment NullableItemEquipment, weapon NullableItemWeapon, ) *Equipment {
	this := Equipment{}
	this.Id = id
	this.Name = name
	this.Incomplete = incomplete
	this.Members = members
	this.Tradeable = tradeable
	this.TradeableOnGe = tradeableOnGe
	this.Stackable = stackable
	this.Stacked = stacked
	this.Noted = noted
	this.Noteable = noteable
	this.LinkedIdItem = linkedIdItem
	this.LinkedIdNoted = linkedIdNoted
	this.LinkedIdPlaceholder = linkedIdPlaceholder
	this.Placeholder = placeholder
	this.Equipable = equipable
	this.EquipableByPlayer = equipableByPlayer
	this.EquipableWeapon = equipableWeapon
	this.Cost = cost
	this.Lowalch = lowalch
	this.Highalch = highalch
	this.Weight = weight
	this.BuyLimit = buyLimit
	this.QuestItem = questItem
	this.ReleaseDate = releaseDate
	this.Duplicate = duplicate
	this.Examine = examine
	this.Icon = icon
	this.WikiName = wikiName
	this.WikiUrl = wikiUrl
	this.WikiExchange = wikiExchange
	this.Equipment = equipment
	this.Weapon = weapon
	return &this
}

// NewEquipmentWithDefaults instantiates a new Equipment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentWithDefaults() *Equipment {
	this := Equipment{}
	return &this
}

// GetId returns the Id field value
func (o *Equipment) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Equipment) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Equipment) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Equipment) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Equipment) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Equipment) SetName(v string) {
	o.Name = v
}

// GetIncomplete returns the Incomplete field value
func (o *Equipment) GetIncomplete() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.Incomplete
}

// GetIncompleteOk returns a tuple with the Incomplete field value
// and a boolean to check if the value has been set.
func (o *Equipment) GetIncompleteOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Incomplete, true
}

// SetIncomplete sets field value
func (o *Equipment) SetIncomplete(v bool) {
	o.Incomplete = v
}

// GetMembers returns the Members field value
func (o *Equipment) GetMembers() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.Members
}

// GetMembersOk returns a tuple with the Members field value
// and a boolean to check if the value has been set.
func (o *Equipment) GetMembersOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Members, true
}

// SetMembers sets field value
func (o *Equipment) SetMembers(v bool) {
	o.Members = v
}

// GetTradeable returns the Tradeable field value
func (o *Equipment) GetTradeable() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.Tradeable
}

// GetTradeableOk returns a tuple with the Tradeable field value
// and a boolean to check if the value has been set.
func (o *Equipment) GetTradeableOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Tradeable, true
}

// SetTradeable sets field value
func (o *Equipment) SetTradeable(v bool) {
	o.Tradeable = v
}

// GetTradeableOnGe returns the TradeableOnGe field value
func (o *Equipment) GetTradeableOnGe() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.TradeableOnGe
}

// GetTradeableOnGeOk returns a tuple with the TradeableOnGe field value
// and a boolean to check if the value has been set.
func (o *Equipment) GetTradeableOnGeOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TradeableOnGe, true
}

// SetTradeableOnGe sets field value
func (o *Equipment) SetTradeableOnGe(v bool) {
	o.TradeableOnGe = v
}

// GetStackable returns the Stackable field value
func (o *Equipment) GetStackable() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.Stackable
}

// GetStackableOk returns a tuple with the Stackable field value
// and a boolean to check if the value has been set.
func (o *Equipment) GetStackableOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Stackable, true
}

// SetStackable sets field value
func (o *Equipment) SetStackable(v bool) {
	o.Stackable = v
}

// GetStacked returns the Stacked field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *Equipment) GetStacked() int32 {
	if o == nil || o.Stacked.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Stacked.Get()
}

// GetStackedOk returns a tuple with the Stacked field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Equipment) GetStackedOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Stacked.Get(), o.Stacked.IsSet()
}

// SetStacked sets field value
func (o *Equipment) SetStacked(v int32) {
	o.Stacked.Set(&v)
}

// GetNoted returns the Noted field value
func (o *Equipment) GetNoted() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.Noted
}

// GetNotedOk returns a tuple with the Noted field value
// and a boolean to check if the value has been set.
func (o *Equipment) GetNotedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Noted, true
}

// SetNoted sets field value
func (o *Equipment) SetNoted(v bool) {
	o.Noted = v
}

// GetNoteable returns the Noteable field value
func (o *Equipment) GetNoteable() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.Noteable
}

// GetNoteableOk returns a tuple with the Noteable field value
// and a boolean to check if the value has been set.
func (o *Equipment) GetNoteableOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Noteable, true
}

// SetNoteable sets field value
func (o *Equipment) SetNoteable(v bool) {
	o.Noteable = v
}

// GetLinkedIdItem returns the LinkedIdItem field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *Equipment) GetLinkedIdItem() int32 {
	if o == nil || o.LinkedIdItem.Get() == nil {
		var ret int32
		return ret
	}

	return *o.LinkedIdItem.Get()
}

// GetLinkedIdItemOk returns a tuple with the LinkedIdItem field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Equipment) GetLinkedIdItemOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LinkedIdItem.Get(), o.LinkedIdItem.IsSet()
}

// SetLinkedIdItem sets field value
func (o *Equipment) SetLinkedIdItem(v int32) {
	o.LinkedIdItem.Set(&v)
}

// GetLinkedIdNoted returns the LinkedIdNoted field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *Equipment) GetLinkedIdNoted() int32 {
	if o == nil || o.LinkedIdNoted.Get() == nil {
		var ret int32
		return ret
	}

	return *o.LinkedIdNoted.Get()
}

// GetLinkedIdNotedOk returns a tuple with the LinkedIdNoted field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Equipment) GetLinkedIdNotedOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LinkedIdNoted.Get(), o.LinkedIdNoted.IsSet()
}

// SetLinkedIdNoted sets field value
func (o *Equipment) SetLinkedIdNoted(v int32) {
	o.LinkedIdNoted.Set(&v)
}

// GetLinkedIdPlaceholder returns the LinkedIdPlaceholder field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *Equipment) GetLinkedIdPlaceholder() int32 {
	if o == nil || o.LinkedIdPlaceholder.Get() == nil {
		var ret int32
		return ret
	}

	return *o.LinkedIdPlaceholder.Get()
}

// GetLinkedIdPlaceholderOk returns a tuple with the LinkedIdPlaceholder field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Equipment) GetLinkedIdPlaceholderOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LinkedIdPlaceholder.Get(), o.LinkedIdPlaceholder.IsSet()
}

// SetLinkedIdPlaceholder sets field value
func (o *Equipment) SetLinkedIdPlaceholder(v int32) {
	o.LinkedIdPlaceholder.Set(&v)
}

// GetPlaceholder returns the Placeholder field value
func (o *Equipment) GetPlaceholder() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.Placeholder
}

// GetPlaceholderOk returns a tuple with the Placeholder field value
// and a boolean to check if the value has been set.
func (o *Equipment) GetPlaceholderOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Placeholder, true
}

// SetPlaceholder sets field value
func (o *Equipment) SetPlaceholder(v bool) {
	o.Placeholder = v
}

// GetEquipable returns the Equipable field value
func (o *Equipment) GetEquipable() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.Equipable
}

// GetEquipableOk returns a tuple with the Equipable field value
// and a boolean to check if the value has been set.
func (o *Equipment) GetEquipableOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Equipable, true
}

// SetEquipable sets field value
func (o *Equipment) SetEquipable(v bool) {
	o.Equipable = v
}

// GetEquipableByPlayer returns the EquipableByPlayer field value
func (o *Equipment) GetEquipableByPlayer() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.EquipableByPlayer
}

// GetEquipableByPlayerOk returns a tuple with the EquipableByPlayer field value
// and a boolean to check if the value has been set.
func (o *Equipment) GetEquipableByPlayerOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EquipableByPlayer, true
}

// SetEquipableByPlayer sets field value
func (o *Equipment) SetEquipableByPlayer(v bool) {
	o.EquipableByPlayer = v
}

// GetEquipableWeapon returns the EquipableWeapon field value
func (o *Equipment) GetEquipableWeapon() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.EquipableWeapon
}

// GetEquipableWeaponOk returns a tuple with the EquipableWeapon field value
// and a boolean to check if the value has been set.
func (o *Equipment) GetEquipableWeaponOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EquipableWeapon, true
}

// SetEquipableWeapon sets field value
func (o *Equipment) SetEquipableWeapon(v bool) {
	o.EquipableWeapon = v
}

// GetCost returns the Cost field value
func (o *Equipment) GetCost() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Cost
}

// GetCostOk returns a tuple with the Cost field value
// and a boolean to check if the value has been set.
func (o *Equipment) GetCostOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cost, true
}

// SetCost sets field value
func (o *Equipment) SetCost(v int32) {
	o.Cost = v
}

// GetLowalch returns the Lowalch field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *Equipment) GetLowalch() int32 {
	if o == nil || o.Lowalch.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Lowalch.Get()
}

// GetLowalchOk returns a tuple with the Lowalch field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Equipment) GetLowalchOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Lowalch.Get(), o.Lowalch.IsSet()
}

// SetLowalch sets field value
func (o *Equipment) SetLowalch(v int32) {
	o.Lowalch.Set(&v)
}

// GetHighalch returns the Highalch field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *Equipment) GetHighalch() int32 {
	if o == nil || o.Highalch.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Highalch.Get()
}

// GetHighalchOk returns a tuple with the Highalch field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Equipment) GetHighalchOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Highalch.Get(), o.Highalch.IsSet()
}

// SetHighalch sets field value
func (o *Equipment) SetHighalch(v int32) {
	o.Highalch.Set(&v)
}

// GetWeight returns the Weight field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *Equipment) GetWeight() float32 {
	if o == nil || o.Weight.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Weight.Get()
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Equipment) GetWeightOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Weight.Get(), o.Weight.IsSet()
}

// SetWeight sets field value
func (o *Equipment) SetWeight(v float32) {
	o.Weight.Set(&v)
}

// GetBuyLimit returns the BuyLimit field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *Equipment) GetBuyLimit() int32 {
	if o == nil || o.BuyLimit.Get() == nil {
		var ret int32
		return ret
	}

	return *o.BuyLimit.Get()
}

// GetBuyLimitOk returns a tuple with the BuyLimit field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Equipment) GetBuyLimitOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BuyLimit.Get(), o.BuyLimit.IsSet()
}

// SetBuyLimit sets field value
func (o *Equipment) SetBuyLimit(v int32) {
	o.BuyLimit.Set(&v)
}

// GetQuestItem returns the QuestItem field value
func (o *Equipment) GetQuestItem() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.QuestItem
}

// GetQuestItemOk returns a tuple with the QuestItem field value
// and a boolean to check if the value has been set.
func (o *Equipment) GetQuestItemOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.QuestItem, true
}

// SetQuestItem sets field value
func (o *Equipment) SetQuestItem(v bool) {
	o.QuestItem = v
}

// GetReleaseDate returns the ReleaseDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Equipment) GetReleaseDate() string {
	if o == nil || o.ReleaseDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.ReleaseDate.Get()
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Equipment) GetReleaseDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ReleaseDate.Get(), o.ReleaseDate.IsSet()
}

// SetReleaseDate sets field value
func (o *Equipment) SetReleaseDate(v string) {
	o.ReleaseDate.Set(&v)
}

// GetDuplicate returns the Duplicate field value
func (o *Equipment) GetDuplicate() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.Duplicate
}

// GetDuplicateOk returns a tuple with the Duplicate field value
// and a boolean to check if the value has been set.
func (o *Equipment) GetDuplicateOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Duplicate, true
}

// SetDuplicate sets field value
func (o *Equipment) SetDuplicate(v bool) {
	o.Duplicate = v
}

// GetExamine returns the Examine field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Equipment) GetExamine() string {
	if o == nil || o.Examine.Get() == nil {
		var ret string
		return ret
	}

	return *o.Examine.Get()
}

// GetExamineOk returns a tuple with the Examine field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Equipment) GetExamineOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Examine.Get(), o.Examine.IsSet()
}

// SetExamine sets field value
func (o *Equipment) SetExamine(v string) {
	o.Examine.Set(&v)
}

// GetIcon returns the Icon field value
func (o *Equipment) GetIcon() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Icon
}

// GetIconOk returns a tuple with the Icon field value
// and a boolean to check if the value has been set.
func (o *Equipment) GetIconOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Icon, true
}

// SetIcon sets field value
func (o *Equipment) SetIcon(v string) {
	o.Icon = v
}

// GetWikiName returns the WikiName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Equipment) GetWikiName() string {
	if o == nil || o.WikiName.Get() == nil {
		var ret string
		return ret
	}

	return *o.WikiName.Get()
}

// GetWikiNameOk returns a tuple with the WikiName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Equipment) GetWikiNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WikiName.Get(), o.WikiName.IsSet()
}

// SetWikiName sets field value
func (o *Equipment) SetWikiName(v string) {
	o.WikiName.Set(&v)
}

// GetWikiUrl returns the WikiUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Equipment) GetWikiUrl() string {
	if o == nil || o.WikiUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.WikiUrl.Get()
}

// GetWikiUrlOk returns a tuple with the WikiUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Equipment) GetWikiUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WikiUrl.Get(), o.WikiUrl.IsSet()
}

// SetWikiUrl sets field value
func (o *Equipment) SetWikiUrl(v string) {
	o.WikiUrl.Set(&v)
}

// GetWikiExchange returns the WikiExchange field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Equipment) GetWikiExchange() string {
	if o == nil || o.WikiExchange.Get() == nil {
		var ret string
		return ret
	}

	return *o.WikiExchange.Get()
}

// GetWikiExchangeOk returns a tuple with the WikiExchange field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Equipment) GetWikiExchangeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WikiExchange.Get(), o.WikiExchange.IsSet()
}

// SetWikiExchange sets field value
func (o *Equipment) SetWikiExchange(v string) {
	o.WikiExchange.Set(&v)
}

// GetEquipment returns the Equipment field value
// If the value is explicit nil, the zero value for ItemEquipment will be returned
func (o *Equipment) GetEquipment() ItemEquipment {
	if o == nil || o.Equipment.Get() == nil {
		var ret ItemEquipment
		return ret
	}

	return *o.Equipment.Get()
}

// GetEquipmentOk returns a tuple with the Equipment field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Equipment) GetEquipmentOk() (*ItemEquipment, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Equipment.Get(), o.Equipment.IsSet()
}

// SetEquipment sets field value
func (o *Equipment) SetEquipment(v ItemEquipment) {
	o.Equipment.Set(&v)
}

// GetWeapon returns the Weapon field value
// If the value is explicit nil, the zero value for ItemWeapon will be returned
func (o *Equipment) GetWeapon() ItemWeapon {
	if o == nil || o.Weapon.Get() == nil {
		var ret ItemWeapon
		return ret
	}

	return *o.Weapon.Get()
}

// GetWeaponOk returns a tuple with the Weapon field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Equipment) GetWeaponOk() (*ItemWeapon, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Weapon.Get(), o.Weapon.IsSet()
}

// SetWeapon sets field value
func (o *Equipment) SetWeapon(v ItemWeapon) {
	o.Weapon.Set(&v)
}

func (o Equipment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["incomplete"] = o.Incomplete
	}
	if true {
		toSerialize["members"] = o.Members
	}
	if true {
		toSerialize["tradeable"] = o.Tradeable
	}
	if true {
		toSerialize["tradeable_on_ge"] = o.TradeableOnGe
	}
	if true {
		toSerialize["stackable"] = o.Stackable
	}
	if true {
		toSerialize["stacked"] = o.Stacked.Get()
	}
	if true {
		toSerialize["noted"] = o.Noted
	}
	if true {
		toSerialize["noteable"] = o.Noteable
	}
	if true {
		toSerialize["linked_id_item"] = o.LinkedIdItem.Get()
	}
	if true {
		toSerialize["linked_id_noted"] = o.LinkedIdNoted.Get()
	}
	if true {
		toSerialize["linked_id_placeholder"] = o.LinkedIdPlaceholder.Get()
	}
	if true {
		toSerialize["placeholder"] = o.Placeholder
	}
	if true {
		toSerialize["equipable"] = o.Equipable
	}
	if true {
		toSerialize["equipable_by_player"] = o.EquipableByPlayer
	}
	if true {
		toSerialize["equipable_weapon"] = o.EquipableWeapon
	}
	if true {
		toSerialize["cost"] = o.Cost
	}
	if true {
		toSerialize["lowalch"] = o.Lowalch.Get()
	}
	if true {
		toSerialize["highalch"] = o.Highalch.Get()
	}
	if true {
		toSerialize["weight"] = o.Weight.Get()
	}
	if true {
		toSerialize["buy_limit"] = o.BuyLimit.Get()
	}
	if true {
		toSerialize["quest_item"] = o.QuestItem
	}
	if true {
		toSerialize["release_date"] = o.ReleaseDate.Get()
	}
	if true {
		toSerialize["duplicate"] = o.Duplicate
	}
	if true {
		toSerialize["examine"] = o.Examine.Get()
	}
	if true {
		toSerialize["icon"] = o.Icon
	}
	if true {
		toSerialize["wiki_name"] = o.WikiName.Get()
	}
	if true {
		toSerialize["wiki_url"] = o.WikiUrl.Get()
	}
	if true {
		toSerialize["wiki_exchange"] = o.WikiExchange.Get()
	}
	if true {
		toSerialize["equipment"] = o.Equipment.Get()
	}
	if true {
		toSerialize["weapon"] = o.Weapon.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableEquipment struct {
	value *Equipment
	isSet bool
}

func (v NullableEquipment) Get() *Equipment {
	return v.value
}

func (v *NullableEquipment) Set(val *Equipment) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipment) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipment(val *Equipment) *NullableEquipment {
	return &NullableEquipment{value: val, isSet: true}
}

func (v NullableEquipment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


