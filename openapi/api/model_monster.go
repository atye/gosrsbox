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

// Monster struct for Monster
type Monster struct {
	// Unique OSRS monster ID number.
	Id string `json:"id"`
	// The name of the monster.
	Name string `json:"name"`
	// The last time (UTC) the monster was updated (in ISO8601 date format).
	LastUpdated NullableString `json:"last_updated"`
	// If the monster has incomplete wiki data.
	Incomplete bool `json:"incomplete"`
	// If the monster is members only, or not.
	Members bool `json:"members"`
	// The release date of the monster (in ISO8601 date format).
	ReleaseDate NullableString `json:"release_date"`
	// The combat level of the monster.
	CombatLevel int32 `json:"combat_level"`
	// The size, in tiles, of the monster.
	Size int32 `json:"size"`
	// The number of hitpoints a monster has.
	Hitpoints NullableInt32 `json:"hitpoints"`
	// The maximum hit of the monster.
	MaxHit NullableInt32 `json:"max_hit"`
	// The attack style (e.g., melee, magic, range) of the monster.
	AttackType []string `json:"attack_type"`
	// The attack speed (in game ticks) of the monster.
	AttackSpeed NullableInt32 `json:"attack_speed"`
	// If the monster is aggressive, or not.
	Aggressive bool `json:"aggressive"`
	// If the monster poisons, or not
	Poisonous bool `json:"poisonous"`
	// If the monster poisons using venom, or not
	Venomous bool `json:"venomous"`
	// If the monster is immune to poison, or not
	ImmunePoison bool `json:"immune_poison"`
	// If the monster is immune to venom, or not
	ImmuneVenom bool `json:"immune_venom"`
	// An array of monster attributes.
	Attributes []string `json:"attributes"`
	// An array of monster category.
	Category []string `json:"category"`
	// If the monster is a potential slayer task.
	SlayerMonster bool `json:"slayer_monster"`
	// The slayer level required to kill the monster.
	SlayerLevel NullableInt32 `json:"slayer_level"`
	// The slayer XP rewarded for a monster kill.
	SlayerXp NullableFloat32 `json:"slayer_xp"`
	// The slayer masters who can assign the monster.
	SlayerMasters []string `json:"slayer_masters"`
	// If the monster is a duplicate.
	Duplicate bool `json:"duplicate"`
	// The examine text of the monster.
	Examine string `json:"examine"`
	// The OSRS Wiki name for the monster.
	WikiName string `json:"wiki_name"`
	// The OSRS Wiki URL (possibly including anchor link).
	WikiUrl string `json:"wiki_url"`
	// The attack level of the monster.
	AttackLevel int32 `json:"attack_level"`
	// The strength level of the monster.
	StrengthLevel int32 `json:"strength_level"`
	// The defence level of the monster.
	DefenceLevel int32 `json:"defence_level"`
	// The magic level of the monster.
	MagicLevel int32 `json:"magic_level"`
	// The ranged level of the monster.
	RangedLevel int32 `json:"ranged_level"`
	// The attack bonus of the monster.
	AttackBonus int32 `json:"attack_bonus"`
	// The strength bonus of the monster.
	StrengthBonus int32 `json:"strength_bonus"`
	// The magic attack of the monster.
	AttackMagic int32 `json:"attack_magic"`
	// The magic bonus of the monster.
	MagicBonus int32 `json:"magic_bonus"`
	// The ranged attack of the monster.
	AttackRanged int32 `json:"attack_ranged"`
	// The ranged bonus of the monster.
	RangedBonus int32 `json:"ranged_bonus"`
	// The defence stab bonus of the monster.
	DefenceStab int32 `json:"defence_stab"`
	// The defence slash bonus of the monster.
	DefenceSlash int32 `json:"defence_slash"`
	// The defence crush bonus of the monster.
	DefenceCrush int32 `json:"defence_crush"`
	// The defence magic bonus of the monster.
	DefenceMagic int32 `json:"defence_magic"`
	// The defence ranged bonus of the monster.
	DefenceRanged int32 `json:"defence_ranged"`
	// An array of monster drop objects.
	Drops []MonsterDrops `json:"drops"`
}

