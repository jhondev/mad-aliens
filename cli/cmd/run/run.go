package run

import (
	"fmt"
	"mad-aliens/pkg/simulation"
	"mad-aliens/pkg/world"
	"mad-aliens/pkg/world/providers"
	"math/rand"

	"github.com/spf13/cobra"
)

const (
	flagWorldFilePath = "path"
	flagAliensNumber  = "naliens"
	flagMaxMoves      = "maxmoves"
	defaultAliens     = 5
	defaultMaxMovs    = 10000
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
			naliens, _ := flags.GetInt(flagAliensNumber)
			maxmoves, _ := flags.GetInt(flagMaxMoves)
			err := runSimulation(path, naliens, maxmoves)
			if err != nil {
				fmt.Printf("error running the simulation: %v", err)
			}
		},
	}
	runCmd.Flags().StringP(flagWorldFilePath, "p", "", "world map file path")
	_ = runCmd.MarkFlagRequired(flagWorldFilePath)
	runCmd.Flags().IntP(flagAliensNumber, "n", defaultAliens, "number of aliens in the simulation")
	runCmd.Flags().IntP(flagMaxMoves, "m", defaultMaxMovs, "max number of moves")

	return runCmd
}

func runSimulation(path string, naliens int, maxmoves int) error {
	prov := providers.NewFile(path)
	wld := world.New(prov)
	sim := simulation.New(wld, naliens, maxmoves, rand.Intn)
	return sim.Run()
}
