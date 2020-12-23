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

// monsternamesCmd represents the monsternames command
var monsternamesCmd = &cobra.Command{
	Use:   "monsternames",
	Short: "Get monsters by wiki name",
	Long: `Get monsters by wiki name. Example:

	osrsboxapi monsternames "Molanisk" "Aberrant spectre"`,
	Run: func(cmd *cobra.Command, args []string) {
		api := osrsboxapi.NewAPI(nil)
		monsters, err := api.GetMonstersByName(context.Background(), args...)
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
	rootCmd.AddCommand(monsternamesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// monsternamesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// monsternamesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
