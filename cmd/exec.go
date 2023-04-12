/*
Copyright © 2022 Frederic Leist <frederic.leist@gmail.com>
*/
package cmd

import (
	"os"
	"strings"

	"github.com/XotoX1337/dogo/log"
	"github.com/XotoX1337/dogo/lookup"
	"github.com/XotoX1337/dogo/terminal"

	"github.com/spf13/cobra"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute a command in a running container",
	Run: func(cmd *cobra.Command, args []string) {
		executeCommand(args)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) != 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		return lookup.Containers(toComplete, false), cobra.ShellCompDirectiveNoFileComp
	},
}

func executeCommand(args []string) {
	if len(args) < 2 {
		log.Fatal("not enough arguments supplied, need at least 2")
	}

	err := terminal.ShellExecute("docker exec -it "+strings.Join(args, " "), terminal.ShellExecuteOpts{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	})
	if err != nil {
		log.Fatal("there was an error executing the command")
		log.Fatal(err.Error())
	}
	os.Exit(0)
}

func init() {
	rootCmd.AddCommand(execCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// execCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// execCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
