package main

import "os"

func exitCommand(config *Config, _ string) error {
	os.Exit(0)
	return nil
}
