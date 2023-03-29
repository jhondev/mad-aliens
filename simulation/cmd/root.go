package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "mad",
		Short: "Mad aliens helps you to simulate an alien invasion",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("mad aliens invasion simulation")
		},
	}
	return rootCmd
}
