package main

import (
	"context"
	"log"

	"github.com/atye/gosrsbox/osrsbox"
	"github.com/atye/gosrsbox/osrsbox/db"
	"github.com/atye/gosrsbox/osrsbox/db/sets"
)

func main() {
	api := osrsbox.NewAPI(nil)
	api = osrsbox.NewAPI(nil)
	api = osrsbox.NewAPI(nil)
	_ = osrsbox.NewAPI(nil)
	monsters, err := api.GetItemSet(context.Background(), sets.ThirdAgeRangeKit)
	if err != nil {
		log.Fatal(err)
	}
	printItems(monsters)
	//log.Println(len(monsters))

	/*items, err = api.GetItemSet(context.Background(), sets.PartyHats)
	if err != nil {
		log.Fatal(err)
	}
	printItems(items)

	items, err = api.GetItemsByQuery(context.Background(), `{ "equipment.prayer": { "$lt": 0 }, "duplicate": false }`)
	if err != nil {
		log.Fatal(err)
	}
	printItems(items)*/
}

func printItems(items []db.Item) {
	for _, item := range items {
		log.Println(item.WikiName)
	}
}

func printMonsters(monsters []db.Monster) {
	for _, monster := range monsters {
		log.Println(monster.WikiName)
	}
}
