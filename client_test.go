package gosrsbox

import (
	"bytes"
	"context"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/atye/gosrsbox/mocks"

	"github.com/golang/mock/gomock"
)

func Test_New(t *testing.T) {
	_ = New(&http.Client{})
	_ = New(nil)
}

func Test_GetAllItems(t *testing.T) {
	type checkFn func(*testing.T, []*Item, error)
	check := func(fns ...checkFn) []checkFn { return fns }

	verifyItems := func(t *testing.T, items []*Item, err error) {
		if len(items) == 0 {
			t.Errorf("expected items, got zero length slice")
		}

		for _, item := range items {
			if item.Name == "" {
				t.Errorf("expected an item name, got empty string")
			}
		}
	}

	verifyNoError := func(t *testing.T, items []*Item, err error) {
		if err != nil {
			t.Errorf("expected no error, got %#v", err)
		}
	}

	verifyError := func(t *testing.T, items []*Item, err error) {
		if err == nil {
			t.Errorf("expected an error, got nil")
		}
	}

	tests := map[string]func(t *testing.T) (*client, []checkFn){
		"success": func(t *testing.T) (*client, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(context.Background(), "GET", "https://raw.githubusercontent.com/osrsbox/osrsbox-db/master/osrsbox/docs/items-complete.json", nil)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: 200,
					Body:       getJSON(t, "testdata/all_items.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
				endpoints: &endpoints{
					items: items,
				},
			}
			return client, check(verifyNoError, verifyItems)
		},
		"http error": func(t *testing.T) (*client, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(context.Background(), "GET", "https://raw.githubusercontent.com/osrsbox/osrsbox-db/master/osrsbox/docs/items-complete.json", nil)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(nil, errors.New("http error"))

			client := &client{
				client: mockHTTPClient,
				endpoints: &endpoints{
					items: items,
				},
			}
			return client, check(verifyError)
		},
		"json error": func(t *testing.T) (*client, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(context.Background(), "GET", "https://raw.githubusercontent.com/osrsbox/osrsbox-db/master/osrsbox/docs/items-complete.json", nil)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewBufferString("bad json")),
				}, nil)

			client := &client{
				client: mockHTTPClient,
				endpoints: &endpoints{
					items: items,
				},
			}
			return client, check(verifyError)
		},
		"nil client": func(t *testing.T) (*client, []checkFn) {
			client := &client{
				client: nil,
				endpoints: &endpoints{
					items: items,
				},
			}
			return client, check(verifyError)
		},
		"bad entity": func(t *testing.T) (*client, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(context.Background(), "GET", "https://raw.githubusercontent.com/osrsbox/osrsbox-db/master/osrsbox/docs/test-complete.json", nil)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: 200,
					Body:       getJSON(t, "testdata/all_items.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
				endpoints: &endpoints{
					items: "test",
				},
			}
			return client, check(verifyError)
		},
		"type error": func(t *testing.T) (*client, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(context.Background(), "GET", "https://raw.githubusercontent.com/osrsbox/osrsbox-db/master/osrsbox/docs/monsters-complete.json", nil)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: 200,
					Body:       getJSON(t, "testdata/all_monsters.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
				endpoints: &endpoints{
					items: monsters,
				},
			}
			return client, check(verifyError)
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client, checkFns := tc(t)

			if len(checkFns) == 0 {
				t.Skipf("Skipping %s because there are no checks in place", name)
			}

			items, err := client.GetAllItems(context.TODO())

			for _, checkFn := range checkFns {
				checkFn(t, items, err)
			}

		})
	}
}

