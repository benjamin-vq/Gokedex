package gokeapi

type Pokedex struct {
	Entries map[string]*Pokemon
}

func NewPokedex() *Pokedex {

	pokedex := Pokedex{
		map[string]*Pokemon{},
	}
	return &pokedex
}
