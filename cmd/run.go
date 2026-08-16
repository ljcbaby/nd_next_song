package cmd

import (
	"fmt"
	"os"

	"github.com/ljcbaby/nd_next_song/conf"
	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "Show next songs",
	Long:  `Show the next songs to play based on the configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := conf.LoadConfig()
		if err != nil {
			fmt.Printf(`Error loading config: %v`, err)
			os.Exit(1)
		}
		fmt.Printf(`%+v`, config)
		os.Exit(0)
	},
}

func init() {
	RootCmd.AddCommand(RunCmd)
}
