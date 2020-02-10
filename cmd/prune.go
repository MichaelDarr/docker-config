package cmd

import (
	"github.com/spf13/cobra"

	"github.com/MichaelDarr/docker-config/internal"
)

var volumes bool

var pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Remove unused Docker assets",
	Run: func(cmd *cobra.Command, args []string) {
		pruneArgs := []string{"system", "prune", "-a"}
		if volumes {
			pruneArgs = append(pruneArgs, "--volumes")
		}
		err := internal.DockerCmd(&pruneArgs)
		internal.PrintErrFatal(err)
	},
}

func init() {
	pruneCmd.Flags().BoolVar(&volumes, "volumes", false, "Also prune docker volumes")
	rootCmd.AddCommand(pruneCmd)
}