package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/atye/gosrsbox/osrsboxapi"
)

func Test_GetPrayersByName(t *testing.T) {
	type checkFn func(t *testing.T, prayers []osrsboxapi.Prayer, expectedNames []string, err error)

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
			api.RunOptions(WithSource(&TestDataUpdater{}))
			err := api.UpdatePrayers()
			if err != nil {
				t.Fatal(err)
			}
			return api, []string{"Burst of Strength", "Thick Skin"}, verifyPrayerNames
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			c, names, checkFn := tc(t)
			prayers, err := c.GetPrayersByName(context.Background(), names...)
			checkFn(t, prayers, names, err)
		})
	}
}
