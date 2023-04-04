package terminal

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/XotoX1337/dogo/platform"
)

func Print(title string, content string) {
	fmt.Printf("%s:\n", title)
	fmt.Println(strings.Repeat("-", len(title)+1))
	for _, line := range strings.Split(content, "\n") {
		fmt.Printf("%s\n", line)
	}
	fmt.Println("")
}

func ShellExecute(command string) (string, error) {

	p := platform.New()
	cmd := exec.Command(p.GetShell(), p.GetExec(), command)
	output, error := cmd.CombinedOutput()
	return string(output), error
}
