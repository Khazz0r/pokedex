package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		inputted_text := scanner.Text()
		cleaned_input := cleanInput(inputted_text)

		if len(cleaned_input) > 0 {
			fmt.Printf("Your command was: %v \n", cleaned_input[0])
		} else {
			fmt.Println("Invalid input, exiting program")
			break
		}
	}
}

func cleanInput(text string) []string {
	textLower := strings.ToLower(text)
	cleanedStrSlice := strings.Fields(textLower)

	return cleanedStrSlice
}
