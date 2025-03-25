package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/kaipov24/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *config, pkx *pokedex, args ...string) error {
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
		if pkx.pokemons == nil {
			pkx.pokemons = make(map[string]pokeapi.Pokemon)
		}
		pkx.pokemons[name] = pokemon
		fmt.Printf("%v has been added to your Pokedex!\n", name)
		fmt.Printf("Pokemons in pokedex %v !\n", len(pkx.pokemons))
		return nil
	} else {
		fmt.Printf("%v escaped!\n", name)
		return nil
	}

}
