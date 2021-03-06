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

// ResponeLinks struct for ResponeLinks
type ResponeLinks struct {
	Parent *ResponeLinksParent `json:"parent,omitempty"`
	Self *ResponeLinksParent `json:"self,omitempty"`
}

// NewResponeLinks instantiates a new ResponeLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponeLinks() *ResponeLinks {
	this := ResponeLinks{}
	return &this
}

// NewResponeLinksWithDefaults instantiates a new ResponeLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponeLinksWithDefaults() *ResponeLinks {
	this := ResponeLinks{}
	return &this
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *ResponeLinks) GetParent() ResponeLinksParent {
	if o == nil || o.Parent == nil {
		var ret ResponeLinksParent
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponeLinks) GetParentOk() (*ResponeLinksParent, bool) {
	if o == nil || o.Parent == nil {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *ResponeLinks) HasParent() bool {
	if o != nil && o.Parent != nil {
		return true
	}

	return false
}

// SetParent gets a reference to the given ResponeLinksParent and assigns it to the Parent field.
func (o *ResponeLinks) SetParent(v ResponeLinksParent) {
	o.Parent = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *ResponeLinks) GetSelf() ResponeLinksParent {
	if o == nil || o.Self == nil {
		var ret ResponeLinksParent
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponeLinks) GetSelfOk() (*ResponeLinksParent, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *ResponeLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given ResponeLinksParent and assigns it to the Self field.
func (o *ResponeLinks) SetSelf(v ResponeLinksParent) {
	o.Self = &v
}

func (o ResponeLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Parent != nil {
		toSerialize["parent"] = o.Parent
	}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	return json.Marshal(toSerialize)
}

type NullableResponeLinks struct {
	value *ResponeLinks
	isSet bool
}

func (v NullableResponeLinks) Get() *ResponeLinks {
	return v.value
}

func (v *NullableResponeLinks) Set(val *ResponeLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableResponeLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableResponeLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponeLinks(val *ResponeLinks) *NullableResponeLinks {
	return &NullableResponeLinks{value: val, isSet: true}
}

func (v NullableResponeLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponeLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


