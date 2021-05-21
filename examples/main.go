package main

import (
	"context"
	"log"

	"github.com/atye/gosrsbox"
	"github.com/atye/gosrsbox/internal/api"
	"github.com/atye/gosrsbox/sets"
)

func main() {
	//Create api client using http.DefaultClient, disable logging
	api := gosrsbox.NewAPI(gosrsbox.APIConfig{})

	// Get slice of items in the Ahrims set
	items, err := api.GetItemSet(context.Background(), sets.Ahrims)
	if err != nil {
		log.Fatal(err)
	}
	printItems(items)

	// Get slice of monsters that drop the Bandos chestplate
	_, err = api.GetMonstersThatDrop(context.Background(), "Bandos chestplate")
	if err != nil {
		log.Fatal(err)
	}
	//printMonsters(monsters)

	// Get items with negative prayer bonus using MongoDB query
	items, err = api.GetItemsByQuery(context.Background(), `{ "equipment.prayer": { "$lt": 0 }, "duplicate": false }`)
	if err != nil {
		log.Fatal(err)
	}
	printItems(items)
}

func printItems(items []api.Item) {
	for _, item := range items {
		log.Println(item.WikiName)
	}
}

/*func printMonsters(monsters []openapi.Monster) {
	for _, monster := range monsters {
		log.Println(monster.GetWikiName())
	}
}*/
