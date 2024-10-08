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
	serverCmd.Flags().IntVarP(&portArg, "port", "p", 0, "listen port")
	serverCmd.Flags().StringVarP(&protocolArg, "protocol", "P", "tcp", "listen protocol, (tcp, udp)")

	return serverCmd
}

var (
	bindArg     string
	portArg     int
	protocolArg string
)

func run(cmd *cobra.Command, args []string) {
	info()
	listen(bindArg, portArg, protocolArg)
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

	// protocolArg
	if !utils.IsValidProtocol(protocolArg) {
		return fmt.Errorf("protocol:%s is invalid", protocolArg)
	}

	return nil
}

func info() {
	fmt.Printf("bind server: %s\n", bindArg)
	fmt.Printf("listen port: %d\n", portArg)
	fmt.Printf("protocol:%s\n", protocolArg)
}
