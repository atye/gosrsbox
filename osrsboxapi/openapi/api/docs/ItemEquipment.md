# ItemEquipment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttackStab** | **int32** | The attack stab bonus of the item. | 
**AttackSlash** | **int32** | The attack slash bonus of the item. | 
**AttackCrush** | **int32** | The attack crush bonus of the item. | 
**AttackMagic** | **int32** | The attack magic bonus of the item. | 
**AttackRanged** | **int32** | The attack ranged bonus of the item. | 
**DefenceStab** | **int32** | The defence stab bonus of the item. | 
**DefenceSlash** | **int32** | The defence slash bonus of the item. | 
**DefenceCrush** | **int32** | The defence crush bonus of the item. | 
**DefenceMagic** | **int32** | The defence magic bonus of the item. | 
**DefenceRanged** | **int32** | The defence ranged bonus of the item. | 
**MeleeStrength** | **int32** | The melee strength bonus of the item. | 
**RangedStrength** | **int32** | The ranged strength bonus of the item. | 
**MagicDamage** | **int32** | The magic damage bonus of the item. | 
**Prayer** | **int32** | The prayer bonus of the item. | 
**Slot** | **string** | The equipment slot associated with the item (e.g., head). | 
**Requirements** | **map[string]interface{}** | An object of requirements {skill: level}. | 

## Methods

### NewItemEquipment

`func NewItemEquipment(attackStab int32, attackSlash int32, attackCrush int32, attackMagic int32, attackRanged int32, defenceStab int32, defenceSlash int32, defenceCrush int32, defenceMagic int32, defenceRanged int32, meleeStrength int32, rangedStrength int32, magicDamage int32, prayer int32, slot string, requirements map[string]interface{}, ) *ItemEquipment`

NewItemEquipment instantiates a new ItemEquipment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemEquipmentWithDefaults

`func NewItemEquipmentWithDefaults() *ItemEquipment`

NewItemEquipmentWithDefaults instantiates a new ItemEquipment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttackStab

`func (o *ItemEquipment) GetAttackStab() int32`

GetAttackStab returns the AttackStab field if non-nil, zero value otherwise.

### GetAttackStabOk

`func (o *ItemEquipment) GetAttackStabOk() (*int32, bool)`

GetAttackStabOk returns a tuple with the AttackStab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackStab

`func (o *ItemEquipment) SetAttackStab(v int32)`

SetAttackStab sets AttackStab field to given value.


### GetAttackSlash

`func (o *ItemEquipment) GetAttackSlash() int32`

GetAttackSlash returns the AttackSlash field if non-nil, zero value otherwise.

### GetAttackSlashOk

`func (o *ItemEquipment) GetAttackSlashOk() (*int32, bool)`

GetAttackSlashOk returns a tuple with the AttackSlash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackSlash

`func (o *ItemEquipment) SetAttackSlash(v int32)`

SetAttackSlash sets AttackSlash field to given value.


### GetAttackCrush

`func (o *ItemEquipment) GetAttackCrush() int32`

GetAttackCrush returns the AttackCrush field if non-nil, zero value otherwise.

### GetAttackCrushOk

`func (o *ItemEquipment) GetAttackCrushOk() (*int32, bool)`

GetAttackCrushOk returns a tuple with the AttackCrush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackCrush

`func (o *ItemEquipment) SetAttackCrush(v int32)`

SetAttackCrush sets AttackCrush field to given value.


### GetAttackMagic

`func (o *ItemEquipment) GetAttackMagic() int32`

GetAttackMagic returns the AttackMagic field if non-nil, zero value otherwise.

### GetAttackMagicOk

`func (o *ItemEquipment) GetAttackMagicOk() (*int32, bool)`

GetAttackMagicOk returns a tuple with the AttackMagic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackMagic

`func (o *ItemEquipment) SetAttackMagic(v int32)`

SetAttackMagic sets AttackMagic field to given value.


### GetAttackRanged

`func (o *ItemEquipment) GetAttackRanged() int32`

GetAttackRanged returns the AttackRanged field if non-nil, zero value otherwise.

### GetAttackRangedOk

`func (o *ItemEquipment) GetAttackRangedOk() (*int32, bool)`

GetAttackRangedOk returns a tuple with the AttackRanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackRanged

`func (o *ItemEquipment) SetAttackRanged(v int32)`

SetAttackRanged sets AttackRanged field to given value.


### GetDefenceStab

`func (o *ItemEquipment) GetDefenceStab() int32`

GetDefenceStab returns the DefenceStab field if non-nil, zero value otherwise.

### GetDefenceStabOk

`func (o *ItemEquipment) GetDefenceStabOk() (*int32, bool)`

GetDefenceStabOk returns a tuple with the DefenceStab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefenceStab

`func (o *ItemEquipment) SetDefenceStab(v int32)`

