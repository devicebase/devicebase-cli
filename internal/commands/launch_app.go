package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func newLaunchAppCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "launch-app <app_name>",
		Short: "Launch an application on the device",
		Long:  "Launch the specified application on the device by its app name or package identifier.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if args[0] == "" {
				fmt.Fprintln(os.Stderr, "Error: app_name cannot be empty")
				os.Exit(1)
			}
			client := mustCreateClient()
			printResult(client.LaunchApp(GetSerial(), args[0]))
		},
	}
}
