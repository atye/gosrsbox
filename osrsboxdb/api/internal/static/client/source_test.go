package client

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/atye/gosrsbox/osrsboxdb"
)

func Test_HttpUpdateItems(t *testing.T) {
	type checkFn func(t *testing.T, items []osrsboxdb.Item, expectedNames []string, err error)
	verifyItemNames := func(t *testing.T, items []osrsboxdb.Item, expectedNames []string, err error) {
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if len(items) != len(expectedNames) {
			t.Errorf("expected %d items, got %d", len(expectedNames), len(items))
		}

		for _, item := range items {
			if !contains(item.Name, expectedNames) {
				t.Errorf("monster name %s doesn't exist in %s", item.Name, expectedNames)
			}
		}

	}

	tests := []struct {
		Name          string
		HttpGet       httpGet
		ExpectedNames []string
		CheckFn       checkFn
	}{
		{
			"items",
			func(string) (*http.Response, error) {
				return &http.Response{
					Body: testDataReadCloser("items.json"),
				}, nil
			},
			[]string{"Dwarf remains", "Toolkit"},
			verifyItemNames,
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			bytes, checkErr := tc.HttpGet.Items()
			var items []osrsboxdb.Item
			err := json.Unmarshal(bytes, &items)
			if err != nil {
				t.Fatal(err)
			}
			tc.CheckFn(t, items, tc.ExpectedNames, checkErr)
		})
	}
}

func Test_HttpUpdateMonsters(t *testing.T) {
	type checkFn func(t *testing.T, items []osrsboxdb.Monster, expectedNames []string, err error)
	verifyItemNames := func(t *testing.T, monsters []osrsboxdb.Monster, expectedNames []string, err error) {
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if len(monsters) != len(expectedNames) {
			t.Errorf("expected %d items, got %d", len(expectedNames), len(monsters))
		}

		for _, monster := range monsters {
			if !contains(monster.Name, expectedNames) {
				t.Errorf("monster name %s doesn't exist in %s", monster.Name, expectedNames)
			}
		}

	}

	tests := []struct {
		Name          string
		HttpGet       httpGet
		ExpectedNames []string
		CheckFn       checkFn
	}{
		{
			"monsters",
			func(string) (*http.Response, error) {
				return &http.Response{
					Body: testDataReadCloser("monsters.json"),
				}, nil
			},
			[]string{"Molanisk", "Aberrant spectre"},
			verifyItemNames,
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			bytes, checkErr := tc.HttpGet.Monsters()
			var monsters []osrsboxdb.Monster
			err := json.Unmarshal(bytes, &monsters)
			if err != nil {
				t.Fatal(err)
			}
			tc.CheckFn(t, monsters, tc.ExpectedNames, checkErr)
		})
	}
}

func Test_HttpUpdatePrayers(t *testing.T) {
	type checkFn func(t *testing.T, items []osrsboxdb.Prayer, expectedNames []string, err error)
	verifyItemNames := func(t *testing.T, prayers []osrsboxdb.Prayer, expectedNames []string, err error) {
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if len(prayers) != len(expectedNames) {
			t.Errorf("expected %d items, got %d", len(expectedNames), len(prayers))
		}

		for _, prayer := range prayers {
			if !contains(prayer.Name, expectedNames) {
				t.Errorf("monster name %s doesn't exist in %s", prayer.Name, expectedNames)
			}
		}

	}

	tests := []struct {
		Name          string
		HttpGet       httpGet
		ExpectedNames []string
		CheckFn       checkFn
	}{
		{
			"prayers",
			func(string) (*http.Response, error) {
				return &http.Response{
					Body: testDataReadCloser("prayers.json"),
				}, nil
			},
			[]string{"Thick Skin", "Burst of Strength"},
			verifyItemNames,
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			bytes, checkErr := tc.HttpGet.Prayers()
			var prayers []osrsboxdb.Prayer
			err := json.Unmarshal(bytes, &prayers)
			if err != nil {
				t.Fatal(err)
			}
			tc.CheckFn(t, prayers, tc.ExpectedNames, checkErr)
		})
	}
}

func testDataReadCloser(name string) io.ReadCloser {
	file, err := os.Open(filepath.Join("testdata", name))
	if err != nil {
		panic(err)
	}
	return file
}

func contains(value string, slice []string) bool {
	for _, str := range slice {
		if str == value {
			return true
		}
	}
	return false
}
