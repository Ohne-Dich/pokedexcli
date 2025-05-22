package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Ohne-Dich/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemon    map[string]pokeapi.PokemonFull
}

func startRepl(cfg *config) {
	commandHelp(cfg, "hi")

	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		commandInput := []string{}

		if len(words) > 1 {
			commandInput = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, commandInput...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explores the Pokemon at location <location_name>",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Try to catch the Pokemon <name>",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon>",
			description: "Inspect a caught Pokemon <name>",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists the Pokemon you caught so far",
			callback:    commandPokedex,
		},
	}
}
