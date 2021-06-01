package client

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/atye/gosrsbox/internal/client/common"
	"github.com/atye/gosrsbox/internal/client/common/mocks"
	"github.com/atye/gosrsbox/models"
	"github.com/atye/gosrsbox/sets"
	"github.com/atye/gosrsbox/slots"
	"github.com/golang/mock/gomock"
)

func TestItems(t *testing.T) {
	t.Run("GetItemsByID", testGetItemsByID)
	t.Run("GetItemsByName", testGetItemsByName)
	t.Run("GetItemSet", testGetItemSet)
	t.Run("GetItemsBySlot", testGetItemsBySlot)
	t.Run("GetItemsError", testGetItemsAPIError)
}

func testGetItemsByID(t *testing.T) {
	type checkFn func(t *testing.T, items []models.Item, expectedID []string, err error)

	verifyItemIDs := func(t *testing.T, items []models.Item, expectedIDs []string, err error) {
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if len(items) != len(expectedIDs) {
			t.Errorf("expected %d items, got %d", len(expectedIDs), len(items))
		}

		for i, item := range items {
			if item.Id != expectedIDs[i] {
				t.Errorf("expected name %s, got %s", expectedIDs[i], item.Id)
			}
		}
	}

	verifyError := func(t *testing.T, items []models.Item, expectedIDs []string, err error) {
		if err == nil {
			t.Errorf("expected an error")
		}
	}

	tests := map[string]func(t *testing.T) (*APIClient, []string, checkFn){
		"success": func(t *testing.T) (*APIClient, []string, checkFn) {
			ctrl := gomock.NewController(t)

			inline := mocks.NewMockItemsResponse(ctrl)
			inline.EXPECT().GetTotal().Return(1)
			inline.EXPECT().GetMaxResults().Return(25)
			inline.EXPECT().GetItems().Return([]models.Item{
				{
					Id: "2",
				},
			})

			executor := mocks.NewMockRequestExecutor(ctrl)
			executor.EXPECT().ExecuteItemsRequest(gomock.Any(), common.Params{Where: `{ "id": { "$in": ["2"] }, "duplicate": false }`}).Return(inline, nil)

			api := NewAPI("")
			api.reqExecutor = executor

			return api, []string{"2"}, verifyItemIDs
		},
		"no IDs": func(t *testing.T) (*APIClient, []string, checkFn) {
			api := NewAPI("")
			return api, []string{}, verifyError
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			api, ids, checkFn := tc(t)
			items, err := api.GetItemsByID(context.Background(), ids...)
			checkFn(t, items, ids, err)
		})
	}
}

