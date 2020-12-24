# Monster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique OSRS item ID number. | 
**Name** | **string** | The name of the monster. | 
**Incomplete** | **bool** | If the monster has incomplete wiki data. | 
**Members** | **bool** | If the monster is members only, or not. | 
**ReleaseDate** | **NullableString** | The release date of the monster (in ISO8601 format). | 
**CombatLevel** | **int32** | The combat level of the monster. | 
**Size** | **int32** | The size, in tiles, of the monster. | 
**Hitpoints** | **int32** | The number of hitpoints a monster has. | 
**MaxHit** | **int32** | The maximum hit of the monster. | 
**AttackType** | **[]string** | The attack style (melee, magic, range) of the monster. | 
**AttackSpeed** | **NullableInt32** | The attack speed (in game ticks) of the monster. | 
**Aggressive** | **bool** | If the monster is aggressive, or not. | 
**Poisonous** | **bool** | If the monster poisons, or not | 
**ImmunePoison** | **bool** | If the monster is immune to poison, or not | 
**ImmuneVenom** | **bool** | If the monster is immune to venom, or not | 
**Attributes** | **[]string** | An array of monster attributes. | 
**Category** | **[]string** | An array of monster category. | 
**SlayerMonster** | **bool** | If the monster is a potential slayer task. | 
**SlayerLevel** | **NullableInt32** | The slayer level required to kill the monster. | 
**SlayerXp** | **NullableFloat32** | The slayer XP rewarded for a monster kill. | 
**SlayerMasters** | **[]string** | The slayer masters who can assign the monster. | 
**Duplicate** | **bool** | If the monster is a duplicate. | 
**Examine** | **string** | The examine text of the monster. | 
**Icon** | **NullableString** | The monster icon  (in base64 encoding). | 
**WikiName** | **string** | The OSRS Wiki name for the monster. | 
**WikiUrl** | **string** | The OSRS Wiki URL (possibly including anchor link). | 
**AttackLevel** | **int32** | The attack level of the monster. | 
**StrengthLevel** | **int32** | The strength level of the monster. | 
**DefenceLevel** | **int32** | The defence level of the monster. | 
**MagicLevel** | **int32** | The magic level of the monster. | 
**RangedLevel** | **int32** | The ranged level of the monster. | 
**AttackStab** | **int32** | The attack stab bonus of the monster. | 
**AttackSlash** | **int32** | The attack slash bonus of the monster. | 
**AttackCrush** | **int32** | The attack crush bonus of the monster. | 
**AttackMagic** | **int32** | The attack magic bonus of the monster. | 
**AttackRanged** | **int32** | The attack ranged bonus of the monster. | 
**DefenceStab** | **int32** | The defence stab bonus of the monster. | 
**DefenceSlash** | **int32** | The defence slash bonus of the monster. | 
**DefenceCrush** | **int32** | The defence crush bonus of the monster. | 
**DefenceMagic** | **int32** | The defence magic bonus of the monster. | 
**DefenceRanged** | **int32** | The defence ranged bonus of the monster. | 
**AttackAccuracy** | **int32** | The attack accuracy bonus of the monster. | 
**MeleeStrength** | **int32** | The melee strength bonus of the monster. | 
**RangedStrength** | **int32** | The ranged strength bonus of the monster. | 
**MagicDamage** | **int32** | The magic damage bonus of the monster. | 
**Drops** | [**[]MonsterDrops**](MonsterDrops.md) | An array of monster drop objects. | 

## Methods

### NewMonster

`func NewMonster(id string, name string, incomplete bool, members bool, releaseDate NullableString, combatLevel int32, size int32, hitpoints int32, maxHit int32, attackType []string, attackSpeed NullableInt32, aggressive bool, poisonous bool, immunePoison bool, immuneVenom bool, attributes []string, category []string, slayerMonster bool, slayerLevel NullableInt32, slayerXp NullableFloat32, slayerMasters []string, duplicate bool, examine string, icon NullableString, wikiName string, wikiUrl string, attackLevel int32, strengthLevel int32, defenceLevel int32, magicLevel int32, rangedLevel int32, attackStab int32, attackSlash int32, attackCrush int32, attackMagic int32, attackRanged int32, defenceStab int32, defenceSlash int32, defenceCrush int32, defenceMagic int32, defenceRanged int32, attackAccuracy int32, meleeStrength int32, rangedStrength int32, magicDamage int32, drops []MonsterDrops, ) *Monster`

