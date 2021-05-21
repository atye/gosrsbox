package api

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"testing"

	openapi "github.com/atye/gosrsbox/internal/openapi/api"
	"github.com/atye/gosrsbox/sets"
	"github.com/atye/gosrsbox/slots"
)

func TestItems(t *testing.T) {
	apiSvr := setupItemsAPISvr()
	defer apiSvr.Close()
	t.Run("GetItemsByID", test_GetItemsByID)
	t.Run("GetItemsByName", test_GetItemsByName)
	t.Run("GetItemSetD", test_GetItemSet)
	t.Run("GetItemsBySlot", test_GetItemsBySlot)
}

func test_GetItemsByID(t *testing.T) {
	type checkFn func(t *testing.T, items []Item, expectedID []string, err error)

	apiSvr := setupItemsAPISvr()
	defer apiSvr.Close()

	verifyItemIDs := func(t *testing.T, items []Item, expectedIDs []string, err error) {
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if len(items) != len(expectedIDs) {
			t.Errorf("expected %d items, got %d", len(expectedIDs), len(items))
		}

		for i, item := range items {
			if item.Id != expectedIDs[i] {
				t.Errorf("expected name %s, got %s", expectedIDs[i], item.Id)
			}
		}
	}

	tests := map[string]func(t *testing.T) (*client, []string, checkFn){
		"success": func(t *testing.T) (*client, []string, checkFn) {
			api := NewAPI(&openapi.Configuration{
				Scheme:     "http",
				HTTPClient: http.DefaultClient,
				Servers: []openapi.ServerConfiguration{
					{
						URL: apiSvr.URL,
					},
				},
			})
			return api, []string{"2"}, verifyItemIDs
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			api, ids, checkFn := tc(t)
			items, err := api.GetItemsByID(context.Background(), ids...)
			checkFn(t, items, ids, err)
		})
	}
}

func test_GetItemsByName(t *testing.T) {
	type checkFn func(t *testing.T, items []Item, expectedNames []string, err error)

	apiSvr := setupItemsAPISvr()
	defer apiSvr.Close()

	verifyItemNames := func(t *testing.T, items []Item, expectedNames []string, err error) {
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if len(items) != len(expectedNames) {
			t.Errorf("expected %d items, got %d", len(expectedNames), len(items))
		}

		for i, item := range items {
			if item.Name != expectedNames[i] {
				t.Errorf("expected name %s, got %s", expectedNames[i], item.Name)
			}
		}
	}

	tests := map[string]func(t *testing.T) (*client, []string, checkFn){
		"success": func(t *testing.T) (*client, []string, checkFn) {
			api := NewAPI(&openapi.Configuration{
				Scheme:     "http",
				HTTPClient: http.DefaultClient,
				Servers: []openapi.ServerConfiguration{
					{
						URL: apiSvr.URL,
					},
				},
			})
			return api, []string{"Abyssal whip", "Abyssal dagger", "Rune platebody", "Dragon scimitar"}, verifyItemNames
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			api, names, checkFn := tc(t)
			items, err := api.GetItemsByName(context.Background(), names...)
			checkFn(t, items, names, err)
		})
	}
}

func test_GetItemSet(t *testing.T) {
	type checkFn func(t *testing.T, items []Item, expectedNames []string, err error)

	apiSvr := setupItemsAPISvr()
	defer apiSvr.Close()

	verifyItemNames := func(t *testing.T, items []Item, expectedNames []string, err error) {
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if len(items) != len(expectedNames) {
			t.Errorf("expected %d items, got %d", len(expectedNames), len(items))
		}

		for i, item := range items {
			if item.Name != expectedNames[i] {
				t.Errorf("expected name %s, got %s", expectedNames[i], item.Name)
			}
		}
	}

	tests := map[string]func(t *testing.T) (*client, sets.SetName, []string, checkFn){
		"success": func(t *testing.T) (*client, sets.SetName, []string, checkFn) {
			api := NewAPI(&openapi.Configuration{
				Scheme:     "http",
				HTTPClient: http.DefaultClient,
				Servers: []openapi.ServerConfiguration{
					{
						URL: apiSvr.URL,
					},
				},
			})
			return api, sets.RuneLg, []string{"Rune full helm", "Rune platebody", "Rune platelegs", "Rune kiteshield"}, verifyItemNames
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			api, setName, names, checkFn := tc(t)
			set, err := api.GetItemSet(context.Background(), setName)
			checkFn(t, set, names, err)
		})
	}
}

