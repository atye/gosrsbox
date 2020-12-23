package main

import (
	"context"
	"log"

	"github.com/atye/gosrsbox/osrsboxapi/api"
	"github.com/atye/gosrsbox/osrsboxapi/sets"
	openapi "github.com/atye/gosrsbox/pkg/openapi/api"
)

func main() {
	//Create api client using http.DefaultClient
	api := api.NewAPI(nil)

	// Get slice of items in the Ahrims set
	items, err := api.GetItemSet(context.Background(), sets.Ahrims)
	if err != nil {
		log.Fatal(err)
	}
	printItems(items)

	// Get slice of monsters that drop the Bandos chestplate
	monsters, err := api.GetMonstersThatDrop(context.Background(), "Bandos chestplate")
	if err != nil {
		log.Fatal(err)
	}
	printMonsters(monsters)

	// Get items with negative prayer bonus using Python query
	_, err = api.GetPrayersByQuery(context.Background(), "equipment.prayer<0")
	if err != nil {
		log.Fatal(err)
	}
	printItems(items)

	// Get items with negative prayer bonus using MongoDB query
	items, err = api.GetItemsByQuery(context.Background(), `{ "equipment.prayer": { "$lt": 0 }, "duplicate": false }`)
	if err != nil {
		log.Fatal(err)
	}
	printItems(items)

	// Get prayers with a faulty MongoDB query that returns an error
	_, err = api.GetPrayersByQuery(context.Background(), `{"name":{"$nin":"test"}}`)
	if err != nil {
		log.Fatal(err)
	}
	printItems(items)
}

func printItems(items []openapi.Item) {
	for _, item := range items {
		log.Println(item.GetWikiName())
	}
}

func printMonsters(monsters []openapi.Monster) {
	for _, monster := range monsters {
		log.Println(monster.GetWikiName())
	}
}
