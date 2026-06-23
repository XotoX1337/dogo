package lookup

import (
	"context"
	"os"
	"strings"

	"github.com/XotoX1337/dogo/constants"
	"github.com/XotoX1337/dogo/log"
	"github.com/docker/docker/api/types/container"
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

func Containers(all bool) []string {
	containerList := ContainerList(container.ListOptions{All: all})
	var containers []string
	for _, c := range containerList {
		containers = append(containers, c.Names[0][1:])
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
func Services() []string {
	containerList := ContainerList(container.ListOptions{All: true})
	var services []string
	for _, c := range containerList {
		services = append(services, c.Image)
	}
	return services
}

func Client() *client.Client {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatal("could not establish a connection with docker. Is docker installed?")
	}
	return cli
}

func ContainerList(options container.ListOptions) []container.Summary {
	cli := Client()
	containerList, err := cli.ContainerList(context.Background(), options)
	if err != nil {
		log.Fatal("could not establish a connection with docker. Is docker installed?")
	}
	return containerList
}

func GenerateServiceMap() {
	containerList := ContainerList(container.ListOptions{All: true})
	for _, c := range containerList {
		index := c.Labels[constants.COMPOSE_CONFIG_FILE_LABEL]
		ServicesMap[index] = ConfigDetails{
			Ready:             true,
			Services:          append(ServicesMap[index].Services, c.Labels[constants.COMPOSE_SERVICE_LABEL]),
			Images:            append(ServicesMap[index].Images, c.Image),
			Containers:        append(ServicesMap[index].Containers, c.Names[0][1:]),
			ImageContainerMap: append(ServicesMap[index].ImageContainerMap, map[string]string{c.Image: c.Names[0][1:]}),
			Config:            index,
		}

	}
}
