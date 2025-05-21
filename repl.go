package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		txt := cleanInput(scanner.Text())
		cmdList := registry()
		function, ok := cmdList[txt[0]]
		if ok {
			err := function.callback(cfg)
			if err != nil {
				fmt.Printf("oh no, something went wrong: %v\n", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	x := []string{}
	for _, word := range strings.Fields(text) {
		x = append(x, strings.ToLower(word))
	}
	return x
}
