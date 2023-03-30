package simulation

import (
	"fmt"
	"mad-aliens/pkg/world"
)

type Simulation struct {
	world   *world.World
	naliens int
}

func New(wld *world.World, naliens int) *Simulation {
	return &Simulation{
		world:   wld,
		naliens: naliens,
	}
}

// Run runs the simulation
// randf: random function to return a pseudo-random number to position an alien
func (sim *Simulation) Run(randf world.RandF) error {
	fmt.Printf("Starting alien invasion simulation with the following parameters:\n\n"+
		"  number of aliens: %d\n\n", sim.naliens)

	fmt.Println("👽 Loading world")
	err := sim.world.Load(sim.naliens, randf)
	if err != nil {
		return err
	}

	fmt.Println("👽 Mad aliens are invading")

	return nil
}
