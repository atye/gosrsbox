package main

import (
	"context"
	"log"

	"github.com/atye/gosrsbox/osrsboxapi"
	"github.com/atye/gosrsbox/osrsboxapi/api"
	"github.com/atye/gosrsbox/osrsboxapi/sets"
)

func main() {
	//Create api client using http.DefaultClient
	api := api.NewAPI(nil)

	set, err := api.GetItemSet(context.Background(), sets.Ahrims)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(set)

	/*var dataSet map[string]osrsboxapi.Item
	err := api.GetJSONFiles(context.Background(), []string{"items-json"}, &dataSet)
	if err != nil {
		log.Fatal(err)
	}

	// Get slice of items in the Third Age Range Kit
	items, err := api.GetItemSet(context.Background(), sets.ThirdAgeRangeKit)
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
	items, err = api.GetItemsByQuery(context.Background(), "equipment.prayer<0")
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

	// Get JSON file from STATIC JSON API as an osrsboxapi.Item
	var excalibur osrsboxapi.Item
	err = api.GetJSONFiles(context.Background(), []string{"items-json/35.json"}, &excalibur)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(excalibur.Name)

	var twoHandedITems map[string]osrsboxapi.Item
	err = api.GetJSONFiles(context.Background(), []string{"items-json-slot/items-2h.json"}, &twoHandedITems)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(len(twoHandedITems))

	// Gather npcs-summary.json data which doesn't contain Items, Monsters, or Prayers data
	// Create my own custom struct and variable to unmarshal npcs-summary.json
	type NPCSummary struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	var data map[string]NPCSummary

	// Pass variable to GetJSONFiles to use it as an unmarshal destination
	// Get multiple JSON datasets concurrently
	var cannonBall osrsboxapi.Item
	err = api.GetJSONFiles(context.Background(), []string{"npcs-summary.json", "items-json/2.json"}, &data, &cannonBall)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(data["2"].Name)
	log.Println(cannonBall.WikiName)*/
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
