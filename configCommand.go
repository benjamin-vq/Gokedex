package main

import "fmt"

func configCommand(config *Config) error {

	var nextUrl string
	var previousUrl string

	if config.nextUrl != nil {
		nextUrl = *config.nextUrl
	}

	if config.previousUrl != nil {
		previousUrl = *config.previousUrl
	}

	fmt.Println("----------- Current Config State -----------")
	fmt.Printf("Next URL: %s\n", nextUrl)
	fmt.Printf("Previous URL: %s\n", previousUrl)
	fmt.Println("----------- Current Config State -----------")

	return nil
}
