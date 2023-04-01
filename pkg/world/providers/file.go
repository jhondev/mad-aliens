package providers

import (
	"bufio"
	"mad-aliens/pkg/world"
	"os"
	"strings"
)

type FileProvider struct {
	path string
}

func NewFile(path string) world.Provider {
	return &FileProvider{path: path}
}

func (fp *FileProvider) GetMap() (*world.Map, error) {
	file, err := os.Open(fp.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	wm := &world.Map{
		Cities:      make([]world.CityName, 0),
		Battlefield: make(world.Battlefield),
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		city := parts[0] // TODO: validation
		wm.Cities = append(wm.Cities, world.CityName(city))
		dirs := make(world.Directions)
		for _, part := range parts[1:] {
			dirParts := strings.Split(part, "=")
			dir := dirParts[0] // TODO: validation
			neighbor := dirParts[1]
			dirs[world.CityName(neighbor)] = world.Dir(dir)
		}
		wm.Battlefield[world.CityName(city)] = world.Status{
			Directions: dirs,
			Aliens:     make(world.AliensStat),
		}
	}
	return wm, nil
}
