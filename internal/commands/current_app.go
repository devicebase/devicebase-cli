package commands

import "github.com/spf13/cobra"

func newCurrentAppCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "current-app",
		Short: "Get the current foreground app",
		Long:  "Get the name of the application currently running in the foreground on the device.",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			client := mustCreateClient()
			printResult(client.CurrentApp(GetSerial()))
		},
	}
}
