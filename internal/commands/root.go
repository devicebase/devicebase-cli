package commands

import (
	"fmt"
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
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if serial == "" {
				return fmt.Errorf("required flag(s) \"--serial\" not set")
			}
			return nil
		},
	}

	rootCmd.PersistentFlags().StringVarP(&serial, "serial", "s", "", "Device serial number (required)")
	rootCmd.MarkPersistentFlagRequired("serial")

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
