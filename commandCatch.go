package main

import (
	"fmt"

	"github.com/kevinjimenez96/pokedex/internal/pokeapi"
)

func commandCatch(params ...string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", params[0])
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
