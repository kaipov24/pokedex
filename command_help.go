package main

import "fmt"

func commandHelp(cfg *config, pkx *pokedex, args ...string) error {
	fmt.Println("Displays a help message")
	return nil
}
