/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"io"
	"os"

	"github.com/spf13/cobra"
)

var Destination string
var File bool

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate completion script",
	PreRun: func(cmd *cobra.Command, args []string) {
		file, _ := cmd.Flags().GetString("dest")
		if file != "" {
			cmd.MarkFlagRequired("file")
		}
	},
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		writer := getWriter(args[0], cmd)
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(writer)
		case "zsh":
			cmd.Root().GenZshCompletion(writer)
		case "fish":
			cmd.Root().GenFishCompletion(writer, true)
		case "powershell":
			cmd.Root().GenPowerShellCompletionWithDesc(writer)
		}
	},
}

func getWriter(terminal string, cmd *cobra.Command) io.Writer {
	writeToFile, _ := cmd.Flags().GetBool("file")
	fileDestination, _ := cmd.Flags().GetString("destination")
	if !writeToFile {
		return os.Stdout
	}
	var writer io.Writer
	switch terminal {
	case "bash":
		homeDir, _ := os.UserHomeDir()
		destination := homeDir + "/.bash_completion.d/dogo-completion.sh"
		if fileDestination != "" {
			destination = fileDestination
		}
		file, err := os.OpenFile(destination, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		writer = file
	case "zsh":
		writer = os.Stdout
	case "fish":
		writer = os.Stdout
	case "powershell":
		writer = os.Stdout
	}

	return writer
}

func init() {
	rootCmd.AddCommand(completionCmd)
	completionCmd.Flags().BoolVarP(&File, "file", "f", false, "write completion to file instead of stdout")
	completionCmd.Flags().StringVarP(&Destination, "dest", "d", "", "specify file destination, defaults to $HOME/.bash-completion.d/dogo-completion.sh")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
