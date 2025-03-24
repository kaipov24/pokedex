package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kaipov24/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func keyContains(m map[string]cliCommand, substr string) bool {
	for key := range m {
		if strings.Contains(key, substr) {
			return true
		}
	}
	return false
}

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	mapCliCommands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Show this help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display the names of 20 location areas",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the names of 20 location areas in reverse",
			callback:    commandMapb,
		},
	}

	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		command := reader.Text()
		if keyContains(mapCliCommands, command) {
			mapCliCommands[command].callback(cfg)

		} else {
			fmt.Println("Unknown command")
		}

	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)

	return words
}
