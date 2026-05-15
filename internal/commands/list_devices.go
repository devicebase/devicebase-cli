package commands

import (
	"github.com/spf13/cobra"
	"github.com/uusense/devicebase-cli/internal/api"
)

func newListDevicesCmd() *cobra.Command {
	var keyword, state string
	var limit int

	cmd := &cobra.Command{
		Use:   "list-devices",
		Short: "List all devices",
		Long:  "Retrieve a list of all devices accessible by the current user. Supports filtering by keyword and state.",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			client := mustCreateClient()

			req := api.ListDevicesRequest{}
			if keyword != "" {
				req.Keyword = &keyword
			}
			if state != "" {
				req.State = &state
			}
			if limit > 0 {
				req.Limit = &limit
			}

			printResult(client.ListDevices(req))
		},
	}

	cmd.Flags().StringVar(&keyword, "keyword", "", "Filter by keyword (brand/model/serial/name)")
	cmd.Flags().StringVar(&state, "state", "", "Filter by state (busy/free/offline)")
	cmd.Flags().IntVar(&limit, "limit", 10, "Maximum number of devices to return")

	return cmd
}