// NewMonster instantiates a new Monster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonster(id string, name string, lastUpdated NullableString, incomplete bool, members bool, releaseDate NullableString, combatLevel int32, size int32, hitpoints NullableInt32, maxHit NullableInt32, attackType []string, attackSpeed NullableInt32, aggressive bool, poisonous bool, venomous bool, immunePoison bool, immuneVenom bool, attributes []string, category []string, slayerMonster bool, slayerLevel NullableInt32, slayerXp NullableFloat32, slayerMasters []string, duplicate bool, examine string, wikiName string, wikiUrl string, attackLevel int32, strengthLevel int32, defenceLevel int32, magicLevel int32, rangedLevel int32, attackBonus int32, strengthBonus int32, attackMagic int32, magicBonus int32, attackRanged int32, rangedBonus int32, defenceStab int32, defenceSlash int32, defenceCrush int32, defenceMagic int32, defenceRanged int32, drops []MonsterDrops) *Monster {
	this := Monster{}
	this.Id = id
	this.Name = name
	this.LastUpdated = lastUpdated
	this.Incomplete = incomplete
	this.Members = members
	this.ReleaseDate = releaseDate
	this.CombatLevel = combatLevel
	this.Size = size
	this.Hitpoints = hitpoints
	this.MaxHit = maxHit
	this.AttackType = attackType
	this.AttackSpeed = attackSpeed
	this.Aggressive = aggressive
	this.Poisonous = poisonous
	this.Venomous = venomous
	this.ImmunePoison = immunePoison
	this.ImmuneVenom = immuneVenom
	this.Attributes = attributes
	this.Category = category
	this.SlayerMonster = slayerMonster
	this.SlayerLevel = slayerLevel
	this.SlayerXp = slayerXp
	this.SlayerMasters = slayerMasters
	this.Duplicate = duplicate
	this.Examine = examine
	this.WikiName = wikiName
	this.WikiUrl = wikiUrl
	this.AttackLevel = attackLevel
	this.StrengthLevel = strengthLevel
	this.DefenceLevel = defenceLevel
	this.MagicLevel = magicLevel
	this.RangedLevel = rangedLevel
	this.AttackBonus = attackBonus
	this.StrengthBonus = strengthBonus
	this.AttackMagic = attackMagic
	this.MagicBonus = magicBonus
	this.AttackRanged = attackRanged
	this.RangedBonus = rangedBonus
	this.DefenceStab = defenceStab
	this.DefenceSlash = defenceSlash
	this.DefenceCrush = defenceCrush
	this.DefenceMagic = defenceMagic
	this.DefenceRanged = defenceRanged
	this.Drops = drops
	return &this
}

// NewMonsterWithDefaults instantiates a new Monster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonsterWithDefaults() *Monster {
	this := Monster{}
	return &this
}

// GetId returns the Id field value
func (o *Monster) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Monster) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Monster) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Monster) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Monster) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Monster) SetName(v string) {
	o.Name = v
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Monster) GetLastUpdated() string {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret string
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Monster) GetLastUpdatedOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *Monster) SetLastUpdated(v string) {
	o.LastUpdated.Set(&v)
}

// GetIncomplete returns the Incomplete field value
func (o *Monster) GetIncomplete() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Incomplete
}

// GetIncompleteOk returns a tuple with the Incomplete field value
// and a boolean to check if the value has been set.
func (o *Monster) GetIncompleteOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Incomplete, true
}

// SetIncomplete sets field value
func (o *Monster) SetIncomplete(v bool) {
	o.Incomplete = v
}

// GetMembers returns the Members field value
func (o *Monster) GetMembers() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Members
}

// GetMembersOk returns a tuple with the Members field value
// and a boolean to check if the value has been set.
func (o *Monster) GetMembersOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Members, true
}

// SetMembers sets field value
func (o *Monster) SetMembers(v bool) {
	o.Members = v
}

