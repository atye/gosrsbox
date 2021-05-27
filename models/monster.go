package models

import "github.com/atye/gosrsbox/internal/api"

type Monster = api.Monster

type NullableMonster = api.NullableMonster

type MonsterDrops = api.MonsterDrops

type NullableMonsterDrops = api.NullableMonsterDrops

func NewNullableMonster(val *Monster) NullableMonster {
	return *api.NewNullableMonster(val)
}

func NewNullableMonsterDrops(val *MonsterDrops) NullableMonsterDrops {
	return *api.NewNullableMonsterDrops(val)
}
