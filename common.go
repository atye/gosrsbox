package gosrsbox

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strings"
)

func getAll(ctx context.Context, client HTTPClient, entity string) (interface{}, error) {
	if client == nil {
		return nil, errors.New("no client configured")
	}

	url := strings.ReplaceAll(completeURL, "replaceme", entity)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	switch entity {
	case "items":

		var itemsMap map[string]*Item
		err = json.NewDecoder(resp.Body).Decode(&itemsMap)
		if err != nil {
			return nil, err
		}

		var items []*Item
		for _, item := range itemsMap {
			items = append(items, item)
		}

		return items, nil

	case "monsters":

		var monstersMap map[string]*Monster
		err = json.NewDecoder(resp.Body).Decode(&monstersMap)
		if err != nil {
			return nil, err
		}

		var monsters []*Monster
		for _, monster := range monstersMap {
			monsters = append(monsters, monster)
		}

		return monsters, nil

	case "prayers":

		var prayersMap map[string]*Prayer
		err = json.NewDecoder(resp.Body).Decode(&prayersMap)
		if err != nil {
			return nil, err
		}

		var prayers []*Prayer
		for _, prayer := range prayersMap {
			prayers = append(prayers, prayer)
		}

		return prayers, nil

	default:
		return nil, fmt.Errorf("Entity %s not supported", entity)
	}
}

func getByName(ctx context.Context, client HTTPClient, endpoint string, names ...string) (interface{}, error) {
	if client == nil {
		return nil, errors.New("no client configured")
	}

	var nameData []string
	var query string

	switch endpoint {
	case "items":

		for _, name := range names {
			nameData = append(nameData, fmt.Sprintf(`"%s"`, makeValidItemName(name)))
		}
		query = fmt.Sprintf(`{ "name": { "$in": [%s] }, "duplicate": false }`, strings.Join(nameData, ", "))

	case "monsters":

		for _, name := range names {
			nameData = append(nameData, fmt.Sprintf(`"%s"`, name))
		}
		query = fmt.Sprintf(`{ "name": { "$in": [%s] }, "duplicate": false }`, strings.Join(nameData, ", "))

	case "prayers":

		for _, name := range names {
			nameData = append(nameData, fmt.Sprintf(`"%s"`, name))
		}
		query = fmt.Sprintf(`{ "name": { "$in": [%s] } }`, strings.Join(nameData, ", "))

	default:
		return nil, fmt.Errorf("Entity %s not supported", endpoint)
	}

	return getWhere(ctx, client, endpoint, query)

}

func getThatDrop(ctx context.Context, client HTTPClient, endpoint string, names ...string) (interface{}, error) {
	if client == nil {
		return nil, errors.New("no client configured")
	}

	var nameData []string
	for _, name := range names {
		nameData = append(nameData, fmt.Sprintf(`"%s"`, makeValidItemName(name)))
	}

	query := fmt.Sprintf(`{ "drops": { "$elemMatch": { "name": { "$in": [%s] } } }, "duplicate": false }`, strings.Join(nameData, ", "))

	return getWhere(ctx, client, endpoint, query)

}

func getWhere(ctx context.Context, client HTTPClient, endpoint, query string) (interface{}, error) {
	if client == nil {
		return nil, errors.New("no client configured")
	}

	url := fmt.Sprintf("%s/%s?where=%s", api, endpoint, url.QueryEscape(query))

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	switch endpoint {
	case "items":

		var items []*Item

		var itemsEnd *itemsEndpoint
		err = json.NewDecoder(resp.Body).Decode(&itemsEnd)
		if err != nil {
			return nil, err
		}
		resp.Body.Close()

		items = append(items, itemsEnd.Items...)

		var pages int
		if itemsEnd.Meta.MaxResults != 0 {
			pages = int(math.Ceil(float64(itemsEnd.Meta.Total) / float64(itemsEnd.Meta.MaxResults)))
		}

		for i := 2; i <= pages; i++ {
			req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s&page=%d", url, i), nil)
			if err != nil {
				return nil, err
			}

			resp, err := client.Do(req)
			if err != nil {
				return nil, err
			}

			var itemsEndTemp *itemsEndpoint
			err = json.NewDecoder(resp.Body).Decode(&itemsEndTemp)
			if err != nil {
				return nil, err
			}
			resp.Body.Close()

			items = append(items, itemsEndTemp.Items...)
		}

		return items, nil

	case "monsters":

		var monsters []*Monster

		var monstersEnd *monstersEndpoint
		err = json.NewDecoder(resp.Body).Decode(&monstersEnd)
		if err != nil {
			return nil, err
		}
		resp.Body.Close()

		monsters = append(monsters, monstersEnd.Monsters...)

		var pages int
		if monstersEnd.Meta.MaxResults != 0 {
			pages = int(math.Ceil(float64(monstersEnd.Meta.Total) / float64(monstersEnd.Meta.MaxResults)))
		}

		for i := 2; i <= pages; i++ {
			req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s&page=%d", url, i), nil)
			if err != nil {
				return nil, err
			}

			resp, err := client.Do(req)
			if err != nil {
				return nil, err
			}

			var monstersEndTemp *monstersEndpoint
			err = json.NewDecoder(resp.Body).Decode(&monstersEndTemp)
			if err != nil {
				return nil, err
			}
			resp.Body.Close()

			monsters = append(monsters, monstersEndTemp.Monsters...)
		}

		return monsters, nil

	case "prayers":

		var prayers []*Prayer

		var prayersEnd *prayersEndpoint
		err = json.NewDecoder(resp.Body).Decode(&prayersEnd)
		if err != nil {
			return nil, err
		}
		resp.Body.Close()

		prayers = append(prayers, prayersEnd.Prayers...)

		var pages int
		if prayersEnd.Meta.MaxResults != 0 {
			pages = int(math.Ceil(float64(prayersEnd.Meta.Total) / float64(prayersEnd.Meta.MaxResults)))
		}

		for i := 2; i <= pages; i++ {
			req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s&page=%d", url, i), nil)
			if err != nil {
				return nil, err
			}

			resp, err := client.Do(req)
			if err != nil {
				return nil, err
			}

			var prayersEndTemp *prayersEndpoint
			err = json.NewDecoder(resp.Body).Decode(&prayersEndTemp)
			if err != nil {
				return nil, err
			}
			resp.Body.Close()

			prayers = append(prayers, prayersEndTemp.Prayers...)
		}

		return prayers, nil

	default:
		return nil, fmt.Errorf("Entity %s not supported", endpoint)
	}
}

func makeValidItemName(name string) string {
	words := strings.Split(name, " ")

	if len(words) > 0 {
		words[0] = strings.Title(words[0])
		if len(words) > 1 {
			for i := 1; i < len(words); i++ {
				words[i] = strings.ToLower(words[i])
			}
		}
	}

	return strings.Join(words, " ")
}
