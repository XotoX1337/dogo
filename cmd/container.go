/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/XotoX1337/dogo/lookup"
	"github.com/XotoX1337/dogo/terminal"
	"github.com/docker/docker/api/types"
	"github.com/spf13/cobra"
)

// containerCmd represents the container command
var containerCmd = &cobra.Command{
	Use:   "container",
	Short: "Prints all available docker containers",
	Run: func(cmd *cobra.Command, args []string) {
		var rows [][]string
		for _, container := range lookup.ContainerList(types.ContainerListOptions{All: true}) {
			rows = append(rows, []string{
				container.Names[0][1:],
				container.Image,
				container.State,
				container.Status,
			})
		}
		terminal.PrintTable("Container", []string{"Name", "Image", "State", "Status"}, rows)
	},
}

func init() {
	listCmd.AddCommand(containerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// containerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// containerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
