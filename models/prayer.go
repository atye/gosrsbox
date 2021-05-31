package models

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
	DrainPerMinute float64 `json:"drain_per_minute"`
	// The OSRS Wiki URL.
	WikiUrl string `json:"wiki_url"`
	// The stat requirements to use the prayer.
	Requirements map[string]int `json:"requirements"`
	// The bonuses a prayer provides.
	Bonuses map[string]int `json:"bonuses"`
	// The prayer icon.
	Icon string `json:"icon"`
}
