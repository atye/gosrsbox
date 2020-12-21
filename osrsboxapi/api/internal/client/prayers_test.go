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

func Test_GetPrayersByName(t *testing.T) {
	type checkFn func(t *testing.T, prayers []osrsboxapi.Prayer, expectedNames []string, err error)

	apiSvr := setupPrayersAPISvr()
	defer apiSvr.Close()

	verifyPrayerNames := func(t *testing.T, prayers []osrsboxapi.Prayer, expectedNames []string, err error) {
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if len(prayers) != len(expectedNames) {
			t.Errorf("expected %d items, got %d", len(expectedNames), len(prayers))
		}

		for i, prayer := range prayers {
			if prayer.Name != expectedNames[i] {
				t.Errorf("expected name %s, got %s", expectedNames[i], prayer.Name)
			}
		}
	}

	tests := map[string]func(t *testing.T) (*client, []string, checkFn){
		"success": func(t *testing.T) (*client, []string, checkFn) {
			api := NewAPI(http.DefaultClient)
			api.apiAddress = apiSvr.URL
			return api, []string{"Burst of Strength", "Thick Skin"}, verifyPrayerNames
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			api, names, checkFn := tc(t)
			prayers, err := api.GetPrayersByName(context.Background(), names...)
			checkFn(t, prayers, names, err)
		})
	}
}

func setupPrayersAPISvr() *httptest.Server {
	return httptest.NewServer((http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.String() {
		case fmt.Sprintf("/prayers?where=%s", url.QueryEscape(`{ "name": { "$in": ["Burst of Strength", "Thick Skin"] } }`)):
			data, err := ioutil.ReadFile(filepath.Join("testdata", "prayers.json"))
			if err != nil {
				panic(err)
			}
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		case fmt.Sprintf("/prayers?where=%s&page=2", url.QueryEscape(`{ "name": { "$in": ["Burst of Strength", "Thick Skin"] } }`)):
			data, err := ioutil.ReadFile(filepath.Join("testdata", "prayers_page2.json"))
			if err != nil {
				panic(err)
			}
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		default:
			panic(fmt.Errorf("%s not supported", r.URL.String()))
		}
	})))
}
