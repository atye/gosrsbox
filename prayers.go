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

// Prayer is an osrsbox prayer.
// https://api.osrsbox.com/swaggerui#/Prayer/get_prayers
type Prayer struct {
	ID             interface{}    `json:"id"`
	Name           string         `json:"name"`
	Members        bool           `json:"members"`
	Description    string         `json:"description"`
	DrainPerMinute float32        `json:"drain_per_minute"`
	WikiURL        string         `json:"wiki_url"`
	Requirements   map[string]int `json:"requirements"`
	Bonuses        map[string]int `json:"bonuses"`
	Icon           string         `json:"icon"`
}

type prayersResponse struct {
	Prayers []*Prayer `json:"_items"`
	Meta    struct {
		Page       int `json:"page"`
		MaxResults int `json:"max_results"`
		Total      int `json:"total"`
	} `json:"_meta"`
	Error *serverError `json:"_error"`
}

func getAllPrayers(ctx context.Context, client HTTPClient) ([]*Prayer, error) {
	if client == nil {
		return nil, errors.New("no client configured")
	}

	req, err := http.NewRequestWithContext(ctx, "GET", prayersCompleteURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error doing request: %w", err)
	}
	defer resp.Body.Close()

	var prayersMap map[string]*Prayer
	err = json.NewDecoder(resp.Body).Decode(&prayersMap)
	if err != nil {
		return nil, fmt.Errorf("error decoding json response: %w", err)
	}

	var prayers []*Prayer
	for _, prayer := range prayersMap {
		prayers = append(prayers, prayer)
	}

	return prayers, nil
}

func getPrayersByName(ctx context.Context, client HTTPClient, names ...string) ([]*Prayer, error) {
	if client == nil {
		return nil, errors.New("no client configured")
	}

	var nameData []string
	var query string

	for _, name := range names {
		nameData = append(nameData, fmt.Sprintf(`"%s"`, name))
	}
	query = fmt.Sprintf(`{ "name": { "$in": [%s] } }`, strings.Join(nameData, ", "))

	return getPrayersWhere(ctx, client, query)
}

func getPrayersWhere(ctx context.Context, client HTTPClient, query string) ([]*Prayer, error) {
	if client == nil {
		return nil, errors.New("no client configured")
	}

	url := fmt.Sprintf("%s/%s?where=%s", api, prayersEndpoint, url.QueryEscape(query))

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error doing request: %w", err)
	}

	var prayers []*Prayer

	var prayersResp *prayersResponse
	err = json.NewDecoder(resp.Body).Decode(&prayersResp)
	if err != nil {
		return nil, fmt.Errorf("error decoding json response: %w", err)
	}
	resp.Body.Close()

	if prayersResp.Error != nil {
		return nil, fmt.Errorf("error from server: %w", prayersResp.Error)
	}

	prayers = append(prayers, prayersResp.Prayers...)

	var pages int
	if prayersResp.Meta.MaxResults != 0 {
		pages = int(math.Ceil(float64(prayersResp.Meta.Total) / float64(prayersResp.Meta.MaxResults)))
	}

	for i := 2; i <= pages; i++ {
		req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s&page=%d", url, i), nil)
		if err != nil {
			return nil, fmt.Errorf("error creating request: %w", err)
		}

		resp, err := client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("error doing request: %w", err)
		}

		var prayersRespTemp *prayersResponse
		err = json.NewDecoder(resp.Body).Decode(&prayersRespTemp)
		if err != nil {
			return nil, fmt.Errorf("error decoding json response: %w", err)
		}
		resp.Body.Close()

		if prayersRespTemp.Error != nil {
			return nil, fmt.Errorf("error from server: %w", prayersRespTemp.Error)
		}

		prayers = append(prayers, prayersRespTemp.Prayers...)
	}

	return prayers, nil
}
