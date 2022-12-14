package functions

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var GlobalServices = map[string]string{}

const COMPOSE_SERVICE_LABEL = "com.docker.compose.service"
const COMPOSE_CONFIG_FILE_LABEL = "com.docker.compose.project.config_files"

func ContainerGet(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {

	options := types.ContainerListOptions{}
	//flag all
	containerList := getContainerList(options)
	containers := make([]string, len(containerList))
	for _, container := range containerList {
		containers = append(containers, container.Names[0][1:])
	}
	return containers, cobra.ShellCompDirectiveNoFileComp
}

func ServiceGet(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {

	options := types.ContainerListOptions{}
	//flag all
	containerList := getContainerList(options)
	services := make([]string, len(containerList))
	for _, container := range containerList {
		services = append(services, container.Image)
		GlobalServices[container.Image] = container.Labels[COMPOSE_CONFIG_FILE_LABEL]
	}
	return services, cobra.ShellCompDirectiveNoFileComp
}

func GetClient() *client.Client {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	return cli
}

func getContainerList(options types.ContainerListOptions) []types.Container {
	cli := GetClient()
	containerList, err := cli.ContainerList(context.Background(), options)
	if err != nil {
		panic(err)
	}
	return containerList
}
