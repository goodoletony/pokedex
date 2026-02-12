package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]

		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue
		}
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

type cliCommand struct {
	name		string
	description	string
	callback	func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"catch": {
			name:		 "catch {pokemon_name}",
			description: "Try to catch a Pokemon and add it to your Pokedex",
			callback:	 commandCatch,
		},
		"exit": {
			name:		 "exit",
			description: "Exit the Pokedex",
			callback:	 commandExit,
		},
		"explore": {
			name:		 "explore {location_area}",
			description: "Lists Pokemon in Location Area",
			callback:	 commandExplore,
		},
		"help": {
			name:		 "help",
			description: "Displays a help message",
			callback:	 commandHelp,
		},
		"inspect": {
			name:		 "inspect {pokemon_name}",
			description: "inspects a Pokemon in Pokedex (if caught)",
			callback:	 commandInspect,
		},
		"map": {
			name:		 "map",
			description: "Gets location areas forward",
			callback:	 commandMap,
		},
		"mapb": {
			name:		 "map",
			description: "Gets location areas backward",
			callback:	 commandMapB,
		},
		"pokedex": {
			name:		 "pokedex",
			description: "Displays Pokedex contents",
			callback:	 commandPokedex,
		},
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
