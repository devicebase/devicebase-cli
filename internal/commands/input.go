package commands

import "github.com/spf13/cobra"

func newInputCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "input <text>",
		Short: "Input text on the device",
		Long:  "Type the specified text into the currently focused text field on the device.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			client := mustCreateClient()
			printResult(client.InputText(GetSerial(), args[0]))
		},
	}
}
