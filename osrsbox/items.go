package osrsbox

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strings"
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

func getAllItems(ctx context.Context, c *client) ([]*Item, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", itemsCompleteURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error doing request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error from server: %w", &serverError{
			Code:    resp.StatusCode,
			Message: "something went wrong",
		})
	}

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

func getItemsByName(ctx context.Context, c *client, names ...string) ([]*Item, error) {
	var nameData []string
	for _, name := range names {
		nameData = append(nameData, fmt.Sprintf(`"%s"`, makeValidItemName(name)))
	}

	query := fmt.Sprintf(`{ "name": { "$in": [%s] }, "duplicate": false }`, strings.Join(nameData, ", "))
	return getItemsWhere(ctx, c, query)
}

func getItemsByWikiName(ctx context.Context, c *client, names ...string) ([]*Item, error) {
	var nameData []string
	for _, name := range names {
		nameData = append(nameData, fmt.Sprintf(`"%s"`, name))
	}

	query := fmt.Sprintf(`{ "wiki_name": { "$in": [%s] }, "duplicate": false }`, strings.Join(nameData, ", "))
	return getItemsWhere(ctx, c, query)
}

func getItemsWhere(ctx context.Context, c *client, query string) ([]*Item, error) {
	url := fmt.Sprintf("%s/%s?where=%s", api, itemsEndpoint, url.QueryEscape(query))

	itemsResp, err := doItemsRespRequest(ctx, c, url)
	if err != nil {
		return nil, err
	}

	items := make([]*Item, itemsResp.Meta.Total)

	for i, item := range itemsResp.Items {
		items[i] = item
	}

	var pages int
	if itemsResp.Meta.MaxResults != 0 {
		pages = int(math.Ceil(float64(itemsResp.Meta.Total) / float64(itemsResp.Meta.MaxResults)))
	}

	if pages > 1 {

		errChan := make(chan error)
		waitChan := make(chan struct{})

		go func() {
			for page := 2; page <= pages; page++ {
				c.wg.Add(1)
				go func(page int) {
					defer c.wg.Done()

					itemsRespTemp, err := doItemsRespRequest(ctx, c, fmt.Sprintf("%s&page=%d", url, page))
					if err != nil {
						errChan <- err
						return
					}

					for i, item := range itemsRespTemp.Items {
						c.mu.Lock()
						items[itemsResp.Meta.MaxResults*(page-1)+i] = item
						c.mu.Unlock()
					}

				}(page)
			}
			c.wg.Wait()
			close(waitChan)
		}()

		select {
		case <-waitChan:
			return items, nil
		case err := <-errChan:
			return nil, err
		}
	}

	return items, nil
}

func doItemsRespRequest(ctx context.Context, c *client, url string) (*itemsResponse, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error doing request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var itemsResp *itemsResponse
		_ = json.NewDecoder(resp.Body).Decode(&itemsResp)
		defer resp.Body.Close()

		if itemsResp != nil && itemsResp.Error != nil {
			return nil, fmt.Errorf("error from server: %w", itemsResp.Error)
		}

		return nil, fmt.Errorf("error some server: %w", &serverError{
			Code:    resp.StatusCode,
			Message: "something went wrong",
		})
	}

	var itemsResp *itemsResponse
	err = json.NewDecoder(resp.Body).Decode(&itemsResp)
	if err != nil {
		return nil, fmt.Errorf("error decoding json response: %w", err)
	}
	defer resp.Body.Close()

	return itemsResp, nil
}
