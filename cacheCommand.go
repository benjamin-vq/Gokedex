package main

import (
	"log"

	"github.com/benjamin-vq/gokedex/internal/gokeapi"
)

func cacheCommand(config *Config, _ string) error {

	locs := gokeapi.GetLocationsCache()
	explore := gokeapi.GetExploreAreasCache()
	pokemons := gokeapi.GetPokemonCache()

	log.Println("--------- Current Keys in Location Cache --------- ")
	for key := range locs.Cache {
		log.Println(key)

	}
	log.Println("--------- Current Keys in Location Cache --------- ")
	log.Println()
	log.Println("--------- Current Keys in Explore Cache --------- ")
	for key := range explore.Cache {
		log.Println(key)

	}
	log.Println("--------- Current Keys in Explore Cache --------- ")
	log.Println()
	log.Println("--------- Current Keys in Pokemon Cache --------- ")
	for key := range pokemons.Cache {
		log.Println(key)

	}
	log.Println("--------- Current Keys in Pokemon Cache --------- ")
	return nil
}
