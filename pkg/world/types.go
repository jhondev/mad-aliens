package world

type CityName string
type Dir string
type Alien string
type Directions = map[CityName]Dir
type AliensStat = map[Alien]bool
type Status struct {
	Directions
	Aliens    AliensStat
	Destroyed bool
}
type Battlefield = map[CityName]Status
type Aliens = map[Alien]CityName
type City struct {
	Name CityName
}
type Map struct {
	Cities []CityName
	Battlefield
}
type World struct {
	provider Provider
	*Map
	Aliens
	IW Battlefield // Information Warfare
}
type Provider interface {
	GetMap() (*Map, error)
}
type RandF func(int) int
