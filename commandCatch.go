package main

import (
	"fmt"
	"math/rand"

	"github.com/kevinjimenez96/pokedex/internal/pokeapi"
)

func commandCatch(params ...string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", params[0])
	pokemon, err := pokeapi.GetPokemon(params[0])

	if err != nil {
		return err
	}
	catchChance := rand.Intn(500)
	if catchChance < pokemon.BaseExperience {
		fmt.Printf("%s escaped!\n", params[0])
	} else {
		fmt.Printf("%s was caught!\nYou may now inspect it with the inspect command.\n", params[0])
		upokedex[params[0]] = pokemon
	}

	return nil
}
