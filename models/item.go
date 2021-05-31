package models

type Item struct {
	// Unique OSRS item ID number.
	Id string `json:"id"`
	// The name of the item.
	Name string `json:"name"`
	// The last time (UTC) the item was updated (in ISO8601 date format).
	LastUpdated string `json:"last_updated"`
	// If the item has incomplete wiki data.
	Incomplete bool `json:"incomplete"`
	// If the item is a members-only.
	Members bool `json:"members"`
	// If the item is tradeable (between players and on the GE).
	Tradeable bool `json:"tradeable"`
	// If the item is tradeable (only on GE).
	TradeableOnGe bool `json:"tradeable_on_ge"`
	// If the item is stackable (in inventory).
	Stackable bool `json:"stackable"`
	// If the item is stacked, indicated by the stack count.
	Stacked int `json:"stacked"`
	// If the item is noted.
	Noted bool `json:"noted"`
	// If the item is noteable.
	Noteable bool `json:"noteable"`
	// The linked ID of the actual item (if noted/placeholder).
	LinkedIdItem int `json:"linked_id_item"`
	// The linked ID of an item in noted form.
	LinkedIdNoted int `json:"linked_id_noted"`
	// The linked ID of an item in placeholder form.
	LinkedIdPlaceholder int `json:"linked_id_placeholder"`
	// If the item is a placeholder.
	Placeholder bool `json:"placeholder"`
	// If the item is equipable (based on right-click menu entry).
	Equipable bool `json:"equipable"`
	// If the item is equipable in-game by a player.
	EquipableByPlayer bool `json:"equipable_by_player"`
	// If the item is an equipable weapon.
	EquipableWeapon bool `json:"equipable_weapon"`
	// The store price of an item.
	Cost int32 `json:"cost"`
	// The low alchemy value of the item (cost * 0.4).
	Lowalch int `json:"lowalch"`
	// The high alchemy value of the item (cost * 0.6).
	Highalch int `json:"highalch"`
	// The weight (in kilograms) of the item.
	Weight float64 `json:"weight"`
	// The Grand Exchange buy limit of the item.
	BuyLimit int `json:"buy_limit"`
	// If the item is associated with a quest.
	QuestItem bool `json:"quest_item"`
	// Date the item was released (in ISO8601 format).
	ReleaseDate string `json:"release_date"`
	// If the item is a duplicate.
	Duplicate bool `json:"duplicate"`
	// The examine text for the item.
	Examine string `json:"examine"`
	// The item icon (in base64 encoding).
	Icon string `json:"icon"`
	// The OSRS Wiki name for the item.
	WikiName string `json:"wiki_name"`
	// The OSRS Wiki URL (possibly including anchor link).
	WikiUrl   string    `json:"wiki_url"`
	Equipment Equipment `json:"equipment"`
	Weapon    Weapon    `json:"weapon"`
}

type Equipment struct {
	// The attack stab bonus of the item.
	AttackStab int `json:"attack_stab"`
	// The attack slash bonus of the item.
	AttackSlash int `json:"attack_slash"`
	// The attack crush bonus of the item.
	AttackCrush int `json:"attack_crush"`
	// The attack magic bonus of the item.
	AttackMagic int `json:"attack_magic"`
	// The attack ranged bonus of the item.
	AttackRanged int `json:"attack_ranged"`
	// The defence stab bonus of the item.
	DefenceStab int `json:"defence_stab"`
	// The defence slash bonus of the item.
	DefenceSlash int `json:"defence_slash"`
	// The defence crush bonus of the item.
	DefenceCrush int `json:"defence_crush"`
	// The defence magic bonus of the item.
	DefenceMagic int `json:"defence_magic"`
	// The defence ranged bonus of the item.
	DefenceRanged int `json:"defence_ranged"`
	// The melee strength bonus of the item.
	MeleeStrength int `json:"melee_strength"`
	// The ranged strength bonus of the item.
	RangedStrength int `json:"ranged_strength"`
	// The magic damage bonus of the item.
	MagicDamage int `json:"magic_damage"`
	// The prayer bonus of the item.
	Prayer int `json:"prayer"`
	// The equipment slot associated with the item (e.g., head).
	Slot string `json:"slot"`
	// An object of requirements {skill: level}.
	Requirements map[string]int `json:"requirements"`
}

type Weapon struct {
	// The attack speed of a weapon (in game ticks).
	AttackSpeed int `json:"attack_speed"`
	// The weapon classification (e.g., axes)
	WeaponType string `json:"weapon_type"`
	// An array of weapon stance information.
	Stances []WeaponStances `json:"stances"`
}

type WeaponStances struct {
	CombatStyle string `json:"combat_style"`
	AttackType  string `json:"attack_type"`
	AttackStyle string `json:"attack_style"`
	Experience  string `json:"experience"`
	Boosts      string `json:"boosts"`
}
