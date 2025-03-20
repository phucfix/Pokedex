package main

import (
    "fmt"
)

func commandPokedex(cfg *config, args ...string) error {
    if len(args) != 0 {
        fmt.Printf("Error: invalid option: %s. This command has no option\n", args[0])
        return nil
    }

    if len(cfg.caughtPokemon) == 0 {
        fmt.Println("You have nothing in your deck! Please catch some pokemon")
        return nil
    }

    fmt.Println("Your Pokedex:")
    for _, pokemon := range cfg.caughtPokemon {
        fmt.Printf("\t- %s\n", pokemon.Name)
    }

    return nil
}
