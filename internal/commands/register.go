package commands

import "github.com/spf13/cobra"

func RegisterCommands(root *cobra.Command) {
	root.AddCommand(
		newTapCmd(),
		newDoubleTapCmd(),
		newLongPressCmd(),
		newSwipeCmd(),
		newBackCmd(),
		newHomeCmd(),
		newLaunchAppCmd(),
		newInputCmd(),
		newClearTextCmd(),
		newCurrentAppCmd(),
		newDumpHierarchyCmd(),
		newScreenshotCmd(),
		newDeviceInfoCmd(),
		newListDevicesCmd(),
	)
}
