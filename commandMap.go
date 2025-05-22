package main

import (
	"fmt"

	"github.com/kevinjimenez96/pokedex/internal/pokeapi"
)

func commandMap(params ...string) error {
	locations, err := pokeapi.GetNextLocations()

	if err != nil {
		return err
	}

	for _, location := range locations {
		fmt.Println(location.Name)
	}

	return nil
}
