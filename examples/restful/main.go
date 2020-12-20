package main

import (
	"context"
	"log"

	"github.com/atye/gosrsbox/osrsboxapi"
	"github.com/atye/gosrsbox/osrsboxapi/api/restful"
	"github.com/atye/gosrsbox/osrsboxapi/sets"
)

func main() {
	api := restful.NewAPI(nil)
	items, err := api.GetItemSet(context.Background(), sets.ThirdAgeRangeKit)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(items)

	monsters, err := api.GetMonstersThatDrop(context.Background(), "Bandos chestplate")
	if err != nil {
		log.Fatal(err)
	}
	printMonsters(monsters)

	items, err = api.GetItemsByQuery(context.Background(), "equipment.prayer<0")
	if err != nil {
		log.Fatal(err)
	}
	printItems(items)

	items, err = api.GetItemsByQuery(context.Background(), `{ "equipment.prayer": { "$lt": 0 }, "duplicate": false }`)
	if err != nil {
		log.Fatal(err)
	}
	printItems(items)
}

func printItems(items []osrsboxapi.Item) {
	for _, item := range items {
		log.Println(item.WikiName)
	}
}

func printMonsters(monsters []osrsboxapi.Monster) {
	for _, monster := range monsters {
		log.Println(monster.WikiName)
	}
}
