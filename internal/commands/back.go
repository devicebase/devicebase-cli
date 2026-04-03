package commands

import "github.com/spf13/cobra"

func newBackCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "back",
		Short: "Press the back button",
		Long:  "Simulate pressing the back button on the device.",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			client := mustCreateClient()
			printResult(client.Back(GetSerial()))
		},
	}
}
