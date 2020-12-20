
go-osrsbox is a Go client library for [osrsbox-api](https://api.osrsbox.com).

## RESTFul Client
"Internet-accessible API with rich-quering including filtering, sorting and projection functionality"

```go get github.com/atye/gosrsbox/osrsboxdb/api/restful```

The `restful` package provides a client for accessing [osrsbox-api](https://api.osrsbox.com).  See [examples](./examples/restful) for usage.

#### Features
 - all calls are made with an `http.Client` to [osrsbox-api](https://api.osrsbox.com)
   
 - no more than 10 concurrent http calls (for now)
   
  - supports MongoDB and Python queries as documented on [osrsbox-api](https://api.osrsbox.com)
 
```
type  API  interface {
	GetItemsByName(ctx context.Context, names ...string) ([]osrsboxdb.Item, error)
	GetItemsByQuery(ctx context.Context, query string) ([]osrsboxdb.Item, error)
	GetItemSet(ctx context.Context, set sets.SetName) ([]osrsboxdb.Item, error)
	GetMonstersByName(ctx context.Context, names ...string) ([]osrsboxdb.Monster, error)
	GetMonstersByQuery(ctx context.Context, query string) ([]osrsboxdb.Monster, error)
	GetMonstersThatDrop(ctx context.Context, items ...string) ([]osrsboxdb.Monster, error)
	GetPrayersByName(ctx context.Context, names ...string) ([]osrsboxdb.Prayer, error)
	GetPrayersByQuery(ctx context.Context, query string) ([]osrsboxdb.Prayer, error)
}
```
```
import (
    "github.com/atye/gosrsbox/osrsboxdb/api/restful"
)

// api client
api := restful.NewAPI(nil)

// Get slice of items in the Third Age Range Kit
items, err := api.GetItemSet(context.Background(), sets.ThirdAgeRangeKit)

// Get slice of monsters that drop the Bandos chestplate
monsters, err := api.GetMonstersThatDrop(context.Background(), "Bandos chestplate")

// Get slice of items with negative prayer bonus using Python query
items, err = api.GetItemsByQuery(context.Background(), "equipment.prayer<0")

// Get slice of items with negative prayer bonus using MongoDB query
items, err = api.GetItemsByQuery(context.Background(), `{ "equipment.prayer": { "$lt": 0 }, "duplicate": false }`)