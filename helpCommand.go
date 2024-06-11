package main

import "fmt"

func helpCommand(cpnfig *Config) error {

	commands := getCliCommands()

	fmt.Println()
	fmt.Println("Welcome to the Gokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()

	return nil
}
