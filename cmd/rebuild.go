/*
Copyright © 2022 Frederic Leist <frederic.leist@gmail.com>
*/
package cmd

import (
	"os"
	"regexp"
	"strings"
	"sync"

	"github.com/XotoX1337/dogo/constants"
	"github.com/XotoX1337/dogo/log"
	"github.com/XotoX1337/dogo/lookup"
	"github.com/XotoX1337/dogo/terminal"
	"github.com/docker/docker/api/types"
	"github.com/spf13/cobra"
)

type configDetails struct {
	services   []string
	containers []string
	images     []string
}

// has all services and containers
// indexed by the corresponding
// docker-compose.yml
var configs = map[string]configDetails{}
var services []string

var rebuildCmdPruneFlag bool

// rebuildCmd represents the rebuild command
var rebuildCmd = &cobra.Command{
	Use:   "rebuild",
	Short: "Rebuild one or many services",
	Run: func(cmd *cobra.Command, args []string) {
		loadServices(args)
		loadConfigs()
		removeDockerConfig()

		var wg sync.WaitGroup
		for config, details := range configs {
			wg.Add(1)
			go func(config string, details configDetails) {
				defer wg.Done()
				rebuild(config, details.images, details.services)
				recreate(config, details.images, details.services)
				startCmd.Run(cmd, details.containers)
			}(config, details)
		}
		wg.Wait()
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return lookup.Services(toComplete, true), cobra.ShellCompDirectiveNoFileComp
	},
}

func rebuild(config string, images []string, services []string) {
	log.Info("rebuilding %v...", images)
	err := terminal.ShellExecute("docker compose -f "+config+" build "+strings.Join(services, " "), terminal.ShellExecuteOpts{})
	if err != nil {
		log.Fatal("could not recreate %v", images)
	}
}

func recreate(config string, images []string, services []string) {
	log.Info("recreating %v...", images)
	err := terminal.ShellExecute("docker compose -f "+config+" create "+strings.Join(services, " "), terminal.ShellExecuteOpts{})
	if err != nil {
		log.Fatal("could not recreate %v", images)
	}
}

func removeDockerConfig() {
	homeDir, _ := os.UserHomeDir()
	os.Remove(homeDir + "/.docker/config.json")
}

func loadConfigs() {
	containerList := lookup.ContainerList(types.ContainerListOptions{All: true})
	for _, service := range services {
		for _, container := range containerList {
			if container.Image != service {
				continue
			}

			index := convertWslPath(container.Labels[constants.COMPOSE_CONFIG_FILE_LABEL])
			serviceName := container.Labels[constants.COMPOSE_SERVICE_LABEL]
			containerName := container.Names[0][1:]
			imageName := container.Image
			configs[index] = configDetails{
				services:   append(configs[index].services, serviceName),
				containers: append(configs[index].containers, containerName),
				images:     append(configs[index].images, imageName),
			}
		}
	}
}

func convertWslPath(path string) string {
	re := regexp.MustCompile(`(?m)\\+wsl\$\\+\w+`)
	match := re.FindString(path)

	if match != "" {
		path = strings.Replace(path, match, "", -1)
		path = strings.ReplaceAll(path, "\\", "/")
	}
	return path
}

func loadServices(args []string) {
	serviceList := lookup.Services("", true)
	for _, arg := range args {
		services = append(services, lookup.Search(serviceList, arg)...)
	}
}

func init() {
	rootCmd.AddCommand(rebuildCmd)
	rebuildCmd.Flags().BoolVarP(&rebuildCmdPruneFlag, "prune", "p", false, "prune build cache")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rebuildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rebuildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
