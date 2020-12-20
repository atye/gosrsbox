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

// prayernamesCmd represents the prayernames command
var prayernamesCmd = &cobra.Command{
	Use:   "prayernames",
	Short: "Get prayers by name",
	Long: `Get prayers by name. Example:

	osrsboxdb prayernames "Thick Skin" "Burst of Strength"`,
	Run: func(cmd *cobra.Command, args []string) {
		api := restful.NewAPI(nil)
		prayers, err := api.GetPrayersByName(context.Background(), args...)
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
	rootCmd.AddCommand(prayernamesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// prayernamesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// prayernamesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
