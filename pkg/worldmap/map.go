package worldmap

type Provider interface {
}

type WorldMap struct {
}

func New(prov Provider) *WorldMap {
	return &WorldMap{}
}
