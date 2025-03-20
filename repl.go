package main

import (
    "strings"
    "bufio"
    "fmt"
    "os"
    "github.com/phucfix/pokedexcli/internal/pokeapi"
)

type config struct {
    pokeapiClient   pokeapi.Client
    nextLocationURL *string
    prevLocationURL *string
    caughtPokemon   map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
    scanner := bufio.NewScanner(os.Stdin)

    for {
        // Print the prompt
        fmt.Print("Pokedex > ")

        // Wait for user input
        scanner.Scan()

        // Get input that user typed (exclude '\n' char)
        userInput := scanner.Text()
        
        // Split the user input to command and parameter
        words := cleanInput(userInput)
        if len(words) == 0 {
            continue
        }

        commandName := words[0]
        args := []string{}
        if len(words) > 1 {
            args = words[1:]
        } 

        command, exists := getCommand()[commandName]        
        if !exists {
            fmt.Printf("Unknown command: %s\n", commandName)
            continue
        }
        if err := command.callback(cfg, args...); err != nil {
            fmt.Printf("%s: %v\n", command.name, err)
            continue
        }
    }
}

func cleanInput(text string) []string {
    output := strings.ToLower(text)
    words := strings.Fields(output)
    return words
}

type cliCommand struct {
    name        string
    description string
    callback    func(*config, ...string) error
}

func getCommand() map[string]cliCommand {
    return map[string]cliCommand {
        "exit": {
            name:        "exit",
            description: "Exit the Pokedex",
            callback:    commandExit,
        },
        "help": {
            name:        "help",
            description: "Displays a help message",
            callback:    commandHelp,
        },
        "map": {
            name:        "map",
            description: "Displays the names of 20 location areas in the Pokemon world",
            callback:    commandMapf,
        },
        "mapb": {
            name:        "mapb",
            description: "Displays the previous 20 locations",
            callback:    commandMapb,
        },
        "explore": {
            name:        "explore <location_name>",
            description: "Displays list of all the Pokemon located in there",
            callback:    commandExplore,
        },
        "catch": {
            name:        "catch <pokemon_name>",
            description: "Catch pokemon and adds them to the Pokedex",
            callback:    commandCatch,
        },
        "inspect": {
            name:        "inspect <caught_pokemon_name>",
            description: "See details about a Pokemon",
            callback:    commandInspect,
        },
        "pokedex": {
            name:        "pokedex",
            description: "Prints a list of all the names of the Pokemon the user has caught",
            callback:    commandPokedex,
        },
    }
}
