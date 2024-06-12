package main

import (
	"log"

	"github.com/benjamin-vq/gokedex/internal/gokeapi"
)

func cacheCommand(config *Config, _ string) error {

	locs := gokeapi.GetLocationsCache()
	explore := gokeapi.GetExploreAreasCache()

	log.Println("--------- Current Keys in Location Cache --------- ")
	for key := range locs.Cache {
		log.Println(key)

	}
	log.Println("--------- Current Keys in Location Cache --------- ")

	log.Println("--------- Current Keys in Explore Cache --------- ")
	for key := range explore.Cache {
		log.Println(key)

	}
	log.Println("--------- Current Keys in Explore Cache --------- ")
	return nil
}
