package client

import (
	"fmt"

	"github.com/breeve/goperf/pkg/utils"
	"github.com/spf13/cobra"
)

func Client() *cobra.Command {
	clientCmd := &cobra.Command{
		Use:     "client",
		Short:   "Client model",
		Long:    `Client model`,
		Run:     run,
		PreRunE: perRun,
	}

	clientCmd.Flags().StringVarP(&serverArg, "server", "s", "127.0.0.1", "server address")
	clientCmd.Flags().IntVarP(&portArg, "port", "p", 0, "listen port")
	clientCmd.Flags().StringVarP(&protocolArg, "protocol", "P", "tcp", "listen protocol, (tcp, udp)")

	return clientCmd
}

var (
	serverArg   string
	portArg     int
	protocolArg string
)

func run(cmd *cobra.Command, args []string) {
	connect(serverArg, portArg, protocolArg)
}

func perRun(cmd *cobra.Command, args []string) error {
	// bindArg
	if !utils.IsValidIp(serverArg) {
		return fmt.Errorf("server address:%s is invalid", serverArg)
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
