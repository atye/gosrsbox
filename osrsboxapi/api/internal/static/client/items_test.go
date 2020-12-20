package client

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"testing"

	"github.com/atye/gosrsbox/osrsboxapi"
	"github.com/atye/gosrsbox/osrsboxapi/sets"
)

func Test_GetItemsByName(t *testing.T) {
	type checkFn func(t *testing.T, items []osrsboxapi.Item, expectedNames []string, err error)

	verifyItemNames := func(t *testing.T, items []osrsboxapi.Item, expectedNames []string, err error) {
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
			api := NewAPI(http.DefaultClient)
			api.RunOptions(WithSource(&TestDataUpdater{ItemsFile: "items.json"}))
			err := api.UpdateItems()
			if err != nil {
				t.Fatal(err)
			}
			return api, []string{"Toolkit", "Dwarf remains"}, verifyItemNames
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			db, names, checkFn := tc(t)
			items, err := db.GetItemsByName(context.Background(), names...)
			checkFn(t, items, names, err)
		})
	}
}

func Test_GetItemSet(t *testing.T) {
	type checkFn func(t *testing.T, items []osrsboxapi.Item, expectedNames []string, err error)

	verifyItemNames := func(t *testing.T, items []osrsboxapi.Item, expectedNames []string, err error) {
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
			api := NewAPI(http.DefaultClient)
			api.RunOptions(WithSource(&TestDataUpdater{ItemsFile: "full_rune.json"}))
			err := api.UpdateItems()
			if err != nil {
				t.Fatal(err)
			}
			return api, sets.RuneLg, []string{"Rune full helm", "Rune platebody", "Rune platelegs", "Rune kiteshield"}, verifyItemNames
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			db, setName, names, checkFn := tc(t)
			set, err := db.GetItemSet(context.Background(), setName)
			checkFn(t, set, names, err)
		})
	}
}

type TestDataUpdater struct {
	ItemsFile string
}

func (t *TestDataUpdater) Items() ([]byte, error) {
	file, err := ioutil.ReadFile(filepath.Join("testdata", t.ItemsFile))
	if err != nil {
		return nil, err
	}

	var itemMap map[string]osrsboxapi.Item
	err = json.Unmarshal(file, &itemMap)
	if err != nil {
		return nil, err
	}

	itemSlice := make([]osrsboxapi.Item, 0, len(itemMap))
	for _, item := range itemMap {
		itemSlice = append(itemSlice, item)
	}

	items, err := json.Marshal(itemSlice)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (t *TestDataUpdater) Monsters() ([]byte, error) {
	file, err := ioutil.ReadFile(filepath.Join("testdata", "monsters.json"))
	if err != nil {
		return nil, err
	}

	var monstersMap map[string]osrsboxapi.Monster
	err = json.Unmarshal(file, &monstersMap)
	if err != nil {
		return nil, err
	}

	monstersSlice := make([]osrsboxapi.Monster, 0, len(monstersMap))
	for _, monster := range monstersMap {
		monstersSlice = append(monstersSlice, monster)

	}

	monsters, err := json.Marshal(monstersSlice)
	if err != nil {
		return nil, err
	}
	return monsters, nil
}

func (t *TestDataUpdater) Prayers() ([]byte, error) {
	file, err := ioutil.ReadFile(filepath.Join("testdata", "prayers.json"))
	if err != nil {
		return nil, err
	}
	var prayersMap map[string]osrsboxapi.Prayer
	err = json.Unmarshal(file, &prayersMap)
	if err != nil {
		return nil, err
	}

	prayersSlice := make([]osrsboxapi.Prayer, 0, len(prayersMap))
	for _, prayer := range prayersMap {
		prayersSlice = append(prayersSlice, prayer)
	}

	prayers, err := json.Marshal(prayersSlice)
	if err != nil {
		return nil, err
	}
	return prayers, nil
}
