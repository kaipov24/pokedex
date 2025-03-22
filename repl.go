package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Displays a help message")
	return nil
}

func keyContains(m map[string]cliCommand, substr string) bool {
	for key := range m {
		if strings.Contains(key, substr) {
			return true
		}
	}
	return false
}

func startRepl() {

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
	}

	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		command := reader.Text()
		if keyContains(mapCliCommands, command) {
			mapCliCommands[command].callback()

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
