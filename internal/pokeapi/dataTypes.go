package pokeapi

import (
	"github.com/Hi-Im-Yuri/pokedex-cli/internal/pokecache"
)

type PokeMap struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		LocationName string `json:"name"`
		URL          string `json:"url"`
	} `json:"results"`
}

type Config struct {
	Next     *string
	Previous *string
	Cache    *pokecache.Cache
}

type PokeEncounters struct {
	PokemonList []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
