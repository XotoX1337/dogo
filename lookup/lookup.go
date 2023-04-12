package lookup

import (
	"context"
	"os"
	"strings"

	"github.com/XotoX1337/dogo/constants"
	"github.com/XotoX1337/dogo/log"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var GlobalServices = map[string]string{}

func Containers(toComplete string, all bool) []string {
	options := types.ContainerListOptions{All: all}
	containerList := ContainerList(options)
	var containers []string
	for _, container := range containerList {
		containers = append(containers, container.Names[0][1:])
	}
	return containers
}

func Search(slice []string, query string) []string {

	if query == "*" {
		log.Warn("need at least one character for wildcard search")
		os.Exit(1)
	}
	var found []string
	for _, element := range slice {
		if strings.HasSuffix(query, "*") {
			if strings.HasPrefix(element, strings.TrimSuffix(query, "*")) {
				found = append(found, element)
			}
		}
		if element == query {
			found = append(found, element)
		}
	}
	return found
}
func Services(toComplete string, all bool) []string {
	options := types.ContainerListOptions{All: all}
	//flag all
	containerList := ContainerList(options)
	var services []string
	for _, container := range containerList {
		services = append(services, container.Image)
		GlobalServices[container.Image] = container.Labels[constants.COMPOSE_CONFIG_FILE_LABEL]

	}
	return services
}

func ServiceConfig(serviceName string) string {
	_, exists := GlobalServices[serviceName]
	if !exists {
		generateServiceMap()
	}
	return GlobalServices[serviceName]
}

func Client() *client.Client {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatal("could not establish a connection with docker. Is docker installed?")
	}
	return cli
}

func ContainerList(options types.ContainerListOptions) []types.Container {
	cli := Client()
	containerList, err := cli.ContainerList(context.Background(), options)
	if err != nil {
		log.Fatal("could not establish a connection with docker. Is docker installed?")
	}
	return containerList
}

func generateServiceMap() {
	Services("", true)
}
