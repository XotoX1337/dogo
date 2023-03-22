package functions

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var GlobalServices = map[string]string{}

const COMPOSE_SERVICE_LABEL = "com.docker.compose.service"
const COMPOSE_CONFIG_FILE_LABEL = "com.docker.compose.project.config_files"

func GetContainers(toComplete string, all bool) []string {
	options := types.ContainerListOptions{All: all}
	containerList := GetContainerList(options)
	containers := make([]string, len(containerList))
	for _, container := range containerList {
		if len(container.Names[0][1:]) <= 0 {
			continue
		}
		containers = append(containers, container.Names[0][1:])
	}
	return containers
}
func GetServices(toComplete string, all bool) []string {
	options := types.ContainerListOptions{All: all}
	//flag all
	containerList := GetContainerList(options)
	services := make([]string, len(containerList))
	for _, container := range containerList {
		if len(container.Image) <= 0 {
			continue
		}
		services = append(services, container.Image)
		GlobalServices[container.Image] = container.Labels[COMPOSE_CONFIG_FILE_LABEL]
	}
	return services
}

func FetchServiceConfig(serviceName string) string {
	_, exists := GlobalServices[serviceName]
	if !exists {
		generateServiceMap()
	}
	return GlobalServices[serviceName]
}

func GetClient() *client.Client {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	return cli
}

func GetContainerList(options types.ContainerListOptions) []types.Container {
	cli := GetClient()
	containerList, err := cli.ContainerList(context.Background(), options)
	if err != nil {
		panic(err)
	}
	return containerList
}

func generateServiceMap() {
	GetServices("", true)
}
