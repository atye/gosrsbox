# Prayer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique prayer ID number. | 
**Name** | **string** | The name of the prayer. | 
**Members** | **bool** | If the prayer is members-only. | 
**Description** | **string** | The prayer description (as show in-game). | 
**DrainPerMinute** | **float32** | The prayer point drain rate per minute. | 
**WikiUrl** | **string** | The OSRS Wiki URL. | 
**Requirements** | **map[string]interface{}** | The stat requirements to use the prayer. | 
**Bonuses** | **map[string]interface{}** | The bonuses a prayer provides. | 
**Icon** | **string** | The prayer icon. | 

## Methods

### NewPrayer

`func NewPrayer(id string, name string, members bool, description string, drainPerMinute float32, wikiUrl string, requirements map[string]interface{}, bonuses map[string]interface{}, icon string, ) *Prayer`

NewPrayer instantiates a new Prayer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrayerWithDefaults

`func NewPrayerWithDefaults() *Prayer`

NewPrayerWithDefaults instantiates a new Prayer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Prayer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Prayer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Prayer) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Prayer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Prayer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Prayer) SetName(v string)`

SetName sets Name field to given value.


### GetMembers

`func (o *Prayer) GetMembers() bool`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Prayer) GetMembersOk() (*bool, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *Prayer) SetMembers(v bool)`

SetMembers sets Members field to given value.


### GetDescription

`func (o *Prayer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Prayer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Prayer) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDrainPerMinute

`func (o *Prayer) GetDrainPerMinute() float32`

GetDrainPerMinute returns the DrainPerMinute field if non-nil, zero value otherwise.

### GetDrainPerMinuteOk

`func (o *Prayer) GetDrainPerMinuteOk() (*float32, bool)`

GetDrainPerMinuteOk returns a tuple with the DrainPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainPerMinute

`func (o *Prayer) SetDrainPerMinute(v float32)`

SetDrainPerMinute sets DrainPerMinute field to given value.


### GetWikiUrl

`func (o *Prayer) GetWikiUrl() string`

GetWikiUrl returns the WikiUrl field if non-nil, zero value otherwise.

### GetWikiUrlOk

`func (o *Prayer) GetWikiUrlOk() (*string, bool)`

GetWikiUrlOk returns a tuple with the WikiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWikiUrl

`func (o *Prayer) SetWikiUrl(v string)`

SetWikiUrl sets WikiUrl field to given value.


### GetRequirements

`func (o *Prayer) GetRequirements() map[string]interface{}`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *Prayer) GetRequirementsOk() (*map[string]interface{}, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *Prayer) SetRequirements(v map[string]interface{})`

SetRequirements sets Requirements field to given value.


### GetBonuses

`func (o *Prayer) GetBonuses() map[string]interface{}`

GetBonuses returns the Bonuses field if non-nil, zero value otherwise.

### GetBonusesOk

`func (o *Prayer) GetBonusesOk() (*map[string]interface{}, bool)`

GetBonusesOk returns a tuple with the Bonuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonuses

`func (o *Prayer) SetBonuses(v map[string]interface{})`

SetBonuses sets Bonuses field to given value.


### GetIcon

`func (o *Prayer) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *Prayer) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *Prayer) SetIcon(v string)`

SetIcon sets Icon field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


