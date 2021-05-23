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
	"github.com/atye/gosrsbox/osrsbox"
)

func Test_Prayers(t *testing.T) {
	apiSvr := setupPrayersAPISvr()
	defer apiSvr.Close()
	t.Run("GetPrayersByID", testGetPrayersByID)
	t.Run("GetPrayersByName", testGetPrayersByName)
}

func testGetPrayersByID(t *testing.T) {
	type checkFn func(t *testing.T, prayers []osrsbox.Prayer, expectedIDs []string, err error)

	apiSvr := setupPrayersAPISvr()
	defer apiSvr.Close()

	verifyPrayerID := func(t *testing.T, prayers []osrsbox.Prayer, expectedIDs []string, err error) {
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if len(prayers) != len(expectedIDs) {
			t.Errorf("expected %d items, got %d", len(expectedIDs), len(prayers))
		}

		for i, prayer := range prayers {
			if prayer.Id != expectedIDs[i] {
				t.Errorf("expected name %s, got %s", expectedIDs[i], prayer.Name)
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
			return api, []string{"2"}, verifyPrayerID
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			api, ids, checkFn := tc(t)
			prayer, err := api.GetPrayersByID(context.Background(), ids...)
			checkFn(t, prayer, ids, err)
		})
	}
}

func testGetPrayersByName(t *testing.T) {
	type checkFn func(t *testing.T, prayers []osrsbox.Prayer, expectedNames []string, err error)

	apiSvr := setupPrayersAPISvr()
	defer apiSvr.Close()

	verifyPrayerNames := func(t *testing.T, prayers []osrsbox.Prayer, expectedNames []string, err error) {
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
			api := NewAPI(&openapi.Configuration{
				Scheme:     "http",
				HTTPClient: http.DefaultClient,
				Servers: []openapi.ServerConfiguration{
					{
						URL: apiSvr.URL,
					},
				},
			})
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
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
		case fmt.Sprintf("/prayers?page=2&where=%s", url.QueryEscape(`{ "name": { "$in": ["Burst of Strength", "Thick Skin"] } }`)):
			data, err := ioutil.ReadFile(filepath.Join("testdata", "prayers_page2.json"))
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
		case fmt.Sprintf("/prayers?where=%s", url.QueryEscape(`{ "id": { "$in": ["2"] }}`)):
			data, err := ioutil.ReadFile(filepath.Join("testdata", "single_prayer.json"))
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