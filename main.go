/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"dogo/cmd"
	"fmt"
	"os/exec"
)

func main() {
	cmd.Execute()
	result := exec.Command("locate", "docker-compose.yml")
	out, err := result.CombinedOutput()
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("%s\n", out)
}
