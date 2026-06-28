package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/Hi-Im-Yuri/pokedex-cli/internal"
	"github.com/Hi-Im-Yuri/pokedex-cli/internal/pokeapi"
	"github.com/Hi-Im-Yuri/pokedex-cli/internal/pokecache"
	"github.com/Hi-Im-Yuri/pokedex-cli/internal/repl"
)

// base api url
const pokeApi string = "https://pokeapi.co/api/v2/location-area/"

func main() {
	//initialize cli tool
	scanner := bufio.NewScanner(os.Stdin)

	//create struct to manage calls to different areas
	var config pokeapi.Config
	config.Cache = pokecache.NewCache(time.Second * 7)

	//cli tool entrypoint
	for {
		fmt.Print("Pokedex > ")

		//get user input
		if scanner.Scan() {
			rawInput := scanner.Text()
			userInputSlice := repl.CleanInput(rawInput)
			//guard against user returning nothing
			var userFlag *string = nil
			if len(userInputSlice) == 0 {
				continue
			} else if len(userInputSlice) > 2 {
				fmt.Println("Too many arguments in user input")
				continue
			} else if len(userInputSlice) == 2 {
				userFlag = &userInputSlice[1]
			}
			userInput := userInputSlice[0]

			//check user input against commandMap
			command, ok := internal.GetCommands()[userInput]
			if ok {
				err := command.Callback(&config, userFlag)
				if err != nil {
					fmt.Printf("%v\n", err)
				}
			} else {
				fmt.Println("Unknown command")
			}

			//checks for any error with the scanner itself
		} else if err := scanner.Err(); err != nil {
			fmt.Printf("There was an error: '%v' with the scanner. Shutting down safely...", err)
			return
		}
	}
}
