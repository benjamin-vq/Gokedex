package main

import (
	"bufio"
	"fmt"
	"os"
)

func initRepl() {

	config := Config{}
	fmt.Print("Gokedex > ")
	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		input := scanner.Text()
		command, ok := getCliCommands()[input]

		if !ok {
			fmt.Println("Unknown command '" + input + "'")
			continue
		}

		err := command.callback(&config)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Print("Gokedex > ")
	}
}

type Config struct {
	previousUrl, nextUrl *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(config *Config) error
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
			description: "Display the current state of the GokeCache",
			callback:    cacheCommand,
		},
	}
}
