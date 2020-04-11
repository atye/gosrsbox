package gosrsbox

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strings"

	"github.com/atye/gosrsbox/internal/lib"
)

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

type itemsResponse struct {
	Items []*Item `json:"_items"`
	Meta  struct {
		Page       int `json:"page"`
		MaxResults int `json:"max_results"`
		Total      int `json:"total"`
	} `json:"_meta"`
	Error *serverError `json:"_error"`
}

func getAllItems(ctx context.Context, client HTTPClient) ([]*Item, error) {
	if client == nil {
		return nil, errors.New("no client configured")
	}

	req, err := http.NewRequestWithContext(ctx, "GET", itemsCompleteURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error doing request: %w", err)
	}
	defer resp.Body.Close()

	var itemsMap map[string]*Item
	err = json.NewDecoder(resp.Body).Decode(&itemsMap)
	if err != nil {
		return nil, fmt.Errorf("error decoding json response: %w", err)
	}

	var items []*Item
	for _, item := range itemsMap {
		items = append(items, item)
	}

	return items, nil
}

func getItemsByName(ctx context.Context, client HTTPClient, names ...string) ([]*Item, error) {
	if client == nil {
		return nil, errors.New("no client configured")
	}

	var nameData []string
	var query string

	for _, name := range names {
		nameData = append(nameData, fmt.Sprintf(`"%s"`, lib.MakeValidItemName(name)))
	}
	query = fmt.Sprintf(`{ "name": { "$in": [%s] }, "duplicate": false }`, strings.Join(nameData, ", "))

	return getItemsWhere(ctx, client, query)
}

func getItemsWhere(ctx context.Context, client HTTPClient, query string) ([]*Item, error) {
	if client == nil {
		return nil, errors.New("no client configured")
	}

	url := fmt.Sprintf("%s/%s?where=%s", api, itemsEndpoint, url.QueryEscape(query))

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error doing request: %w", err)
	}

	var items []*Item

	var itemsResp *itemsResponse
	err = json.NewDecoder(resp.Body).Decode(&itemsResp)
	if err != nil {
		return nil, fmt.Errorf("error decoding json response: %w", err)
	}
	resp.Body.Close()

	if itemsResp.Error != nil {
		return nil, fmt.Errorf("error from server: %w", itemsResp.Error)
	}

	items = append(items, itemsResp.Items...)

	var pages int
	if itemsResp.Meta.MaxResults != 0 {
		pages = int(math.Ceil(float64(itemsResp.Meta.Total) / float64(itemsResp.Meta.MaxResults)))
	}

	for i := 2; i <= pages; i++ {
		req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s&page=%d", url, i), nil)
		if err != nil {
			return nil, fmt.Errorf("error creating request: %w", err)
		}

		resp, err := client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("error doing request: %w", err)
		}

		var itemsRespTemp *itemsResponse
		err = json.NewDecoder(resp.Body).Decode(&itemsRespTemp)
		if err != nil {
			return nil, fmt.Errorf("error decoding json response: %w", err)
		}
		resp.Body.Close()

		if itemsRespTemp.Error != nil {
			return nil, fmt.Errorf("error from server: %w", itemsRespTemp.Error)
		}

		items = append(items, itemsRespTemp.Items...)
	}

	return items, nil
}
