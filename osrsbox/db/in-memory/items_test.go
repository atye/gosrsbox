package inmemory

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/atye/gosrsbox/osrsbox/db"
)

func Test_GetItemsByName(t *testing.T) {
	type checkFn func(t *testing.T, items []db.Item, expectedNames []string, err error)

	verifyItemNames := func(t *testing.T, items []db.Item, expectedNames []string, err error) {
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

	tests := map[string]func(t *testing.T) (*InMemoryClient, []string, checkFn){
		"success": func(t *testing.T) (*InMemoryClient, []string, checkFn) {
			db, err := NewInMemoryClient(WithUpdater(&TestDataUpdater{}))
			if err != nil {
				t.Fatal(err)
			}
			err = db.UpdateItems()
			if err != nil {
				t.Fatal(err)
			}
			return db, []string{"Toolkit", "Dwarf remains"}, verifyItemNames
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

type TestDataUpdater struct{}

func (t *TestDataUpdater) Items() ([]byte, error) {
	file, err := ioutil.ReadFile(filepath.Join("testdata", "items.json"))
	if err != nil {
		return nil, err
	}

	var itemMap map[string]db.Item
	err = json.Unmarshal(file, &itemMap)
	if err != nil {
		return nil, err
	}

	itemSlice := make([]db.Item, 0, len(itemMap))
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

	var monstersMap map[string]db.Monster
	err = json.Unmarshal(file, &monstersMap)
	if err != nil {
		return nil, err
	}

	monstersSlice := make([]db.Monster, 0, len(monstersMap))
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
	var prayersMap map[string]db.Prayer
	err = json.Unmarshal(file, &prayersMap)
	if err != nil {
		return nil, err
	}

	prayersSlice := make([]db.Prayer, 0, len(prayersMap))
	for _, prayer := range prayersMap {
		prayersSlice = append(prayersSlice, prayer)
	}

	prayers, err := json.Marshal(prayersSlice)
	if err != nil {
		return nil, err
	}
	return prayers, nil
}
