package client

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strconv"
	"testing"

	openapi "github.com/atye/gosrsbox/pkg/openapi/api"
)

func Test_GetJSONFiles(t *testing.T) {
	type NPCSummary struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	type checkFn func(t *testing.T, summaries map[string]NPCSummary, expectedNames []string, err error)

	apiSvr := setupJsonAPISvr()
	defer apiSvr.Close()

	api := NewAPI(&openapi.Configuration{HTTPClient: http.DefaultClient})
	api.docsAddress = apiSvr.URL

	verifyNpcNames := func(t *testing.T, summaries map[string]NPCSummary, expectedNames []string, err error) {
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if len(summaries) != len(expectedNames) {
			t.Errorf("expected %d items, got %d", len(expectedNames), len(summaries))
		}

		for k, npc := range summaries {
			i, err := strconv.Atoi(k)
			if err != nil {
				t.Fatal(err)
			}
			if npc.Name != expectedNames[i] {
				t.Errorf("expected %s, got %s", expectedNames[i], npc.Name)
			}
		}
	}

	tests := map[string]func(t *testing.T) func(){
		"npcs-summary": func(t *testing.T) func() {
			return func() {
				var data map[string]NPCSummary
				err := api.GetJSONFiles(context.Background(), []string{"testdata/json-docs/npcs-summary.json"}, &data)
				if err != nil {
					t.Fatal(err)
				}

				verifyNpcNames(t, data, []string{"Tool Leprechaun", "Molanisk"}, err)
			}
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tc(t)()
		})
	}
}

func setupJsonAPISvr() *httptest.Server {
	return httptest.NewServer((http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.String() {
		case "/testdata/json-docs/npcs-summary.json":
			data, err := ioutil.ReadFile(filepath.Join("testdata", "json-docs", "npcs-summary.json"))
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
}
