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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("layer called")
	},
}

// layerCmd represents the layer command
var layerListCmd = &cobra.Command{
	Use:     "list",
	Aliases: listAlias,
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		api.LayersGet(&cfg, workspace)
	},
}

// layerDeleteCmd represents the layer delete command
var layerDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: deleteAlias,
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("require layer name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		api.LayerDelete(&cfg, workspace, args[0])
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
