package cmd

/*
Copyright © 2020 Yuhang Chen <i@yuhang.ch>

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
import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yuhangch/geoserver-cli/api"
)

// styleCmd represents the style command
var styleCmd = &cobra.Command{
	Use:   "style",
	Short: "Manage styles",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("style called")
	},
}

// styleEditCmd represents the style command
var styleEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a style in vim",

	Run: func(cmd *cobra.Command, args []string) {
		ws, name, _ := api.ParseName(args[0], workspace)
		api.StyleEdit(&cfg, ws, name)
	},
}

// styleEditCmd represents the style command
var styleInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "HandleBody style's information",

	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("require style name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		ws, name, _ := api.ParseName(args[0], workspace)
		s, _ := api.StyleGet(&cfg, ws, name)
		p, _ := json.MarshalIndent(s, "", "  ")

		fmt.Printf("%s \n", p)
	},
}

func init() {
	rootCmd.AddCommand(styleCmd)
	styleCmd.AddCommand(styleEditCmd, styleInfoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// styleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// styleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
