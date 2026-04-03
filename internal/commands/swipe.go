package commands

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/uusense/devicebase-cli/internal/api"
)

func newSwipeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "swipe <x1>,<y1>,<x2>,<y2>",
		Short: "Swipe on the device screen",
		Long:  "Perform a swipe gesture from (x1,y1) to (x2,y2) on the device screen.\nCoordinates are specified as x1,y1,x2,y2 (e.g. 100,200,300,400).",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			b, err := parseBounds(args[0])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
				os.Exit(1)
			}
			client := mustCreateClient()
			printResult(client.Swipe(GetSerial(), b))
		},
	}
}

func parseBounds(s string) (api.Bounds, error) {
	parts := strings.SplitN(s, ",", 4)
	if len(parts) != 4 {
		return api.Bounds{}, fmt.Errorf("invalid bounds format %q, expected x1,y1,x2,y2", s)
	}
	vals := make([]int, 4)
	for i, p := range parts {
		v, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil {
			return api.Bounds{}, fmt.Errorf("invalid coordinate at position %d: %w", i, err)
		}
		vals[i] = v
	}
	return api.Bounds{X1: vals[0], Y1: vals[1], X2: vals[2], Y2: vals[3]}, nil
}
