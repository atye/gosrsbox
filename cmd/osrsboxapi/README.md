```go install github.com/atye/gosrsbox/cmd/osrsboxapi```
```
$ which osrsboxapi
/usr/local/bin/osrsboxapi

$ osrsboxapi help
get Items, Monsters, and Prayers by wiki name, and custom queries
	supports MongoDB and Python queries as documented on osrsbox-api

Usage:
  osrsboxapi [command]

Available Commands:
  help         Help about any command
  itemids      Get items by ids
  itemnames    Get items by wiki name
  itemquery    Get items by MongoDB or Python queries
  monsternames Get monsters by wiki name
  monsterquery Get items by MongoDB or Python queries
  prayernames  Get prayers by name
  prayerquery  Get prayers by MongoDB or Python queries

Flags:
      --config string   config file (default is $HOME/.osrsboxapi.yaml)
  -h, --help            help for osrsboxapi
  -t, --toggle          Help message for toggle

Use "osrsboxapi [command] --help" for more information about a command.

$ osrsboxapi help itemnames
Get items by wiki name. Example:

	osrsboxapi itemnames "Abyssal whip" "Abyssal dagger"

Usage:
  osrsboxapi itemnames [flags]

Flags:
  -h, --help   help for itemnames

Global Flags:
      --config string   config file (default is $HOME/.osrsboxapi.yaml)
```