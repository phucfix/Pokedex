package main

import (
    "strings"
    "bufio"
    "fmt"
    "os"
)

func cleanInput(text string) []string {
    output := strings.ToLower(text)
    words := strings.Fields(output)
    return words
}

func startRepl() {
    scanner := bufio.NewScanner(os.Stdin)

    for {
        // Print the prompt
        fmt.Print("Pokedex > ")

        // Wait for user input
        scanner.Scan()

        // Get input that user typed (exclude '\n' char)
        userInput := scanner.Text()
        userCommand := strings.Fields(userInput)[0]

        command, exists := getCommand()[userCommand]        
        if !exists {
            fmt.Printf("Unknown command: %s\n", userCommand)
            continue
        }
        if err := command.callback(); err != nil {
            fmt.Printf("%s: %v\n", command.name, err)
            continue
        }
        if command.name == "exit" {
            os.Exit(0)
        }
    }
}
