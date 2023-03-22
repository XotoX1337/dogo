/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"dogo/functions"

	"github.com/docker/docker/api/types"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start one or many containers",
	Run: func(cmd *cobra.Command, args []string) {
		start(args)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return functions.GetContainers(toComplete, true), cobra.ShellCompDirectiveNoFileComp
	},
}

func start(args []string) {
	cli := functions.GetClient()
	for _, container := range args {
		cli.ContainerStart(context.Background(), container, types.ContainerStartOptions{})
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
