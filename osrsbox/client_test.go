package osrsbox

import (
	"bytes"
	"context"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"testing"

	"github.com/atye/gosrsbox/osrsbox/mocks"
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

	tests := map[string]func(t *testing.T) (*client, context.Context, []checkFn){
		"success": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://raw.githubusercontent.com/osrsbox/osrsbox-db/master/osrsbox/docs/items-complete.json",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       getJSON(t, "testdata/all_items.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
			}
			return client, context.Background(), check(verifyNoError, verifyItems)
		},
		"http error": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://raw.githubusercontent.com/osrsbox/osrsbox-db/master/osrsbox/docs/items-complete.json",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(nil, errors.New("http error"))

			client := &client{
				client: mockHTTPClient,
			}
			return client, context.Background(), check(verifyError)
		},
		"status not ok": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://raw.githubusercontent.com/osrsbox/osrsbox-db/master/osrsbox/docs/items-complete.json",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: 500,
					Body:       ioutil.NopCloser(bytes.NewBufferString("{}")),
				}, nil)

			client := &client{
				client: mockHTTPClient,
			}
			return client, context.Background(), check(verifyError)
		},
		"json error": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://raw.githubusercontent.com/osrsbox/osrsbox-db/master/osrsbox/docs/items-complete.json",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewBufferString("bad json")),
				}, nil)

			client := &client{
				client: mockHTTPClient,
			}
			return client, context.Background(), check(verifyError)
		},
		"nil client": func(t *testing.T) (*client, context.Context, []checkFn) {
			client := &client{
				client: nil,
			}
			return client, context.Background(), check(verifyError)
		},
		"no context": func(t *testing.T) (*client, context.Context, []checkFn) {
			client := &client{
				client: &http.Client{},
			}
			return client, nil, check(verifyError)
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client, ctx, checkFns := tc(t)

			if len(checkFns) == 0 {
				t.Skipf("Skipping %s because there are no checks in place", name)
			}

			items, err := client.GetAllItems(ctx)

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
		if len(items) != 4 {
			t.Errorf("expected four items, got zero length slice")
		}

		for i := range items {
			if i == 0 && items[i].Name != "Abyssal whip" {
				t.Errorf("expected Abyssal whip, got %s", items[i].Name)
			}

			if i == 1 && items[i].Name != "Abyssal dagger" {
				t.Errorf("expected Abyssal dagger, got %s", items[i].Name)
			}

			if i == 2 && items[i].Name != "Rune platebody" {
				t.Errorf("expected Rune platebody, got %s", items[i].Name)
			}

			if i == 3 && items[i].Name != "Dragon scimitar" {
				t.Errorf("expected Dragon scimitar, got %s", items[i].Name)
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

	tests := map[string]func(t *testing.T) (*client, context.Context, []checkFn){
		"success": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			firstReq, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/items?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Abyssal+whip%22%2C+%22Abyssal+dagger%22%2C+%22Dragon+scimitar%22%2C+%22Rune+platebody%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			secondReq, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/items?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Abyssal+whip%22%2C+%22Abyssal+dagger%22%2C+%22Dragon+scimitar%22%2C+%22Rune+platebody%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D&page=2",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(firstReq).
				Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       getJSON(t, "testdata/where_items_page1.json"),
				}, nil)

			mockHTTPClient.EXPECT().
				Do(secondReq).
				Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       getJSON(t, "testdata/where_items_page2.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
			}
			return client, context.Background(), check(verifyNoError, verifyItemNames)
		},
		"nil client": func(t *testing.T) (*client, context.Context, []checkFn) {
			client := &client{
				client: nil,
			}
			return client, context.Background(), check(verifyError)
		},
		"no context": func(t *testing.T) (*client, context.Context, []checkFn) {
			client := &client{
				client: &http.Client{},
			}
			return client, nil, check(verifyError)
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client, ctx, checkFns := tc(t)

			if len(checkFns) == 0 {
				t.Skipf("Skipping %s because there are no checks in place", name)
			}

			items, err := client.GetItemsByName(ctx, "abyssal Whip", "abyssal Dagger", "Dragon Scimitar", "Rune Platebody")

			for _, checkFn := range checkFns {
				checkFn(t, items, err)
			}
		})
	}
}