NewMonster instantiates a new Monster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonsterWithDefaults

`func NewMonsterWithDefaults() *Monster`

NewMonsterWithDefaults instantiates a new Monster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Monster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Monster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Monster) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Monster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Monster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Monster) SetName(v string)`

SetName sets Name field to given value.


### GetIncomplete

`func (o *Monster) GetIncomplete() bool`

GetIncomplete returns the Incomplete field if non-nil, zero value otherwise.

### GetIncompleteOk

`func (o *Monster) GetIncompleteOk() (*bool, bool)`

GetIncompleteOk returns a tuple with the Incomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomplete

`func (o *Monster) SetIncomplete(v bool)`

SetIncomplete sets Incomplete field to given value.


### GetMembers

`func (o *Monster) GetMembers() bool`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Monster) GetMembersOk() (*bool, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *Monster) SetMembers(v bool)`

SetMembers sets Members field to given value.


### GetReleaseDate

`func (o *Monster) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *Monster) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *Monster) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.


### SetReleaseDateNil

`func (o *Monster) SetReleaseDateNil(b bool)`

 SetReleaseDateNil sets the value for ReleaseDate to be an explicit nil

### UnsetReleaseDate
`func (o *Monster) UnsetReleaseDate()`

UnsetReleaseDate ensures that no value is present for ReleaseDate, not even an explicit nil
### GetCombatLevel

`func (o *Monster) GetCombatLevel() int32`

GetCombatLevel returns the CombatLevel field if non-nil, zero value otherwise.

### GetCombatLevelOk

`func (o *Monster) GetCombatLevelOk() (*int32, bool)`

GetCombatLevelOk returns a tuple with the CombatLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombatLevel

`func (o *Monster) SetCombatLevel(v int32)`

SetCombatLevel sets CombatLevel field to given value.


### GetSize

`func (o *Monster) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Monster) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Monster) SetSize(v int32)`

SetSize sets Size field to given value.


### GetHitpoints

`func (o *Monster) GetHitpoints() int32`

GetHitpoints returns the Hitpoints field if non-nil, zero value otherwise.

### GetHitpointsOk

`func (o *Monster) GetHitpointsOk() (*int32, bool)`

GetHitpointsOk returns a tuple with the Hitpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitpoints

`func (o *Monster) SetHitpoints(v int32)`

SetHitpoints sets Hitpoints field to given value.


### GetMaxHit

`func (o *Monster) GetMaxHit() int32`

GetMaxHit returns the MaxHit field if non-nil, zero value otherwise.

### GetMaxHitOk

`func (o *Monster) GetMaxHitOk() (*int32, bool)`

GetMaxHitOk returns a tuple with the MaxHit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHit

`func (o *Monster) SetMaxHit(v int32)`

SetMaxHit sets MaxHit field to given value.


### GetAttackType

`func (o *Monster) GetAttackType() []string`

GetAttackType returns the AttackType field if non-nil, zero value otherwise.

### GetAttackTypeOk

`func (o *Monster) GetAttackTypeOk() (*[]string, bool)`

GetAttackTypeOk returns a tuple with the AttackType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackType

`func (o *Monster) SetAttackType(v []string)`

SetAttackType sets AttackType field to given value.


### GetAttackSpeed

`func (o *Monster) GetAttackSpeed() int32`

GetAttackSpeed returns the AttackSpeed field if non-nil, zero value otherwise.

### GetAttackSpeedOk

`func (o *Monster) GetAttackSpeedOk() (*int32, bool)`

GetAttackSpeedOk returns a tuple with the AttackSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackSpeed

`func (o *Monster) SetAttackSpeed(v int32)`

SetAttackSpeed sets AttackSpeed field to given value.


### SetAttackSpeedNil

`func (o *Monster) SetAttackSpeedNil(b bool)`

 SetAttackSpeedNil sets the value for AttackSpeed to be an explicit nil

### UnsetAttackSpeed
`func (o *Monster) UnsetAttackSpeed()`

UnsetAttackSpeed ensures that no value is present for AttackSpeed, not even an explicit nil
### GetAggressive

`func (o *Monster) GetAggressive() bool`

GetAggressive returns the Aggressive field if non-nil, zero value otherwise.

### GetAggressiveOk

`func (o *Monster) GetAggressiveOk() (*bool, bool)`

GetAggressiveOk returns a tuple with the Aggressive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggressive

`func (o *Monster) SetAggressive(v bool)`

SetAggressive sets Aggressive field to given value.


### GetPoisonous

`func (o *Monster) GetPoisonous() bool`

GetPoisonous returns the Poisonous field if non-nil, zero value otherwise.

### GetPoisonousOk

`func (o *Monster) GetPoisonousOk() (*bool, bool)`

GetPoisonousOk returns a tuple with the Poisonous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoisonous

`func (o *Monster) SetPoisonous(v bool)`

SetPoisonous sets Poisonous field to given value.


### GetImmunePoison

`func (o *Monster) GetImmunePoison() bool`

GetImmunePoison returns the ImmunePoison field if non-nil, zero value otherwise.

### GetImmunePoisonOk

`func (o *Monster) GetImmunePoisonOk() (*bool, bool)`

GetImmunePoisonOk returns a tuple with the ImmunePoison field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmunePoison

`func (o *Monster) SetImmunePoison(v bool)`

SetImmunePoison sets ImmunePoison field to given value.


### GetImmuneVenom

`func (o *Monster) GetImmuneVenom() bool`

GetImmuneVenom returns the ImmuneVenom field if non-nil, zero value otherwise.

### GetImmuneVenomOk

`func (o *Monster) GetImmuneVenomOk() (*bool, bool)`

GetImmuneVenomOk returns a tuple with the ImmuneVenom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmuneVenom

`func (o *Monster) SetImmuneVenom(v bool)`

SetImmuneVenom sets ImmuneVenom field to given value.


### GetAttributes

`func (o *Monster) GetAttributes() []string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Monster) GetAttributesOk() (*[]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Monster) SetAttributes(v []string)`

SetAttributes sets Attributes field to given value.


### GetCategory

`func (o *Monster) GetCategory() []string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Monster) GetCategoryOk() (*[]string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Monster) SetCategory(v []string)`

SetCategory sets Category field to given value.


### GetSlayerMonster

`func (o *Monster) GetSlayerMonster() bool`

GetSlayerMonster returns the SlayerMonster field if non-nil, zero value otherwise.

### GetSlayerMonsterOk

`func (o *Monster) GetSlayerMonsterOk() (*bool, bool)`

GetSlayerMonsterOk returns a tuple with the SlayerMonster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlayerMonster

`func (o *Monster) SetSlayerMonster(v bool)`

SetSlayerMonster sets SlayerMonster field to given value.


### GetSlayerLevel

`func (o *Monster) GetSlayerLevel() int32`

GetSlayerLevel returns the SlayerLevel field if non-nil, zero value otherwise.

### GetSlayerLevelOk

`func (o *Monster) GetSlayerLevelOk() (*int32, bool)`

GetSlayerLevelOk returns a tuple with the SlayerLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlayerLevel

`func (o *Monster) SetSlayerLevel(v int32)`

SetSlayerLevel sets SlayerLevel field to given value.


### SetSlayerLevelNil

`func (o *Monster) SetSlayerLevelNil(b bool)`

 SetSlayerLevelNil sets the value for SlayerLevel to be an explicit nil

### UnsetSlayerLevel
`func (o *Monster) UnsetSlayerLevel()`

UnsetSlayerLevel ensures that no value is present for SlayerLevel, not even an explicit nil
### GetSlayerXp

`func (o *Monster) GetSlayerXp() float32`

GetSlayerXp returns the SlayerXp field if non-nil, zero value otherwise.

### GetSlayerXpOk

`func (o *Monster) GetSlayerXpOk() (*float32, bool)`

GetSlayerXpOk returns a tuple with the SlayerXp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlayerXp

`func (o *Monster) SetSlayerXp(v float32)`

SetSlayerXp sets SlayerXp field to given value.


### SetSlayerXpNil

`func (o *Monster) SetSlayerXpNil(b bool)`

 SetSlayerXpNil sets the value for SlayerXp to be an explicit nil

### UnsetSlayerXp
`func (o *Monster) UnsetSlayerXp()`

UnsetSlayerXp ensures that no value is present for SlayerXp, not even an explicit nil
### GetSlayerMasters

`func (o *Monster) GetSlayerMasters() []string`

GetSlayerMasters returns the SlayerMasters field if non-nil, zero value otherwise.

### GetSlayerMastersOk

`func (o *Monster) GetSlayerMastersOk() (*[]string, bool)`

GetSlayerMastersOk returns a tuple with the SlayerMasters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlayerMasters

`func (o *Monster) SetSlayerMasters(v []string)`

SetSlayerMasters sets SlayerMasters field to given value.


### GetDuplicate

`func (o *Monster) GetDuplicate() bool`

GetDuplicate returns the Duplicate field if non-nil, zero value otherwise.

### GetDuplicateOk

`func (o *Monster) GetDuplicateOk() (*bool, bool)`

GetDuplicateOk returns a tuple with the Duplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicate

`func (o *Monster) SetDuplicate(v bool)`

SetDuplicate sets Duplicate field to given value.


### GetExamine

`func (o *Monster) GetExamine() string`

GetExamine returns the Examine field if non-nil, zero value otherwise.

### GetExamineOk

`func (o *Monster) GetExamineOk() (*string, bool)`

GetExamineOk returns a tuple with the Examine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExamine

`func (o *Monster) SetExamine(v string)`

SetExamine sets Examine field to given value.


### GetIcon

`func (o *Monster) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *Monster) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *Monster) SetIcon(v string)`

SetIcon sets Icon field to given value.


### SetIconNil

`func (o *Monster) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *Monster) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetWikiName

