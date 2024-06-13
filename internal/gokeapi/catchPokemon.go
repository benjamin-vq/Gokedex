package gokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"math/rand"
	"net/http"
	"time"
)

// OOP-brained
var (
	pokemonCache = NewCache(120 * time.Second)
	pokedex      = NewPokedex()
)

func GetPokemonCache() *Cache {
	return pokemonCache
}

func GetPokedex() *Pokedex {
	return pokedex
}

func CatchPokemon(arg *string) (caught bool, name string, err error) {

	p, ok := pokedex.Entries[*arg]

	if ok {
		return true, p.Name, nil
	}

	pokemonUrl := apiBaseURL + "/pokemon/" + *arg

	pokemon, err := getPokemon(&pokemonUrl)

	if err != nil {
		return false, "", err
	}

	baseExp := pokemon.BaseExperience
	chance := catchChance(baseExp)

	if chance > .5 {
		pokedex.Entries[*arg] = pokemon
		caught = true
	}

	return caught, pokemon.Name, nil
}

func getPokemon(pokemonUrl *string) (*Pokemon, error) {

	if cached, hit := getPokemonFromCache(pokemonUrl); hit {
		return cached, nil
	}

	res, err := http.Get(*pokemonUrl)

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode > 299 || err != nil {
		return nil, errors.New("invalid api response")
	}

	pokemon := &Pokemon{}
	err = json.Unmarshal(body, pokemon)

	if err != nil {
		return nil, errors.New("error unmarshalling JSON")
	}

	pokemonCache.Add(*pokemonUrl, body)

	return pokemon, nil
}

func catchChance(baseExp int) float64 {

	f := rand.Float64()

	// Should be the maximum but I don't know how much that is
	maxBaseExp := 1000.0
	normalizedExp := (1 - float64(baseExp)/maxBaseExp)

	chance := normalizedExp * f
	return chance
}

func getPokemonFromCache(url *string) (p *Pokemon, hit bool) {

	entry, hit := pokemonCache.Get(*url)

	if !hit {
		return nil, false
	}

	p = &Pokemon{}
	err := json.Unmarshal(entry, p)

	// Should never happen, if it was cached it means it was previously
	// unmarshalled
	if err != nil {
		return nil, false
	}

	return p, true
}
