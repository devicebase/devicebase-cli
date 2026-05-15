package commands

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/uusense/devicebase-cli/internal/api"
)

func newTapCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "tap <x>,<y>",
		Short: "Tap on the device screen",
		Long:  "Perform a single tap at the specified coordinates on the device screen.\nCoordinates are specified as x,y (e.g. 100,200).",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			p, err := parsePoint(args[0])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
				os.Exit(1)
			}
			client := mustCreateClient()
			printResult(client.Tap(mustGetSerial(), p))
		},
	}
}

func parsePoint(s string) (api.Point, error) {
	parts := strings.SplitN(s, ",", 2)
	if len(parts) != 2 {
		return api.Point{}, fmt.Errorf("invalid point format %q, expected x,y", s)
	}
	x, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return api.Point{}, fmt.Errorf("invalid x coordinate: %w", err)
	}
	y, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return api.Point{}, fmt.Errorf("invalid y coordinate: %w", err)
	}
	return api.Point{X: x, Y: y}, nil
}
