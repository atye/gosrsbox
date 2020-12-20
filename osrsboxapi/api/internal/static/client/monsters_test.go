package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/atye/gosrsbox/osrsboxapi"
)

func Test_GetMonstersByName(t *testing.T) {
	type checkFn func(t *testing.T, monsters []osrsboxapi.Monster, expectedNames []string, err error)

	verifyMonsterNames := func(t *testing.T, monsters []osrsboxapi.Monster, expectedNames []string, err error) {
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if len(monsters) != len(expectedNames) {
			t.Errorf("expected %d items, got %d", len(expectedNames), len(monsters))
		}

		for _, monster := range monsters {
			if !contains(monster.Name, expectedNames) {
				t.Errorf("monster name %s doesn't exist in %s", monster.Name, expectedNames)
			}
		}
	}

	tests := map[string]func(t *testing.T) (*client, []string, checkFn){
		"success": func(t *testing.T) (*client, []string, checkFn) {
			api := NewAPI(http.DefaultClient)
			api.RunOptions(WithSource(&TestDataUpdater{}))
			err := api.UpdateMonsters()
			if err != nil {
				t.Fatal(err)
			}
			return api, []string{"Molanisk", "Aberrant spectre"}, verifyMonsterNames
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			db, names, checkFn := tc(t)
			monsters, err := db.GetMonstersByName(context.Background(), names...)
			checkFn(t, monsters, names, err)
		})
	}
}

func Test_GetMonstersThatDrop(t *testing.T) {
	type checkFn func(t *testing.T, monsters []osrsboxapi.Monster, expectedNames []string, err error)

	verifyMonsterNames := func(t *testing.T, monsters []osrsboxapi.Monster, expectedNames []string, err error) {
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if len(monsters) != len(expectedNames) {
			t.Errorf("expected %d items, got %d", len(expectedNames), len(monsters))
		}

		for _, monster := range monsters {
			if !contains(monster.Name, expectedNames) {
				t.Errorf("monster name %s doesn't exist in %s", monster.Name, expectedNames)
			}
		}
	}

	tests := map[string]func(t *testing.T) (*client, []string, []string, checkFn){
		"success": func(t *testing.T) (*client, []string, []string, checkFn) {
			api := NewAPI(http.DefaultClient)
			api.RunOptions(WithSource(&TestDataUpdater{}))
			err := api.UpdateMonsters()
			if err != nil {
				t.Fatal(err)
			}
			return api, []string{"Grimy ranarr weed"}, []string{"Molanisk", "Aberrant spectre"}, verifyMonsterNames
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			db, item, names, checkFn := tc(t)
			monsters, err := db.GetMonstersThatDrop(context.Background(), item...)
			checkFn(t, monsters, names, err)
		})
	}
}
