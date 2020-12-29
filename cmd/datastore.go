package cmd

/*
Copyright © 2020 Yuhang Chen <i@yuhang.ch>

Licensed under the Apache License, Version 2.0 (the "License");
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

var (
	datastoreFile      string
	datastoreMethod    string
	datastoreFormat    string
	datastoreConfigure string
)

// datastoreCmd represents the datastore command
var datastoreCmd = &cobra.Command{
	Use:   "datastore",
	Short: "Manage data stores.",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("datastore called")
	},
}

// datastoreCmd represents the datastore command
var datastoreListCmd = &cobra.Command{
	Use:     "list",
	Aliases: listAlias,
	Short:   "HandleBody datastores list in a workspace",

	Args: func(cmd *cobra.Command, args []string) error {
		if len(workspace) < 1 && len(args) < 1 {
			return errors.New("requires a workspace name")
		}
		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		var ws string
		if len(workspace) > 0 {
			ws = workspace
		} else {
			ws = args[0]
		}

		api.DataStoresGet(&cfg, ws)
	},
}

// datastoreCreateCmd represents the datastore command
var datastoreCreateCmd = &cobra.Command{
	Use:     "create",
	Aliases: createAlias,
	Short:   "Create a new datastore in a select workspace",

	Args: func(cmd *cobra.Command, args []string) error {

		if len(args) < 1 {
			return errors.New("require datastore name")
		}
		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		ws, name, err := api.ParseName(args[0], workspace)
		if err != nil {
			fmt.Println("datastore name error")
		}
		api.DataStoresCreate(&cfg, ws, name, datastoreFile, datastoreMethod, datastoreFormat, datastoreConfigure)

	},
}

// datastoreCreateCmd represents the datastore command
var datastoreDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: deleteAlias,
	Short:   "Create a new datastore in a select workspace",

	Args: func(cmd *cobra.Command, args []string) error {

		if len(args) < 1 {
			return errors.New("require datastore name")
		}
		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		ws, name, err := api.ParseName(args[0], workspace)
		if err != nil {
			fmt.Println("datastore name error")
		}
		api.DataStoreDelete(&cfg, ws, name, recurse)
	},
}

func init() {
	rootCmd.AddCommand(datastoreCmd)
	datastoreCmd.AddCommand(datastoreListCmd, datastoreCreateCmd, datastoreDeleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// datastoreCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	datastoreCreateCmd.Flags().StringVarP(&datastoreFile, "file", "f", "", "file path")
	datastoreCreateCmd.Flags().StringVarP(&datastoreConfigure, "configure", "", "none", "set configure")
	datastoreCreateCmd.Flags().StringVarP(&datastoreMethod, "method", "", "file", "set datastore method (file,url,external)")
	datastoreCreateCmd.Flags().StringVarP(&datastoreMethod, "format", "", "shp", "set datastore format (e.g. shp)")
}
