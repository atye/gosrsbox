package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strings"

	openapi "github.com/atye/gosrsbox/internal/openapi/api"
	"github.com/atye/gosrsbox/sets"
	"github.com/atye/gosrsbox/slots"
	"golang.org/x/sync/errgroup"
)

// Item struct for an Item
type Item struct {
	// Unique OSRS item ID number.
	Id string `json:"id"`
	// The name of the item.
	Name string `json:"name"`
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
	Cost int `json:"cost"`
	// The low alchemy value of the item (cost * 0.4).
	Lowalch int `json:"lowalch"`
	// The high alchemy value of the item (cost * 0.6).
	Highalch int `json:"highalch"`
	// The weight (in kilograms) of the item.
	Weight float32 `json:"weight"`
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
	WikiUrl string `json:"wiki_url"`
	// The OSRS Wiki Exchange URL.
	WikiExchange string    `json:"wiki_exchange"`
	Equipment    Equipment `json:"equipment"`
	Weapon       Weapon    `json:"weapon"`
}

// Equipment bonuses of equipable armour/weapons.
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
	Requirements map[string]interface{} `json:"requirements"`
}

// Weapon bonuses including attack speed, type and stance.
type Weapon struct {
	// The attack speed of a weapon (in game ticks).
	AttackSpeed int `json:"attack_speed"`
	// The weapon classification (e.g., axes).
	WeaponType string `json:"weapon_type"`
	// An array of weapon stance information.
	Stances []WeaponStances `json:"stances"`
}

// WeaponStances including styles, experience, and boosts
type WeaponStances struct {
	CombatStyle string `json:"combat_style"`
	AttackType  string `json:"attack_type"`
	AttackStyle string `json:"attack_style"`
	Experience  string `json:"experience"`
	Boosts      string `json:"boosts"`
}

func (c *client) GetItemsByID(ctx context.Context, ids ...string) ([]Item, error) {
	if len(ids) == 0 {
		return nil, errors.New("no ids provided")
	}
	return c.GetItemsByQuery(ctx, fmt.Sprintf(`{ "id": { "$in": [%s] }, "duplicate": false }`, strings.Join(quoteStrings(ids...), ", ")))
}

func (c *client) GetItemsByName(ctx context.Context, names ...string) ([]Item, error) {
	if len(names) == 0 {
		return nil, errors.New("no names provided")
	}
	return c.GetItemsByQuery(ctx, fmt.Sprintf(`{ "wiki_name": { "$in": [%s] }, "duplicate": false }`, strings.Join(quoteStrings(names...), ", ")))
}

func (c *client) GetItemSet(ctx context.Context, set sets.SetName) ([]Item, error) {
	if set == nil || len(set) == 0 {
		return nil, errors.New("no set provided")
	}
	return c.GetItemsByName(ctx, set...)
}

func (c *client) GetItemsBySlot(ctx context.Context, slot slots.SlotName) ([]Item, error) {
	if slot == "" || len(slot) == 0 {
		return nil, errors.New("no set provided")
	}
	return c.GetItemsByQuery(ctx, fmt.Sprintf(`{ "equipable_by_player": true, "equipment.slot": %s, "duplicate": false }`, slot))
}

func (c *client) GetItemsByQuery(ctx context.Context, query string) ([]Item, error) {
	resp, err := c.doOpenAPIRequest(ctx, c.openAPIClient.ItemApi.Getitems(ctx).Where(query))
	if err != nil {
		return nil, err
	}

	switch inline := resp.(type) {
	case openapi.InlineResponse200:
		pages := int(math.Ceil(float64(*inline.Meta.Total) / float64(*inline.Meta.MaxResults)))
		items := make([]Item, *inline.Meta.Total)

		respItems, err := openItemsToItems(inline.GetItems())
		if err != nil {
			return nil, err
		}

		_ = copy(items, respItems)

		if pages > 1 {
			var eg errgroup.Group
			for page := 2; page <= pages; page++ {
				page := page
				eg.Go(func() error {
					resp, err := c.doOpenAPIRequest(ctx, c.openAPIClient.ItemApi.Getitems(ctx).Where(query).Page(int32(page)))
					if err != nil {
						return err
					}

					switch inline := resp.(type) {
					case openapi.InlineResponse200:
						respItems, err := openItemsToItems(inline.GetItems())
						if err != nil {
							return err
						}

						for i, item := range respItems {
							items[int(*inline.Meta.MaxResults)*(page-1)+i] = item
						}
					default:
						return fmt.Errorf("unexpected response type %T", inline)
					}

					return nil
				})
			}
			err := eg.Wait()
			if err != nil {
				return nil, err
			}
		}
		return items, nil
	default:
		return nil, fmt.Errorf("unexpected response type %T", inline)
	}
}

func openItemsToItems(openItems []openapi.Item) ([]Item, error) {
	b, err := json.Marshal(openItems)
	if err != nil {
		return nil, err
	}

	var items []Item
	err = json.Unmarshal(b, &items)
	if err != nil {
		return nil, err
	}

	return items, nil
}
