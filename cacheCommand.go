package main

import (
	"log"

	"github.com/benjamin-vq/gokedex/internal/gokeapi"
)

func cacheCommand() error {

	cache := gokeapi.GetLocationsCache()

	log.Println("--------- Current Keys in Location Cache --------- ")
	for key := range cache.Cache {
		log.Printf("%s\n", key)

	}
	log.Println("--------- Current Keys in Location Cache --------- ")
	return nil
}
