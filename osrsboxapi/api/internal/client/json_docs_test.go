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

	"github.com/atye/gosrsbox/osrsboxapi"
)

func Test_GetJSONFiles(t *testing.T) {
	type checkFn func(t *testing.T, summaries map[string]osrsboxapi.NPCSummary, expectedNames []string, err error)

	apiSvr := setupJsonAPISvr()
	defer apiSvr.Close()

	api := NewAPI(http.DefaultClient)
	api.docsAddress = apiSvr.URL

	verifyNpcNames := func(t *testing.T, summaries map[string]osrsboxapi.NPCSummary, expectedNames []string, err error) {
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
				var npcs map[string]osrsboxapi.NPCSummary
				err := api.GetJSONFiles(context.Background(), []string{"testdata/json-docs/npcs-summary.json"}, &npcs)
				if err != nil {
					t.Fatal(err)
				}

				verifyNpcNames(t, npcs, []string{"Tool Leprechaun", "Molanisk"}, err)
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
