package main

import (
	"time"

	"github.com/Ohne-Dich/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		caughtPokemon: map[string]pokeapi.PokemonFull{},
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
