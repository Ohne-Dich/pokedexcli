package main

import (
	"strings"
)

func cleanInput(text string) []string {
	x := []string{}
	for _, word := range strings.Fields(text) {
		x = append(x, strings.ToLower(word))
	}
	return x
}
