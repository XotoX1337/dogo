package terminal

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/XotoX1337/dogo/platform"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
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

// PrintTable renders a titled, rounded-border table with the given header and
// rows and prints it to stdout. An empty row set prints a friendly hint instead.
func PrintTable(title string, header []string, rows [][]string) {
	if len(rows) == 0 {
		fmt.Printf("%s: none\n\n", title)
		return
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleRounded)
	t.SetTitle(title)
	t.Style().Title.Align = text.AlignCenter
	t.Style().Options.SeparateRows = false

	headerRow := make(table.Row, len(header))
	for i, h := range header {
		headerRow[i] = h
	}
	t.AppendHeader(headerRow)

	for _, r := range rows {
		row := make(table.Row, len(r))
		for i, c := range r {
			row[i] = c
		}
		t.AppendRow(row)
	}

	t.Render()
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
