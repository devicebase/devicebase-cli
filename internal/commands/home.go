package commands

import "github.com/spf13/cobra"

func newHomeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "home",
		Short: "Press the home button",
		Long:  "Simulate pressing the home button on the device to return to the home screen.",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			client := mustCreateClient()
			printResult(client.Home(GetSerial()))
		},
	}
}