func Test_GetItemsByName(t *testing.T) {
	type checkFn func(*testing.T, []*Item, error)
	check := func(fns ...checkFn) []checkFn { return fns }

	verifyItemNames := func(t *testing.T, items []*Item, err error) {
		if len(items) != 2 {
			t.Errorf("expected items, got zero length slice")
		}

		for i := range items {
			if i == 0 && items[i].Name != "Abyssal whip" {
				t.Errorf("expected Abyssal whip, got %s", items[i].Name)
			}

			if i == 1 && items[i].Name != "Abyssal dagger" {
				t.Errorf("expected Abyssal dagger, got %s", items[i].Name)
			}
		}
	}

	verifyNoError := func(t *testing.T, items []*Item, err error) {
		if err != nil {
			t.Errorf("expected no error, got %#v", err)
		}
	}

	verifyError := func(t *testing.T, items []*Item, err error) {
		if err == nil {
			t.Errorf("expected an error, got nil")
		}
	}

	tests := map[string]func(t *testing.T) (*client, []checkFn){
		"success": func(t *testing.T) (*client, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(context.Background(), "GET", "https://api.osrsbox.com/items?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Abyssal+whip%22%2C+%22Abyssal+dagger%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D", nil)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: 200,
					Body:       getJSON(t, "testdata/where_items.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
				endpoints: &endpoints{
					items: items,
				},
			}
			return client, check(verifyNoError, verifyItemNames)
		},
		"http error": func(t *testing.T) (*client, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(context.Background(), "GET", "https://api.osrsbox.com/items?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Abyssal+whip%22%2C+%22Abyssal+dagger%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D", nil)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(nil, errors.New("http error"))

			client := &client{
				client: mockHTTPClient,
				endpoints: &endpoints{
					items: items,
				},
			}
			return client, check(verifyError)
		},
		"json error": func(t *testing.T) (*client, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(context.Background(), "GET", "https://api.osrsbox.com/items?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Abyssal+whip%22%2C+%22Abyssal+dagger%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D", nil)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewBufferString("bad json")),
				}, nil)

			client := &client{
				client: mockHTTPClient,
				endpoints: &endpoints{
					items: items,
				},
			}
			return client, check(verifyError)
		},
		"nil client": func(t *testing.T) (*client, []checkFn) {
			client := &client{
				client: nil,
				endpoints: &endpoints{
					items: items,
				},
			}
			return client, check(verifyError)
		},
		"type error": func(t *testing.T) (*client, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(context.Background(), "GET", "https://api.osrsbox.com/monsters?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Abyssal+whip%22%2C+%22Abyssal+dagger%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D", nil)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: 200,
					Body:       getJSON(t, "testdata/where_monsters.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
				endpoints: &endpoints{
					items: monsters,
				},
			}
			return client, check(verifyError)
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client, checkFns := tc(t)

			if len(checkFns) == 0 {
				t.Skipf("Skipping %s because there are no checks in place", name)
			}

			items, err := client.GetItemsByName(context.Background(), "Abyssal whip", "Abyssal dagger")

			for _, checkFn := range checkFns {
				checkFn(t, items, err)
			}
		})
	}
}

