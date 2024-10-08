package pkg

import (
	"os"

	"github.com/breeve/goperf/pkg/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func Run() {
	rootCmd := &cobra.Command{
		Use:   "goperf",
		Short: "TCP、UDP、SCTP Performance Testing Tools",
		Long:  `TCP、UDP、SCTP Performance Testing Tools`,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	rootCmd.SetVersionTemplate("v0.0.1")

	rootCmd.AddCommand(server.Server())

	if err := rootCmd.Execute(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}
