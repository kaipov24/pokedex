package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) < 1 {
		fmt.Println("You have not caught any pokemon yet")
		return nil

	} else {
		fmt.Println("Your Pokedex:")
		for pokemon := range cfg.caughtPokemon {
			fmt.Println("  -", pokemon)
		}
	}

	return nil
}
