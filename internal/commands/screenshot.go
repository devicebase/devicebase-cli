package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var screenshotOutput string

func newScreenshotCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "screenshot",
		Short: "Take a screenshot of the device",
		Long:  "Capture a screenshot of the device screen. The image is saved to a file\nor printed to stdout if no output file is specified.",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			client := mustCreateClient()
			data, err := client.Screenshot(mustGetSerial())
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
				os.Exit(1)
			}

			if screenshotOutput != "" {
				if err := os.WriteFile(screenshotOutput, data, 0644); err != nil {
					fmt.Fprintln(os.Stderr, "Error writing file:", err)
					os.Exit(1)
				}
				fmt.Printf("Screenshot saved to %s\n", screenshotOutput)
			} else {
				os.Stdout.Write(data)
			}
		},
	}

	cmd.Flags().StringVarP(&screenshotOutput, "output", "o", "", "Output file path (default: stdout)")

	return cmd
}