func Test_GetItemsWhere(t *testing.T) {
	type checkFn func(*testing.T, []*Item, error)
	check := func(fns ...checkFn) []checkFn { return fns }

	verifyItemNames := func(t *testing.T, items []*Item, err error) {
		if len(items) != 2 {
			t.Errorf("expected items, got zero length slice")
		}

		for i := range items {
			if i == 0 && items[i].Name != "Abyssal whip" {
				t.Errorf("expected Abyssal whip, got %s", items[i].Name)
			}

			if i == 1 && items[i].Name != "Abyssal dagger" {
				t.Errorf("expected Abyssal dagger, got %s", items[i].Name)
			}
		}
	}

	verifyNoError := func(t *testing.T, items []*Item, err error) {
		if err != nil {
			t.Errorf("expected no error, got %#v", err)
		}
	}

	verifyError := func(t *testing.T, items []*Item, err error) {
		if err == nil {
			t.Errorf("expected an error, got nil")
		}
	}

	tests := map[string]func(t *testing.T) (*client, []checkFn){
		"success": func(t *testing.T) (*client, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(context.Background(), "GET", "https://api.osrsbox.com/items?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Abyssal+whip%22%2C+%22Abyssal+dagger%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D", nil)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: 200,
					Body:       getJSON(t, "testdata/where_items.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
				endpoints: &endpoints{
					items: items,
				},
			}
			return client, check(verifyNoError, verifyItemNames)
		},
		"bad entity": func(t *testing.T) (*client, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(context.Background(), "GET", "https://api.osrsbox.com/test?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Abyssal+whip%22%2C+%22Abyssal+dagger%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D", nil)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: 200,
					Body:       getJSON(t, "testdata/where_items.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
				endpoints: &endpoints{
					items: "test",
				},
			}
			return client, check(verifyError)
		},
		"nil client": func(t *testing.T) (*client, []checkFn) {
			client := &client{
				client: nil,
				endpoints: &endpoints{
					items: items,
				},
			}
			return client, check(verifyError)
		},
		"type error": func(t *testing.T) (*client, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(context.Background(), "GET", "https://api.osrsbox.com/monsters?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Abyssal+whip%22%2C+%22Abyssal+dagger%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D", nil)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: 200,
					Body:       getJSON(t, "testdata/where_monsters.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
				endpoints: &endpoints{
					items: monsters,
				},
			}
			return client, check(verifyError)
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client, checkFns := tc(t)

			if len(checkFns) == 0 {
				t.Skipf("Skipping %s because there are no checks in place", name)
			}

			items, err := client.GetItemsWhere(context.Background(), `{ "name": { "$in": ["Abyssal whip", "Abyssal dagger"] }, "duplicate": false }`)

			for _, checkFn := range checkFns {
				checkFn(t, items, err)
			}

		})
	}
}

func Test_GetAllMonsters(t *testing.T) {
	type checkFn func(*testing.T, []*Monster, error)
	check := func(fns ...checkFn) []checkFn { return fns }

	verifyMonsters := func(t *testing.T, monsters []*Monster, err error) {
		if len(monsters) == 0 {
			t.Errorf("expected items, got zero length slice")
		}

		for _, item := range monsters {
			if item.Name == "" {
				t.Errorf("expected an item name, got empty string")
			}
		}
	}

	verifyNoError := func(t *testing.T, monsters []*Monster, err error) {
		if err != nil {
			t.Errorf("expected no error, got %#v", err)
		}
	}

	verifyError := func(t *testing.T, monsters []*Monster, err error) {
		if err == nil {
			t.Errorf("expected an error, got nil")
		}
	}

	tests := map[string]func(t *testing.T) (*client, []checkFn){
		"success": func(t *testing.T) (*client, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(context.Background(), "GET", "https://raw.githubusercontent.com/osrsbox/osrsbox-db/master/osrsbox/docs/monsters-complete.json", nil)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: 200,
					Body:       getJSON(t, "testdata/all_monsters.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
				endpoints: &endpoints{
					monsters: monsters,
				},
			}
			return client, check(verifyNoError, verifyMonsters)
		},
		"http error": func(t *testing.T) (*client, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(context.Background(), "GET", "https://raw.githubusercontent.com/osrsbox/osrsbox-db/master/osrsbox/docs/monsters-complete.json", nil)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(nil, errors.New("http error"))

			client := &client{
				client: mockHTTPClient,
				endpoints: &endpoints{
					monsters: monsters,
				},
			}
			return client, check(verifyError)
		},
		"json error": func(t *testing.T) (*client, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(context.Background(), "GET", "https://raw.githubusercontent.com/osrsbox/osrsbox-db/master/osrsbox/docs/monsters-complete.json", nil)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewBufferString("bad json")),
				}, nil)

			client := &client{
				client: mockHTTPClient,
				endpoints: &endpoints{
					monsters: monsters,
				},
			}
			return client, check(verifyError)
		},
		"nil client": func(t *testing.T) (*client, []checkFn) {
			client := &client{
				client: nil,
				endpoints: &endpoints{
					monsters: monsters,
				},
			}
			return client, check(verifyError)
		},
		"type error": func(t *testing.T) (*client, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(context.Background(), "GET", "https://raw.githubusercontent.com/osrsbox/osrsbox-db/master/osrsbox/docs/items-complete.json", nil)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: 200,
					Body:       getJSON(t, "testdata/all_items.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
				endpoints: &endpoints{
					monsters: items,
				},
			}
			return client, check(verifyError)
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client, checkFns := tc(t)

			if len(checkFns) == 0 {
				t.Skipf("Skipping %s because there are no checks in place", name)
			}

			monsters, err := client.GetAllMonsters(context.Background())

			for _, checkFn := range checkFns {
				checkFn(t, monsters, err)
			}

		})
	}
}

