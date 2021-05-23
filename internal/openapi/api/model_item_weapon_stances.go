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

// ItemWeaponStances struct for ItemWeaponStances
type ItemWeaponStances struct {
	CombatStyle string `json:"combat_style"`
	AttackType NullableString `json:"attack_type"`
	AttackStyle NullableString `json:"attack_style"`
	Experience NullableString `json:"experience"`
	Boosts NullableString `json:"boosts"`
}

// NewItemWeaponStances instantiates a new ItemWeaponStances object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemWeaponStances(combatStyle string, attackType NullableString, attackStyle NullableString, experience NullableString, boosts NullableString, ) *ItemWeaponStances {
	this := ItemWeaponStances{}
	this.CombatStyle = combatStyle
	this.AttackType = attackType
	this.AttackStyle = attackStyle
	this.Experience = experience
	this.Boosts = boosts
	return &this
}

// NewItemWeaponStancesWithDefaults instantiates a new ItemWeaponStances object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemWeaponStancesWithDefaults() *ItemWeaponStances {
	this := ItemWeaponStances{}
	return &this
}

// GetCombatStyle returns the CombatStyle field value
func (o *ItemWeaponStances) GetCombatStyle() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.CombatStyle
}

// GetCombatStyleOk returns a tuple with the CombatStyle field value
// and a boolean to check if the value has been set.
func (o *ItemWeaponStances) GetCombatStyleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CombatStyle, true
}

// SetCombatStyle sets field value
func (o *ItemWeaponStances) SetCombatStyle(v string) {
	o.CombatStyle = v
}

// GetAttackType returns the AttackType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ItemWeaponStances) GetAttackType() string {
	if o == nil || o.AttackType.Get() == nil {
		var ret string
		return ret
	}

	return *o.AttackType.Get()
}

// GetAttackTypeOk returns a tuple with the AttackType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemWeaponStances) GetAttackTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AttackType.Get(), o.AttackType.IsSet()
}

// SetAttackType sets field value
func (o *ItemWeaponStances) SetAttackType(v string) {
	o.AttackType.Set(&v)
}

// GetAttackStyle returns the AttackStyle field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ItemWeaponStances) GetAttackStyle() string {
	if o == nil || o.AttackStyle.Get() == nil {
		var ret string
		return ret
	}

	return *o.AttackStyle.Get()
}

// GetAttackStyleOk returns a tuple with the AttackStyle field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemWeaponStances) GetAttackStyleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AttackStyle.Get(), o.AttackStyle.IsSet()
}

// SetAttackStyle sets field value
func (o *ItemWeaponStances) SetAttackStyle(v string) {
	o.AttackStyle.Set(&v)
}

// GetExperience returns the Experience field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ItemWeaponStances) GetExperience() string {
	if o == nil || o.Experience.Get() == nil {
		var ret string
		return ret
	}

	return *o.Experience.Get()
}

// GetExperienceOk returns a tuple with the Experience field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemWeaponStances) GetExperienceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Experience.Get(), o.Experience.IsSet()
}

// SetExperience sets field value
func (o *ItemWeaponStances) SetExperience(v string) {
	o.Experience.Set(&v)
}

// GetBoosts returns the Boosts field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ItemWeaponStances) GetBoosts() string {
	if o == nil || o.Boosts.Get() == nil {
		var ret string
		return ret
	}

	return *o.Boosts.Get()
}

// GetBoostsOk returns a tuple with the Boosts field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemWeaponStances) GetBoostsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Boosts.Get(), o.Boosts.IsSet()
}

// SetBoosts sets field value
func (o *ItemWeaponStances) SetBoosts(v string) {
	o.Boosts.Set(&v)
}

func (o ItemWeaponStances) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["combat_style"] = o.CombatStyle
	}
	if true {
		toSerialize["attack_type"] = o.AttackType.Get()
	}
	if true {
		toSerialize["attack_style"] = o.AttackStyle.Get()
	}
	if true {
		toSerialize["experience"] = o.Experience.Get()
	}
	if true {
		toSerialize["boosts"] = o.Boosts.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableItemWeaponStances struct {
	value *ItemWeaponStances
	isSet bool
}

func (v NullableItemWeaponStances) Get() *ItemWeaponStances {
	return v.value
}

func (v *NullableItemWeaponStances) Set(val *ItemWeaponStances) {
	v.value = val
	v.isSet = true
}

func (v NullableItemWeaponStances) IsSet() bool {
	return v.isSet
}

func (v *NullableItemWeaponStances) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemWeaponStances(val *ItemWeaponStances) *NullableItemWeaponStances {
	return &NullableItemWeaponStances{value: val, isSet: true}
}

func (v NullableItemWeaponStances) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemWeaponStances) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