func Test_GetItemsByWikiName(t *testing.T) {
	type checkFn func(*testing.T, []*Item, error)
	check := func(fns ...checkFn) []checkFn { return fns }

	verifyItemNames := func(t *testing.T, items []*Item, err error) {
		if len(items) != 4 {
			t.Errorf("expected four items, got zero length slice")
		}

		for i := range items {
			if i == 0 && items[i].WikiName != "Abyssal whip" {
				t.Errorf("expected Abyssal whip, got %s", items[i].WikiName)
			}

			if i == 1 && items[i].WikiName != "Abyssal dagger (Unpoisoned)" {
				t.Errorf("expected Abyssal dagger, got %s", items[i].WikiName)
			}

			if i == 2 && items[i].WikiName != "Rune platebody" {
				t.Errorf("expected Rune platebody, got %s", items[i].WikiName)
			}

			if i == 3 && items[i].WikiName != "Dragon scimitar" {
				t.Errorf("expected Dragon scimitar, got %s", items[i].WikiName)
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

	tests := map[string]func(t *testing.T) (*client, context.Context, []checkFn){
		"success": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			firstReq, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/items?where=%7B+%22wiki_name%22%3A+%7B+%22%24in%22%3A+%5B%22Abyssal+whip%22%2C+%22Abyssal+dagger+%28Unpoisoned%29%22%2C+%22Dragon+scimitar%22%2C+%22Rune+platebody%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			secondReq, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/items?where=%7B+%22wiki_name%22%3A+%7B+%22%24in%22%3A+%5B%22Abyssal+whip%22%2C+%22Abyssal+dagger+%28Unpoisoned%29%22%2C+%22Dragon+scimitar%22%2C+%22Rune+platebody%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D&page=2",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(firstReq).
				Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       getJSON(t, "testdata/where_items_page1.json"),
				}, nil)

			mockHTTPClient.EXPECT().
				Do(secondReq).
				Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       getJSON(t, "testdata/where_items_page2.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
			}
			return client, context.Background(), check(verifyNoError, verifyItemNames)
		},
		"nil client": func(t *testing.T) (*client, context.Context, []checkFn) {
			client := &client{
				client: nil,
			}
			return client, context.Background(), check(verifyError)
		},
		"no context": func(t *testing.T) (*client, context.Context, []checkFn) {
			client := &client{
				client: &http.Client{},
			}
			return client, nil, check(verifyError)
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client, ctx, checkFns := tc(t)

			if len(checkFns) == 0 {
				t.Skipf("Skipping %s because there are no checks in place", name)
			}

			items, err := client.GetItemsByWikiName(ctx, "Abyssal whip", "Abyssal dagger (Unpoisoned)", "Dragon scimitar", "Rune platebody")

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
		if len(items) != 4 {
			t.Errorf("expected four items, got %d length slice", len(items))
		}

		for i := range items {
			if i == 0 && items[i].Name != "Abyssal whip" {
				t.Errorf("expected Abyssal whip, got %s", items[i].Name)
			}

			if i == 1 && items[i].Name != "Abyssal dagger" {
				t.Errorf("expected Abyssal dagger, got %s", items[i].Name)
			}

			if i == 2 && items[i].Name != "Rune platebody" {
				t.Errorf("expected Rune platebody, got %s", items[i].Name)
			}

			if i == 3 && items[i].Name != "Dragon scimitar" {
				t.Errorf("expected Dragon scimitar, got %s", items[i].Name)
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

	tests := map[string]func(t *testing.T) (*client, context.Context, []checkFn){
		"success": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			firstReq, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/items?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Abyssal+whip%22%2C+%22Abyssal+dagger%22%2C+%22Dragon+scimitar%22%2C+%22Rune+platebody%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			secondReq, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/items?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Abyssal+whip%22%2C+%22Abyssal+dagger%22%2C+%22Dragon+scimitar%22%2C+%22Rune+platebody%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D&page=2",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(firstReq).
				Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       getJSON(t, "testdata/where_items_page1.json"),
				}, nil)

			mockHTTPClient.EXPECT().
				Do(secondReq).
				Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       getJSON(t, "testdata/where_items_page2.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
				wg:     sync.WaitGroup{},
				mu:     sync.Mutex{},
			}
			return client, context.Background(), check(verifyNoError, verifyItemNames)
		},
		"success one page": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/items?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Abyssal+whip%22%2C+%22Abyssal+dagger%22%2C+%22Dragon+scimitar%22%2C+%22Rune+platebody%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       getJSON(t, "testdata/where_items_one_page.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
				wg:     sync.WaitGroup{},
				mu:     sync.Mutex{},
			}
			return client, context.Background(), check(verifyItemNames, verifyNoError)
		},
		"nil client": func(t *testing.T) (*client, context.Context, []checkFn) {
			client := &client{
				client: nil,
			}
			return client, context.Background(), check(verifyError)
		},
		"no context": func(t *testing.T) (*client, context.Context, []checkFn) {
			client := &client{
				client: &http.Client{},
			}
			return client, nil, check(verifyError)
		},
		"request error": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/items?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Abyssal+whip%22%2C+%22Abyssal+dagger%22%2C+%22Dragon+scimitar%22%2C+%22Rune+platebody%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(nil, errors.New("http error"))

			client := &client{
				client: mockHTTPClient,
			}
			return client, context.Background(), check(verifyError)
		},
		"second request error": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			firstReq, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/items?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Abyssal+whip%22%2C+%22Abyssal+dagger%22%2C+%22Dragon+scimitar%22%2C+%22Rune+platebody%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			secondReq, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/items?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Abyssal+whip%22%2C+%22Abyssal+dagger%22%2C+%22Dragon+scimitar%22%2C+%22Rune+platebody%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D&page=2",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(firstReq).
				Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       getJSON(t, "testdata/where_items_page1.json"),
				}, nil)

			mockHTTPClient.EXPECT().
				Do(secondReq).
				Return(nil, errors.New("http error"))

			client := &client{
				client: mockHTTPClient,
			}
			return client, context.Background(), check(verifyError)
		},
		"formatted error": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/items?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Abyssal+whip%22%2C+%22Abyssal+dagger%22%2C+%22Dragon+scimitar%22%2C+%22Rune+platebody%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: http.StatusBadRequest,
					Body:       getJSON(t, "testdata/bad_request_error.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
				wg:     sync.WaitGroup{},
				mu:     sync.Mutex{},
			}
			return client, context.Background(), check(verifyError)
		},
		"other error": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/items?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Abyssal+whip%22%2C+%22Abyssal+dagger%22%2C+%22Dragon+scimitar%22%2C+%22Rune+platebody%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: http.StatusInternalServerError,
					Body:       ioutil.NopCloser(bytes.NewBufferString("{}")),
				}, nil)

			client := &client{
				client: mockHTTPClient,
				wg:     sync.WaitGroup{},
				mu:     sync.Mutex{},
			}
			return client, context.Background(), check(verifyError)
		},
		"json error": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/items?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Abyssal+whip%22%2C+%22Abyssal+dagger%22%2C+%22Dragon+scimitar%22%2C+%22Rune+platebody%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewBufferString("bad json")),
				}, nil)

			client := &client{
				client: mockHTTPClient,
			}
			return client, context.Background(), check(verifyError)
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client, ctx, checkFns := tc(t)

			if len(checkFns) == 0 {
				t.Skipf("Skipping %s because there are no checks in place", name)
			}

			items, err := client.GetItemsWhere(ctx, `{ "name": { "$in": ["Abyssal whip", "Abyssal dagger", "Dragon scimitar", "Rune platebody"] }, "duplicate": false }`)

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
			t.Errorf("expected monsters, got zero length slice")
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

	tests := map[string]func(t *testing.T) (*client, context.Context, []checkFn){
		"success": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://raw.githubusercontent.com/osrsbox/osrsbox-db/master/osrsbox/docs/monsters-complete.json",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       getJSON(t, "testdata/all_monsters.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
			}
			return client, context.Background(), check(verifyNoError, verifyMonsters)
		},
		"http error": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://raw.githubusercontent.com/osrsbox/osrsbox-db/master/osrsbox/docs/monsters-complete.json",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(nil, errors.New("http error"))

			client := &client{
				client: mockHTTPClient,
			}
			return client, context.Background(), check(verifyError)
		},
		"status not ok": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://raw.githubusercontent.com/osrsbox/osrsbox-db/master/osrsbox/docs/monsters-complete.json",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: http.StatusInternalServerError,
					Body:       ioutil.NopCloser(bytes.NewBufferString("{}")),
				}, nil)

			client := &client{
				client: mockHTTPClient,
			}
			return client, context.Background(), check(verifyError)
		},
		"json error": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://raw.githubusercontent.com/osrsbox/osrsbox-db/master/osrsbox/docs/monsters-complete.json",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewBufferString("bad json")),
				}, nil)

			client := &client{
				client: mockHTTPClient,
			}
			return client, context.Background(), check(verifyError)
		},
		"nil client": func(t *testing.T) (*client, context.Context, []checkFn) {
			client := &client{
				client: nil,
			}
			return client, context.Background(), check(verifyError)
		},
		"no context": func(t *testing.T) (*client, context.Context, []checkFn) {
			client := &client{
				client: &http.Client{},
			}
			return client, nil, check(verifyError)
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client, ctx, checkFns := tc(t)

			if len(checkFns) == 0 {
				t.Skipf("Skipping %s because there are no checks in place", name)
			}

			monsters, err := client.GetAllMonsters(ctx)

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
		if len(monsters) != 4 {
			t.Errorf("expected four monsters, got zero length slice")
		}

		for i := range monsters {
			if i == 0 && monsters[i].Name != "Molanisk" {
				t.Errorf("expected Molanisk, got %s", monsters[i].Name)
			}

			if i == 1 && monsters[i].Name != "Aberrant spectre" {
				t.Errorf("expected Aberrant spectre, got %s", monsters[i].Name)
			}

			if i == 2 && monsters[i].Name != "Chaos Elemental" {
				t.Errorf("expected Molanisk, got %s", monsters[i].Name)
			}

			if i == 3 && monsters[i].Name != "Venenatis" {
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

	tests := map[string]func(t *testing.T) (*client, context.Context, []checkFn){
		"success": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			firstReq, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/monsters?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Molanisk%22%2C+%22Aberrant+spectre%22%2C+%22Chaos+Elemental%22%2C+%22Venenatis%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			secondReq, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/monsters?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Molanisk%22%2C+%22Aberrant+spectre%22%2C+%22Chaos+Elemental%22%2C+%22Venenatis%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D&page=2",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(firstReq).
				Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       getJSON(t, "testdata/where_monsters_page1.json"),
				}, nil)

			mockHTTPClient.EXPECT().
				Do(secondReq).
				Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       getJSON(t, "testdata/where_monsters_page2.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
			}
			return client, context.Background(), check(verifyNoError, verifyMonsterNames)
		},
		"nil client": func(t *testing.T) (*client, context.Context, []checkFn) {
			client := &client{
				client: nil,
			}
			return client, context.Background(), check(verifyError)
		},
		"no context": func(t *testing.T) (*client, context.Context, []checkFn) {
			client := &client{
				client: &http.Client{},
			}
			return client, nil, check(verifyError)
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client, ctx, checkFns := tc(t)

			if len(checkFns) == 0 {
				t.Skipf("Skipping %s because there are no checks in place", name)
			}

			monsters, err := client.GetMonstersByName(ctx, "Molanisk", "Aberrant spectre", "Chaos Elemental", "Venenatis")

			for _, checkFn := range checkFns {
				checkFn(t, monsters, err)
			}
		})
	}
}