func Test_GetMonstersByName(t *testing.T) {
	type checkFn func(*testing.T, []*Monster, error)
	check := func(fns ...checkFn) []checkFn { return fns }

	verifyMonsterNames := func(t *testing.T, monsters []*Monster, err error) {
		if len(monsters) != 2 {
			t.Errorf("expected items, got zero length slice")
		}

		for i := range monsters {
			if i == 0 && monsters[i].Name != "Molanisk" {
				t.Errorf("expected Molanisk, got %s", monsters[i].Name)
			}

			if i == 1 && monsters[i].Name != "Aberrant spectre" {
				t.Errorf("expected Aberrant spectre, got %s", monsters[i].Name)
			}
		}
	}

	verifyNoError := func(t *testing.T, items []*Monster, err error) {
		if err != nil {
			t.Errorf("expected no error, got %#v", err)
		}
	}

	verifyError := func(t *testing.T, items []*Monster, err error) {
		if err == nil {
			t.Errorf("expected an error, got nil")
		}
	}

	tests := map[string]func(t *testing.T) (*client, []checkFn){
		"success": func(t *testing.T) (*client, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(context.Background(), "GET", "https://api.osrsbox.com/monsters?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Molanisk%22%2C+%22Aberrant+spectre%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D", nil)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: 200,
					Body:       getJSON(t, "testdata/where_monsters.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
				endpoints: &endpoints{
					monsters: monsters,
				},
			}
			return client, check(verifyNoError, verifyMonsterNames)
		},
		"http error": func(t *testing.T) (*client, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(context.Background(), "GET", "https://api.osrsbox.com/monsters?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Molanisk%22%2C+%22Aberrant+spectre%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D", nil)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(nil, errors.New("http error"))

			client := &client{
				client: mockHTTPClient,
				endpoints: &endpoints{
					monsters: monsters,
				},
			}
			return client, check(verifyError)
		},
		"json error": func(t *testing.T) (*client, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(context.Background(), "GET", "https://api.osrsbox.com/monsters?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Molanisk%22%2C+%22Aberrant+spectre%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D", nil)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewBufferString("bad json")),
				}, nil)

			client := &client{
				client: mockHTTPClient,
				endpoints: &endpoints{
					monsters: monsters,
				},
			}
			return client, check(verifyError)
		},
		"nil client": func(t *testing.T) (*client, []checkFn) {
			client := &client{
				client: nil,
				endpoints: &endpoints{
					monsters: monsters,
				},
			}
			return client, check(verifyError)
		},
		"type error": func(t *testing.T) (*client, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(context.Background(), "GET", "https://api.osrsbox.com/items?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Molanisk%22%2C+%22Aberrant+spectre%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D", nil)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: 200,
					Body:       getJSON(t, "testdata/where_items.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
				endpoints: &endpoints{
					monsters: items,
				},
			}
			return client, check(verifyError)
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client, checkFns := tc(t)

			if len(checkFns) == 0 {
				t.Skipf("Skipping %s because there are no checks in place", name)
			}

			monsters, err := client.GetMonstersByName(context.Background(), "Molanisk", "Aberrant spectre")

			for _, checkFn := range checkFns {
				checkFn(t, monsters, err)
			}
		})
	}
}

