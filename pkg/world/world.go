package world

type City string
type Dir string
type Alien int
type Direction = map[Dir]City
type GPS = map[City]Direction
type Aliens = map[Alien]City
type Map struct {
	Cities []City
	GPS
}
type World struct {
	provider Provider
	*Map
	Aliens
}
type Provider interface {
	GetMap() (*Map, error)
}
type RandF func(int) int

func New(prov Provider) *World {
	return &World{provider: prov}
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
		wld.Aliens[Alien(i)] = city
	}

	return nil
}
