package main

import (
	"context"
	"fmt"
	"log"

	"github.com/atye/gosrsbox"
	"github.com/atye/gosrsbox/models"
	"github.com/atye/gosrsbox/slots"
)

func main() {
	// gosrsbox.WithTracing("http://localhost:9411/api/v2/spans")
	api := gosrsbox.NewAPI("my user agent", gosrsbox.WithTracing("http://localhost:9411/api/v2/spans"))

	// Get items in the Ahrims set
	items, err := api.GetItemsByQuery(context.Background(), `{sgesg}`)
	if err != nil {
		log.Println(err)
	}
	//select {}
	printItems(items)

	// Get items in the Hands slot
	items, err = api.GetItemsBySlot(context.Background(), slots.Ammo)
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

	select {}
}

func printItems(items []models.Item) {
	for _, item := range items {
		fmt.Println(item.WikiName)
	}
}

func printMonsters(monsters []models.Monster) {
	for _, monster := range monsters {
		fmt.Println(monster.WikiName)
	}
}

func printPrayers(prayers []models.Prayer) {
	for _, prayer := range prayers {
		fmt.Println(prayer.Name)
	}
}
