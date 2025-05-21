package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		txt := cleanInput(scanner.Text())
		msg := fmt.Sprintf("Your command was: %v", txt[0])
		fmt.Println(msg)
	}
}