`func (o *Monster) GetWikiName() string`

GetWikiName returns the WikiName field if non-nil, zero value otherwise.

### GetWikiNameOk

`func (o *Monster) GetWikiNameOk() (*string, bool)`

GetWikiNameOk returns a tuple with the WikiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWikiName

`func (o *Monster) SetWikiName(v string)`

SetWikiName sets WikiName field to given value.


### GetWikiUrl

`func (o *Monster) GetWikiUrl() string`

GetWikiUrl returns the WikiUrl field if non-nil, zero value otherwise.

### GetWikiUrlOk

`func (o *Monster) GetWikiUrlOk() (*string, bool)`

GetWikiUrlOk returns a tuple with the WikiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWikiUrl

`func (o *Monster) SetWikiUrl(v string)`

SetWikiUrl sets WikiUrl field to given value.


### GetAttackLevel

`func (o *Monster) GetAttackLevel() int32`

GetAttackLevel returns the AttackLevel field if non-nil, zero value otherwise.

### GetAttackLevelOk

`func (o *Monster) GetAttackLevelOk() (*int32, bool)`

GetAttackLevelOk returns a tuple with the AttackLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackLevel

`func (o *Monster) SetAttackLevel(v int32)`

SetAttackLevel sets AttackLevel field to given value.


