/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/XotoX1337/dogo/log"
	"github.com/XotoX1337/dogo/lookup"
	"github.com/spf13/cobra"
)

var Prune bool

// rebuildCmd represents the rebuild command
var rebuildCmd = &cobra.Command{
	Use:   "rebuild",
	Short: "rebuild one or many services",
	Run: func(cmd *cobra.Command, args []string) {
		removeDockerConfig()
		rebuild(args)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return lookup.Services(toComplete, true), cobra.ShellCompDirectiveNoFileComp
	},
}

func rebuild(args []string) {
	var wg sync.WaitGroup
	var serviceMap = map[string][]string{}

	for _, service := range args {
		config := lookup.ServiceConfig(service)
		_, exists := serviceMap[config]
		if !exists {
			serviceMap[config] = []string{}
		}
		parts := strings.Split(service, "-")
		parts = parts[1:]
		// remove chars before first "-" to get service name
		serviceMap[config] = append(serviceMap[config], strings.Join(parts, "-"))
	}

	for config, services := range serviceMap {
		wg.Add(1)
		file := config
		list := services
		go func() {
			defer wg.Done()
			rebuildServices(file, list)
		}()
	}

	wg.Wait()
}

func rebuildServices(config string, services []string) {

	log.Info(fmt.Sprintf("rebuilding %v...\n", services))
	command := exec.Command("bash", "-c", "docker compose -f "+config+" build --quiet "+strings.Join(services, " "))
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		log.Fatal(fmt.Sprintf("could not rebuild %v", services))
	}
	recreateServices(config, services)
	// done
}

func recreateServices(config string, services []string) {
	fmt.Printf("recreating %v...\n", services)
	command := exec.Command("bash", "-c", "docker compose -f "+config+" create "+strings.Join(services, " "))
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		fmt.Printf("%v", err)
	}
}

func removeDockerConfig() {
	homeDir, _ := os.UserHomeDir()
	os.Remove(homeDir + "/.docker/config.json")
}

func init() {
	rootCmd.AddCommand(rebuildCmd)
	rebuildCmd.Flags().BoolVarP(&Prune, "prune", "p", false, "prune build cache")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rebuildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rebuildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
