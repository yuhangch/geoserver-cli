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
	"github.com/spf13/viper"
	"github.com/yuhangch/geoserver-cli/config"
)

var (
	serverURLFlag      string
	serverUsernameFlag string
	serverPasswordFlag string
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Server commander to control the local server config",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")
	},
}

var serverCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "To add a new server to local config",

	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a server's alias")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		serverCreate(args[0])
	},
}

var serverActivateCmd = &cobra.Command{
	Use:     "activate",
	Aliases: []string{"set"},
	Short:   "To add a new server to local config",

	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a server's alias")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		serverActivate(args[0])
	},
}

var serverDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: deleteAlias,
	Short:   "Delete a server in configuration",

	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a server's alias")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		serverDelete(args[0])
	},
}

var serverListCmd = &cobra.Command{
	Use:     "list",
	Aliases: listAlias,
	Short:   "List servers in configuration",

	Run: func(cmd *cobra.Command, args []string) {
		serverList()
	},
}

func serverList() {
	fmt.Println(cfg.Servers.ServersFmt())
}

func serverCreate(alias string) {
	i := cfg.Servers.IndexOf(alias)
	if i > -1 {
		fmt.Printf("%s already exists\n", alias)
		return
	}

	cfg.Servers = append(cfg.Servers, config.Server{
		Alias:    alias,
		URL:      serverURLFlag,
		Username: serverUsernameFlag,
		Password: serverPasswordFlag,
	})
	viper.Set("Servers", cfg.Servers)
	viper.WriteConfig()
	fmt.Printf("Server %s added to configuration", alias)
}

func serverDelete(alias string) {
	i := cfg.Servers.IndexOf(alias)
	if i < 0 {
		fmt.Println("No such server")
		return
	}
	cfg.Servers = append(cfg.Servers[:i], cfg.Servers[i+1:]...)
	viper.Set("Servers", cfg.Servers)
	viper.WriteConfig()
	fmt.Println("Delete success")

}
func serverActivate(alias string) {
	i := cfg.Servers.IndexOf(alias)
	if i < 0 {
		fmt.Println("No such server")
		return
	}
	cfg.Server = i
	viper.Set("Server", cfg.Server)
	viper.WriteConfig()
	fmt.Println("Now working on " + alias)

}
func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.AddCommand(serverCreateCmd, serverListCmd, serverDeleteCmd, serverActivateCmd)
	serverCreateCmd.Flags().StringVarP(&serverURLFlag, "url", "u", "", "Server url")
	serverCreateCmd.Flags().StringVarP(&serverUsernameFlag, "username", "n", "", "Server username")
	serverCreateCmd.Flags().StringVarP(&serverPasswordFlag, "password", "p", "", "Server password")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
