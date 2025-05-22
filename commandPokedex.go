package main

import (
	"fmt"
)

func commandPokedex(params ...string) error {
	if len(upokedex) == 0 {
		fmt.Println("you have not caught any pokemon")
		return nil
	}

	fmt.Printf("Your Pokedex:\n")

	for name, _ := range upokedex {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}