### GetStrengthLevel

`func (o *Monster) GetStrengthLevel() int32`

GetStrengthLevel returns the StrengthLevel field if non-nil, zero value otherwise.

### GetStrengthLevelOk

`func (o *Monster) GetStrengthLevelOk() (*int32, bool)`

GetStrengthLevelOk returns a tuple with the StrengthLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrengthLevel

`func (o *Monster) SetStrengthLevel(v int32)`

SetStrengthLevel sets StrengthLevel field to given value.


### GetDefenceLevel

`func (o *Monster) GetDefenceLevel() int32`

GetDefenceLevel returns the DefenceLevel field if non-nil, zero value otherwise.

### GetDefenceLevelOk

`func (o *Monster) GetDefenceLevelOk() (*int32, bool)`

GetDefenceLevelOk returns a tuple with the DefenceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefenceLevel

`func (o *Monster) SetDefenceLevel(v int32)`

SetDefenceLevel sets DefenceLevel field to given value.


### GetMagicLevel

`func (o *Monster) GetMagicLevel() int32`

GetMagicLevel returns the MagicLevel field if non-nil, zero value otherwise.

### GetMagicLevelOk

`func (o *Monster) GetMagicLevelOk() (*int32, bool)`

GetMagicLevelOk returns a tuple with the MagicLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagicLevel

