package client

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"testing"

	"github.com/atye/gosrsbox/osrsboxapi"
	"github.com/atye/gosrsbox/osrsboxapi/sets"
)

func Test_GetItemsByName(t *testing.T) {
	type checkFn func(t *testing.T, items []osrsboxapi.Item, expectedNames []string, err error)

	apiSvr := setupItemsAPISvr()
	defer apiSvr.Close()

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
			api.address = apiSvr.URL
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

func Test_GetItemSet(t *testing.T) {
	type checkFn func(t *testing.T, items []osrsboxapi.Item, expectedNames []string, err error)

	apiSvr := setupItemsAPISvr()
	defer apiSvr.Close()

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
			api.address = apiSvr.URL
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

func setupItemsAPISvr() *httptest.Server {
	ts := httptest.NewServer((http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.String() {
		case fmt.Sprintf("/items?where=%s", url.QueryEscape(`{ "wiki_name": { "$in": ["Abyssal whip", "Abyssal dagger", "Rune platebody", "Dragon scimitar"] }, "duplicate": false }`)):
			data, err := ioutil.ReadFile(filepath.Join("testdata", "items.json"))
			if err != nil {
				panic(err)
			}
			w.WriteHeader(http.StatusOK)
			w.Write(data)
			return
		case fmt.Sprintf("/items?where=%s&page=2", url.QueryEscape(`{ "wiki_name": { "$in": ["Abyssal whip", "Abyssal dagger", "Rune platebody", "Dragon scimitar"] }, "duplicate": false }`)):
			data, err := ioutil.ReadFile(filepath.Join("testdata", "items_page2.json"))
			if err != nil {
				panic(err)
			}
			w.WriteHeader(http.StatusOK)
			w.Write(data)
			return
		case fmt.Sprintf("/items?where=%s", url.QueryEscape(`{ "wiki_name": { "$in": ["Rune full helm", "Rune platebody", "Rune platelegs", "Rune kiteshield"] }, "duplicate": false }`)):
			data, err := ioutil.ReadFile(filepath.Join("testdata", "full_rune.json"))
			if err != nil {
				panic(err)
			}
			w.WriteHeader(http.StatusOK)
			w.Write(data)
			return
		case fmt.Sprintf("/items?where=%s", url.QueryEscape(`{ "wiki_name": { "$in": ["Green d'hide body", "Green d'hide chaps", "Green d'hide vambraces"] }, "duplicate": false }`)):
			data, err := ioutil.ReadFile(filepath.Join("testdata", "full_greendhide.json"))
			if err != nil {
				panic(err)
			}
			w.WriteHeader(http.StatusOK)
			w.Write(data)
			return
		default:
			panic(fmt.Errorf("%s not supported", r.URL.String()))
		}
	})))

	return ts
}

func respBodyFromTestData(path string) io.ReadCloser {
	file, err := os.Open(fmt.Sprintf("testdata/%s", path))
	if err != nil {
		panic(err)
	}
	return file
}
