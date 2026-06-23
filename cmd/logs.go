/*
Copyright © 2022 Frederic Leist <frederic.leist@gmail.com>
*/
package cmd

import (
	"context"
	"io"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/XotoX1337/dogo/log"
	"github.com/XotoX1337/dogo/lookup"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/spf13/cobra"
)

var logsCmdFollowFlag bool
var logsCmdTailFlag string
var logsCmdTimestampsFlag bool

// logsCmd represents the logs command
var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Show (and follow) the logs of one or many containers",
	Run: func(cmd *cobra.Command, args []string) {
		logs(args)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return lookup.Containers(toComplete, true), cobra.ShellCompDirectiveNoFileComp
	},
}

func logs(args []string) {
	var containers []string
	for _, argument := range args {
		found := lookup.Search(lookup.Containers("", true), argument)
		if len(found) < 1 {
			log.Info("no container found for %s", argument)
		}
		containers = append(containers, found...)
	}

	// cancel the streams cleanly on ctrl+c
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup
	for _, container := range containers {
		wg.Add(1)
		go func(container string) {
			defer wg.Done()
			streamLogs(ctx, container)
		}(container)
	}
	wg.Wait()
}

func streamLogs(ctx context.Context, container string) {
	cli := lookup.Client()

	info, err := cli.ContainerInspect(ctx, container)
	if err != nil {
		log.Warn("could not inspect container %s", container)
		return
	}

	reader, err := cli.ContainerLogs(ctx, container, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     logsCmdFollowFlag,
		Tail:       logsCmdTailFlag,
		Timestamps: logsCmdTimestampsFlag,
	})
	if err != nil {
		log.Warn("could not read logs of container %s", container)
		return
	}
	defer reader.Close()

	// containers without a TTY produce a multiplexed stream that has to be
	// demultiplexed; TTY containers emit a raw stream we can copy as-is
	if info.Config.Tty {
		_, err = io.Copy(os.Stdout, reader)
	} else {
		_, err = stdcopy.StdCopy(os.Stdout, os.Stderr, reader)
	}
	if err != nil && ctx.Err() == nil {
		log.Warn("error while streaming logs of container %s", container)
	}
}

func init() {
	rootCmd.AddCommand(logsCmd)

	logsCmd.Flags().BoolVarP(&logsCmdFollowFlag, "follow", "f", true, "follow log output (tail -f)")
	logsCmd.Flags().StringVarP(&logsCmdTailFlag, "tail", "n", "all", "number of lines to show from the end of the logs")
	logsCmd.Flags().BoolVarP(&logsCmdTimestampsFlag, "timestamps", "t", false, "show timestamps")
}
