package main

import (
	"context"
	"fmt"
	"log"

	"github.com/atye/gosrsbox"
	"github.com/atye/gosrsbox/osrsbox"
	"github.com/atye/gosrsbox/sets"
)

func main() {
	api := gosrsbox.NewAPI("")

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

	prayers, err := api.GetPrayersByName(context.Background(), "Thick Skin")
	if err != nil {
		log.Fatal(err)
	}
	printPrayers(prayers)

	var out map[string]interface{}
	err = api.GetDocument(context.Background(), "items-json/0.json", &out)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", out["wiki_name"])
}

func printItems(items []osrsbox.Item) {
	for _, item := range items {
		fmt.Println(item.WikiName)
	}
}

func printMonsters(monsters []osrsbox.Monster) {
	for _, monster := range monsters {
		fmt.Println(monster.WikiName)
	}
}

func printPrayers(prayers []osrsbox.Prayer) {
	for _, prayer := range prayers {
		fmt.Println(prayer.Name)
	}
}
