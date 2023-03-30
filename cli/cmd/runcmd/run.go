package runcmd

import (
	"fmt"
	"mad-aliens/pkg/simulation"
	"mad-aliens/pkg/worldmap"
	"mad-aliens/pkg/worldmap/providers"

	"github.com/spf13/cobra"
)

const (
	flagWorldFilePath = "path"
	flagAliensNumber  = "aliens"
	defaultAliens     = 5
)

func New() *cobra.Command {
	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Run the aliens invasion simulation",
		Long:  ``,
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			flags := cmd.Flags()
			path, _ := flags.GetString(flagWorldFilePath)
			aliens, _ := flags.GetInt(flagAliensNumber)
			err := runSimulation(path, aliens)
			if err != nil {
				fmt.Printf("error running the simulation: %v", err)
			}
		},
	}
	runCmd.Flags().StringP(flagWorldFilePath, "p", "", "world map file path")
	_ = runCmd.MarkFlagRequired(flagWorldFilePath)
	runCmd.Flags().IntP(flagAliensNumber, "a", defaultAliens, "number of aliens in the simulation")

	return runCmd
}

func runSimulation(path string, aliens int) error {
	prov := providers.NewFile(path)
	world := worldmap.New(prov)
	sim := simulation.New(world, aliens)
	return sim.Run()
}