func Test_GetMonstersByWikiName(t *testing.T) {
	type checkFn func(*testing.T, []*Monster, error)
	check := func(fns ...checkFn) []checkFn { return fns }

	verifyMonsterNames := func(t *testing.T, monsters []*Monster, err error) {
		if len(monsters) != 4 {
			t.Errorf("expected four monsters, got zero length slice")
		}

		for i := range monsters {
			if i == 0 && monsters[i].WikiName != "Molanisk" {
				t.Errorf("expected Molanisk, got %s", monsters[i].WikiName)
			}

			if i == 1 && monsters[i].WikiName != "Aberrant spectre" {
				t.Errorf("expected Aberrant spectre, got %s", monsters[i].WikiName)
			}

			if i == 2 && monsters[i].WikiName != "Chaos Elemental" {
				t.Errorf("expected Molanisk, got %s", monsters[i].WikiName)
			}

			if i == 3 && monsters[i].WikiName != "Venenatis" {
				t.Errorf("expected Aberrant spectre, got %s", monsters[i].WikiName)
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

	tests := map[string]func(t *testing.T) (*client, context.Context, []checkFn){
		"success": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			firstReq, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/monsters?where=%7B+%22wiki_name%22%3A+%7B+%22%24in%22%3A+%5B%22Molanisk%22%2C+%22Aberrant+spectre%22%2C+%22Chaos+Elemental%22%2C+%22Venenatis%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			secondReq, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/monsters?where=%7B+%22wiki_name%22%3A+%7B+%22%24in%22%3A+%5B%22Molanisk%22%2C+%22Aberrant+spectre%22%2C+%22Chaos+Elemental%22%2C+%22Venenatis%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D&page=2",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(firstReq).
				Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       getJSON(t, "testdata/where_monsters_page1.json"),
				}, nil)

			mockHTTPClient.EXPECT().
				Do(secondReq).
				Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       getJSON(t, "testdata/where_monsters_page2.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
			}
			return client, context.Background(), check(verifyNoError, verifyMonsterNames)
		},
		"nil client": func(t *testing.T) (*client, context.Context, []checkFn) {
			client := &client{
				client: nil,
			}
			return client, context.Background(), check(verifyError)
		},
		"no context": func(t *testing.T) (*client, context.Context, []checkFn) {
			client := &client{
				client: &http.Client{},
			}
			return client, nil, check(verifyError)
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client, ctx, checkFns := tc(t)

			if len(checkFns) == 0 {
				t.Skipf("Skipping %s because there are no checks in place", name)
			}

			monsters, err := client.GetMonstersByWikiName(ctx, "Molanisk", "Aberrant spectre", "Chaos Elemental", "Venenatis")

			for _, checkFn := range checkFns {
				checkFn(t, monsters, err)
			}
		})
	}
}

