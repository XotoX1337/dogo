/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/XotoX1337/dogo/constants"
	"github.com/XotoX1337/dogo/lookup"
	"github.com/XotoX1337/dogo/terminal"
	"github.com/docker/docker/api/types"
	"github.com/spf13/cobra"
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "List all containers and services",
	Run: func(cmd *cobra.Command, args []string) {
		var rows [][]string
		for _, container := range lookup.ContainerList(types.ContainerListOptions{All: true}) {
			service := container.Labels[constants.COMPOSE_SERVICE_LABEL]
			if service == "" {
				service = "-"
			}
			rows = append(rows, []string{
				service,
				container.Names[0][1:],
				container.Image,
				container.State,
				container.Status,
			})
		}
		terminal.PrintTable("Containers & Services", []string{"Service", "Container", "Image", "State", "Status"}, rows)
	},
}

func init() {
	listCmd.AddCommand(allCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// allCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// allCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
