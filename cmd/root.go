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
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yuhangch/geoserver-cli/config"
)

var (
	cfgFile   string
	workspace string
	recurse   bool
	cfg       config.Config

	listAlias   = []string{"ls"}
	renameAlias = []string{"mv"}
	deleteAlias = []string{"rm", "remove", "del"}
	createAlias = []string{"add", "new"}
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gctl",
	Short: "GeoServer CLI control panel",
	Long:  `A cli tool for GeoServer access .`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.geoserver-cli.yaml)")
	rootCmd.PersistentFlags().StringVar(&workspace, "workspace", "", "select a workspace")
	rootCmd.PersistentFlags().BoolVarP(&recurse, "recurse", "r", false, "recurse delete")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		//Find home directory.
		// home, err := os.UserHomeDir()
		// if err != nil {
		// 	fmt.Println(err)
		// 	os.Exit(1)
		// }
		// fmt.Println(home)
		// fmt.Println(os.TempDir())

		// Search config in home directory with name ".geoserver-cli" (without extension).
		viper.AddConfigPath("./")
		viper.SetConfigName(".geoserver-cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		// fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	err := viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalln(err)
	}
}
