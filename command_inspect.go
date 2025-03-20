package main

import (
    "errors"
    "fmt"
)

func commandInspect(cfg *config, args ...string) error {
    if len(args) != 1 {
        return errors.New("You must provide a pokemon name")
    }

    pokemonName := args[0]

    pokemonData, ok := cfg.caughtPokemon[pokemonName];
    if !ok {
        fmt.Println("you have not caught that pokemon")
        return nil
    }

    fmt.Printf("Name: %s\n", pokemonName)
    fmt.Printf("Height: %d\n", pokemonData.Height)
    fmt.Printf("Weight: %d\n", pokemonData.Weight)
    fmt.Printf("Stats:\n")
    for _, pokemonStat := range pokemonData.Stats {
        fmt.Printf("\t-%s: %d\n", pokemonStat.Stat.Name, pokemonStat.BaseStat)
    }
    fmt.Printf("Types:\n")
    for _, pokemonType := range pokemonData.Types {
        fmt.Printf("\t- %s\n", pokemonType.Type.Name)
    }

    return nil
}
