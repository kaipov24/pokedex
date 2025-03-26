# Pokedex CLI

Pokedex CLI is a command-line REPL application that allows users to explore Pokémon locations, catch Pokémon, and manage their personal Pokédex. It interacts with the [PokéAPI](https://pokeapi.co/) to fetch Pokémon and location data.

## Features

- **Explore Locations**: View Pokémon available in specific locations.
- **Catch Pokémon**: Attempt to catch Pokémon and add them to your personal Pokédex.
- **Inspect Pokémon**: View detailed stats and information about caught Pokémon.
- **Navigate Locations**: Browse through paginated lists of locations.
- **View Pokedex**: Display all Pokémon you have caught.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/kaipov24/pokedexcli.git
   cd pokedexcli
   ```

2. Build the project:
   ```bash
   go build -o pokedexcli
   ```

3. Run the application:
   ```bash
   ./pokedexcli
   ```