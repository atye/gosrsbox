package gosrsbox

type itemsEndpoint struct {
	Items []*Item `json:"_items"`
}

// Item is an osrsbox item.
// https://api.osrsbox.com/swaggerui#/Item/get_items
type Item struct {
	ID                  interface{} `json:"id"`
	Name                string      `json:"name"`
	Incomplete          bool        `json:"incomplete"`
	Members             bool        `json:"members"`
	Tradeable           bool        `json:"tradeable"`
	TradeableOnGe       bool        `json:"tradeable_on_ge"`
	Stackable           bool        `json:"stackable"`
	Stacked             int         `json:"stacked"`
	Noted               bool        `json:"noted"`
	Noteable            bool        `json:"noteable"`
	LinkedIDItem        int         `json:"linked_id_item"`
	LinkedIDNoted       int         `json:"linked_id_noted"`
	LinkedIDPlaceholder int         `json:"linked_id_placeholder"`
	Placeholder         bool        `json:"placeholder"`
	Equipable           bool        `json:"equipable"`
	EquipableByPlayer   bool        `json:"equipable_by_player"`
	EquipableWeapon     bool        `json:"equipable_weapon"`
	Cost                int         `json:"cost"`
	Lowalch             int         `json:"lowalch"`
	Highalch            int         `json:"highalch"`
	Weight              float64     `json:"weight"`
	BuyLimit            int         `json:"buy_limit"`
	QuestItem           bool        `json:"quest_item"`
	ReleaseDate         string      `json:"release_date"`
	Duplicate           bool        `json:"duplicate"`
	Examine             string      `json:"examine"`
	Icon                string      `json:"icon"`
	WikiName            string      `json:"wiki_name"`
	WikiURL             string      `json:"wiki_url"`
	Equipment           *Equipment  `json:"equipment"`
	Weapon              *Weapon     `json:"weapon"`
}

// Equipment bonuses of equipable armor/weapons.
// https://api.osrsbox.com/swaggerui#/Equipment/get_equipment
type Equipment struct {
	AttackStab     int            `json:"attack_stab"`
	AttackSlash    int            `json:"attack_slash"`
	AttackCrush    int            `json:"attack_crush"`
	AttackMagic    int            `json:"attack_magic"`
	AttackRanged   int            `json:"attack_ranged"`
	DefenceStab    int            `json:"defence_stab"`
	DefenceSlash   int            `json:"defence_slash"`
	DefenceCrush   int            `json:"defence_crush"`
	DefenceMagic   int            `json:"defence_magic"`
	DefenceRanged  int            `json:"defence_ranged"`
	MeleeStrength  int            `json:"melee_strength"`
	RangedStrength int            `json:"ranged_strength"`
	MagicDamage    int            `json:"magic_damage"`
	Prayer         int            `json:"prayer"`
	Slot           string         `json:"slot"`
	Requirements   map[string]int `json:"requirements"`
}

// Weapon bonuses including attack speed, type and stances.
// https://api.osrsbox.com/swaggerui#/Weapon/get_weapons
type Weapon struct {
	AttackSpeed int       `json:"attack_speed"`
	WeaponType  string    `json:"weapon_type"`
	Stances     []*Stance `json:"stances"`
}

// Stance information for a weapon
type Stance struct {
	CombatStyle string `json:"combat_style"`
	AttackType  string `json:"attack_type"`
	AttackStyle string `json:"attack_style"`
	Experience  string `json:"experience"`
	Boosts      string `json:"boosts"`
}

type monstersEndpoint struct {
	Monsters []*Monster `json:"_items"`
}

// Monster is an osrsbox monster.
// https://api.osrsbox.com/swaggerui#/Monster/get_monsters
type Monster struct {
	ID             interface{} `json:"id"`
	Name           string      `json:"name"`
	Incomplete     bool        `json:"incomplete"`
	Members        bool        `json:"members"`
	ReleaseDate    string      `json:"release_date"`
	CombatLevel    int         `json:"combat_level"`
	Size           int         `json:"size"`
	Hitpoints      int         `json:"hitpoints"`
	MaxHit         int         `json:"max_hit"`
	AttackType     []string    `json:"attack_type"`
	AttackSpeed    int         `json:"attack_speed"`
	Aggressive     bool        `json:"aggressive"`
	Poisonous      bool        `json:"poisonous"`
	ImmunePoison   bool        `json:"immune_poison"`
	ImmuneVenom    bool        `json:"immune_venom"`
	Attributes     []string    `json:"attributes"`
	Category       []string    `json:"category"`
	SlayerMonster  bool        `json:"slayer_monster"`
	SlayerLevel    int         `json:"slayer_level"`
	SlayerXp       float64     `json:"slayer_xp"`
	SlayerMasters  []string    `json:"slayer_masters"`
	Duplicate      bool        `json:"duplicate"`
	Examine        string      `json:"examine"`
	Icon           string      `json:"icon"`
	WikiName       string      `json:"wiki_name"`
	WikiURL        string      `json:"wiki_url"`
	AttackLevel    int         `json:"attack_level"`
	StrengthLevel  int         `json:"strength_level"`
	DefenceLevel   int         `json:"defence_level"`
	MagicLevel     int         `json:"magic_level"`
	RangedLevel    int         `json:"ranged_level"`
	AttackStab     int         `json:"attack_stab"`
	AttackSlash    int         `json:"attack_slash"`
	AttackCrush    int         `json:"attack_crush"`
	AttackMagic    int         `json:"attack_magic"`
	AttackRanged   int         `json:"attack_ranged"`
	DefenceStab    int         `json:"defence_stab"`
	DefenceSlash   int         `json:"defence_slash"`
	DefenceCrush   int         `json:"defence_crush"`
	DefenceMagic   int         `json:"defence_magic"`
	DefenceRanged  int         `json:"defence_ranged"`
	AttackAccuracy int         `json:"attack_accuracy"`
	MeleeStrength  int         `json:"melee_strength"`
	RangedStrength int         `json:"ranged_strength"`
	MagicDamage    int         `json:"magic_damage"`
	Drops          []*Drop     `json:"drops"`
}

// Drop is an entity dropped by a monster.
type Drop struct {
	ID               int         `json:"id"`
	Name             string      `json:"name"`
	Members          bool        `json:"members"`
	Quantity         string      `json:"quantity"`
	Noted            bool        `json:"noted"`
	Rarity           float64     `json:"rarity"`
	DropRequirements interface{} `json:"drop_requirements"`
}

type prayersEndpoint struct {
	Prayers []*Prayer `json:"_items"`
}

// Prayer is an osrsbox prayer.
// https://api.osrsbox.com/swaggerui#/Prayer/get_prayers
type Prayer struct {
	ID             interface{}    `json:"id"`
	Name           string         `json:"name"`
	Members        bool           `json:"members"`
	Description    string         `json:"description"`
	DrainPerMinute float32        `json:"drain_per_minute"`
	WikiURL        string         `json:"wiki_url"`
	Requirements   map[string]int `json:"requirements"`
	Bonuses        map[string]int `json:"bonuses"`
	Icon           string         `json:"icon"`
}