func Test_GetMonstersThatDrop(t *testing.T) {
	type checkFn func(*testing.T, []*Monster, error)
	check := func(fns ...checkFn) []checkFn { return fns }

	verifyMonsterNames := func(t *testing.T, monsters []*Monster, err error) {
		if len(monsters) != 2 {
			t.Errorf("expected items, got zero length slice")
		}

		for i := range monsters {
			if i == 0 && monsters[i].Name != "Molanisk" {
				t.Errorf("expected Molanisk, got %s", monsters[i].Name)
			}

			if i == 1 && monsters[i].Name != "Aberrant spectre" {
				t.Errorf("expected Aberrant spectre, got %s", monsters[i].Name)
			}
		}
	}

	verifyNoError := func(t *testing.T, items []*Monster, err error) {
		if err != nil {
			t.Errorf("expected no error, got %#v", err)
		}
	}

	verifyError := func(t *testing.T, items []*Monster, err error) {
		if err == nil {
			t.Errorf("expected an error, got nil")
		}
	}

	tests := map[string]func(t *testing.T) (*client, []checkFn){
		"success": func(t *testing.T) (*client, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/monsters?where=%7B+%22drops%22%3A+%7B+%22%24elemMatch%22%3A+%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Grimy+ranarr+weed%22%2C+%22Grimy+avantoe%22%5D+%7D+%7D+%7D%2C+%22duplicate%22%3A+false+%7D",
				nil)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: 200,
					Body:       getJSON(t, "testdata/where_monsters.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
				endpoints: &endpoints{
					monsters: monsters,
				},
			}
			return client, check(verifyNoError, verifyMonsterNames)
		},
		"http error": func(t *testing.T) (*client, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/monsters?where=%7B+%22drops%22%3A+%7B+%22%24elemMatch%22%3A+%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Grimy+ranarr+weed%22%2C+%22Grimy+avantoe%22%5D+%7D+%7D+%7D%2C+%22duplicate%22%3A+false+%7D",
				nil)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(nil, errors.New("http error"))

			client := &client{
				client: mockHTTPClient,
				endpoints: &endpoints{
					monsters: monsters,
				},
			}
			return client, check(verifyError)
		},
		"json error": func(t *testing.T) (*client, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/monsters?where=%7B+%22drops%22%3A+%7B+%22%24elemMatch%22%3A+%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Grimy+ranarr+weed%22%2C+%22Grimy+avantoe%22%5D+%7D+%7D+%7D%2C+%22duplicate%22%3A+false+%7D",
				nil)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewBufferString("bad json")),
				}, nil)

			client := &client{
				client: mockHTTPClient,
				endpoints: &endpoints{
					monsters: monsters,
				},
			}
			return client, check(verifyError)
		},
		"nil client": func(t *testing.T) (*client, []checkFn) {
			client := &client{
				client: nil,
				endpoints: &endpoints{
					monsters: monsters,
				},
			}
			return client, check(verifyError)
		},
		"type error": func(t *testing.T) (*client, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/items?where=%7B+%22drops%22%3A+%7B+%22%24elemMatch%22%3A+%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Grimy+ranarr+weed%22%2C+%22Grimy+avantoe%22%5D+%7D+%7D+%7D%2C+%22duplicate%22%3A+false+%7D",
				nil)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: 200,
					Body:       getJSON(t, "testdata/where_items.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
				endpoints: &endpoints{
					monsters: items,
				},
			}
			return client, check(verifyError)
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client, checkFns := tc(t)

			if len(checkFns) == 0 {
				t.Skipf("Skipping %s because there are no checks in place", name)
			}

			monsters, err := client.GetMonstersThatDrop(context.Background(), "Grimy ranarr weed", "Grimy avantoe")

			for _, checkFn := range checkFns {
				checkFn(t, monsters, err)
			}
		})
	}
}