// GetReleaseDate returns the ReleaseDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Monster) GetReleaseDate() string {
	if o == nil || o.ReleaseDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.ReleaseDate.Get()
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Monster) GetReleaseDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ReleaseDate.Get(), o.ReleaseDate.IsSet()
}

// SetReleaseDate sets field value
func (o *Monster) SetReleaseDate(v string) {
	o.ReleaseDate.Set(&v)
}

// GetCombatLevel returns the CombatLevel field value
func (o *Monster) GetCombatLevel() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CombatLevel
}

// GetCombatLevelOk returns a tuple with the CombatLevel field value
// and a boolean to check if the value has been set.
func (o *Monster) GetCombatLevelOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CombatLevel, true
}

// SetCombatLevel sets field value
func (o *Monster) SetCombatLevel(v int32) {
	o.CombatLevel = v
}

// GetSize returns the Size field value
func (o *Monster) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *Monster) GetSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *Monster) SetSize(v int32) {
	o.Size = v
}

// GetHitpoints returns the Hitpoints field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *Monster) GetHitpoints() int32 {
	if o == nil || o.Hitpoints.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Hitpoints.Get()
}

// GetHitpointsOk returns a tuple with the Hitpoints field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Monster) GetHitpointsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Hitpoints.Get(), o.Hitpoints.IsSet()
}

// SetHitpoints sets field value
func (o *Monster) SetHitpoints(v int32) {
	o.Hitpoints.Set(&v)
}

// GetMaxHit returns the MaxHit field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *Monster) GetMaxHit() int32 {
	if o == nil || o.MaxHit.Get() == nil {
		var ret int32
		return ret
	}

	return *o.MaxHit.Get()
}

// GetMaxHitOk returns a tuple with the MaxHit field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Monster) GetMaxHitOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MaxHit.Get(), o.MaxHit.IsSet()
}

// SetMaxHit sets field value
func (o *Monster) SetMaxHit(v int32) {
	o.MaxHit.Set(&v)
}

// GetAttackType returns the AttackType field value
func (o *Monster) GetAttackType() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AttackType
}

// GetAttackTypeOk returns a tuple with the AttackType field value
// and a boolean to check if the value has been set.
func (o *Monster) GetAttackTypeOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AttackType, true
}

// SetAttackType sets field value
func (o *Monster) SetAttackType(v []string) {
	o.AttackType = v
}

// GetAttackSpeed returns the AttackSpeed field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *Monster) GetAttackSpeed() int32 {
	if o == nil || o.AttackSpeed.Get() == nil {
		var ret int32
		return ret
	}

	return *o.AttackSpeed.Get()
}

// GetAttackSpeedOk returns a tuple with the AttackSpeed field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Monster) GetAttackSpeedOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AttackSpeed.Get(), o.AttackSpeed.IsSet()
}

// SetAttackSpeed sets field value
func (o *Monster) SetAttackSpeed(v int32) {
	o.AttackSpeed.Set(&v)
}

// GetAggressive returns the Aggressive field value
func (o *Monster) GetAggressive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Aggressive
}

// GetAggressiveOk returns a tuple with the Aggressive field value
// and a boolean to check if the value has been set.
func (o *Monster) GetAggressiveOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Aggressive, true
}

// SetAggressive sets field value
func (o *Monster) SetAggressive(v bool) {
	o.Aggressive = v
}

// GetPoisonous returns the Poisonous field value
func (o *Monster) GetPoisonous() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Poisonous
}

// GetPoisonousOk returns a tuple with the Poisonous field value
// and a boolean to check if the value has been set.
func (o *Monster) GetPoisonousOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Poisonous, true
}

// SetPoisonous sets field value
func (o *Monster) SetPoisonous(v bool) {
	o.Poisonous = v
}

// GetVenomous returns the Venomous field value
func (o *Monster) GetVenomous() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Venomous
}

// GetVenomousOk returns a tuple with the Venomous field value
// and a boolean to check if the value has been set.
func (o *Monster) GetVenomousOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Venomous, true
}

// SetVenomous sets field value
func (o *Monster) SetVenomous(v bool) {
	o.Venomous = v
}

