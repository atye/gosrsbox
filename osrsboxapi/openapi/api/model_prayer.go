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

// Prayer struct for Prayer
type Prayer struct {
	// Unique prayer ID number.
	Id string `json:"id"`
	// The name of the prayer.
	Name string `json:"name"`
	// If the prayer is members-only.
	Members bool `json:"members"`
	// The prayer description (as show in-game).
	Description string `json:"description"`
	// The prayer point drain rate per minute.
	DrainPerMinute float32 `json:"drain_per_minute"`
	// The OSRS Wiki URL.
	WikiUrl string `json:"wiki_url"`
	// The stat requirements to use the prayer.
	Requirements map[string]interface{} `json:"requirements"`
	// The bonuses a prayer provides.
	Bonuses map[string]interface{} `json:"bonuses"`
	// The prayer icon.
	Icon string `json:"icon"`
}

// NewPrayer instantiates a new Prayer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrayer(id string, name string, members bool, description string, drainPerMinute float32, wikiUrl string, requirements map[string]interface{}, bonuses map[string]interface{}, icon string, ) *Prayer {
	this := Prayer{}
	this.Id = id
	this.Name = name
	this.Members = members
	this.Description = description
	this.DrainPerMinute = drainPerMinute
	this.WikiUrl = wikiUrl
	this.Requirements = requirements
	this.Bonuses = bonuses
	this.Icon = icon
	return &this
}

// NewPrayerWithDefaults instantiates a new Prayer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrayerWithDefaults() *Prayer {
	this := Prayer{}
	return &this
}

// GetId returns the Id field value
func (o *Prayer) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Prayer) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Prayer) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Prayer) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Prayer) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Prayer) SetName(v string) {
	o.Name = v
}

// GetMembers returns the Members field value
func (o *Prayer) GetMembers() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.Members
}

// GetMembersOk returns a tuple with the Members field value
// and a boolean to check if the value has been set.
func (o *Prayer) GetMembersOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Members, true
}

// SetMembers sets field value
func (o *Prayer) SetMembers(v bool) {
	o.Members = v
}

// GetDescription returns the Description field value
func (o *Prayer) GetDescription() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Prayer) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Prayer) SetDescription(v string) {
	o.Description = v
}

// GetDrainPerMinute returns the DrainPerMinute field value
func (o *Prayer) GetDrainPerMinute() float32 {
	if o == nil  {
		var ret float32
		return ret
	}

	return o.DrainPerMinute
}

// GetDrainPerMinuteOk returns a tuple with the DrainPerMinute field value
// and a boolean to check if the value has been set.
func (o *Prayer) GetDrainPerMinuteOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DrainPerMinute, true
}

// SetDrainPerMinute sets field value
func (o *Prayer) SetDrainPerMinute(v float32) {
	o.DrainPerMinute = v
}

// GetWikiUrl returns the WikiUrl field value
func (o *Prayer) GetWikiUrl() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.WikiUrl
}

// GetWikiUrlOk returns a tuple with the WikiUrl field value
// and a boolean to check if the value has been set.
func (o *Prayer) GetWikiUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WikiUrl, true
}

// SetWikiUrl sets field value
func (o *Prayer) SetWikiUrl(v string) {
	o.WikiUrl = v
}

// GetRequirements returns the Requirements field value
func (o *Prayer) GetRequirements() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}

	return o.Requirements
}

// GetRequirementsOk returns a tuple with the Requirements field value
// and a boolean to check if the value has been set.
func (o *Prayer) GetRequirementsOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Requirements, true
}

// SetRequirements sets field value
func (o *Prayer) SetRequirements(v map[string]interface{}) {
	o.Requirements = v
}

// GetBonuses returns the Bonuses field value
func (o *Prayer) GetBonuses() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}

	return o.Bonuses
}

// GetBonusesOk returns a tuple with the Bonuses field value
// and a boolean to check if the value has been set.
func (o *Prayer) GetBonusesOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Bonuses, true
}

// SetBonuses sets field value
func (o *Prayer) SetBonuses(v map[string]interface{}) {
	o.Bonuses = v
}

// GetIcon returns the Icon field value
func (o *Prayer) GetIcon() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Icon
}

// GetIconOk returns a tuple with the Icon field value
// and a boolean to check if the value has been set.
func (o *Prayer) GetIconOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Icon, true
}

// SetIcon sets field value
func (o *Prayer) SetIcon(v string) {
	o.Icon = v
}

func (o Prayer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["members"] = o.Members
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["drain_per_minute"] = o.DrainPerMinute
	}
	if true {
		toSerialize["wiki_url"] = o.WikiUrl
	}
	if true {
		toSerialize["requirements"] = o.Requirements
	}
	if true {
		toSerialize["bonuses"] = o.Bonuses
	}
	if true {
		toSerialize["icon"] = o.Icon
	}
	return json.Marshal(toSerialize)
}

type NullablePrayer struct {
	value *Prayer
	isSet bool
}

func (v NullablePrayer) Get() *Prayer {
	return v.value
}

func (v *NullablePrayer) Set(val *Prayer) {
	v.value = val
	v.isSet = true
}

func (v NullablePrayer) IsSet() bool {
	return v.isSet
}

func (v *NullablePrayer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrayer(val *Prayer) *NullablePrayer {
	return &NullablePrayer{value: val, isSet: true}
}

func (v NullablePrayer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrayer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


