package functions

import (
	"context"
	"github.com/spf13/cobra"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func ContainerGet(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if len(args) != 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	// fetch containerList from docker API
	cli := GetClient()
	containerList, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	containers := make([]string, len(containerList))
	for _, container := range containerList {
		containers = append(containers, container.Names[0][1:])
	}
	return containers, cobra.ShellCompDirectiveNoFileComp
}

//func ServiceGet(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
//	if len(args) != 0 {
//		return nil, cobra.ShellCompDirectiveNoFileComp
//	}
//	cli := GetClient()
//	serviceList, err := cli.ServiceList(context.Background(), types.ServiceListOptions{})
//	if err != nil {
//		panic(err)
//	}
//
//	services := make([]string, len(serviceList))
//	for _, service := range serviceList {
//		services = append(services, service.Names[0][1:])
//	}
//	return containerList, cobra.ShellCompDirectiveNoFileComp
//}

func GetClient() *client.Client {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	return cli
}
