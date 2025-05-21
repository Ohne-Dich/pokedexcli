package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		txt := cleanInput(scanner.Text())
		cmdList := registry()
		function, ok := cmdList[txt[0]]
		if ok {
			err := function.callback()
			if err != nil {
				fmt.Printf("oh no, something went wrong: %v\n", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	cmdList := registry()
	for _, list := range cmdList {
		fmt.Printf("%v: %v\n", list.name, list.description)
	}
	return nil
}

func cleanInput(text string) []string {
	x := []string{}
	for _, word := range strings.Fields(text) {
		x = append(x, strings.ToLower(word))
	}
	return x
}
