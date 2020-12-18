package update

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/atye/gosrsbox/osrsbox/db"
)

type HttpGet func(string) (*http.Response, error)

const (
	itemsCompleteURL    = "https://raw.githubusercontent.com/osrsbox/osrsbox-db/master/osrsbox/docs/items-complete.json"
	monstersCompleteURL = "https://raw.githubusercontent.com/osrsbox/osrsbox-db/master/osrsbox/docs/monsters-complete.json"
	prayersCompleteURL  = "https://raw.githubusercontent.com/osrsbox/osrsbox-db/master/osrsbox/docs/prayers-complete.json"
)

func (f HttpGet) Items() ([]byte, error) {
	resp, err := f(itemsCompleteURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var itemMap map[string]db.Item
	err = json.Unmarshal(body, &itemMap)
	if err != nil {
		return nil, err
	}

	itemSlice := make([]db.Item, 0, len(itemMap))
	for _, item := range itemMap {
		itemSlice = append(itemSlice, item)
	}

	items, err := json.Marshal(itemSlice)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (f HttpGet) Monsters() ([]byte, error) {
	resp, err := f(monstersCompleteURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var monstersMap map[string]db.Monster
	err = json.Unmarshal(body, &monstersMap)
	if err != nil {
		return nil, err
	}

	monstersSlice := make([]db.Monster, 0, len(monstersMap))
	for _, monster := range monstersMap {
		monstersSlice = append(monstersSlice, monster)

	}

	monsters, err := json.Marshal(monstersSlice)
	if err != nil {
		return nil, err
	}

	return monsters, nil
}

func (f HttpGet) Prayers() ([]byte, error) {
	resp, err := f(prayersCompleteURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var prayersMap map[string]db.Prayer
	err = json.Unmarshal(body, &prayersMap)
	if err != nil {
		return nil, err
	}

	prayersSlice := make([]db.Prayer, 0, len(prayersMap))
	for _, prayer := range prayersMap {
		prayersSlice = append(prayersSlice, prayer)
	}

	prayers, err := json.Marshal(prayersSlice)
	if err != nil {
		return nil, err
	}

	return prayers, nil
}
