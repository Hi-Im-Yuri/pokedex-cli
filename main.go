package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	//cli tool entrypoint
	for {
		fmt.Print("Pokedex > ")

		//get user input
		if scanner.Scan() {
			rawInput := scanner.Text()
			userInputSlice := cleanInput(rawInput)
			//guard against user returning nothing
			if len(userInputSlice) == 0 {
				continue
			}
			userInput := userInputSlice[0]

			//check user input against commandMap
			command, ok := getCommands()[userInput]
			if ok {
				err := command.callback()
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
