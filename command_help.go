package main

import (
    "fmt"
)

func commandHelp(cfg *config, args ...string) error {
    if len(args) != 0 {
        fmt.Printf("Error: invalid option: %s. This command has no option\n", args[0])
        return nil
    }

    fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

    for _, cmd := range getCommand() {
        fmt.Printf("%s: %s\n", cmd.name, cmd.description)
    }
    fmt.Println()
    return nil
}
