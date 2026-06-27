package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Hi-Im-Yuri/pokedex-cli/internal"
	"github.com/Hi-Im-Yuri/pokedex-cli/internal/pokeapi"
	"github.com/Hi-Im-Yuri/pokedex-cli/internal/repl"
)

// base api url
const pokeApi string = "https://pokeapi.co/api/v2/location-area/"

func main() {
	//initialize cli tool
	scanner := bufio.NewScanner(os.Stdin)

	//create struct to manage calls to different areas
	var config pokeapi.Config

	//cli tool entrypoint
	for {
		fmt.Print("Pokedex > ")

		//get user input
		if scanner.Scan() {
			rawInput := scanner.Text()
			userInputSlice := repl.CleanInput(rawInput)
			//guard against user returning nothing
			if len(userInputSlice) == 0 {
				continue
			}
			userInput := userInputSlice[0]

			//check user input against commandMap
			command, ok := internal.GetCommands()[userInput]
			if ok {
				err := command.Callback(&config)
				if err != nil {
					fmt.Printf("Error: %v\n", err)
				}
			} else {
				fmt.Println("Unknown command")
			}

			//checks for any error with the scanner itself
		} else if err := scanner.Err(); err != nil {
			fmt.Printf("There was an error: %v with the scanner. Shutting down safely...", err)
			return
		}
	}
}
