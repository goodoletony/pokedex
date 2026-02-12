package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	_ = cfg
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	command := getCommands()
	for k, v := range command {
		fmt.Printf("%v: %v\n", k, v.description)
	}
	return nil
}
