package main

import (
	"bufio"
	"fmt"
	"os"
)

func initRepl() {

	fmt.Print("Gokedex > ")
	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		input := scanner.Text()
		command, ok := getCliCommands()[input]

		if !ok {
			fmt.Println("Unknown command '" + input + "'")
			continue
		}

		command.callback()
		fmt.Print("Gokedex > ")
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCliCommands() map[string]cliCommand {

	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display a helpful message",
			callback:    helpCommand,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Gokedex",
			callback:    exitCommand,
		},
		"map": {
			name:        "map",
			description: "Displays 20 location areas",
			callback:    mapNextCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas",
			callback:    mapPreviousCommand,
		},
		"cached": {
			name:        "cached",
			description: "Display the current state of all caches",
			callback:    cacheCommand,
		},
	}
}
