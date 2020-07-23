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
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yuhangch/geoserver-cli/api"
)

// layerCmd represents the layer command
var layerCmd = &cobra.Command{
	Use:   "layer",
	Short: "Manage layers",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("layer called")
	},
}

// layerCmd represents the layer command
var layerListCmd = &cobra.Command{
	Use:     "list",
	Aliases: listAlias,
	Short:   "Get layers list in a workspace",

	Run: func(cmd *cobra.Command, args []string) {
		var ws string
		if len(args) < 1 && len(workspace) < 1 {
			fmt.Println("require workspace name ")
		}
		if len(workspace) < 1 {
			ws = args[0]
		} else {
			ws = workspace
		}
		api.LayersGet(&cfg, ws)
	},
}

// layerDeleteCmd represents the layer delete command
var layerDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: deleteAlias,
	Short:   "Delete a layer",

	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("require layer name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		ws, name, err := api.ParseName(args[0], workspace)
		if err != nil {
			fmt.Println("datastore name error")
		}
		api.LayerDelete(&cfg, ws, name)
	},
}

func init() {
	rootCmd.AddCommand(layerCmd)
	layerCmd.AddCommand(layerListCmd, layerDeleteCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// layerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// layerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
