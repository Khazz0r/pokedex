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
	caughtEntries map[string]pokeapi.Pokemon
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

		if command, exists := getCommandRegistry()[cleaned_input[0]]; exists {
			args := []string{}
			if len(cleaned_input) > 1 {
				args = cleaned_input[1:]
			}
			err := command.callback(cfg, args...)
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
	callback    func(*config, ...string) error
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
			description: "Traverse the pokemon map forward by 20 locations",
			callback: commandMapf,
		},
		"mapb": {
			name: "mapb",
			description: "Traverse the pokemon map backward by 20 locations",
			callback: commandMapb,
		},
		"explore": {
			name: "explore <location_name>",
			description: "Explore the pokemon in a location of your choice by its name",
			callback: commandExplore,
		},
		"catch": {
			name: "catch <pokemon_name>",
			description: "Catch a pokemon of your choice with a random chance based on its base experience",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect <pokemon_name>",
			description: "Inspect a pokemon that you've caught to see its stats",
			callback: commandInspect,
		},
	}
}