func Test_GetMonstersWhere(t *testing.T) {
	type checkFn func(*testing.T, []*Monster, error)
	check := func(fns ...checkFn) []checkFn { return fns }

	verifyMonsterNames := func(t *testing.T, monsters []*Monster, err error) {
		if len(monsters) != 2 {
			t.Errorf("expected items, got zero length slice")
		}

		for i := range monsters {
			if i == 0 && monsters[i].Name != "Molanisk" {
				t.Errorf("expected Molanisk, got %s", monsters[i].Name)
			}

			if i == 1 && monsters[i].Name != "Aberrant spectre" {
				t.Errorf("expected Aberrant spectre, got %s", monsters[i].Name)
			}
		}
	}

	verifyNoError := func(t *testing.T, items []*Monster, err error) {
		if err != nil {
			t.Errorf("expected no error, got %#v", err)
		}
	}

	verifyError := func(t *testing.T, monsters []*Monster, err error) {
		if err == nil {
			t.Errorf("expected an error, got nil")
		}
	}

	tests := map[string]func(t *testing.T) (*client, []checkFn){
		"success": func(t *testing.T) (*client, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(context.Background(), "GET", "https://api.osrsbox.com/monsters?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Molanisk%22%2C+%22Aberrant+spectre%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D", nil)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: 200,
					Body:       getJSON(t, "testdata/where_monsters.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
				endpoints: &endpoints{
					monsters: monsters,
				},
			}
			return client, check(verifyNoError, verifyMonsterNames)
		},
		"nil client": func(t *testing.T) (*client, []checkFn) {
			client := &client{
				client: nil,
				endpoints: &endpoints{
					monsters: monsters,
				},
			}
			return client, check(verifyError)
		},
		"type error": func(t *testing.T) (*client, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(context.Background(), "GET", "https://api.osrsbox.com/items?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Molanisk%22%2C+%22Aberrant+spectre%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D", nil)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: 200,
					Body:       getJSON(t, "testdata/where_items.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
				endpoints: &endpoints{
					monsters: items,
				},
			}
			return client, check(verifyError)
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client, checkFns := tc(t)

			if len(checkFns) == 0 {
				t.Skipf("Skipping %s because there are no checks in place", name)
			}

			monsters, err := client.GetMonstersWhere(context.Background(), `{ "name": { "$in": ["Molanisk", "Aberrant spectre"] }, "duplicate": false }`)

			for _, checkFn := range checkFns {
				checkFn(t, monsters, err)
			}

		})
	}
}

