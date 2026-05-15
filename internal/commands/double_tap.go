package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func newDoubleTapCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "double-tap <x>,<y>",
		Short: "Double tap on the device screen",
		Long:  "Perform a double tap at the specified coordinates on the device screen.\nCoordinates are specified as x,y (e.g. 100,200).",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			p, err := parsePoint(args[0])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
				os.Exit(1)
			}
			client := mustCreateClient()
			printResult(client.DoubleTap(mustGetSerial(), p))
		},
	}
}
