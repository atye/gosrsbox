/*
 * osrsbox-api
 *
 * An open, free, complete and up-to-date RESTful API for Old School RuneScape (OSRS) items, monsters and prayers.
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ResponeMetadata struct for ResponeMetadata
type ResponeMetadata struct {
	Page *string `json:"page,omitempty"`
	Total *int32 `json:"total,omitempty"`
	MaxResults *int32 `json:"max_results,omitempty"`
}

// NewResponeMetadata instantiates a new ResponeMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponeMetadata() *ResponeMetadata {
	this := ResponeMetadata{}
	return &this
}

// NewResponeMetadataWithDefaults instantiates a new ResponeMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponeMetadataWithDefaults() *ResponeMetadata {
	this := ResponeMetadata{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *ResponeMetadata) GetPage() string {
	if o == nil || o.Page == nil {
		var ret string
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponeMetadata) GetPageOk() (*string, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *ResponeMetadata) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given string and assigns it to the Page field.
func (o *ResponeMetadata) SetPage(v string) {
	o.Page = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ResponeMetadata) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponeMetadata) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ResponeMetadata) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *ResponeMetadata) SetTotal(v int32) {
	o.Total = &v
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *ResponeMetadata) GetMaxResults() int32 {
	if o == nil || o.MaxResults == nil {
		var ret int32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponeMetadata) GetMaxResultsOk() (*int32, bool) {
	if o == nil || o.MaxResults == nil {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *ResponeMetadata) HasMaxResults() bool {
	if o != nil && o.MaxResults != nil {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.
func (o *ResponeMetadata) SetMaxResults(v int32) {
	o.MaxResults = &v
}

func (o ResponeMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.MaxResults != nil {
		toSerialize["max_results"] = o.MaxResults
	}
	return json.Marshal(toSerialize)
}

type NullableResponeMetadata struct {
	value *ResponeMetadata
	isSet bool
}

func (v NullableResponeMetadata) Get() *ResponeMetadata {
	return v.value
}

func (v *NullableResponeMetadata) Set(val *ResponeMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableResponeMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableResponeMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponeMetadata(val *ResponeMetadata) *NullableResponeMetadata {
	return &NullableResponeMetadata{value: val, isSet: true}
}

func (v NullableResponeMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponeMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


