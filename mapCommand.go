package main

import (
	"fmt"

	"github.com/benjamin-vq/gokedex/internal/gokeapi"
)

func mapNextCommand() error {

	locs, err := gokeapi.GetLocations()

	if err != nil {
		fmt.Println("Could not get locations, try again later")
		return err
	}

	for _, loc := range locs.Results {
		fmt.Printf("%s\n", loc.Name)
	}

	return nil
}

func mapPreviousCommand() error {
	locs, err := gokeapi.GetPreviousLocations()

	if err != nil {
		fmt.Println("Could not get locations, try again later")
		return err
	}

	for _, loc := range locs.Results {
		fmt.Printf("%s\n", loc.Name)
	}

	return nil
}
