package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a single target")
	}
	pokemon, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}
	name := pokemon.Name
	fmt.Printf("Throwing a Pokeball at %v...\n", name)
	chance := rand.Intn(256)

	if chance > pokemon.BaseExperience {
		fmt.Printf("%v was caught!\n", name)
		cfg.caughtPokemon[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%v escaped!\n", name)
	}

	return nil
}
