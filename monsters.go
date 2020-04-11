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

type monstersResponse struct {
	Monsters []*Monster `json:"_items"`
	Meta     struct {
		Page       int `json:"page"`
		MaxResults int `json:"max_results"`
		Total      int `json:"total"`
	} `json:"_meta"`
	Error *serverError `json:"_error"`
}

func getAllMonsters(ctx context.Context, client HTTPClient) ([]*Monster, error) {
	if client == nil {
		return nil, errors.New("no client configured")
	}

	req, err := http.NewRequestWithContext(ctx, "GET", monstersCompleteURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error doing request: %w", err)
	}
	defer resp.Body.Close()

	var monstersMap map[string]*Monster
	err = json.NewDecoder(resp.Body).Decode(&monstersMap)
	if err != nil {
		return nil, fmt.Errorf("error decoding json response: %w", err)
	}

	var monsters []*Monster
	for _, monster := range monstersMap {
		monsters = append(monsters, monster)
	}

	return monsters, nil
}

func getMonstersByName(ctx context.Context, client HTTPClient, names ...string) ([]*Monster, error) {
	if client == nil {
		return nil, errors.New("no client configured")
	}

	var nameData []string
	var query string

	for _, name := range names {
		nameData = append(nameData, fmt.Sprintf(`"%s"`, name))
	}
	query = fmt.Sprintf(`{ "name": { "$in": [%s] }, "duplicate": false }`, strings.Join(nameData, ", "))

	return getMonstersWhere(ctx, client, query)

}

func getMonstersThatDrop(ctx context.Context, client HTTPClient, names ...string) ([]*Monster, error) {
	if client == nil {
		return nil, errors.New("no client configured")
	}

	var nameData []string
	for _, name := range names {
		nameData = append(nameData, fmt.Sprintf(`"%s"`, lib.MakeValidItemName(name)))
	}

	query := fmt.Sprintf(`{ "drops": { "$elemMatch": { "name": { "$in": [%s] } } }, "duplicate": false }`, strings.Join(nameData, ", "))

	return getMonstersWhere(ctx, client, query)
}

func getMonstersWhere(ctx context.Context, client HTTPClient, query string) ([]*Monster, error) {
	if client == nil {
		return nil, errors.New("no client configured")
	}

	url := fmt.Sprintf("%s/%s?where=%s", api, monstersEndpoint, url.QueryEscape(query))

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error doing request: %w", err)
	}

	var monsters []*Monster

	var monstersResp *monstersResponse
	err = json.NewDecoder(resp.Body).Decode(&monstersResp)
	if err != nil {
		return nil, fmt.Errorf("error decoding json response: %w", err)
	}
	resp.Body.Close()

	if monstersResp.Error != nil {
		return nil, fmt.Errorf("error from server: %w", monstersResp.Error)
	}

	monsters = append(monsters, monstersResp.Monsters...)

	var pages int
	if monstersResp.Meta.MaxResults != 0 {
		pages = int(math.Ceil(float64(monstersResp.Meta.Total) / float64(monstersResp.Meta.MaxResults)))
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

		var monstersRespTemp *monstersResponse
		err = json.NewDecoder(resp.Body).Decode(&monstersRespTemp)
		if err != nil {
			return nil, fmt.Errorf("error decoding json response: %w", err)
		}
		resp.Body.Close()

		if monstersRespTemp.Error != nil {
			return nil, monstersRespTemp.Error
		}

		monsters = append(monsters, monstersRespTemp.Monsters...)
	}

	return monsters, nil
}
