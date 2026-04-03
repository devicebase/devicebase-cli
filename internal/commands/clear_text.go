package commands

import "github.com/spf13/cobra"

func newClearTextCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "clear-text",
		Short: "Clear text in the current input field",
		Long:  "Clear all text content in the currently focused text field on the device.",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			client := mustCreateClient()
			printResult(client.ClearText(GetSerial()))
		},
	}
}
