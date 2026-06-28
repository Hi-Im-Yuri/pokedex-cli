package pokeapi

import (
	"github.com/Hi-Im-Yuri/pokedex-cli/internal/pokecache"
)

// contains a list of location areas to be used for the map command
type PokeMap struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		LocationName string `json:"name"`
		URL          string `json:"url"`
	} `json:"results"`
}

// struct for important data that needs to be passed between main and commands
type Config struct {
	Next     *string
	Previous *string
	Cache    *pokecache.Cache
	Caught   map[string]PokeData
}

// contains a list of pokemon to be used with the explore command
type PokeEncounters struct {
	PokemonList []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

// contains information about an individual pokemon to be used with the catch and inspect commands
type PokeData struct {
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int
		Stat     struct {
			Name string
			URL  string
		}
	}
	Types []struct {
		Slot int
		Type struct {
			Name string
			URL  string
		}
	}
	Weight int `json:"weight"`
}
