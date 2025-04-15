package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
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
			command.callback()
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
	callback    func() error
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
	}
}
