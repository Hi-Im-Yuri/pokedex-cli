package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/Hi-Im-Yuri/pokedex-cli/internal/pokecache"
)

// entry point url for pokeapi areas
const pokeApi string = "https://pokeapi.co/api/v2/location-area/"

// cache update interval
const interval time.Duration = 7 * time.Second

// gets the json data of a location-area from the pokeApi
func GetArea(url string, cache *pokecache.Cache) (PokeMap, error) {
	var pokeLocations PokeMap
	value, cached := cache.Get(url)
	if cached {
		err := json.Unmarshal(value, &pokeLocations)
		if err != nil {
			return PokeMap{}, err
		}
		return pokeLocations, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return PokeMap{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return PokeMap{}, err
	}
	cache.Add(url, body)
	err = json.Unmarshal(body, &pokeLocations)
	if err != nil {
		return PokeMap{}, err
	}

	return pokeLocations, nil

}
