# ItemWeapon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttackSpeed** | **int32** | The attack speed of a weapon (in game ticks). | 
**WeaponType** | **string** | The weapon classification (e.g., axes) | 
**Stances** | [**[]ItemWeaponStances**](ItemWeaponStances.md) | An array of weapon stance information. | 

## Methods

### NewItemWeapon

`func NewItemWeapon(attackSpeed int32, weaponType string, stances []ItemWeaponStances, ) *ItemWeapon`

NewItemWeapon instantiates a new ItemWeapon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemWeaponWithDefaults

`func NewItemWeaponWithDefaults() *ItemWeapon`

NewItemWeaponWithDefaults instantiates a new ItemWeapon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttackSpeed

`func (o *ItemWeapon) GetAttackSpeed() int32`

GetAttackSpeed returns the AttackSpeed field if non-nil, zero value otherwise.

### GetAttackSpeedOk

`func (o *ItemWeapon) GetAttackSpeedOk() (*int32, bool)`

GetAttackSpeedOk returns a tuple with the AttackSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackSpeed

`func (o *ItemWeapon) SetAttackSpeed(v int32)`

SetAttackSpeed sets AttackSpeed field to given value.


### GetWeaponType

`func (o *ItemWeapon) GetWeaponType() string`

GetWeaponType returns the WeaponType field if non-nil, zero value otherwise.

### GetWeaponTypeOk

`func (o *ItemWeapon) GetWeaponTypeOk() (*string, bool)`

GetWeaponTypeOk returns a tuple with the WeaponType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeaponType

`func (o *ItemWeapon) SetWeaponType(v string)`

SetWeaponType sets WeaponType field to given value.


### GetStances

`func (o *ItemWeapon) GetStances() []ItemWeaponStances`

GetStances returns the Stances field if non-nil, zero value otherwise.

### GetStancesOk

`func (o *ItemWeapon) GetStancesOk() (*[]ItemWeaponStances, bool)`

GetStancesOk returns a tuple with the Stances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStances

`func (o *ItemWeapon) SetStances(v []ItemWeaponStances)`

SetStances sets Stances field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


