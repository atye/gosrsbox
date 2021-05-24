
[gosrsbox](https://godoc.org/github.com/atye/gosrsbox) is a client library for [osrsbox-api](https://api.osrsbox.com) utilizing OpenAPI automation.
# Installing

```go get github.com/atye/gosrsbox```
# Creating a client

```
api := gosrsbox.NewAPI("my user agent")
```
# Features (See examples folder for examples)
- get Items, Monsters, and Prayers by id, wiki name, built-in options, and custom queries
- supports MongoDB and Python queries as documented on [osrsbox-api](https://api.osrsbox.com)
- get any document from the [Static JSON API](https://www.osrsbox.com/projects/osrsbox-db/#the-osrsbox-static-json-api) (good for dumping the database)
# About
Item, Monster, and Prayer models are generated via a modified [OpenAPI specification](openapi/openapi.yaml) with the original being https://api.osrsbox.com/api-docs.
The real API responses don't quite align with the original specification so some things are modified:

- The `_id` property is removed as it conflicts with the `id` property.
- The `id` property is a string, not an integer; except for the `MonsterDrop` model.
- The `page` in the response meta is an integer, not a string.

Concurrent requests to the API server is limited to 10 so that the API server is not overloaded with queries that return multiple pages.

I do my best to make sure the client is up to date with API but if something isn't correct, please open an issue.
