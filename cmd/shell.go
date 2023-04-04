// Package cmd /*
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/XotoX1337/dogo/lookup"
	"github.com/spf13/cobra"
)

// shellCmd represents the ssh command
var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "use shell of a running container",
	Run: func(cmd *cobra.Command, args []string) {
		connect(args)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) != 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		return lookup.Containers(toComplete, false), cobra.ShellCompDirectiveNoFileComp
	},
}

func connect(args []string) {
	if len(args) == 0 {
		fmt.Println("no container given")
		os.Exit(1)
	}
	containerName := args[0]
	command := exec.Command("bash", "-c", "docker exec -it "+containerName+" /bin/bash")
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.AddCommand(shellCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// shellCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// shellCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
