/*
Copyright © 2022 Frederic Leist <frederic.leist@gmail.com>
*/
package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"

	"github.com/XotoX1337/dogo/constants"
	"github.com/XotoX1337/dogo/log"
	"github.com/spf13/cobra"
)

var completionCmdDestinationFlag string
var completionCmdFileFlag bool

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
	Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
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

func getProfilePath(terminal string) string {
	var profilePath string
	homeDir, _ := os.UserHomeDir()

	switch terminal {
	case "bash", "zsh", "fish":
		profilePath = filepath.Join(homeDir)
	case "powershell":
		profilePath = filepath.Join(homeDir, "Documents", "WindowsPowerShell")
	}

	return profilePath
}

func getScriptPath(terminal string, cmd *cobra.Command) string {

	var filename string = "dogo-completion"
	var defaultPath string
	var scriptPath string
	profilePath := getProfilePath(terminal)
	customDest, _ := cmd.Flags().GetString("destination")

	switch terminal {
	case "bash", "zsh", "fish":
		defaultPath = filepath.Join(profilePath, ".bash_completion.d")
		filename += ".sh"
	case "powershell":
		defaultPath = filepath.Join(profilePath)
		filename += ".ps1"
	}
	scriptPath = defaultPath
	if customDest != "" {
		scriptPath = filepath.Join(customDest)
	}
	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		os.MkdirAll(scriptPath, 0644)
	}
	return filepath.Join(scriptPath, filename)
}

func getWriter(terminal string, cmd *cobra.Command) io.Writer {
	writeToFile, _ := cmd.Flags().GetBool("file")

	if !writeToFile {
		return os.Stdout
	}
	scriptPath := getScriptPath(terminal, cmd)
	writer, writeError := os.OpenFile(scriptPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if writeError != nil {
		log.Warn("could not write completion script")
		log.Fatal(writeError.Error())
	}

	appendError := appendToProfile(terminal, scriptPath)
	if appendError != nil {
		log.Warn("could not append completion script")
		log.Fatal(appendError.Error())
	}

	return writer
}

func appendToProfile(terminal string, scriptPath string) error {
	profilePath := getProfilePath(terminal)
	var err error
	var profile string
	if terminal == "powershell" {
		profile = filepath.Join(profilePath, "Microsoft.PowerShell_profile.ps1")
	} else {
		profile = filepath.Join(profilePath, ".bashrc")
	}
	profileContent, err := os.ReadFile(profile)
	r := regexp.MustCompile(constants.PROFILE_PREFIX)
	if !r.Match(profileContent) {
		writer, _ := os.OpenFile(profile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		_, err = writer.WriteString(fmt.Sprintf("\n\n%s\n. %s", constants.PROFILE_PREFIX, scriptPath))
	}

	return err
}

func init() {
	rootCmd.AddCommand(completionCmd)
	completionCmd.Flags().BoolVarP(&completionCmdFileFlag, "file", "f", false, "write completion to file instead of stdout")
	completionCmd.Flags().StringVarP(&completionCmdDestinationFlag, "dest", "d", "", "specify file destination, defaults to $HOME/.bash-completion.d/dogo-completion.sh")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
