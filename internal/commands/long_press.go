package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func newLongPressCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "long-press <x>,<y>",
		Short: "Long press on the device screen",
		Long:  "Perform a long press at the specified coordinates on the device screen.\nCoordinates are specified as x,y (e.g. 100,200).",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			p, err := parsePoint(args[0])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
				os.Exit(1)
			}
			client := mustCreateClient()
			printResult(client.LongPress(GetSerial(), p))
		},
	}
}
