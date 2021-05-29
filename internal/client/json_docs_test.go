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

	"github.com/atye/gosrsbox/internal/api"
	"golang.org/x/sync/semaphore"
)

func TestGetDocument(t *testing.T) {
	t.Run("testGetDocument", testGetDocument)
	t.Run("testGetDocumentError", testGetDocumentError)

}

func testGetDocument(t *testing.T) {
	type NPCSummary struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	apiSvr := setupJsonAPISvr()
	defer apiSvr.Close()

	api := &apiClient{
		docsAddress:   apiSvr.URL,
		openAPIClient: api.NewAPIClient(&api.Configuration{HTTPClient: http.DefaultClient}),
		sem:           semaphore.NewWeighted(int64(10)),
	}

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
				err := api.GetDocument(context.Background(), "testdata/json-docs/npcs-summary.json", &data)
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

func testGetDocumentError(t *testing.T) {
	apiSvr := httptest.NewServer((http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`File not found`))
	})))
	defer apiSvr.Close()

	api := &apiClient{
		docsAddress:   apiSvr.URL,
		openAPIClient: api.NewAPIClient(&api.Configuration{HTTPClient: http.DefaultClient}),
		sem:           semaphore.NewWeighted(int64(10)),
	}

	err := api.GetDocument(context.Background(), "test", new(map[string]interface{}))

	if err == nil {
		t.Errorf("expected non-nil error")
	}

	want := fmt.Errorf("code: %d, message: %s", http.StatusNotFound, http.StatusText(http.StatusNotFound))
	if want.Error() != err.Error() {
		t.Errorf("expected %+v, got %+v", want, err)
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
