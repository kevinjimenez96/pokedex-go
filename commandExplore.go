package main

import (
	"fmt"

	"github.com/kevinjimenez96/pokedex/internal/pokeapi"
)

func commandExplore(params ...string) error {
	fmt.Printf("Exploring %s...\n", params[0])
	locations, err := pokeapi.GetLocation(params[0])

	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")

	for _, pokemonEncounter := range locations.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemonEncounter.Pokemon.Name)
	}

	return nil
}
