package cmd

import (
	"mad-aliens/cli/cmd/runcmd"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "mad",
		Short: "Mad aliens helps you to simulate an alien invasion",
		Long:  ``,
	}
	rootCmd.AddCommand(runcmd.New())

	return rootCmd
}
