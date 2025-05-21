package main

import (
	"encoding/json"
	"errors"
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

	if cfg.next != nil {
		url = *cfg.next
	}

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
	if cfg.previous == nil {
		return errors.New("you're on the first page")
	}
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	url := *cfg.previous

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
