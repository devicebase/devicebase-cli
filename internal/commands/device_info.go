package commands

import "github.com/spf13/cobra"

func newDeviceInfoCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "device-info",
		Short: "Get device information",
		Long:  "Retrieve detailed information about the device including status, hardware info, and connection state.",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			client := mustCreateClient()
			printResult(client.DeviceInfo(GetSerial()))
		},
	}
}
