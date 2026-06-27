package internal

import (
	"fmt"
	"os"
)

type CliCommand struct {
	name        string
	description string
	Callback    func() error
}

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
	}

	return commandMap
}

// exit the cli tool
func CommandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// display a list of commands available to the user
func CommandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, cli := range GetCommands() {
		fmt.Printf("%s: %s\n", cli.name, cli.description)
	}
	return nil
}
