package main

import (
	"fmt"
)

func commandCatch(cfg *config, args ...string) error {
	fmt.Println("Throwing a Pokeball at <pokemon>...")
	return nil
}
