/*
Copyright © 2022 Frederic Leist <frederic.leist@gmail.com>
*/
package cmd

import (
	"context"

	"github.com/XotoX1337/dogo/log"
	"github.com/XotoX1337/dogo/lookup"
	"github.com/docker/docker/api/types"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove one or many containers",
	Run: func(cmd *cobra.Command, args []string) {
		remove(args)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return lookup.Containers(toComplete, true), cobra.ShellCompDirectiveNoFileComp
	},
}

func remove(args []string) {
	for _, argument := range args {
		containerList := lookup.Search(lookup.Containers("", true), argument)
		if len(containerList) < 1 {
			log.Info("no container found for %s", argument)
		}
		removeContainers(containerList)
	}
}

func removeContainers(containers []string) {
	cli := lookup.Client()
	for _, container := range containers {
		log.Info("removing %s...", container)
		err := cli.ContainerRemove(context.Background(), container, types.ContainerRemoveOptions{})
		if err != nil {
			log.Warn("could not remove container %s", container)
		}
	}
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
