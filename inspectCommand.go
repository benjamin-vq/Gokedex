package main

import (
	"fmt"

	"github.com/benjamin-vq/gokedex/internal/gokeapi"
)

func inspectCommand(_ *Config, arg string) error {

	pokedex := gokeapi.GetPokedex()

	p, ok := pokedex.Entries[arg]

	if !ok {
		fmt.Println("You have not caught that pokemon yet")
		return nil
	}

	pokemon := *p
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stats := range pokemon.Stats {
		fmt.Printf(" -%s: %d\n", stats.Stat.Name, stats.BaseStat)
	}
	fmt.Println("Types:")
	for _, types := range pokemon.Types {
		fmt.Printf(" - %s\n", types.Type.Name)
	}

	return nil
}