func Test_GetMonstersThatDrop(t *testing.T) {
	type checkFn func(*testing.T, []*Monster, error)
	check := func(fns ...checkFn) []checkFn { return fns }

	verifyMonsterDrops := func(t *testing.T, monsters []*Monster, err error) {
		if len(monsters) != 4 {
			t.Errorf("expected four monsters, got %d length slice", len(monsters))
		}

		for _, monster := range monsters {
			if !inDrops([]string{"Grimy ranarr weed", "Grimy avantoe", "Grimy snapdragon"}, monster.Drops) {
				t.Errorf("expected %s in drops, got drops %v", "Grimy ranarr weed", monster.Drops)
			}

			if !inDrops([]string{"Grimy ranarr weed", "Grimy avantoe", "Grimy snapdragon"}, monster.Drops) {
				t.Errorf("expected %s in drops, got drops %v", "Grimy avantoe", monster.Drops)
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

	tests := map[string]func(t *testing.T) (*client, context.Context, []checkFn){
		"success": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			firstReq, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/monsters?where=%7B+%22drops%22%3A+%7B+%22%24elemMatch%22%3A+%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Grimy+ranarr+weed%22%2C+%22Grimy+avantoe%22%2C+%22Grimy+snapdragon%22%5D+%7D+%7D+%7D%2C+%22duplicate%22%3A+false+%7D",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			secondReq, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/monsters?where=%7B+%22drops%22%3A+%7B+%22%24elemMatch%22%3A+%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Grimy+ranarr+weed%22%2C+%22Grimy+avantoe%22%2C+%22Grimy+snapdragon%22%5D+%7D+%7D+%7D%2C+%22duplicate%22%3A+false+%7D&page=2",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(firstReq).
				Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       getJSON(t, "testdata/where_monsters_page1.json"),
				}, nil)

			mockHTTPClient.EXPECT().
				Do(secondReq).
				Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       getJSON(t, "testdata/where_monsters_page2.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
			}
			return client, context.Background(), check(verifyNoError, verifyMonsterDrops)
		},
		"nil client": func(t *testing.T) (*client, context.Context, []checkFn) {
			client := &client{
				client: nil,
			}
			return client, context.Background(), check(verifyError)
		},
		"no context": func(t *testing.T) (*client, context.Context, []checkFn) {
			client := &client{
				client: &http.Client{},
			}
			return client, nil, check(verifyError)
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client, ctx, checkFns := tc(t)

			if len(checkFns) == 0 {
				t.Skipf("Skipping %s because there are no checks in place", name)
			}

			monsters, err := client.GetMonstersThatDrop(ctx, "grimy Ranarr weed", "grimy avantoe", "Grimy Snapdragon")

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
		if len(monsters) != 4 {
			t.Errorf("expected four monsters, got %d length slice", len(monsters))
		}

		for i := range monsters {
			if i == 0 && monsters[i].Name != "Molanisk" {
				t.Errorf("expected Molanisk, got %s", monsters[i].Name)
			}

			if i == 1 && monsters[i].Name != "Aberrant spectre" {
				t.Errorf("expected Aberrant spectre, got %s", monsters[i].Name)
			}

			if i == 2 && monsters[i].Name != "Chaos Elemental" {
				t.Errorf("expected Molanisk, got %s", monsters[i].Name)
			}

			if i == 3 && monsters[i].Name != "Venenatis" {
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

	tests := map[string]func(t *testing.T) (*client, context.Context, []checkFn){
		"success": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			firstReq, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/monsters?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Molanisk%22%2C+%22Aberrant+spectre%22%2C+%22Chaos+Elemental%22%2C+%22Venenatis%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			secondReq, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/monsters?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Molanisk%22%2C+%22Aberrant+spectre%22%2C+%22Chaos+Elemental%22%2C+%22Venenatis%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D&page=2",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(firstReq).
				Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       getJSON(t, "testdata/where_monsters_page1.json"),
				}, nil)

			mockHTTPClient.EXPECT().
				Do(secondReq).
				Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       getJSON(t, "testdata/where_monsters_page2.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
				wg:     sync.WaitGroup{},
				mu:     sync.Mutex{},
			}
			return client, context.Background(), check(verifyNoError, verifyMonsterNames)
		},
		"success one page": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			firstReq, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/monsters?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Molanisk%22%2C+%22Aberrant+spectre%22%2C+%22Chaos+Elemental%22%2C+%22Venenatis%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(firstReq).
				Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       getJSON(t, "testdata/where_monsters_one_page.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
			}
			return client, context.Background(), check(verifyNoError, verifyMonsterNames)
		},
		"nil client": func(t *testing.T) (*client, context.Context, []checkFn) {
			client := &client{
				client: nil,
			}
			return client, context.Background(), check(verifyError)
		},
		"no context": func(t *testing.T) (*client, context.Context, []checkFn) {
			client := &client{
				client: &http.Client{},
			}
			return client, nil, check(verifyError)
		},
		"request error": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/monsters?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Molanisk%22%2C+%22Aberrant+spectre%22%2C+%22Chaos+Elemental%22%2C+%22Venenatis%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(nil, errors.New("http error"))

			client := &client{
				client: mockHTTPClient,
			}
			return client, context.Background(), check(verifyError)
		},
		"second request error": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			firstReq, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/monsters?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Molanisk%22%2C+%22Aberrant+spectre%22%2C+%22Chaos+Elemental%22%2C+%22Venenatis%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			secondReq, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/monsters?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Molanisk%22%2C+%22Aberrant+spectre%22%2C+%22Chaos+Elemental%22%2C+%22Venenatis%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D&page=2",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(firstReq).
				Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       getJSON(t, "testdata/where_monsters_page1.json"),
				}, nil)

			mockHTTPClient.EXPECT().
				Do(secondReq).
				Return(nil, errors.New("http error"))

			client := &client{
				client: mockHTTPClient,
			}
			return client, context.Background(), check(verifyError)
		},
		"formatted error": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/monsters?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Molanisk%22%2C+%22Aberrant+spectre%22%2C+%22Chaos+Elemental%22%2C+%22Venenatis%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: http.StatusBadRequest,
					Body:       getJSON(t, "testdata/bad_request_error.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
				wg:     sync.WaitGroup{},
				mu:     sync.Mutex{},
			}
			return client, context.Background(), check(verifyError)
		},
		"other error": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/monsters?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Molanisk%22%2C+%22Aberrant+spectre%22%2C+%22Chaos+Elemental%22%2C+%22Venenatis%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: http.StatusInternalServerError,
					Body:       ioutil.NopCloser(bytes.NewBufferString("{}")),
				}, nil)

			client := &client{
				client: mockHTTPClient,
				wg:     sync.WaitGroup{},
				mu:     sync.Mutex{},
			}
			return client, context.Background(), check(verifyError)
		},
		"json error": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/monsters?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Molanisk%22%2C+%22Aberrant+spectre%22%2C+%22Chaos+Elemental%22%2C+%22Venenatis%22%5D+%7D%2C+%22duplicate%22%3A+false+%7D",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewBufferString("bad json")),
				}, nil)

			client := &client{
				client: mockHTTPClient,
			}
			return client, context.Background(), check(verifyError)
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client, ctx, checkFns := tc(t)

			if len(checkFns) == 0 {
				t.Skipf("Skipping %s because there are no checks in place", name)
			}

			monsters, err := client.GetMonstersWhere(ctx, `{ "name": { "$in": ["Molanisk", "Aberrant spectre", "Chaos Elemental", "Venenatis"] }, "duplicate": false }`)
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

	tests := map[string]func(t *testing.T) (*client, context.Context, []checkFn){
		"success": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://raw.githubusercontent.com/osrsbox/osrsbox-db/master/osrsbox/docs/prayers-complete.json",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       getJSON(t, "testdata/all_monsters.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
			}
			return client, context.Background(), check(verifyNoError, verifyPrayers)
		},
		"http error": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://raw.githubusercontent.com/osrsbox/osrsbox-db/master/osrsbox/docs/prayers-complete.json",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(nil, errors.New("http error"))

			client := &client{
				client: mockHTTPClient,
			}
			return client, context.Background(), check(verifyError)
		},
		"status not ok": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://raw.githubusercontent.com/osrsbox/osrsbox-db/master/osrsbox/docs/prayers-complete.json",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: http.StatusInternalServerError,
					Body:       ioutil.NopCloser(bytes.NewBufferString("{}")),
				}, nil)

			client := &client{
				client: mockHTTPClient,
			}
			return client, context.Background(), check(verifyError)
		},
		"json error": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://raw.githubusercontent.com/osrsbox/osrsbox-db/master/osrsbox/docs/prayers-complete.json",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewBufferString("bad json")),
				}, nil)

			client := &client{
				client: mockHTTPClient,
			}
			return client, context.Background(), check(verifyError)
		},
		"nil client": func(t *testing.T) (*client, context.Context, []checkFn) {
			client := &client{
				client: nil,
			}
			return client, context.Background(), check(verifyError)
		},
		"no context": func(t *testing.T) (*client, context.Context, []checkFn) {
			client := &client{
				client: &http.Client{},
			}
			return client, nil, check(verifyError)
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client, ctx, checkFns := tc(t)

			if len(checkFns) == 0 {
				t.Skipf("Skipping %s because there are no checks in place", name)
			}

			prayers, err := client.GetAllPrayers(ctx)

			for _, checkFn := range checkFns {
				checkFn(t, prayers, err)
			}

		})
	}
}

