/*
 * osrsbox-api
 *
 * An open, free, complete and up-to-date RESTful API for Old School RuneScape (OSRS) items, monsters and prayers.
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// IconsPrayer struct for IconsPrayer
type IconsPrayer struct {
	// Unique OSRS prayer ID number.
	Id int32 `json:"id"`
	// The icon image (in base64 encoding) of the prayer.
	Icon string `json:"icon"`
}

// NewIconsPrayer instantiates a new IconsPrayer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIconsPrayer(id int32, icon string, ) *IconsPrayer {
	this := IconsPrayer{}
	this.Id = id
	this.Icon = icon
	return &this
}

// NewIconsPrayerWithDefaults instantiates a new IconsPrayer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIconsPrayerWithDefaults() *IconsPrayer {
	this := IconsPrayer{}
	return &this
}

// GetId returns the Id field value
func (o *IconsPrayer) GetId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IconsPrayer) GetIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IconsPrayer) SetId(v int32) {
	o.Id = v
}

// GetIcon returns the Icon field value
func (o *IconsPrayer) GetIcon() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Icon
}

// GetIconOk returns a tuple with the Icon field value
// and a boolean to check if the value has been set.
func (o *IconsPrayer) GetIconOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Icon, true
}

// SetIcon sets field value
func (o *IconsPrayer) SetIcon(v string) {
	o.Icon = v
}

func (o IconsPrayer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["icon"] = o.Icon
	}
	return json.Marshal(toSerialize)
}

type NullableIconsPrayer struct {
	value *IconsPrayer
	isSet bool
}

func (v NullableIconsPrayer) Get() *IconsPrayer {
	return v.value
}

func (v *NullableIconsPrayer) Set(val *IconsPrayer) {
	v.value = val
	v.isSet = true
}

func (v NullableIconsPrayer) IsSet() bool {
	return v.isSet
}

func (v *NullableIconsPrayer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIconsPrayer(val *IconsPrayer) *NullableIconsPrayer {
	return &NullableIconsPrayer{value: val, isSet: true}
}

func (v NullableIconsPrayer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIconsPrayer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

