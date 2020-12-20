package client

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"testing"

	"github.com/atye/gosrsbox/osrsboxapi"
)

func Test_GetMonstersByName(t *testing.T) {
	type checkFn func(t *testing.T, monsters []osrsboxapi.Monster, expectedNames []string, err error)

	apiSvr := setupMonstersAPISvr()
	defer apiSvr.Close()

	verifyMonsterNames := func(t *testing.T, monsters []osrsboxapi.Monster, expectedNames []string, err error) {
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

	tests := map[string]func(t *testing.T) (*client, []string, checkFn){
		"success": func(t *testing.T) (*client, []string, checkFn) {
			api := NewAPI(http.DefaultClient)
			api.apiAddress = apiSvr.URL
			return api, []string{"Molanisk", "Aberrant spectre", "Chaos Elemental"}, verifyMonsterNames
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			api, names, checkFn := tc(t)
			monsters, err := api.GetMonstersByName(context.Background(), names...)
			checkFn(t, monsters, names, err)
		})
	}
}

func Test_GetMonstersThatDrop(t *testing.T) {
	type checkFn func(t *testing.T, monsters []osrsboxapi.Monster, expectedNames []string, err error)

	apiSvr := setupMonstersAPISvr()
	defer apiSvr.Close()

	verifyMonsterNames := func(t *testing.T, monsters []osrsboxapi.Monster, expectedNames []string, err error) {
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

	tests := map[string]func(t *testing.T) (*client, []string, []string, checkFn){
		"success": func(t *testing.T) (*client, []string, []string, checkFn) {
			api := NewAPI(http.DefaultClient)
			api.apiAddress = apiSvr.URL
			return api, []string{"Grimy ranarr weed"}, []string{"Molanisk", "Aberrant spectre"}, verifyMonsterNames
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			api, item, names, checkFn := tc(t)
			monsters, err := api.GetMonstersThatDrop(context.Background(), item...)
			checkFn(t, monsters, names, err)
		})
	}
}

func setupMonstersAPISvr() *httptest.Server {
	ts := httptest.NewServer((http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.String() {
		case fmt.Sprintf("/monsters?where=%s", url.QueryEscape(`{ "wiki_name": { "$in": ["Molanisk", "Aberrant spectre", "Chaos Elemental"] }, "duplicate": false }`)):
			data, err := ioutil.ReadFile(filepath.Join("testdata", "monsters.json"))
			if err != nil {
				panic(err)
			}
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		case fmt.Sprintf("/monsters?where=%s&page=2", url.QueryEscape(`{ "wiki_name": { "$in": ["Molanisk", "Aberrant spectre", "Chaos Elemental"] }, "duplicate": false }`)):
			data, err := ioutil.ReadFile(filepath.Join("testdata", "monsters_page2.json"))
			if err != nil {
				panic(err)
			}
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		case fmt.Sprintf("/monsters?where=%s", url.QueryEscape(`{ "drops": { "$elemMatch": { "name": { "$in": ["Grimy ranarr weed"] } } }, "duplicate": false }`)):
			data, err := ioutil.ReadFile(filepath.Join("testdata", "monsters_onepage.json"))
			if err != nil {
				panic(err)
			}
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		default:
			panic(fmt.Errorf("%s not supported", r.URL.String()))
		}
	})))

	return ts
}
