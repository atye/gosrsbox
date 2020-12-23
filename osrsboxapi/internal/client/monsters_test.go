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

	openapi "github.com/atye/gosrsbox/osrsboxapi/openapi/client"
)

func Test_Monsters(t *testing.T) {
	apiSvr := setupMonstersAPISvr()
	defer apiSvr.Close()
	t.Run("GetMonstersByID", test_GetMonstersByID)
	t.Run("GetMonstersByName", test_GetMonstersByName)
	t.Run("GetMonstersThatDrop", test_GetMonstersThatDrop)
}

func test_GetMonstersByID(t *testing.T) {
	type checkFn func(t *testing.T, monsters []openapi.Monster, expectedIDs []string, err error)

	apiSvr := setupMonstersAPISvr()
	defer apiSvr.Close()

	verifyMonsterID := func(t *testing.T, monsters []openapi.Monster, expectedIDs []string, err error) {
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if len(monsters) != len(expectedIDs) {
			t.Errorf("expected %d items, got %d", len(expectedIDs), len(monsters))
		}

		for i, monster := range monsters {
			if monster.GetId() != expectedIDs[i] {
				t.Errorf("expected name %s, got %s", expectedIDs[i], monster.GetName())
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
			return api, []string{"2"}, verifyMonsterID
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			api, ids, checkFn := tc(t)
			monster, err := api.GetMonstersByID(context.Background(), ids...)
			checkFn(t, monster, ids, err)
		})
	}
}

func test_GetMonstersByName(t *testing.T) {
	type checkFn func(t *testing.T, monsters []openapi.Monster, expectedNames []string, err error)

	apiSvr := setupMonstersAPISvr()
	defer apiSvr.Close()

	verifyMonsterNames := func(t *testing.T, monsters []openapi.Monster, expectedNames []string, err error) {
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
			api := NewAPI(&openapi.Configuration{
				Scheme:     "http",
				HTTPClient: http.DefaultClient,
				Servers: []openapi.ServerConfiguration{
					{
						URL: apiSvr.URL,
					},
				},
			})
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

func test_GetMonstersThatDrop(t *testing.T) {
	type checkFn func(t *testing.T, monsters []openapi.Monster, expectedNames []string, err error)

	apiSvr := setupMonstersAPISvr()
	defer apiSvr.Close()

	verifyMonsterNames := func(t *testing.T, monsters []openapi.Monster, expectedNames []string, err error) {
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
			api := NewAPI(&openapi.Configuration{
				Scheme:     "http",
				HTTPClient: http.DefaultClient,
				Servers: []openapi.ServerConfiguration{
					{
						URL: apiSvr.URL,
					},
				},
			})
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
	return httptest.NewServer((http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.String() {
		case fmt.Sprintf("/monsters?where=%s", url.QueryEscape(`{ "wiki_name": { "$in": ["Molanisk", "Aberrant spectre", "Chaos Elemental"] }, "duplicate": false }`)):
			data, err := ioutil.ReadFile(filepath.Join("testdata", "monsters.json"))
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
		case fmt.Sprintf("/monsters?page=2&where=%s", url.QueryEscape(`{ "wiki_name": { "$in": ["Molanisk", "Aberrant spectre", "Chaos Elemental"] }, "duplicate": false }`)):
			data, err := ioutil.ReadFile(filepath.Join("testdata", "monsters_page2.json"))
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
		case fmt.Sprintf("/monsters?where=%s", url.QueryEscape(`{ "drops": { "$elemMatch": { "name": { "$in": ["Grimy ranarr weed"] } } }, "duplicate": false }`)):
			data, err := ioutil.ReadFile(filepath.Join("testdata", "monsters_onepage.json"))
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
		case fmt.Sprintf("/monsters?where=%s", url.QueryEscape(`{ "id": { "$in": ["2"] }, "duplicate": false }`)):
			data, err := ioutil.ReadFile(filepath.Join("testdata", "single_monster.json"))
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
