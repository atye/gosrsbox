package models

import "github.com/atye/gosrsbox/internal/api"

type Item = api.Item

type NullableItem = api.NullableItem

type ItemEquipment = api.ItemEquipment

type NullableItemEquipment = api.NullableItemEquipment

type ItemWeapon = api.ItemWeapon

type NullableItemWeapon = api.NullableItemWeapon

type ItemWeaponStances = api.ItemWeaponStances

type NullableItemWeaponStances = api.NullableItemWeaponStances

func NewNullableItem(val *Item) NullableItem {
	return *api.NewNullableItem(val)
}

func NewNullableItemEquipment(val *ItemEquipment) NullableItemEquipment {
	return *api.NewNullableItemEquipment(val)
}

func NewNullableItemWeapon(val *ItemWeapon) NullableItemWeapon {
	return *api.NewNullableItemWeapon(val)
}

func NewNullableItemWeaponStances(val *ItemWeaponStances) NullableItemWeaponStances {
	return *api.NewNullableItemWeaponStances(val)
}
