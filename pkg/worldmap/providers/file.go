package providers

import "mad-aliens/pkg/worldmap"

type FileProvider struct {
	name string
}

func NewFile(path string) worldmap.Provider {
	return &FileProvider{}
}
