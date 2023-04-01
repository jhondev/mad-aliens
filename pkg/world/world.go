package world

import "fmt"

func New(prov Provider) *World {
	return &World{provider: prov, IW: make(Battlefield)}
}

// Load loads the World using the map data from the provider and positions aliens in random locations
// naliens: number of aliens to be positioned in the map
// randf: random function to return a pseudo-random number to position an alien
func (wld *World) Load(naliens int, randf RandF) error {
	// load map from provider
	data, err := wld.provider.GetMap()
	if err != nil {
		return err
	}
	wld.Map = data

	// load aliens
	wld.Aliens = make(Aliens)
	for i := 1; i <= naliens; i++ {
		randIdx := randf(len(wld.Cities))
		city := wld.Cities[randIdx]
		wld.Aliens[Alien(fmt.Sprint(i))] = city

		status := wld.Battlefield[city]
		status.Aliens[Alien(fmt.Sprint(i))] = true
		wld.Battlefield[city] = status
	}

	return nil
}
