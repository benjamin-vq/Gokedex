package main

import "os"

func exitCommand(config *Config) error {
	os.Exit(0)
	return nil
}
