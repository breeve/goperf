package server

import (
	"fmt"

	"github.com/breeve/goperf/pkg/utils"
	"github.com/spf13/cobra"
)

func Server() *cobra.Command {
	serverCmd := &cobra.Command{
		Use:     "server",
		Short:   "Server model",
		Long:    `Server model`,
		Run:     run,
		PreRunE: perRun,
	}
	serverCmd.Flags().StringVarP(&bindArg, "bind", "b", "127.0.0.1", "bind server")
	serverCmd.Flags().IntVarP(&portArg, "pord", "p", 0, "listen port")

	return serverCmd
}

var (
	bindArg string
	portArg int
)

func run(cmd *cobra.Command, args []string) {
	info()
}

func perRun(cmd *cobra.Command, args []string) error {
	// bindArg
	if !utils.IsValidIp(bindArg) {
		return fmt.Errorf("bind address:%s is invalid", bindArg)
	}

	// portArg
	if !utils.IsValidPort(portArg) {
		return fmt.Errorf("listen port:%d is invalid", portArg)
	}

	return nil
}

func info() {
	fmt.Printf("bind server: %s\n", bindArg)
	fmt.Printf("listen port: %d\n", portArg)
}