func Test_GetAllPrayers(t *testing.T) {
	type checkFn func(*testing.T, []*Prayer, error)
	check := func(fns ...checkFn) []checkFn { return fns }

	verifyPrayers := func(t *testing.T, prayers []*Prayer, err error) {
		if len(prayers) == 0 {
			t.Errorf("expected items, got zero length slice")
		}

		for _, item := range prayers {
			if item.Name == "" {
				t.Errorf("expected an item name, got empty string")
			}
		}
	}

	verifyNoError := func(t *testing.T, prayers []*Prayer, err error) {
		if err != nil {
			t.Errorf("expected no error, got %#v", err)
		}
	}

	verifyError := func(t *testing.T, prayers []*Prayer, err error) {
		if err == nil {
			t.Errorf("expected an error, got nil")
		}
	}

	tests := map[string]func(t *testing.T) (*client, []checkFn){
		"success": func(t *testing.T) (*client, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(context.Background(), "GET", "https://raw.githubusercontent.com/osrsbox/osrsbox-db/master/osrsbox/docs/prayers-complete.json", nil)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: 200,
					Body:       getJSON(t, "testdata/all_monsters.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
				endpoints: &endpoints{
					prayers: prayers,
				},
			}
			return client, check(verifyNoError, verifyPrayers)
		},
		"http error": func(t *testing.T) (*client, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(context.Background(), "GET", "https://raw.githubusercontent.com/osrsbox/osrsbox-db/master/osrsbox/docs/prayers-complete.json", nil)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(nil, errors.New("http error"))

			client := &client{
				client: mockHTTPClient,
				endpoints: &endpoints{
					prayers: prayers,
				},
			}
			return client, check(verifyError)
		},
		"json error": func(t *testing.T) (*client, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(context.Background(), "GET", "https://raw.githubusercontent.com/osrsbox/osrsbox-db/master/osrsbox/docs/prayers-complete.json", nil)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewBufferString("bad json")),
				}, nil)

			client := &client{
				client: mockHTTPClient,
				endpoints: &endpoints{
					prayers: prayers,
				},
			}
			return client, check(verifyError)
		},
		"nil client": func(t *testing.T) (*client, []checkFn) {
			client := &client{
				client: nil,
				endpoints: &endpoints{
					prayers: prayers,
				},
			}
			return client, check(verifyError)
		},
		"type error": func(t *testing.T) (*client, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(context.Background(), "GET", "https://raw.githubusercontent.com/osrsbox/osrsbox-db/master/osrsbox/docs/items-complete.json", nil)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: 200,
					Body:       getJSON(t, "testdata/all_items.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
				endpoints: &endpoints{
					prayers: items,
				},
			}
			return client, check(verifyError)
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client, checkFns := tc(t)

			if len(checkFns) == 0 {
				t.Skipf("Skipping %s because there are no checks in place", name)
			}

			prayers, err := client.GetAllPrayers(context.Background())

			for _, checkFn := range checkFns {
				checkFn(t, prayers, err)
			}

		})
	}
}

func Test_GetPrayersByName(t *testing.T) {
	type checkFn func(*testing.T, []*Prayer, error)
	check := func(fns ...checkFn) []checkFn { return fns }

	verifyMonsterNames := func(t *testing.T, prayers []*Prayer, err error) {
		if len(prayers) != 2 {
			t.Errorf("expected items, got zero length slice")
		}

		for i := range prayers {
			if i == 0 && prayers[i].Name != "Thick Skin" {
				t.Errorf("expected Thick Skin, got %s", prayers[i].Name)
			}

			if i == 1 && prayers[i].Name != "Burst of Strength" {
				t.Errorf("expected Burst of Strength, got %s", prayers[i].Name)
			}
		}
	}

	verifyNoError := func(t *testing.T, items []*Prayer, err error) {
		if err != nil {
			t.Errorf("expected no error, got %#v", err)
		}
	}

	verifyError := func(t *testing.T, items []*Prayer, err error) {
		if err == nil {
			t.Errorf("expected an error, got nil")
		}
	}

	tests := map[string]func(t *testing.T) (*client, []checkFn){
		"success": func(t *testing.T) (*client, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(context.Background(), "GET", "https://api.osrsbox.com/prayers?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Thick+Skin%22%2C+%22Burst+of+Strength%22%5D+%7D+%7D", nil)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: 200,
					Body:       getJSON(t, "testdata/where_prayers.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
				endpoints: &endpoints{
					prayers: prayers,
				},
			}
			return client, check(verifyNoError, verifyMonsterNames)
		},
		"http error": func(t *testing.T) (*client, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(context.Background(), "GET", "https://api.osrsbox.com/prayers?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Thick+Skin%22%2C+%22Burst+of+Strength%22%5D+%7D+%7D", nil)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(nil, errors.New("http error"))

			client := &client{
				client: mockHTTPClient,
				endpoints: &endpoints{
					prayers: prayers,
				},
			}
			return client, check(verifyError)
		},
		"json error": func(t *testing.T) (*client, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(context.Background(), "GET", "https://api.osrsbox.com/prayers?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Thick+Skin%22%2C+%22Burst+of+Strength%22%5D+%7D+%7D", nil)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewBufferString("bad json")),
				}, nil)

			client := &client{
				client: mockHTTPClient,
				endpoints: &endpoints{
					prayers: prayers,
				},
			}
			return client, check(verifyError)
		},
		"nil client": func(t *testing.T) (*client, []checkFn) {
			client := &client{
				client: nil,
				endpoints: &endpoints{
					prayers: prayers,
				},
			}
			return client, check(verifyError)
		},
		"type error": func(t *testing.T) (*client, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(context.Background(), "GET", "https://api.osrsbox.com/items?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Thick+Skin%22%2C+%22Burst+of+Strength%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D", nil)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: 200,
					Body:       getJSON(t, "testdata/where_items.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
				endpoints: &endpoints{
					prayers: items,
				},
			}
			return client, check(verifyError)
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client, checkFns := tc(t)

			if len(checkFns) == 0 {
				t.Skipf("Skipping %s because there are no checks in place", name)
			}

			prayers, err := client.GetPrayersByName(context.Background(), "Thick Skin", "Burst of Strength")

			for _, checkFn := range checkFns {
				checkFn(t, prayers, err)
			}
		})
	}
}

