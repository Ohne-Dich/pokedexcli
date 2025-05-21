package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	next     *string
	previous *string
}

func registry() map[string]cliCommand {
	x := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Helps with the Pokedex",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Gives you the next 20 Map Points",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Gives you the previous 20 Map Points",
			callback:    commandMapb,
		},
	}
	return x
}
