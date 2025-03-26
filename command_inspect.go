package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]
	pokemonInPokedex := cfg.caughtPokemon[name]
	if pokemonInPokedex.Name != "" {
		fmt.Printf("Name: %v\n", pokemonInPokedex.Name)
		fmt.Printf("Height: %v\n", pokemonInPokedex.Height)
		fmt.Printf("Weight: %v\n", pokemonInPokedex.Weight)
		fmt.Println("Stats:")
		for i := 0; i < len(pokemonInPokedex.Stats); i++ {
			fmt.Printf("  -%v: %v\n", pokemonInPokedex.Stats[i].Stat.Name, pokemonInPokedex.Stats[i].BaseStat)
		}
		fmt.Println("Types:")
		for i := 0; i < len(pokemonInPokedex.Types); i++ {
			fmt.Printf("  -%v\n", pokemonInPokedex.Types[i].Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}

	return nil
}
