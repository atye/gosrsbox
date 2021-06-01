package client

import (
	"context"
	"testing"

	"github.com/atye/gosrsbox/internal/client/common"
	"github.com/atye/gosrsbox/internal/client/common/mocks"
	"github.com/golang/mock/gomock"

	"github.com/atye/gosrsbox/models"
)

func TestPrayers(t *testing.T) {
	t.Run("GetPrayersByID", testGetPrayersByID)
	t.Run("GetPrayersByName", testGetPrayersByName)
}

func testGetPrayersByID(t *testing.T) {
	type checkFn func(t *testing.T, prayers []models.Prayer, expectedIDs []string, err error)

	verifyPrayerID := func(t *testing.T, prayers []models.Prayer, expectedIDs []string, err error) {
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

	verifyError := func(t *testing.T, prayers []models.Prayer, expectedIDs []string, err error) {
		if err == nil {
			t.Errorf("expected an error")
		}
	}

	tests := map[string]func(t *testing.T) (*APIClient, []string, checkFn){
		"success": func(t *testing.T) (*APIClient, []string, checkFn) {
			ctrl := gomock.NewController(t)

			inline := mocks.NewMockPrayersResponse(ctrl)
			inline.EXPECT().GetTotal().Return(1)
			inline.EXPECT().GetMaxResults().Return(1)
			inline.EXPECT().GetPrayers().Return([]models.Prayer{
				{
					Id: "2",
				},
			})

			executor := mocks.NewMockRequestExecutor(ctrl)
			executor.EXPECT().ExecutePrayersRequest(gomock.Any(), common.Params{Where: `{ "id": { "$in": ["2"] }}`}).Return(inline, nil)

			api := NewAPI("")
			api.reqExecutor = executor

			return api, []string{"2"}, verifyPrayerID
		},
		"no IDs": func(t *testing.T) (*APIClient, []string, checkFn) {
			api := NewAPI("")
			return api, []string{}, verifyError
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
	type checkFn func(t *testing.T, prayers []models.Prayer, expectedNames []string, err error)

	verifyPrayerNames := func(t *testing.T, prayers []models.Prayer, expectedNames []string, err error) {
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

	verifyError := func(t *testing.T, prayers []models.Prayer, expectedIDs []string, err error) {
		if err == nil {
			t.Errorf("expected an error")
		}
	}

	tests := map[string]func(t *testing.T) (*APIClient, []string, checkFn){
		"success": func(t *testing.T) (*APIClient, []string, checkFn) {
			ctrl := gomock.NewController(t)

			inline := mocks.NewMockPrayersResponse(ctrl)
			inline.EXPECT().GetTotal().Return(2)
			inline.EXPECT().GetMaxResults().Return(1)
			inline.EXPECT().GetPrayers().Return([]models.Prayer{
				{
					Name: "Burst of Strength",
				},
			})

			executor := mocks.NewMockRequestExecutor(ctrl)
			query := `{ "name": { "$in": ["Burst of Strength", "Thick Skin"] } }`
			executor.EXPECT().ExecutePrayersRequest(gomock.Any(), common.Params{Where: query}).Return(inline, nil)

			inlinePage2 := mocks.NewMockPrayersResponse(ctrl)
			inlinePage2.EXPECT().GetPrayers().Return([]models.Prayer{
				{
					Name: "Thick Skin",
				},
			})
			executor.EXPECT().ExecutePrayersRequest(gomock.Any(), common.Params{Where: query, Page: 2}).Return(inlinePage2, nil)

			api := NewAPI("")
			api.reqExecutor = executor

			return api, []string{"Burst of Strength", "Thick Skin"}, verifyPrayerNames
		},
		"no names": func(t *testing.T) (*APIClient, []string, checkFn) {
			api := NewAPI("")
			return api, []string{}, verifyError
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
