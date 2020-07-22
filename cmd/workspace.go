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

// workspaceCmd represents the workspace command
var workspaceCmd = &cobra.Command{
	Use:   "workspace",
	Short: "Manage workspaces",

	Run: func(cmd *cobra.Command, args []string) {

	},
}

// workspaceCmd represents the workspace command
var workspaceListCmd = &cobra.Command{
	Use:     "list",
	Aliases: listAlias,
	Short:   "List workspaces",

	Run: func(cmd *cobra.Command, args []string) {
		api.WorkSpacesGet(&cfg)
	},
}

// workspaceCmd represents the workspace command
var workspaceCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new workspace",

	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a workspace name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println(args[0])
		api.WorkSpacesPost(&cfg, args[0])

	},
}

// workspaceCmd represents the workspace command
var workspaceDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: deleteAlias,
	Short:   "Delete a workspace",

	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a workspace name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Need workspace name")
			return
		}
		// fmt.Println(args[0])
		api.WorkSpaceDelete(&cfg, args[0])

	},
}

// workspaceCmd represents the workspace command
var workspacePutCmd = &cobra.Command{
	Use:     "rename",
	Aliases: renameAlias,
	Short:   "Rename a workspace",

	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("requires both workspace's old name and new name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		api.WorkSpacePut(&cfg, args[0], args[1])

	},
}

// workspaceList to get workspace list from server.

func init() {
	rootCmd.AddCommand(workspaceCmd)
	workspaceCmd.AddCommand(workspaceListCmd, workspaceCreateCmd, workspaceDeleteCmd, workspacePutCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// workspaceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// workspaceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
