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
        userCommand, parameter := parseCommandLine(userInput)
        if userCommand == "" {
            continue
        }

        command, exists := getCommand()[userCommand]        
        if !exists {
            fmt.Printf("Unknown command: %s\n", userCommand)
            continue
        }
        if err := command.callback(cfg, parameter); err != nil {
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

func parseCommandLine(userInput string) (string, string) {
    cleanUserInput := strings.TrimSpace(userInput)
    part := strings.SplitN(cleanUserInput, " ", 2)
    command := part[0]
    parameter := ""
    if len(part) != 1 {
        parameter = part[1]
    }

    return command, parameter
}

type cliCommand struct {
    name        string
    description string
    callback    func(*config, string) error
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
            description: "Displays the next 20 locations",
            callback:    commandMapb,
        },
        "explore": {
            name:        "explore",
            description: "Displays list of all the Pokemon located in there",
            callback:    commandExplore,
        },
    }
}
