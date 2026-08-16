package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

var Version = "dev"

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show current version",
	Long:  `Show the current version of nd_next_song`,
	Run: func(cmd *cobra.Command, args []string) {
		goVersion := fmt.Sprintf("%s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH)

		fmt.Printf(`Version: %s
Go Version: %s
`, Version, goVersion)

		os.Exit(0)
	},
}

func init() {
	RootCmd.AddCommand(VersionCmd)
}
