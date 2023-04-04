/*
Copyright © 2022 Frederic Leist <frederic.leist@gmail.com>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/XotoX1337/dogo/log"
	"github.com/XotoX1337/dogo/lookup"
	containertypes "github.com/docker/docker/api/types/container"
	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop one or many containers",
	Run: func(cmd *cobra.Command, args []string) {
		stop(args)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return lookup.Containers(toComplete, false), cobra.ShellCompDirectiveNoFileComp
	},
}

func stop(args []string) {
	for _, argument := range args {
		containerList := lookup.Search(lookup.Containers("", true), argument)
		stopContainers(containerList)
	}
}

func stopContainers(containers []string) {
	cli := lookup.Client()
	for _, container := range containers {
		log.Info(fmt.Sprintf("stopping %s...", container))
		err := cli.ContainerStop(context.Background(), container, containertypes.StopOptions{})
		if err != nil {
			log.Warn(fmt.Sprintf("could not stop container %s", container))
		}
	}
}

func init() {
	rootCmd.AddCommand(stopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
