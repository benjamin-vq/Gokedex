package main

import (
	"errors"
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

	configCommand(config)

	return nil
}

func mapPreviousCommand(config *Config) error {

	if config.previousUrl == nil {
		// We just map'b to the beginning, reset nextUrl otherwise it
		// may point to the second batch of locations
		// Example: Without this line, here is what happened:
		// launch the program -> map -> mapb -> nextUrl now points to the second batch
		config.nextUrl = nil
		return errors.New("there are no previous locations")
	}

	locs, err := gokeapi.GetLocations(config.previousUrl)

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
