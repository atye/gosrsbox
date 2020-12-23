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

	"github.com/atye/gosrsbox/osrsboxapi/api"
	"github.com/spf13/cobra"
)

// itemnamesCmd represents the itemnames command
var itemnamesCmd = &cobra.Command{
	Use:   "itemnames",
	Short: "Get items by wiki name",
	Long: `Get items by wiki name. Example:

	osrsboxapi itemnames "Abyssal whip" "Abyssal dagger"`,
	Run: func(cmd *cobra.Command, args []string) {
		api := api.NewAPI(nil)
		items, err := api.GetItemsByName(context.Background(), args...)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		data, err := json.MarshalIndent(items, "", "\t")
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "%s\n", string(data))
	},
}

func init() {
	rootCmd.AddCommand(itemnamesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// itemnamesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// itemnamesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