`func (o *Monster) SetMagicLevel(v int32)`

SetMagicLevel sets MagicLevel field to given value.


### GetRangedLevel

`func (o *Monster) GetRangedLevel() int32`

GetRangedLevel returns the RangedLevel field if non-nil, zero value otherwise.

### GetRangedLevelOk

`func (o *Monster) GetRangedLevelOk() (*int32, bool)`

GetRangedLevelOk returns a tuple with the RangedLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangedLevel

`func (o *Monster) SetRangedLevel(v int32)`

SetRangedLevel sets RangedLevel field to given value.


### GetAttackStab

`func (o *Monster) GetAttackStab() int32`

GetAttackStab returns the AttackStab field if non-nil, zero value otherwise.

### GetAttackStabOk

`func (o *Monster) GetAttackStabOk() (*int32, bool)`

GetAttackStabOk returns a tuple with the AttackStab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackStab

`func (o *Monster) SetAttackStab(v int32)`

SetAttackStab sets AttackStab field to given value.


### GetAttackSlash

`func (o *Monster) GetAttackSlash() int32`

GetAttackSlash returns the AttackSlash field if non-nil, zero value otherwise.

### GetAttackSlashOk

`func (o *Monster) GetAttackSlashOk() (*int32, bool)`

GetAttackSlashOk returns a tuple with the AttackSlash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackSlash

`func (o *Monster) SetAttackSlash(v int32)`

SetAttackSlash sets AttackSlash field to given value.


### GetAttackCrush

`func (o *Monster) GetAttackCrush() int32`

GetAttackCrush returns the AttackCrush field if non-nil, zero value otherwise.

### GetAttackCrushOk

`func (o *Monster) GetAttackCrushOk() (*int32, bool)`

GetAttackCrushOk returns a tuple with the AttackCrush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackCrush

`func (o *Monster) SetAttackCrush(v int32)`

SetAttackCrush sets AttackCrush field to given value.


### GetAttackMagic

`func (o *Monster) GetAttackMagic() int32`

GetAttackMagic returns the AttackMagic field if non-nil, zero value otherwise.

### GetAttackMagicOk

`func (o *Monster) GetAttackMagicOk() (*int32, bool)`

GetAttackMagicOk returns a tuple with the AttackMagic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackMagic

`func (o *Monster) SetAttackMagic(v int32)`

SetAttackMagic sets AttackMagic field to given value.


### GetAttackRanged

`func (o *Monster) GetAttackRanged() int32`

GetAttackRanged returns the AttackRanged field if non-nil, zero value otherwise.

### GetAttackRangedOk

`func (o *Monster) GetAttackRangedOk() (*int32, bool)`

GetAttackRangedOk returns a tuple with the AttackRanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackRanged

`func (o *Monster) SetAttackRanged(v int32)`

SetAttackRanged sets AttackRanged field to given value.


### GetDefenceStab

`func (o *Monster) GetDefenceStab() int32`

GetDefenceStab returns the DefenceStab field if non-nil, zero value otherwise.

### GetDefenceStabOk

`func (o *Monster) GetDefenceStabOk() (*int32, bool)`

GetDefenceStabOk returns a tuple with the DefenceStab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefenceStab

`func (o *Monster) SetDefenceStab(v int32)`

SetDefenceStab sets DefenceStab field to given value.


### GetDefenceSlash

`func (o *Monster) GetDefenceSlash() int32`

GetDefenceSlash returns the DefenceSlash field if non-nil, zero value otherwise.

### GetDefenceSlashOk

