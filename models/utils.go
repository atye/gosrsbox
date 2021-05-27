package models

import (
	"time"

	"github.com/atye/gosrsbox/internal/api"
)

type NullableBool = api.NullableBool

type NullableInt = api.NullableInt

type NullableInt32 = api.NullableInt32

type NullableInt64 = api.NullableInt64

type NullableFloat32 = api.NullableFloat32

type NullableFloat64 = api.NullableFloat64

type NullableString = api.NullableString

type NullableTime = api.NullableTime

func NewNullableBool(val bool) NullableBool {
	return *api.NewNullableBool(api.PtrBool(val))
}

func NewNullableInt(val int) NullableInt {
	return *api.NewNullableInt(api.PtrInt(val))
}

func NewNullableInt32(val int32) NullableInt32 {
	return *api.NewNullableInt32(api.PtrInt32(val))
}

func NewNullableInt64(val int64) NullableInt64 {
	return *api.NewNullableInt64(api.PtrInt64(val))
}

func NewNullableFloat32(val float32) NullableFloat32 {
	return *api.NewNullableFloat32(api.PtrFloat32(val))
}

func NewNullableFloat64(val float64) NullableFloat64 {
	return *api.NewNullableFloat64(api.PtrFloat64(val))
}

func NewNullableString(val string) NullableString {
	return *api.NewNullableString(api.PtrString(val))
}

func NewNullableTime(val time.Time) NullableTime {
	return *api.NewNullableTime(api.PtrTime(val))
}
