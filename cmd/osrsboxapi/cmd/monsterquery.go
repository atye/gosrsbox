/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/atye/gosrsbox/osrsboxapi"
	"github.com/spf13/cobra"
)

// monsterqueryCmd represents the monsterquery command
var monsterqueryCmd = &cobra.Command{
	Use:   "monsterquery",
	Short: "Get items by MongoDB or Python queries",
	Long: `Get items by MongoDB or Python queries. Example:

	osrsboxapi monsterquery 'wiki_name=="Abyssal demon"'
	osrsboxapi monsterquery '{ "wiki_name": "Abyssal demon", "duplicate": false }'

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Fprintf(os.Stderr, "%s\n", "no query provided")
			os.Exit(1)
		}
		api := osrsboxapi.NewAPI(nil)
		monsters, err := api.GetMonstersByQuery(context.Background(), args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		data, err := json.MarshalIndent(monsters, "", "\t")
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "%s\n", string(data))
	},
}

func init() {
	rootCmd.AddCommand(monsterqueryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// monsterqueryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// monsterqueryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
