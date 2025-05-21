package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	cmdList := registry()
	for _, list := range cmdList {
		fmt.Printf("%v: %v\n", list.name, list.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	url := "https://pokeapi.co/api/v2/location-area/"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	locations := ResJsonLocations{}
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return err
	}

	cfg.next = locations.Next
	cfg.previous = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {

	return nil
}
