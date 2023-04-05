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

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start one or many containers",
	Run: func(cmd *cobra.Command, args []string) {
		start(args)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return lookup.Containers(toComplete, true), cobra.ShellCompDirectiveNoFileComp
	},
}

func start(args []string) {
	for _, argument := range args {
		containerList := lookup.Search(lookup.Containers("", true), argument)
		startContainers(containerList)
	}
}

func startContainers(containers []string) {
	cli := lookup.Client()
	for _, container := range containers {
		log.Info("starting %s...", container)
		err := cli.ContainerStart(context.Background(), container, types.ContainerStartOptions{})
		if err != nil {
			log.Warn("could not start container %s", container)
		}
	}
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
