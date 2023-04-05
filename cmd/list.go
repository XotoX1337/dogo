/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:       "list",
	Short:     "List all containers & services",
	Args:      cobra.MatchAll(cobra.OnlyValidArgs),
	ValidArgs: []string{"all", "container", "services"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			allCmd.Run(cmd, args)
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
