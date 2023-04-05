/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"path/filepath"
	"strings"

	"github.com/XotoX1337/dogo/log"
	"github.com/XotoX1337/dogo/terminal"
	"github.com/spf13/cobra"
)

var createCmdFileFlag string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:                   "create",
	Short:                 "Create all or a specific service from a docker-compose.yml file",
	Args:                  cobra.MatchAll(cobra.MaximumNArgs(1), cobra.OnlyValidArgs),
	DisableFlagsInUseLine: true,
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.MarkFlagRequired("file")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {

		var service string
		if len(args) == 1 {
			service = args[0]
		}
		file, _ := cmd.Flags().GetString("file")
		create(service, file)
	},
}

func create(service string, file string) {
	cmdSlice := []string{
		"docker compose",
	}
	if file != "" {
		path := filepath.Join(file, "docker-compose.yml")
		cmdSlice = append(cmdSlice, "-f", path)
	}
	cmdSlice = append(cmdSlice, "create")
	if service != "" {
		cmdSlice = append(cmdSlice, service)
	}
	out, err := terminal.ShellExecute(strings.Join(cmdSlice, " "))
	if err != nil {
		log.Warn(out)
	}
	log.Info(out)
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&createCmdFileFlag, "file", "f", "", "write completion to file instead of stdout")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