func testGetItemsByName(t *testing.T) {
	type checkFn func(t *testing.T, items []models.Item, expectedNames []string, err error)

	verifyItemNames := func(t *testing.T, items []models.Item, expectedNames []string, err error) {
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

	verifyError := func(t *testing.T, items []models.Item, expectedIDs []string, err error) {
		if err == nil {
			t.Errorf("expected an error")
		}
	}

	tests := map[string]func(t *testing.T) (*APIClient, []string, checkFn){
		"success": func(t *testing.T) (*APIClient, []string, checkFn) {
			ctrl := gomock.NewController(t)

			inline := mocks.NewMockItemsResponse(ctrl)
			inline.EXPECT().GetTotal().Return(4)
			inline.EXPECT().GetMaxResults().Return(2)
			inline.EXPECT().GetItems().Return([]models.Item{
				{
					Name: "Abyssal whip",
				}, {
					Name: "Abyssal dagger",
				},
			})

			executor := mocks.NewMockRequestExecutor(ctrl)
			query := `{ "wiki_name": { "$in": ["Abyssal whip", "Abyssal dagger", "Rune platebody", "Dragon scimitar"] }, "duplicate": false }`
			executor.EXPECT().ExecuteItemsRequest(gomock.Any(), common.Params{Where: query}).Return(inline, nil)

			inlinePage2 := mocks.NewMockItemsResponse(ctrl)
			inlinePage2.EXPECT().GetItems().Return([]models.Item{
				{
					Name: "Rune platebody",
				},
				{
					Name: "Dragon scimitar",
				},
			})
			executor.EXPECT().ExecuteItemsRequest(gomock.Any(), common.Params{Where: query, Page: 2}).Return(inlinePage2, nil)

			api := NewAPI("")
			api.reqExecutor = executor

			return api, []string{"Abyssal whip", "Abyssal dagger", "Rune platebody", "Dragon scimitar"}, verifyItemNames
		},
		"no names": func(t *testing.T) (*APIClient, []string, checkFn) {
			api := NewAPI("")
			return api, []string{}, verifyError
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

func testGetItemSet(t *testing.T) {
	type checkFn func(t *testing.T, items []models.Item, expectedNames []string, err error)

	verifyItemNames := func(t *testing.T, items []models.Item, expectedNames []string, err error) {
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

	verifyError := func(t *testing.T, items []models.Item, expectedIDs []string, err error) {
		if err == nil {
			t.Errorf("expected an error")
		}
	}

	tests := map[string]func(t *testing.T) (*APIClient, sets.SetName, []string, checkFn){
		"success": func(t *testing.T) (*APIClient, sets.SetName, []string, checkFn) {
			ctrl := gomock.NewController(t)

			inline := mocks.NewMockItemsResponse(ctrl)
			inline.EXPECT().GetTotal().Return(4)
			inline.EXPECT().GetMaxResults().Return(25)
			inline.EXPECT().GetItems().Return([]models.Item{
				{
					Name: "Rune full helm",
				},
				{
					Name: "Rune platebody",
				},
				{
					Name: "Rune platelegs",
				},
				{
					Name: "Rune kiteshield",
				},
			})

			executor := mocks.NewMockRequestExecutor(ctrl)
			query := `{ "wiki_name": { "$in": ["Rune full helm", "Rune platebody", "Rune platelegs", "Rune kiteshield"] }, "duplicate": false }`
			executor.EXPECT().ExecuteItemsRequest(gomock.Any(), common.Params{Where: query}).Return(inline, nil)

			api := NewAPI("")
			api.reqExecutor = executor

			return api, sets.RuneLg, []string{"Rune full helm", "Rune platebody", "Rune platelegs", "Rune kiteshield"}, verifyItemNames
		},
		"no set": func(t *testing.T) (*APIClient, sets.SetName, []string, checkFn) {
			api := NewAPI("")
			return api, nil, nil, verifyError
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			api, setName, names, checkFn := tc(t)
			set, err := api.GetItemSet(context.Background(), setName)
			checkFn(t, set, names, err)
		})
	}
}

func testGetItemsBySlot(t *testing.T) {
	type checkFn func(t *testing.T, items []models.Item, expectedNames []string, err error)

	verifyItemNames := func(t *testing.T, items []models.Item, expectedNames []string, err error) {
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

	verifyError := func(t *testing.T, items []models.Item, expectedIDs []string, err error) {
		if err == nil {
			t.Errorf("expected an error")
		}
	}

	tests := map[string]func(t *testing.T) (*APIClient, slots.SlotName, []string, checkFn){
		"success": func(t *testing.T) (*APIClient, slots.SlotName, []string, checkFn) {
			ctrl := gomock.NewController(t)

			inline := mocks.NewMockItemsResponse(ctrl)
			inline.EXPECT().GetTotal().Return(2)
			inline.EXPECT().GetMaxResults().Return(25)
			inline.EXPECT().GetItems().Return([]models.Item{
				{
					Name: "Longbow",
				},
				{
					Name: "Shortbow",
				},
			})

			executor := mocks.NewMockRequestExecutor(ctrl)
			query := `{ "equipable_by_player": true, "equipment.slot": "2h", "duplicate": false }`
			executor.EXPECT().ExecuteItemsRequest(gomock.Any(), common.Params{Where: query}).Return(inline, nil)

			api := NewAPI("")
			api.reqExecutor = executor

			return api, slots.TwoHanded, []string{"Longbow", "Shortbow"}, verifyItemNames
		},
		"no slot": func(t *testing.T) (*APIClient, slots.SlotName, []string, checkFn) {
			api := NewAPI("")
			return api, "", nil, verifyError
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			api, slotName, names, checkFn := tc(t)
			set, err := api.GetItemsBySlot(context.Background(), slotName)
			checkFn(t, set, names, err)
		})
	}
}

func testGetItemsAPIError(t *testing.T) {
	apiSvr := httptest.NewServer((http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"_status": "ERR", "_error": {"code": 400, "message": "The browser (or proxy) sent a request that this server could not understand."}}`))
	})))
	defer apiSvr.Close()

	api := NewAPI("")

	_, err := api.GetItemsByQuery(context.Background(), `{test}`)

	if err == nil {
		t.Errorf("expected non-nil error")
	}

	want := fmt.Errorf("code %d, message: %s", http.StatusBadRequest, "The browser (or proxy) sent a request that this server could not understand.")
	if want.Error() != err.Error() {
		t.Errorf("expected %+v, got %+v", want, err)
	}
}
