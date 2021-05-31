package openapi_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"testing"

	"github.com/atye/gosrsbox/internal/common"
	"github.com/atye/gosrsbox/internal/openapi"
)

func TestExecuteItemsRequest(t *testing.T) {
	t.Run("execute items", testExecuteItemsRequest)
	t.Run("execute items error", testExecuteItemsRequestError)
}

func TestExecuteMonstersRequest(t *testing.T) {
	t.Run("execute monsters", testExecuteMonstersRequest)
	t.Run("execute monsters error", testExecuteMonstersRequestError)
}

func TestExecutePrayersRequest(t *testing.T) {
	t.Run("prayers execute", testExecutePrayersRequest)
	t.Run("prayers execute error", testExecutePrayersRequestError)
}

func testExecuteItemsRequest(t *testing.T) {
	query := `{ "wiki_name": { "$in": ["Abyssal whip", "Abyssal dagger"] }, "duplicate": false }`
	ts := httptest.NewServer((http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.String() {
		case fmt.Sprintf("/items?page=1&where=%s", url.QueryEscape(query)):
			data, err := ioutil.ReadFile(filepath.Join("testdata", "items.json"))
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
			return
		default:
			t.Fatal(r.URL.Path)
		}
	})))
	defer ts.Close()

	c := openapi.NewClient("", "http", ts.URL)

	resp, err := c.ExecuteItemsRequest(context.Background(), common.Params{Where: query, Page: 1})
	if err != nil {
		t.Fatal(err)
	}

	want := []string{"Abyssal whip", "Abyssal dagger"}

	if err != nil {
		t.Errorf("expected nil error, got %+v", err)
	}

	if len(resp.GetItems()) != 2 {
		t.Errorf("expected 2 items, got %d", len(resp.GetItems()))
	}

	for i, v := range resp.GetItems() {
		if want[i] != v.Name {
			t.Errorf("expected %s, got %s", want[i], v.Name)
		}
	}

	if resp.GetTotal() != 2 {
		t.Errorf("expected total 2, got %d", resp.GetTotal())
	}

	if resp.GetMaxResults() != 2 {
		t.Errorf("expected max results 2, got %d", resp.GetMaxResults())
	}
}

func testExecuteItemsRequestError(t *testing.T) {
	query := `{foo}`
	ts := httptest.NewServer((http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"_status": "ERR", "_error": {"code": 400, "message": "The browser (or proxy) sent a request that this server could not understand."}}`))
	})))
	defer ts.Close()

	c := openapi.NewClient("", "http", ts.URL)

	_, err := c.ExecuteItemsRequest(context.Background(), common.Params{Where: query})

	if err == nil {
		t.Errorf("expected an error, got nil")
	}

	want := fmt.Errorf("code %d, message: %s", http.StatusBadRequest, "The browser (or proxy) sent a request that this server could not understand.")
	if want.Error() != err.Error() {
		t.Errorf("expected %+v, got %+v", want, err)
	}
}

func testExecuteMonstersRequest(t *testing.T) {
	query := `{ "wiki_name": { "$in": ["Molanisk", "Aberrant spectre", "Chaos Elemental"] }, "duplicate": false }`
	ts := httptest.NewServer((http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.String() {
		case fmt.Sprintf("/monsters?page=1&where=%s", url.QueryEscape(query)):
			data, err := ioutil.ReadFile(filepath.Join("testdata", "monsters.json"))
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
			return
		default:
			t.Fatal(r.URL.Path)
		}
	})))
	defer ts.Close()

	c := openapi.NewClient("", "http", ts.URL)

	resp, err := c.ExecuteMonstersRequest(context.Background(), common.Params{Where: query, Page: 1})
	if err != nil {
		t.Fatal(err)
	}

	want := []string{"Molanisk", "Aberrant spectre"}

	if err != nil {
		t.Errorf("expected nil error, got %+v", err)
	}

	if len(resp.GetMonsters()) != 2 {
		t.Errorf("expected 2 monsters, got %d", len(resp.GetMonsters()))
	}

	for i, v := range resp.GetMonsters() {
		if want[i] != v.Name {
			t.Errorf("expected %s, got %s", want[i], v.Name)
		}
	}

	if resp.GetTotal() != 2 {
		t.Errorf("expected total 2, got %d", resp.GetTotal())
	}

	if resp.GetMaxResults() != 2 {
		t.Errorf("expected max results 2, got %d", resp.GetMaxResults())
	}
}

func testExecuteMonstersRequestError(t *testing.T) {
	query := `{foo}`
	ts := httptest.NewServer((http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"_status": "ERR", "_error": {"code": 400, "message": "The browser (or proxy) sent a request that this server could not understand."}}`))
	})))
	defer ts.Close()

	c := openapi.NewClient("", "http", ts.URL)

	_, err := c.ExecuteMonstersRequest(context.Background(), common.Params{Where: query})

	if err == nil {
		t.Errorf("expected an error, got nil")
	}

	want := fmt.Errorf("code %d, message: %s", http.StatusBadRequest, "The browser (or proxy) sent a request that this server could not understand.")
	if want.Error() != err.Error() {
		t.Errorf("expected %+v, got %+v", want, err)
	}
}

func testExecutePrayersRequest(t *testing.T) {
	query := `{ "name": { "$in": ["Burst of Strength"] } }`
	ts := httptest.NewServer((http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.String() {
		case fmt.Sprintf("/prayers?page=1&where=%s", url.QueryEscape(query)):
			data, err := ioutil.ReadFile(filepath.Join("testdata", "prayers.json"))
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
			return
		default:
			t.Fatal(r.URL.Path)
		}
	})))
	defer ts.Close()

	c := openapi.NewClient("", "http", ts.URL)

	resp, err := c.ExecutePrayersRequest(context.Background(), common.Params{Where: query, Page: 1})
	if err != nil {
		t.Fatal(err)
	}

	want := []string{"Burst of Strength", "Thick Skin"}

	if err != nil {
		t.Errorf("expected nil error, got %+v", err)
	}

	if len(resp.GetPrayers()) != 1 {
		t.Errorf("expected 1 prayer, got %d", len(resp.GetPrayers()))
	}

	for i, v := range resp.GetPrayers() {
		if want[i] != v.Name {
			t.Errorf("expected %s, got %s", want[i], v.Name)
		}
	}

	if resp.GetTotal() != 1 {
		t.Errorf("expected total 1, got %d", resp.GetTotal())
	}

	if resp.GetMaxResults() != 1 {
		t.Errorf("expected max results 1, got %d", resp.GetMaxResults())
	}
}

func testExecutePrayersRequestError(t *testing.T) {
	query := `{foo}`
	ts := httptest.NewServer((http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"_status": "ERR", "_error": {"code": 400, "message": "The browser (or proxy) sent a request that this server could not understand."}}`))
	})))
	defer ts.Close()

	c := openapi.NewClient("", "http", ts.URL)

	_, err := c.ExecutePrayersRequest(context.Background(), common.Params{Where: query})

	if err == nil {
		t.Errorf("expected an error, got nil")
	}

	want := fmt.Errorf("code %d, message: %s", http.StatusBadRequest, "The browser (or proxy) sent a request that this server could not understand.")
	if want.Error() != err.Error() {
		t.Errorf("expected %+v, got %+v", want, err)
	}
}
