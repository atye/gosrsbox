package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/atye/gosrsbox/osrsboxdb"
)

type httpGet func(string) (*http.Response, error)

const (
	itemsCompleteURL    = "https://raw.githubusercontent.com/osrsbox/osrsbox-db/master/osrsbox/docs/items-complete.json"
	monstersCompleteURL = "https://raw.githubusercontent.com/osrsbox/osrsbox-db/master/osrsbox/docs/monsters-complete.json"
	prayersCompleteURL  = "https://raw.githubusercontent.com/osrsbox/osrsbox-db/master/osrsbox/docs/prayers-complete.json"
)

func FromHttpClient(c *http.Client) httpGet {
	return httpGet(c.Get)
}

func (f httpGet) Items() ([]byte, error) {
	resp, err := f(itemsCompleteURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//itemMap can be in any order
	var itemMap map[string]osrsboxdb.Item
	err = json.Unmarshal(body, &itemMap)
	if err != nil {
		return nil, err
	}

	itemSlice := make([]osrsboxdb.Item, 0, len(itemMap))
	for _, item := range itemMap {
		itemSlice = append(itemSlice, item)
	}

	items, err := json.Marshal(itemSlice)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (f httpGet) Monsters() ([]byte, error) {
	resp, err := f(monstersCompleteURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var monstersMap map[string]osrsboxdb.Monster
	err = json.Unmarshal(body, &monstersMap)
	if err != nil {
		return nil, err
	}

	monstersSlice := make([]osrsboxdb.Monster, 0, len(monstersMap))
	for _, monster := range monstersMap {
		monstersSlice = append(monstersSlice, monster)

	}

	monsters, err := json.Marshal(monstersSlice)
	if err != nil {
		return nil, err
	}

	return monsters, nil
}

func (f httpGet) Prayers() ([]byte, error) {
	resp, err := f(prayersCompleteURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var prayersMap map[string]osrsboxdb.Prayer
	err = json.Unmarshal(body, &prayersMap)
	if err != nil {
		return nil, err
	}

	prayersSlice := make([]osrsboxdb.Prayer, 0, len(prayersMap))
	for _, prayer := range prayersMap {
		prayersSlice = append(prayersSlice, prayer)
	}

	prayers, err := json.Marshal(prayersSlice)
	if err != nil {
		return nil, err
	}

	return prayers, nil
}
