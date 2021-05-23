[gosrsbox](https://godoc.org/github.com/atye/gosrsbox/osrsboxapi) is a client library for [osrsbox-api](https://api.osrsbox.com) utilizing OpenAPI automation.
# Installing
```go get github.com/atye/gosrsbox```

# Creating a client
```
api := gosrsbox.NewAPI("my user agent")
```
#### Features
- get Items, Monsters, and Prayers by id, wiki name, built-in options, and custom queries
- supports MongoDB and Python queries as documented on [osrsbox-api](https://api.osrsbox.com)

### See examples folder for examples.
