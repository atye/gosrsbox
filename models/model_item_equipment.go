/*
 * osrsbox-api
 *
 * An open, free, complete and up-to-date RESTful API for Old School RuneScape (OSRS) items, monsters and prayers.
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
)

// ItemEquipment The equipment bonuses of equipable armour/weapons.
type ItemEquipment struct {
	// The attack stab bonus of the item.
	AttackStab int32 `json:"attack_stab"`
	// The attack slash bonus of the item.
	AttackSlash int32 `json:"attack_slash"`
	// The attack crush bonus of the item.
	AttackCrush int32 `json:"attack_crush"`
	// The attack magic bonus of the item.
	AttackMagic int32 `json:"attack_magic"`
	// The attack ranged bonus of the item.
	AttackRanged int32 `json:"attack_ranged"`
	// The defence stab bonus of the item.
	DefenceStab int32 `json:"defence_stab"`
	// The defence slash bonus of the item.
	DefenceSlash int32 `json:"defence_slash"`
	// The defence crush bonus of the item.
	DefenceCrush int32 `json:"defence_crush"`
	// The defence magic bonus of the item.
	DefenceMagic int32 `json:"defence_magic"`
	// The defence ranged bonus of the item.
	DefenceRanged int32 `json:"defence_ranged"`
	// The melee strength bonus of the item.
	MeleeStrength int32 `json:"melee_strength"`
	// The ranged strength bonus of the item.
	RangedStrength int32 `json:"ranged_strength"`
	// The magic damage bonus of the item.
	MagicDamage int32 `json:"magic_damage"`
	// The prayer bonus of the item.
	Prayer int32 `json:"prayer"`
	// The equipment slot associated with the item (e.g., head).
	Slot string `json:"slot"`
	// An object of requirements {skill: level}.
	Requirements map[string]interface{} `json:"requirements"`
}

// NewItemEquipment instantiates a new ItemEquipment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemEquipment(attackStab int32, attackSlash int32, attackCrush int32, attackMagic int32, attackRanged int32, defenceStab int32, defenceSlash int32, defenceCrush int32, defenceMagic int32, defenceRanged int32, meleeStrength int32, rangedStrength int32, magicDamage int32, prayer int32, slot string, requirements map[string]interface{}) *ItemEquipment {
	this := ItemEquipment{}
	this.AttackStab = attackStab
	this.AttackSlash = attackSlash
	this.AttackCrush = attackCrush
	this.AttackMagic = attackMagic
	this.AttackRanged = attackRanged
	this.DefenceStab = defenceStab
	this.DefenceSlash = defenceSlash
	this.DefenceCrush = defenceCrush
	this.DefenceMagic = defenceMagic
	this.DefenceRanged = defenceRanged
	this.MeleeStrength = meleeStrength
	this.RangedStrength = rangedStrength
	this.MagicDamage = magicDamage
	this.Prayer = prayer
	this.Slot = slot
	this.Requirements = requirements
	return &this
}

// NewItemEquipmentWithDefaults instantiates a new ItemEquipment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemEquipmentWithDefaults() *ItemEquipment {
	this := ItemEquipment{}
	return &this
}

// GetAttackStab returns the AttackStab field value
func (o *ItemEquipment) GetAttackStab() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AttackStab
}

// GetAttackStabOk returns a tuple with the AttackStab field value
// and a boolean to check if the value has been set.
func (o *ItemEquipment) GetAttackStabOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttackStab, true
}

// SetAttackStab sets field value
func (o *ItemEquipment) SetAttackStab(v int32) {
	o.AttackStab = v
}

// GetAttackSlash returns the AttackSlash field value
func (o *ItemEquipment) GetAttackSlash() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AttackSlash
}

// GetAttackSlashOk returns a tuple with the AttackSlash field value
// and a boolean to check if the value has been set.
func (o *ItemEquipment) GetAttackSlashOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttackSlash, true
}

// SetAttackSlash sets field value
func (o *ItemEquipment) SetAttackSlash(v int32) {
	o.AttackSlash = v
}

// GetAttackCrush returns the AttackCrush field value
func (o *ItemEquipment) GetAttackCrush() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AttackCrush
}

// GetAttackCrushOk returns a tuple with the AttackCrush field value
// and a boolean to check if the value has been set.
func (o *ItemEquipment) GetAttackCrushOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttackCrush, true
}

// SetAttackCrush sets field value
func (o *ItemEquipment) SetAttackCrush(v int32) {
	o.AttackCrush = v
}

// GetAttackMagic returns the AttackMagic field value
func (o *ItemEquipment) GetAttackMagic() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AttackMagic
}

// GetAttackMagicOk returns a tuple with the AttackMagic field value
// and a boolean to check if the value has been set.
func (o *ItemEquipment) GetAttackMagicOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttackMagic, true
}

// SetAttackMagic sets field value
func (o *ItemEquipment) SetAttackMagic(v int32) {
	o.AttackMagic = v
}

// GetAttackRanged returns the AttackRanged field value
func (o *ItemEquipment) GetAttackRanged() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AttackRanged
}

// GetAttackRangedOk returns a tuple with the AttackRanged field value
// and a boolean to check if the value has been set.
func (o *ItemEquipment) GetAttackRangedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttackRanged, true
}

// SetAttackRanged sets field value
func (o *ItemEquipment) SetAttackRanged(v int32) {
	o.AttackRanged = v
}

// GetDefenceStab returns the DefenceStab field value
func (o *ItemEquipment) GetDefenceStab() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DefenceStab
}

// GetDefenceStabOk returns a tuple with the DefenceStab field value
// and a boolean to check if the value has been set.
func (o *ItemEquipment) GetDefenceStabOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefenceStab, true
}

// SetDefenceStab sets field value
func (o *ItemEquipment) SetDefenceStab(v int32) {
	o.DefenceStab = v
}

// GetDefenceSlash returns the DefenceSlash field value
func (o *ItemEquipment) GetDefenceSlash() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DefenceSlash
}

// GetDefenceSlashOk returns a tuple with the DefenceSlash field value
// and a boolean to check if the value has been set.
func (o *ItemEquipment) GetDefenceSlashOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefenceSlash, true
}

// SetDefenceSlash sets field value
func (o *ItemEquipment) SetDefenceSlash(v int32) {
	o.DefenceSlash = v
}

// GetDefenceCrush returns the DefenceCrush field value
func (o *ItemEquipment) GetDefenceCrush() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DefenceCrush
}

// GetDefenceCrushOk returns a tuple with the DefenceCrush field value
// and a boolean to check if the value has been set.
func (o *ItemEquipment) GetDefenceCrushOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefenceCrush, true
}

// SetDefenceCrush sets field value
func (o *ItemEquipment) SetDefenceCrush(v int32) {
	o.DefenceCrush = v
}

// GetDefenceMagic returns the DefenceMagic field value
func (o *ItemEquipment) GetDefenceMagic() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DefenceMagic
}

// GetDefenceMagicOk returns a tuple with the DefenceMagic field value
// and a boolean to check if the value has been set.
func (o *ItemEquipment) GetDefenceMagicOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefenceMagic, true
}

// SetDefenceMagic sets field value
func (o *ItemEquipment) SetDefenceMagic(v int32) {
	o.DefenceMagic = v
}

// GetDefenceRanged returns the DefenceRanged field value
func (o *ItemEquipment) GetDefenceRanged() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DefenceRanged
}

// GetDefenceRangedOk returns a tuple with the DefenceRanged field value
// and a boolean to check if the value has been set.
func (o *ItemEquipment) GetDefenceRangedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefenceRanged, true
}

// SetDefenceRanged sets field value
func (o *ItemEquipment) SetDefenceRanged(v int32) {
	o.DefenceRanged = v
}

// GetMeleeStrength returns the MeleeStrength field value
func (o *ItemEquipment) GetMeleeStrength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MeleeStrength
}

// GetMeleeStrengthOk returns a tuple with the MeleeStrength field value
// and a boolean to check if the value has been set.
func (o *ItemEquipment) GetMeleeStrengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MeleeStrength, true
}

// SetMeleeStrength sets field value
func (o *ItemEquipment) SetMeleeStrength(v int32) {
	o.MeleeStrength = v
}

// GetRangedStrength returns the RangedStrength field value
func (o *ItemEquipment) GetRangedStrength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RangedStrength
}

// GetRangedStrengthOk returns a tuple with the RangedStrength field value
// and a boolean to check if the value has been set.
func (o *ItemEquipment) GetRangedStrengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RangedStrength, true
}

// SetRangedStrength sets field value
func (o *ItemEquipment) SetRangedStrength(v int32) {
	o.RangedStrength = v
}

// GetMagicDamage returns the MagicDamage field value
func (o *ItemEquipment) GetMagicDamage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MagicDamage
}

// GetMagicDamageOk returns a tuple with the MagicDamage field value
// and a boolean to check if the value has been set.
func (o *ItemEquipment) GetMagicDamageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MagicDamage, true
}

// SetMagicDamage sets field value
func (o *ItemEquipment) SetMagicDamage(v int32) {
	o.MagicDamage = v
}

// GetPrayer returns the Prayer field value
func (o *ItemEquipment) GetPrayer() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Prayer
}

// GetPrayerOk returns a tuple with the Prayer field value
// and a boolean to check if the value has been set.
func (o *ItemEquipment) GetPrayerOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prayer, true
}

// SetPrayer sets field value
func (o *ItemEquipment) SetPrayer(v int32) {
	o.Prayer = v
}

// GetSlot returns the Slot field value
func (o *ItemEquipment) GetSlot() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slot
}

// GetSlotOk returns a tuple with the Slot field value
// and a boolean to check if the value has been set.
func (o *ItemEquipment) GetSlotOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slot, true
}

// SetSlot sets field value
func (o *ItemEquipment) SetSlot(v string) {
	o.Slot = v
}

// GetRequirements returns the Requirements field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *ItemEquipment) GetRequirements() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Requirements
}

// GetRequirementsOk returns a tuple with the Requirements field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemEquipment) GetRequirementsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Requirements == nil {
		return nil, false
	}
	return &o.Requirements, true
}

// SetRequirements sets field value
func (o *ItemEquipment) SetRequirements(v map[string]interface{}) {
	o.Requirements = v
}

func (o ItemEquipment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["attack_stab"] = o.AttackStab
	}
	if true {
		toSerialize["attack_slash"] = o.AttackSlash
	}
	if true {
		toSerialize["attack_crush"] = o.AttackCrush
	}
	if true {
		toSerialize["attack_magic"] = o.AttackMagic
	}
	if true {
		toSerialize["attack_ranged"] = o.AttackRanged
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
		toSerialize["melee_strength"] = o.MeleeStrength
	}
	if true {
		toSerialize["ranged_strength"] = o.RangedStrength
	}
	if true {
		toSerialize["magic_damage"] = o.MagicDamage
	}
	if true {
		toSerialize["prayer"] = o.Prayer
	}
	if true {
		toSerialize["slot"] = o.Slot
	}
	if o.Requirements != nil {
		toSerialize["requirements"] = o.Requirements
	}
	return json.Marshal(toSerialize)
}

type NullableItemEquipment struct {
	value *ItemEquipment
	isSet bool
}

func (v NullableItemEquipment) Get() *ItemEquipment {
	return v.value
}

func (v *NullableItemEquipment) Set(val *ItemEquipment) {
	v.value = val
	v.isSet = true
}

func (v NullableItemEquipment) IsSet() bool {
	return v.isSet
}

func (v *NullableItemEquipment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemEquipment(val *ItemEquipment) *NullableItemEquipment {
	return &NullableItemEquipment{value: val, isSet: true}
}

func (v NullableItemEquipment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemEquipment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