// GetImmunePoison returns the ImmunePoison field value
func (o *Monster) GetImmunePoison() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ImmunePoison
}

// GetImmunePoisonOk returns a tuple with the ImmunePoison field value
// and a boolean to check if the value has been set.
func (o *Monster) GetImmunePoisonOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ImmunePoison, true
}

// SetImmunePoison sets field value
func (o *Monster) SetImmunePoison(v bool) {
	o.ImmunePoison = v
}

// GetImmuneVenom returns the ImmuneVenom field value
func (o *Monster) GetImmuneVenom() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ImmuneVenom
}

// GetImmuneVenomOk returns a tuple with the ImmuneVenom field value
// and a boolean to check if the value has been set.
func (o *Monster) GetImmuneVenomOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ImmuneVenom, true
}

// SetImmuneVenom sets field value
func (o *Monster) SetImmuneVenom(v bool) {
	o.ImmuneVenom = v
}

// GetAttributes returns the Attributes field value
func (o *Monster) GetAttributes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *Monster) GetAttributesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *Monster) SetAttributes(v []string) {
	o.Attributes = v
}

// GetCategory returns the Category field value
func (o *Monster) GetCategory() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *Monster) GetCategoryOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *Monster) SetCategory(v []string) {
	o.Category = v
}

// GetSlayerMonster returns the SlayerMonster field value
func (o *Monster) GetSlayerMonster() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SlayerMonster
}

// GetSlayerMonsterOk returns a tuple with the SlayerMonster field value
// and a boolean to check if the value has been set.
func (o *Monster) GetSlayerMonsterOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SlayerMonster, true
}

// SetSlayerMonster sets field value
func (o *Monster) SetSlayerMonster(v bool) {
	o.SlayerMonster = v
}

// GetSlayerLevel returns the SlayerLevel field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *Monster) GetSlayerLevel() int32 {
	if o == nil || o.SlayerLevel.Get() == nil {
		var ret int32
		return ret
	}

	return *o.SlayerLevel.Get()
}

// GetSlayerLevelOk returns a tuple with the SlayerLevel field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Monster) GetSlayerLevelOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SlayerLevel.Get(), o.SlayerLevel.IsSet()
}

// SetSlayerLevel sets field value
func (o *Monster) SetSlayerLevel(v int32) {
	o.SlayerLevel.Set(&v)
}

// GetSlayerXp returns the SlayerXp field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *Monster) GetSlayerXp() float32 {
	if o == nil || o.SlayerXp.Get() == nil {
		var ret float32
		return ret
	}

	return *o.SlayerXp.Get()
}

// GetSlayerXpOk returns a tuple with the SlayerXp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Monster) GetSlayerXpOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SlayerXp.Get(), o.SlayerXp.IsSet()
}

// SetSlayerXp sets field value
func (o *Monster) SetSlayerXp(v float32) {
	o.SlayerXp.Set(&v)
}

// GetSlayerMasters returns the SlayerMasters field value
func (o *Monster) GetSlayerMasters() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SlayerMasters
}

// GetSlayerMastersOk returns a tuple with the SlayerMasters field value
// and a boolean to check if the value has been set.
func (o *Monster) GetSlayerMastersOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SlayerMasters, true
}

// SetSlayerMasters sets field value
func (o *Monster) SetSlayerMasters(v []string) {
	o.SlayerMasters = v
}

// GetDuplicate returns the Duplicate field value
func (o *Monster) GetDuplicate() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Duplicate
}

// GetDuplicateOk returns a tuple with the Duplicate field value
// and a boolean to check if the value has been set.
func (o *Monster) GetDuplicateOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Duplicate, true
}

// SetDuplicate sets field value
func (o *Monster) SetDuplicate(v bool) {
	o.Duplicate = v
}

// GetExamine returns the Examine field value
func (o *Monster) GetExamine() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Examine
}

// GetExamineOk returns a tuple with the Examine field value
// and a boolean to check if the value has been set.
func (o *Monster) GetExamineOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Examine, true
}

// SetExamine sets field value
func (o *Monster) SetExamine(v string) {
	o.Examine = v
}

