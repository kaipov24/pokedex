package main

import (
	"fmt"
	"time"

	"github.com/kaipov24/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	fmt.Println(pokeClient)
	startRepl(cfg)
}
