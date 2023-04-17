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

type Service struct {
	Service   string
	Container string
	Image     string
}
type ConfigDetails struct {
	Ready bool

	Services          []string
	Containers        []string
	Images            []string
	ImageContainerMap []map[string]string
	Config            string
}

var ServicesMap = map[string]ConfigDetails{}

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
	}
	return services
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

func GenerateServiceMap() {
	containerList := ContainerList(types.ContainerListOptions{All: true})
	for _, container := range containerList {
		index := container.Labels[constants.COMPOSE_CONFIG_FILE_LABEL]
		ServicesMap[index] = ConfigDetails{
			Ready:             true,
			Services:          append(ServicesMap[index].Services, container.Labels[constants.COMPOSE_SERVICE_LABEL]),
			Images:            append(ServicesMap[index].Images, container.Image),
			Containers:        append(ServicesMap[index].Containers, container.Names[0][1:]),
			ImageContainerMap: append(ServicesMap[index].ImageContainerMap, map[string]string{container.Image: container.Names[0][1:]}),
			Config:            index,
		}

	}
}
