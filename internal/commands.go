package internal

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/Hi-Im-Yuri/pokedex-cli/internal/pokeapi"
)

type CliCommand struct {
	name        string
	description string
	Callback    func(*pokeapi.Config, *string) error
}

// base api url for locations
const pokeApi string = "https://pokeapi.co/api/v2/location-area/"

// base api url for pokemon
const pokemonURL string = "https://pokeapi.co/api/v2/pokemon/"

// getCommands allows access to the commandMap map containing all commands
func GetCommands() map[string]CliCommand {
	//add any additional commands to the cli tool here. Define them below
	var commandMap = map[string]CliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message. Can be used with help <command> to display information about one command",
			Callback:    CommandHelp,
		},
		"map": {
			name:        "map",
			description: "Cycles through map locations",
			Callback:    CommandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "cycles backwards to previous location",
			Callback:    CommandMapB,
		},
		"explore": {
			name:        "explore",
			description: "explore an area using explore <location-name> from map",
			Callback:    CommandExplore,
		},
		"catch": {
			name:        "catch",
			description: "attempt to catch a pokemon using catch <pokemon-name>",
			Callback:    CommandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "display details about pokemon you have caught using inspect <pokemon-name>",
			Callback:    CommandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Access your pokedex! See all the pokemon you've caught so far.",
			Callback:    CommandPokedex,
		},
	}

	return commandMap
}

// exit the cli tool
func CommandExit(config *pokeapi.Config, userFlag *string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// display a list of commands available to the user
func CommandHelp(config *pokeapi.Config, userFlag *string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n")
	if userFlag != nil {
		command, exists := GetCommands()[*userFlag]
		if exists {
			fmt.Printf("%s: %s\n", command.name, command.description)
		} else {
			fmt.Printf("no command found for: %s", *userFlag)
		}
		return nil
	}
	for _, command := range GetCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

// displays the next possible pokemap
func CommandMap(config *pokeapi.Config, userFlag *string) error {
	var url string
	if config.Next == nil {
		url = pokeApi
	} else {
		url = *config.Next
	}
	areas, err := pokeapi.GetArea(url, config.Cache)
	if err != nil {
		return fmt.Errorf("Error: '%w' getting maps from api", err)
	}

	config.Next = areas.Next
	config.Previous = areas.Previous
	locationNames := []string{}

	for _, area := range areas.Results {
		locationNames = append(locationNames, area.LocationName)
	}

	for _, locationName := range locationNames {
		fmt.Println(locationName)
	}
	return nil

}

func CommandMapB(config *pokeapi.Config, userFlag *string) error {
	url := config.Previous
	if url == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	areas, err := pokeapi.GetArea(*url, config.Cache)
	if err != nil {
		return fmt.Errorf("Error: '%w' getting maps from api", err)
	}
	config.Next = areas.Next
	config.Previous = areas.Previous
	locationNames := []string{}

	for _, area := range areas.Results {
		locationNames = append(locationNames, area.LocationName)
	}

	for _, locationName := range locationNames {
		fmt.Println(locationName)
	}
	return nil

}

func CommandExplore(config *pokeapi.Config, userFlag *string) error {
	if userFlag == nil {
		fmt.Printf("please input explore <location-name>")
		return nil
	}
	encounters, err := pokeapi.GetExplore(*userFlag, config.Cache)
	if err != nil {
		return fmt.Errorf("Error: '%w' getting exploration data from api", err)
	}

	for _, encounter := range encounters.PokemonList {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}

func CommandCatch(config *pokeapi.Config, userFlag *string) error {
	if userFlag == nil {
		fmt.Printf("please input catch <pokemon-name>\n")
		return nil
	} else if pm, caught := config.Caught[*userFlag]; caught {
		fmt.Printf("You have already caught %s\n", pm.Name)
		return nil
	}
	pokemon, err := pokeapi.GetPokemon(*userFlag, config.Cache)
	if err != nil {
		return fmt.Errorf("Error: '%w' getting pokemon data from api", err)
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if rand.Int31n(350) >= int32(pokemon.BaseExperience) {
		config.Caught[pokemon.Name] = pokemon
		fmt.Printf("%s was caught!\n", pokemon.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}

func CommandInspect(config *pokeapi.Config, userFlag *string) error {
	if userFlag == nil {
		fmt.Printf("please input inspect <pokemon-name>\n")
		return nil
	}
	pokemon, caught := config.Caught[*userFlag]
	if !caught {
		fmt.Printf("You have not caught %s yet.", *userFlag)
		return nil
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pt := range pokemon.Types {
		fmt.Printf("  - %s\n", pt.Type.Name)
	}
	return nil
}

func CommandPokedex(config *pokeapi.Config, userFlag *string) error {
	if len(config.Caught) == 0 {
		fmt.Println("You have not caught any pokemon yet! Try using catch <pokemon-name> to catch 'em all!")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for pokemon, _ := range config.Caught {
		fmt.Printf("  - %s\n", pokemon)
	}
	return nil
}
