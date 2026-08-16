package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	DataDir string
	Debug   bool
	Dev     bool
)

var RootCmd = &cobra.Command{
	Use:   "nd_next_song",
	Short: "song select",
	Long:  `A song selection tool for choosing the next song to play`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().BoolVar(&Debug, "debug", false, "start with debug mode")
	RootCmd.PersistentFlags().BoolVar(&Dev, "dev", false, "start with dev mode")
}