// GetWikiName returns the WikiName field value
func (o *Monster) GetWikiName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WikiName
}

// GetWikiNameOk returns a tuple with the WikiName field value
// and a boolean to check if the value has been set.
func (o *Monster) GetWikiNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WikiName, true
}

// SetWikiName sets field value
func (o *Monster) SetWikiName(v string) {
	o.WikiName = v
}

// GetWikiUrl returns the WikiUrl field value
func (o *Monster) GetWikiUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WikiUrl
}

// GetWikiUrlOk returns a tuple with the WikiUrl field value
// and a boolean to check if the value has been set.
func (o *Monster) GetWikiUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WikiUrl, true
}

// SetWikiUrl sets field value
func (o *Monster) SetWikiUrl(v string) {
	o.WikiUrl = v
}

// GetAttackLevel returns the AttackLevel field value
func (o *Monster) GetAttackLevel() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AttackLevel
}

// GetAttackLevelOk returns a tuple with the AttackLevel field value
// and a boolean to check if the value has been set.
func (o *Monster) GetAttackLevelOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AttackLevel, true
}

// SetAttackLevel sets field value
func (o *Monster) SetAttackLevel(v int32) {
	o.AttackLevel = v
}

// GetStrengthLevel returns the StrengthLevel field value
func (o *Monster) GetStrengthLevel() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StrengthLevel
}

// GetStrengthLevelOk returns a tuple with the StrengthLevel field value
// and a boolean to check if the value has been set.
func (o *Monster) GetStrengthLevelOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StrengthLevel, true
}

// SetStrengthLevel sets field value
func (o *Monster) SetStrengthLevel(v int32) {
	o.StrengthLevel = v
}

// GetDefenceLevel returns the DefenceLevel field value
func (o *Monster) GetDefenceLevel() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DefenceLevel
}

// GetDefenceLevelOk returns a tuple with the DefenceLevel field value
// and a boolean to check if the value has been set.
func (o *Monster) GetDefenceLevelOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DefenceLevel, true
}

// SetDefenceLevel sets field value
func (o *Monster) SetDefenceLevel(v int32) {
	o.DefenceLevel = v
}

// GetMagicLevel returns the MagicLevel field value
func (o *Monster) GetMagicLevel() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MagicLevel
}

// GetMagicLevelOk returns a tuple with the MagicLevel field value
// and a boolean to check if the value has been set.
func (o *Monster) GetMagicLevelOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MagicLevel, true
}

// SetMagicLevel sets field value
func (o *Monster) SetMagicLevel(v int32) {
	o.MagicLevel = v
}

// GetRangedLevel returns the RangedLevel field value
func (o *Monster) GetRangedLevel() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RangedLevel
}

// GetRangedLevelOk returns a tuple with the RangedLevel field value
// and a boolean to check if the value has been set.
func (o *Monster) GetRangedLevelOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RangedLevel, true
}

// SetRangedLevel sets field value
func (o *Monster) SetRangedLevel(v int32) {
	o.RangedLevel = v
}

// GetAttackBonus returns the AttackBonus field value
func (o *Monster) GetAttackBonus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AttackBonus
}

// GetAttackBonusOk returns a tuple with the AttackBonus field value
// and a boolean to check if the value has been set.
func (o *Monster) GetAttackBonusOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AttackBonus, true
}

// SetAttackBonus sets field value
func (o *Monster) SetAttackBonus(v int32) {
	o.AttackBonus = v
}

// GetStrengthBonus returns the StrengthBonus field value
func (o *Monster) GetStrengthBonus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StrengthBonus
}

// GetStrengthBonusOk returns a tuple with the StrengthBonus field value
// and a boolean to check if the value has been set.
func (o *Monster) GetStrengthBonusOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StrengthBonus, true
}

// SetStrengthBonus sets field value
func (o *Monster) SetStrengthBonus(v int32) {
	o.StrengthBonus = v
}

// GetAttackMagic returns the AttackMagic field value
func (o *Monster) GetAttackMagic() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AttackMagic
}

