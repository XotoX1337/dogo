/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/XotoX1337/dogo/lookup"
	"github.com/XotoX1337/dogo/terminal"
	"github.com/jedib0t/go-pretty/v6/list"
	"github.com/spf13/cobra"
)

// servicesCmd represents the services command
var servicesCmd = &cobra.Command{
	Use:   "services",
	Short: "Prints all available docker services",
	Run: func(cmd *cobra.Command, args []string) {
		l := list.NewWriter()
		l.SetStyle(list.StyleBulletCircle)
		for _, service := range lookup.Services("", true) {
			l.AppendItem(service)
		}
		terminal.Print("Services", l.Render())
	},
}

func init() {
	listCmd.AddCommand(servicesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// servicesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// servicesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
