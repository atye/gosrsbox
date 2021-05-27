package models

import "github.com/atye/gosrsbox/internal/api"

type Prayer = api.Prayer

type NullablePrayer = api.NullablePrayer

func NewNullablePrayer(val *Prayer) NullablePrayer {
	return *api.NewNullablePrayer(val)
}
