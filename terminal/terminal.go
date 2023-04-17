package terminal

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/XotoX1337/dogo/platform"
)

type ShellExecuteOpts struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

func Print(title string, content string) {
	fmt.Printf("%s:\n", title)
	fmt.Println(strings.Repeat("-", len(title)+1))
	for _, line := range strings.Split(content, "\n") {
		fmt.Printf("%s\n", line)
	}
	fmt.Println("")
}

func ShellExecute(command string, opts ShellExecuteOpts) error {

	p := platform.New()
	cmd := exec.Command(p.GetShell(), p.GetExec(), command)
	if opts.Stdout != nil {
		cmd.Stdout = opts.Stdout
	} else {
		cmd.Stdout = os.Stdout
	}

	if opts.Stdin != nil {
		cmd.Stdin = opts.Stdin
	}
	if opts.Stderr != nil {
		cmd.Stderr = opts.Stderr
	}

	return cmd.Run()
}
