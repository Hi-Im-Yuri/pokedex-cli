package pokeapi

import (
	"testing"
	"time"

	"github.com/Hi-Im-Yuri/pokedex-cli/internal/pokecache"
)

func TestGetArea(t *testing.T) {
	const apiUrl string = "https://pokeapi.co/api/v2/location-area/"
	config := Config{
		Next:     nil,
		Previous: nil,
		Cache:    pokecache.NewCache(time.Second * 5),
	}
	areas, err := GetArea(apiUrl, config.Cache)
	if err != nil {
		t.Fatalf("error: %v getting locations from api", err)
	}

	for _, location := range areas.Results {
		if len(location.LocationName) <= 0 {
			t.Errorf("length of location name was 0. Returned: %v\n", location)
		}
	}
}
