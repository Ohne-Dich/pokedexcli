package main

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
	return x
}