func Test_GetPrayersByName(t *testing.T) {
	type checkFn func(*testing.T, []*Prayer, error)
	check := func(fns ...checkFn) []checkFn { return fns }

	verifyPrayerNames := func(t *testing.T, prayers []*Prayer, err error) {
		if len(prayers) != 4 {
			t.Errorf("expected four prayers, got %d length slice", len(prayers))
		}

		for i := range prayers {
			if i == 0 && prayers[i].Name != "Thick Skin" {
				t.Errorf("expected Thick Skin, got %s", prayers[i].Name)
			}

			if i == 1 && prayers[i].Name != "Burst of Strength" {
				t.Errorf("expected Burst of Strength, got %s", prayers[i].Name)
			}

			if i == 2 && prayers[i].Name != "Smite" {
				t.Errorf("expected Smite, got %s", prayers[i].Name)
			}

			if i == 3 && prayers[i].Name != "Rigour" {
				t.Errorf("expected Rigour, got %s", prayers[i].Name)
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

	tests := map[string]func(t *testing.T) (*client, context.Context, []checkFn){
		"success": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			firstReq, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/prayers?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Thick+Skin%22%2C+%22Burst+of+Strength%22%2C+%22Smite%22%2C+%22Rigour%22%5D+%7D+%7D",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			secondReq, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/prayers?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Thick+Skin%22%2C+%22Burst+of+Strength%22%2C+%22Smite%22%2C+%22Rigour%22%5D+%7D+%7D&page=2",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(firstReq).
				Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       getJSON(t, "testdata/where_prayers_page1.json"),
				}, nil)

			mockHTTPClient.EXPECT().
				Do(secondReq).
				Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       getJSON(t, "testdata/where_prayers_page2.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
			}
			return client, context.Background(), check(verifyNoError, verifyPrayerNames)
		},
		"nil client": func(t *testing.T) (*client, context.Context, []checkFn) {
			client := &client{
				client: nil,
			}
			return client, context.Background(), check(verifyError)
		},
		"no context": func(t *testing.T) (*client, context.Context, []checkFn) {
			client := &client{
				client: &http.Client{},
			}
			return client, nil, check(verifyError)
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client, ctx, checkFns := tc(t)

			if len(checkFns) == 0 {
				t.Skipf("Skipping %s because there are no checks in place", name)
			}

			prayers, err := client.GetPrayersByName(ctx, "Thick Skin", "Burst of Strength", "Smite", "Rigour")

			for _, checkFn := range checkFns {
				checkFn(t, prayers, err)
			}
		})
	}
}

