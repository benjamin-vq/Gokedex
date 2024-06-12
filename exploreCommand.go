package main

import (
	"fmt"

	"github.com/benjamin-vq/gokedex/internal/gokeapi"
)

func exploreCommand(_ *Config, area string) error {

	if area == "" {
		return fmt.Errorf("explore area can not be empty")
	}

	fmt.Printf("Exploring %s...\n", area)
	exploreArea, err := gokeapi.GetExploreAreas(&area)

	if err != nil {
		return err
	}

	encounters := exploreArea.PokemonEncounters

	fmt.Println("Found Pokemon:")
	for _, encounter := range encounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