// GetAttackMagicOk returns a tuple with the AttackMagic field value
// and a boolean to check if the value has been set.
func (o *Monster) GetAttackMagicOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AttackMagic, true
}

// SetAttackMagic sets field value
func (o *Monster) SetAttackMagic(v int32) {
	o.AttackMagic = v
}

// GetMagicBonus returns the MagicBonus field value
func (o *Monster) GetMagicBonus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MagicBonus
}

// GetMagicBonusOk returns a tuple with the MagicBonus field value
// and a boolean to check if the value has been set.
func (o *Monster) GetMagicBonusOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MagicBonus, true
}

// SetMagicBonus sets field value
func (o *Monster) SetMagicBonus(v int32) {
	o.MagicBonus = v
}

// GetAttackRanged returns the AttackRanged field value
func (o *Monster) GetAttackRanged() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AttackRanged
}

// GetAttackRangedOk returns a tuple with the AttackRanged field value
// and a boolean to check if the value has been set.
func (o *Monster) GetAttackRangedOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AttackRanged, true
}

// SetAttackRanged sets field value
func (o *Monster) SetAttackRanged(v int32) {
	o.AttackRanged = v
}

// GetRangedBonus returns the RangedBonus field value
func (o *Monster) GetRangedBonus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RangedBonus
}

// GetRangedBonusOk returns a tuple with the RangedBonus field value
// and a boolean to check if the value has been set.
func (o *Monster) GetRangedBonusOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RangedBonus, true
}

// SetRangedBonus sets field value
func (o *Monster) SetRangedBonus(v int32) {
	o.RangedBonus = v
}

// GetDefenceStab returns the DefenceStab field value
func (o *Monster) GetDefenceStab() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DefenceStab
}

// GetDefenceStabOk returns a tuple with the DefenceStab field value
// and a boolean to check if the value has been set.
func (o *Monster) GetDefenceStabOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DefenceStab, true
}

// SetDefenceStab sets field value
func (o *Monster) SetDefenceStab(v int32) {
	o.DefenceStab = v
}

// GetDefenceSlash returns the DefenceSlash field value
func (o *Monster) GetDefenceSlash() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DefenceSlash
}

// GetDefenceSlashOk returns a tuple with the DefenceSlash field value
// and a boolean to check if the value has been set.
func (o *Monster) GetDefenceSlashOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DefenceSlash, true
}

// SetDefenceSlash sets field value
func (o *Monster) SetDefenceSlash(v int32) {
	o.DefenceSlash = v
}

// GetDefenceCrush returns the DefenceCrush field value
func (o *Monster) GetDefenceCrush() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DefenceCrush
}

// GetDefenceCrushOk returns a tuple with the DefenceCrush field value
// and a boolean to check if the value has been set.
func (o *Monster) GetDefenceCrushOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DefenceCrush, true
}

// SetDefenceCrush sets field value
func (o *Monster) SetDefenceCrush(v int32) {
	o.DefenceCrush = v
}

// GetDefenceMagic returns the DefenceMagic field value
func (o *Monster) GetDefenceMagic() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DefenceMagic
}

// GetDefenceMagicOk returns a tuple with the DefenceMagic field value
// and a boolean to check if the value has been set.
func (o *Monster) GetDefenceMagicOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DefenceMagic, true
}

// SetDefenceMagic sets field value
func (o *Monster) SetDefenceMagic(v int32) {
	o.DefenceMagic = v
}

// GetDefenceRanged returns the DefenceRanged field value
func (o *Monster) GetDefenceRanged() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DefenceRanged
}

// GetDefenceRangedOk returns a tuple with the DefenceRanged field value
// and a boolean to check if the value has been set.
func (o *Monster) GetDefenceRangedOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DefenceRanged, true
}

// SetDefenceRanged sets field value
func (o *Monster) SetDefenceRanged(v int32) {
	o.DefenceRanged = v
}

// GetDrops returns the Drops field value
func (o *Monster) GetDrops() []MonsterDrops {
	if o == nil {
		var ret []MonsterDrops
		return ret
	}

	return o.Drops
}

