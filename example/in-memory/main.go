package main

/*import (
	"context"
	"log"

	"github.com/atye/gosrsbox/osrsbox"
	"github.com/atye/gosrsbox/osrsbox/db"
	"github.com/atye/gosrsbox/osrsbox/db/sets"
)

func main() {
	log.Println("Creating in-memory client...")
	api, err := osrsbox.NewInMemoryAPI(nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Done")

	items, err := api.GetItemSet(context.Background(), sets.RuneGoldTrimmedLg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(len(items))

	monsters, err := api.GetMonstersThatDrop(context.Background(), "Bandos chestplate")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(len(monsters))

	items, err = api.GetItemSet(context.Background(), sets.PartyHats)
	if err != nil {
		log.Fatal(err)
	}
	printItems(items)
}

func printItems(items []db.Item) {
	for _, item := range items {
		log.Println(item.Name)
	}
}

func printMonsters(monsters []db.Monster) {
	for _, monster := range monsters {
		log.Println(monster.Name)
	}
}*/
