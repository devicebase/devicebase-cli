package commands

import (
	"fmt"
	"os"

	"github.com/uusense/devicebase-cli/internal/api"
)

func mustCreateClient() *api.Client {
	c, err := api.NewClient()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
	return c
}

func printResult(data []byte, err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
	if len(data) > 0 {
		fmt.Println(string(data))
	}
}
