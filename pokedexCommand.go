package main

import (
	"fmt"

	"github.com/benjamin-vq/gokedex/internal/gokeapi"
)

func pokedexCommand(_ *Config, _ string) error {

	pokedex := gokeapi.GetPokedex()

	fmt.Printf("Your pokedex contains %d entries:\n", len(pokedex.Entries))
	for _, entry := range pokedex.Entries {
		fmt.Printf(" - %s\n", entry.Name)
	}

	return nil
}
