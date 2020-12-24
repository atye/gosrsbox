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

// itemidsCmd represents the itemids command
var itemidsCmd = &cobra.Command{
	Use:   "itemids",
	Short: "Get items by ids",
	Long: `Get items by ids. Example:

	osrsboxapi itemids "1" "2"`,
	Run: func(cmd *cobra.Command, args []string) {
		api := osrsboxapi.NewAPI(&osrsboxapi.APIConfig{Logger: nil})
		items, err := api.GetItemsByID(context.Background(), args...)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		data, err := json.Marshal(items)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "%s\n", string(data))
	},
}

func init() {
	rootCmd.AddCommand(itemidsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// itemidsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// itemidsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
