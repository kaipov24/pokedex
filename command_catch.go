package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	name := args[0]
	fmt.Printf("Throwing a Pokeball at %v...\n", name)

	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	chance := rand.Intn(pokemon.BaseExperience / 10)
	if chance == ((pokemon.BaseExperience / 10) - 1) {
		fmt.Printf("%v was caught!\n", name)
		return nil
	} else {
		fmt.Printf("%v escaped!\n", name)
		return nil
	}

}
