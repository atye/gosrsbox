/*
Package osrsbox provides a client for the osrsbox-api (https://api.osrsbox.com).

The API has /equipment and /weapons endpoints but those datasets are subsets of /items. So as far as this client is concerned, those entities are items.

See the example below to find equipment with prayer bonuses of less than 0.

Example:

	// New osrsbox client
	client := osrsbox.New(&http.Client{})

	//Get items by name
	items, err := client.GetItemsByName(context.TODO(), "Dragon scimitar", "Rune platebody")
	if err != nil {
		//
	}

	for _, item := range items {
		fmt.Printf("%s: %d\n", item.Name, item.Highalch)
	}

	//Find items with a prayer bonus of less than 0 with a MongoDB query
	items, err = client.GetItemsWhere(context.TODO(), `{ "equipment.prayer": { "$lt": 0 }, "duplicate": false }`)
	if err != nil {
		//
	}

	for _, item := range items {
		fmt.Printf("%s: %d\n", item.Name, item.Equipment.Prayer)
	}

	//Find monsters that drop the Smoke battlestaff
	monsters, err := client.GetMonstersThatDrop(context.TODO(), "Smoke battlestaff")
	if err != nil {
		//
	}

	for _, monster := range monsters {
		fmt.Printf("%s\n", monster.Name)
	}


*/
package osrsbox
