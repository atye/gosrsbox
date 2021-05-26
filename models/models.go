package models

import (
	"github.com/atye/gosrsbox/internal/api"
)

type Item = api.Item

type ItemEquipment = api.ItemEquipment

type ItemWeapon = api.ItemWeapon

type ItemWeaponStances = api.ItemWeaponStances

type Monster = api.Monster

type MonsterDrops = api.MonsterDrops

type Prayer = api.Prayer

var (
	NewItem                      = api.NewItem
	NewNullableItem              = api.NewNullableItem
	NewItemEquipment             = api.NewItemEquipment
	NewNullableItemEquipment     = api.NewNullableItemEquipment
	NewItemWeapon                = api.NewItemWeapon
	NewNullableItemWeapon        = api.NewNullableItemWeapon
	NewItemWeaponStances         = api.NewItemWeaponStances
	NewNullableItemWeaponStances = api.NewNullableItemWeaponStances

	NewMonster              = api.NewMonster
	NewNullableMonster      = api.NewNullableMonster
	NewMonsterDrops         = api.NewMonsterDrops
	NewNullableMonsterDrops = api.NewNullableMonsterDrops

	NewPrayer        = api.NewPrayer
	NewNullablePraye = api.NewNullablePrayer

	PtrBool            = api.PtrBool
	PtrInt             = api.PtrInt
	PtrInt32           = api.PtrInt32
	PtrInt64           = api.PtrInt64
	PtrFloat32         = api.PtrFloat32
	PtrFloat64         = api.PtrFloat64
	PtrString          = api.PtrString
	PtrTime            = api.PtrTime
	NewNullableBool    = api.NewNullableBool
	NewNullableInt     = api.NewNullableInt
	NewNullableInt32   = api.NewNullableInt32
	NewNullableInt64   = api.NewNullableInt64
	NewNullableFloat32 = api.NewNullableFloat32
	NewNullableFloat64 = api.NewNullableFloat64
	NewNullableString  = api.NewNullableString
	NewNullableTime    = api.NewNullableTime
)
