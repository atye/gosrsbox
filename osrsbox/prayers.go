package osrsbox

import (
	"context"
	"encoding/json"
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

func getAllPrayers(ctx context.Context, c *client) ([]*Prayer, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", prayersCompleteURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error doing request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error from server: %w", &serverError{
			Code:    resp.StatusCode,
			Message: "something went wrong",
		})
	}

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

func getPrayersByName(ctx context.Context, c *client, names ...string) ([]*Prayer, error) {
	var nameData []string
	var query string

	for _, name := range names {
		nameData = append(nameData, fmt.Sprintf(`"%s"`, name))
	}
	query = fmt.Sprintf(`{ "name": { "$in": [%s] } }`, strings.Join(nameData, ", "))

	return getPrayersWhere(ctx, c, query)
}

func getPrayersWhere(ctx context.Context, c *client, query string) ([]*Prayer, error) {
	url := fmt.Sprintf("%s/%s?where=%s", api, prayersEndpoint, url.QueryEscape(query))

	prayersResp, err := doPrayersRespRequest(ctx, c, url)
	if err != nil {
		return nil, err
	}

	prayers := make([]*Prayer, prayersResp.Meta.Total)

	for i, prayer := range prayersResp.Prayers {
		prayers[i] = prayer
	}

	var pages int
	if prayersResp.Meta.MaxResults != 0 {
		pages = int(math.Ceil(float64(prayersResp.Meta.Total) / float64(prayersResp.Meta.MaxResults)))
	}

	if pages > 1 {

		errChan := make(chan error)
		waitChan := make(chan struct{})

		go func() {
			for page := 2; page <= pages; page++ {
				c.wg.Add(1)
				go func(page int) {
					defer c.wg.Done()

					prayersRespTemp, err := doPrayersRespRequest(ctx, c, fmt.Sprintf("%s&page=%d", url, page))
					if err != nil {
						errChan <- err
						return
					}
					for i, prayer := range prayersRespTemp.Prayers {
						c.mu.Lock()
						prayers[prayersRespTemp.Meta.MaxResults*(page-1)+i] = prayer
						c.mu.Unlock()
					}

				}(page)
			}
			c.wg.Wait()
			close(waitChan)
		}()

		select {
		case <-waitChan:
			return prayers, nil
		case err := <-errChan:
			return nil, err
		}
	}

	return prayers, nil
}

func doPrayersRespRequest(ctx context.Context, c *client, url string) (*prayersResponse, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error doing request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var prayersResp *prayersResponse
		_ = json.NewDecoder(resp.Body).Decode(&prayersResp)
		defer resp.Body.Close()

		if prayersResp != nil && prayersResp.Error != nil {
			return nil, fmt.Errorf("error from server: %w", prayersResp.Error)
		}

		return nil, fmt.Errorf("error some server: %w", &serverError{
			Code:    resp.StatusCode,
			Message: "something went wrong",
		})
	}

	var prayersResp *prayersResponse
	err = json.NewDecoder(resp.Body).Decode(&prayersResp)
	if err != nil {
		return nil, fmt.Errorf("error decoding json response: %w", err)
	}
	defer resp.Body.Close()

	return prayersResp, nil
}
