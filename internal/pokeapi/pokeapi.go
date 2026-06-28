package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Hi-Im-Yuri/pokedex-cli/internal/pokecache"
)

// entry point url for pokeapi areas
const pokeApi string = "https://pokeapi.co/api/v2/location-area/"

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

// allows user to explore an area shown in map taking the userInput string and returning a struct
// containing a list of available pokemon
func GetExplore(userInput string, cache *pokecache.Cache) (PokeEncounters, error) {
	explorationUrl := pokeApi + userInput
	var encounters PokeEncounters
	data, cached := cache.Get(explorationUrl)
	if cached {
		err := json.Unmarshal(data, &encounters)
		if err != nil {
			return PokeEncounters{}, err
		}
		return encounters, nil
	}
	res, err := http.Get(explorationUrl)
	if err != nil {
		return PokeEncounters{}, err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return PokeEncounters{}, err
	}
	cache.Add(explorationUrl, resBody)
	err = json.Unmarshal(resBody, &encounters)
	if err != nil {
		return PokeEncounters{}, err
	}
	return encounters, nil
}
