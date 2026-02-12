package main

import (
	"time"
	"github.com/goodoletony/pokedexcli/internal/pokeapi"
)
type config struct {
	pokeapiClient pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
	pokemonSpecimens	map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Minute * 30),
		pokemonSpecimens: make(map[string]pokeapi.Pokemon),
	}
	startRepl(&cfg)
}
