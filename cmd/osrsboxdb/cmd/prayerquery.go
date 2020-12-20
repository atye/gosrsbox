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

	"github.com/atye/gosrsbox/osrsboxdb/api/restful"
	"github.com/spf13/cobra"
)

// prayerqueryCmd represents the prayerquery command
var prayerqueryCmd = &cobra.Command{
	Use:   "prayerquery",
	Short: "Get prayers by MongoDB or Python queries",
	Long: `Get items by MongoDB or Python queries. Example:

	osrsboxdb prayerquery 'name=="Thick Skin"'
	osrsboxdb prayerquery '{ "name": "Thick Skin"}'

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Fprintf(os.Stderr, "%s\n", "no query provided")
			os.Exit(1)
		}
		api := restful.NewAPI(nil)
		prayers, err := api.GetPrayersByQuery(context.Background(), args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		data, err := json.MarshalIndent(prayers, "", "\t")
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "%s\n", string(data))
	},
}

func init() {
	rootCmd.AddCommand(prayerqueryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// prayerqueryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// prayerqueryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
