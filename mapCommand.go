package main

import (
	"fmt"

	"github.com/benjamin-vq/gokedex/internal/gokeapi"
)

func mapNextCommand(config *Config) error {

	locs, err := gokeapi.GetLocations(config.nextUrl)

	if err != nil {
		return err
	}

	config.previousUrl = locs.Previous
	config.nextUrl = locs.Next

	for _, loc := range locs.Results {
		fmt.Printf("%s\n", loc.Name)
	}

	return nil
}

func mapPreviousCommand(config *Config) error {
	locs, err := gokeapi.GetPreviousLocations(config.previousUrl)

	if err != nil {
		return err
	}

	config.previousUrl = locs.Previous
	config.nextUrl = locs.Next

	for _, loc := range locs.Results {
		fmt.Printf("%s\n", loc.Name)
	}

	return nil
}
