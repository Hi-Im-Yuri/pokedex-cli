package internal

import (
	"fmt"
	"os"

	"github.com/Hi-Im-Yuri/pokedex-cli/internal/pokeapi"
)

type CliCommand struct {
	name        string
	description string
	Callback    func(*pokeapi.Config) error
}

// base api url
const pokeApi string = "https://pokeapi.co/api/v2/location-area/"

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
			description: "Displays a help message",
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
	}

	return commandMap
}

// exit the cli tool
func CommandExit(config *pokeapi.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// display a list of commands available to the user
func CommandHelp(config *pokeapi.Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, cli := range GetCommands() {
		fmt.Printf("%s: %s\n", cli.name, cli.description)
	}
	return nil
}

// displays the next possible pokemap
func CommandMap(config *pokeapi.Config) error {
	var url string
	if config.Next == nil {
		url = pokeApi
	} else {
		url = *config.Next
	}
	areas, err := pokeapi.GetArea(url)
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

func CommandMapB(config *pokeapi.Config) error {
	url := config.Previous
	if url == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	areas, err := pokeapi.GetArea(*url)
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
