/*
Copyright © 2022 Frederic Leist <frederic.leist@gmail.com>
*/
package cmd

import (
	"context"

	"github.com/XotoX1337/dogo/log"
	"github.com/XotoX1337/dogo/lookup"
	"github.com/docker/docker/api/types/container"
	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop one or many containers",
	Run: func(cmd *cobra.Command, args []string) {
		stop(args)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return lookup.Containers(false), cobra.ShellCompDirectiveNoFileComp
	},
}

func stop(args []string) {
	for _, argument := range args {
		containerList := lookup.Search(lookup.Containers(true), argument)
		if len(containerList) < 1 {
			log.Info("no container found for %s", argument)
		}
		stopContainers(containerList)
	}
}

func stopContainers(containers []string) {
	cli := lookup.Client()
	defer cli.Close()
	for _, c := range containers {
		info, err := cli.ContainerInspect(context.Background(), c)
		if err != nil {
			log.Warn("could not inspect container %s: %s", c, err)
			continue
		}
		if !info.State.Running {
			log.Info("container %s is not running", c)
			continue
		}
		log.Info("stopping %s...", c)
		err = cli.ContainerStop(context.Background(), c, container.StopOptions{})
		if err != nil {
			log.Warn("could not stop container %s", c)
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