`func (o *Monster) GetDefenceSlashOk() (*int32, bool)`

GetDefenceSlashOk returns a tuple with the DefenceSlash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefenceSlash

`func (o *Monster) SetDefenceSlash(v int32)`

SetDefenceSlash sets DefenceSlash field to given value.


### GetDefenceCrush

`func (o *Monster) GetDefenceCrush() int32`

GetDefenceCrush returns the DefenceCrush field if non-nil, zero value otherwise.

### GetDefenceCrushOk

`func (o *Monster) GetDefenceCrushOk() (*int32, bool)`

GetDefenceCrushOk returns a tuple with the DefenceCrush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefenceCrush

`func (o *Monster) SetDefenceCrush(v int32)`

SetDefenceCrush sets DefenceCrush field to given value.


### GetDefenceMagic

`func (o *Monster) GetDefenceMagic() int32`

GetDefenceMagic returns the DefenceMagic field if non-nil, zero value otherwise.

### GetDefenceMagicOk

`func (o *Monster) GetDefenceMagicOk() (*int32, bool)`

GetDefenceMagicOk returns a tuple with the DefenceMagic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefenceMagic

`func (o *Monster) SetDefenceMagic(v int32)`

SetDefenceMagic sets DefenceMagic field to given value.


### GetDefenceRanged

`func (o *Monster) GetDefenceRanged() int32`

GetDefenceRanged returns the DefenceRanged field if non-nil, zero value otherwise.

### GetDefenceRangedOk

`func (o *Monster) GetDefenceRangedOk() (*int32, bool)`

GetDefenceRangedOk returns a tuple with the DefenceRanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefenceRanged

`func (o *Monster) SetDefenceRanged(v int32)`

SetDefenceRanged sets DefenceRanged field to given value.


### GetAttackAccuracy

`func (o *Monster) GetAttackAccuracy() int32`

GetAttackAccuracy returns the AttackAccuracy field if non-nil, zero value otherwise.

### GetAttackAccuracyOk

`func (o *Monster) GetAttackAccuracyOk() (*int32, bool)`

GetAttackAccuracyOk returns a tuple with the AttackAccuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackAccuracy

`func (o *Monster) SetAttackAccuracy(v int32)`

SetAttackAccuracy sets AttackAccuracy field to given value.


### GetMeleeStrength

`func (o *Monster) GetMeleeStrength() int32`

GetMeleeStrength returns the MeleeStrength field if non-nil, zero value otherwise.

### GetMeleeStrengthOk

`func (o *Monster) GetMeleeStrengthOk() (*int32, bool)`

GetMeleeStrengthOk returns a tuple with the MeleeStrength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeleeStrength

`func (o *Monster) SetMeleeStrength(v int32)`

SetMeleeStrength sets MeleeStrength field to given value.


### GetRangedStrength

`func (o *Monster) GetRangedStrength() int32`

GetRangedStrength returns the RangedStrength field if non-nil, zero value otherwise.

### GetRangedStrengthOk

`func (o *Monster) GetRangedStrengthOk() (*int32, bool)`

GetRangedStrengthOk returns a tuple with the RangedStrength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangedStrength

`func (o *Monster) SetRangedStrength(v int32)`

SetRangedStrength sets RangedStrength field to given value.


### GetMagicDamage

`func (o *Monster) GetMagicDamage() int32`

GetMagicDamage returns the MagicDamage field if non-nil, zero value otherwise.

### GetMagicDamageOk

`func (o *Monster) GetMagicDamageOk() (*int32, bool)`

GetMagicDamageOk returns a tuple with the MagicDamage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagicDamage

`func (o *Monster) SetMagicDamage(v int32)`

SetMagicDamage sets MagicDamage field to given value.


### GetDrops

`func (o *Monster) GetDrops() []MonsterDrops`

GetDrops returns the Drops field if non-nil, zero value otherwise.

### GetDropsOk

`func (o *Monster) GetDropsOk() (*[]MonsterDrops, bool)`

GetDropsOk returns a tuple with the Drops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrops

`func (o *Monster) SetDrops(v []MonsterDrops)`

SetDrops sets Drops field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


