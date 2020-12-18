package main

import (
	"context"
	"log"

	"github.com/atye/gosrsbox/osrsbox"
	armorsets "github.com/atye/gosrsbox/osrsbox/db/armor-sets"
)

func main() {
	log.Println("Creating in-memory client...")
	api, err := osrsbox.NewInMemoryClient()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Done")
	log.Println("Getting set...")
	items, err := api.GetItemsByName(context.Background(), armorsets.RuneGoldTrimmedLg...)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Done")
	for _, item := range items {
		log.Println(item.Name)
	}

	log.Println("Getting set...")
	items, err = api.GetItemsByName(context.Background(), armorsets.RuneGoldTrimmedLg...)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Done")
	for _, item := range items {
		log.Println(item.Name)
	}

	log.Println("Creating api client...")
	networkAPI := osrsbox.NewAPIClient()
	log.Println("Getting set...")
	items, err = networkAPI.GetItemsByName(context.Background(), armorsets.DragonLg...)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Done")
	for _, item := range items {
		log.Println(item.Name)
	}

	log.Println("Getting set...")
	items, err = networkAPI.GetItemsByName(context.Background(), armorsets.SteelGoldTrimmedLg...)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Done")
	for _, item := range items {
		log.Println(item.Name)
	}

}
