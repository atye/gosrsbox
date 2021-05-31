package models

type Monster struct {
	// Unique OSRS monster ID number.
	Id string `json:"id"`
	// The name of the monster.
	Name string `json:"name"`
	// The last time (UTC) the monster was updated (in ISO8601 date format).
	LastUpdated string `json:"last_updated"`
	// If the monster has incomplete wiki data.
	Incomplete bool `json:"incomplete"`
	// If the monster is members only, or not.
	Members bool `json:"members"`
	// The release date of the monster (in ISO8601 date format).
	ReleaseDate string `json:"release_date"`
	// The combat level of the monster.
	CombatLevel int `json:"combat_level"`
	// The size, in tiles, of the monster.
	Size int `json:"size"`
	// The number of hitpoints a monster has.
	Hitpoints int `json:"hitpoints"`
	// The maximum hit of the monster.
	MaxHit int `json:"max_hit"`
	// The attack style (e.g., melee, magic, range) of the monster.
	AttackType []string `json:"attack_type"`
	// The attack speed (in game ticks) of the monster.
	AttackSpeed int `json:"attack_speed"`
	// If the monster is aggressive, or not.
	Aggressive bool `json:"aggressive"`
	// If the monster poisons, or not
	Poisonous bool `json:"poisonous"`
	// If the monster poisons using venom, or not
	Venomous bool `json:"venomous"`
	// If the monster is immune to poison, or not
	ImmunePoison bool `json:"immune_poison"`
	// If the monster is immune to venom, or not
	ImmuneVenom bool `json:"immune_venom"`
	// An array of monster attributes.
	Attributes []string `json:"attributes"`
	// An array of monster category.
	Category []string `json:"category"`
	// If the monster is a potential slayer task.
	SlayerMonster bool `json:"slayer_monster"`
	// The slayer level required to kill the monster.
	SlayerLevel int `json:"slayer_level"`
	// The slayer XP rewarded for a monster kill.
	SlayerXp float32 `json:"slayer_xp"`
	// The slayer masters who can assign the monster.
	SlayerMasters []string `json:"slayer_masters"`
	// If the monster is a duplicate.
	Duplicate bool `json:"duplicate"`
	// The examine text of the monster.
	Examine string `json:"examine"`
	// The OSRS Wiki name for the monster.
	WikiName string `json:"wiki_name"`
	// The OSRS Wiki URL (possibly including anchor link).
	WikiUrl string `json:"wiki_url"`
	// The attack level of the monster.
	AttackLevel int `json:"attack_level"`
	// The strength level of the monster.
	StrengthLevel int `json:"strength_level"`
	// The defence level of the monster.
	DefenceLevel int `json:"defence_level"`
	// The magic level of the monster.
	MagicLevel int `json:"magic_level"`
	// The ranged level of the monster.
	RangedLevel int `json:"ranged_level"`
	// The attack bonus of the monster.
	AttackBonus int `json:"attack_bonus"`
	// The strength bonus of the monster.
	StrengthBonus int `json:"strength_bonus"`
	// The magic attack of the monster.
	AttackMagic int `json:"attack_magic"`
	// The magic bonus of the monster.
	MagicBonus int `json:"magic_bonus"`
	// The ranged attack of the monster.
	AttackRanged int `json:"attack_ranged"`
	// The ranged bonus of the monster.
	RangedBonus int `json:"ranged_bonus"`
	// The defence stab bonus of the monster.
	DefenceStab int `json:"defence_stab"`
	// The defence slash bonus of the monster.
	DefenceSlash int `json:"defence_slash"`
	// The defence crush bonus of the monster.
	DefenceCrush int `json:"defence_crush"`
	// The defence magic bonus of the monster.
	DefenceMagic int `json:"defence_magic"`
	// The defence ranged bonus of the monster.
	DefenceRanged int `json:"defence_ranged"`
	// An array of monster drop objects.
	Drops []MonsterDrops `json:"drops"`
}

type MonsterDrops struct {
	// The ID number of the item drop.
	Id int `json:"id"`
	// The name of the item drop.
	Name string `json:"name"`
	// If the drop is a members-only item.
	Members bool `json:"members"`
	// The quantity of the item drop (integer, comma-separated or range).
	Quantity string `json:"quantity"`
	// If the item drop is noted, or not.
	Noted bool `json:"noted"`
	// The rarity of the item drop (as a float out of 1.0).
	Rarity float32 `json:"rarity"`
	// Number of rolls from the drop.
	Rolls int `json:"rolls"`
}
