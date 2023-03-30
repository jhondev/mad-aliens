package simulation

import (
	"fmt"
	"mad-aliens/pkg/worldmap"
)

type Simulation struct {
	world  *worldmap.WorldMap
	aliens int
}

func New(world *worldmap.WorldMap, aliens int) *Simulation {
	return &Simulation{
		world:  world,
		aliens: aliens,
	}
}

func (sim *Simulation) Run() error {
	fmt.Printf("Starting alien simulation with the following parameters:\n\n"+
		"  aliens: %d\n\n", sim.aliens)
	fmt.Println("Simulation running")
	return nil
}
