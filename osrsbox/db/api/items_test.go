package api

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

	"github.com/atye/gosrsbox/osrsbox/db"
)

func Test_GetItemsByName(t *testing.T) {
	type checkFn func(t *testing.T, items []db.Item, expectedNames []string, err error)

	apiSvr := setupItemsAPISvr()
	defer apiSvr.Close()

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

	tests := map[string]func(t *testing.T) (*APIClient, []string, checkFn){
		"success": func(t *testing.T) (*APIClient, []string, checkFn) {
			api := NewAPIClient()
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
		case fmt.Sprintf("/items?where=%s", url.QueryEscape(`{ "wiki_name": { "$in": ["Abyssal whip", "Abyssal dagger", "Rune platebody", "Dragon scimitar"] }, "duplicate": false }&page=2`)):
			data, err := ioutil.ReadFile(filepath.Join("testdata", "items_page2.json"))
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
