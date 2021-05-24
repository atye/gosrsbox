package main

import (
	"context"
	"fmt"
	"log"

	"github.com/atye/gosrsbox"
	"github.com/atye/gosrsbox/openapi/api"
	"github.com/atye/gosrsbox/sets"
	"github.com/atye/gosrsbox/slots"
)

func main() {
	api := gosrsbox.NewAPI("my user agent")

	// Get items in the Ahrims set
	items, err := api.GetItemSet(context.Background(), sets.Ahrims)
	if err != nil {
		log.Fatal(err)
	}
	printItems(items)

	// Get items in the Hands slot
	items, err = api.GetItemsBySlot(context.Background(), slots.Hands)
	if err != nil {
		log.Fatal(err)
	}
	printItems(items)

	// Get monsters that drop the Bandos chestplate
	monsters, err := api.GetMonstersThatDrop(context.Background(), "Bandos chestplate")
	if err != nil {
		log.Fatal(err)
	}
	printMonsters(monsters)

	// Get items with negative prayer bonus using MongoDB query
	items, err = api.GetItemsByQuery(context.Background(), `{ "equipment.prayer": { "$lt": 0 }, "duplicate": false }`)
	if err != nil {
		log.Fatal(err)
	}
	printItems(items)

	// Get items with negative prayer bonus using Python query
	items, err = api.GetItemsByQuery(context.Background(), `equipment.prayer<0`)
	if err != nil {
		log.Fatal(err)
	}
	printItems(items)

	// Get the Thick Skin prayer
	prayers, err := api.GetPrayersByName(context.Background(), "Thick Skin")
	if err != nil {
		log.Fatal(err)
	}
	printPrayers(prayers)

	// Get the 0th item from the Static JSON API
	var out map[string]interface{}
	err = api.GetDocument(context.Background(), "items-json/0.json", &out)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", out["wiki_name"])
}

func printItems(items []api.Item) {
	for _, item := range items {
		fmt.Println(item.GetWikiName())
	}
}

func printMonsters(monsters []api.Monster) {
	for _, monster := range monsters {
		fmt.Println(monster.GetWikiName())
	}
}

func printPrayers(prayers []api.Prayer) {
	for _, prayer := range prayers {
		fmt.Println(prayer.GetName())
	}
}
