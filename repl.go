package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Khazz0r/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		inputted_text := scanner.Text()
		cleaned_input := cleanInput(inputted_text)

		if len(cleaned_input) == 0 {
			continue
		}

		if command, ok := getCommandRegistry()[cleaned_input[0]]; ok {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	textLower := strings.ToLower(text)
	cleanedStrSlice := strings.Fields(textLower)

	return cleanedStrSlice
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommandRegistry() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name: "map",
			description: "Explore the pokemon map forward by 20 locations",
			callback: commandMapf,
		},
		"mapb": {
			name: "mapb",
			description: "Explore the pokemon map backward by 20 locations",
			callback: commandMapb,
		},
	}
}
