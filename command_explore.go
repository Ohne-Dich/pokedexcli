package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}
	locationsResp, err := cfg.pokeapiClient.GetLocations(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationsResp.Name)
	fmt.Println("Found Pokemon: ")

	for _, mon := range locationsResp.PokemonEncounters {
		fmt.Println(mon.Pokemon.Name)
	}
	return nil
}
