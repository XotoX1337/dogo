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

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove one or many containers",
	Run: func(cmd *cobra.Command, args []string) {
		remove(args)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return functions.GetContainers(toComplete, true), cobra.ShellCompDirectiveNoFileComp
	},
}

func remove(args []string) {

	cli := functions.GetClient()
	for _, container := range args {
		cli.ContainerRemove(context.Background(), container, types.ContainerRemoveOptions{})
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
