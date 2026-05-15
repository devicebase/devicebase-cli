package commands

import (
	"os"

	"github.com/spf13/cobra"
)

var Version = "1.0.0"

var serial string

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "devicebase",
		Short:   "Devicebase - A CLI tool for device control via HTTP API",
		Long:    "Devicebase is a cross-platform CLI tool that interfaces with the Devicebase HTTP API\nto control remote devices. It supports tap, swipe, input text, launch apps,\nand many other device control operations.",
		Version: Version,
	}

	rootCmd.PersistentFlags().StringVarP(&serial, "serial", "s", "", "Device serial number")

	return rootCmd
}

func GetSerial() string {
	return serial
}

func Execute() {
	rootCmd := NewRootCmd()
	RegisterCommands(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
