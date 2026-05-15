package commands

import "github.com/spf13/cobra"

func newDumpHierarchyCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "dump-hierarchy",
		Short: "Dump the UI hierarchy",
		Long:  "Dump the current UI hierarchy (accessibility tree) of the device screen as a JSON structure.",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			client := mustCreateClient()
			printResult(client.DumpHierarchy(mustGetSerial()))
		},
	}
}
