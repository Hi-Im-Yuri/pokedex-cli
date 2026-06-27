package pokeapi

import (
	"encoding/json"
	"net/http"
)

// entry point url for pokeapi areas
const pokeApi string = "https://pokeapi.co/api/v2/location-area/"

// gets the json data of a location-area from the pokeApi
func GetArea(url string) (PokeMap, error) {
	res, err := http.Get(url)
	if err != nil {
		return PokeMap{}, err
	}
	defer res.Body.Close()

	var pokeLocations PokeMap
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&pokeLocations)
	if err != nil {
		return PokeMap{}, err
	}

	return pokeLocations, nil

}