func test_GetItemsBySlot(t *testing.T) {
	type checkFn func(t *testing.T, items []Item, expectedNames []string, err error)

	apiSvr := setupItemsAPISvr()
	defer apiSvr.Close()

	verifyItemNames := func(t *testing.T, items []Item, expectedNames []string, err error) {
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if len(items) != len(expectedNames) {
			t.Errorf("expected %d items, got %d", len(expectedNames), len(items))
		}

		for i, item := range items {
			if item.Name != expectedNames[i] {
				t.Errorf("expected name %s, got %s", expectedNames[i], item.Name)
			}
		}
	}

	tests := map[string]func(t *testing.T) (*client, slots.SlotName, []string, checkFn){
		"success": func(t *testing.T) (*client, slots.SlotName, []string, checkFn) {
			api := NewAPI(&openapi.Configuration{
				Scheme:     "http",
				HTTPClient: http.DefaultClient,
				Servers: []openapi.ServerConfiguration{
					{
						URL: apiSvr.URL,
					},
				},
			})
			return api, slots.TwoHanded, []string{"Longbow", "Shortbow"}, verifyItemNames
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			api, slotName, names, checkFn := tc(t)
			set, err := api.GetItemsBySlot(context.Background(), slotName)
			checkFn(t, set, names, err)
		})
	}
}

func setupItemsAPISvr() *httptest.Server {
	return httptest.NewServer((http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.String() {
		case fmt.Sprintf("/items?where=%s", url.QueryEscape(`{ "wiki_name": { "$in": ["Abyssal whip", "Abyssal dagger", "Rune platebody", "Dragon scimitar"] }, "duplicate": false }`)):
			data, err := ioutil.ReadFile(filepath.Join("testdata", "items.json"))
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
			return
		case fmt.Sprintf("/items?page=2&where=%s", url.QueryEscape(`{ "wiki_name": { "$in": ["Abyssal whip", "Abyssal dagger", "Rune platebody", "Dragon scimitar"] }, "duplicate": false }`)):
			data, err := ioutil.ReadFile(filepath.Join("testdata", "items_page2.json"))
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
			return
		case fmt.Sprintf("/items?where=%s", url.QueryEscape(`{ "wiki_name": { "$in": ["Rune full helm", "Rune platebody", "Rune platelegs", "Rune kiteshield"] }, "duplicate": false }`)):
			data, err := ioutil.ReadFile(filepath.Join("testdata", "full_rune.json"))
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
			return
		case fmt.Sprintf("/items?where=%s", url.QueryEscape(`{ "wiki_name": { "$in": ["Green d'hide body", "Green d'hide chaps", "Green d'hide vambraces"] }, "duplicate": false }`)):
			data, err := ioutil.ReadFile(filepath.Join("testdata", "full_greendhide.json"))
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
			return
		case fmt.Sprintf("/items?where=%s", url.QueryEscape(`{ "equipable_by_player": true, "equipment.slot": "2h", "duplicate": false }`)):
			data, err := ioutil.ReadFile(filepath.Join("testdata", "items_2h.json"))
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
			return
		case fmt.Sprintf("/items?where=%s", url.QueryEscape(`{ "id": { "$in": ["2"] }, "duplicate": false }`)):
			data, err := ioutil.ReadFile(filepath.Join("testdata", "single_item.json"))
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
			return
		default:
			panic(fmt.Errorf("%s not supported", r.URL.String()))
		}
	})))
}

/*func TestConvert(t *testing.T) {
	var openItem []openapi.Item
	err := json.Unmarshal([]byte(testItems), &openItem)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("OPEN ITEM: %+v\n", openItem)

	item, err := OpenItemsToItems(openItem[0])
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", item)
}*/

