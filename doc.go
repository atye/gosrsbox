/*
Package osrsbox provides a wrapper for osrsbox-api (https://api.osrsbox.com).

An InMemoryAPI downloads all entities from the osrsbox-db project at once and stores them in memory in JSON. gjson is used to query the InMemoryAPI.

An ExternalAPI uses the default http client to make all queries to osrsbox-api.

Example InMemoryAPI:

	log.Println("Creating in-memory client...")
	api, err := osrsbox.NewInMemoryClient(nil)
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


Example ExternalAPI:

	api := osrsbox.NewAPIClient(nil)

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
}
*/
package osrsbox
