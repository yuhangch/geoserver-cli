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
	"io"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// styleCmd represents the style command
var styleCmd = &cobra.Command{
	Use:   "style",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("style called")
	},
}

func pipe(cmds ...*exec.Cmd) {
	for i, cmd := range cmds {
		if i > 0 {
			out, _ := cmds[i-1].StdoutPipe()
			in, _ := cmd.StdinPipe()
			go func() {
				defer func() {
					out.Close()
					in.Close()
				}()
				io.Copy(in, out)
			}()
		}
	}
}

// styleEditCmd represents the style command
var styleEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		c1 := exec.Command("cat go.mod")
		c2 := exec.Command("vi -")
		c2.Stdin, _ = c1.StdoutPipe()
		c2.Stdout = os.Stdout
		_ = c2.Start()
		_ = c1.Run()
		_ = c2.Wait()

	},
}

func init() {
	rootCmd.AddCommand(styleCmd)
	styleCmd.AddCommand(styleEditCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// styleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// styleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
