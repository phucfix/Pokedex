package main

import (
    "fmt"
)

func commandExit() error {
    fmt.Printf("Closing the Pokedex... Goodbye!\n")
    return nil
}

func commandHelp() error {
    fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")

    for k, v := range getCommand() {
        fmt.Printf("%s: %s\n", k, v.description)
    }
    return nil
}