// GetDropsOk returns a tuple with the Drops field value
// and a boolean to check if the value has been set.
func (o *Monster) GetDropsOk() (*[]MonsterDrops, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Drops, true
}

// SetDrops sets field value
func (o *Monster) SetDrops(v []MonsterDrops) {
	o.Drops = v
}

func (o Monster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["last_updated"] = o.LastUpdated.Get()
	}
	if true {
		toSerialize["incomplete"] = o.Incomplete
	}
	if true {
		toSerialize["members"] = o.Members
	}
	if true {
		toSerialize["release_date"] = o.ReleaseDate.Get()
	}
	if true {
		toSerialize["combat_level"] = o.CombatLevel
	}
	if true {
		toSerialize["size"] = o.Size
	}
	if true {
		toSerialize["hitpoints"] = o.Hitpoints.Get()
	}
	if true {
		toSerialize["max_hit"] = o.MaxHit.Get()
	}
	if true {
		toSerialize["attack_type"] = o.AttackType
	}
	if true {
		toSerialize["attack_speed"] = o.AttackSpeed.Get()
	}
	if true {
		toSerialize["aggressive"] = o.Aggressive
	}
	if true {
		toSerialize["poisonous"] = o.Poisonous
	}
	if true {
		toSerialize["venomous"] = o.Venomous
	}
	if true {
		toSerialize["immune_poison"] = o.ImmunePoison
	}
	if true {
		toSerialize["immune_venom"] = o.ImmuneVenom
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["category"] = o.Category
	}
	if true {
		toSerialize["slayer_monster"] = o.SlayerMonster
	}
	if true {
		toSerialize["slayer_level"] = o.SlayerLevel.Get()
	}
	if true {
		toSerialize["slayer_xp"] = o.SlayerXp.Get()
	}
	if true {
		toSerialize["slayer_masters"] = o.SlayerMasters
	}
	if true {
		toSerialize["duplicate"] = o.Duplicate
	}
	if true {
		toSerialize["examine"] = o.Examine
	}
	if true {
		toSerialize["wiki_name"] = o.WikiName
	}
	if true {
		toSerialize["wiki_url"] = o.WikiUrl
	}
	if true {
		toSerialize["attack_level"] = o.AttackLevel
	}
	if true {
		toSerialize["strength_level"] = o.StrengthLevel
	}
	if true {
		toSerialize["defence_level"] = o.DefenceLevel
	}
	if true {
		toSerialize["magic_level"] = o.MagicLevel
	}
	if true {
		toSerialize["ranged_level"] = o.RangedLevel
	}
	if true {
		toSerialize["attack_bonus"] = o.AttackBonus
	}
	if true {
		toSerialize["strength_bonus"] = o.StrengthBonus
	}
	if true {
		toSerialize["attack_magic"] = o.AttackMagic
	}
	if true {
		toSerialize["magic_bonus"] = o.MagicBonus
	}
	if true {
		toSerialize["attack_ranged"] = o.AttackRanged
	}
	if true {
		toSerialize["ranged_bonus"] = o.RangedBonus
	}
	if true {
		toSerialize["defence_stab"] = o.DefenceStab
	}
	if true {
		toSerialize["defence_slash"] = o.DefenceSlash
	}
	if true {
		toSerialize["defence_crush"] = o.DefenceCrush
	}
	if true {
		toSerialize["defence_magic"] = o.DefenceMagic
	}
	if true {
		toSerialize["defence_ranged"] = o.DefenceRanged
	}
	if true {
		toSerialize["drops"] = o.Drops
	}
	return json.Marshal(toSerialize)
}

type NullableMonster struct {
	value *Monster
	isSet bool
}

func (v NullableMonster) Get() *Monster {
	return v.value
}

func (v *NullableMonster) Set(val *Monster) {
	v.value = val
	v.isSet = true
}

func (v NullableMonster) IsSet() bool {
	return v.isSet
}

func (v *NullableMonster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonster(val *Monster) *NullableMonster {
	return &NullableMonster{value: val, isSet: true}
}

func (v NullableMonster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


