package main

import (
	"errors"
	"fmt"

	"github.com/benjamin-vq/gokedex/internal/gokeapi"
)

func catchCommand(_ *Config, arg string) error {

	if arg == "" {
		return errors.New("catch argument can not be empty")
	}

	caught, name, err := gokeapi.CatchPokemon(&arg)

	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	if caught {
		fmt.Printf("%s was caught!\n", name)
	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	return nil
}
