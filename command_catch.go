package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("pokemon name not provided")
	}

	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	time.Sleep(time.Second * 2)

	const threshold = 50
	randomNum := rand.Intn(pokemon.BaseExperience)
	if randomNum > threshold {
		return fmt.Errorf("%s escaped!\n", pokemonName)
	}

	cfg.pokemonSpecimens[pokemonName] = pokemon
	fmt.Printf("%s was caught!\n", pokemonName)
	return nil
}
