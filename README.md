[![Build Status](https://travis-ci.org/atye/gosrsbox.svg?branch=master)](https://travis-ci.org/atye/gosrsbox) [![Coverage Status](https://coveralls.io/repos/github/atye/gosrsbox/badge.svg?branch=add-travis-ci)](https://coveralls.io/github/atye/gosrsbox?branch=add-travis-ci)

gosrsbox is a client for the osrsbox-api; https://api.osrsbox.com.

The client supports getting items, monsters, and prayers. The API has /equipment and /weapons endpoints but those datasets are subsets of items.
So as far as this client is concerned, those entities are items. See the godoc example for this in action.

https://godoc.org/github.com/atye/gosrsbox

### Features
- Get all items, monsters, and prayers
- Get items, monsters, and prayers by name
- Get items, monsters, and prayers by custom MongoDB query
- Get monsters that drop specific items

### Development
I will continue to think of common use cases, such as getting entities by name, that this client could wrap.

If you want something added or have an idea, feel free to open an issue.