func Test_GetPrayersWhere(t *testing.T) {
	type checkFn func(*testing.T, []*Prayer, error)
	check := func(fns ...checkFn) []checkFn { return fns }

	verifyPrayers := func(t *testing.T, prayers []*Prayer, err error) {
		if len(prayers) != 2 {
			t.Errorf("expected items, got zero length slice")
		}

		for i := range prayers {
			if i == 0 && prayers[i].Name != "Thick Skin" {
				t.Errorf("expected Thick Skin, got %s", prayers[i].Name)
			}

			if i == 1 && prayers[i].Name != "Burst of Strength" {
				t.Errorf("expected Burst of Strength, got %s", prayers[i].Name)
			}
		}
	}

	verifyNoError := func(t *testing.T, items []*Prayer, err error) {
		if err != nil {
			t.Errorf("expected no error, got %#v", err)
		}
	}

	verifyError := func(t *testing.T, prayers []*Prayer, err error) {
		if err == nil {
			t.Errorf("expected an error, got nil")
		}
	}

	tests := map[string]func(t *testing.T) (*client, []checkFn){
		"success": func(t *testing.T) (*client, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(context.Background(), "GET", "https://api.osrsbox.com/prayers?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Thick+Skin%22%2C+%22Burst+of+Strength%22%5D+%7D+%7D", nil)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: 200,
					Body:       getJSON(t, "testdata/where_prayers.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
				endpoints: &endpoints{
					prayers: prayers,
				},
			}
			return client, check(verifyNoError, verifyPrayers)
		},
		"nil client": func(t *testing.T) (*client, []checkFn) {
			client := &client{
				client: nil,
				endpoints: &endpoints{
					prayers: prayers,
				},
			}
			return client, check(verifyError)
		},
		"type error": func(t *testing.T) (*client, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(context.Background(), "GET", "https://api.osrsbox.com/items?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Thick+Skin%22%2C+%22Burst+of+Strength%22%5D+%7D+%7D", nil)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: 200,
					Body:       getJSON(t, "testdata/where_items.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
				endpoints: &endpoints{
					prayers: items,
				},
			}
			return client, check(verifyError)
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client, checkFns := tc(t)

			if len(checkFns) == 0 {
				t.Skipf("Skipping %s because there are no checks in place", name)
			}

			prayers, err := client.GetPrayersWhere(context.Background(), `{ "name": { "$in": ["Thick Skin", "Burst of Strength"] } }`)

			for _, checkFn := range checkFns {
				checkFn(t, prayers, err)
			}

		})
	}
}

func getJSON(t *testing.T, file string) io.ReadCloser {
	t.Helper()
	json, err := os.Open(file)
	if err != nil {
		t.Fatal(err)
	}
	return json
}