func Test_GetPrayersWhere(t *testing.T) {
	type checkFn func(*testing.T, []*Prayer, error)
	check := func(fns ...checkFn) []checkFn { return fns }

	verifyPrayerNames := func(t *testing.T, prayers []*Prayer, err error) {
		if len(prayers) != 4 {
			t.Errorf("expected four prayers, got %d length slice", len(prayers))
		}

		for i := range prayers {
			if i == 0 && prayers[i].Name != "Thick Skin" {
				t.Errorf("expected Thick Skin, got %s", prayers[i].Name)
			}

			if i == 1 && prayers[i].Name != "Burst of Strength" {
				t.Errorf("expected Burst of Strength, got %s", prayers[i].Name)
			}

			if i == 2 && prayers[i].Name != "Smite" {
				t.Errorf("expected Smite, got %s", prayers[i].Name)
			}

			if i == 3 && prayers[i].Name != "Rigour" {
				t.Errorf("expected Rigour, got %s", prayers[i].Name)
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

	tests := map[string]func(t *testing.T) (*client, context.Context, []checkFn){
		"success": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			firstReq, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/prayers?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Thick+Skin%22%2C+%22Burst+of+Strength%22%2C+%22Smite%22%2C+%22Rigour%22%5D+%7D+%7D",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			secondReq, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/prayers?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Thick+Skin%22%2C+%22Burst+of+Strength%22%2C+%22Smite%22%2C+%22Rigour%22%5D+%7D+%7D&page=2",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(firstReq).
				Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       getJSON(t, "testdata/where_prayers_page1.json"),
				}, nil)

			mockHTTPClient.EXPECT().
				Do(secondReq).
				Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       getJSON(t, "testdata/where_prayers_page2.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
			}
			return client, context.Background(), check(verifyNoError, verifyPrayerNames)
		},
		"success one page": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			firstReq, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/prayers?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Thick+Skin%22%2C+%22Burst+of+Strength%22%2C+%22Smite%22%2C+%22Rigour%22%5D+%7D+%7D",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(firstReq).
				Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       getJSON(t, "testdata/where_prayers_one_page.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
			}
			return client, context.Background(), check(verifyNoError, verifyPrayerNames)
		},
		"nil client": func(t *testing.T) (*client, context.Context, []checkFn) {
			client := &client{
				client: nil,
			}
			return client, context.Background(), check(verifyError)
		},
		"no context": func(t *testing.T) (*client, context.Context, []checkFn) {
			client := &client{
				client: &http.Client{},
			}
			return client, nil, check(verifyError)
		},
		"request error": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/prayers?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Thick+Skin%22%2C+%22Burst+of+Strength%22%2C+%22Smite%22%2C+%22Rigour%22%5D+%7D+%7D",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(nil, errors.New("http error"))

			client := &client{
				client: mockHTTPClient,
			}
			return client, context.Background(), check(verifyError)
		},
		"second request error": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			firstReq, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/prayers?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Thick+Skin%22%2C+%22Burst+of+Strength%22%2C+%22Smite%22%2C+%22Rigour%22%5D+%7D+%7D",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			secondReq, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/prayers?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Thick+Skin%22%2C+%22Burst+of+Strength%22%2C+%22Smite%22%2C+%22Rigour%22%5D+%7D+%7D&page=2",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(firstReq).
				Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       getJSON(t, "testdata/where_prayers_page1.json"),
				}, nil)

			mockHTTPClient.EXPECT().
				Do(secondReq).
				Return(nil, errors.New("http error"))

			client := &client{
				client: mockHTTPClient,
			}
			return client, context.Background(), check(verifyError)
		},
		"formatted error": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/prayers?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Thick+Skin%22%2C+%22Burst+of+Strength%22%2C+%22Smite%22%2C+%22Rigour%22%5D+%7D+%7D",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: http.StatusBadRequest,
					Body:       getJSON(t, "testdata/bad_request_error.json"),
				}, nil)

			client := &client{
				client: mockHTTPClient,
				wg:     sync.WaitGroup{},
				mu:     sync.Mutex{},
			}
			return client, context.Background(), check(verifyError)
		},
		"other error": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/prayers?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Thick+Skin%22%2C+%22Burst+of+Strength%22%2C+%22Smite%22%2C+%22Rigour%22%5D+%7D+%7D",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: http.StatusInternalServerError,
					Body:       ioutil.NopCloser(bytes.NewBufferString("{}")),
				}, nil)

			client := &client{
				client: mockHTTPClient,
				wg:     sync.WaitGroup{},
				mu:     sync.Mutex{},
			}
			return client, context.Background(), check(verifyError)
		},
		"json error": func(t *testing.T) (*client, context.Context, []checkFn) {
			ctrl := gomock.NewController(t)

			mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

			req, err := http.NewRequestWithContext(
				context.Background(),
				"GET",
				"https://api.osrsbox.com/prayers?where=%7B+%22name%22%3A+%7B+%22%24in%22%3A+%5B%22Thick+Skin%22%2C+%22Burst+of+Strength%22%2C+%22Smite%22%2C+%22Rigour%22%5D+%7D+%7D",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			mockHTTPClient.EXPECT().
				Do(req).
				Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewBufferString("bad json")),
				}, nil)

			client := &client{
				client: mockHTTPClient,
			}
			return client, context.Background(), check(verifyError)
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client, ctx, checkFns := tc(t)

			if len(checkFns) == 0 {
				t.Skipf("Skipping %s because there are no checks in place", name)
			}

			prayers, err := client.GetPrayersWhere(ctx, `{ "name": { "$in": ["Thick Skin", "Burst of Strength", "Smite", "Rigour"] } }`)

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

func inDrops(names []string, drops []*Drop) bool {
	for _, drop := range drops {
		for _, name := range names {
			if name == drop.Name {
				return true
			}
		}
	}
	return false
}