package inmemory

import (
	"context"
	"testing"

	"github.com/atye/gosrsbox/osrsbox/db"
)

func Test_GetMonstersByName(t *testing.T) {
	type checkFn func(t *testing.T, monsters []db.Monster, expectedNames []string, err error)

	verifyMonsterNames := func(t *testing.T, monsters []db.Monster, expectedNames []string, err error) {
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if len(monsters) != len(expectedNames) {
			t.Errorf("expected %d items, got %d", len(expectedNames), len(monsters))
		}

		if monsters[0].Name != expectedNames[0] {
			t.Errorf("expected name %s, got %s", monsters[0].Name, expectedNames[0])
		}

		if monsters[1].Name != expectedNames[1] {
			t.Errorf("expected name %s, got %s", monsters[1].Name, expectedNames[1])
		}
	}

	tests := map[string]func(t *testing.T) (*InMemoryClient, []string, checkFn){
		"success": func(t *testing.T) (*InMemoryClient, []string, checkFn) {
			c, err := NewInMemoryClient(WithUpdater(&TestDataUpdater{}))
			if err != nil {
				t.Fatal(err)
			}
			err = c.UpdateMonsters()
			if err != nil {
				t.Fatal(err)
			}
			return c, []string{"Molanisk", "Aberrant spectre"}, verifyMonsterNames
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
	type checkFn func(t *testing.T, monsters []db.Monster, expectedNames []string, err error)

	verifyMonsterNames := func(t *testing.T, monsters []db.Monster, expectedNames []string, err error) {
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if len(monsters) != len(expectedNames) {
			t.Errorf("expected %d items, got %d", len(expectedNames), len(monsters))
		}

		for i, monster := range monsters {
			if monster.Name != expectedNames[i] {
				t.Errorf("expected name %s, got %s", expectedNames[i], monster.Name)
			}
		}
	}

	tests := map[string]func(t *testing.T) (*InMemoryClient, []string, []string, checkFn){
		"success": func(t *testing.T) (*InMemoryClient, []string, []string, checkFn) {
			db, err := NewInMemoryClient(WithUpdater(&TestDataUpdater{}))
			if err != nil {
				t.Fatal(err)
			}
			err = db.UpdateMonsters()
			if err != nil {
				t.Fatal(err)
			}
			return db, []string{"Grimy ranarr weed"}, []string{"Molanisk", "Aberrant spectre"}, verifyMonsterNames
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
