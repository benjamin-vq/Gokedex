package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/benjamin-vq/gokedex/internal/gokeapi"
)

func initRepl() {

	config := Config{
		cache: *gokeapi.NewCache(120 * time.Second),
	}
	fmt.Print("Gokedex > ")
	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {

		cmd, arg := parseInput(scanner.Text())

		command, ok := getCliCommands()[cmd]

		if !ok {
			fmt.Println("Invalid command")
			continue
		}

		err := command.callback(&config, arg)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Print("Gokedex > ")
	}
}

type Config struct {
	cache                gokeapi.Cache
	previousUrl, nextUrl *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(config *Config, arg string) error
}

func parseInput(text string) (command string, arg string) {
	sliced := strings.Split(text, " ")

	if len(sliced) == 1 {
		return sliced[0], ""
	}

	return sliced[0], sliced[1]
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
		"explore": {
			name:        "explore [AREA]",
			description: "Displays more information about the provided area",
			callback:    exploreCommand,
		},
		"catch": {
			name:        "catch [NAME|ID]",
			description: "Attempt to catch a pokemon by providing a name or an id",
			callback:    catchCommand,
		},
		"inspect": {
			name:        "inspect [NAME]",
			description: "Display information about a pokemon in your pokedex",
			callback:    inspectCommand,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Display a list of the current pokemons in your pokedex",
			callback:    pokedexCommand,
		},
		// Extra commands, for debugging purposes
		"cached": {
			name:        "cached",
			description: "Display the current state of the GokeCache",
			callback:    cacheCommand,
		},
		"config": {
			name:        "config",
			description: "Display the current state of the previous and next URLs",
			callback:    configCommand,
		},
	}
}
