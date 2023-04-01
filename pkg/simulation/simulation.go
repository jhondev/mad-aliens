package simulation

import (
	"fmt"
	"mad-aliens/pkg/world"
)

type Simulation struct {
	world    *world.World
	naliens  int
	maxmoves int
	// randf returns, as an int, a non-negative pseudo-random number
	// in the half-open interval [0,n) from the default Source.
	// It panics if n <= 0
	randf world.RandF
}

func New(
	wld *world.World,
	naliens int,
	maxmoves int,
	randf world.RandF) *Simulation {
	return &Simulation{
		world:    wld,
		naliens:  naliens,
		maxmoves: maxmoves,
		randf:    randf,
	}
}

// Run runs the simulation
func (sim *Simulation) Run() error {
	fmt.Printf("Starting alien invasion simulation with the following setup:\n\n"+
		"  number of aliens: %d\n"+
		"  max moves:        %d\n\n", sim.naliens, sim.maxmoves)

	fmt.Println("👽 Loading world.")
	err := sim.world.Load(sim.naliens, sim.randf)
	if err != nil {
		return err
	}

	fmt.Printf("👽 Mad aliens are invading:\n\n")
	for i := 0; i < sim.maxmoves; i++ {
		if finished := sim.checkBattlefield(); finished {
			break
		}
		sim.moveAliens()
	}
	fmt.Println("👽 The mad aliens invasion has finished with the following report:")
	return nil
}

func (sim *Simulation) checkBattlefield() bool {
	fmt.Printf("   👾 Checking for battlefield status:\n\n")
	for cityName, status := range sim.world.Battlefield {
		if len(status.Aliens) < 2 {
			continue
		}

		fmt.Printf("      💣 A dead battle is happening in %s city\n", cityName)
		fmt.Printf("      💀 %s has been destroyed by aliens %s\n\n",
			cityName, world.JoinAliens(status.Aliens, ", "))

		status.Destroyed = true
		// save the historical data
		sim.world.IW[cityName] = status

		sim.cleanUpBattlefield(cityName)
	}

	// return true if all aliens were destroyed
	return len(sim.world.Aliens) == 0
}

func (sim *Simulation) cleanUpBattlefield(cityName world.CityName) {
	// remove paths to the city
	for _, st := range sim.world.Battlefield {
		delete(st.Directions, cityName)
	}

	// remove destroyed aliens in the city from the world
	for alien := range sim.world.Battlefield[cityName].Aliens {
		delete(sim.world.Aliens, alien)
	}

	// remove city from the battlefield
	delete(sim.world.Battlefield, cityName)
}

func (sim *Simulation) moveAliens() {
	for alien, cityName := range sim.world.Aliens {
		status, ok := sim.world.Battlefield[cityName]
		if !ok {
			continue
		}
		newCity := sim.randPath(status.Directions)
		sim.world.Aliens[alien] = newCity
		delete(status.Aliens, alien)
	}
	fmt.Printf("      🛸 Aliens have moved\n\n")
}

func (sim *Simulation) randPath(dirs world.Directions) world.CityName {
	randIdx := sim.randf(len(dirs) - 1)
	dirIdx := 0
	for cityName := range dirs {
		if dirIdx == randIdx {
			return cityName
		}
		dirIdx++
	}
	return ""
}