var testItems = `
[{
	"_id": "5e865e2d28d93c7e7b2668af",
	"id": "4151",
	"name": "Abyssal whip",
	"incomplete": false,
	"members": true,
	"tradeable": true,
	"tradeable_on_ge": true,
	"stackable": false,
	"stacked": null,
	"noted": false,
	"noteable": true,
	"linked_id_item": null,
	"linked_id_noted": 4152,
	"linked_id_placeholder": 14032,
	"placeholder": false,
	"equipable": true,
	"equipable_by_player": true,
	"equipable_weapon": true,
	"cost": 120001,
	"lowalch": 48000,
	"highalch": 72000,
	"weight": 0.453,
	"buy_limit": 70,
	"quest_item": false,
	"release_date": "2005-01-26",
	"duplicate": false,
	"examine": "A weapon from the abyss.",
	"icon": "iVBORw0KGgoAAAANSUhEUgAAACQAAAAgCAYAAAB6kdqOAAABvUlEQVR4Xu2Xv26DMBDG4QEyZECKIkVCKIoyVR26dOrQpUOHDn3/V3F7nL7689kGAo6z9JNuiH1wP+6PIU3zr2Jq3bRVkQ/Y7feBvfcn93o8upfDwT11XQKwOGQMwfYx9O7rcnaf52GEezuFgNdfn4JQ7RjAQojZLAAMcPJbAOH73PcloBTIQiEAG8MB7Pt6MQ+wSW1QAs6Kh1A/NgtnHyKMcZMUCP3AIBw0XcqmgU9qb4W0JwTIwuRAYHETF4FSIMkORjn1xCnjayy8lN/vLVY7TglfvBRGTJpZrpXM+my1f6DhPRdJs8M3nAPibGDseY9hVgHxzbh3chC22dkPQ8EnufddpBgI69Lk3B8hnPrwOsPEw7FYqUzoOsqBUtps8XUASh0bYbxZxUCcJZzCNnjOBGgDDBRD8d4tQAVgRAqEs2jusLOGEhaCgTQTgApLp/s5A0RBGMg3shx2YZb0fZUz9issD4VpsR4PkJ+uYbczJQr94rW7KT3yDJcegLtqeuRlAFa+0bdoeuTxPV2538Ixt1D86Vutr8IR16CSndzfoSpQLIDh09dmscL5lJICMUClw3JKD8tGXiVgfgACr1tEhnw7UAAAAABJRU5ErkJggg==",
	"wiki_name": "Abyssal whip",
	"wiki_url": "https://oldschool.runescape.wiki/w/Abyssal_whip",
	"equipment": {
		"attack_stab": 0,
		"attack_slash": 82,
		"attack_crush": 0,
		"attack_magic": 0,
		"attack_ranged": 0,
		"defence_stab": 0,
		"defence_slash": 0,
		"defence_crush": 0,
		"defence_magic": 0,
		"defence_ranged": 0,
		"melee_strength": 82,
		"ranged_strength": 0,
		"magic_damage": 0,
		"prayer": 0,
		"slot": "weapon",
		"requirements": {
			"attack": 70
		}
	},
	"weapon": {
		"attack_speed": 4,
		"weapon_type": "whips",
		"stances": [
			{
				"combat_style": "flick",
				"attack_type": "slash",
				"attack_style": "accurate",
				"experience": "attack",
				"boosts": null
			},
			{
				"combat_style": "lash",
				"attack_type": "slash",
				"attack_style": "controlled",
				"experience": "shared",
				"boosts": null
			},
			{
				"combat_style": "deflect",
				"attack_type": "slash",
				"attack_style": "defensive",
				"experience": "defence",
				"boosts": null
			}
		]
	}
},
{
	"_id": "5e865e2f28d93c7e7b268a55",
	"id": "13265",
	"name": "Abyssal dagger",
	"incomplete": false,
	"members": true,
	"tradeable": true,
	"tradeable_on_ge": true,
	"stackable": false,
	"stacked": null,
	"noted": false,
	"noteable": true,
	"linked_id_item": null,
	"linked_id_noted": 13266,
	"linked_id_placeholder": 18629,
	"placeholder": false,
	"equipable": true,
	"equipable_by_player": true,
	"equipable_weapon": true,
	"cost": 115001,
	"lowalch": 46000,
	"highalch": 69000,
	"weight": 0.453,
	"buy_limit": 8,
	"quest_item": false,
	"release_date": "2015-10-01",
	"duplicate": false,
	"examine": "Something sharp from the body of a defeated Abyssal Sire.",
	"icon": "iVBORw0KGgoAAAANSUhEUgAAACQAAAAgCAYAAAB6kdqOAAABNElEQVR4Xu2WPQvCQAyGr2MHhw5CKRSKiKM46CAIgv/AxcXVxc2f4789+7bmcobzY7CJgw8EOvTahzdJqXN/zMn881Klf2nt3NNSknotMi6KkNCiabqST/gSaZHZXSIlMqiMFCEJZRHwXkZJBHwuI09+GV5dmpPUvKjJxIlAYp47X5VlJ9LUtaYMyEIiMpXpZKItA7hNVAZtAo+ffUihVSiTNlEKNDOrqGIheXJA+mQgtW0FaGY2NkKc0KHI/aUc+XVVhfYpC7EMCutNIjwzavOTlqGXxwIDytA/Cs+M0VqDXkCKyGTkqQF5FDFMhsi6NSYRXGPFSUjerUAvhDqNc39sV3zf1jJsljoshITOrdS1KYKQQUqZ341YiDbNdIZoduQwG8gQvGlGbUph2qY0PyXzJ+IGkmlDupvU/MQAAAAASUVORK5CYII=",
	"wiki_name": "Abyssal dagger (Unpoisoned)",
	"wiki_url": "https://oldschool.runescape.wiki/w/Abyssal_dagger#Unpoisoned",
	"equipment": {
		"attack_stab": 75,
		"attack_slash": 40,
		"attack_crush": -4,
		"attack_magic": 1,
		"attack_ranged": 0,
		"defence_stab": 0,
		"defence_slash": 0,
		"defence_crush": 0,
		"defence_magic": 1,
		"defence_ranged": 0,
		"melee_strength": 75,
		"ranged_strength": 0,
		"magic_damage": 0,
		"prayer": 0,
		"slot": "weapon",
		"requirements": {
			"attack": 70
		}
	},
	"weapon": {
		"attack_speed": 4,
		"weapon_type": "stabbing_swords",
		"stances": [
			{
				"combat_style": "stab",
				"attack_type": "stab",
				"attack_style": "accurate",
				"experience": "attack",
				"boosts": null
			},
			{
				"combat_style": "lunge",
				"attack_type": "stab",
				"attack_style": "aggressive",
				"experience": "strength",
				"boosts": null
			},
			{
				"combat_style": "slash",
				"attack_type": "slash",
				"attack_style": "aggressive",
				"experience": "strength",
				"boosts": null
			},
			{
				"combat_style": "block",
				"attack_type": "stab",
				"attack_style": "defensive",
				"experience": "defence",
				"boosts": null
			}
		]
	}
}]
`
