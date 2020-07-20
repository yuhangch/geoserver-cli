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

var (
	datastoreFile string
)

// datastoreCmd represents the datastore command
var datastoreCmd = &cobra.Command{
	Use:   "datastore",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("datastore called")
	},
}

// datastoreCmd represents the datastore command
var datastoreListCmd = &cobra.Command{
	Use:     "list",
	Aliases: listAlias,
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(workspace) < 1 {
			return errors.New("requires a workspace name")
		}
		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		api.DataStoresGet(&cfg, workspace)
	},
}

// datastoreCreateCmd represents the datastore command
var datastoreCreateCmd = &cobra.Command{
	Use:     "create",
	Aliases: createAlias,
	Short:   "create a new datastore in a select workspace",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(workspace) < 1 {
			return errors.New("requires a workspace name")
		}
		if len(args) < 1 {
			return errors.New("require datastore name")
		}
		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		api.DataStoresCreate(&cfg, workspace, args[0], datastoreFile)
	},
}

// datastoreCreateCmd represents the datastore command
var datastoreDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: deleteAlias,
	Short:   "create a new datastore in a select workspace",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(workspace) < 1 {
			return errors.New("requires a workspace name")
		}
		if len(args) < 1 {
			return errors.New("require datastore name")
		}
		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		api.DataStoreDelete(&cfg, workspace, args[0], recurse)
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
}
