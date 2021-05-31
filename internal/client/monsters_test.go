package client

import (
	"context"
	"testing"

	"github.com/atye/gosrsbox/internal/common"
	"github.com/atye/gosrsbox/internal/common/mocks"
	"github.com/golang/mock/gomock"

	"github.com/atye/gosrsbox/models"
)

func TestMonsters(t *testing.T) {
	t.Run("GetMonstersByID", testGetMonstersByID)
	t.Run("GetMonstersByName", testGetMonstersByName)
	t.Run("GetMonstersThatDrop", testGetMonstersThatDrop)
}

func testGetMonstersByID(t *testing.T) {
	type checkFn func(t *testing.T, monsters []models.Monster, expectedIDs []string, err error)

	verifyMonsterID := func(t *testing.T, monsters []models.Monster, expectedIDs []string, err error) {
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if len(monsters) != len(expectedIDs) {
			t.Errorf("expected %d items, got %d", len(expectedIDs), len(monsters))
		}

		for i, monster := range monsters {
			if monster.Id != expectedIDs[i] {
				t.Errorf("expected name %s, got %s", expectedIDs[i], monster.Name)
			}
		}
	}

	verifyError := func(t *testing.T, monsters []models.Monster, expectedIDs []string, err error) {
		if err == nil {
			t.Errorf("expected an error")
		}
	}

	tests := map[string]func(t *testing.T) (*APIClient, []string, checkFn){
		"success": func(t *testing.T) (*APIClient, []string, checkFn) {
			ctrl := gomock.NewController(t)

			inline := mocks.NewMockMonstersResponse(ctrl)
			inline.EXPECT().GetTotal().Return(1)
			inline.EXPECT().GetMaxResults().Return(25)
			inline.EXPECT().GetMonsters().Return([]models.Monster{
				{
					Id: "2",
				},
			})

			executor := mocks.NewMockRequestExecutor(ctrl)
			executor.EXPECT().ExecuteMonstersRequest(gomock.Any(), common.Params{Where: `{ "id": { "$in": ["2"] }, "duplicate": false }`}).Return(inline, nil)

			api := NewAPI("")
			api.reqExecutor = executor

			return api, []string{"2"}, verifyMonsterID
		},
		"no IDs": func(t *testing.T) (*APIClient, []string, checkFn) {
			api := NewAPI("")
			return api, []string{}, verifyError
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			api, ids, checkFn := tc(t)
			monster, err := api.GetMonstersByID(context.Background(), ids...)
			checkFn(t, monster, ids, err)
		})
	}
}

func testGetMonstersByName(t *testing.T) {
	type checkFn func(t *testing.T, monsters []models.Monster, expectedNames []string, err error)

	verifyMonsterNames := func(t *testing.T, monsters []models.Monster, expectedNames []string, err error) {
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if len(monsters) != len(expectedNames) {
			t.Errorf("expected %d items, got %d", len(expectedNames), len(monsters))
		}

		for i, monster := range monsters {
			if monster.Name != expectedNames[i] {
				t.Errorf("expected name %s, got %s", expectedNames[i], monster.Name)
			}
		}
	}

	verifyError := func(t *testing.T, monsters []models.Monster, expectedIDs []string, err error) {
		if err == nil {
			t.Errorf("expected an error")
		}
	}

	tests := map[string]func(t *testing.T) (*APIClient, []string, checkFn){
		"success": func(t *testing.T) (*APIClient, []string, checkFn) {
			ctrl := gomock.NewController(t)

			inline := mocks.NewMockMonstersResponse(ctrl)
			inline.EXPECT().GetTotal().Return(3)
			inline.EXPECT().GetMaxResults().Return(2)
			inline.EXPECT().GetMonsters().Return([]models.Monster{
				{
					Name: "Molanisk",
				}, {
					Name: "Aberrant spectre",
				},
			})

			executor := mocks.NewMockRequestExecutor(ctrl)
			query := `{ "wiki_name": { "$in": ["Molanisk", "Aberrant spectre", "Chaos Elemental"] }, "duplicate": false }`
			executor.EXPECT().ExecuteMonstersRequest(gomock.Any(), common.Params{Where: query}).Return(inline, nil)

			inlinePage2 := mocks.NewMockMonstersResponse(ctrl)
			inlinePage2.EXPECT().GetMonsters().Return([]models.Monster{
				{
					Name: "Chaos Elemental",
				},
			})
			executor.EXPECT().ExecuteMonstersRequest(gomock.Any(), common.Params{Where: query, Page: 2}).Return(inlinePage2, nil)

			api := NewAPI("")
			api.reqExecutor = executor

			return api, []string{"Molanisk", "Aberrant spectre", "Chaos Elemental"}, verifyMonsterNames
		},
		"no names": func(t *testing.T) (*APIClient, []string, checkFn) {
			api := NewAPI("")
			return api, []string{}, verifyError
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			api, names, checkFn := tc(t)
			monsters, err := api.GetMonstersByName(context.Background(), names...)
			checkFn(t, monsters, names, err)
		})
	}
}

func testGetMonstersThatDrop(t *testing.T) {
	type checkFn func(t *testing.T, monsters []models.Monster, expectedNames []string, err error)

	verifyMonsterNames := func(t *testing.T, monsters []models.Monster, expectedNames []string, err error) {
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if len(monsters) != len(expectedNames) {
			t.Errorf("expected %d items, got %d", len(expectedNames), len(monsters))
		}

		for i, monster := range monsters {
			if monster.Name != expectedNames[i] {
				t.Errorf("expected name %s, got %s", expectedNames[i], monster.Name)
			}
		}
	}

	verifyError := func(t *testing.T, monsters []models.Monster, expectedIDs []string, err error) {
		if err == nil {
			t.Errorf("expected an error")
		}
	}

	tests := map[string]func(t *testing.T) (*APIClient, []string, []string, checkFn){
		"success": func(t *testing.T) (*APIClient, []string, []string, checkFn) {
			ctrl := gomock.NewController(t)

			inline := mocks.NewMockMonstersResponse(ctrl)
			inline.EXPECT().GetTotal().Return(2)
			inline.EXPECT().GetMaxResults().Return(25)
			inline.EXPECT().GetMonsters().Return([]models.Monster{
				{
					Name: "Molanisk",
				},
				{
					Name: "Aberrant spectre",
				},
			})

			executor := mocks.NewMockRequestExecutor(ctrl)
			query := `{ "drops": { "$elemMatch": { "name": { "$in": ["Grimy ranarr weed"] } } }, "duplicate": false }`
			executor.EXPECT().ExecuteMonstersRequest(gomock.Any(), common.Params{Where: query}).Return(inline, nil)

			api := NewAPI("")
			api.reqExecutor = executor

			return api, []string{"Grimy ranarr weed"}, []string{"Molanisk", "Aberrant spectre"}, verifyMonsterNames
		},
		"no drops": func(t *testing.T) (*APIClient, []string, []string, checkFn) {
			api := NewAPI("")
			return api, []string{}, nil, verifyError
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			api, item, names, checkFn := tc(t)
			monsters, err := api.GetMonstersThatDrop(context.Background(), item...)
			checkFn(t, monsters, names, err)
		})
	}
}