SetDefenceStab sets DefenceStab field to given value.


### GetDefenceSlash

`func (o *ItemEquipment) GetDefenceSlash() int32`

GetDefenceSlash returns the DefenceSlash field if non-nil, zero value otherwise.

### GetDefenceSlashOk

`func (o *ItemEquipment) GetDefenceSlashOk() (*int32, bool)`

GetDefenceSlashOk returns a tuple with the DefenceSlash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefenceSlash

`func (o *ItemEquipment) SetDefenceSlash(v int32)`

SetDefenceSlash sets DefenceSlash field to given value.


### GetDefenceCrush

`func (o *ItemEquipment) GetDefenceCrush() int32`

GetDefenceCrush returns the DefenceCrush field if non-nil, zero value otherwise.

### GetDefenceCrushOk

`func (o *ItemEquipment) GetDefenceCrushOk() (*int32, bool)`

GetDefenceCrushOk returns a tuple with the DefenceCrush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefenceCrush

`func (o *ItemEquipment) SetDefenceCrush(v int32)`

SetDefenceCrush sets DefenceCrush field to given value.


### GetDefenceMagic

`func (o *ItemEquipment) GetDefenceMagic() int32`

GetDefenceMagic returns the DefenceMagic field if non-nil, zero value otherwise.

### GetDefenceMagicOk

`func (o *ItemEquipment) GetDefenceMagicOk() (*int32, bool)`

GetDefenceMagicOk returns a tuple with the DefenceMagic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefenceMagic

`func (o *ItemEquipment) SetDefenceMagic(v int32)`

SetDefenceMagic sets DefenceMagic field to given value.


### GetDefenceRanged

`func (o *ItemEquipment) GetDefenceRanged() int32`

GetDefenceRanged returns the DefenceRanged field if non-nil, zero value otherwise.

### GetDefenceRangedOk

`func (o *ItemEquipment) GetDefenceRangedOk() (*int32, bool)`

GetDefenceRangedOk returns a tuple with the DefenceRanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefenceRanged

`func (o *ItemEquipment) SetDefenceRanged(v int32)`

SetDefenceRanged sets DefenceRanged field to given value.


### GetMeleeStrength

`func (o *ItemEquipment) GetMeleeStrength() int32`

GetMeleeStrength returns the MeleeStrength field if non-nil, zero value otherwise.

### GetMeleeStrengthOk

`func (o *ItemEquipment) GetMeleeStrengthOk() (*int32, bool)`

GetMeleeStrengthOk returns a tuple with the MeleeStrength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeleeStrength

`func (o *ItemEquipment) SetMeleeStrength(v int32)`

SetMeleeStrength sets MeleeStrength field to given value.


### GetRangedStrength

`func (o *ItemEquipment) GetRangedStrength() int32`

GetRangedStrength returns the RangedStrength field if non-nil, zero value otherwise.

### GetRangedStrengthOk

`func (o *ItemEquipment) GetRangedStrengthOk() (*int32, bool)`

GetRangedStrengthOk returns a tuple with the RangedStrength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangedStrength

`func (o *ItemEquipment) SetRangedStrength(v int32)`

SetRangedStrength sets RangedStrength field to given value.


### GetMagicDamage

`func (o *ItemEquipment) GetMagicDamage() int32`

GetMagicDamage returns the MagicDamage field if non-nil, zero value otherwise.

### GetMagicDamageOk

`func (o *ItemEquipment) GetMagicDamageOk() (*int32, bool)`

GetMagicDamageOk returns a tuple with the MagicDamage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagicDamage

`func (o *ItemEquipment) SetMagicDamage(v int32)`

SetMagicDamage sets MagicDamage field to given value.


### GetPrayer

`func (o *ItemEquipment) GetPrayer() int32`

GetPrayer returns the Prayer field if non-nil, zero value otherwise.

### GetPrayerOk

`func (o *ItemEquipment) GetPrayerOk() (*int32, bool)`

GetPrayerOk returns a tuple with the Prayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrayer

`func (o *ItemEquipment) SetPrayer(v int32)`

SetPrayer sets Prayer field to given value.


### GetSlot

`func (o *ItemEquipment) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *ItemEquipment) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *ItemEquipment) SetSlot(v string)`

SetSlot sets Slot field to given value.


### GetRequirements

`func (o *ItemEquipment) GetRequirements() map[string]interface{}`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *ItemEquipment) GetRequirementsOk() (*map[string]interface{}, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *ItemEquipment) SetRequirements(v map[string]interface{})`

SetRequirements sets Requirements field to given value.


### SetRequirementsNil

`func (o *ItemEquipment) SetRequirementsNil(b bool)`

 SetRequirementsNil sets the value for Requirements to be an explicit nil

### UnsetRequirements
`func (o *ItemEquipment) UnsetRequirements()`

UnsetRequirements ensures that no value is present for Requirements, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


