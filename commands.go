package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, command := range getCommandRegistry() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println("")
	return nil
}

func commandMapf(cfg *config, args ...string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}
	location, err := cfg.pokeapiClient.ListPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found these pokemon: ")
	for _, pokemonEncounter := range location.PokemonEncounters {
		fmt.Println(pokemonEncounter.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a valid pokemon name")
	}
	pokemon, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	catchChance := rand.Intn(pokemon.BaseExperience + 100)
	if catchChance < 100 {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.caughtEntries[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a valid pokemon name")
	}

	pokemonInfo, exists := cfg.caughtEntries[args[0]]
	if exists {
		fmt.Printf("Name: %s\n", pokemonInfo.Name)
		fmt.Printf("Height: %d\n", pokemonInfo.Height)
		fmt.Printf("Weight: %d\n", pokemonInfo.Weight)
		fmt.Println("Stats:")
		for i := 0; i < len(pokemonInfo.Stats); i++ {
			fmt.Printf("  -%s: %d\n", pokemonInfo.Stats[i].Stat.Name, pokemonInfo.Stats[i].BaseStat)
		}
		fmt.Println("Types:")
		for i := 0; i < len(pokemonInfo.Types); i++ {
			fmt.Printf("  - %s\n", pokemonInfo.Types[i].Type.Name)
		}
	} else {
		fmt.Printf("You have not caught %s yet, can't inspect\n", args[0])
	}

	return nil
}
